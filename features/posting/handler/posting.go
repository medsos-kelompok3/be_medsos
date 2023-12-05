package handler

import (
	"be_medsos/features/posting"
	cld "be_medsos/utils/cld"
	"context"
	"net/http"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	golangjwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type PostingController struct {
	p      posting.Service
	cl     *cloudinary.Cloudinary
	ct     context.Context
	folder string
}

func New(p posting.Service, cld *cloudinary.Cloudinary, ctx context.Context, uploadparam string) posting.Handler {
	return &PostingController{
		p:      p,
		cl:     cld,
		ct:     ctx,
		folder: uploadparam,
	}
}

// Add implements posting.Handler.
func (pc *PostingController) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(PostingRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}
		formHeader, err := c.FormFile("gambar_posting")
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError, map[string]any{
					"message": "formheader error",
				})

		}
		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError, map[string]any{
					"message": "formfile error",
				})
		}

		link, err := cld.UploadImage(pc.cl, pc.ct, formFile, pc.folder)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "harap pilih gambar",
					"data":    nil,
				})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": "kesalahan pada server",
					"data":    nil,
				})
			}
		}
		var inputProcess = new(posting.Posting)
		inputProcess.GambarPosting = link
		inputProcess.Caption = input.Caption

		result, err := pc.p.AddPosting(c.Get("user").(*golangjwt.Token), *inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			return c.JSON(statusCode, map[string]any{
				"message": message,
			})
		}

		var response = new(PostingResponse)
		response.ID = result.ID
		response.Caption = result.Caption
		response.GambarPosting = result.GambarPosting
		response.UserName = result.UserName

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

package handler

import (
	"be_medsos/features/models"
	"be_medsos/features/posting"
	cld "be_medsos/utils/cld"
	"context"
	"net/http"
	"strconv"
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

// Delete implements posting.Handler.
func (pc *PostingController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		postingID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID user tidak valid",
			})
		}

		err = pc.p.HapusPosting(c.Get("user").(*golangjwt.Token), uint(postingID))
		if err != nil {
			c.Logger().Error("ERROR Delete postingan, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika menghapus postingan"

			if strings.Contains(err.Error(), "tidak ditemukan") {
				statusCode = http.StatusNotFound
				message = "postingan tidak ditemukan"
			} else if strings.Contains(err.Error(), "tidak memiliki izin") {
				statusCode = http.StatusForbidden
				message = "Anda tidak memiliki izin untuk menghapus postingan ini"
			}

			return c.JSON(statusCode, map[string]interface{}{
				"message": message,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete postingan",
		})
	}
}

// Update implements posting.Handler.
func (pc *PostingController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(PutPostingRequest)
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

		postingID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID posting tidak valid",
			})
		}

		var inputProcess = models.Posting{
			ID:            uint(postingID),
			Caption:       input.Caption,
			GambarPosting: link,
		}

		result, err := pc.p.UpdatePosting(c.Get("user").(*golangjwt.Token), inputProcess)

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

		var response = new(PutResponse)
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

// GetAll implements posting.Handler.
func (pc *PostingController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if page <= 0 {
			page = 1
		}
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		if limit <= 0 {
			limit = 5
		}
		results, err := pc.p.SemuaPosting(page, limit)
		if err != nil {
			c.Logger().Error("Error fetching coupons: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to retrieve coupon data",
			})
		}
		var response []PostingResponse
		for _, result := range results {
			response = append(response, PostingResponse{
				ID:            result.ID,
				Caption:       result.Caption,
				GambarPosting: result.GambarPosting,
				UserName:      result.UserName,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":    "Success fetching all coupon data",
			"data":       response,
			"pagination": map[string]interface{}{"page": page, "limit": limit},
		})
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
		var inputProcess = new(models.Posting)
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

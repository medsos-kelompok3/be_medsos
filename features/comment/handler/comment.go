package handler

import (
	"be_medsos/features/comment"
	"be_medsos/features/models"
	"net/http"
	"strconv"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

type CommentController struct {
	c comment.Service
}

func New(c comment.Service) comment.Handler {
	return &CommentController{
		c: c,
	}
}

// Delete implements comment.Handler.
func (cc *CommentController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID comment tidak valid",
			})
		}
		err = cc.c.HapusComment(c.Get("user").(*golangjwt.Token), uint(commentID))
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
			"message": "success delete comment",
		})
	}
}

// Add implements comment.Handler.
func (cc *CommentController) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CommentRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}
		var inputProses = new(models.Comment)
		inputProses.PostingID = input.PostingID
		inputProses.IsiComment = input.IsiComment

		result, err := cc.c.AddComment(c.Get("user").(*golangjwt.Token), *inputProses)

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

		var response = new(CommentResponse)
		response.ID = result.ID
		response.PostingID = result.PostingID
		response.IsiComment = result.IsiComment
		response.UserName = result.UserName

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

// Update implements comment.Handler.
func (cc *CommentController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(PutCommentRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		commentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID posting tidak valid",
			})
		}

		var inputProcess = models.Comment{
			ID:         uint(commentID),
			IsiComment: input.IsiComment,
		}

		result, err := cc.c.UpdateComment(c.Get("user").(*golangjwt.Token), inputProcess)

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

		var response = new(PutCommentResponse)
		response.ID = result.ID
		response.PostingID = result.PostingID
		response.IsiComment = result.IsiComment
		response.UserName = result.UserName

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (cc *CommentController) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		commentID, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID comment tidak valid",
				"data":    nil,
			})
		}
		comment, err := cc.c.GetOne(c.Get("user").(*golangjwt.Token), uint(commentID))
		if err != nil {
			if strings.Contains(err.Error(), "ditemukan") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "komen tidak ditemukan",
					"data":    nil,
				})

			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error  pada server",
				"data":    nil,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "ID comment tidak valid",
			"data":    comment,
		})
	}
}

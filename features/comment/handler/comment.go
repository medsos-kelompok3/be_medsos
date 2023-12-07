package handler

import (
	"be_medsos/features/comment"
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

// Add implements comment.Handler.
func (cc *CommentController) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CommentRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}
		var inputProses = new(comment.Comment)
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

		var inputProcess = comment.Comment{
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

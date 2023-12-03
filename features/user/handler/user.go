package handler

import (
	"be_medsos/features/user"
	"be_medsos/helper/jwt"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	srv user.Service
}

func New(s user.Service) user.Handler {
	return &UserController{
		srv: s,
	}
}

// Login implements user.Handler.
func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := uc.srv.Login(input.Username, input.Password)

		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]any{
					"message": "pengguna tidak ditemukan",
				})
			}
			if strings.Contains(err.Error(), "password salah") {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "password salah",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "pengguna tidak di temukan",
			})
		}

		strToken, err := jwt.GenerateJWT(result.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(LoginResponse)
		response.Username = result.Username
		response.ID = result.ID
		response.Token = strToken

		return c.JSON(http.StatusOK, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

package handler

import (
	"be_medsos/features/user"
	"be_medsos/helper/jwt"
	"net/http"
	"strconv"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
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
			"message": "success login data",
			"data":    response,
		})
	}
}

func (uc *UserController) GetListUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userid, err := jwt.ExtractToken(c.Get("user").(*gojwt.Token))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"message": "tidak ada kuasa untuk mengakses",
			})
		}

		return c.JSON(http.StatusOK, userid)
	}
}

// GetAllUser implements user.Handler.
func (uc *UserController) GetAllUserByUsername() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		user, err := uc.srv.DapatUser(username)
		if err != nil {
			c.Logger().Errorf("Failed to get user: %v", err)

			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "user not found",
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error",
			})
		}

		var response = new(GetResponse)
		response.ID = user.ID
		response.Username = user.Username
		response.Bio = user.Bio
		response.Avatar = user.Avatar

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user details",
			"user":    response,
		})
	}
}

// Delete implements user.Handler.
func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID user tidak valid",
			})
		}

		err = uc.srv.HapusUser(c.Get("user").(*gojwt.Token), uint(userID))
		if err != nil {
			c.Logger().Error("ERROR Delete User, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika menghapus user"

			if strings.Contains(err.Error(), "tidak ditemukan") {
				statusCode = http.StatusNotFound
				message = "user tidak ditemukan"
			} else if strings.Contains(err.Error(), "tidak memiliki izin") {
				statusCode = http.StatusForbidden
				message = "Anda tidak memiliki izin untuk menghapus user ini"
			}

			return c.JSON(statusCode, map[string]interface{}{
				"message": message,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user",
		})
	}
}

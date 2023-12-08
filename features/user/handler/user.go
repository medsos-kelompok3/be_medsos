package handler

import (
	"be_medsos/features/models"
	"be_medsos/features/user"
	"be_medsos/helper/jwt"
	cld "be_medsos/utils/cld"
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	gojwt "github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	srv    user.Service
	cl     *cloudinary.Cloudinary
	ct     context.Context
	folder string
}

func New(s user.Service, cld *cloudinary.Cloudinary, ctx context.Context, uploadparam string) user.Handler {
	return &UserController{
		srv:    s,
		cl:     cld,
		ct:     ctx,
		folder: uploadparam,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterReq)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}
		var processInput = new(models.User)
		processInput.Username = input.Username
		processInput.Email = input.Email
		processInput.Address = input.Address
		processInput.Password = input.Password
		result := uc.srv.AddUser(*processInput)
		if result != nil {
			//di sini mendengarkan error message untuk http code yg sesuai
			//"message": "Email sudah didaftarkan" 409
			if strings.Contains(result.Error(), "didaftarkan") {
				return c.JSON(http.StatusConflict, map[string]any{
					"message": "Data sudah didaftarkan, harap login/mendaftar dengan username/email baru",
				})
			}
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "Harap melengkapi data sesuai dengan format",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success login data",
		})
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
		response.Password = result.Password
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
		userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
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

// update
func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(PutRequest)
		userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID user tidak valid",
				"data":    nil,
			})
		}
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Harap Login dulu",
				"data":    nil,
			})
		}
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
				"data":    nil,
			})
		}

		formHeader, err := c.FormFile("avatar")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				var inputProcess = new(models.User)
				inputProcess.Avatar = ""
				inputProcess.ID = input.ID
				inputProcess.Address = input.Address
				inputProcess.Password = input.Password
				inputProcess.NewPassword = input.NewPassword
				inputProcess.Bio = input.Bio
				inputProcess.Email = input.Email
				inputProcess.Username = input.Username

				result, err := uc.srv.UpdateUser(c.Get("user").(*gojwt.Token), *inputProcess)

				if err != nil {
					c.Logger().Error("ERROR Register, explain:", err.Error())
					var statusCode = http.StatusInternalServerError
					var message = "terjadi permasalahan ketika memproses data"

					if strings.Contains(err.Error(), "terdaftar") {
						statusCode = http.StatusBadRequest
						message = "data yang diinputkan sudah terdaftar ada sistem"
					}
					if strings.Contains(err.Error(), "yang lama") {
						statusCode = http.StatusBadRequest
						message = "harap masukkan password yang lama jika ingin mengganti password"
					}

					return c.JSON(statusCode, map[string]any{
						"message": message,
					})
				}

				var response = new(PutResponse)
				response.ID = result.ID
				response.Username = result.Username
				response.Email = result.Email
				response.Bio = result.Bio
				response.Address = result.Address
				response.Avatar = result.Avatar

				return c.JSON(http.StatusCreated, map[string]any{
					"message": "success create data",
					"data":    response,
				})
			}
			return c.JSON(
				http.StatusBadRequest, map[string]any{
					"message": "formheader error",
					"data":    nil,
				})

		}

		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(
				http.StatusBadRequest, map[string]any{
					"message": "formfile error",
					"data":    nil,
				})
		}
		defer formFile.Close()

		link, err := cld.UploadImage(uc.cl, uc.ct, formFile, uc.folder)
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

		var inputProcess = new(models.User)
		inputProcess.Avatar = link
		inputProcess.ID = input.ID
		inputProcess.Address = input.Address
		inputProcess.Password = input.Password
		inputProcess.NewPassword = input.NewPassword
		inputProcess.Bio = input.Bio
		inputProcess.Email = input.Email
		inputProcess.Username = input.Username

		result, err := uc.srv.UpdateUser(c.Get("user").(*gojwt.Token), *inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}
			if strings.Contains(err.Error(), "yang lama") {
				statusCode = http.StatusBadRequest
				message = "harap masukkan password yang lama jika ingin mengganti password"
			}

			return c.JSON(statusCode, map[string]any{
				"message": message,
			})
		}

		var response = new(PutResponse)
		response.ID = result.ID
		response.Username = result.Username
		response.Email = result.Email
		response.Bio = result.Bio
		response.Address = result.Address
		response.Avatar = result.Avatar

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

// get user details
func (uc *UserController) GetUserDetails() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID user tidak valid",
				"data":    nil,
			})
		}
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Harap Login dulu",
				"data":    nil,
			})
		}

		// ngambil dari repo
		proses, err := uc.srv.GetUserDetails(c.Get("user").(*gojwt.Token), uint(userID))
		if err != nil {
			if strings.Contains(err.Error(), "ditemukann") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "User tidak ditemukan",
					"data":    nil,
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Server error",
				"data":    nil,
			})
		}
		response := GetResponse{
			ID:       proses.ID,
			Username: proses.Username,
			Email:    proses.Email,
			Address:  proses.Address,
			Bio:      proses.Bio,
			Avatar:   proses.Avatar,
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Server error",
			"data":    response,
		})
	}
}

// get user profiles
func (uc *UserController) GetUserProfiles() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "ID user tidak valid",
				"data":    nil,
			})
		}
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Harap Login dulu",
				"data":    nil,
			})
		}

		proses, postings, err := uc.srv.GetUserProfiles(c.Get("user").(*gojwt.Token), uint(userID))
		if err != nil {
			if strings.Contains(err.Error(), "ditemukann") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "User tidak ditemukan",
					"data":    nil,
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Server error",
				"data":    nil,
			})
		}
		var response GetProfilResponse
		response.ID = proses.ID
		response.Username = proses.Username
		response.Address = proses.Address
		response.Bio = proses.Bio
		response.Avatar = proses.Avatar

		response.Posts = append(response.Posts, postings)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Data bberhasil dimuat",
			"data":    response,
		})
	}
}

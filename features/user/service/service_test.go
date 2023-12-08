package service_test

import (
	"be_medsos/features/models"
	"be_medsos/features/user/mocks"
	"be_medsos/features/user/service"
	eMock "be_medsos/helper/enkrip/mocks"
	"be_medsos/helper/jwt"
	"errors"
	"testing"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var userID = uint(1)
var str, _ = jwt.GenerateJWT(userID)
var token, _ = gojwt.Parse(str, func(t *gojwt.Token) (interface{}, error) {
	return []byte("$!1gnK3yyy!!!"), nil
})

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	enkrip := eMock.NewHashInterface(t)
	m := service.New(repo, enkrip)

	var userData = models.User{Username: "john_doe", Password: "password123"}
	var repoData = models.User{Username: "john_doe", Password: "some string"}
	var falseUserData = models.User{Username: "", Password: ""}

	t.Run("Success Case", func(t *testing.T) {
		enkrip.On("HashPassword", repoData.Password).Return("some string", nil).Once()

		userData.Password = "some string"
		repo.On("AddUser", repoData).Return(nil).Once()
		err := m.AddUser(repoData)

		enkrip.AssertExpectations(t)
		repo.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Failed Case - Empty Username and Password", func(t *testing.T) {
		err := m.AddUser(falseUserData)

		assert.Error(t, err)
		assert.Equal(t, "username and password are required", err.Error())

		repo.AssertExpectations(t)
	})
}

func TestLogin_Success(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	inputName := "testuser"
	inputPassword := "testpassword"

	repo.On("Login", inputName).Return(models.User{
		ID:       1,
		Username: inputName,
		Password: "hashedpassword",
	}, nil).Once()

	hashMock.On("Compare", "hashedpassword", inputPassword).Return(nil).Once()

	result, err := userService.Login(inputName, inputPassword)

	repo.AssertExpectations(t)
	hashMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, inputName, result.Username)
}

func TestLogin_EmptyUsernameOrPassword(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	result, err := userService.Login("", "")

	assert.Error(t, err)
	assert.Equal(t, models.User{}, result)
	assert.Equal(t, "username and password are required", err.Error())
}

func TestLogin_UserNotFound(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	inputName := "data tidak ditemukan"
	inputPassword := "testpassword"

	repo.On("Login", inputName).Return(models.User{}, errors.New("data tidak ditemukan")).Once()

	result, err := userService.Login(inputName, inputPassword)

	repo.AssertExpectations(t)
	hashMock.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, models.User{}, result)
	assert.Equal(t, "terjadi kesalahan pada sistem", err.Error())
}

func TestLogin_IncorrectPassword(t *testing.T) {
	repo := mocks.NewRepository(t)
	hashMock := eMock.NewHashInterface(t)

	userService := service.New(repo, hashMock)

	inputName := "testuser"
	inputPassword := "password salah"

	repo.On("Login", inputName).Return(models.User{
		ID:       1,
		Username: inputName,
		Password: "hashedpassword",
	}, nil).Once()

	hashMock.On("Compare", "hashedpassword", inputPassword).Return(errors.New("password salah")).Once()

	result, err := userService.Login(inputName, inputPassword)

	repo.AssertExpectations(t)
	hashMock.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, models.User{}, result)
	assert.Equal(t, "password salah", err.Error())
}

func TestDapatUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	userService := service.New(repo, nil)

	username := "john_doe"
	expectedUser := models.User{
		Username: username,
	}

	t.Run("Success Case", func(t *testing.T) {
		repo.On("GetUserByUsername", username).Return(expectedUser, nil).Once()

		result, err := userService.DapatUser(username)

		assert.Nil(t, err)
		assert.Equal(t, expectedUser, result)

		repo.AssertExpectations(t)
	})

	t.Run("Error Case - Repository Failure", func(t *testing.T) {
		repo.On("GetUserByUsername", username).Return(models.User{}, errors.New("database error")).Once()

		_, err := userService.DapatUser(username)

		assert.Error(t, err)
		assert.Equal(t, "failed to retrieve inserted Data", err.Error())

		repo.AssertExpectations(t)
	})

}

type MockRepository struct{}

func (mr *MockRepository) Insert(newUser models.User) (models.User, error) {
	return models.User{ID: uint(1), Username: "jerry"}, nil
}
func (mr *MockRepository) Login(Username string) (models.User, error) {
	return models.User{}, nil
}

type FalseMockRepository struct{}

func (fmr *FalseMockRepository) Insert(newUser models.User) (models.User, error) {
	return models.User{}, errors.New("something happend")
}
func (fmr *FalseMockRepository) Login(Username string) (models.User, error) {
	return models.User{}, nil
}

type MockEnkrip struct{}

func (me *MockEnkrip) Compare(hashed string, input string) error {
	return nil
}
func (me *MockEnkrip) HashPassword(input string) (string, error) {
	return "some string", nil
}

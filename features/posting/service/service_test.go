package service_test

import (
	"be_medsos/features/posting"
	"be_medsos/features/posting/mocks"
	"be_medsos/features/posting/service"
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

func TestAddPosting(t *testing.T) {
	repo := mocks.NewRepository(t)
	m := service.New(repo)

	var repoData = posting.Posting{Caption: "hoax", GambarPosting: "wwww.facebook.com", UserName: "budi"}
	var falseData = posting.Posting{}

	t.Run("Success Case", func(t *testing.T) {
		repo.On("InsertPosting", userID, repoData).Return(repoData, nil).Once()
		res, err := m.AddPosting(token, repoData)

		assert.Nil(t, err)
		assert.Equal(t, "budi", res.UserName)
		assert.Equal(t, "hoax", res.Caption)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		repo.On("InsertPosting", userID, falseData).Return(falseData, errors.New("ERROR db")).Once()
		res, err := m.AddPosting(token, falseData)
		assert.Error(t, err)
		assert.Equal(t, "", res.UserName)
		assert.Equal(t, "", res.Caption)
		repo.AssertExpectations(t)
	})
}

func TestSemuaPosting(t *testing.T) {
	repo := mocks.NewRepository(t)
	m := service.New(repo)

	t.Run("Success Case", func(t *testing.T) {
		page := 1
		limit := 5

		// Set up expectations for mock dependencies
		repo.On("GetTanpaPosting", page, limit).Return([]posting.Posting{
			{ID: 1, Caption: "Test Caption 1", GambarPosting: "www.fawa.com"},
			{ID: 2, Caption: "Test Caption 2", GambarPosting: "www.fawa.com"},
			// Add other postings as needed
		}, nil).Once()

		// Invoke the service method
		result, err := m.SemuaPosting(page, limit)

		// Assertions
		assert.Nil(t, err)
		assert.Len(t, result, 2) // Adjust the expected length based on the mocked data

		// Verify that the expectations were met
		repo.AssertExpectations(t)
	})

	t.Run("Failure Case - Repository Error", func(t *testing.T) {
		// Set up input data
		page := 1
		limit := 5

		// Set up expectations for mock dependencies
		repo.On("GetTanpaPosting", page, limit).Return(nil, errors.New("repository error")).Once()

		// Invoke the service method
		result, err := m.SemuaPosting(page, limit)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, result)

		// Verify that the expectations were met
		repo.AssertExpectations(t)
	})
}

func TestUpdatePosting(t *testing.T) {
	repo := mocks.NewRepository(t)
	m := service.New(repo)

	var repoData = posting.Posting{ID: 1, Caption: "updated caption", GambarPosting: "wwww.updated.com", UserName: "budi"}
	var falseData = posting.Posting{ID: 2}

	t.Run("Success Case", func(t *testing.T) {
		repo.On("UpdatePosting", repoData).Return(repoData, nil).Once()
		res, err := m.UpdatePosting(token, repoData)

		assert.Nil(t, err)
		assert.Equal(t, "budi", res.UserName)
		assert.Equal(t, "updated caption", res.Caption)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Post Not Found", func(t *testing.T) {
		repo.On("UpdatePosting", falseData).Return(falseData, errors.New("Post not found")).Once()
		res, err := m.UpdatePosting(token, falseData)
		assert.Error(t, err)
		assert.Equal(t, "", res.UserName)
		assert.Equal(t, "", res.Caption)
		repo.AssertExpectations(t)
	})
}

func TestHapusPosting(t *testing.T) {
	repo := mocks.NewRepository(t)
	m := service.New(repo)

	// Set up input data
	postingID := uint(1)
	testToken := &gojwt.Token{
		Claims: gojwt.MapClaims{
			"id": float64(1), // Replace with the appropriate claim key and value
		},
		// Set other fields as needed, e.g., Header, Method, etc.
	}

	t.Run("Success Case", func(t *testing.T) {
		// Set up expectations for mock dependencies
		repo.On("DeletePosting", postingID).Return(nil).Once()

		// Invoke the service method
		err := m.HapusPosting(testToken, postingID)

		// Assertions
		assert.Nil(t, err)

		// Verify that the expectations were met
		repo.AssertExpectations(t)
	})

	t.Run("Failure Case - Repository Error", func(t *testing.T) {
		// Set up expectations for mock dependencies
		repo.On("DeletePosting", postingID).Return(errors.New("repository error")).Once()

		// Invoke the service method
		err := m.HapusPosting(testToken, postingID)

		// Assertions
		assert.Error(t, err)

		// Verify that the expectations were met
		repo.AssertExpectations(t)
	})
}

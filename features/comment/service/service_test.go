package service_test

import (
	"be_medsos/helper/jwt"

	gojwt "github.com/golang-jwt/jwt/v5"
)

var userID = uint(1)
var str, _ = jwt.GenerateJWT(userID)
var token, _ = gojwt.Parse(str, func(t *gojwt.Token) (interface{}, error) {
	return []byte("$!1gnK3yyy!!!"), nil
})

// func TestAddComment(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	m := service.New(repo)

// 	userID := uint(1)
// 	postingID := uint(2)
// 	newComment := models.Comment{
// 		PostingID:  uint(2),
// 		IsiComment: "Test comment",
// 	}

// 	errorComment := models.Comment{
// 		PostingID:  uint(2),
// 		IsiComment: "",
// 	}

// 	t.Run("Success Case", func(t *testing.T) {
// 		repo.On("InsertComment", userID, postingID, newComment).Return(newComment, nil).Once()
// 		result, err := m.AddComment(token, newComment)

// 		assert.Nil(t, err)
// 		assert.Equal(t, newComment, result)

// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Failed Case", func(t *testing.T) {
// 		repo.On("InsertComment", userID, postingID, errorComment).Return(errorComment, nil).Once()
// 		result, err := m.AddComment(token, errorComment)

// 		assert.Nil(t, err)
// 		assert.Equal(t, errorComment, result)

// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdateComment(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	m := service.New(repo)

// 	token := &gojwt.Token{
// 		Claims: gojwt.MapClaims{
// 			"id": float64(1),
// 		},
// 	}
// 	input := models.Comment{
// 		IsiComment: "Updated comment",
// 		PostingID:  uint(2),
// 	}

// 	salah := models.Comment{
// 		IsiComment: "",
// 		PostingID:  uint(2),
// 	}
// 	t.Run("Success Case", func(t *testing.T) {
// 		repo.On("UpdateComment", input).Return(input, nil).Once()
// 		result, err := m.UpdateComment(token, input)
// 		assert.Nil(t, err)
// 		assert.Equal(t, input, result)

// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed Case", func(t *testing.T) {
// 		repo.On("UpdateComment", salah).Return(salah, nil).Once()
// 		result, err := m.UpdateComment(token, salah)
// 		assert.Nil(t, err)
// 		assert.Equal(t, salah, result)

// 		repo.AssertExpectations(t)
// 	})

// }

// func TestDelete(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	m := service.New(repo)

// 	commentID := uint(2)
// 	token := &gojwt.Token{
// 		Claims: gojwt.MapClaims{
// 			"id": float64(1),
// 		},
// 	}
// 	SalahID := uint(0)
// 	t.Run("Success Case", func(t *testing.T) {
// 		repo.On("DeleteComment", commentID).Return(nil).Once()
// 		err := m.HapusComment(token, commentID)
// 		assert.Nil(t, err)

// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Failed Case", func(t *testing.T) {
// 		repo.On("DeleteComment", SalahID).Return(nil).Once()
// 		err := m.HapusComment(token, SalahID)
// 		assert.Nil(t, err)

// 		repo.AssertExpectations(t)
// 	})
// }

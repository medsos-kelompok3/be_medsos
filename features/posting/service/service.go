package service

import (
	"be_medsos/features/posting"
	"be_medsos/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type PostingService struct {
	repo posting.Repository
}

// AddPosting implements posting.Service.
func (ps *PostingService) AddPosting(token *golangjwt.Token, newPosting posting.Posting) (posting.Posting, error) {
	userId, err := jwt.ExtractToken(token)
	if err != nil {
		return posting.Posting{}, err
	}
	result, err := ps.repo.InsertPosting(userId, newPosting)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return posting.Posting{}, errors.New("dobel input")
		}
		return posting.Posting{}, errors.New("terjadi kesalahan")
	}
	return result, nil
}

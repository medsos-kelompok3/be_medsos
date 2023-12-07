package service

import (
	"be_medsos/features/posting"
	"be_medsos/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type PostingService struct {
	r posting.Repository
}

// UpdatePosting implements posting.Service.
func (ps *PostingService) UpdatePosting(token *golangjwt.Token, input posting.Posting) (posting.Posting, error) {
	respons, err := ps.r.UpdatePosting(input)
	if err != nil {

		return posting.Posting{}, errors.New("kesalahan pada database")
	}
	return respons, nil
}

func New(model posting.Repository) posting.Service {
	return &PostingService{
		r: model,
	}
}

// SemuaPosting implements posting.Service.
func (ps *PostingService) SemuaPosting(page int, limit int) ([]posting.Posting, error) {
	result, err := ps.r.GetTanpaPosting(page, limit)
	if err != nil {
		return nil, errors.New("failed to retrieve inserted coupon")
	}
	return result, nil
}

// AddPosting implements posting.Service.
func (ps *PostingService) AddPosting(token *golangjwt.Token, newPosting posting.Posting) (posting.Posting, error) {
	userId, err := jwt.ExtractToken(token)
	if err != nil {
		return posting.Posting{}, err
	}
	result, err := ps.r.InsertPosting(userId, newPosting)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return posting.Posting{}, errors.New("dobel input")
		}
		return posting.Posting{}, errors.New("terjadi kesalahan")
	}
	return result, nil
}

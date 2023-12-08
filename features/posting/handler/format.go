package handler

import (
	"be_medsos/features/models"
	"mime/multipart"
)

type PostingRequest struct {
	Caption       string `json:"caption" form:"caption"`
	GambarPosting string `json:"gambar_posting" form:"gambar_posting"`
}

type PostingResponse struct {
	ID            uint   `json:"id"`
	Caption       string `json:"caption"`
	GambarPosting string `json:"gambar_posting"`
	UserName      string `json:"user_name"`
	Avatar        string `json:"avatar"`
	CreatedAt     string `json:"createdat"`
}

type PutPostingRequest struct {
	ID            uint           `json:"id" form:"id"`
	Caption       string         `json:"caption" form:"caption"`
	GambarPosting multipart.File `json:"gambar_posting" form:"gambar_posting"`
}

type PutResponse struct {
	ID            uint   `json:"id" form:"id"`
	Caption       string `json:"caption" form:"caption"`
	GambarPosting string `json:"gambar_posting" form:"gambar_posting"`
	UserName      string `json:"user_name"`
	Avatar        string `json:"avatar"`

	CreatedAt string `json:"createdat"`
}

type GetResponse struct {
	ID            uint             `json:"id"`
	Caption       string           `json:"caption"`
	GambarPosting string           `json:"gambar_posting"`
	UserName      string           `json:"user_name"`
	Avatar        string           `json:"avatar"`
	CreatedAt     string           `json:"createdat"`
	Comments      []models.Comment `json:"comments"`
}

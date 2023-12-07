package handler

import "mime/multipart"

type PostingRequest struct {
	Caption       string `json:"caption" form:"caption"`
	GambarPosting string `json:"gambar_posting" form:"gambar_posting"`
}

type PostingResponse struct {
	ID            uint   `json:"id"`
	Caption       string `json:"caption"`
	GambarPosting string `json:"gambar_posting"`
	UserName      string `json:"user_name"`
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
}

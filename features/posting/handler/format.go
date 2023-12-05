package handler

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

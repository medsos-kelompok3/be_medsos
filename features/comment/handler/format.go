package handler

type CommentRequest struct {
	PostingID  uint   `json:"posting_id" form:"posting_id"`
	IsiComment string `json:"isi_comment" form:"isi_comment"`
}

type CommentResponse struct {
	ID         uint   `json:"id"`
	PostingID  uint   `json:"posting_id"`
	IsiComment string `json:"isi_comment"`
	UserName   string `json:"user_name"`
}

type PutCommentRequest struct {
	ID         uint   `json:"id" form:"id"`
	IsiComment string `json:"isi_comment" form:"isi_comment"`
}

type PutCommentResponse struct {
	ID         uint   `json:"id"`
	PostingID  uint   `json:"posting_id"`
	IsiComment string `json:"isi_comment"`
	UserName   string `json:"user_name"`
}

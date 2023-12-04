package handler

import "mime/multipart"

type RegisterReq struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

type RegisResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}

type PutRequest struct {
	ID       uint           `json:"id" form:"id"`
	Username string         `json:"username" form:"username"`
	Email    string         `json:"email" form:"email"`
	Address  string         `json:"address" form:"address"`
	Bio      string         `json:"bio" form:"bio"`
	Avatar   multipart.File `json:"avatar" form:"avatar"`
}

type PutResponse struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	Bio      string `json:"bio" form:"bio"`
	Avatar   string `json:"avatar" form:"avatar"`
}

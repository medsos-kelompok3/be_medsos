package models

import "gorm.io/gorm"

type User struct {
	ID          uint
	Username    string
	Email       string
	Address     string
	Bio         string
	Avatar      string
	Password    string
	NewPassword string
}

type Posting struct {
	ID            uint
	Caption       string
	GambarPosting string
	UserID        uint
	UserName      string
	Avatar        string
	CreatedAt     string
	CommentCount  int64
}

type Comment struct {
	ID         uint
	PostingID  uint
	UserID     uint
	UserName   string
	Avatar     string
	IsiComment string
	CreatedAt  string
}

type UserModel struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Address  string
	Bio      string
	Avatar   string
	Password string
}

type PostingModel struct {
	gorm.Model
	Caption       string
	GambarPosting string
	UserName      string
	User_id       uint
	Avatar        string
}

type CommentModel struct {
	gorm.Model
	PostingID  uint
	IsiComment string
	UserName   string
	UserID     uint
	Avatar     string
}

package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"size:50;unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"` // 序列化时忽略密码
	Email    string `gorm:"size:100;unique;not null" json:"email"`
	Posts    []Post `json:"posts,omitempty"`
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title    string    `gorm:"size:200;not null" json:"title"`
	Content  string    `gorm:"type:text;not null" json:"content"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	User     User      `json:"user,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"size:500;not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `json:"user,omitempty"`
	PostID  uint   `gorm:"not null" json:"post_id"`
	Post    Post   `json:"post,omitempty"`
}

package Request

import (
	"time"
)

type CreatePostRequest struct {
	Title     string    `json:"title" bson:"title"  binding:"required"`
	Email     string    `json:"email" bson:"email" binding:"required"`
	Token     string    `json:"token" bson:"token" binding:"required"`
	Content   string    `json:"content" bson:"content"  binding:"required"`
	Image     string    `json:"image" bson:"image"  binding:"required"`
	User      string    `json:"user" bson:"user"  binding:"required"`
	CreateAt  time.Time `json:"create_at,omitempty" bson:"create_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type UpdatePost struct {
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Token     string    `json:"token,omitempty" bson:"token,omitempty"`
	Content   string    `json:"content,omitempty" bson:"content,omitempty"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty"`
	User      string    `json:"user,omitempty" bson:"user,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type TokenRequest struct {
	Email string `json:"email" bson:"email" binding:"required"`
	Token string `json:"token" bson:"token"binding:"required"`
}

type FromReq struct {
	Title     string    `form:"title" bson:"title"`
	Email     string    `form:"email" bson:"email"`
	Token     string    `form:"token" bson:"token"`
	Content   string    `form:"content" bson:"content"`
	Image     string    `form:"image" bson:"image"`
	User      string    `form:"User" bson:"user"`
	CreateAt  time.Time `form:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `form:"updated_at" bson:"updated_at,omitempty"`
}

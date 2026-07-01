package services

import (
	"sync"
)

type Post struct {
	ID          string `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type PostService struct {
	mu    sync.Mutex
	posts []Post
}

func NewPostService() *PostService {
	return &PostService{
		posts: []Post{},
	}
}

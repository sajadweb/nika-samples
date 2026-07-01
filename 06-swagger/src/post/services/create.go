package services

import (
	"context"
	"fmt"

	"nikaapp/src/post/dto"
)

func (s *PostService) Create(ctx context.Context, input *dto.CreatePostDto) (Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%d", len(s.posts)+1)
	newPost := Post{
		ID:          id,
		Subject:     input.Subject,
		Description: input.Description,
	}
	s.posts = append(s.posts, newPost)
	return newPost, nil
}

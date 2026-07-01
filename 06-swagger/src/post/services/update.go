package services

import (
	"context"
	"errors"
	"nikaapp/src/post/dto"
)

func (s *PostService) Update(ctx context.Context, body *dto.UpdatePostDto, input *dto.FindOnePostDto) (Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.posts {
		if p.ID == input.ID {
			upPost := p
			if body.Subject != "" {
				upPost.Subject = body.Subject
			}
			if body.Description != "" {
				upPost.Description = body.Description
			}
			s.posts[i] = upPost
			return upPost, nil
		}
	}
	return Post{}, errors.New("post not found")
}

package services

import (
	"context"
	"errors"
	"nikaapp/src/post/dto"
)

func (s *PostService) Delete(ctx context.Context, input *dto.FindOnePostDto) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.posts {
		if p.ID == input.ID {
			s.posts = append(s.posts[:i], s.posts[i+1:]...)
			return nil
		}
	}
	return errors.New("post not found")
}

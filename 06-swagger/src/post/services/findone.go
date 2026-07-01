package services

import (
	"context"

	"nikaapp/src/post/dto"
	"errors" 
)

 
func (s *PostService) FindOne(ctx context.Context, input *dto.FindOnePostDto) (Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, p := range s.posts {
		if p.ID == input.ID {
			return p, nil
		}
	}
	return Post{}, errors.New("post not found" + input.ID)
}

package services

import (
	"context"

	"nikaapp/src/post/dto" 
)
 

func (s *PostService) Find(ctx context.Context, input *dto.ListPostDto) ([]Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	postsCopy := make([]Post, len(s.posts))
	copy(postsCopy, s.posts)
	return postsCopy,nil
}

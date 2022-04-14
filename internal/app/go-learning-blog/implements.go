package golearningblog

import (
	"context"
	blog "github.com/x666ep/go-learning-blog/pkg/api/go-learning-blog/v1"
)

func (s *Implementation) GetPosts(ctx context.Context, in *blog.Empty) (*blog.Posts, error) {
	posts := []*blog.Post{
		{
			Author: "kserov",
			Head:   "TestPost",
			Body:   "super body",
		},
	}

	return &blog.Posts{Posts: posts}, nil
}

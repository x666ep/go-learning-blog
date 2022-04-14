package golearningblog

import (
	blog "github.com/x666ep/go-learning-blog/pkg/api/go-learning-blog/v1"
)

type Implementation struct {
	blog.UnimplementedGoLearningBlogServiceServer
}

func NewGoLearningBlogApi() *Implementation {
	return &Implementation{}
}

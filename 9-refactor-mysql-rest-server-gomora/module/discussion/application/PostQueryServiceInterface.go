package application

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/infrastructure/service/types"
)

// PostQueryServiceInterface holds the implementable method for the Post query service
type PostQueryServiceInterface interface {
	GetPosts(ctx context.Context, data types.GetPost) ([]entity.Post, error)
	GetPostByID(ctx context.Context, data types.GetPost) ([]entity.Post, error)
}

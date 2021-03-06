package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "rest-server/configs/hystrix"
	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// DiscussionCommandRepositoryCircuitBreaker circuit breaker for post command repository
type DiscussionCommandRepositoryCircuitBreaker struct {
	repository.DiscussionCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// =======================================POST=======================================

// DeletePostByID is the decorator for the the post repository delete by id method
func (repository *DiscussionCommandRepositoryCircuitBreaker) DeletePostByID(postID int64) error {
	hystrix.ConfigureCommand("delete_post_by_id", config.Settings())
	errors := hystrix.Go("delete_post_by_id", func() error {
		err := repository.DiscussionCommandRepositoryInterface.DeletePostByID(postID)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertPost decorator pattern to insert post
func (repository *DiscussionCommandRepositoryCircuitBreaker) InsertPost(data repositoryTypes.CreatePost) (entity.Post, error) {
	output := make(chan entity.Post, 1)
	hystrix.ConfigureCommand("insert_post", config.Settings())
	errors := hystrix.Go("insert_post", func() error {
		post, err := repository.DiscussionCommandRepositoryInterface.InsertPost(data)
		if err != nil {
			return err
		}

		output <- post
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Post{}, err
	}
}

// UpdatePostByID is the decorator for the post repository update post method
func (repository *DiscussionCommandRepositoryCircuitBreaker) UpdatePostByID(data repositoryTypes.UpdatePost) (entity.Post, error) {
	output := make(chan entity.Post, 1)
	hystrix.ConfigureCommand("update_post", config.Settings())
	errors := hystrix.Go("update_post", func() error {
		post, err := repository.DiscussionCommandRepositoryInterface.UpdatePostByID(data)
		if err != nil {
			return err
		}

		output <- post
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Post{}, err
	}
}

// =======================================COMMENT=======================================

// DeleteCommentByID is the decorator for the the Comment repository delete by id method
func (repository *DiscussionCommandRepositoryCircuitBreaker) DeleteCommentByID(commentID int64) error {
	hystrix.ConfigureCommand("delete_comment_by_id", config.Settings())
	errors := hystrix.Go("delete_comment_by_id", func() error {
		err := repository.DiscussionCommandRepositoryInterface.DeleteCommentByID(commentID)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertComment decorator pattern to insert Comment
func (repository *DiscussionCommandRepositoryCircuitBreaker) InsertComment(data repositoryTypes.CreateComment) (entity.Comment, error) {
	output := make(chan entity.Comment, 1)
	hystrix.ConfigureCommand("insert_comment", config.Settings())
	errors := hystrix.Go("insert_comment", func() error {
		comment, err := repository.DiscussionCommandRepositoryInterface.InsertComment(data)
		if err != nil {
			return err
		}

		output <- comment
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Comment{}, err
	}
}

// UpdateCommentByID is the decorator for the Comment repository update Comment method
func (repository *DiscussionCommandRepositoryCircuitBreaker) UpdateCommentByID(data repositoryTypes.UpdateComment) (entity.Comment, error) {
	output := make(chan entity.Comment, 1)
	hystrix.ConfigureCommand("update_comment", config.Settings())
	errors := hystrix.Go("update_comment", func() error {
		comment, err := repository.DiscussionCommandRepositoryInterface.UpdateCommentByID(data)
		if err != nil {
			return err
		}

		output <- comment
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Comment{}, err
	}
}

package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nooble/task/audio-short-api/pkg/api/generated"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/logging"
)

func (r *mutationResolver) CreateAudioShort(ctx context.Context, input model.AudioShortInput) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Create Audio Short")
	return r.store.Create(ctx, &input)
}

func (r *mutationResolver) UpdateAudioShort(ctx context.Context, id string, input model.AudioShortInput) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Update Audio Short with id " + id)
	return r.store.Update(ctx, id, &input)
}

func (r *mutationResolver) DeleteAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Delete Audio Short with id " + id)
	return r.store.Delete(ctx, id)
}

func (r *queryResolver) GetAudioShorts(ctx context.Context, page *int, limit *int) ([]*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Get Audio Shorts")
	return r.store.GetAll(ctx, uint16(*page), uint16(*limit))
}

func (r *queryResolver) GetAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Get Audio Short with id " + id)
	return r.store.GetByID(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

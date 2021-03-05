package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/pkg/errors"

	"github.com/nooble/task/audio-short-api/pkg/api/generated"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/logging"
)

func (r *mutationResolver) CreateAudioShort(ctx context.Context, input model.AudioShortInput) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Create Audio Short")
	short, err := r.store.Create(ctx, &input)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageCreateFailed).Error())
		return nil, errors.New(ErrorMessageCreateFailed)
	}
	return short, nil
}

func (r *mutationResolver) UpdateAudioShort(ctx context.Context, id string, input model.AudioShortInput) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Update Audio Short with id " + id)
	short, err := r.store.Update(ctx, id, &input)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageUpdateFailed).Error())
		return nil, errors.New(ErrorMessageUpdateFailed)
	}
	return short, nil
}

func (r *mutationResolver) DeleteAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Delete Audio Short with id " + id)
	short, err := r.store.Delete(ctx, id)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageDeleteFailed).Error())
		return nil, errors.New(ErrorMessageDeleteFailed)
	}
	return short, nil
}

func (r *mutationResolver) HardDeleteAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Hard Delete Audio Short with id " + id)
	short, err := r.store.HardDelete(ctx, id)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageHardDeleteFailed).Error())
		return nil, errors.New(ErrorMessageHardDeleteFailed)
	}
	return short, nil
}

func (r *queryResolver) GetAudioShorts(ctx context.Context, page *int, limit *int) ([]*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Get Audio Shorts")
	shorts, err := r.store.GetAll(ctx, uint16(*page), uint16(*limit))
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageReadFailed).Error())
		return nil, errors.New(ErrorMessageReadFailed)
	}
	return shorts, nil
}

func (r *queryResolver) GetAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Get Audio Short with id " + id)
	short, err := r.store.GetByID(ctx, id)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageReadFailed).Error())
		return nil, errors.New(ErrorMessageReadFailed)
	}
	return short, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

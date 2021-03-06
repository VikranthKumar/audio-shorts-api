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
	short, err := r.shortsStore.Create(ctx, &input)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageCreateFailed).Error())
		return nil, errors.New(ErrorMessageCreateFailed)
	}
	return short, nil
}

func (r *mutationResolver) UpdateAudioShort(ctx context.Context, id string, input model.AudioShortInput) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Update Audio Short With ID " + id)
	short, err := r.shortsStore.Update(ctx, id, &input)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageUpdateFailed).Error())
		return nil, errors.New(ErrorMessageUpdateFailed)
	}
	return short, nil
}

func (r *mutationResolver) DeleteAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Delete Audio Short With ID " + id)
	short, err := r.shortsStore.Delete(ctx, id)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageDeleteFailed).Error())
		return nil, errors.New(ErrorMessageDeleteFailed)
	}
	return short, nil
}

func (r *mutationResolver) HardDeleteAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Hard Delete Audio Short With ID " + id)
	short, err := r.shortsStore.HardDelete(ctx, id)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageHardDeleteFailed).Error())
		return nil, errors.New(ErrorMessageHardDeleteFailed)
	}
	return short, nil
}

func (r *queryResolver) GetAudioShorts(ctx context.Context, page *int, limit *int) ([]*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Get Audio Shorts")
	if *page < 1 {
		return nil, errors.New(ErrorMessageBadRequest)
	}
	shorts, err := r.shortsStore.GetAll(ctx, uint16(*page)-1, uint16(*limit))
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageReadFailed).Error())
		return nil, errors.New(ErrorMessageReadFailed)
	}
	return shorts, nil
}

func (r *queryResolver) GetAudioShort(ctx context.Context, id string) (*model.AudioShort, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Get Audio Short With ID " + id)
	short, err := r.shortsStore.GetByID(ctx, id)
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageReadFailed).Error())
		return nil, errors.New(ErrorMessageReadFailed)
	}
	return short, nil
}

func (r *queryResolver) GetCreators(ctx context.Context, page *int, limit *int) ([]*model.Creator, error) {
	ctx = logging.NewContext(ctx)
	logging.WithContext(ctx).Info("Get Creators")
	if *page < 1 {
		return nil, errors.New(ErrorMessageBadRequest)
	}
	shorts, err := r.creatorsStore.GetAll(ctx, uint16(*page)-1, uint16(*limit))
	if err != nil {
		logging.WithContext(ctx).Error(errors.Wrap(err, ErrorMessageReadFailed).Error())
		return nil, errors.New(ErrorMessageReadFailed)
	}
	return shorts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

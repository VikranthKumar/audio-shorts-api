package store

import (
	"context"
	"database/sql"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"sync"
)

type (
	AudioShortsStore interface {
		GetByID(ctx context.Context, id string) (short *model.AudioShort, err error)
		GetAll(ctx context.Context, page, limit uint16) (shorts []*model.AudioShort, err error)
		Create(ctx context.Context, input *model.AudioShortInput) (short *model.AudioShort, err error)
		Update(ctx context.Context, id string, input *model.AudioShortInput) (short *model.AudioShort, err error)
		Delete(ctx context.Context, id string) (short *model.AudioShort, err error)
	}

	shortsStore struct {
		db *sql.DB
		sync.RWMutex
	}
)

func New(db *sql.DB) (AudioShortsStore, error) {
	return &shortsStore{
		db: db,
	}, nil
}

func (s *shortsStore) GetByID(ctx context.Context, id string) (short *model.AudioShort, err error) {
	return
}

func (s *shortsStore) GetAll(ctx context.Context, page, limit uint16) (shorts []*model.AudioShort, err error) {
	return
}

func (s *shortsStore) Create(ctx context.Context, input *model.AudioShortInput) (short *model.AudioShort, err error) {
	return
}

func (s *shortsStore) Update(ctx context.Context, id string, input *model.AudioShortInput) (short *model.AudioShort, err error) {
	return
}

func (s *shortsStore) Delete(ctx context.Context, id string) (short *model.AudioShort, err error) {
	return
}

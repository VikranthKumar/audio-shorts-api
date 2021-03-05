package store

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/logging"
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
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New("cannot start transaction")
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error("failed to rollback")
			}
		}
	}()

	short, err = findOneByID(ctx, tx, id)
	if err != nil {
		return nil, errors.New("cannot find audio short with ID " + id)
	}
	logging.WithContext(ctx).Debug(short.Title)

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("cannot commit transaction")
	}
	return
}

func (s *shortsStore) GetAll(ctx context.Context, page, limit uint16) (shorts []*model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New("cannot start transaction")
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error("failed to rollback")
			}
		}
	}()

	shorts, err = findAll(ctx, tx, page, limit)
	if err != nil {
		return nil, errors.New("failed to find audio shorts")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("cannot commit transaction")
	}
	return
}

func (s *shortsStore) Create(ctx context.Context, input *model.AudioShortInput) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New("cannot start transaction")
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			logging.WithContext(ctx).Error(err.Error())
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error("failed to rollback")
			}
		}
	}()

	err = createOne(ctx, tx, input)
	if err != nil {
		return nil, errors.New("failed to create audio short")
	}
	short, err = findOneByUnique(ctx, tx, input.Title, input.Creator.ID)
	if err != nil {
		return nil, errors.New("cannot find audio short with ID ")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("cannot commit transaction")
	}
	return
}

func (s *shortsStore) Update(ctx context.Context, id string, input *model.AudioShortInput) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New("cannot start transaction")
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			logging.WithContext(ctx).Error(err.Error())
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error("failed to rollback")
			}
		}
	}()

	err = updateOne(ctx, tx, id, input)
	if err != nil {
		return nil, errors.New("failed to update audio short")
	}
	short, err = findOneByID(ctx, tx, id)
	if err != nil {
		return nil, errors.New("cannot find audio short with ID " + id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("cannot commit transaction")
	}
	return
}

func (s *shortsStore) Delete(ctx context.Context, id string) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	short = &model.AudioShort{}
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New("cannot start transaction")
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error("failed to rollback")
			}
		}
	}()

	short, err = findOneByID(ctx, tx, id)
	if err != nil {
		return nil, errors.New("cannot find audio short with ID " + id)
	}
	err = deleteOne(ctx, tx, id)
	if err != nil {
		return nil, errors.New("failed to delete with ID " + id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("cannot commit transaction")
	}
	return
}

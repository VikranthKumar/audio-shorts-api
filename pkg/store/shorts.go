package store

import (
	"context"
	"database/sql"
	"sync"

	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/logging"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=shorts.go -destination=shorts_mock.go -package=store AudioShortsStore
type (
	AudioShortsStore interface {
		GetByID(ctx context.Context, id string) (short *model.AudioShort, err error)
		GetAll(ctx context.Context, page, limit uint16) (shorts []*model.AudioShort, err error)
		Create(ctx context.Context, input *model.AudioShortInput) (short *model.AudioShort, err error)
		Update(ctx context.Context, id string, input *model.AudioShortInput) (short *model.AudioShort, err error)
		Delete(ctx context.Context, id string) (short *model.AudioShort, err error)
		HardDelete(ctx context.Context, id string) (short *model.AudioShort, err error)
	}

	shortsStore struct {
		db *sql.DB
		sync.RWMutex
	}
)

func NewShortsStore(db *sql.DB) (AudioShortsStore, error) {
	return &shortsStore{
		db: db,
	}, nil
}

func (s *shortsStore) GetByID(ctx context.Context, id string) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageTransactionFailed)
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error(ErrorMessageRollbackFailed)
			}
		}
	}()

	short, err = findOneByID(ctx, tx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageFindFailed+" ID:"+id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCommitFailed)
	}
	return
}

func (s *shortsStore) GetAll(ctx context.Context, page, limit uint16) (shorts []*model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageTransactionFailed)
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error(ErrorMessageRollbackFailed)
			}
		}
	}()

	shorts, err = findAllShorts(ctx, tx, page, limit)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageFindFailed)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCommitFailed)
	}
	return
}

func (s *shortsStore) Create(ctx context.Context, input *model.AudioShortInput) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageTransactionFailed)
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error(ErrorMessageRollbackFailed)
			}
		}
	}()

	err = createOne(ctx, tx, input)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCreateFailed)
	}
	short, err = findOneByUnique(ctx, tx, input.Title, input.Creator.ID)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageFindFailed)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCommitFailed)
	}
	return
}

func (s *shortsStore) Update(ctx context.Context, id string, input *model.AudioShortInput) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageTransactionFailed)
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error(ErrorMessageRollbackFailed)
			}
		}
	}()

	err = updateOne(ctx, tx, id, input)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageUpdateFailed+" ID:"+id)
	}
	short, err = findOneByID(ctx, tx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageFindFailed+" ID:"+id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCommitFailed)
	}
	return
}

func (s *shortsStore) Delete(ctx context.Context, id string) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	short = &model.AudioShort{}
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageTransactionFailed)
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error(ErrorMessageRollbackFailed)
			}
		}
	}()

	err = softDeleteOne(ctx, tx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageDeleteFailed+" ID:"+id)
	}
	short, err = findOneByID(ctx, tx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageFindFailed+" ID:"+id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCommitFailed)
	}
	return
}

func (s *shortsStore) HardDelete(ctx context.Context, id string) (short *model.AudioShort, err error) {
	s.Lock()
	defer s.Unlock()

	short = &model.AudioShort{}
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageTransactionFailed)
	}

	defer func() {
		// evaluated when function returns
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logging.WithContext(ctx).Error(ErrorMessageRollbackFailed)
			}
		}
	}()

	short, err = findOneByID(ctx, tx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageFindFailed+" ID:"+id)
	}
	err = hardDeleteOne(ctx, tx, id)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageDeleteFailed+" ID:"+id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCommitFailed)
	}
	return
}

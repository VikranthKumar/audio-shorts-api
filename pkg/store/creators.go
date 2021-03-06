package store

import (
	"context"
	"database/sql"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/logging"
	"github.com/pkg/errors"
	"sync"
)

//go:generate mockgen -source=creators.go -destination=creators_mock.go -package=store CreatorsStore
type (
	CreatorsStore interface {
		GetAll(ctx context.Context, page, limit uint16) (shorts []*model.Creator, err error)
	}

	creatorsStore struct {
		db *sql.DB
		sync.RWMutex
	}
)

func NewCreatorsStore(db *sql.DB) (*creatorsStore, error) {
	return &creatorsStore{
		db: db,
	}, nil
}

func (s *creatorsStore) GetAll(ctx context.Context, page, limit uint16) (creators []*model.Creator, err error) {
	s.RLock()
	defer s.RUnlock()

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

	creators, err = findAllCreators(ctx, tx, page, limit)
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageFindFailed)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, ErrorMessageCommitFailed)
	}
	return
}

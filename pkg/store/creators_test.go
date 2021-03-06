package store

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nooble/task/audio-short-api/pkg/logging"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestCreatorsStore_GetAll(t *testing.T) {
	var (
		ID    = "1"
		name  = "hi"
		email = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := NewCreatorsStore(db)
	assert.NoError(t, err)

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT id, name, email FROM creators ORDER BY id ASC LIMIT $1 OFFSET $2")).
		WithArgs(1, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(ID, name, email))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.GetAll(ctx, 0, 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, ID, resp[0].ID)
	assert.Equal(t, name, resp[0].Name)
	assert.Equal(t, email, resp[0].Email)
}

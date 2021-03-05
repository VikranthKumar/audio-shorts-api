package store

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/logging"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestShortsStore_Create(t *testing.T) {
	var (
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		audioFile   = "a"
		creatorID   = "1"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := New(db)
	assert.NoError(t, err)

	input := &model.AudioShortInput{
		Title:       title,
		Description: description,
		Category:    category,
		AudioFile:   audioFile,
		Creator:     &model.CreatorInput{ID: creatorID},
	}

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(
		regexp.QuoteMeta("INSERT INTO audio_shorts( title, description, status, category, audio_file, creator_id ) VALUES ($1, $2, $3, $4, $5, $6 )")).
		WithArgs(title, description, "active", category, audioFile, creatorID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT a.id, a.title, a.description, a.category, a.audio_file, c.username, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.title = $1 AND a.creator_id = $2")).
		WithArgs(title, creatorID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "category", "audio_file", "username", "email"}).
			AddRow("1", title, description, category, audioFile, "itsme", "mockemail@gmail.com"))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.Create(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, "1", resp.ID)
	assert.Equal(t, title, resp.Title)
}

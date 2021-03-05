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

func TestShortsStore_GetByID(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		audioFile   = "a"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := New(db)
	assert.NoError(t, err)

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT a.title, a.description, a.category, a.audio_file, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1 AND a.status != 'deleted'")).
		WithArgs(ID).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description", "category", "audio_file", "username", "email"}).
			AddRow(title, description, category, audioFile, name, email))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.GetByID(ctx, ID)

	assert.NoError(t, err)
	assert.Equal(t, ID, resp.ID)
	assert.Equal(t, title, resp.Title)
	assert.Equal(t, name, resp.Creator.Name)
	assert.Equal(t, email, resp.Creator.Email)
}

func TestShortsStore_GetAll(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		audioFile   = "a"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := New(db)
	assert.NoError(t, err)

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT a.id, a.title, a.description, a.category, a.audio_file, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.status != 'deleted' ORDER BY a.id ASC LIMIT $1 OFFSET $2")).
		WithArgs(1, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "category", "audio_file", "username", "email"}).
			AddRow(ID, title, description, category, audioFile, name, email))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.GetAll(ctx, 0, 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, ID, resp[0].ID)
	assert.Equal(t, title, resp[0].Title)
	assert.Equal(t, name, resp[0].Creator.Name)
	assert.Equal(t, email, resp[0].Creator.Email)
}

func TestShortsStore_Create(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		status      = model.StatusActive
		audioFile   = "a"
		creatorID   = "1"
		name        = "hi"
		email       = "mockemail@gmail.com"
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
		WithArgs(title, description, status, category, audioFile, creatorID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT a.id, a.title, a.description, a.category, a.audio_file, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.title = $1 AND a.creator_id = $2 AND a.status != 'deleted'")).
		WithArgs(title, creatorID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "category", "audio_file", "name", "email"}).
			AddRow(ID, title, description, category, audioFile, name, email))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.Create(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, ID, resp.ID)
	assert.Equal(t, title, resp.Title)
	assert.Equal(t, name, resp.Creator.Name)
	assert.Equal(t, email, resp.Creator.Email)
}

func TestShortsStore_Update(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		audioFile   = "a"
		creatorID   = "1"
		name        = "hi"
		email       = "mockemail@gmail.com"
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
		regexp.QuoteMeta("UPDATE audio_shorts SET title = $1, description = $2, category = $3, audio_file = $4, creator_id = $5 WHERE id = $6")).
		WithArgs(title, description, category, audioFile, creatorID, ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT a.title, a.description, a.category, a.audio_file, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1 AND a.status != 'deleted'")).
		WithArgs(ID).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description", "category", "audio_file", "name", "email"}).
			AddRow(title, description, category, audioFile, name, email))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.Update(ctx, ID, input)

	assert.NoError(t, err)
	assert.Equal(t, ID, resp.ID)
	assert.Equal(t, title, resp.Title)
	assert.Equal(t, name, resp.Creator.Name)
	assert.Equal(t, email, resp.Creator.Email)
}

func TestShortsStore_Delete(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		status      = model.StatusDeleted
		audioFile   = "a"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := New(db)
	assert.NoError(t, err)

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT a.title, a.description, a.category, a.audio_file, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1 AND a.status != 'deleted'")).
		WithArgs(ID).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description", "category", "audio_file", "name", "email"}).
			AddRow(title, description, category, audioFile, name, email))
	sqlMock.ExpectExec(
		regexp.QuoteMeta("UPDATE audio_shorts SET status = $1 WHERE id = $2")).
		WithArgs(status, ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.Delete(ctx, ID)

	assert.NoError(t, err)
	assert.Equal(t, ID, resp.ID)
	assert.Equal(t, title, resp.Title)
	assert.Equal(t, name, resp.Creator.Name)
	assert.Equal(t, email, resp.Creator.Email)
}

func TestShortsStore_HardDelete(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		audioFile   = "a"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := New(db)
	assert.NoError(t, err)

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(
		regexp.QuoteMeta("SELECT a.title, a.description, a.category, a.audio_file, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1 AND a.status != 'deleted'")).
		WithArgs(ID).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description", "category", "audio_file", "name", "email"}).
			AddRow(title, description, category, audioFile, name, email))
	sqlMock.ExpectExec(
		regexp.QuoteMeta("DELETE FROM audio_shorts WHERE id = $1")).
		WithArgs(ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	ctx := logging.NewContext(context.Background())
	resp, err := store.HardDelete(ctx, ID)

	assert.NoError(t, err)
	assert.Equal(t, ID, resp.ID)
	assert.Equal(t, title, resp.Title)
	assert.Equal(t, name, resp.Creator.Name)
	assert.Equal(t, email, resp.Creator.Email)
}

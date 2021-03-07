package store

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/logging"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestShortsStore_GetByID(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		status      = model.StatusActive
		category    = model.CategoryNews
		audioFile   = "a"
		creatorID   = "1"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := NewShortsStore(db)
	assert.NoError(t, err)

	t.Run("happy path", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1")).
			WithArgs(ID).
			WillReturnRows(sqlmock.NewRows([]string{"title", "description", "status", "category", "audio_file", "id", "name", "email"}).
				AddRow(title, description, status, category, audioFile, creatorID, name, email)).RowsWillBeClosed()
		sqlMock.ExpectCommit()

		ctx := logging.NewContext(context.Background())
		resp, err := store.GetByID(ctx, ID)

		assert.NoError(t, err)
		assert.Equal(t, ID, resp.ID)
		assert.Equal(t, title, resp.Title)
		assert.Equal(t, name, resp.Creator.Name)
		assert.Equal(t, email, resp.Creator.Email)
	})

	t.Run("sad path - no rows", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1")).
			WithArgs(ID).
			WillReturnError(sql.ErrNoRows)
		sqlMock.ExpectRollback()

		ctx := logging.NewContext(context.Background())
		resp, err := store.GetByID(ctx, ID)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestShortsStore_GetAll(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		status      = model.StatusActive
		category    = model.CategoryNews
		audioFile   = "a"
		creatorID   = "1"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := NewShortsStore(db)
	assert.NoError(t, err)

	t.Run("happy path", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.id, a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id ORDER BY a.id ASC LIMIT $1 OFFSET $2")).
			WithArgs(1, 0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "category", "audio_file", "id", "name", "email"}).
				AddRow(ID, title, description, status, category, audioFile, creatorID, name, email)).RowsWillBeClosed()
		sqlMock.ExpectCommit()

		ctx := logging.NewContext(context.Background())
		resp, err := store.GetAll(ctx, 0, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(resp))
		assert.Equal(t, ID, resp[0].ID)
		assert.Equal(t, title, resp[0].Title)
		assert.Equal(t, name, resp[0].Creator.Name)
		assert.Equal(t, email, resp[0].Creator.Email)
	})

	t.Run("sad path - no rows", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.id, a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id ORDER BY a.id ASC LIMIT $1 OFFSET $2")).
			WithArgs(1, 0).
			WillReturnError(sql.ErrNoRows)
		sqlMock.ExpectCommit()

		ctx := logging.NewContext(context.Background())
		resp, err := store.GetAll(ctx, 0, 1)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
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
	store, err := NewShortsStore(db)
	assert.NoError(t, err)

	input := &model.AudioShortInput{
		Title:       title,
		Description: description,
		Category:    category,
		AudioFile:   audioFile,
		Creator:     &model.CreatorInput{ID: creatorID},
	}

	t.Run("happy path", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(
			regexp.QuoteMeta("INSERT INTO audio_shorts( title, description, status, category, audio_file, creator_id ) VALUES ($1, $2, $3, $4, $5, $6 )")).
			WithArgs(title, description, status, category, audioFile, creatorID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.id, a.title, a.description, a.status, a.category, a.audio_file, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.title = $1 AND a.creator_id = $2")).
			WithArgs(title, creatorID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "status", "category", "audio_file", "name", "email"}).
				AddRow(ID, title, description, status, category, audioFile, name, email))
		sqlMock.ExpectCommit()

		ctx := logging.NewContext(context.Background())
		resp, err := store.Create(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, ID, resp.ID)
		assert.Equal(t, title, resp.Title)
		assert.Equal(t, name, resp.Creator.Name)
		assert.Equal(t, email, resp.Creator.Email)
	})

	t.Run("sad path - failed insert", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(
			regexp.QuoteMeta("INSERT INTO audio_shorts( title, description, status, category, audio_file, creator_id ) VALUES ($1, $2, $3, $4, $5, $6 )")).
			WithArgs(title, description, status, category, audioFile, creatorID).
			WillReturnError(errors.New("some error"))
		sqlMock.ExpectRollback()

		ctx := logging.NewContext(context.Background())
		resp, err := store.Create(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestShortsStore_Update(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		status      = model.StatusActive
		category    = model.CategoryNews
		audioFile   = "a"
		creatorID   = "1"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := NewShortsStore(db)
	assert.NoError(t, err)

	input := &model.AudioShortInput{
		Title:       title,
		Description: description,
		Category:    category,
		AudioFile:   audioFile,
		Creator:     &model.CreatorInput{ID: creatorID},
	}

	t.Run("happy path", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(
			regexp.QuoteMeta("UPDATE audio_shorts SET title = $1, description = $2, category = $3, audio_file = $4, creator_id = $5 WHERE id = $6")).
			WithArgs(title, description, category, audioFile, creatorID, ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1")).
			WithArgs(ID).
			WillReturnRows(sqlmock.NewRows([]string{"title", "description", "status", "category", "audio_file", "id", "name", "email"}).
				AddRow(title, description, status, category, audioFile, creatorID, name, email))
		sqlMock.ExpectCommit()

		ctx := logging.NewContext(context.Background())
		resp, err := store.Update(ctx, ID, input)

		assert.NoError(t, err)
		assert.Equal(t, ID, resp.ID)
		assert.Equal(t, title, resp.Title)
		assert.Equal(t, name, resp.Creator.Name)
		assert.Equal(t, email, resp.Creator.Email)
	})

	t.Run("sad path - failed update", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(
			regexp.QuoteMeta("UPDATE audio_shorts SET title = $1, description = $2, category = $3, audio_file = $4, creator_id = $5 WHERE id = $6")).
			WithArgs(title, description, category, audioFile, creatorID, ID).
			WillReturnError(errors.New("some error"))
		sqlMock.ExpectRollback()

		ctx := logging.NewContext(context.Background())
		resp, err := store.Update(ctx, ID, input)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestShortsStore_Delete(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		category    = model.CategoryNews
		status      = model.StatusDeleted
		audioFile   = "a"
		creatorID   = "1"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := NewShortsStore(db)
	assert.NoError(t, err)

	t.Run("happy path", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(
			regexp.QuoteMeta("UPDATE audio_shorts SET status = $1 WHERE id = $2")).
			WithArgs(status, ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1")).
			WithArgs(ID).
			WillReturnRows(sqlmock.NewRows([]string{"title", "description", "status", "category", "audio_file", "id", "name", "email"}).
				AddRow(title, description, status, category, audioFile, creatorID, name, email))
		sqlMock.ExpectCommit()

		ctx := logging.NewContext(context.Background())
		resp, err := store.Delete(ctx, ID)

		assert.NoError(t, err)
		assert.Equal(t, ID, resp.ID)
		assert.Equal(t, title, resp.Title)
		assert.Equal(t, name, resp.Creator.Name)
		assert.Equal(t, email, resp.Creator.Email)
	})

	t.Run("sad path - failed delete", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec(
			regexp.QuoteMeta("UPDATE audio_shorts SET status = $1 WHERE id = $2")).
			WithArgs(status, ID).
			WillReturnError(errors.New("some error"))
		sqlMock.ExpectRollback()

		ctx := logging.NewContext(context.Background())
		resp, err := store.Delete(ctx, ID)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestShortsStore_HardDelete(t *testing.T) {
	var (
		ID          = "1"
		title       = "abc"
		description = "abcs"
		status      = model.StatusActive
		category    = model.CategoryNews
		audioFile   = "a"
		creatorID   = "1"
		name        = "hi"
		email       = "mockemail@gmail.com"
	)
	db, sqlMock, err := sqlmock.New()
	assert.NoError(t, err)
	store, err := NewShortsStore(db)
	assert.NoError(t, err)

	t.Run("happy path", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1")).
			WithArgs(ID).
			WillReturnRows(sqlmock.NewRows([]string{"title", "description", "status", "category", "audio_file", "id", "name", "email"}).
				AddRow(title, description, status, category, audioFile, creatorID, name, email))
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
	})

	t.Run("sad path - failed delete", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(
			regexp.QuoteMeta("SELECT a.title, a.description, a.status, a.category, a.audio_file, c.id, c.name, c.email FROM audio_shorts AS a,creators AS c WHERE c.id = a.creator_id AND a.id = $1")).
			WithArgs(ID).
			WillReturnRows(sqlmock.NewRows([]string{"title", "description", "status", "category", "audio_file", "id", "name", "email"}).
				AddRow(title, description, status, category, audioFile, creatorID, name, email))
		sqlMock.ExpectExec(
			regexp.QuoteMeta("DELETE FROM audio_shorts WHERE id = $1")).
			WithArgs(ID).
			WillReturnError(errors.New("some error"))
		sqlMock.ExpectRollback()

		ctx := logging.NewContext(context.Background())
		resp, err := store.HardDelete(ctx, ID)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

package store

import (
	"context"
	"database/sql"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
)

func findOneByID(ctx context.Context, tx *sql.Tx, id string) (short *model.AudioShort, err error) {
	var (
		title       string
		description string
		category    string
		audioFile   string
		name        string
		email       string
	)
	query := "SELECT " +
		"a.title, " +
		"a.description, " +
		"a.category, " +
		"a.audio_file, " +
		"c.name, " +
		"c.email " +
		"FROM audio_shorts AS a," +
		"creators AS c " +
		"WHERE " +
		"c.id = a.creator_id " +
		"AND a.id = $1"

	row := tx.QueryRowContext(ctx, query, id)
	err = row.Scan(&title, &description, &category, &audioFile, &name, &email)
	short = &model.AudioShort{
		ID:          id,
		Title:       title,
		Description: description,
		Category:    model.Category(category),
		AudioFile:   audioFile,
		Creator: &model.Creator{
			Name:  name,
			Email: email,
		},
	}
	return
}

func findOneByUnique(ctx context.Context, tx *sql.Tx, inputTitle string, creatorID string) (short *model.AudioShort, err error) {
	var (
		id          string
		title       string
		description string
		category    string
		audioFile   string
		name        string
		email       string
	)
	query := "SELECT " +
		"a.id, " +
		"a.title, " +
		"a.description, " +
		"a.category, " +
		"a.audio_file, " +
		"c.name, " +
		"c.email " +
		"FROM audio_shorts AS a," +
		"creators AS c " +
		"WHERE " +
		"c.id = a.creator_id " +
		"AND a.title = $1 " +
		"AND a.creator_id = $2"

	row := tx.QueryRowContext(ctx, query, inputTitle, creatorID)
	err = row.Scan(&id, &title, &description, &category, &audioFile, &name, &email)
	short = &model.AudioShort{
		ID:          id,
		Title:       title,
		Description: description,
		Category:    model.Category(category),
		AudioFile:   audioFile,
		Creator: &model.Creator{
			Name:  name,
			Email: email,
		},
	}
	return
}

func findAll(ctx context.Context, tx *sql.Tx, page, limit uint16) (shorts []*model.AudioShort, err error) {
	shorts = make([]*model.AudioShort, 0, limit) // set cap at limit
	var (
		id          string
		title       string
		description string
		category    string
		audioFile   string
		name        string
		email       string
	)
	query := "SELECT " +
		"a.id, " +
		"a.title, " +
		"a.description, " +
		"a.category, " +
		"a.audio_file, " +
		"c.name, " +
		"c.email " +
		"FROM audio_shorts AS a," +
		"creators AS c " +
		"WHERE " +
		"c.id = a.creator_id " +
		"ORDER BY a.id ASC " +
		"LIMIT $1 " +
		"OFFSET $2"

	rows, err := tx.QueryContext(ctx, query, limit, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &title, &description, &category, &audioFile, &name, &email)
		if err != nil {
			return nil, err
		}
		short := &model.AudioShort{
			ID:          id,
			Title:       title,
			Description: description,
			Category:    model.Category(category),
			AudioFile:   audioFile,
			Creator: &model.Creator{
				Name:  name,
				Email: email,
			},
		}
		shorts = append(shorts, short)
	}
	return
}

func createOne(ctx context.Context, tx *sql.Tx, input *model.AudioShortInput) (err error) {
	query := "INSERT INTO " +
		"audio_shorts( " +
		"title, " +
		"description, " +
		"status, " +
		"category, " +
		"audio_file, " +
		"creator_id " +
		") VALUES (" +
		"$1, " +
		"$2, " +
		"$3, " + // set default status as 'active'
		"$4, " +
		"$5, " +
		"$6 " +
		")"

	_, err = tx.ExecContext(ctx, query, input.Title, input.Description, "active", input.Category.String(), input.AudioFile, input.Creator.ID)
	return
}

func updateOne(ctx context.Context, tx *sql.Tx, id string, input *model.AudioShortInput) (err error) {
	query := "UPDATE " +
		"audio_shorts " +
		"SET " +
		"title = $1, " +
		"description = $2, " +
		"category = $3, " +
		"audio_file = $4, " +
		"creator_id = $5 " +
		"WHERE id = $6"

	_, err = tx.ExecContext(ctx, query, input.Title, input.Description, input.Category.String(), input.AudioFile, input.Creator.ID, id)
	return
}

func deleteOne(ctx context.Context, tx *sql.Tx, id string) (err error) {
	query := "DELETE FROM " +
		"audio_shorts " +
		"WHERE id = $1"

	_, err = tx.ExecContext(ctx, query, id)
	return
}

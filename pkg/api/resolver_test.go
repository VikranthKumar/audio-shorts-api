package api

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang/mock/gomock"
	"github.com/nooble/task/audio-short-api/pkg/api/generated"
	"github.com/nooble/task/audio-short-api/pkg/api/model"
	"github.com/nooble/task/audio-short-api/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolver_Query(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockAudioShortsStore(ctrl)
	resolver, err := New(mockStore)
	assert.NoError(t, err)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

	short := &model.AudioShort{
		ID:          "1",
		Title:       "abc",
		Description: "abcs",
		Category:    model.CategoryNews,
		AudioFile:   "a",
		Creator:     &model.Creator{},
	}
	mockStore.EXPECT().GetByID(gomock.Any(), "1").Return(short, nil)

	var resp struct {
		GetAudioShort struct{ Title, Description string }
	}
	q := `
		query {
			getAudioShort(id: "1") {
				title,
				description
			}
		}`
	c.MustPost(q, &resp)
	assert.Equal(t, "abc", resp.GetAudioShort.Title)
	assert.Equal(t, "abcs", resp.GetAudioShort.Description)
}

func TestResolver_Mutation(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockAudioShortsStore(ctrl)
	resolver, err := New(mockStore)
	assert.NoError(t, err)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

	short := &model.AudioShort{
		ID:          "1",
		Title:       "abc",
		Description: "abcs",
		Category:    model.CategoryNews,
		AudioFile:   "a",
		Creator:     &model.Creator{},
	}
	input := &model.AudioShortInput{
		Title:       "abc",
		Description: "abcs",
		Category:    model.CategoryNews,
		AudioFile:   "a",
		Creator:     &model.CreatorInput{ID: "1"},
	}
	mockStore.EXPECT().Update(gomock.Any(), "1", input).Return(short, nil)

	var resp struct {
		UpdateAudioShort struct{ Title, Description string }
	}
	m := `
		mutation {
			updateAudioShort(id: "1", input: {
				title: "abc",
				description: "abcs",
				category: story,
				audio_file: "a",
				creator: {
					id: "1"
				}
			}) {
				title,
				description
			}
		}`
	c.MustPost(m, &resp)
	assert.Equal(t, "abc", resp.UpdateAudioShort.Title)
	assert.Equal(t, "abcs", resp.UpdateAudioShort.Description)
}

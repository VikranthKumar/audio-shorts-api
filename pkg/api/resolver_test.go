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

func TestQueryResolver_GetAudioShort(t *testing.T) {
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

	t.Run("happy path", func(t *testing.T) {
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
	})

	t.Run("sad path - no args", func(t *testing.T) {
		var resp struct {
			GetAudioShort struct{ Title, Description string }
		}
		q := `
		query {
			getAudioShort {
				title,
				description
			}
		}`
		assert.Panics(t, func() {
			c.MustPost(q, &resp)
		})
	})
}

func TestQueryResolver_GetAudioShorts(t *testing.T) {
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

	t.Run("happy path", func(t *testing.T) {
		mockStore.EXPECT().GetAll(gomock.Any(), uint16(0), uint16(1)).Return([]*model.AudioShort{short}, nil)
		var resp struct {
			GetAudioShorts []struct{ Title, Description string }
		}
		q := `
		query {
			getAudioShorts(page: 1, limit: 1) {
				title,
				description
			}
		}`
		c.MustPost(q, &resp)
		assert.Equal(t, "abc", resp.GetAudioShorts[0].Title)
		assert.Equal(t, "abcs", resp.GetAudioShorts[0].Description)
	})

	t.Run("happy path - default args", func(t *testing.T) {
		mockStore.EXPECT().GetAll(gomock.Any(), uint16(0), uint16(10)).Return([]*model.AudioShort{short}, nil)
		var resp struct {
			GetAudioShorts []struct{ Title, Description string }
		}
		q := `
		query {
			getAudioShorts {
				title,
				description
			}
		}`
		c.MustPost(q, &resp)
		assert.Equal(t, "abc", resp.GetAudioShorts[0].Title)
		assert.Equal(t, "abcs", resp.GetAudioShorts[0].Description)
	})

	t.Run("sad path - page is 0", func(t *testing.T) {
		var resp struct {
			GetAudioShorts []struct{ Title, Description string }
		}
		q := `
		query {
			getAudioShorts(page: 0, limit: 1) {
				title,
				description
			}
		}`
		assert.Panics(t, func() {
			c.MustPost(q, &resp)
		})
	})
}

func TestMutationResolver_UpdateAudioShort(t *testing.T) {
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
				category: news,
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

func TestMutationResolver_CreateAudioShort(t *testing.T) {
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
	mockStore.EXPECT().Create(gomock.Any(), input).Return(short, nil)

	var resp struct {
		CreateAudioShort struct{ Title, Description string }
	}
	m := `
		mutation {
			createAudioShort(input: {
				title: "abc",
				description: "abcs",
				category: news,
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
	assert.Equal(t, "abc", resp.CreateAudioShort.Title)
	assert.Equal(t, "abcs", resp.CreateAudioShort.Description)
}

func TestMutationResolver_DeleteAudioShort(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockAudioShortsStore(ctrl)
	resolver, err := New(mockStore)
	assert.NoError(t, err)
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

	ID := "1"
	short := &model.AudioShort{
		ID:          "1",
		Title:       "abc",
		Description: "abcs",
		Category:    model.CategoryNews,
		AudioFile:   "a",
		Creator:     &model.Creator{},
	}
	mockStore.EXPECT().Delete(gomock.Any(), ID).Return(short, nil)

	var resp struct {
		DeleteAudioShort struct{ Title, Description string }
	}
	m := `
		mutation {
			deleteAudioShort(id: "1")
			{
				title,
				description
			}
		}`
	c.MustPost(m, &resp)
	assert.Equal(t, "abc", resp.DeleteAudioShort.Title)
	assert.Equal(t, "abcs", resp.DeleteAudioShort.Description)
}

func TestMutationResolver_HardDeleteAudioShort(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockAudioShortsStore(ctrl)
	resolver, err := New(mockStore)
	assert.NoError(t, err)

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})))

	ID := "1"
	short := &model.AudioShort{
		ID:          "1",
		Title:       "abc",
		Description: "abcs",
		Category:    model.CategoryNews,
		AudioFile:   "a",
		Creator:     &model.Creator{},
	}
	mockStore.EXPECT().HardDelete(gomock.Any(), ID).Return(short, nil)

	var resp struct {
		HardDeleteAudioShort struct{ Title, Description string }
	}
	m := `
		mutation {
			hardDeleteAudioShort(id: "1")
			{
				title,
				description
			}
		}`
	c.MustPost(m, &resp)
	assert.Equal(t, "abc", resp.HardDeleteAudioShort.Title)
	assert.Equal(t, "abcs", resp.HardDeleteAudioShort.Description)
}

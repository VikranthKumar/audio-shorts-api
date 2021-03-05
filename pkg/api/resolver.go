package api

import "github.com/nooble/task/audio-short-api/pkg/store"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	store store.AudioShortsStore
}

func New(store store.AudioShortsStore) (*Resolver, error) {
	return &Resolver{store: store}, nil
}

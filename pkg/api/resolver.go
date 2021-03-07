package api

import "github.com/nooble/task/audio-short-api/pkg/store"

//go:generate go run github.com/99designs/gqlgen

// Resolver has reference to shortsStore and creatorsStore
type Resolver struct {
	shortsStore   store.AudioShortsStore
	creatorsStore store.CreatorsStore
}

func New(shortsStore store.AudioShortsStore, creatorsStore store.CreatorsStore) (*Resolver, error) {
	return &Resolver{shortsStore: shortsStore, creatorsStore: creatorsStore}, nil
}

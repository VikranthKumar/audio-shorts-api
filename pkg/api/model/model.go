// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AudioShort struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      Status   `json:"status"`
	Category    Category `json:"category"`
	AudioFile   string   `json:"audio_file"`
	Creator     *Creator `json:"creator"`
}

type AudioShortInput struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Category    Category      `json:"category"`
	AudioFile   string        `json:"audio_file"`
	Creator     *CreatorInput `json:"creator"`
}

type Creator struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreatorInput struct {
	ID string `json:"id"`
}

type Category string

const (
	CategoryNews   Category = "news"
	CategoryGossip Category = "gossip"
	CategoryReview Category = "review"
	CategoryStory  Category = "story"
)

var AllCategory = []Category{
	CategoryNews,
	CategoryGossip,
	CategoryReview,
	CategoryStory,
}

func (e Category) IsValid() bool {
	switch e {
	case CategoryNews, CategoryGossip, CategoryReview, CategoryStory:
		return true
	}
	return false
}

func (e Category) String() string {
	return string(e)
}

func (e *Category) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Category(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}

func (e Category) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusActive  Status = "active"
	StatusBanned  Status = "banned"
	StatusDeleted Status = "deleted"
)

var AllStatus = []Status{
	StatusActive,
	StatusBanned,
	StatusDeleted,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusActive, StatusBanned, StatusDeleted:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

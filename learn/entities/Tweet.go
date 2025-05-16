package entities

import "github.com/google/uuid"

type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New().String(),
	}
	return &tweet

}

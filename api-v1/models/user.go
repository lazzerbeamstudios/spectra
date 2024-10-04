package models

import (
	"time"
)

type Profile struct {
	ID      int       `json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type ProfileUpdate struct {
	_    struct{} `json:"-" additionalProperties:"true"`
	ID   int      `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
}

type ProfileUpdateEmail struct {
	_     struct{} `json:"-" additionalProperties:"true"`
	ID    int      `json:"id,omitempty"`
	Email string   `json:"email,omitempty"`
}

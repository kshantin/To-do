package models

import "time"

type Note struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Name    string `json:"name"`
	NumNote int    `json:"num_note"`
}

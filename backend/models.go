package main

import (
	"time"
)

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description *string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"create_at"`
}

var ValidStatuses = map[string]bool{
	"todo": true,
	"in_progress": true,
	"done": true,
}
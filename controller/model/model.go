package model

import "time"

type Model struct {
	Username string `json:"username"`
	Level int `json:"level"`
	Points []int `json:"points"`
	Total int `json:"total"`
	Date time.Time `json:"date"`
}

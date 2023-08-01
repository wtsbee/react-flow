package model

import "time"

type DataDetails struct {
	Label string `json:"label"`
}
type Node struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type"`
	X         float64   `json:"x"`
	Y         float64   `json:"y"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NodeResponse struct {
	ID       string          `json:"id"`
	Type     string          `json:"type"`
	Position PositionDetails `json:"position"`
	Data     DataDetails     `json:"data"`
}

type PositionDetails struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

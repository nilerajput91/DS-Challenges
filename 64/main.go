package main

import "time"

type Order struct {
	ID        string   `json:"id"`
	Item      []string `json:"item"`
	TotalCost float32  `json:"totalCost"`
	CreatedAt time.Time `json:"createdAt"`
}

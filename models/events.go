package models

import "time"

type Events struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	DateTime    time.Time
	UserID      int `json:"userid"`
}

var Event = []Events{
	{ID: 1, Name: "Blue Train", Description: "John Coltrane", Location: "lekki", UserID: 1},
	{ID: 2, Name: "train", Description: "Coltrane", Location: "lekki", UserID: 2},
	{ID: 3, Name: "Blue", Description: "John", Location: "lekki", UserID: 3},
}

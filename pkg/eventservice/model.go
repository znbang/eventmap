package eventservice

import (
	"time"
)

type EventInput struct {
	ID        string    `json:"id" param:"id"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	URL       string    `json:"url"`
	Detail    string    `json:"detail"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Zoom      int       `json:"zoom"`
}

type Event struct {
	ID        string    `json:"id"`
	UserID    string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	URL       string    `json:"url"`
	Detail    string    `json:"detail"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Zoom      int       `json:"zoom"`
}

type EventList struct {
	Page  int     `json:"page"`
	Size  int     `json:"size"`
	Total int     `json:"total"`
	Items []Event `json:"items"`
}

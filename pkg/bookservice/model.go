package bookservice

import (
	"strings"
	"time"
)

type BookInput struct {
	ID     string `json:"id" param:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

type Book struct {
	ID        string    `json:"id" param:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserID    string    `json:"-"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	URL       string    `json:"url"`
	Job       *BookJob  `json:"job"`
}

type BookList struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int    `json:"total"`
	Items []Book `json:"items"`
}

type Chapter struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	BookID    string    `json:"-"`
	Page      int       `json:"page"`
	URL       string    `json:"-"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

func (c *Chapter) Lines() []string {
	return strings.Split(c.Body, "\n")
}

type ChapterList struct {
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Total int       `json:"total"`
	Items []Chapter `json:"items"`
}

type TocView struct {
	Book     Book        `json:"book"`
	Chapters ChapterList `json:"chapters"`
}

type ChapterView struct {
	Book    Book    `json:"book"`
	Chapter Chapter `json:"chapter"`
	Page    int     `json:"page"`
	Total   int     `json:"total"`
}

const StatusDone = 0
const StatusPending = 1
const StatusRunning = 2

type BookJob struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Status    int       `json:"status"`
	BookID    string    `json:"bookId"`
	Message   string    `json:"message"`
}

func (j *BookJob) IsPending() bool {
	return j.Status == StatusPending
}

func (j *BookJob) IsRunning() bool {
	return j.Status == StatusRunning
}

func (j *BookJob) IsDone() bool {
	return j.Status == StatusDone
}

type BookJobList struct {
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Total int       `json:"total"`
	Items []BookJob `json:"items"`
}

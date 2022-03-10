package bookservice

import "errors"

var (
	ErrNoSuchBook    = errors.New("no such book")
	ErrNoSuchBookJob = errors.New("no such book job")
	ErrNoSuchChapter = errors.New("no such chapter")
)

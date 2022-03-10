package eventservice

import "time"

type EventRepository interface {
	FindByEndDateAfter(events *[]Event, endDate time.Time) error
	CountAll(total *int, q string) error
	FindAll(events *[]Event, q string, page int, size int) error
	CountByUserId(total *int, userId string, q string) error
	FindByUserId(events *[]Event, userId string, q string, page int, size int) error
	FindById(event *Event, id string) error
	FindByIdAndUserId(event *Event, id string, userId string) error
	DeleteByIdAndUserId(id string, userId string) error
	Create(event *Event) error
	Save(event *Event) error
}

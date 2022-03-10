package eventservice

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/znbang/eventmap/pkg/uuid"
	"time"
)

type EventService struct {
	EventRepository
}

func NewEventService(eventRepository EventRepository) *EventService {
	return &EventService{
		EventRepository: eventRepository,
	}
}

func (s *EventService) GetActiveEvents(events *EventList) error {
	var items = make([]Event, 0)

	d := time.Now()
	d = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())

	if err := s.EventRepository.FindByEndDateAfter(&items, d); err != nil {
		return err
	}

	events.Page = 1
	events.Size = len(items)
	events.Total = 1
	events.Items = items

	return nil
}

func (s *EventService) GetEvents(events *EventList, q string, page int, size int) error {
	var (
		total int
		items []Event
	)

	if err := s.EventRepository.CountAll(&total, q); err != nil {
		return err
	}

	if err := s.EventRepository.FindAll(&items, q, page, size); err != nil {
		return err
	}

	events.Page = page
	events.Size = size
	events.Total = (total + size - 1) / size
	events.Items = items

	return nil
}

func (s *EventService) GetEventsByUser(events *EventList, userId string, q string, page int, size int) error {
	var (
		total int
		items []Event
	)

	if err := s.EventRepository.CountByUserId(&total, userId, q); err != nil {
		return err
	}

	if err := s.EventRepository.FindByUserId(&items, userId, q, page, size); err != nil {
		return err
	}

	events.Page = page
	events.Size = size
	events.Total = (total + size - 1) / size
	events.Items = items

	return nil
}

func (s *EventService) CreateEvent(userId string, input *Event) error {
	p := bluemonday.UGCPolicy()
	input.Name = p.Sanitize(input.Name)
	input.URL = p.Sanitize(input.URL)
	input.Location = p.Sanitize(input.Location)
	input.Detail = p.Sanitize(input.Detail)
	input.ID = uuid.ULID()
	input.UserID = userId

	return s.EventRepository.Create(input)
}

func (s *EventService) UpdateEvent(userId string, input Event) error {
	var event Event

	if err := s.EventRepository.FindByIdAndUserId(&event, input.ID, userId); err != nil {
		return err
	}

	p := bluemonday.UGCPolicy()
	event.Name = p.Sanitize(input.Name)
	event.URL = p.Sanitize(input.URL)
	event.Location = p.Sanitize(input.Location)
	event.Detail = p.Sanitize(input.Detail)
	event.Lat = input.Lat
	event.Lng = input.Lng
	event.Zoom = input.Zoom
	// PostgreSQL
	// event.StartDate = time.Date(input.StartDate.Year(), input.StartDate.Month(), input.StartDate.Day(), 0, 0, 0, 0, input.StartDate.Location())
	// event.EndDate = time.Date(input.EndDate.Year(), input.EndDate.Month(), input.EndDate.Day(), 0, 0, 0, 0, input.EndDate.Location())
	event.StartDate = input.StartDate
	event.EndDate = input.EndDate

	return s.EventRepository.Save(&event)
}

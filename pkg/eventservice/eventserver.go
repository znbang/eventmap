package eventservice

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/mitchellh/mapstructure"
	v1 "github.com/znbang/eventmap/gen/event/v1"
	"github.com/znbang/eventmap/internal/mvc"
	"github.com/znbang/eventmap/pkg/auth"
	"github.com/znbang/eventmap/pkg/userservice"
	"github.com/znbang/validation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventServer struct {
	eventService *EventService
}

func NewEventServer(eventService *EventService) *EventServer {
	return &EventServer{
		eventService: eventService,
	}
}

func (s *EventServer) CreateEvent(ctx context.Context, r *connect.Request[v1.CreateEventRequest]) (*connect.Response[v1.CreateEventResponse], error) {
	var (
		user     userservice.User
		event Event
		validate = validation.New()
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if err := mapstructure.Decode(r.Msg, &event); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	event.StartDate = r.Msg.StartDate.AsTime()
	event.EndDate = r.Msg.EndDate.AsTime()

	if err := validateCreateEvent(validate, event); err != nil {
		return nil, err
	}

	if err := s.eventService.CreateEvent(user.ID, &event); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.CreateEventResponse{
		Id: event.ID,
	}), nil
}

func validateCreateEvent(validate *validation.Validation, input Event) error {
	if validate.Required(input.Name, "name", "events.required.name") {
		validate.Max(input.Name, 255, "name", "events.size.name")
	}
	validate.Required(input.StartDate, "startDate", "events.required.date")
	if input.StartDate.Unix() == 0 {
		validate.Errors["startDate"] = "events.required.date"
	}
	validate.Required(input.EndDate, "endDate", "events.required.date")
	if input.EndDate.Unix() == 0 {
		validate.Errors["endDate"] = "events.required.date"
	}
	if validate.Required(input.Location, "location", "events.required.location") {
		validate.Max(input.Location, 255, "location", "events.size.location")
	}
	validate.Required(input.Lat, "lat", "events.required.location")
	validate.Required(input.Lng, "lng", "events.required.location")
	validate.Required(input.Zoom, "zoom", "events.required.zoom")
	validate.Required(input.Detail, "detail", "events.required.detail")
	if len(input.URL) > 0 {
		validate.URL(input.URL, "url", "events.invalid.url")
	}

	if validate.HasErrors() {
		return mvc.NewValidationError(validate)
	}

	return nil
}

func (s *EventServer) UpdateEvent(ctx context.Context, r *connect.Request[v1.UpdateEventRequest]) (*connect.Response[v1.UpdateEventResponse], error) {
	var (
		user     userservice.User
		event    Event
		validate = validation.New()
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if err := mapstructure.Decode(r.Msg, &event); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	event.StartDate = r.Msg.StartDate.AsTime()
	event.EndDate = r.Msg.EndDate.AsTime()

	if err := validateUpdateEvent(validate, event); err != nil {
		return nil, err
	}

	if err := s.eventService.UpdateEvent(user.ID, event); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.UpdateEventResponse{
		Id: event.ID,
	}), nil
}

func validateUpdateEvent(validate *validation.Validation, input Event) error {
	validate.Required(input.ID, "id", "events.id.date")
	return validateCreateEvent(validate, input)
}

func (s *EventServer) DeleteEvent(ctx context.Context, r *connect.Request[v1.DeleteEventRequest]) (*connect.Response[v1.DeleteEventResponse], error) {
	var user userservice.User

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.eventService.eventRepository.DeleteByIdAndUserId(r.Msg.Id, user.ID); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteEventResponse{}), nil
}

func (s *EventServer) GetEvent(ctx context.Context, r *connect.Request[v1.GetEventRequest]) (*connect.Response[v1.GetEventResponse], error) {
	var event Event

	if err := s.eventService.eventRepository.FindById(&event, r.Msg.Id); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.GetEventResponse{
		Event: convertEventToPB(event),
	}), nil
}

func (s *EventServer) ListActiveEvent(ctx context.Context, r *connect.Request[v1.ListActiveEventRequest]) (*connect.Response[v1.ListActiveEventResponse], error) {
	var events EventList

	if err := s.eventService.GetActiveEvents(&events); err != nil {
		return nil, err
	}

	resp := connect.NewResponse(&v1.ListActiveEventResponse{
		Page:  int32(events.Page),
		Size:  int32(events.Size),
		Total: int32(events.Total),
		Items: convertEventItems(events.Items),
	})

	return resp, nil
}

func (s *EventServer) ListEvent(ctx context.Context, r *connect.Request[v1.ListEventRequest]) (*connect.Response[v1.ListEventResponse], error) {
	var (
		events     EventList
		page, size = mvc.PageSize(r.Msg.Page, r.Msg.Size)
	)

	if err := s.eventService.GetEvents(&events, r.Msg.Filter, page, size); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ListEventResponse{
		Page:  int32(events.Page),
		Size:  int32(events.Size),
		Total: int32(events.Total),
		Items: convertEventItems(events.Items),
	}), nil
}

func (s *EventServer) ListUserEvent(ctx context.Context, r *connect.Request[v1.ListUserEventRequest]) (*connect.Response[v1.ListUserEventResponse], error) {
	var (
		user       userservice.User
		events     EventList
		page, size = mvc.PageSize(r.Msg.Page, r.Msg.Size)
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.eventService.GetEventsByUser(&events, user.ID, r.Msg.Filter, page, size); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ListUserEventResponse{
		Page:  int32(events.Page),
		Size:  int32(events.Size),
		Total: int32(events.Total),
		Items: convertEventItems(events.Items),
	}), nil
}

func convertEventItems(src []Event) []*v1.Event {
	items := make([]*v1.Event, 0, len(src))
	for _, item := range src {
		items = append(items, convertEventToPB(item))
	}
	return items
}

func convertEventToPB(src Event) *v1.Event {
	return &v1.Event{
		Id:        src.ID,
		UserId:    src.UserID,
		StartDate: timestamppb.New(src.StartDate),
		EndDate:   timestamppb.New(src.EndDate),
		Name:      src.Name,
		Location:  src.Location,
		Url:       src.URL,
		Detail:    src.Detail,
		Lat:       src.Lat,
		Lng:       src.Lng,
		Zoom:      int32(src.Zoom),
		CreatedAt: timestamppb.New(src.CreatedAt),
		UpdatedAt: timestamppb.New(src.UpdatedAt),
	}
}

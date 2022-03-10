package repository

import (
	"errors"
	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/pkg/eventservice"
	"gorm.io/gorm"
	"time"
)

type eventRepository struct {
	*dbx.DB
}

func NewEventRepository(db *dbx.DB) eventservice.EventRepository {
	return &eventRepository{
		DB: db,
	}
}

func (r eventRepository) FindByEndDateAfter(events *[]eventservice.Event, endDate time.Time) error {
	return r.DB.Select("id, name, start_date, end_date, lat, lng").Where("end_date>=?", endDate).Order("start_date desc").Find(events).Error
}

func (r eventRepository) CountAll(total *int, q string) error {
	var total64 int64

	if q == "" {
		if err := r.DB.Model(&eventservice.Event{}).Count(&total64).Error; err != nil {
			return err
		}
	} else {
		term := "%" + q + "%"
		if err := r.DB.Model(&eventservice.Event{}).Where("name like ? or location like ?", term, term).Count(&total64).Error; err != nil {
			return err
		}
	}

	*total = int(total64)

	return nil
}

func (r eventRepository) FindAll(events *[]eventservice.Event, q string, page int, size int) error {
	offset := size * (page - 1)

	if q == "" {
		return r.DB.Order("start_date desc").Limit(size).Offset(offset).Find(events).Error
	} else {
		term := "%" + q + "%"
		return r.DB.Where("name like ? or location like ?", term, term).Order("start_date desc").Limit(size).Offset(offset).Find(events).Error
	}
}

func (r eventRepository) CountByUserId(total *int, userId string, q string) error {
	var total64 int64

	if q == "" {
		if err := r.DB.Model(&eventservice.Event{}).Where("user_id=?", userId).Count(&total64).Error; err != nil {
			return err
		}
	} else {
		term := "%" + q + "%"
		if err := r.DB.Model(&eventservice.Event{}).Where("user_id=? and (name like ? or location like ?)", userId, term, term).Count(&total64).Error; err != nil {
			return err
		}
	}

	*total = int(total64)

	return nil
}

func (r eventRepository) FindByUserId(events *[]eventservice.Event, userId string, q string, page int, size int) error {
	offset := size * (page - 1)

	if q == "" {
		return r.DB.Where("user_id=?", userId).Order("start_date desc").Limit(size).Offset(offset).Find(events).Error
	} else {
		term := "%" + q + "%"
		return r.DB.Where("user_id=? and (name like ? or location like ?)", userId, term, term).Order("start_date desc").Limit(size).Offset(offset).Find(events).Error
	}
}

func (r eventRepository) FindById(event *eventservice.Event, id string) error {
	if err := r.DB.Where("id=?", id).First(event).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return eventservice.ErrNoSuchEvent
	} else {
		return err
	}
}

func (r eventRepository) FindByIdAndUserId(event *eventservice.Event, id string, userId string) error {
	if err := r.DB.Where("id=? and user_id=?", id, userId).First(event).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return eventservice.ErrNoSuchEvent
	} else {
		return err
	}
}

func (r eventRepository) DeleteByIdAndUserId(id string, userId string) error {
	return r.DB.Where("id=? and user_id=?", id, userId).Delete(&eventservice.Event{}).Error
}

func (r eventRepository) Create(event *eventservice.Event) error {
	return r.DB.Create(event).Error
}

func (r eventRepository) Save(event *eventservice.Event) error {
	return r.DB.Save(event).Error
}

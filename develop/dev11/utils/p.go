package utils

import (
	"errors"
	"time"

	"github.com/Jhnvlglmlbrt/develop/dev11/internal/models"
)

// Валидация параметров методов /create_event, /delete_event, /update_event
func ValidateEventParams(event models.Event) error {
	// Проверка формата даты
	_, errParseDate := time.Parse("2006-01-02", event.Date)
	if errParseDate != nil {
		return errors.New("invalid date format")
	}

	// Проверка формата времени
	_, errParseTime := time.Parse("15:04", event.Time)
	if errParseTime != nil {
		return errors.New("invalid time format")
	}

	return nil
}

func ValidateUpdateEventParams(updateEvent models.UpdEvent) error {
	// Проверка формата даты
	_, errParseDate := time.Parse("2006-01-02", updateEvent.Date)
	if errParseDate != nil {
		return errors.New("invalid date format")
	}

	// Проверка формата времени
	_, errParseTime := time.Parse("15:04", updateEvent.Time)
	if errParseTime != nil {
		return errors.New("invalid time format")
	}

	// Проверка формата новой даты
	_, errParseNewDate := time.Parse("2006-01-02", updateEvent.NewData.Date)
	if errParseNewDate != nil {
		return errors.New("invalid new date format")
	}

	// Проверка формата нового времени
	_, errParseNewTime := time.Parse("15:04", updateEvent.NewData.Time)
	if errParseNewTime != nil {
		return errors.New("invalid new time format")
	}

	return nil
}

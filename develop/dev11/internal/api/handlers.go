package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Jhnvlglmlbrt/develop/dev11/internal/cache"
	"github.com/Jhnvlglmlbrt/develop/dev11/internal/models"
	"github.com/Jhnvlglmlbrt/develop/dev11/logger"
	"github.com/Jhnvlglmlbrt/develop/dev11/utils"
)

type CacheHandler struct {
	c *cache.Cache
}

func NewHandlers(c *cache.Cache) *CacheHandler {
	return &CacheHandler{
		c: c,
	}
}

func (ch *CacheHandler) createOrUpdateEvent(w http.ResponseWriter, r *http.Request, create bool) {
	methodNotAllowed := http.StatusMethodNotAllowed
	badRequest := http.StatusBadRequest
	internalServerError := http.StatusInternalServerError

	if r.Method != http.MethodPost {
		logger.Logger(w, methodNotAllowed, "Method not allowed", nil)
		return
	}

	var decoded models.Event
	var isUpdate bool

	if create {
		if decodingBodyErr := json.NewDecoder(r.Body).Decode(&decoded); decodingBodyErr != nil {
			logger.Logger(w, badRequest, "Error decoding request body", decodingBodyErr)
			return
		}

		if decoded.UserId <= 0 {
			logger.Logger(w, badRequest, "Invalid UserId", errors.New("UserId is required"))
			return
		}

		if err := utils.ValidateEventParams(decoded); err != nil {
			logger.Logger(w, badRequest, "Validation error", err)
			return
		}
	} else {
		var updEvent models.UpdEvent
		if decodingBodyErr := json.NewDecoder(r.Body).Decode(&updEvent); decodingBodyErr != nil {
			logger.Logger(w, badRequest, "Error decoding request body", decodingBodyErr)
			return
		}

		if err := utils.ValidateUpdateEventParams(updEvent); err != nil {
			logger.Logger(w, badRequest, "Validation error", err)
			return
		}

		decoded = models.Event{Date: updEvent.Date, Time: updEvent.Time, UserId: updEvent.UserId}
		isUpdate = true
	}

	if create {
		event := models.NewEvent(decoded.Date, decoded.Time, decoded.UserId)
		ch.c.Create(event)
	} else {
		statusCode := ch.c.Update(decoded, decoded.Date, decoded.Time)

		switch statusCode {
		case http.StatusOK:
			logger.Logger(w, http.StatusOK, "Event updated", nil)
		case http.StatusNotFound:
			logger.Logger(w, http.StatusNotFound, "Event not found", nil)
		case http.StatusInternalServerError:
			logger.Logger(w, internalServerError, "Internal server error", nil)
		default:
			logger.Logger(w, internalServerError, "Unexpected status code", nil)
		}
	}

	if create && !isUpdate {
		logger.Logger(w, http.StatusCreated, "Event created", nil)
	}
}

func (ch *CacheHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	ch.createOrUpdateEvent(w, r, true)
}

func (ch *CacheHandler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	ch.createOrUpdateEvent(w, r, false)
}

func (ch *CacheHandler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	methodNotAllowed := http.StatusMethodNotAllowed
	badRequest := http.StatusBadRequest
	internalServerError := http.StatusInternalServerError

	if r.Method != http.MethodPost {
		logger.Logger(w, methodNotAllowed, "Method not allowed", nil)
		return
	}

	var decoded models.Event

	if decodingBodyErr := json.NewDecoder(r.Body).Decode(&decoded); decodingBodyErr != nil {
		logger.Logger(w, badRequest, "Error decoding request body", decodingBodyErr)
		return
	}

	if err := utils.ValidateEventParams(decoded); err != nil {
		logger.Logger(w, badRequest, "Validation error", err)
		return
	}

	statusCode := ch.c.Delete(decoded.Date, decoded.Time)

	switch statusCode {
	case http.StatusOK:
		logger.Logger(w, http.StatusOK, "Event deleted", nil)
	case http.StatusNotFound:
		logger.Logger(w, http.StatusNotFound, "Event not found", nil)
	default:
		logger.Logger(w, internalServerError, "Unexpected status code", nil)
	}
}

func (ch *CacheHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	ch.deleteEvent(w, r)
}

func (ch *CacheHandler) getEventsHandler(w http.ResponseWriter, r *http.Request, readFunc func(string) ([]*models.Event, bool), errorMessage, successMessage string) {
	if r.Method != http.MethodGet {
		logger.Logger(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	dateQuery := r.URL.Query().Get("date")

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		logger.Logger(w, http.StatusBadRequest, "Invalid date format", errParse)
		return
	}

	val, ok := readFunc(dateQuery)
	if ok {
		logger.Logger(w, http.StatusOK, successMessage, val)
	} else {
		logger.Logger(w, http.StatusNotFound, errorMessage, nil)
	}
}

func (ch *CacheHandler) GetDayEventHandler(w http.ResponseWriter, r *http.Request) {
	ch.getEventsHandler(w, r, ch.c.ReadDay, "No events found for the specified date", "Events retrieved")
}

func (ch *CacheHandler) GetWeekEventHandler(w http.ResponseWriter, r *http.Request) {
	ch.getEventsHandler(w, r, ch.c.ReadWeek, "No events found for the specified week", "Events retrieved")
}

func (ch *CacheHandler) GetMonthEventHandler(w http.ResponseWriter, r *http.Request) {
	ch.getEventsHandler(w, r, ch.c.ReadMonth, "No events found for the specified month", "Events retrieved")
}

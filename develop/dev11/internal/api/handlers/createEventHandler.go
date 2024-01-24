package api

import (
	"encoding/json"
	"net/http"

	"github.com/Jhnvlglmlbrt/develop/dev11/internal/cache"
	"github.com/Jhnvlglmlbrt/develop/dev11/internal/models"
	"github.com/Jhnvlglmlbrt/develop/dev11/logger"
	"github.com/Jhnvlglmlbrt/develop/dev11/utils"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodPost {
		// Используем универсальный логгер для ошибок с кодом 405 (Method Not Allowed)
		logger.Logger(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event

	if decodingBodyErr := decoder.Decode(&decoded); decodingBodyErr != nil {
		// Используем универсальный логгер для ошибок с кодом 400 (Bad Request)
		logger.Logger(w, http.StatusBadRequest, "Error decoding request body", decodingBodyErr)
		return
	}

	// Вспомогательная функция для валидации параметров метода /create_event
	if err := utils.ValidateEventParams(decoded); err != nil {
		// Используем универсальный логгер для ошибок с кодом 400 (Bad Request)
		logger.Logger(w, http.StatusBadRequest, "Validation error", err)
		return
	}

	event := models.NewEvent(decoded.Date, decoded.Time, decoded.UserId)
	c.Create(event)

	// Используем универсальный логгер для успешных ответов с кодом 201 (Created)
	logger.Logger(w, http.StatusCreated, "Event created", nil)
}

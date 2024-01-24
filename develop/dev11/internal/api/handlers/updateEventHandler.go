package api

import (
	"encoding/json"
	"net/http"

	"github.com/Jhnvlglmlbrt/develop/dev11/internal/cache"
	"github.com/Jhnvlglmlbrt/develop/dev11/internal/models"
	"github.com/Jhnvlglmlbrt/develop/dev11/logger"
	"github.com/Jhnvlglmlbrt/develop/dev11/utils"
)

func UpdateEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodPost {
		// Используем универсальный логгер для ошибок с кодом 405 (Method Not Allowed)
		logger.Logger(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.UpdEvent

	if decodingBodyErr := decoder.Decode(&decoded); decodingBodyErr != nil {
		// Используем универсальный логгер для ошибок с кодом 400 (Bad Request)
		logger.Logger(w, http.StatusBadRequest, "Error decoding request body", decodingBodyErr)
		return
	}

	// Валидация параметров
	if err := utils.ValidateUpdateEventParams(decoded); err != nil {
		// Используем универсальный логгер для ошибок с кодом 400 (Bad Request)
		logger.Logger(w, http.StatusBadRequest, "Validation error", err)
		return
	}

	// Обновление события
	statusCode := c.Update(models.Event{Date: decoded.Date, Time: decoded.Time, UserId: decoded.UserId}, decoded.NewData.Date, decoded.NewData.Time)

	// Проверяем статус код после обновления и возвращаем соответствующий HTTP статус
	switch statusCode {
	case http.StatusOK:
		// Используем универсальный логгер для успешных ответов с кодом 200 (OK)
		logger.Logger(w, http.StatusOK, "Event updated", nil)
	case http.StatusNotFound:
		// Используем универсальный логгер для ошибок с кодом 404 (Not Found)
		logger.Logger(w, http.StatusNotFound, "Event not found", nil)
	case http.StatusInternalServerError:
		// Используем универсальный логгер для ошибок с кодом 500 (Internal Server Error)
		logger.Logger(w, http.StatusInternalServerError, "Internal server error", nil)
	default:
		// Если возвратили неизвестный статус код, логгируем ошибку
		logger.Logger(w, http.StatusInternalServerError, "Unexpected status code", nil)
	}
}

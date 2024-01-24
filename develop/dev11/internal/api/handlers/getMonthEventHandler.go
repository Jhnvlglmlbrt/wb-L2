package api

import (
	"net/http"
	"time"

	"github.com/Jhnvlglmlbrt/develop/dev11/internal/cache"
	"github.com/Jhnvlglmlbrt/develop/dev11/logger"
)

func GetMonthEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	if r.Method != http.MethodGet {
		logger.Logger(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	dateQuery := r.URL.Query().Get("date")

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		logger.Logger(w, http.StatusBadRequest, "Invalid date format", errParse)
		return
	}

	val, ok := c.ReadMonth(dateQuery)
	if ok {
		logger.Logger(w, http.StatusOK, "Events retrieved", val)
	} else {
		logger.Logger(w, http.StatusNotFound, "No events found for the specified month", nil)
	}
}

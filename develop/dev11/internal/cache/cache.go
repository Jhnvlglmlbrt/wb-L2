package cache

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Jhnvlglmlbrt/develop/dev11/internal/models"
)

type Cache struct {
	m     sync.RWMutex
	cache map[string][]*models.Event
}

func NewCache() *Cache {
	return &Cache{
		m:     sync.RWMutex{},
		cache: make(map[string][]*models.Event),
	}
}

func (c *Cache) Create(e *models.Event) {
	c.m.Lock()
	defer c.m.Unlock()
	c.cache[e.Date] = append(c.cache[e.Date], e)
}

func (c *Cache) ReadDay(date string) ([]*models.Event, bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	val, ok := c.cache[date]
	// fmt.Println(c.cache)
	if ok {
		return val, true
	}
	log.Println("Err at getting", date, ": no events detected")
	return nil, false
}

func (c *Cache) ReadWeek(date string) ([]*models.Event, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	today, err := parseDate(date)
	if err != nil {
		log.Println("Error at parsing date: ", err)
		return nil, false
	}

	var result []*models.Event
	for _, ev := range c.cache {
		for _, event := range ev {
			evDate, err := parseDate(event.Date)
			if err != nil {
				log.Println("Error at parsing date: ", err)
				return nil, false
			}

			if evDate.After(today) && evDate.Before(today.Add(time.Hour*24*7)) {
				result = append(result, event)
			}
		}
	}

	return result, len(result) > 0
}

func (c *Cache) ReadMonth(date string) ([]*models.Event, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	today, err := parseDate(date)
	if err != nil {
		log.Println("Error at parsing date: ", err)
		return nil, false
	}

	var result []*models.Event
	for _, ev := range c.cache {
		for _, event := range ev {
			evDate, err := parseDate(event.Date)
			if err != nil {
				log.Println("Error at parsing date: ", err)
				return nil, false
			}

			if evDate.After(today) && evDate.Before(today.AddDate(0, 1, 0)) {
				result = append(result, event)
			}
		}
	}

	return result, len(result) > 0
}

func (c *Cache) Delete(date, eventTime string) int {
	c.m.Lock()
	defer c.m.Unlock()

	events := c.cache[date]
	for i, event := range events {
		if event.Time == eventTime {
			c.cache[date] = append(events[:i], events[i+1:]...)
			if len(c.cache[date]) == 0 {
				delete(c.cache, date)
			}
			return http.StatusOK
		}
	}

	return http.StatusNotFound
}

func (c *Cache) Update(e models.Event, newDate, newTime string) int {
	c.m.Lock()
	defer c.m.Unlock()

	events := c.cache[e.Date]
	for i, event := range events {
		if event.Time == e.Time && event.Date == e.Date {
			if newTime != "" {
				event.Time = newTime
			}
			if newDate != "" {
				event.Date = newDate
				c.cache[newDate] = append(c.cache[newDate], event)
				c.cache[e.Date] = append(events[:i], events[i+1:]...)
			}
			return http.StatusOK
		}
	}

	return http.StatusNotFound
}

func parseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

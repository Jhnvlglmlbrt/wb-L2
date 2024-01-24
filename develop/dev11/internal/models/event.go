package models

import (
	time2 "time"
)

func NewEvent(date, time string, userId int) *Event {
	Uid := time2.Now().Unix()
	return &Event{
		UserId: userId,
		Date:   date,
		Time:   time,
		Uid:    Uid,
	}
}

type Event struct {
	UserId int    `json:"user_id"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Uid    int64  `json:"uid"`
}

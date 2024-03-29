package models

type UpdEvent struct {
	UserId  int    `json:"user_id"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Uid     int64  `json:"uid"`
	NewData struct {
		UserId int    `json:"user_id"`
		Date   string `json:"date"`
		Time   string `json:"time"`
	} `json:"new_data"`
}

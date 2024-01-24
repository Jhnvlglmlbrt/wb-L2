package response

type Error struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

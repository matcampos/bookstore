package errormodel

type Error struct {
	Code     int       `json:"code"`
	Messages []Message `json:"messages"`
	Field    string    `json:"field,omitempty"`
}

type Message struct {
	Pt string `json:"pt"`
	En string `json:"en"`
}

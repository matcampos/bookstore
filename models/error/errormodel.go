package errormodel

// Error struct receives three parameters: Code which is an integer to identify the error, Messages is a Message struct array and Field is a string to specify the field which was sent wrong.
type Error struct {
	Code     int       `json:"code"`
	Messages []Message `json:"messages"`
	Field    string    `json:"field,omitempty"`
}

// Message struct receives two parameters, Pt is the message in Portuguese and En is the message in English.
type Message struct {
	Pt string `json:"pt"`
	En string `json:"en"`
}

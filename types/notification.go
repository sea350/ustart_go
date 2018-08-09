package types

import (
	"time"
)

type Notification struct {
	Content   string    `json:"Content"`
	Hyperlink string    `json:"Hyperlink"`
	Timestamp time.Time `json:"Timestamp"`
	UserID    string    `json:"UserID"`
}

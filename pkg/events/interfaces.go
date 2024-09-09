package events

import "time"

type EventInterface interface {
	GetName() string
	GetDate() time.Time
	GetPayload() interface{}
}

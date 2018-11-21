package model

import (
	"time"
)

// Counter is basic unit about count
type Counter struct {
	UUID  string    `json:"uuid"`
	Count int       `json:"count"`
	Life  time.Time `json:"life"`
}

// Empty is create empty Counter
func Empty() Counter {
	return Counter{Life: time.Unix(0, 0), UUID: "empty", Count: 0}
}

// CounterCreate is request object for create count
type CounterCreate struct {
	Duration time.Duration `json:"duration"`
}

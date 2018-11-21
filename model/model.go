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

// CounterCreate is request object for create count
type CounterCreate struct {
	Duration time.Duration `json:"duration"`
}

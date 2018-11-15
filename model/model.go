package model

import (
"time"
)

type Counter struct {
	UUID string `json:uuid`
	Count int `json:count`
	Life time.Time `json:life`
}

type CounterCreate struct {
	Duration time.Duration
}
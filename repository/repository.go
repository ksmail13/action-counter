package repository

import (
	"time"

	"github.com/ksmail13/action-counter/model"
)

type Repository interface {
	Set(duration time.Duration) (model.Counter, error)
	Get(key string) (model.Counter, error)
	Increase(key string) (model.Counter, error)
	Decrease(key string) (model.Counter, error)
	Clear() error
}

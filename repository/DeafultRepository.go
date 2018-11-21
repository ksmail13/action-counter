package repository

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ksmail13/action-counter/errors"
	"github.com/ksmail13/action-counter/model"
)

func Default() Repository {
	return &defaultRepository{storage: make(map[string]model.Counter)}
}

// DefaultRepository is repository using map and simple timer
// it does not support threadsafe operation
type defaultRepository struct {
	storage map[string]model.Counter
}

func (r defaultRepository) deleteAfterDuration(uuid string, duration time.Duration) {
	deleteTimer := time.NewTimer(time.Second * duration)
	go func() {
		<-deleteTimer.C
		delete(r.storage, uuid)
	}()
}

func (r *defaultRepository) Set(duration time.Duration) (model.Counter, error) {
	uuid := uuid.New().String()
	counter := model.Counter{UUID: uuid, Count: 1, Life: time.Now().Add(duration)}
	log.Printf("create %s", uuid)

	if r.storage == nil {
		r.Clear()
	}

	r.storage[uuid] = counter

	r.deleteAfterDuration(uuid, duration)
	return counter, nil
}

func (r *defaultRepository) Get(key string) (model.Counter, error) {
	val, exist := r.storage[key]
	if !exist {
		err := errors.NotFound(key)
		return model.Empty(), err
	}

	return val, nil
}

func (r *defaultRepository) Increase(key string) (model.Counter, error) {
	val, err := r.Get(key)

	if err == nil {
		val.Count++
	}

	return val, err
}

func (r *defaultRepository) Decrease(key string) (model.Counter, error) {
	val, err := r.Get(key)

	if err == nil {
		val.Count--
	}

	return val, err
}

func (r *defaultRepository) Clear() error {
	r.storage = map[string]model.Counter{}

	return nil
}

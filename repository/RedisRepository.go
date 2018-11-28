package repository

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/ksmail13/action-counter/config"
	"github.com/ksmail13/action-counter/model"
)

// Redis function is Create RedisRepository
func Redis(conf *config.Config) Repository {
	repo := redisRepository{}
	repo.client = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
		DB:       conf.RedisDB})

	return repo
}

type redisRepository struct {
	client *redis.Client
}

func (r redisRepository) Set(duration time.Duration) (model.Counter, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		log.Printf("error while create UUID (%s)", err)
		return model.Empty(), err
	}
	key := id.String()
	count := model.Counter{UUID: key, Count: 1, Life: time.Now().Add(duration)}

	return r.setData(key, count, duration)
}

func (r redisRepository) setData(key string, count model.Counter, duration time.Duration) (model.Counter, error) {
	err := r.client.Set(key, count, duration*time.Second).Err()

	if err != nil {
		log.Printf("error while set redis (%s)", err)
		return model.Empty(), err
	}

	return count, nil
}

func (r redisRepository) Get(key string) (model.Counter, error) {
	saved, err := r.client.Get(key).Result()

	if err != nil {
		log.Printf("error while get redis (%s)", err)
		return model.Empty(), err
	}
	log.Printf("result %s", saved)
	count := &model.Counter{}
	err = json.Unmarshal([]byte(saved), count)
	if err != nil {
		log.Printf("error while unmarshall (%s)", err)
		return model.Empty(), err
	}

	return *count, nil
}

func (r redisRepository) Increase(key string) (model.Counter, error) {
	count, err := r.Get(key)
	if err != nil {
		return count, err
	}

	count.Count++
	return r.setData(key, count, time.Duration(time.Since(count.Life).Seconds()))
}

func (r redisRepository) Decrease(key string) (model.Counter, error) {
	count, err := r.Get(key)
	if err != nil {
		return count, err
	}

	count.Count--
	return r.setData(key, count, time.Duration(time.Since(count.Life).Seconds()))
}

func (r redisRepository) Clear() error {
	err := r.client.FlushAll().Err()
	if err != nil {
		log.Printf("error during clear (%s)", err)
	}
	return nil
}

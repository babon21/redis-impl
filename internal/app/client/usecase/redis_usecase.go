package usecase

import (
	"io"
	"net/http"
)

type RedisUsecase interface {
	Set(body io.Reader) (*http.Response, error)
	Get(key string) (*http.Response, error)
	Del(key string) (*http.Response, error)
	Keys(body io.Reader) (*http.Response, error)

	HGet(body io.Reader) (*http.Response, error)
	HSet(body io.Reader) (*http.Response, error)

	LGet(body io.Reader) (*http.Response, error)
	LSet(body io.Reader) (*http.Response, error)
	LPush(body io.Reader) (*http.Response, error)

	Expire(body io.Reader) (*http.Response, error)
}

type redisUsecase struct {
	redisGateway RedisGateway
}

func NewRedisUsecase(redisGateway RedisGateway) RedisUsecase {
	return &redisUsecase{redisGateway: redisGateway}
}

func (r *redisUsecase) Set(body io.Reader) (*http.Response, error) {
	return r.redisGateway.Set(body)
}

func (r *redisUsecase) Get(key string) (*http.Response, error) {
	return r.redisGateway.Get(key)
}

func (r *redisUsecase) Del(key string) (*http.Response, error) {
	return r.redisGateway.Del(key)
}

func (r *redisUsecase) Keys(body io.Reader) (*http.Response, error) {
	return r.redisGateway.Keys(body)
}

func (r *redisUsecase) HGet(body io.Reader) (*http.Response, error) {
	return r.redisGateway.HGet(body)
}

func (r *redisUsecase) HSet(body io.Reader) (*http.Response, error) {
	return r.redisGateway.HSet(body)
}

func (r *redisUsecase) LGet(body io.Reader) (*http.Response, error) {
	return r.redisGateway.LGet(body)
}

func (r *redisUsecase) LSet(body io.Reader) (*http.Response, error) {
	return r.redisGateway.LSet(body)
}

func (r *redisUsecase) LPush(body io.Reader) (*http.Response, error) {
	return r.redisGateway.LPush(body)
}

func (r *redisUsecase) Expire(body io.Reader) (*http.Response, error) {
	return r.redisGateway.Expire(body)
}

package gateway

import (
	"github.com/babon21/redis-impl/internal/app/client/usecase"
	"io"
	"net/http"
)

type RedisGatewayImpl struct {
	redisServerUrl string
}

func NewRedisGateway(redisServerUrl string) usecase.RedisGateway {
	return &RedisGatewayImpl{redisServerUrl: redisServerUrl}
}

func (r *RedisGatewayImpl) Set(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/string"
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) Get(key string) (*http.Response, error) {
	return http.Get(r.redisServerUrl + "/cache/string/" + key)
}

func (r *RedisGatewayImpl) Del(key string) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/keys/" + key
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) Keys(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/keys"
	request, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) HGet(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/map"
	request, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) HSet(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/map"
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) LGet(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/list"
	request, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) LSet(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/list"
	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) LPush(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/list"
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

func (r *RedisGatewayImpl) Expire(body io.Reader) (*http.Response, error) {
	url := r.redisServerUrl + "/cache/keys/expire"
	request, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(request)
}

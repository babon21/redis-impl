package repository

import (
	"github.com/babon21/redis-impl/internal/app/server/domain"
	"github.com/babon21/redis-impl/internal/app/server/usecase"
	"regexp"
	"sync"
	"time"
)

type InMemoryRedis struct {
	store    sync.Map
	interval int
}

func NewInMemoryRedisStore() usecase.RedisStore {
	store := InMemoryRedis{interval: 20}
	go store.startPeriodicDeleteExpirationKey()
	return &store
}

type storeValue struct {
	value  interface{}
	expiry time.Time
}

func (r *InMemoryRedis) Set(key string, value string) {
	r.store.Store(key, storeValue{
		value: value,
	})
}

func (r *InMemoryRedis) Get(key string) (string, bool, error) {
	value, exists := r.load(key)
	if !exists {
		return "", false, nil
	}

	strValue, ok := value.value.(string)
	if !ok {
		return "", false, domain.ErrWrongType
	}

	return strValue, true, nil
}

func (r *InMemoryRedis) load(key string) (storeValue, bool) {
	val, ok := r.store.Load(key)
	if !ok {
		return storeValue{}, false
	}

	value := val.(storeValue)
	if r.tryDeleteKeyIfExpire(key, value) {
		return storeValue{}, false
	}

	return value, true
}

func (r *InMemoryRedis) Del(key string) bool {
	if _, found := r.store.LoadAndDelete(key); found {
		return true
	}
	return false
}

func (r *InMemoryRedis) Keys(pattern string) ([]string, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, 10)
	r.store.Range(func(key, value interface{}) bool {
		keyString := key.(string)
		if reg.MatchString(keyString) {
			result = append(result, keyString)
		}
		return true
	})
	return result, nil
}

func (r *InMemoryRedis) startPeriodicDeleteExpirationKey() {
	ticker := time.NewTicker(time.Second * time.Duration(r.interval))
	for range ticker.C {
		r.store.Range(func(key, val interface{}) bool {
			keyString := key.(string)
			value := val.(storeValue)
			r.tryDeleteKeyIfExpire(keyString, value)

			return true
		})
	}
}

func (r *InMemoryRedis) HGet(key string, field string) (string, bool, error) {
	value, exists := r.load(key)
	if !exists {
		return "", false, nil
	}

	storedMap, ok := value.value.(map[string]string)
	if !ok {
		return "", false, domain.ErrWrongType
	}

	v, ok := storedMap[field]
	if !ok {
		return "", false, nil
	}

	return v, true, nil
}

func (r *InMemoryRedis) HSet(key string, field string, value string) error {
	val, exists := r.load(key)
	if !exists {
		newMap := make(map[string]string)
		r.store.Store(key, storeValue{
			value: newMap,
		})

		newMap[field] = value
		return nil
	}

	storedMap, ok := val.value.(map[string]string)
	if !ok {
		return domain.ErrWrongType
	}

	storedMap[field] = value
	return nil
}

func (r *InMemoryRedis) LGet(key string, index int) (string, error) {
	val, exists := r.load(key)
	if !exists {
		return "", domain.ErrNoSuchKey
	}

	list, ok := val.value.([]string)
	if !ok {
		return "", domain.ErrWrongType
	}

	arrLength := len(list)
	if index < 0 || index >= arrLength {
		return "", domain.ErrIndexOutOfRange
	}

	return list[index], nil
}

func (r *InMemoryRedis) LSet(key string, index int, value string) error {
	val, exists := r.load(key)
	if !exists {
		return domain.ErrNoSuchKey
	}

	list, ok := val.value.([]string)
	if !ok {
		return domain.ErrWrongType
	}

	arrLength := len(list)
	if index < 0 || index >= arrLength {
		return domain.ErrIndexOutOfRange
	}

	list[index] = value
	return nil
}

func (r *InMemoryRedis) LPush(key string, value string) (int, error) {
	val, exists := r.load(key)
	if !exists {
		newArr := make([]string, 0, 5)
		newArr = append(newArr, value)
		r.store.Store(key, storeValue{
			value: newArr,
		})

		return len(newArr), nil
	}

	arr, ok := val.value.([]string)
	if !ok {
		return -1, domain.ErrWrongType
	}

	arr = append(arr, value)
	r.store.Store(key, storeValue{
		value: arr,
	})
	return len(arr), nil
}

func (r *InMemoryRedis) Expire(key string, duration int) bool {
	val, exists := r.load(key)
	if !exists {
		return false
	}

	val.expiry = time.Now().Add(time.Second * time.Duration(duration))
	r.store.Store(key, val)
	return true
}

func (r *InMemoryRedis) checkKeyExpiration(val storeValue) bool {
	var defaultTime time.Time
	if val.expiry == defaultTime {
		return false
	}

	if val.expiry.After(time.Now()) {
		return false
	}
	return true
}

func (r *InMemoryRedis) tryDeleteKeyIfExpire(key string, val storeValue) bool {
	if r.checkKeyExpiration(val) {
		r.Del(key)
		return true
	}
	return false
}

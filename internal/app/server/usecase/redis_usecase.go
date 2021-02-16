package usecase

type RedisUsecase interface {
	Set(key string, value string)
	Get(key string) (string, bool, error)
	Del(key string) bool
	Keys(pattern string) ([]string, error)

	HGet(key string, field string) (string, bool, error)
	HSet(key string, pairs []FieldValue) (int, error)

	LGet(key string, index int) (string, error)
	LSet(key string, index int, value string) error
	LPush(key string, values []string) (int, error)

	Expire(key string, duration int) bool
}

type redisUsecase struct {
	redisStore RedisStore
}

func NewRedisUsecase(redisStore RedisStore) RedisUsecase {
	return &redisUsecase{
		redisStore: redisStore,
	}
}

func (r *redisUsecase) Set(key string, value string) {
	r.redisStore.Set(key, value)
}

func (r *redisUsecase) Get(key string) (string, bool, error) {
	return r.redisStore.Get(key)
}

func (r *redisUsecase) Del(key string) bool {
	return r.redisStore.Del(key)
}

func (r *redisUsecase) Keys(pattern string) ([]string, error) {
	return r.redisStore.Keys(pattern)
}

func (r *redisUsecase) HGet(key string, field string) (string, bool, error) {
	return r.redisStore.HGet(key, field)
}

func (r *redisUsecase) HSet(key string, pairs []FieldValue) (int, error) {
	for _, pair := range pairs {
		if err := r.redisStore.HSet(key, pair.Field, pair.Value); err != nil {
			return -1, err
		}
	}
	return len(pairs), nil
}

func (r *redisUsecase) LGet(key string, index int) (string, error) {
	return r.redisStore.LGet(key, index)
}

func (r *redisUsecase) LSet(key string, index int, value string) error {
	return r.redisStore.LSet(key, index, value)
}

func (r *redisUsecase) LPush(key string, values []string) (int, error) {
	var err error
	size := 0

	for i := len(values) - 1; i >= 0; i-- {
		size, err = r.redisStore.LPush(key, values[i])
		if err != nil {
			return -1, err
		}
	}
	return size, nil
}

func (r *redisUsecase) Expire(key string, duration int) bool {
	return r.redisStore.Expire(key, duration)
}

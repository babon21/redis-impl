package usecase

type RedisRepository interface {
	Set(key string, value string)
	Get(key string) (string, bool, error)
	Del(key string) bool
	Keys(pattern string) ([]string, error)

	HGet(key string, field string) (string, bool, error)
	HSet(key string, field string, value string) (bool, error)

	LGet(key string, index int) (string, error)
	LSet(key string, index int, value string) error
	LPush(key string, value string) (int, error)

	Expire(key string, duration int) bool
}

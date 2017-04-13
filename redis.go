package monitoring

import (
	"gopkg.in/redis.v6"
	"fmt"
)

type RedisCollector struct {

}

type RedisConfig struct {
	Host string
	Port int
	Password string

}

func CreateRedisCollector(c RedisConfig) *redis.Client  {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
	})


}

func(r *RedisCollector) Collect() map[string]int {
	return make(map[string]int)
}
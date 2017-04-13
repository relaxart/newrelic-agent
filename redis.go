package monitoring

import (
	"fmt"
	"gopkg.in/redis.v5"
	"time"
)

type RedisCollector struct {
	Name     string
	redisCli *redis.Client
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

func CreateRedisCollector(c RedisConfig, name string) RedisCollector {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
	})

	return RedisCollector{
		Name:     name,
		redisCli: rds,
	}

}

func (c *RedisCollector) Process(ch chan Stat) {
	for {
		s := CreateStat(c.Name)
		s.Metrics.Add("connections", 403)
		s.Metrics.Add("keys", 2512451)
		ch <- s
		time.Sleep(10 * time.Second)
	}
}

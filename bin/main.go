package main

import (
	"monitoring"
)

func main() {
	client := new(monitoring.RedisCollector)
	client.Name = "Redis"
	config := monitoring.CollectorConfig{
		Host:  "http://localhost",
		Key:   "2134213421sdjfhsjdahfkjsdf",
		Delay: 10,
	}
	collectors := monitoring.CreateCollector(config)
	collectors.Add(client)
	collectors.Start()
	collectors.Wait()
}

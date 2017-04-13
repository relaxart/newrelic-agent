package main

import (
	"fmt"
	"monitoring"
)

func main() {
	fmt.Println("Start processing.......")
	client := monitoring.RedisCollector{}
	config := monitoring.CollectorConfig{
		Host:  "http://localhost",
		Key:   "2134213421sdjfhsjdahfkjsdf",
		Delay: 60,
	}
	collectors := monitoring.CreateCollector(config)
	collectors.Add(client)
	fmt.Println(collectors)
}

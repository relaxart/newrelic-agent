package monitoring

import (
	"fmt"
	"time"
)

type CollectorConfig struct {
	Host string
	Key string
	Delay time.Duration
}

func CreateCollector(c CollectorConfig) *Collectors {
	collector := new(Collectors)
	collector.config = c
	return collector
}

type Collectors struct {
	config CollectorConfig
	collectors []Collector
}

func (c *Collectors) Add(collector Collector)  {
	c.collectors = append(c.collectors, collector)
}

func (c *Collectors) All() []Collector  {
	return c.collectors
}

func (c *Collectors) Start()  {
	fmt.Println("Start processing")
	go c.process()
}

func (c *Collectors) process()  {
	for {
		time.Sleep(c.config.Delay * time.Millisecond)
		fmt.Println("Send data to NewRelic")
	}
}

func (c *Collectors) send()  {
	fmt.Println("Send stat")
}
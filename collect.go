package monitoring

import (
	"fmt"
	"sync"
	"time"
)

type CollectorConfig struct {
	Host  string
	Key   string
	Delay int
}

func CreateCollector(c CollectorConfig) *Collectors {
	collector := new(Collectors)
	collector.config = c
	collector.wg = sync.WaitGroup{}

	return collector
}

type Collectors struct {
	config     CollectorConfig
	wg         sync.WaitGroup
	h          Harvest
	collectors []Collector
}

func (c *Collectors) Add(collector Collector) {
	c.collectors = append(c.collectors, collector)
}

func (c *Collectors) All() []Collector {
	return c.collectors
}

func (c *Collectors) Start() {
	fmt.Println("Start data processing")
	c.wg.Add(1)
	ch := make(chan Stat)

	for _, col := range c.All() {
		go col.Process(ch)
	}
	go c.collect(ch)
	go c.harvest()
}

func (c *Collectors) Wait() {
	c.wg.Wait()
}

func (c *Collectors) collect(collectCh chan Stat) {
	for stat := range collectCh {
		c.h.Add(stat)
	}
}

func (c *Collectors) harvest() {
	for {
		time.Sleep(time.Duration(c.config.Delay) * time.Second)
		go c.process(c.h.Harvest())
	}
}

func (c *Collectors) process(stats []Stat) {
	fmt.Println("Send data")
	for _, stat := range stats {
		fmt.Println(stat)
	}
}
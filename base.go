package monitoring

type Collector interface {
	Process(ch chan Stat)
}

type BaseCollector struct {
	Name string
}

type Stat struct {
	Name    string
	Metrics Metrics
}

func CreateStat(name string) Stat {
	return Stat{
		Name:    name,
		Metrics: CreateMetrics(),
	}
}

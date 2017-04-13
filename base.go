package monitoring

type Collector interface {
	Collect() map[string]int
}

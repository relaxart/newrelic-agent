package monitoring

import "sync"

var mutex = &sync.Mutex{}

type Harvest struct {
	stats []Stat
}

func (h *Harvest) Add(s Stat) {
	mutex.Lock()
	h.stats = append(h.stats, s)
	mutex.Unlock()
}

func (h *Harvest) Harvest() []Stat {
	mutex.Lock()
	defer func() {
		h.stats = nil
		mutex.Unlock()
	}()

	return h.stats
}

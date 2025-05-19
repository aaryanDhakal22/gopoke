package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	cache    map[string]CacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *PokeCache {
	newPc := &PokeCache{
		cache:    make(map[string]CacheEntry),
		interval: interval,
		mu:       &sync.RWMutex{},
	}
	go newPc.reapLoop()
	return newPc
}

func (pc *PokeCache) Add(key string, val []byte) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (pc *PokeCache) Get(key string) ([]byte, bool) {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	if ret, ok := pc.cache[key]; ok {
		return ret.val, true
	}
	return nil, false
}

func (pc *PokeCache) delete(keys []string) {
	if len(keys) == 0 {
		return
	}
	pc.mu.Lock()
	defer pc.mu.Unlock()

	for _, key := range keys {
		delete(pc.cache, key)
	}
}
func (pc *PokeCache) reapLoop() {
	ticker := time.NewTicker(pc.interval * time.Second)
	for {
		// This line is blocked till the ticker ticks
		<-ticker.C
		go func() {
			mark := []string{}
			for k, v := range pc.cache {
				delta := v.createdAt.Sub(time.Now())
				if delta.Abs() >= pc.interval {
					mark = append(mark, k)
				} else {
					continue
				}
			}
			pc.delete(mark)
		}()
	}
}

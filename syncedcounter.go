package dizzy

import "sync"

type syncedCounter struct {
	mutex sync.Mutex
	c     int
}

func (sc *syncedCounter) add() {
	sc.mutex.Lock()
	sc.c++
	sc.mutex.Unlock()
}

func (sc *syncedCounter) count() int {
	sc.mutex.Lock()
	c := sc.c
	sc.mutex.Unlock()
	return c
}

func (sc *syncedCounter) set(i int) {
	sc.mutex.Lock()
	sc.c = i
	sc.mutex.Unlock()
}

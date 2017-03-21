package dining_philosophers

import (
	"sync"
)

type mutexPhilosopher struct {
	Id    int
	Count int
	left  *sync.Mutex
	right *sync.Mutex
}

func (p *mutexPhilosopher) Loop(wg *sync.WaitGroup) {
	wg.Wait()
	for {
		if p.Id == NUM_PHILOSOPHERS-1 {
			p.right.Lock()
			p.left.Lock()
		} else {
			p.left.Lock()
			p.right.Lock()
		}
		p.Count += 1

		p.left.Unlock()
		p.right.Unlock()
	}
}

func newMutexes() (fs []sync.Mutex) {
	fs = make([]sync.Mutex, NUM_PHILOSOPHERS)

	return
}

func NewMutexPhilosopher() []mutexPhilosopher {
	fs := newMutexes()
	ps := make([]mutexPhilosopher, NUM_PHILOSOPHERS)
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		ps[i] = mutexPhilosopher{
			Id:    i,
			left:  &fs[i],
			right: &fs[(i+1)%NUM_PHILOSOPHERS],
		}
	}
	return ps
}

package dining_philosophers

import (
	"sync"
)

const NUM_PHILOSOPHERS = 5

type fork struct{}

type Philosopher struct {
	Id    int
	Count int
	left  chan fork
	right chan fork
}

func (p *Philosopher) Loop(wg *sync.WaitGroup) {
	wg.Wait()
	for {
		p.eat()
		p.think()
	}
}

func (p *Philosopher) eat() {
	if p.Id == NUM_PHILOSOPHERS-1 {
		<-p.right
		<-p.left
	} else {
		<-p.left
		<-p.right
	}
}

func (p *Philosopher) think() {
	p.Count += 1
	p.right <- fork{}
	p.left <- fork{}
}

func NewChannelPhilosopher() []Philosopher {
	fs := newForks()
	ps := make([]Philosopher, NUM_PHILOSOPHERS)
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		ps[i] = Philosopher{
			Id:    i,
			left:  fs[i],
			right: fs[(i+1)%NUM_PHILOSOPHERS]}
	}
	return ps
}

func newForks() []chan fork {
	fs := make([]chan fork, NUM_PHILOSOPHERS)
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		fs[i] = make(chan fork, 1)
		fs[i] <- fork{}
	}
	return fs
}

package main

import (
	"fmt"
	"sync"
	"time"
)

const NUM_PHILOSOPHERS = 5

var wg sync.WaitGroup

type fork struct{}

type Philosopher struct {
	id    int
	count int
	left  chan fork
	right chan fork
}

func (p *Philosopher) loop() {
	wg.Wait()
	for {
		p.eat()
		p.think()
	}
}

func (p *Philosopher) eat() {
	if p.id == NUM_PHILOSOPHERS-1 {
		<-p.right
		<-p.left
	} else {
		<-p.left
		<-p.right
	}
}

func (p *Philosopher) think() {
	p.count += 1
	p.right <- fork{}
	p.left <- fork{}
}

func main() {
	ps := newPhilosophers()
	wg.Add(1)

	for i, _ := range ps {
		go ps[i].loop()
	}

	wg.Done()

	time.Sleep(5 * time.Second)

	for _, p := range ps {
		fmt.Println("Thinker", p.id, "count", p.count)
	}
}

func newPhilosophers() []Philosopher {
	fs := newForks()
	ps := make([]Philosopher, NUM_PHILOSOPHERS)
	for i := 0; i < NUM_PHILOSOPHERS; i++ {
		ps[i] = Philosopher{
			id:    i,
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

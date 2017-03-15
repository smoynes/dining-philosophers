package main

import (
	"fmt"
	"sync"
	"time"
)

type fork struct{}

type Philosopher struct {
	id    int
	count int
	left  chan fork
	right chan fork
}

func (p *Philosopher) eat() {
	if (p.id % 2) == 0 {
		<-p.left
		<-p.right
	} else {
		<-p.right
		<-p.left
	}
}

func (p *Philosopher) think() {
	p.right <- fork{}
	p.left <- fork{}
	p.count += 1
}

func (p *Philosopher) loop() {
	wg.Wait()
	for {
		p.eat()
		p.think()
	}
}

var wg sync.WaitGroup

func main() {
	ps := newPhilosophers()

	wg.Add(1)

	for _, p := range ps {
		go p.loop()
	}

	wg.Done()

	time.Sleep(5 * time.Second)

	for _, p := range ps {
		fmt.Println("Thinker", p.id, "count", p.count)
	}
}

func newPhilosophers() (ps [5]*Philosopher) {
	fs := newForks()

	for i := 0; i < 5; i++ {
		ps[i] = &Philosopher{
			id:    i,
			left:  fs[i],
			right: fs[(i+1)%5]}
	}
	return
}

func newForks() (fs [5]chan fork) {
	for i := 0; i < 5; i++ {
		fs[i] = make(chan fork, 1)
		fs[i] <- fork{}
	}
	return fs
}

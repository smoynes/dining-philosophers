package main

import (
	"fmt"
	"sync"
	"time"

	dining_philosophers "github.com/smoynes/dining-philosophers"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ps := dining_philosophers.NewChannelPhilosopher()
	for i, _ := range ps {
		go ps[i].Loop(&wg)
	}

	wg.Done()

	time.Sleep(5 * time.Second)

	for _, p := range ps {
		fmt.Println("Chan Thinker", p.Id, "count", p.Count)
	}

	wg.Add(1)
	mps := dining_philosophers.NewMutexPhilosopher()
	for i, _ := range ps {
		go mps[i].Loop(&wg)
	}

	wg.Done()

	time.Sleep(5 * time.Second)

	for _, p := range mps {
		fmt.Println("Mutex Thinker", p.Id, "count", p.Count)
	}

}

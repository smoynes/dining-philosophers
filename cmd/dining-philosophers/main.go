package main

import (
	"fmt"
	"sync"
	"time"

	dinig_philosophers "github.com/smoynes/dining-philosophers"
)

const NUM_PHILOSOPHERS = 5

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ps := dining_philosophers.NewChannelPhilosopher()

	for i, _ := range ps {
		go ps[i].loop(wg)
	}

	wg.Done()

	time.Sleep(5 * time.Second)

	for _, p := range ps {
		fmt.Println("Thinker", p.id, "count", p.count)
	}
}

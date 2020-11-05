package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	leftCS, rightCS *ChopStick
	id              int
}

func (p Philosopher) eat(wg *sync.WaitGroup) {

	defer wg.Done()

	if rand.Float32() < 0.5 {
		p.leftCS.Lock()
		p.rightCS.Lock()
	} else {
		p.rightCS.Lock()
		p.leftCS.Lock()
	}

	for i := 0; i < 3; i++ {

		fmt.Println("Starting to eat", p.id+1)

		fmt.Println("Finishing eating", p.id+1)
	}

	if rand.Float32() < 0.5 {
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	} else {
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(5)

	chopstics := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopstics[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			chopstics[i],
			chopstics[(i+1)%5], i}
	}

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

// As we saw the race condition
// Now lets apply lock

func (t *TestNumber) incrementValue(m *sync.Mutex, wg sync.WaitGroup) {
	// m.Lock will block other go routines
	m.Lock()
	t.Value += 1
	wg.Done()
	m.Unlock()
}

type TestNumber struct {
	Value int
}

// Mutex should always be passed by ref
func learnMutex() {
	var m sync.Mutex
	var wg sync.WaitGroup
	t := TestNumber{
		Value: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		t.incrementValue(&m, wg)
	}
	wg.Wait()
	fmt.Println(t.Value)
}

package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/singleton"
	"sync"
)

func main() {
	// add wg wait
	wg := sync.WaitGroup{}

	howMany := 30
	for i := 0; i < howMany; i++ {
		wg.Add(1)
		go func(counter int) {
			s := singleton.GetCounter()
			for n := 0; n < counter; n++ {
				s.Increase()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	s := singleton.GetCounter()
	s.Increase()
	fmt.Println("s.N: ", s.N)
}

package singleton

import (
	"fmt"
	"sync"
)

var createLock = &sync.Mutex{}
var incLock = &sync.Mutex{}

type singleCounter struct {
	N int
}

func (receiver *singleCounter) Increase() int {
	if receiver == nil {
		receiver = GetCounter()
		return 0
	}
	incLock.Lock()
	defer incLock.Unlock()
	receiver.N++
	return receiver.N
}

var singleInstance *singleCounter

func GetCounter() *singleCounter {
	if singleInstance == nil {
		createLock.Lock()
		defer createLock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating singleCounter instance now.")
			singleInstance = &singleCounter{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

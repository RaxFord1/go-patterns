package main

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/observer"
	"time"
)

func main() {
	c1 := observer.NewClient("Client 1", 13)
	c2 := observer.NewClient("Client 2", 23)
	c3 := observer.NewClient("Client 3", 42)

	autoML := observer.AutoMLTrainer{}

	autoML.AddSubscriber(c1)
	autoML.AddSubscriber(c2)
	autoML.AddSubscriber(c3)

	err := autoML.StartTraining()
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)
	c1.UpdateName("qwe")

	time.Sleep(122 * time.Second)
}

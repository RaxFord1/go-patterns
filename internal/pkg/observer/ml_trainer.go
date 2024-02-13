package observer

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/adapter/datasets"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
	"time"
)

type MLTrainer interface {
	SetModel(*ml_model.Model)
	SetDataset(dataset *datasets.IDataset)
	Status() bool // true - training in progress
	StartTraining() error
}

// training ml models automatically and updating clients, when training is ready

type AutoMLTrainer struct {
	isTraining  bool
	Subscribers []ISubscriber

	Model   *ml_model.Model
	Dataset *datasets.IDataset
}

func (a *AutoMLTrainer) Status() bool {
	return a.isTraining
}

func (a *AutoMLTrainer) StartTraining() error {
	//if a.Model == nil {
	//	return errors.New("empty model")
	//}
	//if a.Dataset == nil {
	//	return errors.New("empty dataset")
	//}

	// doing some processing
	go func() {
		time.Sleep(5 * time.Second)

		// Notify everybody about end of training
		a.isTraining = true
		a.Notify()
	}()
	return nil
}

func (a *AutoMLTrainer) SetModel(model *ml_model.Model) {
	a.Model = model
}

func (a *AutoMLTrainer) SetDataset(dataset *datasets.IDataset) {
	a.Dataset = dataset
}

func (a *AutoMLTrainer) AddSubscriber(subscriber ISubscriber) {
	if subscriber == nil {
		return
	}
	a.Subscribers = append(a.Subscribers, subscriber)
}

func (a *AutoMLTrainer) RemoveSubscriber(subscriber ISubscriber) {
	if subscriber == nil {
		return
	}
	for i, s := range a.Subscribers {
		if s == subscriber {
			a.Subscribers = append(a.Subscribers[:i], a.Subscribers[i+1:]...)
		}
	}
}

func (a *AutoMLTrainer) Notify() {
	for _, subscriber := range a.Subscribers {
		subscriber.Update(context.Background())
	}
}

package pipelines

import "fmt"

// SupervisedLearningPipeline - Concrete implementation
type SupervisedLearningPipeline struct{}

func (s *SupervisedLearningPipeline) PreprocessData() {
	fmt.Println("Preprocessing data for supervised learning.")
}

func (s *SupervisedLearningPipeline) TrainModel() {
	fmt.Println("Training model using supervised learning.")
}

func (s *SupervisedLearningPipeline) EvaluateModel() {
	fmt.Println("Evaluating supervised learning model.")
}

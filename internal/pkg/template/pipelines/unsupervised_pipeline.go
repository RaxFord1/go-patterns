package pipelines

import "fmt"

// UnsupervisedLearningPipeline - Another concrete implementation
type UnsupervisedLearningPipeline struct{}

func (u *UnsupervisedLearningPipeline) PreprocessData() {
	fmt.Println("Preprocessing data for unsupervised learning.")
}

func (u *UnsupervisedLearningPipeline) TrainModel() {
	fmt.Println("Training model using unsupervised learning.")
}

func (u *UnsupervisedLearningPipeline) EvaluateModel() {
	fmt.Println("Evaluating unsupervised learning model.")
}

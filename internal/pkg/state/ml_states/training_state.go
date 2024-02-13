package ml_states

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/state"
)

type TrainingState struct{}

func (t *TrainingState) ProcessData(context *state.MLProcess) {
	fmt.Println("Currently training model...")
}

func (t *TrainingState) TrainModel(context *state.MLProcess) {
	fmt.Println("Training model in Training State.")
	// Transition to the next state: Evaluation
	context.SetState(&EvaluationState{})
}

func (t *TrainingState) EvaluateModel(context *state.MLProcess) {
	fmt.Println("Currently training model...")
}
func (t *TrainingState) DeployModel(context *state.MLProcess) {
	fmt.Println("Currently training model...")
}

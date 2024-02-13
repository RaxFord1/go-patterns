package ml_states

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/state"
)

type EvaluationState struct{}

func (e *EvaluationState) ProcessData(context *state.MLProcess) {
	fmt.Println("Currently evaluating model...")
}
func (e *EvaluationState) TrainModel(context *state.MLProcess) {
	fmt.Println("Currently evaluating model...")
}

func (e *EvaluationState) EvaluateModel(context *state.MLProcess) {
	fmt.Println("Evaluating model in Evaluation State.")
	// Transition to the next state: Deployment
	context.SetState(&DeploymentState{})
}

func (e *EvaluationState) DeployModel(context *state.MLProcess) {
	fmt.Println("Currently evaluating model...")
}

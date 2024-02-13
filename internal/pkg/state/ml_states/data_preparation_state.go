package ml_states

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/state"
)

type DataPreparationState struct{}

func (d *DataPreparationState) ProcessData(context *state.MLProcess) {
	fmt.Println("Processing data in Data Preparation State.")
	// Transition to the next state: Training
	context.SetState(&TrainingState{})
}

func (d *DataPreparationState) TrainModel(context *state.MLProcess) {
	fmt.Println("Currently processing data...")
}
func (d *DataPreparationState) EvaluateModel(context *state.MLProcess) {
	fmt.Println("Currently processing data...")
}
func (d *DataPreparationState) DeployModel(context *state.MLProcess) {
	fmt.Println("Currently processing data...")
}

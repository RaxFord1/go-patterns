package ml_states

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/state"
)

type DeploymentState struct{}

func (d *DeploymentState) ProcessData(context *state.MLProcess) {
	fmt.Println("Currently deploying model...")
}
func (d *DeploymentState) TrainModel(context *state.MLProcess) {
	fmt.Println("Currently deploying model...")
}
func (d *DeploymentState) EvaluateModel(context *state.MLProcess) {
	fmt.Println("Currently deploying model...")
}

func (d *DeploymentState) DeployModel(context *state.MLProcess) {
	fmt.Println("Deploying model in Deployment State.")
	// Here, you might end the process or set a new state
}

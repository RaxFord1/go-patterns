package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/state/ml_states"
)

func main() {
	mlProcess := ml_states.NewMLProcess()

	// Start with data preparation
	fmt.Println("Starting Data Preparation...")
	mlProcess.ProcessData()

	// doesn't change a state
	mlProcess.EvaluateModel()
	mlProcess.DeployModel()

	// Proceed to model training
	fmt.Println("Proceeding to Model Training...")
	mlProcess.TrainModel()

	// doesn't change a state
	mlProcess.ProcessData()
	mlProcess.DeployModel()

	// Move to model evaluation
	fmt.Println("Proceeding to Model Evaluation...")
	mlProcess.EvaluateModel()

	// Finally, deploy the model
	fmt.Println("Proceeding to Model Deployment...")
	mlProcess.DeployModel()

	fmt.Println("ML process completed.")
}

package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/builder/data_models"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/flyweights"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
	"sync"
)

func runModel(model ml_model.Model, data []float64, wg *sync.WaitGroup) {
	defer wg.Done()

	// Perform some operations with the model
	result := model.Predict(data)
	fmt.Printf("Model result: %v\n", result)
}

func main() {
	// Create a weights factory
	weightsFactory := flyweights.NewModelWeightsFlyweightFactory()

	// Load shared weights for models
	modelName := "model-1"

	// Create an array of data to process
	dataToProcess := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{45, 78.564, 442, 6878, 9512, 9},
		{43, 23, 564},
	}

	var wg sync.WaitGroup

	for _, data := range dataToProcess {
		// Initialize each model with shared weights
		model := data_models.NewRegressionModel() // Replace with your model initialization
		model.LoadWeights(weightsFactory.GetModelWeights(modelName))
		wg.Add(1)
		// Run each model in a separate goroutine
		go runModel(model, data, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All models have completed processing.")

	modelName2 := "model-2"
	for _, data := range dataToProcess {
		// Initialize each model with shared weights
		model := data_models.NewRegressionModel() // Replace with your model initialization
		model.LoadWeights(weightsFactory.GetModelWeights(modelName2))
		wg.Add(1)
		// Run each model in a separate goroutine
		go runModel(model, data, &wg)
	}

	wg.Wait()
	fmt.Println("All models have completed processing.")

}

package main

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/builder/builders"
)

func main() {
	// Example data
	var data [][]float64
	var target []float64

	// Create a regression builder
	regressionBuilder := builders.NewRegressionBuilder()

	// Create a director with the builder
	director := builders.NewDirector(regressionBuilder)

	// Construct the model using the director
	model := director.Construct(data, target)

	// Use the model
	model.Predict([]float64{})

	// Classification example
	classificationBuilder := builders.NewClassificationBuilder()

	// Create a director with the builder
	director = builders.NewDirector(classificationBuilder)

	// Construct the model using the director
	model = director.Construct(data, target)

	// Use the model
	model.Predict([]float64{})
}

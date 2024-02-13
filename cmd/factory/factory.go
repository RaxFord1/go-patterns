package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/factory"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
)

func main() {
	// create all models
	linearRegression, err := factory.NewModelFactory("linear_regression")
	if err != nil {
		fmt.Println(err)
	}
	neuralNetwork, err := factory.NewModelFactory("neural_network")
	if err != nil {
		fmt.Println(err)
	}
	decisionTree, err := factory.NewModelFactory("decision_tree")
	if err != nil {
		fmt.Println(err)
	}

	// create data and labels
	data := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	labels := []float64{
		1,
		2,
		3,
		4,
	}

	// train all models with for loop
	for _, model := range []ml_model.Model{linearRegression, neuralNetwork, decisionTree} {
		model.Train(data, labels)
	}

	fmt.Println()

	// predict all models with for loop
	for _, model := range []ml_model.Model{linearRegression, neuralNetwork, decisionTree} {
		model.Predict(data[0])
	}
}

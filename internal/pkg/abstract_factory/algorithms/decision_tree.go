package algorithms

import "fmt"

// DecisionTree is a concrete implementation of Algorithm for classification
type DecisionTree struct{}

func (dt *DecisionTree) Train() {
	fmt.Println("Тренування дерева рішень")
}

func (dt *DecisionTree) Predict() string {
	return "Результат дерева рішень"
}

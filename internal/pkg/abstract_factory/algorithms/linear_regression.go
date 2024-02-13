package algorithms

import "fmt"

// LinearRegression is a concrete implementation of Algorithm for regression
type LinearRegression struct{}

func (lr *LinearRegression) Train() {
	fmt.Println("Тренування лінійної регресії")
}

func (lr *LinearRegression) Predict() string {
	return "Результат лінійної регресії"
}

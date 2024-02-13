package datasets

import "fmt"

// RegressionDataset is a concrete implementation of Dataset for regression
type RegressionDataset struct{}

func (rd *RegressionDataset) LoadData() {
	fmt.Println("Завантаження даних для регресії")
}

func (rd *RegressionDataset) Preprocess() {
	fmt.Println("Попередня обробка даних для регресії")
}

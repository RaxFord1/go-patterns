package datasets

import "fmt"

// ClassificationDataset is a concrete implementation of Dataset for classification
type ClassificationDataset struct{}

func (cd *ClassificationDataset) LoadData() {
	fmt.Println("Завантаження даних для класифікації")
}

func (cd *ClassificationDataset) Preprocess() {
	fmt.Println("Попередня обробка даних для класифікації")
}

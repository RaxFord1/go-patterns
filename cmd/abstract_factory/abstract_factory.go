package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/abstract_factory/factories"
)

// Функція для запуску сценарію машинного навчання
func runMachineLearningScenario(factory factories.AbstractFactory) {
	// Використання фабрики для створення алгоритму та датасету
	algorithm := factory.CreateAlgorithm()
	dataset := factory.CreateDataset()

	// Робота з датасетом
	dataset.LoadData()
	dataset.Preprocess()

	// Робота з алгоритмом
	algorithm.Train()
	result := algorithm.Predict()

	// Вивід результату
	fmt.Println(result)
}

func main() {
	fmt.Println()
	// Створюємо фабрику для регресійних завдань
	regressionFactory := &factories.RegressionFactory{}
	fmt.Println("Регресійний сценарій:")
	runMachineLearningScenario(regressionFactory)

	fmt.Println()
	// Створюємо фабрику для класифікаційних завдань
	classificationFactory := &factories.ClassificationFactory{}
	fmt.Println("Класифікаційний сценарій:")
	runMachineLearningScenario(classificationFactory)
}

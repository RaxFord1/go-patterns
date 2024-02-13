package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/adapter"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/adapter/datasets"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/adapter/models"
)

func main() {
	predictor := models.NewPredictor([]float64{1, 2, 3})
	fmt.Println(predictor.Predict([]float64{1, 2, 3}))

	dataset := &datasets.CSVDataset{}
	err := dataset.LoadData([][]interface{}{{0.2, 1, 2, 3, "A", "B", "C"}, {0.8, 4, 5, 2, "D", "E", "F"}})
	if err != nil {
		return
	}

	mapName := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
	}
	inputTypes := []adapter.InputTypes{
		adapter.Float, adapter.Int, adapter.Int, adapter.Int, adapter.String, adapter.String, adapter.String,
	}

	hotEncoder := adapter.NewCsvToPredictorAdapter(mapName, inputTypes, predictor, dataset)
	result, err := hotEncoder.PredictWithCSVData()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

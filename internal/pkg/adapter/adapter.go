package adapter

import (
	"errors"
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/adapter/datasets"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/adapter/models"
)

type InputTypes int

const (
	None   InputTypes = iota
	String            // needs to be hot-encoded
	Int               // needs to be hot-encoded
	Float
)

func (it InputTypes) CompareType(inpType interface{}) error {
	inpTypeEquivalent := None
	switch inpType.(type) {
	case string:
		inpTypeEquivalent = String
	case int:
		inpTypeEquivalent = Int
	case float64:
		inpTypeEquivalent = Float
	default:
		return errors.New("invalid input type")
	}
	if it != inpTypeEquivalent {
		return errors.New("invalid input type")
	}
	return nil
}

type CsvToPredictorAdapter struct {
	nameMap    map[string]int
	inputTypes []InputTypes
	predictor  models.IPredictor
	dataset    datasets.IDataset
}

func (cd *CsvToPredictorAdapter) LoadData(dataset [][]interface{}) error {
	return cd.dataset.LoadData(dataset)
}

func (cd *CsvToPredictorAdapter) GetData() [][]interface{} {
	return cd.dataset.GetData()
}

func (cd *CsvToPredictorAdapter) stringToInt(s string) (float64, error) {
	if val, ok := cd.nameMap[s]; ok {
		return float64(val), nil
	}
	return 0, errors.New("not Found value for string: " + s)
}

func (cd *CsvToPredictorAdapter) prepareItem(v interface{}) ([]float64, error) {
	switch v.(type) {
	case string:
		if val, ok := cd.nameMap[v.(string)]; ok {
			return OneHotEncoder(val, len(cd.nameMap))
		}
		return nil, errors.New("not Found value for string: " + v.(string))
	case int:
		return OneHotEncoder(v.(int), len(cd.nameMap))
	case float64:
		return []float64{v.(float64)}, nil
	default:
		return nil, errors.New("wrong type of input")
	}
}

func (cd *CsvToPredictorAdapter) PrepareRow(data []interface{}) ([]float64, error) {
	if len(data) != len(cd.inputTypes) {
		return nil, errors.New("wrong len of inputs")
	}

	var row []float64
	for idx, v := range data {
		// compare Types of data and InputTypes
		if err := cd.inputTypes[idx].CompareType(v); err != nil {
			return nil, err
		}

		preparedItem, err := cd.prepareItem(v)
		if err != nil {
			return nil, err
		}
		row = append(row, preparedItem...)
	}
	return row, nil
}

func (cd *CsvToPredictorAdapter) PredictWithCSVData() ([]float64, error) {
	fmt.Println("Завантаження даних для класифікації")

	data := cd.dataset.GetData()
	result := make([]float64, len(data))
	for idx, row := range data {
		preparedRow, err := cd.PrepareRow(row)
		if err != nil {
			return nil, err
		}

		callResult := cd.predictor.Predict(preparedRow)
		result[idx] = callResult
	}

	return result, nil
}

func OneHotEncoder(inp int, length int) ([]float64, error) {
	if inp >= length || inp < 0 {
		return nil, errors.New("wrong input. inp > length || inp < 0")
	}
	result := make([]float64, length)
	result[inp] = 1
	return result, nil
}

func NewCsvToPredictorAdapter(nameMap map[string]int, inputTypes []InputTypes, predictor models.IPredictor, dataset datasets.IDataset) *CsvToPredictorAdapter {
	return &CsvToPredictorAdapter{
		nameMap:    nameMap,
		inputTypes: inputTypes,
		predictor:  predictor,
		dataset:    dataset,
	}
}

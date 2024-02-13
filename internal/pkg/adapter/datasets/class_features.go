package datasets

import "fmt"

type IDataset interface {
	LoadData(dataset [][]interface{}) error
	GetData() [][]interface{}
}

type CSVDataset struct {
	dataset [][]interface{}
}

func (cd *CSVDataset) LoadData(dataset [][]interface{}) error {
	fmt.Println("Завантаження даних")
	cd.dataset = dataset
	return nil
}

func (cd *CSVDataset) GetData() [][]interface{} {
	return cd.dataset
}

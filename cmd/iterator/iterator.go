package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/flyweights"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/iterator"
	"path/filepath"
)

var weightsFactory = flyweights.NewModelWeightsFlyweightFactory()

func ProcessFile(filePath string) {
	fmt.Printf("Processing file: %s\n", filePath)
	// Define the file path based on the modelName
	modelNameBase := filepath.Base(filePath)
	modelNameFileNameExtension := filepath.Ext(modelNameBase)
	modelName := modelNameBase[0 : len(modelNameBase)-len(modelNameFileNameExtension)]
	weightsFactory.GetModelWeights(modelName)
}

// iterate over directory concurrently loading weights from each file
func main() {
	path := "D:\\University\\4\\2\\Архітектура та проєктування програмного забезпечення\\labs\\labs\\assets\\model_weights"
	iter, err := iterator.NewDirectoryIterator(path)
	if err != nil {
		panic(err)
	}

	for iter.HasNext() {
		file := iter.GetNext().(string)

		fmt.Println(file)
		go ProcessFile(file)
	}
}

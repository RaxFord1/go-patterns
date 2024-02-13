package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/bridge/data_handlers"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/bridge/execution_environments"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/builder/data_models"
)

var wordToNumber = map[string]float64{
	"the": 1.0, "of": 2.0, "and": 3.0, "a": 4.0, "to": 5.0,
	"in": 6.0, "is": 7.0, "you": 8.0, "that": 9.0, "it": 10.0,
	"he": 11.0, "was": 12.0, "for": 13.0, "on": 14.0, "are": 15.0,
	"as": 16.0, "with": 17.0, "his": 18.0, "they": 19.0, "I": 20.0,
	"at": 21.0, "be": 22.0, "this": 23.0, "have": 24.0, "from": 25.0,
	"or": 26.0, "one": 27.0, "had": 28.0, "by": 29.0, "word": 30.0,
	"but": 31.0, "not": 32.0, "what": 33.0, "all": 34.0, "were": 35.0,
	"we": 36.0, "when": 37.0, "your": 38.0, "can": 39.0, "said": 40.0,
	"there": 41.0, "use": 42.0, "an": 43.0, "each": 44.0, "which": 45.0,
	"she": 46.0, "do": 47.0, "how": 48.0, "their": 49.0, "if": 50.0,
	"quick": 51.0, "brown": 52.0, "fox": 53.0, "jumps": 54.0, "over": 55.0,
	"lazy": 56.0, "dog": 57.0, "then": 58.0, "runs": 59.0, "other": 60.0,
	"side": 61.0, "where": 62.0, "finds": 63.0, "know": 64.0,
}

func main() {
	// ml models
	classificationModel := data_models.NewRegressionModel()
	nlpModel := data_models.NewNLPModel(wordToNumber)

	// environments
	cpuEnv := &execution_environments.CPUEnvironment{
		CPUName: "CPU-1",
	}
	gpuEnv := &execution_environments.GPUEnvironment{
		GPUName: "GPU-1",
	}
	distributedEnv := &execution_environments.DistributedEnvironment{
		CPUName: []string{"CPU-1", "CPU-2"},
		GPUName: []string{"GPU-1", "GPU-2", "GPU-3", "GPU-4"},
	}

	// Structured data handlers
	dataHandler := data_handlers.StructuredDataHandlers{}
	dataHandler.SetEnvironment(cpuEnv)
	dataHandler.SetModel(classificationModel)

	inputData := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Handle data
	result := dataHandler.HandleData(inputData)
	fmt.Println(result)

	// Change the environment to GPU
	dataHandler.SetEnvironment(gpuEnv)

	// Handle the same data with a different environment
	newResult := dataHandler.HandleData(inputData)
	fmt.Println(newResult)

	dataHandler.SetEnvironment(distributedEnv)
	newResult = dataHandler.HandleData(inputData)
	fmt.Println(newResult)

	// NLP data handlers
	dataHandlerNLP := data_handlers.UnstructuredDataHandlers{}
	dataHandlerNLP.SetEnvironment(distributedEnv)
	dataHandlerNLP.SetModel(nlpModel)
	inputDataNLP := "the quick brown fox jumps over the lazy dog and then runs to the other side where it finds a word but does not know what to do with it"
	dataHandlerNLP.HandleData(inputDataNLP)
	fmt.Println(inputDataNLP + " " + dataHandlerNLP.LastAnswer)
}

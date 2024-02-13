package flyweights

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/iterator"
	"os"
	"strconv"
	"strings"
	"sync"
)

var mtx = &sync.RWMutex{}

type ModelWeights struct {
	Weights []float64
}

// ModelWeightsFlyweightFactory manages the sharing of ModelWeights
type ModelWeightsFlyweightFactory struct {
	weightsCache map[string]*ModelWeights
}

func NewModelWeightsFlyweightFactory() *ModelWeightsFlyweightFactory {
	return &ModelWeightsFlyweightFactory{
		weightsCache: make(map[string]*ModelWeights),
	}
}

func (f *ModelWeightsFlyweightFactory) getModelWeights(modelName string) (*ModelWeights, bool) {
	mtx.RLock()
	defer mtx.RUnlock()
	weights, exists := f.weightsCache[modelName]
	return weights, exists
}

func (f *ModelWeightsFlyweightFactory) GetModelWeights(modelName string) *ModelWeights {
	if weights, exists := f.getModelWeights(modelName); exists {
		fmt.Printf("Returning cached weights for model %s\n", modelName)
		return weights // Return cached weights
	}

	// Assume LoadModelWeights is a function that loads weights for a given model
	weights := LoadModelWeights(modelName)
	mtx.Lock()
	defer mtx.Unlock()
	f.weightsCache[modelName] = weights // Cache the newly loaded weights

	return weights
}

// LoadModelWeights simulates loading model weights (to be replaced with actual implementation)
func LoadModelWeights(modelName string) *ModelWeights {
	fmt.Printf("Loading weights for model: %s\n", modelName)
	// Define the file path based on the modelName
	filePath := "D:\\University\\4\\2\\Архітектура та проєктування програмного забезпечення\\labs\\labs\\assets\\model_weights\\" + modelName + ".weights"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Unable to open the file: " + err.Error())
	}
	defer file.Close()

	fileIterator, err := iterator.NewLineByLineIterator(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer fileIterator.Close()

	var weights []float64
	for fileIterator.HasNext() {
		line := fileIterator.GetNext()
		if line == nil {
			break
		}
		for _, w := range strings.Split(line.(string), " ") {
			weight, err := strconv.ParseFloat(w, 64)
			if err != nil {
				panic("Error parsing weight: " + err.Error())
			}
			weights = append(weights, weight)
		}
	}

	return &ModelWeights{Weights: weights}
}

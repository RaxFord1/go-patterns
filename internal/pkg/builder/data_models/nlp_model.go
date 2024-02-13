package data_models

import (
	"fmt"
	"hash/fnv"
	"strings"
)

// NLPModel is a concrete model type for text tasks
type NLPModel struct {
	MapWordToNumber map[string]float64
	MapNumberToWord map[float64]string
	TrainedModel    interface{} // This would typically be some type of statistical model structure
}

func (r *NLPModel) Tokenize(text string) []string {
	return strings.Fields(text)
}

func (r *NLPModel) Embed(tokens []string) []float64 {
	var embeddings []float64
	for _, token := range tokens {
		if number, exists := r.MapWordToNumber[token]; exists {
			embeddings = append(embeddings, number)
		} else {
			// Handle unknown words, perhaps with a default number or error
			embeddings = append(embeddings, 0.0)
		}
	}
	return embeddings
}

func (r *NLPModel) PredictNextWord(text []float64) float64 {
	hashValue := hashText(text)
	keys := make([]string, 0, len(r.MapWordToNumber))
	for k := range r.MapWordToNumber {
		keys = append(keys, k)
	}
	pseudoRandomIndex := hashValue % uint32(len(keys))
	selectedWord := keys[pseudoRandomIndex]
	return r.MapWordToNumber[selectedWord]
}

func (r *NLPModel) FromEmbedding(embedding float64) string {
	if word, exists := r.MapNumberToWord[embedding]; exists {
		return word
	}
	return "UNKNOWN"
}

// The following methods implement the Model interface for NLPModel

func (r *NLPModel) Train(data [][]float64, labels []float64) {
	// Implement training logic here
	fmt.Println("NLPModel Train data...")
}

func (r *NLPModel) Predict(input []float64) float64 {
	// Implement prediction logic here
	// For the sake of example, return a dummy value
	fmt.Println("NLPModel Predict data...")
	return 0.0
}

func hashText(text []float64) uint32 {
	h := fnv.New32a()
	for _, num := range text {
		// Converting float64 to byte slice
		bs := make([]byte, 8)
		for i := range bs {
			bs[i] = byte(int(num) >> (8 * i))
		}
		h.Write(bs)
	}
	return h.Sum32()
}

// NewNLPModel creates a new instance of a NLPModel
func NewNLPModel(wordToNumber map[string]float64) *NLPModel {
	var numberToWord = make(map[float64]string)
	for word, number := range wordToNumber {
		numberToWord[number] = word
	}
	return &NLPModel{
		MapWordToNumber: wordToNumber,
		MapNumberToWord: numberToWord,
	}
}

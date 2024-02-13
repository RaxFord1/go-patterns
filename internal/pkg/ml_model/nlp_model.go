package ml_model

// NLPModel defines the interface for an NLP model with typical NLP functions
type NLPModel interface {
	Tokenize(text string) []string
	Embed(tokens []string) []float64
	PredictNextWord(text []float64) float64
	FromEmbedding(embeddings float64) string
	Model
}

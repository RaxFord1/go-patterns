package ml_model

type Model interface {
	Train(data [][]float64, labels []float64)
	Predict(data []float64) float64
}

type ModelWithLoadWeights interface {
	LoadWeights(weights []float64)
	Model
}

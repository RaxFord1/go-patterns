package ml_model

// Model is the complex object that is being built
type StructuredDataModel interface {
	SetData(data [][]float64, target []float64)
	Preprocess()
	TuneHyperparameters()
	Model
}

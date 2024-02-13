package builders

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
)

// ModelBuilder is the builder interface for creating machine learning models
type ModelBuilder interface {
	SetData(data [][]float64, target []float64) ModelBuilder
	Preprocess() ModelBuilder
	TuneHyperparameters() ModelBuilder
	TrainModel() ModelBuilder
	Build() ml_model.StructuredDataModel
}

// Director orchestrates the building steps
type Director struct {
	builder ModelBuilder
}

func NewDirector(b ModelBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Construct(data [][]float64, target []float64) ml_model.StructuredDataModel {
	return d.builder.SetData(data, target).
		Preprocess().
		TuneHyperparameters().
		TrainModel().
		Build()
}

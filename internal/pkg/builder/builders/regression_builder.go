package builders

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/builder/data_models"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
)

// RegressionBuilder is a concrete builder for regression models
type RegressionBuilder struct {
	model ml_model.StructuredDataModel
}

func NewRegressionBuilder() *RegressionBuilder {
	// Initialize with a regression model
	return &RegressionBuilder{model: data_models.NewRegressionModel()}
}

func (b *RegressionBuilder) SetData(data [][]float64, target []float64) ModelBuilder {
	// Set the data for the model
	b.model.SetData(data, target)
	return b
}

func (b *RegressionBuilder) Preprocess() ModelBuilder {
	// Preprocess the data
	b.model.Preprocess()
	return b
}

func (b *RegressionBuilder) TuneHyperparameters() ModelBuilder {
	// Tune hyperparameters for the model
	b.model.TuneHyperparameters()
	return b
}

func (b *RegressionBuilder) TrainModel() ModelBuilder {
	// Train the model
	b.model.Train()
	return b
}

func (b *RegressionBuilder) Build() ml_model.StructuredDataModel {
	// Return the completed model
	return b.model
}

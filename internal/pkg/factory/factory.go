package factory

import (
	"errors"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/factory/factory_models"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
)

type ModelFactory interface {
	CreateModel() ml_model.Model
}

func NewModelFactory(modelType string) (ml_model.Model, error) {
	switch modelType {
	case "linear_regression":
		return factory_models.NewLinearRegression(), nil
	case "neural_network":
		return factory_models.NewNeuralNetwork(), nil
	case "decision_tree":
		return factory_models.NewDecisionTree(), nil
	default:
		return nil, errors.New("unknown model type")
	}
}

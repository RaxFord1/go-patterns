package factories

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/abstract_factory/algorithms"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/abstract_factory/datasets"
)

// RegressionFactory is a concrete factory for regression tasks
type RegressionFactory struct{}

func (r *RegressionFactory) CreateAlgorithm() algorithms.Algorithm {
	return &algorithms.LinearRegression{}
}

func (r *RegressionFactory) CreateDataset() datasets.Dataset {
	return &datasets.RegressionDataset{}
}

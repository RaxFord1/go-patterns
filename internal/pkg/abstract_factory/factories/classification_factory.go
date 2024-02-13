package factories

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/abstract_factory/algorithms"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/abstract_factory/datasets"
)

// ClassificationFactory is a concrete factory for classification tasks
type ClassificationFactory struct{}

func (c *ClassificationFactory) CreateAlgorithm() algorithms.Algorithm {
	return &algorithms.DecisionTree{}
}

func (c *ClassificationFactory) CreateDataset() datasets.Dataset {
	return &datasets.ClassificationDataset{}
}

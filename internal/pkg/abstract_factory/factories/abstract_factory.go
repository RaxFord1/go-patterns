package factories

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/abstract_factory/algorithms"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/abstract_factory/datasets"
)

// AbstractFactory is an interface for our abstract factory
type AbstractFactory interface {
	CreateAlgorithm() algorithms.Algorithm
	CreateDataset() datasets.Dataset
}

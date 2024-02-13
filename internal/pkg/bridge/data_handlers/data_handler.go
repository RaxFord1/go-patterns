package data_handlers

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/bridge/execution_environments"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
)

type IDataHandler interface {
	HandleData(data interface{}) float64
	SetEnvironment(env execution_environments.ExecutionEnvironment)
	SetModel(model ml_model.Model)
}

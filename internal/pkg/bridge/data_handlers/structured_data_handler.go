package data_handlers

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/bridge/execution_environments"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
)

type StructuredDataHandlers struct {
	Environment execution_environments.ExecutionEnvironment
	Model       ml_model.StructuredDataModel
}

func (handler *StructuredDataHandlers) HandleData(data interface{}) float64 {
	switch v := data.(type) {
	case []float64:
		return handler.Environment.ExecuteTask(handler.Model.Predict, v)
	default:
		panic("Data type not supported")
	}
}

func (handler *StructuredDataHandlers) SetEnvironment(env execution_environments.ExecutionEnvironment) {
	handler.Environment = env
}

func (handler *StructuredDataHandlers) SetModel(model ml_model.Model) {
	if model, ok := model.(ml_model.StructuredDataModel); ok {
		handler.Model = model
	} else {
		panic("Model type not supported")
	}
}

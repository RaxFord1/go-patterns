package data_handlers

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/bridge/execution_environments"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/ml_model"
)

type UnstructuredDataHandlers struct {
	Environment execution_environments.ExecutionEnvironment
	Model       ml_model.NLPModel
	LastAnswer  string
}

func (handler *UnstructuredDataHandlers) HandleData(data interface{}) float64 {
	var strData string
	var ok bool
	if strData, ok = data.(string); !ok {
		panic("Data type not supported")
	}

	tokenizedData := handler.Model.Tokenize(strData)
	embeddedData := handler.Model.Embed(tokenizedData)
	result := handler.Environment.ExecuteTask(handler.Model.PredictNextWord, embeddedData)
	handler.LastAnswer = handler.Model.FromEmbedding(result)

	return result
}

func (handler *UnstructuredDataHandlers) SetEnvironment(env execution_environments.ExecutionEnvironment) {
	handler.Environment = env
}

func (handler *UnstructuredDataHandlers) SetModel(model ml_model.Model) {
	if model, ok := model.(ml_model.NLPModel); ok {
		handler.Model = model
	} else {
		panic("Model type not supported")
	}
}

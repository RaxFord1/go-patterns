package template

// IMLPipeline - Template pattern interface
type IMLPipeline interface {
	PreprocessData()
	TrainModel()
	EvaluateModel()
}

type MlPipeline struct {
	mlPipeline IMLPipeline
}

func (o *MlPipeline) RunPipeline() error {
	o.mlPipeline.PreprocessData()
	o.mlPipeline.TrainModel()
	o.mlPipeline.EvaluateModel()
	return nil
}
func (o *MlPipeline) ChangePipeline(pipeline IMLPipeline) {
	o.mlPipeline = pipeline
}

package execution_environments

type ExecuteFunc func(floatArr []float64) float64

type ExecutionEnvironment interface {
	ExecuteTask(funcToRun ExecuteFunc, inputData []float64) float64
}

package execution_environments

import "fmt"

type GPUEnvironment struct {
	GPUName string
}

func (gpu *GPUEnvironment) ExecuteTask(funcToRun ExecuteFunc, inputData []float64) float64 {
	// CPU-specific execution logic
	fmt.Println("Executing task on: " + gpu.GPUName)
	return funcToRun(inputData)
}

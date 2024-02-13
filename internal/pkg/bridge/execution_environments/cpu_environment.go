package execution_environments

import "fmt"

type CPUEnvironment struct {
	CPUName string
}

func (cpu *CPUEnvironment) ExecuteTask(funcToRun ExecuteFunc, inputData []float64) float64 {
	// CPU-specific execution logic
	fmt.Println("Executing task on: " + cpu.CPUName)
	return funcToRun(inputData)
}

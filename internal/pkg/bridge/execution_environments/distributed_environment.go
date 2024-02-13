package execution_environments

import (
	"fmt"
	"strings"
)

type DistributedEnvironment struct {
	// CPU-specific fields
	GPUName []string
	CPUName []string
}

func (distriburedEnvs *DistributedEnvironment) ExecuteTask(funcToRun ExecuteFunc, inputData []float64) float64 {
	fmt.Println("Executing task on CPUs: " + distriburedEnvs.StringCPU() + "and GPUs: " + distriburedEnvs.StringGPU())
	return funcToRun(inputData)
}

func (distriburedEnvs *DistributedEnvironment) StringCPU() string {
	sb := strings.Builder{}
	for idx, cpu := range distriburedEnvs.CPUName {
		if idx != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(cpu)
	}

	return sb.String()
}

func (distriburedEnvs *DistributedEnvironment) StringGPU() string {
	sb := strings.Builder{}
	for idx, gpu := range distriburedEnvs.GPUName {
		if idx != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(gpu)
	}

	return sb.String()
}

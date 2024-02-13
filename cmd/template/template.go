package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/template"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/template/pipelines"
)

func main() {
	mlPipeline := &template.MlPipeline{}

	supervised := &pipelines.SupervisedLearningPipeline{}
	mlPipeline.ChangePipeline(supervised)
	if err := mlPipeline.RunPipeline(); err != nil {
		fmt.Println(err)
	}

	unsupervised := &pipelines.UnsupervisedLearningPipeline{}
	mlPipeline.ChangePipeline(unsupervised)
	if err := mlPipeline.RunPipeline(); err != nil {
		fmt.Println(err)
	}
}

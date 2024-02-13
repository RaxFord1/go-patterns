package main

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/prototype"
)

func main() {
	nn := prototype.NewNN([][]float64{{1, 2, 3}, {4, 5}}, "nn")
	fmt.Println(nn)

	fmt.Println()

	nnClone := nn.Clone()
	fmt.Println(nnClone)
}

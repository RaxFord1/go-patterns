package prototype

import (
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/singleton"
	"strings"
)

type NN struct {
	layers []Prototype
	Name   string
	count  int
}

func (r NN) Clone() Prototype {
	newNN := NN{}
	newNN.layers = make([]Prototype, len(r.layers))
	for i, layer := range r.layers {
		newNN.layers[i] = layer.Clone()
	}

	newNN.count = singleton.GetCounter().Increase()

	newNN.Name = r.Name

	return newNN
}

func NewNN(params [][]float64, name string) NN {
	layers := make([]Prototype, len(params))
	for i, param := range params {
		layers[i] = NewLayer(param, i)
	}
	return NN{
		layers: layers,
		Name:   name,
	}
}

func (r NN) String() string {
	namePrefix := r.Name
	if r.count != 0 {
		namePrefix = fmt.Sprintf("%s-%d", r.Name, r.count)
	}
	layersArr := make([]string, len(r.layers))
	for i, layer := range r.layers {
		layersArr[i] = layer.String()
	}
	layersString := "{" + strings.Join(layersArr, ", ") + "}"
	return fmt.Sprintf("NN: {Layers: %s, Name: %s, count: %d}", layersString, namePrefix, r.count)
}

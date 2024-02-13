package prototype

import "fmt"

type Layers struct {
	params []float64
	n      int
}

func (r Layers) Clone() Prototype {
	newLayers := make([]float64, len(r.params))
	copy(newLayers, r.params)

	return Layers{
		params: newLayers,
		n:      r.n,
	}
}

func (r Layers) String() string {
	return fmt.Sprintf("%d: {params: %v}", r.n, r.params)
}

func NewLayer(params []float64, n int) Layers {
	return Layers{
		params: params,
		n:      n,
	}
}

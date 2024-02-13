package models

type IPredictor interface {
	Predict(x []float64) float64
}

type Predictor struct {
	parameters []float64
}

func (p *Predictor) Predict(x []float64) float64 {
	sum := 0.0
	for i := 0; i < len(p.parameters); i++ {
		for j := 0; j < len(x); j++ {
			sum += p.parameters[i] * x[j]
		}
	}
	return sum
}

func NewPredictor(parameters []float64) *Predictor {
	return &Predictor{
		parameters: parameters,
	}
}

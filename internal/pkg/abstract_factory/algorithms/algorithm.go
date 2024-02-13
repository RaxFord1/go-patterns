package algorithms

// Algorithm is an interface for machine learning algorithms
type Algorithm interface {
	Train()
	Predict() string
}

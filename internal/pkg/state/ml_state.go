package state

type MLState interface {
	ProcessData(context *MLProcess)
	TrainModel(context *MLProcess)
	EvaluateModel(context *MLProcess)
	DeployModel(context *MLProcess)
}

type MLProcess struct {
	State MLState
}

func (m *MLProcess) SetState(state MLState) {
	m.State = state
}

func (m *MLProcess) ProcessData() {
	m.State.ProcessData(m)
}

func (m *MLProcess) TrainModel() {
	m.State.TrainModel(m)
}

func (m *MLProcess) EvaluateModel() {
	m.State.EvaluateModel(m)
}

func (m *MLProcess) DeployModel() {
	m.State.DeployModel(m)
}

package ml_states

import "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/state"

func NewMLProcess() *state.MLProcess {
	return &state.MLProcess{State: &DataPreparationState{}}
}

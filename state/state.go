package state

//go:generate bash -c "rm -Rf mocks"
//go:generate mockery --all --case snake --dir . --output ./mocks --outpkg mock_state --disable-version-string --with-expecter

import (
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/model"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/storage"
)

type State struct {
	*model.TxManager
	*model.ForkIdState
	*model.L1InfoTreeState
	*model.BatchState
	storage.BlockStorer
}

func NewState(storageImpl storage.Storer) *State {
	res := &State{

		model.NewTxManager(storageImpl),
		model.NewForkIdState(storageImpl),
		model.NewL1InfoTreeManager(storageImpl),
		model.NewBatchState(storageImpl),
		storageImpl,
	}
	return res
}
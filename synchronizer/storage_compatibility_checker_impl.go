package synchronizer

import (
	"context"

	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/common/syncinterfaces"
)

type StorageCompatibilityCheckerImpl struct {
	stateCompatibility   syncinterfaces.StateStorageCompatibilityCheckerInterface
	etherman             syncinterfaces.EthermanChainQuerier
	OverrideStorageCheck bool
}

func NewSanityStorageCheckerImpl(stateCompatibility syncinterfaces.StateStorageCompatibilityCheckerInterface,
	etherman syncinterfaces.EthermanChainQuerier,
	overrideStorageCheck bool) *StorageCompatibilityCheckerImpl {
	return &StorageCompatibilityCheckerImpl{
		stateCompatibility:   stateCompatibility,
		etherman:             etherman,
		OverrideStorageCheck: overrideStorageCheck,
	}
}

func (s *StorageCompatibilityCheckerImpl) CheckAndUpdateStorage(ctx context.Context) error {
	localRun := s.getDataFromEtherman()
	return s.stateCompatibility.CheckAndUpdateStorage(ctx, localRun, s.OverrideStorageCheck, nil)
}

func (s *StorageCompatibilityCheckerImpl) getDataFromEtherman() entities.StorageContentsBoundData {
	data := entities.StorageContentsBoundData{
		RollupID:  uint64(s.etherman.GetRollupID()),
		L1ChainID: s.etherman.GetL1ChainID(),
	}
	return data
}

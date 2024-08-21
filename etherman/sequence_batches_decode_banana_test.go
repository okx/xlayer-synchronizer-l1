package etherman

import (
	"encoding/hex"
	"testing"

	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

const (
	txDataBananaHex   = "b910e0f900000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000066c4b6531317f62f023ccd827099d9070ac5eb453040af849485827b42eb6ac92fe112530000000000000000000000005b06837a43bdc3dd9f114558daf4b26ed49842ed00000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000120b0000007a000000000b00000006000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000120b00000006000000000b000000060000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001b0b00000006000000000b00000006000000000b00000006000000000000000000"
	expedtedSeqBanana = `BatchNumber: 2
L1InfoRoot: 0xa958fee09bf3c4d782ab6839f81c8f40316521e19ba088b9c1a580c9a1438506
SequencerAddr: 0x5b06837A43bdC3dD9F114558DAf4B26ed49842Ed
TxHash: 0x2b3fa84d57739604e5fbbc921fca588842923c64e28e0636609ee7ba1bea8c64
Nonce: 1
Coinbase: 0x5b06837A43bdC3dD9F114558DAf4B26ed49842Ed
PolygonZkEVMBatchData: nil
___EtrogSequenceData:ForcedTimestamp: 0
___EtrogSequenceData:ForcedGlobalExitRoot: 0000000000000000000000000000000000000000000000000000000000000000
___EtrogSequenceData:ForcedBlockHashL1: 0000000000000000000000000000000000000000000000000000000000000000
___EtrogSequenceData:Transactions: 0b0000007a000000000b0000000600000000
SequencedBatchElderberryData: nil
BananaData: CounterL1InfoRoot: 1 MaxSequenceTimestamp: 1724167763 ExpectedFinalAccInputHash: 0x1317f62f023ccd827099d9070ac5eb453040af849485827b42eb6ac92fe11253 DataAvailabilityMsg(0):
Metadata: SourceBatchData:  RollupFlavor: kEVM CallFunctionName: sequenceBatchesBanana ForkName: banana
`
)

func TestSequencedBatchBananaDecode(t *testing.T) {
	sut, err := NewDecodeSequenceBatchesBanana()
	require.NoError(t, err)
	require.NotNil(t, sut)
	txData, err := hex.DecodeString(txDataBananaHex)
	require.NoError(t, err)
	lastBatchNumber := uint64(4)
	sequencer := common.HexToAddress("0x5b06837A43bdC3dD9F114558DAf4B26ed49842Ed")
	txHash := common.HexToHash("0x2b3fa84d57739604e5fbbc921fca588842923c64e28e0636609ee7ba1bea8c64")
	nonce := uint64(1)
	l1InfoRoot := common.HexToHash("0xa958fee09bf3c4d782ab6839f81c8f40316521e19ba088b9c1a580c9a1438506")

	res, err := sut.DecodeSequenceBatches(txData, lastBatchNumber, sequencer, txHash, nonce, l1InfoRoot)
	require.NoError(t, err)
	require.NotNil(t, res)
	res0 := res[0].String()
	log.Debug(res0)
	require.Equal(t, expedtedSeqBanana, res0)

}
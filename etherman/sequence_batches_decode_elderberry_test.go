package etherman

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

const (
	txDataElderberryHex = "def57e5400000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000065f838a100000000000000000000000000000000000000000000000000000000000000010000000000000000000000007597b12b953bffe1457d89e7e4fe3da149b45d8800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003cc0b00000890000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000117000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000000000000000000000000000000000000000000"
	expectedElderberry  = `BatchNumber: 2
L1InfoRoot: 0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757
SequencerAddr: 0x7597b12B953Bffe1457d89E7E4fe3dA149b45D88
TxHash: 0x44b7aacaf535bd947803c88c18e63358c8ddd44fbb24950efbb5abb50f938cef
Nonce: 1
Coinbase: 0x7597b12B953Bffe1457d89E7E4fe3dA149b45D88
PolygonZkEVMBatchData: nil
___PolygonRollupBaseEtrogBatchData:ForcedTimestamp: 0
___PolygonRollupBaseEtrogBatchData:ForcedGlobalExitRoot: 0000000000000000000000000000000000000000000000000000000000000000
___PolygonRollupBaseEtrogBatchData:ForcedBlockHashL1: 0000000000000000000000000000000000000000000000000000000000000000
___PolygonRollupBaseEtrogBatchData:Transactions: 0b00000890000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000117000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b00000003000000000b0000000300000000
___SequencedBatchElderberryData:MaxSequenceTimestamp 1710766241
___SequencedBatchElderberryData:InitSequencedBatchNumber 1
Metadata: SourceBatchData:  RollupFlavor: kEVM CallFunctionName: sequenceBatchesElderberry ForkName: elderberry
`
)

func TestSequencedBatchElderberryDecode(t *testing.T) {
	sut, err := NewDecodeSequenceBatchesElderberry()
	require.NoError(t, err)
	require.NotNil(t, sut)
	txData, err := hex.DecodeString(txDataElderberryHex)
	require.NoError(t, err)
	lastBatchNumber := uint64(2)
	sequencer := common.HexToAddress("0x7597b12b953bffe1457d89e7e4fe3da149b45d88")
	txHash := common.HexToHash("0x44b7aacaf535bd947803c88c18e63358c8ddd44fbb24950efbb5abb50f938cef")
	nonce := uint64(1)
	l1InfoRoot := common.HexToHash("0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757")
	res, err := sut.DecodeSequenceBatches(txData, lastBatchNumber, sequencer, txHash, nonce, l1InfoRoot)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 1, len(res))
	res0 := res[0].String()
	require.Equal(t, expectedElderberry, res0)
}

func TestSequencedBatchElderberryMatchMethodId(t *testing.T) {
	sut, err := NewDecodeSequenceBatchesElderberry()
	require.NoError(t, err)
	require.False(t, sut.MatchMethodId([]byte{0x01, 0x02, 0x03, 0x04}))
	require.True(t, sut.MatchMethodId(methodIDSequenceBatchesElderberry))
}

func TestSequencedBatchElderberryNameMethodId(t *testing.T) {
	sut, err := NewDecodeSequenceBatchesElderberry()
	require.NoError(t, err)
	require.Equal(t, "", sut.NameMethodID([]byte{0x01, 0x02, 0x03, 0x04}))
	require.Equal(t, "sequenceBatchesElderberry", sut.NameMethodID(methodIDSequenceBatchesElderberry))
}
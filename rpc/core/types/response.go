package core_types

import (
	cmn "github.com/DSiSc/apigateway/common"
	apitypes "github.com/DSiSc/apigateway/core/types"
	"github.com/DSiSc/craft/types"
)

type ResultEcho struct {
	Value string `json:"value"`
}

type Blockdata struct {
	Number           cmn.Uint64           `json:"number"`
	Hash             cmn.Hash             `json:"hash"`
	ParentHash       cmn.Hash             `json:"parentHash"`
	MixHash          cmn.Hash             `json:"mixHash"`
	StateRoot        cmn.Hash             `json:"stateRoot"`
	Miner            apitypes.Address     `json:"miner"`
	Timestamp        cmn.Uint64           `json:"timestamp"`
	TransactionsRoot cmn.Hash             `json:"transactionsRoot"`
	ReceiptsRoot     cmn.Hash             `json:"receiptsRoot"`
	Transactions     []*types.Transaction `json:"transactions"`
}

type RPCTransaction struct {
	BlockHash        cmn.Hash       `json:"blockHash"`
	BlockNumber      cmn.Uint64     `json:"blockNumber"`
	From             *apitypes.Address `json:"from"`
	Gas              cmn.Uint64     `json:"gas"`
	GasPrice         *cmn.Big       `json:"gasPrice"`
	Hash             *cmn.Hash      `json:"hash"`
	Input            cmn.Bytes      `json:"input"`
	Nonce            *cmn.Uint64    `json:"nonce"`
	To               *apitypes.Address `json:"to"`
	TransactionIndex cmn.Uint       `json:"transactionIndex"`
	Value            *cmn.Big       `json:"value"`
	V                *cmn.Big       `json:"v"`
	R                *cmn.Big       `json:"r"`
	S                *cmn.Big       `json:"s"`
}

type RPCReceipt struct {
	// Consensus fields
	BlockHash         cmn.Hash       `json:"blockHash"`
	BlockNumber       *cmn.Big       `json:"blockNumber"`
	TransactionHash   *cmn.Hash      `json:"transactionHash"`
	TransactionIndex  cmn.Uint       `json:"transactionIndex"`
	From              *apitypes.Address `json:"from"`
	To                *apitypes.Address `json:"to"`
	Root              []byte         `json:"root"`
	Status            *cmn.Uint64    `json:"status"`
	GasUsed           *cmn.Uint64    `json:"gasUsed"`
	CumulativeGasUsed *cmn.Uint64    `json:"cumulativeGasUsed"`
	LogsBloom         *types.Bloom   `json:"logsBloom"`
	Logs              []*types.Log   `json:"logs"`
	ContractAddress   *apitypes.Address `json:"contractAddress"`
}

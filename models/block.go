package models

type Block struct {
	BlockID     string      `json:"blockID"`
	BlockHeader BlockHeader `json:"block_header"`
}

type BlockHeader struct {
	RawData          BlockRawData `json:"raw_data"`
	WitnessSignature string       `json:"witness_signature"`
}

type BlockRawData struct {
	Number         int64  `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int    `json:"version"`
	Timestamp      int64  `json:"timestamp"`
}

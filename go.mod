module github.com/novychok/niletrongrid

go 1.22.3

require (
	github.com/btcsuite/btcd v0.22.3
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/okx/go-wallet-sdk/coins/tron v0.0.0-20240412022709-61460eb99685
)

replace github.com/okx/go-wallet-sdk/coins/tron => github.com/TheTeaParty/go-wallet-sdk/coins/tron v0.0.0-20240612112222-e93ad228616e

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
)

require (
	github.com/TheTeaParty/trongrid v0.0.2
	github.com/btcsuite/btcd/btcec/v2 v2.3.3 // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/golang/protobuf v1.5.4
	github.com/okx/go-wallet-sdk/crypto v0.0.2 // indirect
	github.com/okx/go-wallet-sdk/util v0.0.2 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	google.golang.org/protobuf v1.34.2
)

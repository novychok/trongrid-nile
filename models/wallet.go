package models

import (
	client "github.com/TheTeaParty/trongrid"
)

type Wallet struct {
	FromWalletAddr   string
	FromPrivKey      string
	FromContractAddr string

	ToWalletAddr string
	ToPrivKey    string

	NileClient client.Client
}

func NewWallet(fromWalletAddr, fromPrivKey, fromContractAddr, toWalletAddr, toPrivKey string,
	nileClient client.Client) *Wallet {
	return &Wallet{
		FromWalletAddr:   fromWalletAddr,
		FromPrivKey:      fromPrivKey,
		FromContractAddr: fromContractAddr,

		ToWalletAddr: toWalletAddr,
		ToPrivKey:    toPrivKey,

		NileClient: nileClient,
	}
}

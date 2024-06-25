package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	client "github.com/TheTeaParty/trongrid"
	"github.com/btcsuite/btcd/btcec"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"

	"github.com/okx/go-wallet-sdk/coins/tron"
)

type wallet struct {
	fromWalletAddr string
	toWalletAddr   string
	fromPrivKey    string
	toPrivKey      string
	nileClient     client.Client
}

func (w *wallet) getAddress(walletPrivKey string) (string, string, error) {

	privateKeyBytes, err := hex.DecodeString(walletPrivKey)
	if err != nil {
		return "", "", fmt.Errorf("error to get private key bytes: %s", err)
	}

	_, publicKey := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)
	publicKeyHex := hex.EncodeToString(publicKey.SerializeUncompressed())

	address, err := tron.GetAddressByPublicKey(publicKeyHex)
	if err != nil {
		return "", "", fmt.Errorf("error to get address by public key: %s", err)
	}

	return address, publicKeyHex, nil
}

func (w *wallet) PrepareValidAddr() (string, string, error) {

	fromWalletAddr, _, err := w.getAddress(w.fromPrivKey)
	if err != nil {
		return "", "", err
	}

	toWalletAddr, _, err := w.getAddress(w.toPrivKey)
	if err != nil {
		return "", "", err
	}

	if !tron.ValidateAddress(fromWalletAddr) {
		return "", "", fmt.Errorf("not valid from wallet address: %s", fromWalletAddr)
	}

	if !tron.ValidateAddress(toWalletAddr) {
		return "", "", fmt.Errorf("not valid to wallet address: %s", toWalletAddr)
	}

	return fromWalletAddr, toWalletAddr, nil
}

func (w *wallet) ExecuteTransaction(nowBlock *client.Block, fromWalletAddr, toWalletAddr string, amount int64) (string, error) {

	nowBlockNumber := nowBlock.BlockHeader.RawData.Number
	nowBlockTimestamp := nowBlock.BlockHeader.RawData.Timestamp
	nowRefBlockHash := nowBlock.BlockID

	currentTime := time.Now()
	k1 := make([]byte, 8)
	binary.BigEndian.PutUint64(k1, uint64(nowBlockNumber))
	k2, _ := hex.DecodeString(nowRefBlockHash)

	txStr, err := tron.NewTransfer(
		fromWalletAddr,
		toWalletAddr,
		amount,
		hex.EncodeToString(k1[6:8]),
		hex.EncodeToString(k2[8:16]),
		currentTime.UnixMilli()+3600*1000,
		nowBlockTimestamp)
	if err != nil {
		return "", fmt.Errorf("error to make a transfer: %s", err)
	}

	return txStr, nil
}

func (w *wallet) SignTransaction(d2 string) (string, error) {

	signStart, err := tron.SignStart(d2)
	if err != nil {
		return "", fmt.Errorf("error to start sign the transaction: %s", err)
	}

	provKeyHex, _ := hex.DecodeString(w.fromPrivKey)
	sign, err := tron.Sign(signStart, secp256k1.PrivKeyFromBytes(provKeyHex))
	if err != nil {
		return "", fmt.Errorf("error to sign the transaction: %s", err)
	}

	signEnd, err := tron.SignEnd(d2, sign)
	if err != nil {
		return "", fmt.Errorf("error to sign end the transaction: %s", err)
	}

	return signEnd, nil
}

func NewWallet(fromWalletAddr, toWalletAddr, fromPrivKey, toPrivKey string,
	nileClient client.Client) *wallet {
	return &wallet{
		fromWalletAddr: fromWalletAddr,
		toWalletAddr:   toWalletAddr,
		fromPrivKey:    fromPrivKey,
		toPrivKey:      toPrivKey,
		nileClient:     nileClient,
	}
}

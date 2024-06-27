package service

import (
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/golang/protobuf/ptypes"
	"github.com/novychok/niletrongrid/models"
	"github.com/okx/go-wallet-sdk/coins/tron"
	"github.com/okx/go-wallet-sdk/coins/tron/pb"
	"google.golang.org/protobuf/proto"
)

func SignTransaction(txStr, fromPrivKey string) (string, error) {

	signStart, err := tron.SignStart(txStr)
	if err != nil {
		return "", err
	}

	privKeyHex, _ := hex.DecodeString(fromPrivKey)
	sign, err := tron.Sign(signStart, secp256k1.PrivKeyFromBytes(privKeyHex))
	if err != nil {
		return "", err
	}

	signEnd, err := tron.SignEnd(txStr, sign)
	if err != nil {
		return "", err
	}

	return signEnd, nil
}

func GetAddress(walletPrivKey string) (string, string, error) {

	privateKeyBytes, err := hex.DecodeString(walletPrivKey)
	if err != nil {
		return "", "", err
	}

	_, publicKey := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)
	publicKeyHex := hex.EncodeToString(publicKey.SerializeUncompressed())

	address, err := tron.GetAddressByPublicKey(publicKeyHex)
	if err != nil {
		return "", "", err
	}

	return address, publicKeyHex, nil
}

func PrepareValidAddr(fromPrivKey, toPrivKey string) (string, string, error) {

	fromWalletAddr, _, err := GetAddress(fromPrivKey)
	if err != nil {
		return "", "", err
	}

	toWalletAddr, _, err := GetAddress(toPrivKey)
	if err != nil {
		return "", "", err
	}

	if !tron.ValidateAddress(fromWalletAddr) {
		return "", "", err
	}

	if !tron.ValidateAddress(toWalletAddr) {
		return "", "", err
	}

	return fromWalletAddr, toWalletAddr, nil
}

func NewTransfer(nowBlock *models.Block, fromWalletAddr, toWalletAddr string, amount int64) (string, error) {

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
		return "", err
	}

	return txStr, nil
}

func NewDelegateResourceTransfer(transferContract *pb.DelegateResourceContract, delegateRequest models.Transaction) (string, error) {

	param, err := ptypes.MarshalAny(transferContract)
	if err != nil {
		return "", err
	}

	contract := &pb.Transaction_Contract{
		Type:      pb.Transaction_Contract_DelegateResourceContract,
		Parameter: param,
	}

	raw := new(pb.TransactionRaw)
	refBytes, err := hex.DecodeString(delegateRequest.RawData.RefBlockBytes)
	if err != nil {
		return "", err
	}

	raw.RefBlockBytes = refBytes
	refHash, err := hex.DecodeString(delegateRequest.RawData.RefBlockHash)
	if err != nil {
		return "", err
	}

	raw.RefBlockHash = refHash
	raw.Expiration = delegateRequest.RawData.Expiration
	raw.Timestamp = delegateRequest.RawData.Timestamp
	raw.Contract = []*pb.Transaction_Contract{contract}

	trans := pb.Transaction{
		RawData: raw,
	}

	data, err := proto.Marshal(&trans)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

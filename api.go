package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"

	client "github.com/TheTeaParty/trongrid"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/novychok/trontransaction/delegate"
	"github.com/okx/go-wallet-sdk/coins/tron"
	"github.com/okx/go-wallet-sdk/coins/tron/pb"
)

var nileBaseUrl = "https://nile.trongrid.io/"

var (
	getAccountResourceUrl = nileBaseUrl + "wallet/getaccountresource"
	delegateResourceUrl   = nileBaseUrl + "wallet/delegateresource"
	freezeBalanceV2Url    = nileBaseUrl + "wallet/freezebalancev2"
)

type api struct{}

func (w *wallet) DelegateResource() error {
	payload := strings.NewReader("{\"owner_address\":\"TLmzAddwWkmXXwNzXcKYsQ81h8cGLzeUtN\",\"receiver_address\":\"TNYpXvjX6j7E2NPK4Dwd1WXJjfUiRwcgQg\",\"balance\":1000000,\"resource\":\"ENERGY\",\"lock\":false,\"visible\":true}")

	req, err := http.NewRequest("POST", delegateResourceUrl, payload)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	fmt.Println(string(body))

	var delegateRequest delegate.Transaction
	if err := json.Unmarshal(body, &delegateRequest); err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}

	fromWalletAd, _ := tron.GetAddressHash(w.fromWalletAddr)
	toWalletAd, _ := tron.GetAddressHash(w.toWalletAddr)

	transferContract := &pb.DelegateResourceContract{
		OwnerAddress:    fromWalletAd,
		ReceiverAddress: toWalletAd,
		Resource:        1,
		Balance:         1000000,
		Lock:            false,
		LockPeriod:      10,
	}

	param, err := ptypes.MarshalAny(transferContract)
	if err != nil {
		return err
	}

	contract := &pb.Transaction_Contract{
		Type:      pb.Transaction_Contract_DelegateResourceContract,
		Parameter: param,
	}

	raw := new(pb.TransactionRaw)
	refBytes, err := hex.DecodeString(delegateRequest.RawData.RefBlockBytes)
	if err != nil {
		return err
	}

	raw.RefBlockBytes = refBytes
	refHash, err := hex.DecodeString(delegateRequest.RawData.RefBlockHash)
	if err != nil {
		return err
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
		return err
	}
	hexString := hex.EncodeToString(data)

	signStart, err := tron.SignStart(hexString)
	if err != nil {
		return err
	}

	provKeyHex, _ := hex.DecodeString(w.fromPrivKey)
	sign, err := tron.Sign(signStart, secp256k1.PrivKeyFromBytes(provKeyHex))
	if err != nil {
		return err
	}

	signEnd, err := tron.SignEnd(hexString, sign)
	if err != nil {
		return err
	}

	asd, err := w.nileClient.BroadcastHex(context.TODO(), &client.BroadcastHexRequest{
		Transaction: strings.ToUpper(signEnd),
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(asd)

	return nil
}

func (w *wallet) FreezeBalanceV2() error {

	payload := strings.NewReader("{\"owner_address\":\"TLmzAddwWkmXXwNzXcKYsQ81h8cGLzeUtN\",\"frozen_balance\":1000000,\"resource\":\"ENERGY\",\"visible\":true}")

	req, _ := http.NewRequest("POST", freezeBalanceV2Url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	var acc delegate.Transaction
	if err := json.NewDecoder(res.Body).Decode(&acc); err != nil {
		return fmt.Errorf("error to decode the response body %s", err)
	}

	amount := new(big.Int)
	amount.SetInt64(1000000)
	txStr, err := tron.NewTRC20TokenTransfer(w.fromWalletAddr, w.toWalletAddr, "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj",
		amount, 10000000, acc.RawData.RefBlockBytes, acc.RawData.RefBlockHash, acc.RawData.Expiration, acc.RawData.Timestamp)
	if err != nil {
		return err
	}

	signedTransaction, err := w.SignTransaction(txStr)
	if err != nil {
		log.Println(err)
	}

	response, err := w.nileClient.BroadcastHex(context.TODO(), &client.BroadcastHexRequest{
		Transaction: strings.ToUpper(signedTransaction),
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", response)

	return nil
}

// func (w *wallet) GetAccountResource() (*AccountResourceResponse, error) {

// 	payload := strings.NewReader("{\"owner_address\":\"TZ4UXDV5ZhNW7fb2AMSbgfAEZ7hWsnYS2g\",\"frozen_balance\":10000000,\"frozen_duration\":3,\"resource\":\"ENERGY\",\"visible\":true}")

// 	accountResourceBytes, err := json.Marshal(accountResource)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to marshal the request %s", err)
// 	}

// 	payload := strings.NewReader(string(accountResourceBytes))
// 	req, err := http.NewRequest("POST", getAccountResourceUrl, payload)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to prepare the request %s", err)
// 	}

// 	req.Header.Add("accept", "application/json")
// 	req.Header.Add("content-type", "application/json")

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to Do the request %s", err)
// 	}
// 	defer res.Body.Close()

// 	var accountResourceResponse AccountResourceResponse
// 	if err := json.NewDecoder(res.Body).Decode(&accountResourceResponse); err != nil {
// 		return nil, fmt.Errorf("error to decode the response body %s", err)
// 	}

// 	return &accountResourceResponse, nil
// }

func NewApi() *api {
	return &api{}
}

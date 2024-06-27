package impl

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/big"

	"github.com/novychok/niletrongrid/models"

	"github.com/novychok/niletrongrid/service"
	"github.com/okx/go-wallet-sdk/coins/tron"
	"github.com/okx/go-wallet-sdk/coins/tron/pb"
)

var (
	freezeBalanceV2Url   = "wallet/freezebalancev2"
	broadcastHexUrl      = "wallet/broadcasthex"
	delegateResourceUrl  = "wallet/delegateresource"
	nowBlockUrl          = "walletsolidity/getnowblock"
	createTransactionUrl = "wallet/createtransaction"
)

type srv struct {
	l *slog.Logger

	wallet *models.Wallet
}

func (s *srv) CreateTransaction(ctx context.Context, payload *models.CreateTransactionPayload) error {

	l := s.l.With(slog.String("method", "CreateTransaction"))

	nowBlockResponse, err := service.ApiCall(ctx, nowBlockUrl, "", service.METHOD_GET)
	if err != nil {
		l.ErrorContext(ctx, "failed to get now block from API call", "err", err)
		return err
	}

	var nowBlock models.Block
	if err := json.Unmarshal(nowBlockResponse, &nowBlock); err != nil {
		l.ErrorContext(ctx, "failed to unmarshal api response", "err", err)
		return err
	}

	txStr, err := service.NewTransfer(&nowBlock, s.wallet.FromWalletAddr, s.wallet.ToWalletAddr, payload.Amount)
	if err != nil {
		l.ErrorContext(ctx, "failed to make a new transfer", "err", err)
		return err
	}

	signedTx, err := service.SignTransaction(txStr, s.wallet.FromPrivKey)
	if err != nil {
		l.ErrorContext(ctx, "failed to sign the transaction", "err", err)
		return err
	}

	createTransactionResponse, err := service.BroadcastHex(ctx, signedTx, broadcastHexUrl, service.METHOD_POST)
	if err != nil {
		l.ErrorContext(ctx, "failed to get broadcast hex response from API call", "err", err)
		return err
	}

	l.Info("broadcast hex response from create transaction resource", "CreateTransaction", string(createTransactionResponse))
	return nil
}

func (s *srv) DelegateResource(ctx context.Context, payload *models.DelegateResourcePayload) error {

	l := s.l.With(slog.String("method", "DelegateResource"))

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		l.ErrorContext(ctx, "failed to marshal the payload", "err", err)
		return err
	}

	delegateResourceTrsResponse, err := service.ApiCall(ctx, delegateResourceUrl, string(jsonPayload), service.METHOD_POST)
	if err != nil {
		l.ErrorContext(ctx, "failed to get delegate resource response from API call", "err", err)
		return err
	}

	var delegateResourceTrs models.Transaction
	if err := json.Unmarshal(delegateResourceTrsResponse, &delegateResourceTrs); err != nil {
		l.ErrorContext(ctx, "failed to unmarshal api response", "err", err)
		return err
	}

	fromWalletAddr, err := tron.GetAddressHash(s.wallet.FromWalletAddr)
	if err != nil {
		l.ErrorContext(ctx, "failed to get from wallet address hash", "err", err)
		return err
	}
	toWalletAddr, err := tron.GetAddressHash(s.wallet.ToWalletAddr)
	if err != nil {
		l.ErrorContext(ctx, "failed to get to wallet address hash", "err", err)
		return err
	}

	transferContract := &pb.DelegateResourceContract{
		OwnerAddress:    fromWalletAddr,
		ReceiverAddress: toWalletAddr,
		Resource:        1, // TODO: hardcoded resource
		Balance:         delegateResourceTrs.RawData.Contract[0].Parameter.Value.Balance,
		Lock:            false,
	}

	txStr, err := service.NewDelegateResourceTransfer(transferContract, delegateResourceTrs)
	if err != nil {
		l.ErrorContext(ctx, "failed to execute delegate resource transfer", "err", err)
		return err
	}

	signedTx, err := service.SignTransaction(txStr, s.wallet.FromPrivKey)
	if err != nil {
		l.ErrorContext(ctx, "failed to sign the transaction", "err", err)
		return err
	}

	delegateResourceResponse, err := service.BroadcastHex(ctx, signedTx, broadcastHexUrl, service.METHOD_POST)
	if err != nil {
		l.ErrorContext(ctx, "failed to get broadcast hex response from API call", "err", err)
		return err
	}

	l.Info("broadcast hex response from delegate resource", "DelegateResource", string(delegateResourceResponse))
	return nil

}

func (s *srv) FreezeBalanceV2(ctx context.Context, payload *models.FreezeBalanceV2Payload) error {

	l := s.l.With(slog.String("method", "FreezeBalanceV2"))

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		l.ErrorContext(ctx, "failed to marshal the payload", "err", err)
		return err
	}

	freezeTrsResponse, err := service.ApiCall(ctx, freezeBalanceV2Url, string(jsonPayload), service.METHOD_POST)
	if err != nil {
		l.ErrorContext(ctx, "failed to get freeze transaction response from API call", "err", err)
		return err
	}

	var freezeTransaction models.Transaction
	if err := json.Unmarshal(freezeTrsResponse, &freezeTransaction); err != nil {
		l.ErrorContext(ctx, "failed to unmarshal api response", "err", err)
		return err
	}

	amount := new(big.Int)
	amount.SetUint64(uint64(payload.FrozenBalance))

	txStr, err := tron.NewTRC20TokenTransfer(s.wallet.FromWalletAddr,
		s.wallet.ToWalletAddr,
		s.wallet.FromContractAddr,
		amount,
		10000000, // TODO: not hardcoded need
		freezeTransaction.RawData.RefBlockBytes,
		freezeTransaction.RawData.RefBlockHash,
		freezeTransaction.RawData.Expiration,
		freezeTransaction.RawData.Timestamp,
	)
	if err != nil {
		l.ErrorContext(ctx, "failed to create new TRC20 token transfer", "err", err)
		return err
	}

	signedTx, err := service.SignTransaction(txStr, s.wallet.FromPrivKey)
	if err != nil {
		l.ErrorContext(ctx, "failed to sign the transaction", "err", err)
		return err
	}

	delegateResourceResponse, err := service.BroadcastHex(ctx, signedTx, broadcastHexUrl, service.METHOD_POST)
	if err != nil {
		l.ErrorContext(ctx, "failed to get broadcast hex response from API call", "err", err)
		return err
	}

	l.Info("broadcast hex response from freeze balance", "FreezeBalanceV2", string(delegateResourceResponse))
	return nil
}

func NewAccountResourceSrv(l *slog.Logger, wallet *models.Wallet) service.AccountResource {
	return &srv{
		l:      l,
		wallet: wallet,
	}
}

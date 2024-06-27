package main

import (
	"context"

	client "github.com/TheTeaParty/trongrid"
	"github.com/novychok/niletrongrid/models"
	"github.com/novychok/niletrongrid/service/impl"
	"github.com/novychok/niletrongrid/util/log"
)

var (
	fromWalletAddr   = "TLmzAddwWkmXXwNzXcKYsQ81h8cGLzeUtN"
	fromPrivKey      = "776c2c3fd798cb556f312340cf033c872e9eab6a753d29911a15987a8a448f5d"
	fromContractAddr = "TXLAQ63Xg1NAzckPwKHvzw7CSEmLMEqcdj"

	toWalletAddr = "TNYpXvjX6j7E2NPK4Dwd1WXJjfUiRwcgQg"
	toPrivKey    = "3b157e80fd9a943cfdf99cc0afe17c511a1f4e3485834f94bd5fc310bcf1d152"
)

func main() {

	clientOpts := client.WithNetwork(client.NetworkNileTestnet)
	nileClient := client.New(clientOpts)

	l := log.New()

	wallet := models.NewWallet(fromWalletAddr,
		fromPrivKey,
		fromContractAddr,
		toWalletAddr,
		toPrivKey,
		nileClient)

	accountResourceSrv := impl.NewAccountResourceSrv(l, wallet)

	ctx := context.Background()

	// freezeBalanceV2Payload := &models.FreezeBalanceV2Payload{
	// 	OwnerAddress:  wallet.FromWalletAddr,
	// 	FrozenBalance: 1000000,
	// 	Resource:      "ENERGY",
	// 	Visible:       true,
	// }
	// if err := accountResourceSrv.FreezeBalanceV2(ctx, freezeBalanceV2Payload); err != nil {
	// 	return
	// }

	// delegateResourcePayload := &models.DelegateResourcePayload{
	// 	OwnerAddress:    wallet.FromWalletAddr,
	// 	ReceiverAddress: wallet.ToWalletAddr,
	// 	Balance:         1000000,
	// 	Resource:        "ENERGY",
	// 	Lock:            false,
	// 	Visible:         true,
	// }
	// if err := accountResourceSrv.DelegateResource(ctx, delegateResourcePayload); err != nil {
	// 	return
	// }

	createTransactionPayload := &models.CreateTransactionPayload{
		Amount:  1000000,
		Visible: true,
	}
	if err := accountResourceSrv.CreateTransaction(ctx, createTransactionPayload); err != nil {
		return
	}

}

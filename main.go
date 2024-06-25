package main

import (
	"fmt"

	client "github.com/TheTeaParty/trongrid"
)

var (
	// fromWalletAddr = "TJPdb3QNSyRPHQJsYFXup5n6DsKBnE1dXc"
	// toWalletAddr   = "TMLuLPS3nbkqLaP53QEAG6zzUuBZxnxcQD"
	// fromPrivKey    = "5176a02c53a3b64aeb119255bf3420188fa62d57ebd0ba98a53262ff644695b2"
	// toPrivKey      = "52eeeff66ae21c5a27107554ce8b1dcb0bd78cdcdb030d907ce4ed5f6be314d8"

	fromWalletAddr = "TLmzAddwWkmXXwNzXcKYsQ81h8cGLzeUtN"
	toWalletAddr   = "TNYpXvjX6j7E2NPK4Dwd1WXJjfUiRwcgQg"
	fromPrivKey    = "776c2c3fd798cb556f312340cf033c872e9eab6a753d29911a15987a8a448f5d"
	toPrivKey      = "3b157e80fd9a943cfdf99cc0afe17c511a1f4e3485834f94bd5fc310bcf1d152"
)

func main() {

	clientOpts := client.WithNetwork(client.NetworkNileTestnet)
	nileClient := client.New(clientOpts)

	wallet := NewWallet(fromWalletAddr, toWalletAddr, fromPrivKey, toPrivKey, nileClient)

	// nowBlock, err := nileClient.GetNowBlock(context.TODO())
	// if err != nil {
	// 	log.Println(err)
	// }

	// fromWalletAddr, toWalletAddr, err := wallet.PrepareValidAddr()
	// if err != nil {
	// 	log.Println(err)
	// }

	// var amount int64 = 100000000
	// txStr, err := wallet.ExecuteTransaction(nowBlock, fromWalletAddr, toWalletAddr, amount)
	// if err != nil {
	// 	log.Println(err)
	// }

	// signedTransaction, err := wallet.SignTransaction(txStr)
	// if err != nil {
	// 	log.Println(err)
	// }

	// _, err = nileClient.BroadcastHex(context.TODO(), &client.BroadcastHexRequest{
	// 	Transaction: strings.ToUpper(signedTransaction),
	// })
	// if err != nil {
	// 	log.Println(err)
	// }

	// _, err := wallet.GetAccountResource()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("------------------------------------------")
	// fmt.Printf("%+v\n", accountResource)
	// fmt.Println("------------------------------------------")

	// fmt.Println("------------------------------------------")
	// fmt.Println(resp)
	// fmt.Println("------------------------------------------")

	// fmt.Println("------------------------------------------")
	// wallet.FreezeBalanceV2()
	// fmt.Println("------------------------------------------")

	fmt.Println("------------------------------------------")
	wallet.DelegateResource()
	fmt.Println("------------------------------------------")
}

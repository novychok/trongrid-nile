package models

type FreezeBalanceV2Payload struct {
	OwnerAddress  string `json:"owner_address"`
	FrozenBalance int    `json:"frozen_balance"`
	Resource      string `json:"resource"`
	Visible       bool   `json:"visible"`
}

type DelegateResourcePayload struct {
	OwnerAddress    string `json:"owner_address"`
	ReceiverAddress string `json:"receiver_address"`
	Balance         int    `json:"balance"`
	Resource        string `json:"resource"`
	Lock            bool   `json:"lock"`
	Visible         bool   `json:"visible"`
}

type CreateTransactionPayload struct {
	Amount  int64 `json:"amount"`
	Visible bool  `json:"visible"`
}

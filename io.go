package main

type AccountResource struct {
	AccountAddress string `json:"address"`
	Visible        bool   `json:"visible"`
}

type AccountResourceResponse struct {
	FreeNetUsed       int64 `json:"freeNetUsed"`
	FreeNetLimit      int64 `json:"freeNetLimit"`
	NetUsed           int64 `json:"NetUsed"`
	NetLimit          int64 `json:"NetLimit"`
	TotalNetLimit     int64 `json:"TotalNetLimit"`
	TotalNetWeight    int64 `json:"TotalNetWeight"`
	TronPowerLimit    int64 `json:"tronPowerLimit"`
	EnergyLimit       int64 `json:"EnergyLimit"`
	TotalEnergyLimit  int64 `json:"TotalEnergyLimit"`
	TotalEnergyWeight int64 `json:"TotalEnergyWeight"`
}

type FreezeBalanceV2 struct {
}

type FreezeBalanceV2Response struct {
}

type DelegateResource struct {
}

type DelegateResourceResponse struct {
}

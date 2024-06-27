package service

import (
	"context"

	"github.com/novychok/niletrongrid/models"
)

type AccountResource interface {
	CreateTransaction(ctx context.Context, payload *models.CreateTransactionPayload) error
	FreezeBalanceV2(ctx context.Context, payload *models.FreezeBalanceV2Payload) error
	DelegateResource(ctx context.Context, payload *models.DelegateResourcePayload) error
}

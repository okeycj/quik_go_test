package dto

import "github.com/shopspring/decimal"

type CreateWalletDTO struct {
	ID     uint64          `json:"id" form:"id" binding:"required"`
	Amount decimal.Decimal `json:"amount" form:"amount" binding:"required"`
	UserID uint64          `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type UpdateWalletDTO struct {
	Amount decimal.Decimal `json:"amount" form:"amount" binding:"required"`
	UserID uint64          `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type CreditOrDebitWalletDTO struct {
	Amount decimal.Decimal `json:"amount" form:"amount" binding:"required"`
}

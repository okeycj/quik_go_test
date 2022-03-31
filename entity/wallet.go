package entity

import "github.com/shopspring/decimal"

type Wallet struct {
	ID     uint64          `gorm:"primary_key:auto_increment" json:"id"`
	Amount decimal.Decimal `gorm:"type:decimal(10,2)" json:"amount"`
	UserID uint64          `gorm:"not null" json:"-"`
	User   User            `gorm:"foreignKey:UserID;constaint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}

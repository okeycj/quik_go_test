package repository

import (
	"errors"

	"github.com/okeycj/quik_go_test/entity"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(wallet entity.Wallet) entity.Wallet
	CreditWallet(walletID string, amount decimal.Decimal) interface{}
	DebitWallet(walletID string, amount decimal.Decimal) interface{}
	FindByID(walletID string) (tx *gorm.DB)
	GetBalance(wallet string) entity.Wallet
}

type walletConnection struct {
	connection *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletConnection{
		connection: db,
	}
}

func (db *walletConnection) CreateWallet(wallet entity.Wallet) entity.Wallet {
	db.connection.Save(&wallet)
	return wallet
}

func (db *walletConnection) CreditWallet(walletID string, amount decimal.Decimal) interface{} {
	var wallet entity.Wallet

	if amount.IsNegative() {
		return nil
	}

	res := db.connection.Model(&wallet).Preload("User").First(&wallet, walletID)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if res.Error == nil {
			wallet.Amount = amount.Add(wallet.Amount)
			db.connection.Save(&wallet)
			return wallet
		}
		return nil
	}
	return nil
}

func (db *walletConnection) DebitWallet(walletID string, amount decimal.Decimal) interface{} {
	var wallet entity.Wallet

	if amount.IsNegative() {
		return nil
	}
	res := db.connection.Model(&wallet).Preload("User").First(&wallet, walletID)
	if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if res.Error == nil {
			if wallet.Amount.IsZero() {
				return nil
			}
			wallet.Amount = wallet.Amount.Sub(amount)
			db.connection.Save(&wallet)
			return wallet
		}
		return nil
	}
	return nil
}

func (db *walletConnection) FindByID(walletID string) (tx *gorm.DB) {
	var wallet entity.Wallet
	return db.connection.First(&wallet, walletID)
}

func (db *walletConnection) GetBalance(walletID string) entity.Wallet {
	var wallet entity.Wallet
	db.connection.Model(&wallet).Preload("User").First(&wallet, walletID)
	return wallet
}

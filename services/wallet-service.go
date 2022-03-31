package services

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"github.com/okeycj/quik_go_test/dto"
	"github.com/okeycj/quik_go_test/entity"
	"github.com/okeycj/quik_go_test/repository"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletService interface {
	CreateWallet(wallet dto.CreateWalletDTO) entity.Wallet
	CreditWallet(walletID string, amount decimal.Decimal) interface{}
	DebitWallet(walletID string, amount decimal.Decimal) interface{}
	GetWalletBalance(walletID string) entity.Wallet
	WalletAssists(walletID string) bool
}

type walletService struct {
	walletRepository repository.WalletRepository
}

func NewWalletService(walletRepo repository.WalletRepository) WalletService {
	return &walletService{
		walletRepository: walletRepo,
	}
}

func (service *walletService) CreateWallet(wallet dto.CreateWalletDTO) entity.Wallet {
	walletToCreate := entity.Wallet{}
	err := smapping.FillStruct(&walletToCreate, smapping.MapFields(&wallet))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.walletRepository.CreateWallet(walletToCreate)
	return res
}

func (service *walletService) CreditWallet(walletID string, amount decimal.Decimal) interface{} {
	res := service.walletRepository.CreditWallet(walletID, amount)
	if _, ok := res.(entity.Wallet); ok {
		return res
	}
	return nil
}

func (service *walletService) DebitWallet(walletID string, amount decimal.Decimal) interface{} {
	res := service.walletRepository.DebitWallet(walletID, amount)
	if _, ok := res.(entity.Wallet); ok {
		return res
	}
	return nil
}

func (service *walletService) GetWalletBalance(walletID string) entity.Wallet {
	return service.walletRepository.GetBalance(walletID)
}

func (service *walletService) WalletAssists(walletID string) bool {
	res := service.walletRepository.FindByID(walletID)
	return !(errors.Is(res.Error, gorm.ErrRecordNotFound))
}

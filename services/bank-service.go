package services

import (
	"fmt"

	"github.com/ElladanTasartir/go-bank-api/models"
	"gorm.io/gorm"
)

type BankService struct {
	database *gorm.DB
}

func NewBankService(connection *gorm.DB) *BankService {
	bankService := &BankService{
		database: connection,
	}

	return bankService
}

func (b *BankService) CreateBank(bank *models.Bank) error {
	result := b.database.Create(bank)
	if result.Error != nil {
		fmt.Printf("Error %v was thrown", result.Error)
		return result.Error
	}

	return nil
}

func (b *BankService) GetBank(bank *models.Bank, id string) error {
	result := b.database.Where("id = ?", id).First(bank)
	if result.Error != nil {
		fmt.Printf("Error %v was thrown", result.Error)
		return result.Error
	}

	return nil
}

func (b *BankService) GetBanks(banks *[]models.Bank) error {
	result := b.database.Find(banks)
	if result.Error != nil {
		fmt.Printf("Error %v was thrown", result.Error)
		return result.Error
	}

	return nil
}

func (b *BankService) DeleteBanks(bank *models.Bank, id string) error {
	result := b.database.Where("id = ?", id).Delete(bank)
	if result.Error != nil {
		fmt.Printf("Error %v was thrown", result.Error)
		return result.Error
	}

	return nil
}

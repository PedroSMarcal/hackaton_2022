package models

import (
	"time"

	"gorm.io/gorm"
)

type ConsilicaoBancaria struct {
	ID            uint
	Data          time.Time
	OperationTipe string
	Entry         bool
	Category      string
	Description   string
	Document      string
	Value         string
	TotalAmmount  int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// type Transaction struct {
// 	Total      int `json:"total"`
// 	TotalPages int `json:"totalPages"`
// 	Page       int `json:"page"`
// 	Results    []struct {
// 		ID                 string      `json:"id"`
// 		Description        string      `json:"description"`
// 		DescriptionRaw     string      `json:"descriptionRaw"`
// 		CurrencyCode       string      `json:"currencyCode"`
// 		Amount             int         `json:"amount"`
// 		Date               time.Time   `json:"date"`
// 		Category           interface{} `json:"category"`
// 		Balance            float64     `json:"balance"`
// 		AccountID          string      `json:"accountId"`
// 		ProviderCode       interface{} `json:"providerCode"`
// 		Status             string      `json:"status"`
// 		PaymentData        interface{} `json:"paymentData"`
// 		Type               string      `json:"type"`
// 		CreditCardMetadata interface{} `json:"creditCardMetadata"`
// 		Merchant           interface{} `json:"merchant"`
// 	} `json:"results"`
// }

// type (
// 	Payer struct {
// 		Name           string `json:"name"`
// 		BranchNumber   string `json:"branchNumber"`
// 		AccountNumber  string `json:"accountNumber"`
// 		RoutingNumber  string `json:"routingNumber"`
// 		DocumentNumber struct {
// 			Type  string `json:"type"`
// 			Value string `json:"value"`
// 		} `json:"documentNumber"`
// 	}
// )

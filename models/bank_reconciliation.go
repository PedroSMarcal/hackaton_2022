package models

import (
	"time"
)

type BankReconciliation struct {
	Date          time.Time
	OperationTipe string
	Entry         bool
	Category      string
	Description   string
	CpfCnpj       string
	Value         float64
	Documento     string
	TotalAmmount  float64
}

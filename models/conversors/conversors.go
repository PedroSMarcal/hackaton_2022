package conversors

import (
	"github.com/PedroSMarcal/hackaton2022/models"
	"github.com/PedroSMarcal/hackaton2022/models/responses"
)

func checkPositiveOrNegative(value float64) bool {
	return value > 0
}

func TransactionToBankAccount(transaction []responses.TransactionResults) []models.BankReconciliation {
	bankReconciliation := []models.BankReconciliation{}
	convert := models.BankReconciliation{}

	for _, value := range transaction {
		convert = models.BankReconciliation{
			Date:          value.Date,
			OperationTipe: value.PaymentData.PaymentMethod,
			Entry:         value.Type,
			Category:      value.Category,
			Description:   value.PaymentData.Payer.Name,
			CpfCnpj:       value.PaymentData.Payer.DocumentNumber.Type,
			Value:         value.Amount,
			Documento:     value.PaymentData.Payer.DocumentNumber.Value,
			TotalAmmount:  value.Balance,
		}

		bankReconciliation = append(bankReconciliation, convert)
	}

	return bankReconciliation
}

func TransactionToCashFlow(transaction []responses.TransactionResults) []models.CashFlow {
	cashFlow := []models.CashFlow{}
	convert := models.CashFlow{}

	for _, value := range transaction {
		if decide := checkPositiveOrNegative(value.Amount); decide {
			convert = models.CashFlow{
				Data:    value.Date,
				Account: value.Category,
				Amount:  value.Balance,
				Entry:   value.Amount,
				Out:     0,
			}

			cashFlow = append(cashFlow, convert)
			continue
		}

		convert = models.CashFlow{
			Data:    value.Date,
			Account: value.Category,
			Amount:  value.Balance,
			Entry:   0,
			Out:     value.Amount,
		}

		cashFlow = append(cashFlow, convert)
	}

	return cashFlow
}

package conversors

import (
	"github.com/PedroSMarcal/hackaton2022/models"
	"github.com/PedroSMarcal/hackaton2022/models/responses"
)

func checkType(value string) bool {
	if "POSTED" == value {
		return true
	}
	return false
}

func TransactionToBankAccount(transaction []responses.TransactionResults) []models.BankReconciliation {
	bankReconciliation := []models.BankReconciliation{}
	convert := models.BankReconciliation{}

	for _, value := range transaction {
		convert = models.BankReconciliation{
			Date:          value.Date,
			OperationTipe: value.Description,
			Entry:         checkType(value.Status),
			Category:      "",
			Description:   value.DescriptionRaw,
			CpfCnpj:       value.PaymentData.Payer.DocumentNumber.Type,
			Value:         value.Amount,
			Documento:     value.PaymentData.Payer.DocumentNumber.Value,
			TotalAmmount:  value.Balance,
		}

		bankReconciliation = append(bankReconciliation, convert)
	}

	return bankReconciliation
}

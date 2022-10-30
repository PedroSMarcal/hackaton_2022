package service

import (
	"github.com/PedroSMarcal/hackaton2022/adapters"
	"github.com/PedroSMarcal/hackaton2022/models"
	"github.com/PedroSMarcal/hackaton2022/models/conversors"
)

func ConvertTransactionsToBankRelation() ([]models.BankReconciliation, error) {
	transaction, err := adapters.DoRequestToPluggy()
	if err != nil {
		return nil, err
	}

	response := conversors.TransactionToBankAccount(transaction.Results)

	return response, nil
}

package service

import (
	"sort"

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

	sort.Slice(response, func(i, j int) bool {
		return response[i].Date.Before(response[j].Date)
	})

	return response, nil
}

func ConvertTransactionToCashFlowReturn() ([]models.CashFlow, error) {
	transaction, err := adapters.DoRequestToPluggy()
	if err != nil {
		return nil, err
	}

	response := conversors.TransactionToCashFlow(transaction.Results)

	sort.Slice(response, func(i, j int) bool {
		return response[i].Data.Before(response[j].Data)
	})

	return response, nil
}

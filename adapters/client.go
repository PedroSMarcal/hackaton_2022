package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PedroSMarcal/hackaton2022/common"
	"github.com/PedroSMarcal/hackaton2022/configs"
	"github.com/PedroSMarcal/hackaton2022/models/responses"
)

func Auth() error {

	body := GetClient{
		ClientID:     configs.EnvVariable.ClientID,
		ClientSecret: configs.EnvVariable.ClientSecret,
	}

	bodybytes, _ := json.Marshal(&body)

	bodyreq := io.NopCloser(bytes.NewBuffer(bodybytes))

	url := "https://api.pluggy.ai/auth"

	req, err := http.NewRequest(http.MethodPost, url, bodyreq)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	bodyres, _ := io.ReadAll(res.Body)

	api := responses.APIKey{}

	err = json.Unmarshal(bodyres, &api)
	if err != nil {
		return err
	}

	common.ApiKey = api.APIKey
	return nil
}

func GetConnectorID() error {

	url := "https://api.pluggy.ai/connectors?countries=BR"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", common.ApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	response := responses.ConnectorResponse{}
	json.Unmarshal(body, &response)

	common.ConnectorKey = response.ID

	return nil
}

func GetItens() error {

	url := "https://api.pluggy.ai/items"

	payload := strings.NewReader("{\"parameters\":{\"user\":\"user-ok\",\"password\":\"password-ok\"},\"connectorId\":8}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-API-KEY", common.ApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	response := responses.ItemsResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Print(err)
		return err
	}

	common.ItemKey = response.ID

	return nil
}

func GetAccount() error {
	var url string = "https://api.pluggy.ai/accounts?itemId=5129027e-57d6-4c6c-9aa3-0cd4ad6efac8"
	// var url string = fmt.Sprintf("https://api.pluggy.ai/accounts?itemId=%s", common.ItemKey)

	// fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", common.ApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	response := responses.AccountResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Print(err)
		return err
	}

	for _, value := range response.Results {
		if value.Name == "Conta corrente" {
			common.AccountKey = value.ID
		}
	}

	return nil
}

func GetTransaction(transaction *responses.TransactionResponse) error {

	url := fmt.Sprintf("https://api.pluggy.ai/transactions?accountId=%s", common.AccountKey)

	fmt.Println(&url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", common.ApiKey)

	fmt.Println(common.ApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(body, &transaction)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func DoRequestToPluggy() (responses.TransactionResponse, error) {
	transaction := responses.TransactionResponse{}
	Auth()
	GetItens()
	GetAccount()
	if err := GetTransaction(&transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}

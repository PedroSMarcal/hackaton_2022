package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PedroSMarcal/hackaton2022/common"
	"github.com/PedroSMarcal/hackaton2022/configs"
)

type Parameters struct {
	User     string
	Password string
}

type GetClient struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type BodyConnecorId struct {
	ConnectorID string
	Parameters  Parameters
}

// OK
func GetAPPEnv() {

	body := GetClient{
		ClientID:     configs.EnvVariable.ClientID,
		ClientSecret: configs.EnvVariable.ClientSecret,
	}

	bodybytes, _ := json.Marshal(&body)

	bodyreq := io.NopCloser(bytes.NewBuffer(bodybytes))

	url := "https://api.pluggy.ai/auth"

	req, _ := http.NewRequest(http.MethodPost, url, bodyreq)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	bodyres, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(bodyres))
}

func CreateConnectionWithBank() {

	url := "https://api.pluggy.ai/items"

	parameters := BodyConnecorId{
		ConnectorID: common.BankAccount,
		Parameters: Parameters{
			User:     common.User,
			Password: common.Password,
		},
	}

	bodyByte, _ := json.Marshal(parameters)

	bodyReq := io.NopCloser(bytes.NewBuffer(bodyByte))

	req, _ := http.NewRequest("POST", url, bodyReq)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

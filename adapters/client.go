package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PedroSMarcal/hackaton2022/common"
	"github.com/PedroSMarcal/hackaton2022/configs"
)

type Parameters struct {
	User     string
	Password string
}

type BodyConnecorId struct {
	ConnectorID string
	Parameters  Parameters
}

func GetAPPEnv() {

	reader := fmt.Sprintf(`{clientId:%s,clientSecret:%s}`, configs.EnvVariable.ClientID, configs.EnvVariable.ClientSecret)

	url := "https://api.pluggy.ai/auth"

	payload := strings.NewReader(reader)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
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

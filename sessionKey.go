package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// {
// 	"output_ResponseCode": "INS-0",
// 	"output_ResponseDesc": "Request processed successfully",
// 	"output_SessionID": "e14b7829a4ee44b4826a9ca12b773e92"
//   }

type ResponseData struct {
	Output_ResponseCode string `json:"output_ResponseCode"`
	Output_ResponseDesc string `json:"output_ResponseDesc"`
	Output_SessionID    string `json:"output_SessionID"`
}

func getSessionKey() (string, error) {
	url := "https://openapi.m-pesa.com:443/sandbox/ipg/v2/vodacomLES/getSession/"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", encrypted))
	req.Header.Set("Origin", "*")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("error making request: %v", err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading Response body")
	}

	var data ResponseData

	error := json.Unmarshal(body, &data)
	if error != nil {
		fmt.Println("Error unmarshaling data")
		return "", error
	}

	SessionID := data.Output_SessionID

	return SessionID, nil
}

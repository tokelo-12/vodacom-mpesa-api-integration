package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func c2bPay(c *gin.Context) {

	// The API endpoint
	apiURL := "https://openapi.m-pesa.com/sandbox/ipg/v2/vodacomLES/c2bPayment/singleStage/"
	key, err := getSessionKey()

	if err != nil {
		fmt.Printf("Failed to get Session-Key: %v", err)
	}

	apiKey := ecrypt(base64PublicKeyString(), key)

	// Prepare JSON payload
	payload := map[string]string{
		"input_Amount":                   "10",
		"input_Country":                  "LES",
		"input_Currency":                 "LSL",
		"input_CustomerMSISDN":           "000000000001",
		"input_ServiceProviderCode":      "000000",
		"input_ThirdPartyConversationID": "asv02e5958774f7ba228d83d0d689761",
		"input_TransactionReference":     "T1234C",
		"input_PurchasedItemsDesc":       "Shoes",
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// custom HTTP client with a timeout and TLS configuration
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Use secure settings in production
		},
	}

	// HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("public_key", base64PublicKeyString())
	req.Header.Set("Origin", "127.0.0.1")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print response details
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)

	//Unmarshal JSON to object pointer

	var data Response

	error := json.Unmarshal(body, &data)

	if error != nil {
		fmt.Printf("Error unmarshaling JSON: %v", error)
	}

	// Marshal the struct to json

	c.JSON(http.StatusOK, data)
}

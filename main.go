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
)

var app_key = "kxHXi2jfcSJ0nLJtE4e095AmjZTwXXar"
var base64PublicKeyString = "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArv9yxA69XQKBo24BaF/D+fvlqmGdYjqLQ5WtNBb5tquqGvAvG3WMFETVUSow/LizQalxj2ElMVrUmzu5mGGkxK08bWEXF7a1DEvtVJs6nppIlFJc2SnrU14AOrIrB28ogm58JjAl5BOQawOXD5dfSk7MaAA82pVHoIqEu0FxA8BOKU+RGTihRU+ptw1j4bsAJYiPbSX6i71gfPvwHPYamM0bfI4CmlsUUR3KvCG24rB6FNPcRBhM3jDuv8ae2kC33w9hEq8qNB55uw51vK7hyXoAa+U7IqP1y6nBdlN25gkxEA8yrsl1678cspeXr+3ciRyqoRgj9RD/ONbJhhxFvt1cLBh+qwK2eqISfBb06eRnNeC71oBokDm3zyCnkOtMDGl7IvnMfZfEPFCfg5QgJVk1msPpRvQxmEsrX9MQRyFVzgy2CWNIb7c+jPapyrNwoUbANlN8adU1m6yOuoX7F49x+OjiG2se0EJ6nafeKUXw/+hiJZvELUYgzKUtMAZVTNZfT8jjb58j8GVtuS+6TM2AutbejaCV84ZK58E2CRJqhmjQibEUO6KPdD7oTlEkFy52Y1uOOBXgYpqMzufNPmfdqqqSM4dU70PO8ogyKGiLAIxCetMjjm6FCMEA3Kc8K0Ig7/XtFm9By6VxTJK1Mg36TlHaZKP6VzVLXMtesJECAwEAAQ=="

var encrypted string

// func main() {
// 	// 1. Decoding the Base64-encoded Public Key
// 	decodedKey, err := base64.StdEncoding.DecodeString(base64PublicKeyString)
// 	if err != nil {
// 		fmt.Print("Failed to decode base64 string: %v", err)
// 	}

// 	var pub interface{}
// 	pub, _ = x509.ParsePKIXPublicKey(decodedKey)
// 	publicKey := pub.(*rsa.PublicKey)

// 	// 2. Creating an RSA Cipher Instance
// 	rsaCipher, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(app_key), nil)

// 	if err != nil {
// 		return
// 	}

// 	// 3. Encoding the API Key
// 	encryptedAPIKey := base64.StdEncoding.EncodeToString(rsaCipher)
// 	fmt.Println("Encrypted API Key:", encryptedAPIKey)
// }

func main() {
	ecrypt(base64PublicKeyString, app_key)

	c2bPay()

}

func c2bPay() {
	// Set up the API endpoint and public key (ensure these are accurate for the API context)
	apiURL := "https://openapi.m-pesa.com/sandbox/ipg/v2/vodacomLES/c2bPayment/singleStage/"
	key, err := getSessionKey()

	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	apiKey := ecrypt(base64PublicKeyString, key)

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

	// Set up a custom HTTP client with a timeout and TLS configuration
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Use secure settings in production
		},
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("public_key", base64PublicKeyString)
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
	// fmt.Println("Headers:", resp.Header)
	fmt.Printf("Body: %s\n", body)
}

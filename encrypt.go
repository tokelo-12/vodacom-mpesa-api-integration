package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"log"
)

func ecrypt(publicKey string, key string) string {
	// Step 1: Base64-encoded public key string
	pubKeyString := publicKey

	// Step 2: Decoding the Base64-encoded public key string into bytes
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyString)
	if err != nil {
		log.Fatalf("Failed to decode base64 public key string: %v", err)
	}

	// Step 3: Parsing the decoded public key bytes into an RSA public key
	pubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		log.Fatalf("Failed to parse RSA public key: %v", err)
	}

	// Step 4: API key to encrypt
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		log.Fatalf("The public key is not an RSA key")
	}

	// Step 5: Encrypting the API key using the RSA public key
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, []byte(key))
	if err != nil {
		log.Fatalf("Failed to encrypt API key: %v", err)
	}

	// Step 6: Encoding the encrypted API key in Base64 format
	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedData)

	return encryptedBase64
}

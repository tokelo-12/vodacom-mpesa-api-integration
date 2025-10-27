package main

import "github.com/gin-gonic/gin"

type Response struct {
	OutputResponseCode             string `json:"output_ResponseCode"`
	OutputResponseDesc             string `json:"output_ResponseDesc"`
	OutputTransactionID            string `json:"output_TransactionID"`
	OutputConversationID           string `json:"output_ConversationID"`
	OutputThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
}

type UserData struct {
	PhoneNumber string `json:"phone_number"`
}

func main() {

	endpoints()

}

func endpoints() {
	router := gin.Default()

	router.GET("/c2b", c2bPay)
	router.Run("localhost:8888")
}

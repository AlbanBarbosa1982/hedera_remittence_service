package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

type TransferInput struct {
	AlbanAccountID string `json:"albanAccountId"`
	IrisAccountID  string `json:"irisAccountId"`
	Amount         int64  `json:"amount"`
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print the current working directory
	cwd, _ := os.Getwd()
	log.Println("Current working directory:", cwd)

	router := mux.NewRouter()

	router.HandleFunc("/transfer", transferHandler).Methods("GET", "POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	// Use CORS middleware
	corsHandler := handlers.CORS(headersOk, originsOk, methodsOk)(router)

	// Increase the maximum number of idle connections
	http.DefaultTransport.(*http.Transport).MaxIdleConns = 100
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100

	err = http.ListenAndServe(":8000", corsHandler)
	if err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}

func transferHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("transferHandler invoked") // Log when the function is invoked

	if r.Method != http.MethodPost {
		log.Println("Invalid request method") // Log if invalid request method
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input TransferInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err) // Log error reading request body
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	albanAccountID, err := hedera.AccountIDFromString(input.AlbanAccountID)
	if err != nil {
		log.Printf("Failed to parse Alban's account ID: %v\n", err) // Log failure to parse Alban's account ID
		http.Error(w, "Failed to parse Alban's account ID", http.StatusBadRequest)
		return
	}

	irisAccountID, err := hedera.AccountIDFromString(input.IrisAccountID)
	if err != nil {
		log.Printf("Failed to parse Iris's account ID: %v\n", err) // Log failure to parse Iris's account ID
		http.Error(w, "Failed to parse Iris's account ID", http.StatusBadRequest)
		return
	}

	client := hedera.ClientForTestnet()
	operatorPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("OPERATOR_PRIVATE_KEY"))
	if err != nil {
		log.Printf("Failed to parse operator private key: %v\n", err) // Log failure to parse operator private key
		http.Error(w, "Failed to parse operator private key", http.StatusInternalServerError)
		return
	}
	operatorAccountID := albanAccountID
	client.SetOperator(operatorAccountID, operatorPrivateKey)

	usdcTokenID, err := hedera.TokenIDFromString("0.0.3120049")
	if err != nil {
		log.Printf("Failed to parse USDC token ID: %v", err) // Log failure to parse USDC token ID
		http.Error(w, "Failed to parse USDC token ID", http.StatusInternalServerError)
		return
	}

	_, err = hedera.NewTransferTransaction().
		AddTokenTransfer(usdcTokenID, albanAccountID, -input.Amount).
		AddTokenTransfer(usdcTokenID, irisAccountID, input.Amount).
		SetTransactionID(hedera.TransactionIDGenerate(operatorAccountID)).
		Execute(client)
	if err != nil {
		log.Printf("Failed to transfer USDC: %v", err) // Log if failed to transfer USDC
		http.Error(w, "Failed to transfer USDC", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "USDC transferred successfully!")
}

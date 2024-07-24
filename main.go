package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func main() {
	// Get environment variables
	keyvaultURL := os.Getenv("KEYVAULT_URL")
	if keyvaultURL == "" {
		log.Fatal("KEYVAULT_URL environment variable is not set")
	}
	secretName := os.Getenv("SECRET_NAME")
	if secretName == "" {
		log.Fatal("SECRET_NAME environment variable is not set")
	}
	clientID := os.Getenv("CLIENT_ID")
	tenantID := os.Getenv("TENANT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID == "" || tenantID == "" || clientSecret == "" {
		log.Fatal("CLIENT_ID, TENANT_ID, or CLIENT_SECRET environment variables are not set")
	}

	// Create a credential using the ClientSecretCredential
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		log.Fatalf("failed to create client secret credential: %v", err)
	}

	// Initialize the Key Vault client
	client, err := azsecrets.NewClient(keyvaultURL, cred, nil)
	if err != nil {
		log.Fatalf("failed to create Key Vault client: %v", err)
	}

	// Retrieve the secret from Key Vault
	secretBundle, err := client.GetSecret(context.Background(), secretName, "", nil)
	if err != nil {
		log.Fatalf("failed to get secret from Key Vault: %v", err)
	}

	// Print the secret value
	fmt.Printf("Secret value: %s\n", *secretBundle.Value)
}

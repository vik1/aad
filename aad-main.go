package main

import (
	"context"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"log"
)

func main() {
	keyvaultURL := os.Getenv("KEYVAULT_URL")
	if keyvaultURL == "" {
		log.Fatal("KEYVAULT_URL environment variable is not set")
	}
	secretName := os.Getenv("SECRET_NAME")
	if secretName == "" {
		log.Fatal("SECRET_NAME environment variable is not set")
	}

	// Azure AD Workload Identity webhook will inject the following env vars
	// 	AZURE_CLIENT_ID with the clientID set in the service account annotation
	// 	AZURE_TENANT_ID with the tenantID set in the service account annotation. If not defined, then
	// 	the tenantID provided via azure-wi-webhook-config for the webhook will be used.
	// 	AZURE_FEDERATED_TOKEN_FILE is the service account token path
	// 	AZURE_AUTHORITY_HOST is the AAD authority hostname
	clientID := os.Getenv("AZURE_CLIENT_ID")
	tenantID := os.Getenv("AZURE_TENANT_ID")
	tokenFilePath := os.Getenv("AZURE_FEDERATED_TOKEN_FILE")
	authorityHost := os.Getenv("AZURE_AUTHORITY_HOST")

	if clientID == "" {
		log.Fatal("AZURE_CLIENT_ID environment variable is not set")
	}
	if tenantID == "" {
		log.Fatal("AZURE_TENANT_ID environment variable is not set")
	}
	if tokenFilePath == "" {
		log.Fatal("AZURE_FEDERATED_TOKEN_FILE environment variable is not set")
	}
	if authorityHost == "" {
		log.Fatal("AZURE_AUTHORITY_HOST environment variable is not set")
	}

	cred, err := newClientAssertionCredential(tenantID, clientID, authorityHost, tokenFilePath, nil)
	if err != nil {
		log.Fatal(err)
	}

	// initialize keyvault client
	client, err := azsecrets.NewClient(keyvaultURL, cred, &azsecrets.ClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		secretBundle, err := client.GetSecret(context.Background(), secretName, "", nil)
		if err != nil {
			log.Printf("failed to get secret from keyvault")
			//os.Exit(1)
		} 
		log.Printf("successfully got secret")

		// wait for 60 seconds before polling again
		time.Sleep(60 * time.Second)
	}
}

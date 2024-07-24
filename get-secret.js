const { ClientSecretCredential } = require('@azure/identity');
const { SecretClient } = require('@azure/keyvault-secrets');
require('dotenv').config(); // Load environment variables from .env file

async function main() {
    // Get environment variables
    const keyvaultUrl = process.env.KEYVAULT_URL;
    const secretName = process.env.SECRET_NAME;
    const clientId = process.env.CLIENT_ID;
    const tenantId = process.env.TENANT_ID;
    const clientSecret = process.env.CLIENT_SECRET;

    if (!keyvaultUrl) {
        throw new Error('KEYVAULT_URL environment variable is not set');
    }
    if (!secretName) {
        throw new Error('SECRET_NAME environment variable is not set');
    }
    if (!clientId || !tenantId || !clientSecret) {
        throw new Error('CLIENT_ID, TENANT_ID, or CLIENT_SECRET environment variables are not set');
    }

    // Create a credential using the ClientSecretCredential
    const credential = new ClientSecretCredential(tenantId, clientId, clientSecret);

    // Initialize the Key Vault client
    const client = new SecretClient(keyvaultUrl, credential);

    try {
        // Retrieve the secret from Key Vault
        const secret = await client.getSecret(secretName);
        // Print the secret value
        console.log(`Secret value: ${secret.value}`);
    } catch (err) {
        console.error(`Failed to get secret from Key Vault: ${err.message}`);
    }
}

main().catch((err) => {
    console.error(`Error running script: ${err.message}`);
});

!!! note
    Ensure the relevant environment variables are set to allow Quorum to authenticate with Vault, see [Configuring Quorum to use Vault](Configuring-Quorum-to-use-Vault.md) for details

New accounts can be created with the `geth account new` CLI command.  New accounts can be stored in password encrypted files (as with `geth`) or in a Hashicorp Vault.

The `--hashicorp` and related CLI options tell Quorum to store the account in a vault and provide it with the necessary connection details:
```bash
$ geth account new --help | grep -i hashicorp
HASHICORP VAULT OPTIONS:
  --hashicorp                   Store the newly created account in a Hashicorp Vault
  --hashicorp.url value         Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/
  --hashicorp.approle value     Vault path for an enabled Approle auth method (requires VAULT_ROLE_ID and VAULT_SECRET_ID env vars to be set) (default: "approle")
  --hashicorp.clientcert value  Path to a PEM-encoded client certificate. Required when communicating with the Vault server using TLS
  --hashicorp.clientkey value   Path to an unencrypted, PEM-encoded private key which corresponds to the matching client certificate
  --hashicorp.cacert value      Path to a PEM-encoded CA certificate file. Used to verify the Vault server's SSL certificate
  --hashicorp.engine value      Vault path for an enabled KV v2 secret engine
  --hashicorp.name value        Vault path for a new or existing KV secret.  Any existing secret at the path will be overwritten
  --hashicorp.accountid value   ID for account/address component of Vault secret
  --hashicorp.keyid value       ID for private key component of Vault secret
```
For example, to create and store a new account in a Vault that is not using TLS:
```bash
geth account new --hashicorp --hashicorp.url http://127.0.0.1:8200 \ 
                 --hashicorp.name mySecret --hashicorp.engine kv \
                 --hashicorp.accountid account --hashicorp.keyid key
```

This account will be stored at the Vault path `http://127.0.0.1:8200/v1/kv/data/mySecret` (see the Vault docs for more information on how Vault stores secrets).  The account is stored in the vault as 2 parts: 

* The account address (`hashicorp.accountid`)
* The account private key (`hashicorp.keyid`)

``` json
{
    ...
    "data" : {
      "account" : "a0580332e76c97E4367De186825Ca33231408d42", // hex string representation of the address
      "key" : "12b51404091ae14e044186f20355ed6f369109357660ef26376be3cc0f47dba8" // hex string representation of private key 
    }
    ...
}
```

Saving a new account to an existing secret will overwrite the values stored at that secret. Previous versions may be retained and be retrievable depending on how the K/V secrets engine is configured.

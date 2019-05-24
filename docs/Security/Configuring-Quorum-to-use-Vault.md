To use Vault-based accounts with Quorum:

1. Configure the connection to the Vault
1. Provide the necessary authentication credentials  

## Configuring connection to Vault
Vault configuration is provided in the Quorum `.toml` configuration file:
``` bash
geth --config /path/to/config.toml ...
```

Quorum can connect to multiple vaults, and multiple accounts can be stored in each vault.  

See the [toml documentation](https://github.com/toml-lang/toml) for syntax details.

Each vault is configured as a toml table in an array, i.e. `[[Node.HashicorpVaults]]`.

Each vault's configuration is composed of two parts:

1. `[Node.HashicorpVaults.Client]` - toml table to configure the Quorum connection to the Vault server
1. `[[Node.HashicorpVaults.Secrets]]` - toml array of tables to configure the secrets (i.e. accounts) to be retrieved from the vault 


### Client configuration fields
      
```toml
[Node.HashicorpVaults.Client]
Url = "hostname:port of the Vault server" 
Approle = "(optional, defaults to 'approle') Vault path for an enabled Approle auth method"
ClientCert = "(optional) path to a PEM-encoded client certificate. Required when communicating with the Vault server using TLS"
ClientKey = "(optional) path to an unencrypted, PEM-encoded private key which corresponds to the matching client certificate"
CaCert = "(optional) path to a PEM-encoded CA certificate file. Used to verify the Vault server's SSL certificate"
EnvVarPrefix = "(optional) prefix to apply to environment variable names when fetching authentication credentials"
```

!!! caution
    The Vault client library used by Quorum can also be configured by setting the [default Vault environment variables](https://www.vaultproject.io/docs/commands/#environment-variables) (e.g. setting `VAULT_CACERT` instead of `CaCert`).  Any values set in the `.toml` file will take precedence.  
    
    It is not recommended to configure the client using these environment variables as this configuration will be applied to all clients unless explicitly overridden.  Be aware that if these environment variables are set the actual configuration being applied to the created clients may not be the same as has been configured in the `.toml`.

### Secret configuration fields
```toml
[[Node.HashicorpVaults.Secrets]]
Name = "name of a Vault KV secret"
SecretEngine = "name of an enabled KV v2 secret engine"
Version = "version of the secret to retrieve, must be non-zero"
AccountID = "ID for account/address component of Vault secret"
KeyID = "ID for private key component of Vault secret"
```

### Example configuration
```toml
[[Node.HashicorpVaults]]
# configuration of first vault
[Node.HashicorpVaults.Client]
Url = "https://localhost:8200"
Approle = "my-auth"
ClientCert = "/path/to/client.pem"
ClientKey = "/path/to/client.key"
CaCert = "/path/to/ca.pem"
EnvVarPrefix = "A"

[[Node.HashicorpVaults.Secrets]]
Name = "node1"
SecretEngine = "kv"
AccountID = "account"
KeyID = "key"

[[Node.HashicorpVaults.Secrets]]
Name = "node1"
SecretEngine = "kv"
AccountID = "account"
KeyID = "key"
Version = 1

[[Node.HashicorpVaults]]
# configuration of second vault
[Node.HashicorpVaults.Client]
Url = "http://localhost:8201"
EnvVarPrefix = "B"

[[Node.HashicorpVaults.Secrets]]
Name = "node1"
SecretEngine = "engine"
AccountID = "acctId"
KeyID = "keyId"
```

A Quorum node using this configuration will attempt to retrieve 3 accounts from 2 vaults:

1. Vault: `https://localhost:8200`
    1. Secret: `https://localhost:8200/v1/kv/data/node1?version=0`
    1. Secret: `https://localhost:8200/v1/kv/data/node1?version=1`
    
    where both secrets are stored in the vault as: 
    ```json
    {
       "account": "<string hex representation of account address>",
       "key": "<string hex representation of account private key>"
    }
    ```
    
1. Vault: `http://localhost:8201`
    1. Secret: `http://localhost:8201/v1/engine/data/node1?version=0`
    
        where the secret is stored in the vault as:
        ```json
        {
           "acctId": "<string hex representation of account address>",
           "keyId": "<string hex representation of account private key>"
        }
        ```

## Authenticating with Vault
Quorum authenticates with Vault servers using credentials provided as environment variables. 

### Global environment variables
If only using one vault, or [creating a new account](Creating-new-accounts-in-Vault.md), authentication credentials can be provided by setting the following environment variables:

* If using the AppRole auth method:
    1. `VAULT_ROLE_ID`
    1. `VAULT_SECRET_ID`

    These credentials are obtained as detailed in the  [Vault AppRole documentation](https://www.vaultproject.io/docs/auth/approle.html#configuration).

* If using a single auth token (e.g. the root token or a token obtained by already authorising with a different Auth method), set:
    1. `VAULT_TOKEN`

### Prefixed environment variables
When connecting to multiple vaults, it is necessary to be able to provide different authentication credentials for each vault.  This is made possible by prefixing each vault's environment variables with a unique string.  

The prefixes are set using the `EnvVarPrefix` option in the `[Node.HashicorpVaults.Client]` configuration. 

* If using the AppRole auth method, set:
    1. `<PREFIX>_VAULT_ROLE_ID`
    1. `<PREFIX>_VAULT_SECRET_ID`

    These credentials are obtained as outlined in the AppRole documentation.

* If using a single auth token (e.g. the root token or a token obtained by already authorising with a different Auth method), set:
    1. `<PREFIX>_VAULT_TOKEN`  

For the two vault example config above, and assuming that both vaults use AppRole authentication, the following environment variables would need to be set:

* `A_VAULT_ROLE_ID`, `A_VAULT_SECRET_ID`
* `B_VAULT_ROLE_ID`, `B_VAULT_SECRET_ID` 

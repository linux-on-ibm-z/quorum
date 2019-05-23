As with geth, Quorum accounts can be stored in password-protected `keystore` files (Ledger and Trezor hardware wallets are not yet supported).  See the [geth documentation](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) for details on using file-based accounts. 
 
In addition to this, Quorum also supports the storage of accounts in a Hashicorp Vault.  This section details how to set up a Vault, configure Quorum to use the Vault and how to create new accounts in the Vault for use with Quorum.

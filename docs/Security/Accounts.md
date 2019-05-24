# Managing accounts

As with geth, Quorum accounts can be stored in password-protected `keystore` files (Ledger and Trezor hardware wallets are not yet supported).  See the [geth documentation](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) for details on using file-based accounts. 
 
In addition to this, Quorum also supports the storage of accounts in a Hashicorp Vault.  This section details how to [set up a Vault](Configuring-Hashicorp-Vault.md), [configure Quorum to use the Vault](Configuring-Quorum-to-use-Vault.md) and how to [create new accounts in the Vault](Creating-new-accounts-in-Vault.md) for use with Quorum.

## Managing accounts in a Hashicorp Vault

Managing Quorum accounts in a Hashicorp Vault offers several benefits over using standard `keystore` files:

* Your account private keys are stored in a Hashicorp Vault which can be deployed on separate infrastructure to your Quorum node  

* Quorum retrieves account private keys from the Vault **when needed** and uses those keys to sign.  Keys are never written to disk and are not held in memory indefinitely.

* New accounts can be created in Quorum and stored directly to a Vault using the Quorum CLI.

* Quorum can retrieve accounts from multiple Vaults 

* Communication with the Vault can be secured using TLS

* Vault enables you to configure permissions on a per-secret basis to ensure account private keys can only be accessed by authorised users/applications 

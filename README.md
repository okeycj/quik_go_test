# Overview

This is a Golang API for managing the wallets of the players of an online casino, where one can create an account, credit a wallet, debit a wallet and also get a wallet balance .

# Tech Stack

* Golang
* Gin
* MySQL Database

## Local Setup

To setup the api on your local machine,

1. Clone this project.

2. Copy and paste the `.env.example` file in the same directory.

3. Rename the newly copied file to `.env`.

4. Setup your local database, open the `.env` file and assign values of `DB_NAME`, `DB_USER`, `DB_PASS`, `DB_HOST`, `JWT_SECRET`
   to the name, username, password, host and your jwt secret respectively for the newly created database.
   
5. Run `go run main.go` to run the application locally

## Endpoints

| Http verb | Path                        | Controller.action                  | Used for                                    |
| --------- | --------------------------- | ---------------------------------- | ------------------------------------------- |
| POST      | `/api/v1/auth/register`     | authController.Login               | Login a to a user account                   |
| GET       | `/api/v1/auth/login`        | authController.Register            | Register an account for a user              |
| POST      | `/api/v1/wallet/:id/balance`| walletController.GetWalletBalance  | Get a user wallet account balance           |
| POST      | `/api/v1/wallet/:id/credit` | walletController.CreditWallet      | Credit a user wallet                        |
| GET       | `/api/v1/wallet/:id/debit`  | walletController.DebitWallet       | Debit a user wallet                         |





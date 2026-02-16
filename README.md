# Documentation 

### Project Structure
wallet-app/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── domain/
        └── user.go
│   ├── repository/
        └── user_repository_impl.go
        └── user_repository.go
│   ├── usecase/
        └── wallet_usecase.go
│   └── handler/
        └── wallet_handler.go
│
└── pkg/
    └── database.go
│   ├── Postman/
│       └── Wallet.postman_collection.json
│   ├── Database/
        └── wallet_db.sql


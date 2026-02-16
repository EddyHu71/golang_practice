# Documentation 

### Project Structure

golang_wallet/
│
├── cmd/
│   └── main.go                     # Application entry point
│
├── internal/
│   ├── domain/
│   │   └── user.go                 # Core business entity
│   │
│   ├── repository/
│   │   ├── user_repository.go      # Repository interface
│   │   └── user_repository_impl.go # GORM implementation
│   │
│   ├── usecase/
│   │   └── wallet_usecase.go       # Business logic layer
│   │
│   └── handler/
│       └── wallet_handler.go       # HTTP handler layer
│
├── pkg/
│   └── database.go                 # Database initialization
│
├── Postman/
│   └── Wallet.postman_collection.json  # API testing collection
│
└── Database/
    └── wallet_db.sql               # Database schema & seed data



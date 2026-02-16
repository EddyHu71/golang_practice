<img width="1140" height="185" alt="image" src="https://github.com/user-attachments/assets/d321a6bf-4f40-48ff-8538-f61eee0c509e" /># Documentation 

### Project Structure

<pre> ```golang_wallet/
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
``` </pre>



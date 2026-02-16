# Documentation 

### Project Structure

<pre>golang_wallet/
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
</pre>

How to run the project:
1. Run the xampp then import the databaase that contains in Database folder
2. Go to CMD folder then type ``` go run main.go ``` in terminal
3. If you want to test the API, you can use postman that contains in Postman folder



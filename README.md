# GoLib: A RESTful Library Management API 📚⚙️

## ✨ Overview
Welcome to **GoLib**, a robust REST API designed for managing a modern library system. Built with Go (Golang), Fiber, and PostgreSQL, this backend service provides a clean, scalable, and efficient foundation for a library application. It follows Clean Architecture principles, separating business logic from implementation details, making it highly maintainable and testable.

## 🔋 Key Features
  * 🔐 **JWT Authentication** — Secure endpoints using JSON Web Tokens (JWT), ensuring that only authenticated users can access protected resources.
  * 🏗️ **Clean Architecture** — Organized into distinct layers (API, Service, Repository) for a clear separation of concerns, making the codebase easy to understand, test, and scale.
  * 📦 **Full CRUD Operations** — Comprehensive Create, Read, Update, and Delete functionality for all core entities:
      * **Users**: Manage user data with secure password hashing (bcrypt).
      * **Customers/Members**: Manage library members.
      * **Books**: Manage book catalog information.
      * **Book Stocks**: Track individual copies of each book.
      * **Journals**: Handle borrowing and returning transactions.
  * 🖼️ **Media Uploads** — A dedicated endpoint to handle static file uploads, such as book covers.
  * 🛡️ **Request Validation** — Built-in validation for incoming data using `go-playground/validator` to ensure data integrity before it reaches the business logic.
  * Postgres **Database Integration** — Utilizes PostgreSQL for robust and reliable data storage, with a powerful SQL query builder (`goqu`) for safe and efficient database interactions.
  * 🗑️ **Soft Deletes** — Implements soft deletion for critical data, preserving records for audit trails instead of permanently erasing them.
  * ⚙️ **Centralized Configuration** — Manages all environment-specific settings (database credentials, JWT secrets) securely through a `.env` file.

## 🧑‍💻 How It Works
1.  **User authenticates** by sending their email and password to the `/login` endpoint to receive a JWT.
2.  **The client includes the JWT** as a Bearer Token in the `Authorization` header for all subsequent requests to protected endpoints.
3.  **The JWT Middleware** intercepts and validates the token before allowing the request to proceed.
4.  **The API layer** receives the request, validates the data, and calls the appropriate method in the **Service layer**.
5.  **The Service layer** executes the core business logic and coordinates with the **Repository layer**.
6.  **The Repository layer** builds and executes SQL queries to interact with the PostgreSQL database.
7.  **A structured JSON response** is returned to the client.

## ⚙️ Tech Stack
* 🚀 **Go (Golang)**
* ⚡ **Fiber** (Web Framework)
* 🐘 **PostgreSQL** (Database)
* ⛓️ **Go-qu** (SQL Query Builder)
* 🔐 **golang-jwt/jwt** (JWT Implementation)
* 🛡️ **bcrypt** (Password Hashing)
* ✅ **go-playground/validator** (Struct Validation)
* 📄 **godotenv** (Environment Variable Management)

## 🚀 Getting Started
Follow these steps to get the GoLib up and running on your local machine.

### Prerequisites
* [Go](https://go.dev/doc/install) (version 1.18 or higher)
* [PostgreSQL](https://www.postgresql.org/download/)
* A tool to interact with your database (e.g., TablePlus, DBeaver, or a VS Code extension)

### Installation & Setup
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/LouisFernando1204/golib-backend.git
    cd golib-backend
    ```

2.  **Set up environment variables:**
      * Rename the `.env.example` file to `.env`.
      * Open the `.env` file and fill in your database credentials and a secure JWT secret key.
    
    <!-- end list -->
    ```env
    # Server Configuration
    SERVER_HOST=localhost
    SERVER_PORT=9000

    # Database Configuration
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=golang_restapi
    DB_USER=your_postgres_user
    DB_PASS=your_postgres_password
    DB_TZ=Asia/Jakarta

    # JWT Configuration
    JWT_KEY=your_super_secret_and_long_jwt_key
    JWT_EXP=60 # Expiration in minutes
    ```

3.  **Set up the database:**
      * Start your PostgreSQL server.
      * Create a new database with the name you specified in `DB_NAME`.
      * Run the necessary SQL scripts (SetupDatabase.sql) to create the tables.

4.  **Run the application:**
    ```bash
    go run main.go
    ```
    The server should now be running on `http://localhost:9000`.

## 🤝 Contributor
  * 🧑‍💻 **Louis Fernando** : [@LouisFernando1204](https://github.com/LouisFernando1204)

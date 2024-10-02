# SecureGoAuth

## Description
**SecureGoAuth** is a simple Go-based REST API that implements JWT (JSON Web Token) authentication. Users can log in with credentials, receive a JWT, and access protected routes using the token.

## Features
- User login with JWT token generation.
- JWT-based token verification for protected routes.

## Prerequisites
- Go 1.23.1 or later
- Gorilla Mux and JWT libraries (automatically installed via `go mod`)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/SecureGoAuth.git
   cd SecureGoAuth
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run .
   ```

## API Endpoints

### 1. `/login` (POST)
- **Description**: Authenticates the user and provides a JWT token.
- **Request**: 
   - Method: `POST`
   - Body (JSON):
     ```json
     {
       "username": "user1",
       "password": "password1"
     }
     ```
- **Response**: A JWT token is set in a cookie.

### 2. `/welcome` (GET)
- **Description**: A protected route that requires a valid JWT token.
- **Request**: 
   - Method: `GET`
   - Cookie: `token`
- **Response**: User-specific data if the token is valid.

## Testing with cURL

### 1. Login and store JWT:
```bash
curl -X POST http://localhost:8000/login -H "Content-Type: application/json" -d '{"username":"user1", "password":"password1"}' -c cookie.txt
```

### 2. Access protected route:
```bash
curl -X GET http://localhost:8000/welcome -b cookie.txt
```


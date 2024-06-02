# AuthN-AuthZ

## Product Requirements Document (PRD)

### Feature: User Authentication and Authorization

#### Overview
This document outlines the requirements for implementing a robust user authentication and authorization system in a Go application. The system will support user registration, login, password management, and role-based access control.

#### Objectives
- Ensure secure user authentication using JWT tokens.
- Implement role-based access control to restrict access to certain endpoints.
- Provide a simple API for user registration and login.
- Implement password reset functionality.

### Functional Requirements

#### 1. User Registration
- **Endpoint**: `/api/register`
- **Method**: `POST`
- **Description**: Allow users to create a new account.
- **Request Payload**:
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string"
  }
  ```
- **Response**:
  - **Success**: `201 Created`
    ```json
    {
      "message": "User registered successfully"
    }
    ```
  - **Failure**: `400 Bad Request`
    ```json
    {
      "error": "User already exists"
    }
    ```

#### 2. User Login
- **Endpoint**: `/api/login`
- **Method**: `POST`
- **Description**: Allow users to log in and receive a JWT token.
- **Request Payload**:
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "token": "jwt-token"
    }
    ```
  - **Failure**: `401 Unauthorized`
    ```json
    {
      "error": "Invalid credentials"
    }
    ```

#### 3. Password Reset
- **Endpoint**: `/api/reset-password`
- **Method**: `POST`
- **Description**: Allow users to reset their password.
- **Request Payload**:
  ```json
  {
    "email": "string"
  }
  ```
- **Response**:
  - **Success**: `200 OK`
    ```json
    {
      "message": "Password reset link sent"
    }
    ```
  - **Failure**: `404 Not Found`
    ```json
    {
      "error": "Email not found"
    }
    ```

#### 4. Role-Based Access Control
- **Description**: Implement roles such as `Admin`, `User`, and restrict access to certain endpoints based on roles.
- **Roles**:
  - `Admin`: Access to all endpoints.
  - `User`: Access to general user endpoints.
- **Endpoints**:
  - `/api/admin/*`: Accessible only by `Admin`.
  - `/api/user/*`: Accessible by both `Admin` and `User`.

### Non-Functional Requirements

#### 1. Security
- Use bcrypt for password hashing.
- JWT for token generation and verification.

#### 2. Performance
- The system should handle up to 1000 concurrent users.
- Token generation and verification should be performant.

#### 3. Scalability
- The system should be able to scale horizontally.


### Technical Requirements

#### 1. Language and Framework
- **Language**: Go
- **Framework**: Gin (for HTTP server)

#### 2. Database
- **Type**: PostgreSQL
- **Tables**:
  - `users`:
    ```sql
    CREATE TABLE users (
      id SERIAL PRIMARY KEY,
      username VARCHAR(50) UNIQUE NOT NULL,
      email VARCHAR(100) UNIQUE NOT NULL,
      password VARCHAR(255) NOT NULL,
      role VARCHAR(20) DEFAULT 'User',
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    SELECT * FROM users;
    ```

#### 3. Libraries and Tools
- `github.com/gin-gonic/gin` for HTTP server
- `github.com/dgrijalva/jwt-go` for JWT
- `golang.org/x/crypto/bcrypt` for password hashing
- `github.com/jinzhu/gorm` for ORM
- `github.com/swaggo/swag` for Swagger documentation

### Milestones and Timeline

#### 1. Phase 1: User Registration and Login (2 weeks)
- Implement user registration endpoint.
- Implement user login endpoint.
- Implement JWT token generation.

#### 2. Phase 2: Password Management (1 week)
- Implement password reset endpoint.

#### 3. Phase 3: Role-Based Access Control (2 weeks)
- Implement role-based access control.
- Protect endpoints based on user roles.

#### 4. Phase 4: Testing and Documentation (1 week)
- Write unit tests and integration tests.

### Risks and Mitigation
- **Security Risks**: Ensure all data is encrypted, use secure libraries, conduct regular security audits.
- **Scalability Issues**: Design the system to be stateless where possible, use a load balancer for distributing traffic.

### Conclusion
This PRD outlines the requirements and specifications for implementing a secure and scalable user authentication and authorization system in a Go application. The proposed system aims to enhance the security and usability of the application while providing a robust foundation for future features and expansions.

### Validations and Screenshots


#### **Successful registration**

<img width="1003" alt="Screenshot 2024-05-25 at 1 33 17 PM" src="https://github.com/tarunngusain08/AuthN-AuthZ/assets/36428256/c2e17287-ff12-4f24-ab25-a8ac6c7edfe0">



#### **Duplicacy Checks**

<img width="1001" alt="Screenshot 2024-05-25 at 1 29 56 PM" src="https://github.com/tarunngusain08/AuthN-AuthZ/assets/36428256/730cade5-2754-4bef-b62e-0e154b537cc3">



#### **Database Entries For Registered Users**

<img width="1512" alt="Screenshot 2024-05-25 at 1 30 05 PM" src="https://github.com/tarunngusain08/AuthN-AuthZ/assets/36428256/637fbe02-8d0e-4b27-9792-944ad81597d6">
<img width="1008" alt="Screenshot 2024-05-25 at 1 33 56 PM" src="https://github.com/tarunngusain08/AuthN-AuthZ/assets/36428256/0fd79ac0-389b-4115-9617-bfbcc8fdf61b">



#### **Logging in User and JWT Token Generation**
<img width="997" alt="Screenshot 2024-06-01 at 7 05 52 PM" src="https://github.com/tarunngusain08/AuthN-AuthZ/assets/36428256/fe1ba847-62b0-45a2-80e1-e0d2d9e0a89f">

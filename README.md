# API Docs
The API Docs of simple e-commerse API, built with Golang, Gin framework, JWT authentication, and MySQL database. The API provides several routes for user authentication, managing transactions in a shop, and user setting or preferences.

##  **Authentication**
The authentication section provides endpoints for user registration and login to secure access to the API.

### 1.  Register New User
> `POST`  */api/register*

This endpoint allows users to create a new account by providing the necessary information.  
<details>
   <summary><b>Example</b></summary>
   
   ##### Payload
   -  `username` (string): The desired username for the new user.
   -  `fullname` (string): The full name of the user.
   -  `password` (string): The password for the new user account.
   
   ##### Request
   ```http
   POST /api/register
   Content-Type: application/json
   
   {
      "username": "john_doe",
      "fullname": "John Doe",
      "password": "P@ssw0rd"
   }
   ```
   
   ##### Response
   ```http
   HTTP/1.1 201 Created
   Content-Type: application/json
   
   {
      "message": "User registration successful."
   }
   ```
</details>

### 2. User Login
> `POST` */api/login*

This endpoint allows users to authenticate themselves and obtain an access token (JWT) for accessing protected routes.
<details>
   <summary><b>Example</b></summary>
   
   ##### Payload
   - username (string): The username of the registered user.
   - password (string): The corresponding password for the user.
   
   ##### Request
   ```http
   POST /api/login
   Content-Type: application/json

   {
      "username": "john_doe",
      "password": "P@ssw0rd"
   }
   ```
   
   ##### Response
   ```http
   HTTP/1.1 200 OK
   Content-Type: application/json

   {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...<JWT Token>"
   }
   ```
</details>

## User Setting

### 1. Update Avatar Image
> `POST` */api/setting/avatar-image*

Update avatar image with specific resolution and dimension, this endpoint have image validation
<details>
   <summary><b>Example</b></summary>
   
   ##### Request
   ```http
   POST /api/setting/avatar-image HTTP/1.1
   Host: example.com
   Content-Type: multipart/form-data; boundary="avatar-image"

   --avatar-image
   Content-Disposition: form-data; name="file"; filename="your-avatar-image.png"
   Content-Type: image/png

   <binary data of the image file>
   --avatar-image--
   ```
   
   ##### Response
   ```http
   HTTP/1.1 200 OK
   Content-Type: application/json

   {
      "message": "File upload successfully !!"
   }
   ```
</details>

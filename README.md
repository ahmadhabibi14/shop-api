# API Docs
The API Docs of simple e-commerse API, built with Golang, Gin framework, JWT authentication, and MySQL database. The API provides several routes for user authentication, managing transactions in a shop, and user setting or preferences.

## Public routes

-  `POST` */api/login*
   A Login route to Authenticate user and get Authorization
   *  Payload:
      ```json
      {
         "username": "string",
         "password": "string"
      }
      ```
   *  Response:
      ```json
      {
         "token": "string"
      }
      ```
-  `POST` */api/register*
   Register to make user, if username is exist, you gotta try another username
   *  Payload:
      ```json
      {
         "username": "string",
         "fullname": "string",
         "password": "string"
      }
      ```
   *  Response:
      ```json
      {
         "message": "User created successfully !"
      }
      ```
      redirect to */api/login*

## Protected Route (auth required)
-  `POST`

## Test ::::

- **Authentication**
```sh
# Register
curl -X POST -d '{"username": "vl3k0", "password": "wow123"}' localhost:8080/api/register

# Login
curl -X POST -d '{"username": "vl3k0", "password": "wow123"}' localhost:8080/api/login
```
- **Request to protected endpoint**
```sh
# Via Header
curl -H "Authorization: Bearer <JWT_TOKEN>" <API_ENDPOINT_URL>

# Via query URL
curl -X GET http://localhost:8080/<API_ENDPOINT_URL>?token=<TokenJWT>
```

- **Post New Transaction**
```sh
curl -X POST \
   > -H "Authorization: Bearer <JWT_TOKEN>" \
   > -d '{"menu": "Mobil hotwheels", "price": 150000, "qty": 67, "payment": "BANK Jago", "total": 4}' \
   > http://localhost:8080/<API_ENDPOINT_URL>
```

## TODO

- [x] Catch `user_id` from jwt, convert data type from interface to string
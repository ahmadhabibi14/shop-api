## Test ::::

- **Authentication**
```sh
# Register
curl -X POST localhost:8080/api/register \
   > -d '{"username": "vl3k0", "fullname": "Your Full Name", "password": "wow123"}'

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

- **Update Avatar**
```sh
curl -X POST \
   > -F "file=@/home/habi/test.png" \
   > -H "Content-Type: multipart/form-data" \
   > -H "Authorization: Bearer <JWT_TOKEN>" \
   > http://localhost:8080/api/setting/avatar-image
```

## TODO

- [x] Catch `user_id` from jwt, convert data type from interface to string
- [ ] INSERT to avatar file name to database
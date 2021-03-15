# money-send-api

## usage

### Add user
```
curl -X POST 'http://localhost:1323/users' -H 'Content-Type: application/json'-d '{"name":"J.Y Park"}'
```

### Find all users
```
curl -X GET 'http://localhost:1323/users'
```

### Delete user
```
curl -X DELETE 'http://localhost:1323/users/delete/1'
```
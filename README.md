# money-send-api

## usage

### Login
```
curl -X POST 'http://localhost:1323/login' -H 'Content-Type: application/json' -d '{"name":"taro","password":"taro_pass"}'
```

### Regist
```
curl -X POST 'http://localhost:1323/regist' -H 'Content-Type: application/json' -H 'Authorization: Bearer {YOUR_JWT}' -d '{"id":123,"name":"taro", "password":"taro_pass", "balance":1000, "is_balance_receivable":true}'
```

### Find all users
```
curl -X GET 'http://localhost:1323/users/list'
```

### Find user
```
curl -X GET 'http://localhost:1323/users/123'
```

### Update all user's balance
```
curl -X PUT 'http://localhost:1323/users/balance' -H 'Content-Type: application/json' -d '{"balance":10000}'
```

### Update user's balance
```
curl -X PUT 'http://localhost:1323/users/balance/123' -d 'balance=2000'
```

### Delete user
```
curl -X DELETE 'http://localhost:1323/users/delete/123'
```

## Tests

### 1. infrastructure
```
docker-compose exec app go test -v ./infrastructure
```
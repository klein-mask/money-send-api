# money-send-api

## usage

### Add user
```
curl -X POST 'http://localhost:1323/users/add' -H 'Content-Type: application/json' -d '{"id":123,"name":"taro", "balance":1000, "is_balance_receivable":true}'
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
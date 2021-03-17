# money-send-api

## usage

### Add user
```
curl -X POST 'http://localhost:1323/users' -H 'Content-Type: application/json' -d '{"id":123,"name":"taro", "balance":1000}'
```

### Find all users
```
curl -X GET 'http://localhost:1323/users'
```

### Find user
```
curl -X GET 'http://localhost:1323/users/123'
```

### Update user's balance
```
curl -X PUT 'http://localhost:1323/users/balance/123' -d 'balance=2000'
```

### Delete user
```
curl -X DELETE 'http://localhost:1323/users/delete/123'
```


### References
- [GOのORMを分かりやすくまとめてみた【GORM公式ドキュメントの焼き回し】](https://qiita.com/gold-kou/items/45a95d61d253184b0f33)
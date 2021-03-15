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

### Delete user
```
curl -X DELETE 'http://localhost:1323/users/delete/1'
```


### References
- [GOのORMを分かりやすくまとめてみた【GORM公式ドキュメントの焼き回し】](https://qiita.com/gold-kou/items/45a95d61d253184b0f33)
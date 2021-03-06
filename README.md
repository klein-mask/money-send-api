# 💰 Money-send-api
![screenshot](https://user-images.githubusercontent.com/50162453/111777908-41f7f180-88f7-11eb-94bf-8c2289e467dd.png)

**This api is send or receive money users.**

## 📂 Directories
```
├── README.md
├── docker
│   ├── golang
│   │   └── Dockerfile
│   └── postgres
│       ├── Dockerfile
│       └── init
│           └── create-table.sql
├── docker-compose.yml
├── images
│   ├── er.png
│   └── screenshot.PNG
└── src
    ├── docs
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── domain
    │   └── user.go
    ├── go.mod
    ├── go.sum
    ├── infrastructure
    │   ├── router.go
    │   ├── router_test.go
    │   └── sql_handler.go
    ├── interfaces
    │   ├── api
    │   │   └── user_controller.go
    │   └── database
    │       ├── sql_handler.go
    │       └── user_repository.go
    ├── main.go
    ├── money-send-api
    └── usecase
        ├── user_interactor.go
        └── user_repository.go
```

### docker
- **Create docker image build files**

### src
- **Golang api code files**

### src/docs
- **Swagger document files**


## 🍺 Usage

### 1. Clone this repository and cd this dir.
```
git clone https://github.com/klein-mask/money-send-api.git
cd money-send-api
```

### 2. Build and start docker containers.
```
docker-compose up -d --build
```

### 3. Serve API

go run or build
#### go run
```
docker-compose exec app go run main.go&
```

#### build run
```
docker-compose exec app go build
docker-compose exec app ./money-send-api&
```

### 4. End serve API
```
docker-compose down
```
---

## 📘 Swagger document
### http://localhost:1323/swagger/index.html

## 🎁 API Examples
### Regist
- **Regist new user account**

#### Request
```
curl -X POST 'http://localhost:1323/regist' -H 'Content-Type: application/json' -d '{"id":9999, "name":"example_user", "password":"example_user_pass", "balance":1000, "is_balance_receivable":true}'
```

#### Response
```
{
  "ID": 9999,
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z",
  "DeletedAt": null,
  "name": "example_user",
  "password": "***********",
  "balance": 1000,
  "is_balance_receivable": true
}
```
### Login
- **Login your user account**

#### Request
```
curl -X POST 'http://localhost:1323/login' -H 'Content-Type: application/json' -d '{"name":"example_user","password":"example_user_pass"}'
```

#### Response
```
{
    "token":"Bearer {YOUR_JWT}"
}
```

---

### 💭 <span style="color: pink;">Later api must use jwt token in header</span>

### GetUsers
- **Get registed user list**

#### Request
```
curl -X GET 'http://localhost:1323/api/users/list' -H 'Authorization: Bearer {YOUR_JWT}'
```

#### Response
```
[
  {
    "ID": 9999,
    "CreatedAt": "2021-03-19T08:40:47.848351Z",
    "UpdatedAt": "2021-03-19T08:40:47.848351Z",
    "DeletedAt": null,
    "name": "example_user",
    "password": "***********",
    "balance": 1000,
    "is_balance_receivable": true
  }
]
```

### GetUser
- **Get single user account by user id**

#### Request
```
curl -X GET 'http://localhost:1323/api/users/list/{USER_ID}' -H 'Authorization: Bearer {YOUR_JWT}'
```

#### Response
```
{
  "ID": 9999,
  "CreatedAt": "2021-03-19T09:27:49.64932Z",
  "UpdatedAt": "2021-03-19T09:27:49.64932Z",
  "DeletedAt": null,
  "name": "example_user",
  "password": "***********",
  "balance": 1000,
  "is_balance_receivable": true
}
```

### UpdateAllUserBalance
- **Update all user's balance**
- **Request balance value add current user's balance**
- **Update target is user flag "is_balance_receivable=true"**

|  current user's balance  |  request balance value  | is_balance_receivable flag | new user's balance |
| ---- | ---- | ---- | ---- |
|  100  |  300  |  true  |  400  |
|  1000  |  -200  |  true  |  800  |
|  1000  |  300  |  false  |  1000  |

#### Request
```
curl -X PUT 'http://localhost:1323/api/users/balance' -H 'Content-Type: application/json' -H 'Authorization: Bearer {YOUR_JWT}' -d '{"balance":10000}'
```

#### Response
```
"[Success] : Updated all user's balance."
```

### UpdateUserBalance
- **Update single user's balance by user id**
- `This API does not depend on flag "is_balance_receivable"`

|  current user's balance  |  request balance value  | is_balance_receivable flag | new user's balance |
| ---- | ---- | ---- | ---- |
|  100  |  300  |  true  |  400  |
|  1000  |  300  |  false  |  1300  |

#### Request
```
curl -X PUT 'http://localhost:1323/api/users/balance/{USER_ID}' -H 'Content-Type: application/json' -H 'Authorization: Bearer {YOUR_JWT}' -d '{"balance":10000}'
```

#### Response
```
[Success] : Updated balance User id <{USER_ID}>.
```

### DeleteUser
- **Delete user account by user id**

#### Request
```
curl -X DELETE 'http://localhost:1323/api/users/delete/{USER_ID}' -H 'Authorization: Bearer {YOUR_JWT}'
```

#### Response
```
[Success] : Deleted User id <{USER_ID}>.
```

## 🏧 Tests
```
docker-compose exec app go test -v ./infrastructure
```

## Database
### users
![er](https://user-images.githubusercontent.com/50162453/111777952-4d4b1d00-88f7-11eb-9002-c05ce4e1d153.png)

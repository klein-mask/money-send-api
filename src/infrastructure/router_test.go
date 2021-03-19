package infrastructure

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "strings"
    "strconv"
    "encoding/json"
    "money-send-api/domain"
    _ "bytes"
)

var (
    testUsers []TestUser = initTestUsers()
    jwtToken string
)

type TestUser struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Password string `json:"password"`
    Balance string `json:"balance"`
    IsBalanceReceivable string `json:"is_balance_receivable"`
    JwtToken string `json: "-"`
}

type JWT struct {
    Token string
}

func initTestUsers() []TestUser {
    var tus []TestUser 
    tus = append(tus, TestUser{"999999999", "test_user_1", "test_user_1_pass", "100", "true", ""})
    tus = append(tus, TestUser{"999999998", "test_user_2", "test_user_2_pass", "200", "false", ""})
    tus = append(tus, TestUser{"999999997", "test_user_3", "test_user_3_pass", "300", "true", ""})
    return tus
}

func newTestPostgresDSN() *PostgresDSN {
    pd := new(PostgresDSN)
    pd.host = "test_postgres"
    pd.user = "test_admin"
    pd.password = "test_admin_pass"
    pd.dbname = "test_app"
    pd.port = "5432"
    pd.sslmode = "disable"

    return pd
}



func TestHealthcheckHandler(t *testing.T) {
    router := NewRouter(newTestPostgresDSN())

    req := httptest.NewRequest("GET", "/healthcheck", nil)
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "healthcheck ok", rec.Body.String())
}

func TestRegist(t *testing.T) {
    router := NewRouter(newTestPostgresDSN())

    user := testUsers[0]
    // 登録テスト用ユーザーのレコードが既に存在する場合は事前に削除する
    //NewSqlHandler().DeleteById(&domain.User{}, user.Id)

    //jsonDataBytes, _ := json.Marshal(&user)
    //t.Log(string(jsonDataBytes))
    
    jsonData := `{"id":` + user.Id + `,"name":"` + user.Name + `","password":"` + user.Password + `","balance":` + user.Balance + `,"is_balance_receivable":` + user.IsBalanceReceivable + `}`
    bodyReader := strings.NewReader(jsonData)
 
    req := httptest.NewRequest("POST", "/regist", bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
 
    rec := httptest.NewRecorder()
 
    router.ServeHTTP(rec, req)

    u := domain.User{}
    json.Unmarshal([]byte(rec.Body.String()), &u)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, user.Name, u.Name)
}

func TestLogin(t *testing.T) {
    router := NewRouter(newTestPostgresDSN())

    user := testUsers[0]
    jsonData := `{"name":"` + user.Name + `","password":"` + user.Password + `"}`
    bodyReader := strings.NewReader(jsonData)
    req := httptest.NewRequest("POST", "/login", bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
 
    rec := httptest.NewRecorder()
    router.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)

    jwt := JWT{}
    json.Unmarshal([]byte(rec.Body.String()), &jwt)


    jwtToken = jwt.Token
}

func TestGetAllUsers(t *testing.T) {
    router := NewRouter(newTestPostgresDSN())

    req := httptest.NewRequest("GET", "/api/users/list", nil)
    req.Header.Add("Authorization", jwtToken)

    rec := httptest.NewRecorder()
    router.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUser(t *testing.T) {
    user := testUsers[0]

    router := NewRouter(newTestPostgresDSN())

    req := httptest.NewRequest("GET", "/api/users/list/" + user.Id, nil)
    req.Header.Add("Authorization", jwtToken)

    rec := httptest.NewRecorder()
    router.ServeHTTP(rec, req)

    u := domain.User{}
    json.Unmarshal([]byte(rec.Body.String()), &u)
    assert.Equal(t, http.StatusOK, rec.Code)

    assert.Equal(t, user.Name, u.Name)
}

func TestUpdateAllBalance(t *testing.T) {
    users := testUsers
    handler := NewSqlHandler(newTestPostgresDSN())
    router := NewRouter(newTestPostgresDSN())

    for i := 0; i < len(users); i++ {
        // テスト用ユーザー3人分を初期化
        handler.DeleteById(&domain.User{}, users[i].Id)

        u := domain.User{}
        idInt64, _ := strconv.ParseInt(users[i].Id, 10, 64)
        u.ID = uint(idInt64)
        u.Name = users[i].Name
        u.Password = users[i].Password
        u.Balance, _ = strconv.ParseInt(users[i].Balance, 10, 64)
        u.IsBalanceReceivable, _ = strconv.ParseBool(users[i].IsBalanceReceivable)
        handler.Create(&u)
    }

    const addBalance int64 = 10000

    requestData := `{"balance":` + strconv.FormatInt(addBalance, 10) + "}"
    bodyReader := strings.NewReader(requestData)

    req := httptest.NewRequest("PUT", "/api/users/balance", bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", jwtToken)

    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)

    // 指定した金額分追加されているか確認
    for i := 0; i < len(users); i++ {
        u := domain.User{}
        handler.FindById(&u, users[i].Id)
        newBalance := u.Balance
        oldBalance, _ := strconv.ParseInt(users[i].Balance, 10, 64)
        if u.IsBalanceReceivable {
            oldBalance += addBalance
        }
        assert.Equal(t, oldBalance, newBalance)
    }
}


func TestUpdateBalance(t *testing.T) {
    user := testUsers[0]
    handler := NewSqlHandler(newTestPostgresDSN())
    router := NewRouter(newTestPostgresDSN())

    handler.DeleteById(&domain.User{}, user.Id)
    u := domain.User{}
    idInt64, _ := strconv.ParseInt(user.Id, 10, 64)
    u.ID = uint(idInt64)
    u.Name = user.Name
    u.Password = user.Password
    u.Balance, _ = strconv.ParseInt(user.Balance, 10, 64)
    u.IsBalanceReceivable, _ = strconv.ParseBool(user.IsBalanceReceivable)
    handler.Create(&u)

    const addBalance int64 = 10000

    requestData := `{"balance":` + strconv.FormatInt(addBalance, 10) + "}"
    bodyReader := strings.NewReader(requestData)

    req := httptest.NewRequest("PUT", "/api/users/balance/" + user.Id, bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", jwtToken)

    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)

    oldBalance, _ := strconv.ParseInt(user.Balance, 10, 64)
    handler.FindById(&u, user.Id)
    addedBalance := u.Balance
    assert.Equal(t, (oldBalance + addBalance), addedBalance)


    const subBalance int64 = -500

    requestData = `{"balance":` + strconv.FormatInt(subBalance, 10) + "}"
    bodyReader = strings.NewReader(requestData)

    req = httptest.NewRequest("PUT", "/api/users/balance/" + user.Id, bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Authorization", jwtToken)

    rec = httptest.NewRecorder()

    router.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)

    handler.FindById(&u, user.Id)
    subedBalance := u.Balance
    assert.Equal(t, (addedBalance + subBalance), subedBalance)

}

func TestDeleteUser(t *testing.T) {
    user := testUsers[0]
    router := NewRouter(newTestPostgresDSN())
    handler := NewSqlHandler(newTestPostgresDSN())

    req := httptest.NewRequest("DELETE", "/api/users/delete/" + user.Id, nil)
    req.Header.Add("Authorization", jwtToken)

    rec := httptest.NewRecorder()
    router.ServeHTTP(rec, req)

    u := domain.User{}
    handler.FindById(&u, user.Id)
    assert.Equal(t, "", u.Name)
}

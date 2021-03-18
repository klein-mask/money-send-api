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
)

type TestUser struct {
    id string
    name string
    balance string
    isBalanceReceivable string
}

func TestHealthcheckHandler(t *testing.T) {
    router := NewRouter()

    req := httptest.NewRequest("GET", "/healthcheck", nil)
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "healthcheck ok", rec.Body.String())
}

func testUsers() []TestUser {
    var users []TestUser
    users = append(users, TestUser{"999999999", "test_user_1", "100", "true"})
    users = append(users, TestUser{"999999998", "test_user_2", "100", "false"})
    users = append(users, TestUser{"999999997", "test_user_3", "100", "true"})

    return users
}

func TestAddUser(t *testing.T) {

    user := testUsers()[0]

    // 登録テスト用ユーザーのレコードが既に存在する場合は事前に削除する
    NewSqlHandler().DeleteById(&domain.User{}, user.id)
    
    router := NewRouter()
    
    jsonData := `{"id":` + user.id + `,"name":"` + user.name + `","balance":` + user.balance + `,"is_balance_receivable":` + user.isBalanceReceivable + `}`

    bodyReader := strings.NewReader(jsonData)
 
    req := httptest.NewRequest("POST", "/users/add", bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
 
    rec := httptest.NewRecorder()
 
    router.ServeHTTP(rec, req)

    u := domain.User{}
    json.Unmarshal([]byte(rec.Body.String()), &u)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, user.name, u.Name)
}

func TestGetAllUsers(t *testing.T) {
    router := NewRouter()

    req := httptest.NewRequest("GET", "/users/list", nil)
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUser(t *testing.T) {
    user := testUsers()[0]

    router := NewRouter()

    req := httptest.NewRequest("GET", "/users/list/" + user.id, nil)
    rec := httptest.NewRecorder()
    router.ServeHTTP(rec, req)

    u := domain.User{}
    json.Unmarshal([]byte(rec.Body.String()), &u)
    assert.Equal(t, user.name, u.Name)
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateAllBalance(t *testing.T) {
    users := testUsers()
    handler := NewSqlHandler()
    router := NewRouter()

    for i := 0; i < len(users); i++ {
        // テスト用ユーザー3人分を初期化
        handler.DeleteById(&domain.User{}, users[i].id)

        u := domain.User{}
        idInt64, _ := strconv.ParseInt(users[i].id, 10, 64)
        u.ID = uint(idInt64)
        u.Name = users[i].name
        u.Balance, _ = strconv.ParseInt(users[i].balance, 10, 64)
        u.IsBalanceReceivable, _ = strconv.ParseBool(users[i].isBalanceReceivable)
        handler.Create(&u)
    }

    const addBalance int64 = 10000

    requestData := `{"balance":` + strconv.FormatInt(addBalance, 10) + "}"
    //requestData := `{"balance":10000}`
    bodyReader := strings.NewReader(requestData)

    req := httptest.NewRequest("PUT", "/users/balance", bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")

    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)

    // 指定した金額分追加されているか確認
    for i := 0; i < len(users); i++ {
        u := domain.User{}
        handler.FindById(&u, users[i].id)
        newBalance := u.Balance
        oldBalance, _ := strconv.ParseInt(users[i].balance, 10, 64)
        if u.IsBalanceReceivable {
            oldBalance += addBalance
        }
        assert.Equal(t, newBalance, oldBalance)
    }
}

package infrastructure

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "strings"
    "encoding/json"
    "money-send-api/domain"
    //"github.com/labstack/echo"
)

const (
    id string = "999999999"
    name string = "test_user"
    balance string = "100"
    isBalanceReceivable string = "true"

    updateUserBalance string = "50"
    updateAllUserBalance string = "200"
)

func TestHealthcheckHandler(t *testing.T) {
    router := NewRouter()

    req := httptest.NewRequest("GET", "/healthcheck", nil)
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "healthcheck ok", rec.Body.String())
}

func TestAddUser(t *testing.T) {
    // テスト用ユーザーのレコードは事前に削除する
    NewSqlHandler().DeleteById(&domain.User{}, id)

    router := NewRouter()
    
    jsonData := `{"id":` + id + `,"name":"` + name + `","balance":` + balance + `,"is_balance_receivable":` + isBalanceReceivable + `}`

    bodyReader := strings.NewReader(jsonData)
 
    req := httptest.NewRequest("POST", "/users/add", bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
 
    rec := httptest.NewRecorder()
 
    router.ServeHTTP(rec, req)

    u := domain.User{}
    json.Unmarshal([]byte(rec.Body.String()), &u)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, name, u.Name)
}

func TestGetAllUsers(t *testing.T) {
    router := NewRouter()

    req := httptest.NewRequest("GET", "/users/list", nil)
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetUser(t *testing.T) {
    router := NewRouter()

    req := httptest.NewRequest("GET", "/users/list/" + id, nil)
    rec := httptest.NewRecorder()

    router.ServeHTTP(rec, req)

    u := domain.User{}
    json.Unmarshal([]byte(rec.Body.String()), &u)
    assert.Equal(t, name, u.Name)
    assert.Equal(t, http.StatusOK, rec.Code)
}

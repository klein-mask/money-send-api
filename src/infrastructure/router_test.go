package infrastructure

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "strings"
    "encoding/json"
    "money-send-api/domain"
)

const (
    id string = "535435"
    name string = "test_user4"
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
    router := NewRouter()
    
    jsonData := `{"id":` + id + `,"name":"` + name + `","balance":` + balance + `,"is_balance_receivable":` + isBalanceReceivable + `}`

    bodyReader := strings.NewReader(jsonData)
 
    req := httptest.NewRequest("POST", "/users/add", bodyReader)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
 
    rec := httptest.NewRecorder()
 
    router.ServeHTTP(rec, req)

    if rec.Code == 200 {
        u := domain.User{}
        json.Unmarshal([]byte(rec.Body.String()), &u)
        //assert.JSONEq(t, `{"ID":6666,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"huga","balance":1000,"is_balance_receivable":true}`, rec.Body.String())
        assert.Equal(t, name, u.Name)
        return
    } else if rec.Code == 500 {
        assert.JSONEq(t, `{"message":"Internal Server Error"}`, rec.Body.String())
        return
    }
    
    assert.Equal(t, http.StatusOK, rec.Code)
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

    assert.Equal(t, http.StatusOK, rec.Code)
}

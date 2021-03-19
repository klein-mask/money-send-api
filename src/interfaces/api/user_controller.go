package controllers

import (
    "net/http"
    "money-send-api/domain"
    "money-send-api/interfaces/database"
    "money-send-api/usecase"
    "github.com/labstack/echo/v4"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "time"
    _ "fmt"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

type LoginUser struct {
    Name string `json:"name"`
    Password string `json:"password"`
}

type JsonData struct {
    Balance int64 `json:"balance"`
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func passwordToHash(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(hash), err
}


// @Summary login user account.
// @Description login registed user account by id/pass
// @Param data body string true "data"
// @Success 200 {object} map[string]string{token=string}
// @Failure 401 {error} http.StatusUnauthorized
// @Failure 500 {error} http.StatusInternalServerError
// @Router /login [post]
func (controller *UserController) Login(c echo.Context) error {
    lu := LoginUser{}
    c.Bind(&lu)

    name := lu.Name
    password := lu.Password

    err := controller.Interactor.Login(name, password)

    if err != nil {
        return err
        //return c.JSON(err.Code, err.Message)
    }
    token := jwt.New(jwt.SigningMethodHS256)

    // set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["name"] = name
    claims["admin"] = true
    claims["iat"] = time.Now().Unix()
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // generate encoded token and send it as response
    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, map[string]string{
        "token": "Bearer " + t,
    })

    return echo.ErrUnauthorized
}

// @Summary Regist user account.
// @Description Regist is create new user account
// @Param data body string true "data"
// @Success 200
// @Failure 500 {error} http.StatusInternalServerError
// @Router /regist [post]
func (controller *UserController) Regist(c echo.Context) error {
    u := domain.User{}
    c.Bind(&u)

    hashedPassword, err := passwordToHash(u.Password)
    if err != nil {
        return err
    }

    u.Password = hashedPassword

    err = controller.Interactor.Regist(u)
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, u)
}

// @Summary get all users
// @Description get all users from database, use gorm
// @Success 200
// @Failure 500 {error} Error
// @Security Bearer
// @Router /api/users/list [get]
func (controller *UserController) GetAllUsers(c echo.Context) error {
    users, err := controller.Interactor.GetAllUsers()
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, users)
}

// @Summary get user by id
// @Description get user by id
// @Param user_id path string true "user_id"
// @Success 200
// @Failure 500 {error} Error
// @Security Bearer
// @Router /api/users/list/{user_id} [get]
func (controller *UserController) GetUser(c echo.Context) error {
    userId := c.Param("user_id")
    user, err := controller.Interactor.GetUser(userId)
    if err != nil {
        return err
    } else if user.Name == "" {
        msg := "[Not Found] : User id <" + userId + "> is not found."
        return c.JSON(http.StatusOK, msg)
    }
    return c.JSON(http.StatusOK, user)
}

// @Summary update all user's balance
// @Description update all user's balance add or sub
// @Param data body string true "data"
// @Success 200
// @Failure 500 {error} Error
// @Security Bearer
// @Router /api/users/balance [post]
func (controller *UserController) UpdateAllBalance(c echo.Context) error {
    jsonData := JsonData{}
    c.Bind(&jsonData)
    err := controller.Interactor.UpdateAllBalance(jsonData.Balance)
    if err != nil {
        return err
    }
    msg := "[Success] : Updated all user's balance."
    return c.JSON(http.StatusOK, msg)
}

// @Summary update user's balance by user id
// @Description update user's balance add or sub by user id
// @Param user_id path string true "user_id"
// @Param data body string true "data"
// @Success 200
// @Failure 500 {error} Error
// @Security Bearer
// @Router /api/users/balance/{user_id} [post]
func (controller *UserController) UpdateBalance(c echo.Context) error {
    userId := c.Param("user_id")
    jsonData := JsonData{}
    c.Bind(&jsonData)
    err := controller.Interactor.UpdateBalance(userId, jsonData.Balance)
    if err != nil {
        return err
    }
    msg := "[Success] : Updated balance User id <" + userId + ">."
    return c.String(http.StatusOK, msg)
}

// @Summary delete user
// @Description delete user record by id
// @Param user_id path string true "user_id"
// @Success 200
// @Failure 500 {error} Error
// @Security Bearer
// @Router /api/users/delete/{user_id} [post]
func (controller *UserController) DeleteUser(c echo.Context) error {
    userId := c.Param("user_id")
    err := controller.Interactor.DeleteUser(userId)
    if err != nil {
        return err
    }
    msg := "[Success] : Deleted User id <" + userId + ">."
    return c.String(http.StatusOK, msg)
}
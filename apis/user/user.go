package user

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/microcats/docking/apis/modules/recover"
    "github.com/microcats/docking/apis/models"
)

type User struct {
    Username        string `form:"username" binding:"required"`
    Password        string `form:"password" binding:"required"`
    ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
    Email           string `form:"email" binding:"required,email"`
}

func add(c *gin.Context) {
    defer recover.HttpResponse(c)
    var user User
    if err := c.Bind(&user); err != nil {
        panic(&recover.HttpErrorHandler{http.StatusBadRequest, recover.CodeValidateError, err})
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(&recover.HttpErrorHandler{http.StatusInternalServerError, recover.CodeValidateError, err})
    }

    result, err := models.NewUser(user.Username, string(hashedPassword), user.Email).Add()
    if result != true {
        panic(&recover.HttpErrorHandler{http.StatusInternalServerError, recover.CodeModelError, err})
    }

    c.JSON(http.StatusCreated, gin.H{
        "code": "200",
        "message": "Success",
        "data": result,
    })
}

func get(c *gin.Context) {
    username := c.Param("username")
    u := new(models.User)
    u.Username = username
    user, _ := u.Get()

    c.JSON(http.StatusOK, gin.H{
        "code": "200",
        "message": "success",
        "data": user,
    })
}

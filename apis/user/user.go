package user

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/microcats/docking/apis/models"
)

type User struct {
    Username        string `form:"username" binding:"required"`
    Password        string `form:"password" binding:"required"`
    ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
    Email           string `form:"email" binding:"required,email"`
}

func add(c *gin.Context) {
    var user User
    if c.Bind(&user) == nil {
        log.Println(user.Username)
        log.Println(user.Password)
        log.Println(user.ConfirmPassword)
        log.Println(user.Email)
        models.NewUser(user.Username, user.Password, user.Email).Add()
        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"status": "1"})
    }
}

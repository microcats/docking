package user

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
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
    if c.Bind(&user) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code": "1",})
    } else {
        password := []byte(user.Password)
        hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "1"})
        }

        result := models.NewUser(user.Username, string(hashedPassword), user.Email).Add()
        if result == true {
            c.JSON(http.StatusOK, gin.H{
                "code": "200",
                "message": "success",
                "data": user,
            })
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"status": "1"})
        }
    }
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

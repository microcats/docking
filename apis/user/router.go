package user

import (
    "github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
    user := router.Group("/user")
    {
        user.GET("/:username", get)
        user.POST("", add)
    }
}

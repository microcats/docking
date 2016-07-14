package main

import (
    "github.com/gin-gonic/gin"
    "github.com/microcats/docking/apis/user"
)

func main() {
    router := gin.Default()
    user.Router(router)
    router.Run(":8080")
}

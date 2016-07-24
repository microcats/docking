package recover

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

const (
    CodeValidateError = 40
    CodeModelError = 50
)

type HttpErrorHandler struct {
    Status  int
    Code    int
    Err     error
}

func HttpResponse(c *gin.Context) {
    if r := recover(); r != nil {
        if _, ok := r.(error); !ok {
            switch f := r.(type) {
            case *HttpErrorHandler:
                c.JSON(f.Status, f.getMessage())
                c.Abort()
            }
        }
    }
}

func (h *HttpErrorHandler) getMessage() *gin.H {
    fmt.Println(h.Err)
    return &gin.H{
        "code": h.Code,
        "message": h.Err.Error(),
    }
}

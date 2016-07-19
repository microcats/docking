package recover

import (
    "github.com/gin-gonic/gin"
)

type HttpErrorHandler struct {
    Context *gin.Context
    Status  int
    //code    int
    //message string
}

func HttpResponse() {
    if r := recover(); r != nil {
        if _, ok := r.(error); !ok {
            switch f := r.(type) {
            case *HttpErrorHandler:
                f.Context.JSON(f.Status, gin.H{"status": f.Status})
            }
        }
    }
}

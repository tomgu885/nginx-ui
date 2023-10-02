package middleware

import (
    "github.com/gin-gonic/gin"
    "nginx-ui/pkg/helper"
)

func XRequestID() gin.HandlerFunc {
    return func(c *gin.Context) {
        _traceId := c.GetHeader("x-request-id")
        if _traceId == "" {
            _traceId = helper.RandStr(50)
        }

        c.Set("x-request-id", _traceId)
    }
}

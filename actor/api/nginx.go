package api

import (
    "github.com/gin-gonic/gin"
    "github.com/tufanbarisyildirim/gonginx"
    "nginx-ui/pkg/nginx"
)

// @Router
func Update(c *gin.Context) {
    cfg := nginx.NginxProxy{}
    cfg.
}

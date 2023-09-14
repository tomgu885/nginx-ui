package server

import (
    "github.com/jpillora/overseer"
    "net/http"
    "nginx-ui/server/internal/boot"
    "nginx-ui/server/internal/logger"
    "nginx-ui/server/internal/nginx"
    "nginx-ui/server/internal/upgrader"
    "nginx-ui/server/router"
)

func GetRuntimeInfo() (r upgrader.RuntimeInfo, err error) {
    return upgrader.GetRuntimeInfo()
}

func Program(state overseer.State) {
    defer logger.Sync()

    logger.Info("Nginx config dir path: " + nginx.GetConfPath())

    boot.Kernel()

    if state.Listener != nil {
        err := http.Serve(state.Listener, router.InitRouter())
        if err != nil {
            logger.Error(err)
        }
    }

    logger.Info("Server exiting")
}

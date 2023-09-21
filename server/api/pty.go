package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"nginx-ui/pkg/logger"
	"nginx-ui/server/internal/pty"
)

func Pty(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// upgrade http to websocket
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(err)
		return
	}

	defer ws.Close()

	p, err := pty.NewPipeLine(ws)

	if err != nil {
		logger.Error(err)
		return
	}

	defer p.Pty.Close()

	errorChan := make(chan error, 1)
	go p.ReadPtyAndWriteWs(errorChan)
	go p.ReadWsAndWritePty(errorChan)

	err = <-errorChan

	if err != nil {
		logger.Error(err)
	}

	return
}

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"log"
	"nginx-ui/server"
	"nginx-ui/server/settings"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "config", "app.ini", "Specify the configuration file")
	flag.Parse()

	settings.Init(confPath)
	fmt.Println("settings.ServerSettings.RunMode", settings.ServerSettings.RunMode)
	gin.SetMode(settings.ServerSettings.RunMode)

	r, err := server.GetRuntimeInfo()

	if err != nil {
		log.Fatalln(err)
	}

	overseer.Run(overseer.Config{
		Program:          server.Program,
		Address:          fmt.Sprintf("%s:%s", settings.ServerSettings.HttpHost, settings.ServerSettings.HttpPort),
		Fetcher:          &fetcher.File{Path: r.ExPath},
		TerminateTimeout: 0,
	})
}

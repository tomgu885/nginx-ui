package main

import (
	"fmt"
	"nginx-ui/cmd/actor/cli"
)

func main() {
	//router.InitRouter().Run(":8080")
	if err := cli.Execute(); err != nil {
		fmt.Println("err:", err)
	}
}

package main

import (
	"gotestProject/app/common"
	"gotestProject/app/web"
)

func main() {
	common.App.Default()
	web.RunServer()
}

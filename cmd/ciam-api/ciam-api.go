package main

import (
	"github.com/chang144/ciam/internal/apiserver"
	"github.com/chang144/golunzi/cli"
)

func main() {
	cli.SetConfigIn("./configs/")
	apiserver.NewApp("ciam").Run()
}

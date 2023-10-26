package main

import (
	"github.com/chang144/golunzi/cli"
	"github.com/chang144/iam/internal/apiserver"
)

func main() {
	cli.SetConfigIn("./configs/")
	apiserver.NewApp("iam").Run()
}

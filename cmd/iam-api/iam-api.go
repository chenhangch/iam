package main

import (
	"github.com/chenhangch/golunzi/cli"
	"github.com/chenhangch/iam/internal/apiserver"
)

func main() {
	cli.SetConfigIn("./configs/")
	apiserver.NewApp("iam").Run()
}

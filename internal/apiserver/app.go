package apiserver

import (
	"fmt"
	"github.com/chang144/ciam/internal/apiserver/options"
	"github.com/chang144/golunzi/cli"
	"github.com/fatih/color"
)

const commandDesc = `The CIAM API server
Find more information at https://github.com/chang144/ciam
`

// NewApp 创建一个带着默认参数的 app cli 对象
func NewApp(basename string) *cli.AppCli {
	opts := options.NewOptions()
	application := cli.NewAppCli("ciam-api",
		basename,
		cli.WithOptions(opts),
		cli.WithDescription(commandDesc),
		cli.WithRunFunc(run(opts)),
	)

	return application
}

// run 在apiserver启动时执行的函数
func run(opts *options.ApiServerOptions) cli.RunFunc {
	return func(basename string) error {

		fmt.Printf(color.YellowString(opts.String()))

		return nil
	}
}

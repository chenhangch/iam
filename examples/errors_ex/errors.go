package main

import (
	"fmt"
	"github.com/chang144/golunzi/errors"
	code "github.com/marmotedu/sample-code"
)

// main **不**支持错误堆栈
func main() {
	if err := funcA(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func funcA() error {
	if err := funcB(); err != nil {
		return errors.Wrap(err, "get User failed")
	}
	return nil
}

func funcB() error {
	return errors.WithCode(code.ErrDatabase, "user '' not found")
}

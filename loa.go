package main

import (
	"fmt"

	"github.com/oamazing/loa/cmd"
)

func init() {
	fmt.Println("init")
}

func main() {
	_ = GetConfig()

	cmd.Execute()
}

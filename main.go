package main

import (
	"github.com/server2565543706/gin-vue/src/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail:", err.Error())
		os.Exit(-1)
	}
}

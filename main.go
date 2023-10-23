package main

import (
	"fmt"
	"os"

	"github.com/timdin/vfs/cmd"
	"github.com/timdin/vfs/configs"
	"github.com/timdin/vfs/storage"
)

func main() {
	configs.LoadConfig()
	storage := storage.InitStorage()
	cmd := cmd.InitCmd(storage)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

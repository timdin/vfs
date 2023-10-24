package main

import (
	"fmt"
	"os"

	"github.com/timdin/vfs/cmd"
	"github.com/timdin/vfs/configs"
	"github.com/timdin/vfs/storage"
)

func main() {
	config := configs.LoadConfig()
	storage := storage.InitStorage(config)
	cmd := cmd.InitCmd(storage)
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

package main

import (
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
		os.Exit(1)
	}
}

package main

import (
	"github.com/timdin/vfs/cmd"
	"github.com/timdin/vfs/configs"
	"github.com/timdin/vfs/storage"
)

func main() {
	configs.LoadConfig()
	storage.InitStorage()
	cmd.Execute()
}

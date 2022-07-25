package main

import (
	"github.com/qxoo/mywallet/client"
	"github.com/qxoo/mywallet/config"
)

const PATH = "./datadir"

func main() {
	client := client.NewCmdClient(config.Config.EthUrl, config.Config.DataDir)
	client.Run()
}

package main

import "github.com/qxoo/mywallet/client"

const PATH = "./datadir"

func main() {
	client := client.NewCmdClient("http://localhost:8545", PATH)
	client.Run()
}

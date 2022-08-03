package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, path, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(path), "..")

	Config.Root = root
	cfg := filepath.Join(root, "config.json")

	fl, err := os.Open(cfg)
	if err != nil {
		log.Fatalln("Open Config file fail: ", err)
	}

	err = json.NewDecoder(fl).Decode(&Config)
	if err != nil {
		log.Fatalln("Load Config file fail: ", err)
	}
}

type Configuration struct {
	EthUrl         string `json:"eth_url"`
	DataDir        string `json:"data_dir"`
	GasLimit       uint64 `json:"gas_limit"`
	MytokenAddress string `json:"mytoken_address"`
	Root           string
}

var Config Configuration

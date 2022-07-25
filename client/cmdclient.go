package client

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/qxoo/mywallet/db"
	"github.com/qxoo/mywallet/wallet"
)

func getDB(dir string) *db.DB {
	mydb, err := db.NewDB(dir)
	if err != nil {
		log.Fatalln("Db load fail: ", err)
	}
	return mydb
}

type CmdClient struct {
	Url  string
	Path string
}

func NewCmdClient(url, path string) *CmdClient {
	return &CmdClient{Url: url, Path: path}
}

func (cli CmdClient) Help() {
	fmt.Println("Command")
	fmt.Println()
	fmt.Println("\tcreate -pass PASSWORD -name NAME --for create new wallet")
	fmt.Println("\tshow --for show all wallet")
	fmt.Println("\tdelete -pass PASSWROD -name NAME --for delete wallet")
	fmt.Println("\timport -words \"xx xx xx ... \" --for import wallet by mnemonic")
	fmt.Println("\tbalance -pass PASSWROD -name NAME --for query account balance")
	fmt.Println("\ttransfer -pass PASSWORD -name NAME -to TOADDR -value VALUE --for transfer from acct to toaddr")
}

func (cli CmdClient) CreateWallet(pass string, name string) {
	mydb := getDB(cli.Path)
	defer mydb.Close()

	exists, err := mydb.Exists(name)
	if exists {
		log.Fatalln("Name Repeat: ", name)
	}
	if err != nil {
		log.Fatalln("Query DB error: ", err)
	}

	w, err := wallet.NewWallet(cli.Path, pass)
	if err != nil {
		log.Fatalln("Create Fail: ", err)
	}

	address := w.Account.Address.Hex()
	if name == "" {
		name = address
	}
	if err := mydb.SaveAddress(name, address); err != nil {
		log.Fatalln("Query DB error: ", err)
	}
	fmt.Println("Create Wallet: ", w.Words)
}

func (cli CmdClient) Show() {
	mydb := getDB(cli.Path)
	defer mydb.Close()

	data, err := mydb.GetAll()
	if err != nil {
		log.Fatalln("Query DB error: ", err)
	}

	fmt.Println("Wallet List")
	fmt.Println()
	fmt.Println("\tName \tAddress")
	fmt.Println("----------------------------------")
	for k, v := range data {
		fmt.Printf("\t%s \t%s\n", k, v)
	}
}

func (cli CmdClient) DeleteWallet(pass string, name string) {
	mydb := getDB(cli.Path)
	defer mydb.Close()

	addr, err := mydb.GetAddress(name)
	if err != nil {
		log.Fatalln("Query Db error: ", err)
	}

	err = wallet.DeleteWallet(cli.Path, pass, addr)
	if err != nil {
		log.Fatalln("Delete Wallet error: ", err)
	}
	err = mydb.Delete(name)
	if err != nil {
		log.Fatalln("Delete Wallet error: ", err)
	}
	fmt.Println("Delete Wallet Success: ", name)
}

func (cli CmdClient) ImportWallet(pass string, name string, words string) {
	mydb := getDB(cli.Path)
	defer mydb.Close()

	exists, err := mydb.Exists(name)
	if exists {
		log.Fatalln("Name Repeat: ", name)
	}
	if err != nil {
		log.Fatalln("Query DB error: ", err)
	}

	w, err := wallet.ImportWallet(words, cli.Path, pass)
	if err != nil {
		log.Fatalln("Create Fail: ", err)
	}

	address := w.Account.Address.Hex()
	if name == "" {
		name = address
	}
	if err := mydb.SaveAddress(name, address); err != nil {
		log.Fatalln("Query DB error: ", err)
	}
	fmt.Println("Import Wallet: ", name)
}

func (cli CmdClient) Transfer(pass string, name string, toaddr string, value float64) {
	mydb := getDB(cli.Path)
	defer mydb.Close()

	addr, err := mydb.GetAddress(name)
	if err != nil {
		log.Fatalln("Query Db error: ", err)
	}

	w, err := wallet.LoadWallet(cli.Path, pass, addr)
	if err != nil {
		log.Fatalln("Load Wallet error: ", err)
	}

	err = w.InitEthClient(cli.Url)
	if err != nil {
		log.Fatalln("Init EthClient error: ", err)
	}

	err = w.Transfer(toaddr, value)
	if err != nil {
		log.Fatalln("Transfer error: ", err)
	}
	fmt.Println("Transfer Success.")
}

func (cli CmdClient) GetBalance(pass string, name string) {
	mydb := getDB(cli.Path)
	defer mydb.Close()

	addr, err := mydb.GetAddress(name)
	if err != nil {
		log.Fatalln("Query Db error: ", err)
	}

	w, err := wallet.LoadWallet(cli.Path, pass, addr)
	if err != nil {
		log.Fatalln("Load Wallet error: ", err)
	}

	err = w.InitEthClient(cli.Url)
	if err != nil {
		log.Fatalln("Init EthClient error: ", err)
	}

	balance, err := w.GetBalance()
	if err != nil {
		log.Fatalln("Get Balacne error: ", err)
	}
	fmt.Println("Balance: ", balance)
}

func (cli CmdClient) Run() {
	if len(os.Args) < 2 {
		log.Fatal("Args Error")
	}

	fmt.Println("==================================")
	fmt.Println()
	switch os.Args[1] {
	case "create":
		cmd := flag.NewFlagSet("create", flag.ExitOnError)
		cmd_pass := cmd.String("pass", "", "PASSWORD")
		cmd_name := cmd.String("name", "", "NAME")
		if err := cmd.Parse(os.Args[2:]); err != nil {
			log.Fatal("Args Error")
		}
		cli.CreateWallet(*cmd_pass, *cmd_name)
	case "show":
		cli.Show()
	case "delete":
		cmd := flag.NewFlagSet("delete", flag.ExitOnError)
		cmd_pass := cmd.String("pass", "", "PASSWORD")
		cmd_name := cmd.String("name", "", "NAME")
		if err := cmd.Parse(os.Args[2:]); err != nil {
			log.Fatal("Args Error")
		}
		cli.DeleteWallet(*cmd_pass, *cmd_name)
	case "import":
		cmd := flag.NewFlagSet("import", flag.ExitOnError)
		cmd_pass := cmd.String("pass", "", "PASSWORD")
		cmd_name := cmd.String("name", "", "NAME")
		cmd_words := cmd.String("words", "", "WORDS")
		if err := cmd.Parse(os.Args[2:]); err != nil {
			log.Fatal("Args Error")
		}
		cli.ImportWallet(*cmd_pass, *cmd_name, *cmd_words)
	case "transfer":
		cmd := flag.NewFlagSet("transfer", flag.ExitOnError)
		cmd_pass := cmd.String("pass", "", "PASSWORD")
		cmd_name := cmd.String("name", "", "NAME")
		cmd_to := cmd.String("to", "", "TOADDR")
		cmd_value := cmd.Float64("value", 0, "VALUE")
		if err := cmd.Parse(os.Args[2:]); err != nil {
			log.Fatal("Args Error")
		}
		cli.Transfer(*cmd_pass, *cmd_name, *cmd_to, *cmd_value)
	case "balance":
		cmd := flag.NewFlagSet("balace", flag.ExitOnError)
		cmd_pass := cmd.String("pass", "", "PASSWORD")
		cmd_name := cmd.String("name", "", "NAME")
		if err := cmd.Parse(os.Args[2:]); err != nil {
			log.Fatal("Args Error")
		}
		cli.GetBalance(*cmd_pass, *cmd_name)
	case "help":
		cli.Help()
	default:
		cli.Help()
	}
	fmt.Println()
	fmt.Println("==================================")
}

package mytoken

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/qxoo/mywallet/config"
	"github.com/qxoo/mywallet/sol"
	"github.com/qxoo/mywallet/wallet"
)

type TokenWallet struct {
	url    string
	wallet *wallet.Wallet
}

func NewTokenWallet(url, path, pass string, address string) (*TokenWallet, error) {
	w, err := wallet.LoadWallet(path, pass, address)
	return &TokenWallet{wallet: w, url: url}, err
}

func (tw *TokenWallet) auth() (*bind.TransactOpts, error) {
	if err := tw.wallet.KeyStore.Unlock(tw.wallet.Account, tw.wallet.Pass); err != nil {
		return nil, err
	}
	if err := tw.wallet.InitEthClient(tw.url); err != nil {
		return nil, err
	}
	defer tw.wallet.CloseClient()

	chainid, err := tw.wallet.Client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return bind.NewKeyStoreTransactorWithChainID(tw.wallet.KeyStore, tw.wallet.Account, chainid)
}

func (tw *TokenWallet) getSol() (*sol.Sol, error) {
	address := common.HexToAddress(config.Config.MytokenAddress)
	return sol.NewSol(address, tw.wallet.Client)
}

func (tw *TokenWallet) Deploy() (string, string, error) {
	auth, err := tw.auth()
	if err != nil {
		return "", "", err
	}

	parsed, err := sol.SolMetaData.GetAbi()
	if err != nil {
		return "", "", err
	}
	if parsed == nil {
		return "", "", errors.New("ABI Error")
	}

	address, tx, _, err := sol.DeploySol(auth, tw.wallet.Client, "Mytoken")
	return address.Hex(), tx.Hash().Hex(), nil
}

func (tw *TokenWallet) Mint(toaddr string, value int64) (string, error) {
	auth, err := tw.auth()
	if err != nil {
		return "", err
	}

	instance, err := tw.getSol()
	if err != nil {
		return "", err
	}
	tx, err := instance.Mint(auth, common.HexToAddress(toaddr), big.NewInt(value))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (tw *TokenWallet) Transfer(toaddr string, value int64) (string, error) {
	auth, err := tw.auth()
	if err != nil {
		return "", err
	}

	instance, err := tw.getSol()
	if err != nil {
		return "", err
	}

	tx, err := instance.Transfer(auth, common.HexToAddress(toaddr), big.NewInt(value))
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (tw *TokenWallet) Balacne(owner string) (int64, error) {
	if err := tw.wallet.InitEthClient(tw.url); err != nil {
		return 0, err
	}
	defer tw.wallet.CloseClient()

	instance, err := tw.getSol()
	if err != nil {
		return 0, err
	}

	val, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(owner))
	if err != nil {
		return 0, err
	}
	return val.Int64(), nil
}

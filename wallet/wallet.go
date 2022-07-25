package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tyler-smith/go-bip39"
)

func CreateWords() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func DerivePrivateKey(path accounts.DerivationPath, masterKey *hdkeychain.ExtendedKey) (*ecdsa.PrivateKey, error) {
	var err error
	key := masterKey
	for _, n := range path {
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}
	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}
	privateKeyEcdsa := privateKey.ToECDSA()
	return privateKeyEcdsa, nil
}

func PrivateFromWords(words string) (*ecdsa.PrivateKey, *common.Address, error) {
	path, err := accounts.ParseDerivationPath("m/44'/60'/0'/0/0")
	if err != nil {
		return nil, nil, err
	}

	seed, err := bip39.NewSeedWithErrorChecking(words, "")
	if err != nil {
		return nil, nil, err
	}
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, nil, err
	}
	privateKey, err := DerivePrivateKey(path, masterKey)
	if err != nil {
		return nil, nil, err
	}
	_publicKey := privateKey.Public()
	publicKey, ok := _publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("public error")
	}
	address := crypto.PubkeyToAddress(*publicKey)
	return privateKey, &address, nil
}

type Wallet struct {
	Account  accounts.Account
	KeyStore *keystore.KeyStore
	Pass     string
	Words    string
	Client   *ethclient.Client
}

func NewWallet(path string, pass string) (*Wallet, error) {
	words, err := CreateWords()
	if err != nil {
		return nil, err
	}
	privatekey, _, err := PrivateFromWords(words)
	if err != nil {
		return nil, err
	}
	wallet := &Wallet{Words: words}
	if err := wallet.Store(path, pass, privatekey); err != nil {
		return nil, err
	}

	return wallet, nil
}

func ImportWallet(words string, path string, pass string) (*Wallet, error) {
	privatekey, _, err := PrivateFromWords(words)
	if err != nil {
		return nil, err
	}
	wallet := &Wallet{Words: words}
	if err := wallet.Store(path, pass, privatekey); err != nil {
		return nil, err
	}

	return wallet, nil
}

func Wallets(path string) []string {
	ks := keystore.NewKeyStore(path, keystore.LightScryptN, keystore.LightScryptP)
	add_list := []string{}
	for _, account := range ks.Accounts() {
		add_list = append(add_list, account.Address.Hex())
	}
	return add_list
}

func LoadWallet(path, pass string, address string) (*Wallet, error) {
	ks := keystore.NewKeyStore(path, keystore.LightScryptN, keystore.LightScryptP)
	_address := common.HexToAddress(address)
	if !ks.HasAddress(_address) {
		return nil, fmt.Errorf("Address %s not exists", address)
	}
	account := accounts.Account{Address: _address}
	signAccount, err := ks.Find(account)
	if err != nil {
		return nil, err
	}
	return &Wallet{Account: signAccount, KeyStore: ks, Pass: pass}, nil
}

func DeleteWallet(path, pass string, address string) error {
	ks := keystore.NewKeyStore(path, keystore.LightScryptN, keystore.LightScryptP)
	_address := common.HexToAddress(address)
	if !ks.HasAddress(_address) {
		return nil
	}
	account := accounts.Account{Address: _address}
	return ks.Delete(account, pass)
}

func (w *Wallet) InitEthClient(url string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	w.Client = client
	return nil
}

func (w *Wallet) Store(path string, pass string, privatekey *ecdsa.PrivateKey) error {
	ks := keystore.NewKeyStore(path, keystore.LightScryptN, keystore.LightScryptP)
	acc, err := ks.ImportECDSA(privatekey, pass)
	if err != nil {
		return err
	}
	w.Account = acc
	w.KeyStore = ks
	return nil
}

func (w *Wallet) Transfer(toaddr string, value float64) error {
	if w.Client == nil {
		return fmt.Errorf("Please Init EthClient")
	}

	if err := w.KeyStore.Unlock(w.Account, w.Pass); err != nil {
		return err
	}

	to_addr := common.HexToAddress(toaddr)

	nonce, err := w.Client.PendingNonceAt(context.Background(), w.Account.Address)
	if err != nil {
		return err
	}

	gasprice, err := w.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	amount := big.NewInt(int64(value * math.Pow10(18)))
	data := []byte("Transfer")
	gaslimit := uint64(30000)

	tx := types.NewTransaction(nonce, to_addr, amount, gaslimit, gasprice, data)

	chainid, err := w.Client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := w.KeyStore.SignTx(w.Account, tx, chainid)
	if err != nil {
		return err
	}

	return w.Client.SendTransaction(context.Background(), signedTx)
}

func (w Wallet) GetBalance() (float64, error) {
	if w.Client == nil {
		return 0, fmt.Errorf("Please Init EthClient")
	}

	balance, err := w.Client.BalanceAt(context.Background(), w.Account.Address, nil)
	if err != nil {
		return 0, err
	}

	tmp := new(big.Float)
	tmp.SetString(balance.String())

	_balance := new(big.Float).Quo(tmp, big.NewFloat(math.Pow10(18)))
	val, _ := _balance.Float64()
	return val, nil
}

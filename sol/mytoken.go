// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sol

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SolMetaData contains all meta data concerning the Sol contract.
var SolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001284380380620012848339818101604052810190620000379190620002e5565b80600490805190602001906200004f92919062000098565b5033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200039a565b828054620000a69062000365565b90600052602060002090601f016020900481019282620000ca576000855562000116565b82601f10620000e557805160ff191683800117855562000116565b8280016001018555821562000116579182015b8281111562000115578251825591602001919060010190620000f8565b5b50905062000125919062000129565b5090565b5b80821115620001445760008160009055506001016200012a565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001b18262000166565b810181811067ffffffffffffffff82111715620001d357620001d262000177565b5b80604052505050565b6000620001e862000148565b9050620001f68282620001a6565b919050565b600067ffffffffffffffff82111562000219576200021862000177565b5b620002248262000166565b9050602081019050919050565b60005b838110156200025157808201518184015260208101905062000234565b8381111562000261576000848401525b50505050565b60006200027e6200027884620001fb565b620001dc565b9050828152602081018484840111156200029d576200029c62000161565b5b620002aa84828562000231565b509392505050565b600082601f830112620002ca57620002c96200015c565b5b8151620002dc84826020860162000267565b91505092915050565b600060208284031215620002fe57620002fd62000152565b5b600082015167ffffffffffffffff8111156200031f576200031e62000157565b5b6200032d84828501620002b2565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200037e57607f821691505b60208210810362000394576200039362000336565b5b50919050565b610eda80620003aa6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a082311461012757806395d89b4114610157578063a9059cbb14610175578063dd62ed3e146101a557610088565b8063095ea7b31461008d57806318160ddd146100bd57806323b872dd146100db57806340c10f191461010b575b600080fd5b6100a760048036038101906100a29190610b70565b6101d5565b6040516100b49190610bcb565b60405180910390f35b6100c56102ff565b6040516100d29190610bf5565b60405180910390f35b6100f560048036038101906100f09190610c10565b610309565b6040516101029190610bcb565b60405180910390f35b61012560048036038101906101209190610b70565b61060a565b005b610141600480360381019061013c9190610c63565b61076d565b60405161014e9190610bf5565b60405180910390f35b61015f6107b5565b60405161016c9190610d29565b60405180910390f35b61018f600480360381019061018a9190610b70565b610843565b60405161019c9190610bcb565b60405180910390f35b6101bf60048036038101906101ba9190610d4b565b610a50565b6040516101cc9190610bf5565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361020f57600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516102ed9190610bf5565b60405180910390a36001905092915050565b6000600254905090565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561035657600080fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211156103df57600080fd5b816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104299190610dba565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104b59190610dee565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461057f9190610dba565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600190509392505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066457600080fd5b806002546106729190610dee565b600281905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546106c29190610dee565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107619190610bf5565b60405180910390a35050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600480546107c290610e73565b80601f01602080910402602001604051908101604052809291908181526020018280546107ee90610e73565b801561083b5780601f106108105761010080835404028352916020019161083b565b820191906000526020600020905b81548152906001019060200180831161081e57829003601f168201915b505050505081565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561089057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108c957600080fd5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109139190610dba565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461099f9190610dee565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a3e9190610bf5565b60405180910390a36001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b0782610adc565b9050919050565b610b1781610afc565b8114610b2257600080fd5b50565b600081359050610b3481610b0e565b92915050565b6000819050919050565b610b4d81610b3a565b8114610b5857600080fd5b50565b600081359050610b6a81610b44565b92915050565b60008060408385031215610b8757610b86610ad7565b5b6000610b9585828601610b25565b9250506020610ba685828601610b5b565b9150509250929050565b60008115159050919050565b610bc581610bb0565b82525050565b6000602082019050610be06000830184610bbc565b92915050565b610bef81610b3a565b82525050565b6000602082019050610c0a6000830184610be6565b92915050565b600080600060608486031215610c2957610c28610ad7565b5b6000610c3786828701610b25565b9350506020610c4886828701610b25565b9250506040610c5986828701610b5b565b9150509250925092565b600060208284031215610c7957610c78610ad7565b5b6000610c8784828501610b25565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610cca578082015181840152602081019050610caf565b83811115610cd9576000848401525b50505050565b6000601f19601f8301169050919050565b6000610cfb82610c90565b610d058185610c9b565b9350610d15818560208601610cac565b610d1e81610cdf565b840191505092915050565b60006020820190508181036000830152610d438184610cf0565b905092915050565b60008060408385031215610d6257610d61610ad7565b5b6000610d7085828601610b25565b9250506020610d8185828601610b25565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610dc582610b3a565b9150610dd083610b3a565b925082821015610de357610de2610d8b565b5b828203905092915050565b6000610df982610b3a565b9150610e0483610b3a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610e3957610e38610d8b565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e8b57607f821691505b602082108103610e9e57610e9d610e44565b5b5091905056fea2646970667358221220a5cb2ae3d3164a24a2500d3b3296dbaaf1d6a65e47ddf6731ef182da64b8689f64736f6c634300080d0033",
}

// SolABI is the input ABI used to generate the binding from.
// Deprecated: Use SolMetaData.ABI instead.
var SolABI = SolMetaData.ABI

// SolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolMetaData.Bin instead.
var SolBin = SolMetaData.Bin

// DeploySol deploys a new Ethereum contract, binding an instance of Sol to it.
func DeploySol(auth *bind.TransactOpts, backend bind.ContractBackend, _symbol string) (common.Address, *types.Transaction, *Sol, error) {
	parsed, err := SolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolBin), backend, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// Sol is an auto generated Go binding around an Ethereum contract.
type Sol struct {
	SolCaller     // Read-only binding to the contract
	SolTransactor // Write-only binding to the contract
	SolFilterer   // Log filterer for contract events
}

// SolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolSession struct {
	Contract     *Sol              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolCallerSession struct {
	Contract *SolCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolTransactorSession struct {
	Contract     *SolTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolRaw struct {
	Contract *Sol // Generic contract binding to access the raw methods on
}

// SolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolCallerRaw struct {
	Contract *SolCaller // Generic read-only contract binding to access the raw methods on
}

// SolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolTransactorRaw struct {
	Contract *SolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSol creates a new instance of Sol, bound to a specific deployed contract.
func NewSol(address common.Address, backend bind.ContractBackend) (*Sol, error) {
	contract, err := bindSol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// NewSolCaller creates a new read-only instance of Sol, bound to a specific deployed contract.
func NewSolCaller(address common.Address, caller bind.ContractCaller) (*SolCaller, error) {
	contract, err := bindSol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolCaller{contract: contract}, nil
}

// NewSolTransactor creates a new write-only instance of Sol, bound to a specific deployed contract.
func NewSolTransactor(address common.Address, transactor bind.ContractTransactor) (*SolTransactor, error) {
	contract, err := bindSol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolTransactor{contract: contract}, nil
}

// NewSolFilterer creates a new log filterer instance of Sol, bound to a specific deployed contract.
func NewSolFilterer(address common.Address, filterer bind.ContractFilterer) (*SolFilterer, error) {
	contract, err := bindSol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolFilterer{contract: contract}, nil
}

// bindSol binds a generic wrapper to an already deployed contract.
func bindSol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.SolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sol *SolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sol *SolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Sol.Contract.Allowance(&_Sol.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Sol *SolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Sol.Contract.Allowance(&_Sol.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Sol *SolCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Sol *SolSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Sol.Contract.BalanceOf(&_Sol.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Sol *SolCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Sol.Contract.BalanceOf(&_Sol.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sol *SolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sol *SolSession) Symbol() (string, error) {
	return _Sol.Contract.Symbol(&_Sol.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sol *SolCallerSession) Symbol() (string, error) {
	return _Sol.Contract.Symbol(&_Sol.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sol *SolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sol *SolSession) TotalSupply() (*big.Int, error) {
	return _Sol.Contract.TotalSupply(&_Sol.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Sol *SolCallerSession) TotalSupply() (*big.Int, error) {
	return _Sol.Contract.TotalSupply(&_Sol.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Sol *SolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Sol *SolSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Approve(&_Sol.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Sol *SolTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Approve(&_Sol.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Sol *SolTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Sol *SolSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Mint(&_Sol.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Sol *SolTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Mint(&_Sol.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Sol *SolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Sol *SolSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Transfer(&_Sol.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Sol *SolTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.Transfer(&_Sol.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Sol *SolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Sol *SolSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.TransferFrom(&_Sol.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Sol *SolTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sol.Contract.TransferFrom(&_Sol.TransactOpts, from, to, value)
}

// SolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Sol contract.
type SolApprovalIterator struct {
	Event *SolApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolApproval represents a Approval event raised by the Sol contract.
type SolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sol *SolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Sol.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SolApprovalIterator{contract: _Sol.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sol *SolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Sol.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolApproval)
				if err := _Sol.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Sol *SolFilterer) ParseApproval(log types.Log) (*SolApproval, error) {
	event := new(SolApproval)
	if err := _Sol.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Sol contract.
type SolTransferIterator struct {
	Event *SolTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolTransfer represents a Transfer event raised by the Sol contract.
type SolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Sol *SolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Sol.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SolTransferIterator{contract: _Sol.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Sol *SolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Sol.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolTransfer)
				if err := _Sol.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Sol *SolFilterer) ParseTransfer(log types.Log) (*SolTransfer, error) {
	event := new(SolTransfer)
	if err := _Sol.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

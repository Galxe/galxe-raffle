// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// IDrandOracleRandom is an auto generated low-level Go binding around an user-defined struct.
type IDrandOracleRandom struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}

// RaffleMetaData contains all meta data concerning the Raffle contract.
var RaffleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vkey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_drandOracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commitRandomness\",\"inputs\":[{\"name\":\"_questID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"drandOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuest\",\"inputs\":[{\"name\":\"_questID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"random\",\"type\":\"tuple\",\"internalType\":\"structIDrandOracle.Random\",\"components\":[{\"name\":\"round\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_participantCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_winnerCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasParticipated\",\"inputs\":[{\"name\":\"_verifyID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"participate\",\"inputs\":[{\"name\":\"_questID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_user\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_verifyID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reveal\",\"inputs\":[{\"name\":\"_questID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_publicValues\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDrandOracle\",\"inputs\":[{\"name\":\"_drandOracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSigner\",\"inputs\":[{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVerifier\",\"inputs\":[{\"name\":\"_verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVkey\",\"inputs\":[{\"name\":\"_vkey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vkey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CommitRandomness\",\"inputs\":[{\"name\":\"questID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"roundID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DrandOracleUpdated\",\"inputs\":[{\"name\":\"drandOracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Participate\",\"inputs\":[{\"name\":\"participantID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"questID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"user\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"verifyID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reveal\",\"inputs\":[{\"name\":\"questID\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"participantCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"winnerCount\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerUpdated\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierUpdated\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VkeyUpdated\",\"inputs\":[{\"name\":\"vkey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuestID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRandomness\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"QuestAlreadyRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuestNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuestNotExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuestRandomnessAlreadyCommitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuestStillActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"VerifyIdAlreadyUsed\",\"inputs\":[{\"name\":\"verifyId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// RaffleABI is the input ABI used to generate the binding from.
// Deprecated: Use RaffleMetaData.ABI instead.
var RaffleABI = RaffleMetaData.ABI

// Raffle is an auto generated Go binding around an Ethereum contract.
type Raffle struct {
	RaffleCaller     // Read-only binding to the contract
	RaffleTransactor // Write-only binding to the contract
	RaffleFilterer   // Log filterer for contract events
}

// RaffleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaffleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaffleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaffleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaffleSession struct {
	Contract     *Raffle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaffleCallerSession struct {
	Contract *RaffleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RaffleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaffleTransactorSession struct {
	Contract     *RaffleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaffleRaw struct {
	Contract *Raffle // Generic contract binding to access the raw methods on
}

// RaffleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaffleCallerRaw struct {
	Contract *RaffleCaller // Generic read-only contract binding to access the raw methods on
}

// RaffleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaffleTransactorRaw struct {
	Contract *RaffleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaffle creates a new instance of Raffle, bound to a specific deployed contract.
func NewRaffle(address common.Address, backend bind.ContractBackend) (*Raffle, error) {
	contract, err := bindRaffle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Raffle{RaffleCaller: RaffleCaller{contract: contract}, RaffleTransactor: RaffleTransactor{contract: contract}, RaffleFilterer: RaffleFilterer{contract: contract}}, nil
}

// NewRaffleCaller creates a new read-only instance of Raffle, bound to a specific deployed contract.
func NewRaffleCaller(address common.Address, caller bind.ContractCaller) (*RaffleCaller, error) {
	contract, err := bindRaffle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleCaller{contract: contract}, nil
}

// NewRaffleTransactor creates a new write-only instance of Raffle, bound to a specific deployed contract.
func NewRaffleTransactor(address common.Address, transactor bind.ContractTransactor) (*RaffleTransactor, error) {
	contract, err := bindRaffle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleTransactor{contract: contract}, nil
}

// NewRaffleFilterer creates a new log filterer instance of Raffle, bound to a specific deployed contract.
func NewRaffleFilterer(address common.Address, filterer bind.ContractFilterer) (*RaffleFilterer, error) {
	contract, err := bindRaffle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaffleFilterer{contract: contract}, nil
}

// bindRaffle binds a generic wrapper to an already deployed contract.
func bindRaffle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RaffleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Raffle *RaffleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Raffle.Contract.RaffleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Raffle *RaffleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.Contract.RaffleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Raffle *RaffleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Raffle.Contract.RaffleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Raffle *RaffleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Raffle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Raffle *RaffleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Raffle *RaffleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Raffle.Contract.contract.Transact(opts, method, params...)
}

// DrandOracle is a free data retrieval call binding the contract method 0x09c21de0.
//
// Solidity: function drandOracle() view returns(address)
func (_Raffle *RaffleCaller) DrandOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "drandOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DrandOracle is a free data retrieval call binding the contract method 0x09c21de0.
//
// Solidity: function drandOracle() view returns(address)
func (_Raffle *RaffleSession) DrandOracle() (common.Address, error) {
	return _Raffle.Contract.DrandOracle(&_Raffle.CallOpts)
}

// DrandOracle is a free data retrieval call binding the contract method 0x09c21de0.
//
// Solidity: function drandOracle() view returns(address)
func (_Raffle *RaffleCallerSession) DrandOracle() (common.Address, error) {
	return _Raffle.Contract.DrandOracle(&_Raffle.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Raffle *RaffleCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Raffle *RaffleSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Raffle.Contract.Eip712Domain(&_Raffle.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Raffle *RaffleCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Raffle.Contract.Eip712Domain(&_Raffle.CallOpts)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 _questID) view returns(bool _active, (uint64,uint64,bytes32,bytes) random, uint256 _participantCount, uint256 _winnerCount, bytes32 _merkleRoot)
func (_Raffle *RaffleCaller) GetQuest(opts *bind.CallOpts, _questID *big.Int) (struct {
	Active           bool
	Random           IDrandOracleRandom
	ParticipantCount *big.Int
	WinnerCount      *big.Int
	MerkleRoot       [32]byte
}, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "getQuest", _questID)

	outstruct := new(struct {
		Active           bool
		Random           IDrandOracleRandom
		ParticipantCount *big.Int
		WinnerCount      *big.Int
		MerkleRoot       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Random = *abi.ConvertType(out[1], new(IDrandOracleRandom)).(*IDrandOracleRandom)
	outstruct.ParticipantCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WinnerCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MerkleRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 _questID) view returns(bool _active, (uint64,uint64,bytes32,bytes) random, uint256 _participantCount, uint256 _winnerCount, bytes32 _merkleRoot)
func (_Raffle *RaffleSession) GetQuest(_questID *big.Int) (struct {
	Active           bool
	Random           IDrandOracleRandom
	ParticipantCount *big.Int
	WinnerCount      *big.Int
	MerkleRoot       [32]byte
}, error) {
	return _Raffle.Contract.GetQuest(&_Raffle.CallOpts, _questID)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 _questID) view returns(bool _active, (uint64,uint64,bytes32,bytes) random, uint256 _participantCount, uint256 _winnerCount, bytes32 _merkleRoot)
func (_Raffle *RaffleCallerSession) GetQuest(_questID *big.Int) (struct {
	Active           bool
	Random           IDrandOracleRandom
	ParticipantCount *big.Int
	WinnerCount      *big.Int
	MerkleRoot       [32]byte
}, error) {
	return _Raffle.Contract.GetQuest(&_Raffle.CallOpts, _questID)
}

// HasParticipated is a free data retrieval call binding the contract method 0xd74e4ca8.
//
// Solidity: function hasParticipated(uint256 _verifyID) view returns(bool)
func (_Raffle *RaffleCaller) HasParticipated(opts *bind.CallOpts, _verifyID *big.Int) (bool, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "hasParticipated", _verifyID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasParticipated is a free data retrieval call binding the contract method 0xd74e4ca8.
//
// Solidity: function hasParticipated(uint256 _verifyID) view returns(bool)
func (_Raffle *RaffleSession) HasParticipated(_verifyID *big.Int) (bool, error) {
	return _Raffle.Contract.HasParticipated(&_Raffle.CallOpts, _verifyID)
}

// HasParticipated is a free data retrieval call binding the contract method 0xd74e4ca8.
//
// Solidity: function hasParticipated(uint256 _verifyID) view returns(bool)
func (_Raffle *RaffleCallerSession) HasParticipated(_verifyID *big.Int) (bool, error) {
	return _Raffle.Contract.HasParticipated(&_Raffle.CallOpts, _verifyID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Raffle *RaffleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Raffle *RaffleSession) Owner() (common.Address, error) {
	return _Raffle.Contract.Owner(&_Raffle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Raffle *RaffleCallerSession) Owner() (common.Address, error) {
	return _Raffle.Contract.Owner(&_Raffle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Raffle *RaffleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Raffle *RaffleSession) Paused() (bool, error) {
	return _Raffle.Contract.Paused(&_Raffle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Raffle *RaffleCallerSession) Paused() (bool, error) {
	return _Raffle.Contract.Paused(&_Raffle.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Raffle *RaffleCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Raffle *RaffleSession) PendingOwner() (common.Address, error) {
	return _Raffle.Contract.PendingOwner(&_Raffle.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Raffle *RaffleCallerSession) PendingOwner() (common.Address, error) {
	return _Raffle.Contract.PendingOwner(&_Raffle.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Raffle *RaffleCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Raffle *RaffleSession) Signer() (common.Address, error) {
	return _Raffle.Contract.Signer(&_Raffle.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Raffle *RaffleCallerSession) Signer() (common.Address, error) {
	return _Raffle.Contract.Signer(&_Raffle.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Raffle *RaffleCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Raffle *RaffleSession) Verifier() (common.Address, error) {
	return _Raffle.Contract.Verifier(&_Raffle.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Raffle *RaffleCallerSession) Verifier() (common.Address, error) {
	return _Raffle.Contract.Verifier(&_Raffle.CallOpts)
}

// Vkey is a free data retrieval call binding the contract method 0xc69b0eb1.
//
// Solidity: function vkey() view returns(bytes32)
func (_Raffle *RaffleCaller) Vkey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Raffle.contract.Call(opts, &out, "vkey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Vkey is a free data retrieval call binding the contract method 0xc69b0eb1.
//
// Solidity: function vkey() view returns(bytes32)
func (_Raffle *RaffleSession) Vkey() ([32]byte, error) {
	return _Raffle.Contract.Vkey(&_Raffle.CallOpts)
}

// Vkey is a free data retrieval call binding the contract method 0xc69b0eb1.
//
// Solidity: function vkey() view returns(bytes32)
func (_Raffle *RaffleCallerSession) Vkey() ([32]byte, error) {
	return _Raffle.Contract.Vkey(&_Raffle.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Raffle *RaffleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Raffle *RaffleSession) AcceptOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.AcceptOwnership(&_Raffle.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Raffle *RaffleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.AcceptOwnership(&_Raffle.TransactOpts)
}

// CommitRandomness is a paid mutator transaction binding the contract method 0x88c84146.
//
// Solidity: function commitRandomness(uint64 _questID, uint64 _timestamp, bytes _signature) returns()
func (_Raffle *RaffleTransactor) CommitRandomness(opts *bind.TransactOpts, _questID uint64, _timestamp uint64, _signature []byte) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "commitRandomness", _questID, _timestamp, _signature)
}

// CommitRandomness is a paid mutator transaction binding the contract method 0x88c84146.
//
// Solidity: function commitRandomness(uint64 _questID, uint64 _timestamp, bytes _signature) returns()
func (_Raffle *RaffleSession) CommitRandomness(_questID uint64, _timestamp uint64, _signature []byte) (*types.Transaction, error) {
	return _Raffle.Contract.CommitRandomness(&_Raffle.TransactOpts, _questID, _timestamp, _signature)
}

// CommitRandomness is a paid mutator transaction binding the contract method 0x88c84146.
//
// Solidity: function commitRandomness(uint64 _questID, uint64 _timestamp, bytes _signature) returns()
func (_Raffle *RaffleTransactorSession) CommitRandomness(_questID uint64, _timestamp uint64, _signature []byte) (*types.Transaction, error) {
	return _Raffle.Contract.CommitRandomness(&_Raffle.TransactOpts, _questID, _timestamp, _signature)
}

// Participate is a paid mutator transaction binding the contract method 0xc8cbf5e3.
//
// Solidity: function participate(uint64 _questID, uint256 _user, uint64 _verifyID, bytes _signature) returns()
func (_Raffle *RaffleTransactor) Participate(opts *bind.TransactOpts, _questID uint64, _user *big.Int, _verifyID uint64, _signature []byte) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "participate", _questID, _user, _verifyID, _signature)
}

// Participate is a paid mutator transaction binding the contract method 0xc8cbf5e3.
//
// Solidity: function participate(uint64 _questID, uint256 _user, uint64 _verifyID, bytes _signature) returns()
func (_Raffle *RaffleSession) Participate(_questID uint64, _user *big.Int, _verifyID uint64, _signature []byte) (*types.Transaction, error) {
	return _Raffle.Contract.Participate(&_Raffle.TransactOpts, _questID, _user, _verifyID, _signature)
}

// Participate is a paid mutator transaction binding the contract method 0xc8cbf5e3.
//
// Solidity: function participate(uint64 _questID, uint256 _user, uint64 _verifyID, bytes _signature) returns()
func (_Raffle *RaffleTransactorSession) Participate(_questID uint64, _user *big.Int, _verifyID uint64, _signature []byte) (*types.Transaction, error) {
	return _Raffle.Contract.Participate(&_Raffle.TransactOpts, _questID, _user, _verifyID, _signature)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Raffle *RaffleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Raffle *RaffleSession) Pause() (*types.Transaction, error) {
	return _Raffle.Contract.Pause(&_Raffle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Raffle *RaffleTransactorSession) Pause() (*types.Transaction, error) {
	return _Raffle.Contract.Pause(&_Raffle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.RenounceOwnership(&_Raffle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Raffle *RaffleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Raffle.Contract.RenounceOwnership(&_Raffle.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xd131c722.
//
// Solidity: function reveal(uint64 _questID, bytes _publicValues, bytes _proofBytes) returns()
func (_Raffle *RaffleTransactor) Reveal(opts *bind.TransactOpts, _questID uint64, _publicValues []byte, _proofBytes []byte) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "reveal", _questID, _publicValues, _proofBytes)
}

// Reveal is a paid mutator transaction binding the contract method 0xd131c722.
//
// Solidity: function reveal(uint64 _questID, bytes _publicValues, bytes _proofBytes) returns()
func (_Raffle *RaffleSession) Reveal(_questID uint64, _publicValues []byte, _proofBytes []byte) (*types.Transaction, error) {
	return _Raffle.Contract.Reveal(&_Raffle.TransactOpts, _questID, _publicValues, _proofBytes)
}

// Reveal is a paid mutator transaction binding the contract method 0xd131c722.
//
// Solidity: function reveal(uint64 _questID, bytes _publicValues, bytes _proofBytes) returns()
func (_Raffle *RaffleTransactorSession) Reveal(_questID uint64, _publicValues []byte, _proofBytes []byte) (*types.Transaction, error) {
	return _Raffle.Contract.Reveal(&_Raffle.TransactOpts, _questID, _publicValues, _proofBytes)
}

// SetDrandOracle is a paid mutator transaction binding the contract method 0x2ce7db68.
//
// Solidity: function setDrandOracle(address _drandOracle) returns()
func (_Raffle *RaffleTransactor) SetDrandOracle(opts *bind.TransactOpts, _drandOracle common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setDrandOracle", _drandOracle)
}

// SetDrandOracle is a paid mutator transaction binding the contract method 0x2ce7db68.
//
// Solidity: function setDrandOracle(address _drandOracle) returns()
func (_Raffle *RaffleSession) SetDrandOracle(_drandOracle common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetDrandOracle(&_Raffle.TransactOpts, _drandOracle)
}

// SetDrandOracle is a paid mutator transaction binding the contract method 0x2ce7db68.
//
// Solidity: function setDrandOracle(address _drandOracle) returns()
func (_Raffle *RaffleTransactorSession) SetDrandOracle(_drandOracle common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetDrandOracle(&_Raffle.TransactOpts, _drandOracle)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Raffle *RaffleTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setSigner", _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Raffle *RaffleSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetSigner(&_Raffle.TransactOpts, _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_Raffle *RaffleTransactorSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetSigner(&_Raffle.TransactOpts, _signer)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _verifier) returns()
func (_Raffle *RaffleTransactor) SetVerifier(opts *bind.TransactOpts, _verifier common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setVerifier", _verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _verifier) returns()
func (_Raffle *RaffleSession) SetVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetVerifier(&_Raffle.TransactOpts, _verifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _verifier) returns()
func (_Raffle *RaffleTransactorSession) SetVerifier(_verifier common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.SetVerifier(&_Raffle.TransactOpts, _verifier)
}

// SetVkey is a paid mutator transaction binding the contract method 0x972ccdaf.
//
// Solidity: function setVkey(bytes32 _vkey) returns()
func (_Raffle *RaffleTransactor) SetVkey(opts *bind.TransactOpts, _vkey [32]byte) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "setVkey", _vkey)
}

// SetVkey is a paid mutator transaction binding the contract method 0x972ccdaf.
//
// Solidity: function setVkey(bytes32 _vkey) returns()
func (_Raffle *RaffleSession) SetVkey(_vkey [32]byte) (*types.Transaction, error) {
	return _Raffle.Contract.SetVkey(&_Raffle.TransactOpts, _vkey)
}

// SetVkey is a paid mutator transaction binding the contract method 0x972ccdaf.
//
// Solidity: function setVkey(bytes32 _vkey) returns()
func (_Raffle *RaffleTransactorSession) SetVkey(_vkey [32]byte) (*types.Transaction, error) {
	return _Raffle.Contract.SetVkey(&_Raffle.TransactOpts, _vkey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.TransferOwnership(&_Raffle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Raffle *RaffleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Raffle.Contract.TransferOwnership(&_Raffle.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Raffle *RaffleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Raffle.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Raffle *RaffleSession) Unpause() (*types.Transaction, error) {
	return _Raffle.Contract.Unpause(&_Raffle.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Raffle *RaffleTransactorSession) Unpause() (*types.Transaction, error) {
	return _Raffle.Contract.Unpause(&_Raffle.TransactOpts)
}

// RaffleCommitRandomnessIterator is returned from FilterCommitRandomness and is used to iterate over the raw logs and unpacked data for CommitRandomness events raised by the Raffle contract.
type RaffleCommitRandomnessIterator struct {
	Event *RaffleCommitRandomness // Event containing the contract specifics and raw log

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
func (it *RaffleCommitRandomnessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleCommitRandomness)
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
		it.Event = new(RaffleCommitRandomness)
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
func (it *RaffleCommitRandomnessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleCommitRandomnessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleCommitRandomness represents a CommitRandomness event raised by the Raffle contract.
type RaffleCommitRandomness struct {
	QuestID    uint64
	RoundID    uint64
	Randomness [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitRandomness is a free log retrieval operation binding the contract event 0x1dbc60586a945ceca44922ef613ca23dc725b0fd96c644af43119d1876eca608.
//
// Solidity: event CommitRandomness(uint64 questID, uint64 roundID, bytes32 randomness)
func (_Raffle *RaffleFilterer) FilterCommitRandomness(opts *bind.FilterOpts) (*RaffleCommitRandomnessIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "CommitRandomness")
	if err != nil {
		return nil, err
	}
	return &RaffleCommitRandomnessIterator{contract: _Raffle.contract, event: "CommitRandomness", logs: logs, sub: sub}, nil
}

// WatchCommitRandomness is a free log subscription operation binding the contract event 0x1dbc60586a945ceca44922ef613ca23dc725b0fd96c644af43119d1876eca608.
//
// Solidity: event CommitRandomness(uint64 questID, uint64 roundID, bytes32 randomness)
func (_Raffle *RaffleFilterer) WatchCommitRandomness(opts *bind.WatchOpts, sink chan<- *RaffleCommitRandomness) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "CommitRandomness")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleCommitRandomness)
				if err := _Raffle.contract.UnpackLog(event, "CommitRandomness", log); err != nil {
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

// ParseCommitRandomness is a log parse operation binding the contract event 0x1dbc60586a945ceca44922ef613ca23dc725b0fd96c644af43119d1876eca608.
//
// Solidity: event CommitRandomness(uint64 questID, uint64 roundID, bytes32 randomness)
func (_Raffle *RaffleFilterer) ParseCommitRandomness(log types.Log) (*RaffleCommitRandomness, error) {
	event := new(RaffleCommitRandomness)
	if err := _Raffle.contract.UnpackLog(event, "CommitRandomness", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleDrandOracleUpdatedIterator is returned from FilterDrandOracleUpdated and is used to iterate over the raw logs and unpacked data for DrandOracleUpdated events raised by the Raffle contract.
type RaffleDrandOracleUpdatedIterator struct {
	Event *RaffleDrandOracleUpdated // Event containing the contract specifics and raw log

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
func (it *RaffleDrandOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleDrandOracleUpdated)
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
		it.Event = new(RaffleDrandOracleUpdated)
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
func (it *RaffleDrandOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleDrandOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleDrandOracleUpdated represents a DrandOracleUpdated event raised by the Raffle contract.
type RaffleDrandOracleUpdated struct {
	DrandOracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDrandOracleUpdated is a free log retrieval operation binding the contract event 0x0bf1c7001241496ed737646ad58bc2d42d5c086ee9df98ab228601bb47110c7f.
//
// Solidity: event DrandOracleUpdated(address drandOracle)
func (_Raffle *RaffleFilterer) FilterDrandOracleUpdated(opts *bind.FilterOpts) (*RaffleDrandOracleUpdatedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "DrandOracleUpdated")
	if err != nil {
		return nil, err
	}
	return &RaffleDrandOracleUpdatedIterator{contract: _Raffle.contract, event: "DrandOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchDrandOracleUpdated is a free log subscription operation binding the contract event 0x0bf1c7001241496ed737646ad58bc2d42d5c086ee9df98ab228601bb47110c7f.
//
// Solidity: event DrandOracleUpdated(address drandOracle)
func (_Raffle *RaffleFilterer) WatchDrandOracleUpdated(opts *bind.WatchOpts, sink chan<- *RaffleDrandOracleUpdated) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "DrandOracleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleDrandOracleUpdated)
				if err := _Raffle.contract.UnpackLog(event, "DrandOracleUpdated", log); err != nil {
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

// ParseDrandOracleUpdated is a log parse operation binding the contract event 0x0bf1c7001241496ed737646ad58bc2d42d5c086ee9df98ab228601bb47110c7f.
//
// Solidity: event DrandOracleUpdated(address drandOracle)
func (_Raffle *RaffleFilterer) ParseDrandOracleUpdated(log types.Log) (*RaffleDrandOracleUpdated, error) {
	event := new(RaffleDrandOracleUpdated)
	if err := _Raffle.contract.UnpackLog(event, "DrandOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Raffle contract.
type RaffleEIP712DomainChangedIterator struct {
	Event *RaffleEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *RaffleEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleEIP712DomainChanged)
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
		it.Event = new(RaffleEIP712DomainChanged)
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
func (it *RaffleEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleEIP712DomainChanged represents a EIP712DomainChanged event raised by the Raffle contract.
type RaffleEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Raffle *RaffleFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*RaffleEIP712DomainChangedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &RaffleEIP712DomainChangedIterator{contract: _Raffle.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Raffle *RaffleFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *RaffleEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleEIP712DomainChanged)
				if err := _Raffle.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Raffle *RaffleFilterer) ParseEIP712DomainChanged(log types.Log) (*RaffleEIP712DomainChanged, error) {
	event := new(RaffleEIP712DomainChanged)
	if err := _Raffle.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Raffle contract.
type RaffleOwnershipTransferStartedIterator struct {
	Event *RaffleOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *RaffleOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleOwnershipTransferStarted)
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
		it.Event = new(RaffleOwnershipTransferStarted)
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
func (it *RaffleOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Raffle contract.
type RaffleOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RaffleOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleOwnershipTransferStartedIterator{contract: _Raffle.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *RaffleOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleOwnershipTransferStarted)
				if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) ParseOwnershipTransferStarted(log types.Log) (*RaffleOwnershipTransferStarted, error) {
	event := new(RaffleOwnershipTransferStarted)
	if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Raffle contract.
type RaffleOwnershipTransferredIterator struct {
	Event *RaffleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RaffleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleOwnershipTransferred)
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
		it.Event = new(RaffleOwnershipTransferred)
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
func (it *RaffleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleOwnershipTransferred represents a OwnershipTransferred event raised by the Raffle contract.
type RaffleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RaffleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleOwnershipTransferredIterator{contract: _Raffle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RaffleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleOwnershipTransferred)
				if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Raffle *RaffleFilterer) ParseOwnershipTransferred(log types.Log) (*RaffleOwnershipTransferred, error) {
	event := new(RaffleOwnershipTransferred)
	if err := _Raffle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleParticipateIterator is returned from FilterParticipate and is used to iterate over the raw logs and unpacked data for Participate events raised by the Raffle contract.
type RaffleParticipateIterator struct {
	Event *RaffleParticipate // Event containing the contract specifics and raw log

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
func (it *RaffleParticipateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleParticipate)
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
		it.Event = new(RaffleParticipate)
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
func (it *RaffleParticipateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleParticipateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleParticipate represents a Participate event raised by the Raffle contract.
type RaffleParticipate struct {
	ParticipantID *big.Int
	QuestID       uint64
	User          *big.Int
	VerifyID      uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterParticipate is a free log retrieval operation binding the contract event 0x3225dff9fa5146919afeaf97b89eebbc0eedf6971bae7bd9a329e9d4f46c4d04.
//
// Solidity: event Participate(uint256 participantID, uint64 questID, uint256 user, uint64 verifyID)
func (_Raffle *RaffleFilterer) FilterParticipate(opts *bind.FilterOpts) (*RaffleParticipateIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "Participate")
	if err != nil {
		return nil, err
	}
	return &RaffleParticipateIterator{contract: _Raffle.contract, event: "Participate", logs: logs, sub: sub}, nil
}

// WatchParticipate is a free log subscription operation binding the contract event 0x3225dff9fa5146919afeaf97b89eebbc0eedf6971bae7bd9a329e9d4f46c4d04.
//
// Solidity: event Participate(uint256 participantID, uint64 questID, uint256 user, uint64 verifyID)
func (_Raffle *RaffleFilterer) WatchParticipate(opts *bind.WatchOpts, sink chan<- *RaffleParticipate) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "Participate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleParticipate)
				if err := _Raffle.contract.UnpackLog(event, "Participate", log); err != nil {
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

// ParseParticipate is a log parse operation binding the contract event 0x3225dff9fa5146919afeaf97b89eebbc0eedf6971bae7bd9a329e9d4f46c4d04.
//
// Solidity: event Participate(uint256 participantID, uint64 questID, uint256 user, uint64 verifyID)
func (_Raffle *RaffleFilterer) ParseParticipate(log types.Log) (*RaffleParticipate, error) {
	event := new(RaffleParticipate)
	if err := _Raffle.contract.UnpackLog(event, "Participate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RafflePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Raffle contract.
type RafflePausedIterator struct {
	Event *RafflePaused // Event containing the contract specifics and raw log

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
func (it *RafflePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RafflePaused)
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
		it.Event = new(RafflePaused)
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
func (it *RafflePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RafflePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RafflePaused represents a Paused event raised by the Raffle contract.
type RafflePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Raffle *RaffleFilterer) FilterPaused(opts *bind.FilterOpts) (*RafflePausedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RafflePausedIterator{contract: _Raffle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Raffle *RaffleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RafflePaused) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RafflePaused)
				if err := _Raffle.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Raffle *RaffleFilterer) ParsePaused(log types.Log) (*RafflePaused, error) {
	event := new(RafflePaused)
	if err := _Raffle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleRevealIterator is returned from FilterReveal and is used to iterate over the raw logs and unpacked data for Reveal events raised by the Raffle contract.
type RaffleRevealIterator struct {
	Event *RaffleReveal // Event containing the contract specifics and raw log

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
func (it *RaffleRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleReveal)
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
		it.Event = new(RaffleReveal)
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
func (it *RaffleRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleReveal represents a Reveal event raised by the Raffle contract.
type RaffleReveal struct {
	QuestID          uint64
	ParticipantCount uint32
	WinnerCount      uint32
	Randomness       [32]byte
	MerkleRoot       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReveal is a free log retrieval operation binding the contract event 0xa7e58d47fffa9ed45ed22db8f05e0f79a19f861b3d5fbebf5dfc52441f98996f.
//
// Solidity: event Reveal(uint64 questID, uint32 participantCount, uint32 winnerCount, bytes32 randomness, bytes32 merkleRoot)
func (_Raffle *RaffleFilterer) FilterReveal(opts *bind.FilterOpts) (*RaffleRevealIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "Reveal")
	if err != nil {
		return nil, err
	}
	return &RaffleRevealIterator{contract: _Raffle.contract, event: "Reveal", logs: logs, sub: sub}, nil
}

// WatchReveal is a free log subscription operation binding the contract event 0xa7e58d47fffa9ed45ed22db8f05e0f79a19f861b3d5fbebf5dfc52441f98996f.
//
// Solidity: event Reveal(uint64 questID, uint32 participantCount, uint32 winnerCount, bytes32 randomness, bytes32 merkleRoot)
func (_Raffle *RaffleFilterer) WatchReveal(opts *bind.WatchOpts, sink chan<- *RaffleReveal) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "Reveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleReveal)
				if err := _Raffle.contract.UnpackLog(event, "Reveal", log); err != nil {
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

// ParseReveal is a log parse operation binding the contract event 0xa7e58d47fffa9ed45ed22db8f05e0f79a19f861b3d5fbebf5dfc52441f98996f.
//
// Solidity: event Reveal(uint64 questID, uint32 participantCount, uint32 winnerCount, bytes32 randomness, bytes32 merkleRoot)
func (_Raffle *RaffleFilterer) ParseReveal(log types.Log) (*RaffleReveal, error) {
	event := new(RaffleReveal)
	if err := _Raffle.contract.UnpackLog(event, "Reveal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the Raffle contract.
type RaffleSignerUpdatedIterator struct {
	Event *RaffleSignerUpdated // Event containing the contract specifics and raw log

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
func (it *RaffleSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleSignerUpdated)
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
		it.Event = new(RaffleSignerUpdated)
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
func (it *RaffleSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleSignerUpdated represents a SignerUpdated event raised by the Raffle contract.
type RaffleSignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_Raffle *RaffleFilterer) FilterSignerUpdated(opts *bind.FilterOpts) (*RaffleSignerUpdatedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return &RaffleSignerUpdatedIterator{contract: _Raffle.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_Raffle *RaffleFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *RaffleSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleSignerUpdated)
				if err := _Raffle.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_Raffle *RaffleFilterer) ParseSignerUpdated(log types.Log) (*RaffleSignerUpdated, error) {
	event := new(RaffleSignerUpdated)
	if err := _Raffle.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Raffle contract.
type RaffleUnpausedIterator struct {
	Event *RaffleUnpaused // Event containing the contract specifics and raw log

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
func (it *RaffleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleUnpaused)
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
		it.Event = new(RaffleUnpaused)
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
func (it *RaffleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleUnpaused represents a Unpaused event raised by the Raffle contract.
type RaffleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Raffle *RaffleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RaffleUnpausedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RaffleUnpausedIterator{contract: _Raffle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Raffle *RaffleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RaffleUnpaused) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleUnpaused)
				if err := _Raffle.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Raffle *RaffleFilterer) ParseUnpaused(log types.Log) (*RaffleUnpaused, error) {
	event := new(RaffleUnpaused)
	if err := _Raffle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleVerifierUpdatedIterator is returned from FilterVerifierUpdated and is used to iterate over the raw logs and unpacked data for VerifierUpdated events raised by the Raffle contract.
type RaffleVerifierUpdatedIterator struct {
	Event *RaffleVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *RaffleVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleVerifierUpdated)
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
		it.Event = new(RaffleVerifierUpdated)
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
func (it *RaffleVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleVerifierUpdated represents a VerifierUpdated event raised by the Raffle contract.
type RaffleVerifierUpdated struct {
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierUpdated is a free log retrieval operation binding the contract event 0xd24015cc99cc1700cafca3042840a1d8ac1e3964fd2e0e37ea29c654056ee327.
//
// Solidity: event VerifierUpdated(address verifier)
func (_Raffle *RaffleFilterer) FilterVerifierUpdated(opts *bind.FilterOpts) (*RaffleVerifierUpdatedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &RaffleVerifierUpdatedIterator{contract: _Raffle.contract, event: "VerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchVerifierUpdated is a free log subscription operation binding the contract event 0xd24015cc99cc1700cafca3042840a1d8ac1e3964fd2e0e37ea29c654056ee327.
//
// Solidity: event VerifierUpdated(address verifier)
func (_Raffle *RaffleFilterer) WatchVerifierUpdated(opts *bind.WatchOpts, sink chan<- *RaffleVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleVerifierUpdated)
				if err := _Raffle.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
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

// ParseVerifierUpdated is a log parse operation binding the contract event 0xd24015cc99cc1700cafca3042840a1d8ac1e3964fd2e0e37ea29c654056ee327.
//
// Solidity: event VerifierUpdated(address verifier)
func (_Raffle *RaffleFilterer) ParseVerifierUpdated(log types.Log) (*RaffleVerifierUpdated, error) {
	event := new(RaffleVerifierUpdated)
	if err := _Raffle.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleVkeyUpdatedIterator is returned from FilterVkeyUpdated and is used to iterate over the raw logs and unpacked data for VkeyUpdated events raised by the Raffle contract.
type RaffleVkeyUpdatedIterator struct {
	Event *RaffleVkeyUpdated // Event containing the contract specifics and raw log

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
func (it *RaffleVkeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleVkeyUpdated)
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
		it.Event = new(RaffleVkeyUpdated)
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
func (it *RaffleVkeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleVkeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleVkeyUpdated represents a VkeyUpdated event raised by the Raffle contract.
type RaffleVkeyUpdated struct {
	Vkey [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterVkeyUpdated is a free log retrieval operation binding the contract event 0x127f2989a02817cf91f825d7e0fbbf2799274632f070f40be967d57c923acbda.
//
// Solidity: event VkeyUpdated(bytes32 vkey)
func (_Raffle *RaffleFilterer) FilterVkeyUpdated(opts *bind.FilterOpts) (*RaffleVkeyUpdatedIterator, error) {

	logs, sub, err := _Raffle.contract.FilterLogs(opts, "VkeyUpdated")
	if err != nil {
		return nil, err
	}
	return &RaffleVkeyUpdatedIterator{contract: _Raffle.contract, event: "VkeyUpdated", logs: logs, sub: sub}, nil
}

// WatchVkeyUpdated is a free log subscription operation binding the contract event 0x127f2989a02817cf91f825d7e0fbbf2799274632f070f40be967d57c923acbda.
//
// Solidity: event VkeyUpdated(bytes32 vkey)
func (_Raffle *RaffleFilterer) WatchVkeyUpdated(opts *bind.WatchOpts, sink chan<- *RaffleVkeyUpdated) (event.Subscription, error) {

	logs, sub, err := _Raffle.contract.WatchLogs(opts, "VkeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleVkeyUpdated)
				if err := _Raffle.contract.UnpackLog(event, "VkeyUpdated", log); err != nil {
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

// ParseVkeyUpdated is a log parse operation binding the contract event 0x127f2989a02817cf91f825d7e0fbbf2799274632f070f40be967d57c923acbda.
//
// Solidity: event VkeyUpdated(bytes32 vkey)
func (_Raffle *RaffleFilterer) ParseVkeyUpdated(log types.Log) (*RaffleVkeyUpdated, error) {
	event := new(RaffleVkeyUpdated)
	if err := _Raffle.contract.UnpackLog(event, "VkeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

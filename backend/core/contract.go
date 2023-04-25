// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dataName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_identificationCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_author\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_projectName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dataVolume\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dataTraceability\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_creationOrganization\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dataFileHash\",\"type\":\"string\"}],\"name\":\"createExperimentVoucher\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"experimentId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"ExperimentCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"experimentId\",\"type\":\"bytes32\"}],\"name\":\"getExperiment\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"dataName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identificationCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataVolume\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataTraceability\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"creationOrganization\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataFileHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastExperimentId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetExperiment is a free data retrieval call binding the contract method 0x60bfc1f4.
//
// Solidity: function getExperiment(bytes32 experimentId) view returns(string dataName, string identificationCode, string author, string projectName, string dataVolume, string version, string dataTraceability, string creationOrganization, string dataFileHash, uint256 timestamp)
func (_Contract *ContractCaller) GetExperiment(opts *bind.CallOpts, experimentId [32]byte) (struct {
	DataName             string
	IdentificationCode   string
	Author               string
	ProjectName          string
	DataVolume           string
	Version              string
	DataTraceability     string
	CreationOrganization string
	DataFileHash         string
	Timestamp            *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getExperiment", experimentId)

	outstruct := new(struct {
		DataName             string
		IdentificationCode   string
		Author               string
		ProjectName          string
		DataVolume           string
		Version              string
		DataTraceability     string
		CreationOrganization string
		DataFileHash         string
		Timestamp            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.IdentificationCode = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Author = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ProjectName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.DataVolume = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.DataTraceability = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.CreationOrganization = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.DataFileHash = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetExperiment is a free data retrieval call binding the contract method 0x60bfc1f4.
//
// Solidity: function getExperiment(bytes32 experimentId) view returns(string dataName, string identificationCode, string author, string projectName, string dataVolume, string version, string dataTraceability, string creationOrganization, string dataFileHash, uint256 timestamp)
func (_Contract *ContractSession) GetExperiment(experimentId [32]byte) (struct {
	DataName             string
	IdentificationCode   string
	Author               string
	ProjectName          string
	DataVolume           string
	Version              string
	DataTraceability     string
	CreationOrganization string
	DataFileHash         string
	Timestamp            *big.Int
}, error) {
	return _Contract.Contract.GetExperiment(&_Contract.CallOpts, experimentId)
}

// GetExperiment is a free data retrieval call binding the contract method 0x60bfc1f4.
//
// Solidity: function getExperiment(bytes32 experimentId) view returns(string dataName, string identificationCode, string author, string projectName, string dataVolume, string version, string dataTraceability, string creationOrganization, string dataFileHash, uint256 timestamp)
func (_Contract *ContractCallerSession) GetExperiment(experimentId [32]byte) (struct {
	DataName             string
	IdentificationCode   string
	Author               string
	ProjectName          string
	DataVolume           string
	Version              string
	DataTraceability     string
	CreationOrganization string
	DataFileHash         string
	Timestamp            *big.Int
}, error) {
	return _Contract.Contract.GetExperiment(&_Contract.CallOpts, experimentId)
}

// LastExperimentId is a free data retrieval call binding the contract method 0xe90a474c.
//
// Solidity: function lastExperimentId() view returns(bytes32)
func (_Contract *ContractCaller) LastExperimentId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastExperimentId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastExperimentId is a free data retrieval call binding the contract method 0xe90a474c.
//
// Solidity: function lastExperimentId() view returns(bytes32)
func (_Contract *ContractSession) LastExperimentId() ([32]byte, error) {
	return _Contract.Contract.LastExperimentId(&_Contract.CallOpts)
}

// LastExperimentId is a free data retrieval call binding the contract method 0xe90a474c.
//
// Solidity: function lastExperimentId() view returns(bytes32)
func (_Contract *ContractCallerSession) LastExperimentId() ([32]byte, error) {
	return _Contract.Contract.LastExperimentId(&_Contract.CallOpts)
}

// CreateExperimentVoucher is a paid mutator transaction binding the contract method 0x9c8918d2.
//
// Solidity: function createExperimentVoucher(string _dataName, string _identificationCode, string _author, string _projectName, string _dataVolume, string _version, string _dataTraceability, string _creationOrganization, string _dataFileHash) returns(bytes32)
func (_Contract *ContractTransactor) CreateExperimentVoucher(opts *bind.TransactOpts, _dataName string, _identificationCode string, _author string, _projectName string, _dataVolume string, _version string, _dataTraceability string, _creationOrganization string, _dataFileHash string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createExperimentVoucher", _dataName, _identificationCode, _author, _projectName, _dataVolume, _version, _dataTraceability, _creationOrganization, _dataFileHash)
}

// CreateExperimentVoucher is a paid mutator transaction binding the contract method 0x9c8918d2.
//
// Solidity: function createExperimentVoucher(string _dataName, string _identificationCode, string _author, string _projectName, string _dataVolume, string _version, string _dataTraceability, string _creationOrganization, string _dataFileHash) returns(bytes32)
func (_Contract *ContractSession) CreateExperimentVoucher(_dataName string, _identificationCode string, _author string, _projectName string, _dataVolume string, _version string, _dataTraceability string, _creationOrganization string, _dataFileHash string) (*types.Transaction, error) {
	return _Contract.Contract.CreateExperimentVoucher(&_Contract.TransactOpts, _dataName, _identificationCode, _author, _projectName, _dataVolume, _version, _dataTraceability, _creationOrganization, _dataFileHash)
}

// CreateExperimentVoucher is a paid mutator transaction binding the contract method 0x9c8918d2.
//
// Solidity: function createExperimentVoucher(string _dataName, string _identificationCode, string _author, string _projectName, string _dataVolume, string _version, string _dataTraceability, string _creationOrganization, string _dataFileHash) returns(bytes32)
func (_Contract *ContractTransactorSession) CreateExperimentVoucher(_dataName string, _identificationCode string, _author string, _projectName string, _dataVolume string, _version string, _dataTraceability string, _creationOrganization string, _dataFileHash string) (*types.Transaction, error) {
	return _Contract.Contract.CreateExperimentVoucher(&_Contract.TransactOpts, _dataName, _identificationCode, _author, _projectName, _dataVolume, _version, _dataTraceability, _creationOrganization, _dataFileHash)
}

// ContractExperimentCreatedIterator is returned from FilterExperimentCreated and is used to iterate over the raw logs and unpacked data for ExperimentCreated events raised by the Contract contract.
type ContractExperimentCreatedIterator struct {
	Event *ContractExperimentCreated // Event containing the contract specifics and raw log

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
func (it *ContractExperimentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExperimentCreated)
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
		it.Event = new(ContractExperimentCreated)
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
func (it *ContractExperimentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExperimentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExperimentCreated represents a ExperimentCreated event raised by the Contract contract.
type ContractExperimentCreated struct {
	ExperimentId    [32]byte
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExperimentCreated is a free log retrieval operation binding the contract event 0xd76d7eb8be06bc1dec5eb602e0b7eb0cb167bef0c690cd38ae9b543dfed1340a.
//
// Solidity: event ExperimentCreated(bytes32 indexed experimentId, bytes32 indexed transactionHash)
func (_Contract *ContractFilterer) FilterExperimentCreated(opts *bind.FilterOpts, experimentId [][32]byte, transactionHash [][32]byte) (*ContractExperimentCreatedIterator, error) {

	var experimentIdRule []interface{}
	for _, experimentIdItem := range experimentId {
		experimentIdRule = append(experimentIdRule, experimentIdItem)
	}
	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExperimentCreated", experimentIdRule, transactionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractExperimentCreatedIterator{contract: _Contract.contract, event: "ExperimentCreated", logs: logs, sub: sub}, nil
}

// WatchExperimentCreated is a free log subscription operation binding the contract event 0xd76d7eb8be06bc1dec5eb602e0b7eb0cb167bef0c690cd38ae9b543dfed1340a.
//
// Solidity: event ExperimentCreated(bytes32 indexed experimentId, bytes32 indexed transactionHash)
func (_Contract *ContractFilterer) WatchExperimentCreated(opts *bind.WatchOpts, sink chan<- *ContractExperimentCreated, experimentId [][32]byte, transactionHash [][32]byte) (event.Subscription, error) {

	var experimentIdRule []interface{}
	for _, experimentIdItem := range experimentId {
		experimentIdRule = append(experimentIdRule, experimentIdItem)
	}
	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExperimentCreated", experimentIdRule, transactionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExperimentCreated)
				if err := _Contract.contract.UnpackLog(event, "ExperimentCreated", log); err != nil {
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

// ParseExperimentCreated is a log parse operation binding the contract event 0xd76d7eb8be06bc1dec5eb602e0b7eb0cb167bef0c690cd38ae9b543dfed1340a.
//
// Solidity: event ExperimentCreated(bytes32 indexed experimentId, bytes32 indexed transactionHash)
func (_Contract *ContractFilterer) ParseExperimentCreated(log types.Log) (*ContractExperimentCreated, error) {
	event := new(ContractExperimentCreated)
	if err := _Contract.contract.UnpackLog(event, "ExperimentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

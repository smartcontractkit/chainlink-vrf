package testbeaconvrfconsumer

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

type VRFBeaconTypesOutputServed struct {
	Height            uint64
	ConfirmationDelay *big.Int
	ProofG1X          *big.Int
	ProofG1Y          *big.Int
}

var ArbSysMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"uniqueId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexInBatch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arbBlockNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethBlockNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"L2ToL1Transaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hash\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arbBlockNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethBlockNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"callvalue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"L2ToL1Tx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reserved\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"SendMerkleUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"arbBlockNum\",\"type\":\"uint256\"}],\"name\":\"arbBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbOSVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStorageGasAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTopLevelCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unused\",\"type\":\"address\"}],\"name\":\"mapL1SenderContractAddressToL2Alias\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myCallersAddressWithoutAliasing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendMerkleTreeState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"partials\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendTxToL1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wasMyCallersAddressAliased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

var ArbSysABI = ArbSysMetaData.ABI

type ArbSys struct {
	ArbSysCaller
	ArbSysTransactor
	ArbSysFilterer
}

type ArbSysCaller struct {
	contract *bind.BoundContract
}

type ArbSysTransactor struct {
	contract *bind.BoundContract
}

type ArbSysFilterer struct {
	contract *bind.BoundContract
}

type ArbSysSession struct {
	Contract     *ArbSys
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbSysCallerSession struct {
	Contract *ArbSysCaller
	CallOpts bind.CallOpts
}

type ArbSysTransactorSession struct {
	Contract     *ArbSysTransactor
	TransactOpts bind.TransactOpts
}

type ArbSysRaw struct {
	Contract *ArbSys
}

type ArbSysCallerRaw struct {
	Contract *ArbSysCaller
}

type ArbSysTransactorRaw struct {
	Contract *ArbSysTransactor
}

func NewArbSys(address common.Address, backend bind.ContractBackend) (*ArbSys, error) {
	contract, err := bindArbSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbSys{ArbSysCaller: ArbSysCaller{contract: contract}, ArbSysTransactor: ArbSysTransactor{contract: contract}, ArbSysFilterer: ArbSysFilterer{contract: contract}}, nil
}

func NewArbSysCaller(address common.Address, caller bind.ContractCaller) (*ArbSysCaller, error) {
	contract, err := bindArbSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysCaller{contract: contract}, nil
}

func NewArbSysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbSysTransactor, error) {
	contract, err := bindArbSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbSysTransactor{contract: contract}, nil
}

func NewArbSysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbSysFilterer, error) {
	contract, err := bindArbSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbSysFilterer{contract: contract}, nil
}

func bindArbSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbSysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ArbSys *ArbSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.ArbSysCaller.contract.Call(opts, result, method, params...)
}

func (_ArbSys *ArbSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transfer(opts)
}

func (_ArbSys *ArbSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.ArbSysTransactor.contract.Transact(opts, method, params...)
}

func (_ArbSys *ArbSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbSys.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbSys *ArbSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transfer(opts)
}

func (_ArbSys *ArbSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbSys.Contract.contract.Transact(opts, method, params...)
}

func (_ArbSys *ArbSysCaller) ArbBlockHash(opts *bind.CallOpts, arbBlockNum *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "arbBlockHash", arbBlockNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_ArbSys *ArbSysSession) ArbBlockHash(arbBlockNum *big.Int) ([32]byte, error) {
	return _ArbSys.Contract.ArbBlockHash(&_ArbSys.CallOpts, arbBlockNum)
}

func (_ArbSys *ArbSysCallerSession) ArbBlockHash(arbBlockNum *big.Int) ([32]byte, error) {
	return _ArbSys.Contract.ArbBlockHash(&_ArbSys.CallOpts, arbBlockNum)
}

func (_ArbSys *ArbSysCaller) ArbBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "arbBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbSys *ArbSysSession) ArbBlockNumber() (*big.Int, error) {
	return _ArbSys.Contract.ArbBlockNumber(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) ArbBlockNumber() (*big.Int, error) {
	return _ArbSys.Contract.ArbBlockNumber(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCaller) ArbChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "arbChainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbSys *ArbSysSession) ArbChainID() (*big.Int, error) {
	return _ArbSys.Contract.ArbChainID(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) ArbChainID() (*big.Int, error) {
	return _ArbSys.Contract.ArbChainID(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCaller) ArbOSVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "arbOSVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbSys *ArbSysSession) ArbOSVersion() (*big.Int, error) {
	return _ArbSys.Contract.ArbOSVersion(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) ArbOSVersion() (*big.Int, error) {
	return _ArbSys.Contract.ArbOSVersion(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCaller) GetStorageGasAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "getStorageGasAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbSys *ArbSysSession) GetStorageGasAvailable() (*big.Int, error) {
	return _ArbSys.Contract.GetStorageGasAvailable(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) GetStorageGasAvailable() (*big.Int, error) {
	return _ArbSys.Contract.GetStorageGasAvailable(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCaller) IsTopLevelCall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "isTopLevelCall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ArbSys *ArbSysSession) IsTopLevelCall() (bool, error) {
	return _ArbSys.Contract.IsTopLevelCall(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) IsTopLevelCall() (bool, error) {
	return _ArbSys.Contract.IsTopLevelCall(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCaller) MapL1SenderContractAddressToL2Alias(opts *bind.CallOpts, sender common.Address, unused common.Address) (common.Address, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "mapL1SenderContractAddressToL2Alias", sender, unused)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbSys *ArbSysSession) MapL1SenderContractAddressToL2Alias(sender common.Address, unused common.Address) (common.Address, error) {
	return _ArbSys.Contract.MapL1SenderContractAddressToL2Alias(&_ArbSys.CallOpts, sender, unused)
}

func (_ArbSys *ArbSysCallerSession) MapL1SenderContractAddressToL2Alias(sender common.Address, unused common.Address) (common.Address, error) {
	return _ArbSys.Contract.MapL1SenderContractAddressToL2Alias(&_ArbSys.CallOpts, sender, unused)
}

func (_ArbSys *ArbSysCaller) MyCallersAddressWithoutAliasing(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "myCallersAddressWithoutAliasing")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbSys *ArbSysSession) MyCallersAddressWithoutAliasing() (common.Address, error) {
	return _ArbSys.Contract.MyCallersAddressWithoutAliasing(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) MyCallersAddressWithoutAliasing() (common.Address, error) {
	return _ArbSys.Contract.MyCallersAddressWithoutAliasing(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCaller) SendMerkleTreeState(opts *bind.CallOpts) (struct {
	Size     *big.Int
	Root     [32]byte
	Partials [][32]byte
}, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "sendMerkleTreeState")

	outstruct := new(struct {
		Size     *big.Int
		Root     [32]byte
		Partials [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Size = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Root = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Partials = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

func (_ArbSys *ArbSysSession) SendMerkleTreeState() (struct {
	Size     *big.Int
	Root     [32]byte
	Partials [][32]byte
}, error) {
	return _ArbSys.Contract.SendMerkleTreeState(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) SendMerkleTreeState() (struct {
	Size     *big.Int
	Root     [32]byte
	Partials [][32]byte
}, error) {
	return _ArbSys.Contract.SendMerkleTreeState(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCaller) WasMyCallersAddressAliased(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ArbSys.contract.Call(opts, &out, "wasMyCallersAddressAliased")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ArbSys *ArbSysSession) WasMyCallersAddressAliased() (bool, error) {
	return _ArbSys.Contract.WasMyCallersAddressAliased(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysCallerSession) WasMyCallersAddressAliased() (bool, error) {
	return _ArbSys.Contract.WasMyCallersAddressAliased(&_ArbSys.CallOpts)
}

func (_ArbSys *ArbSysTransactor) SendTxToL1(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "sendTxToL1", destination, data)
}

func (_ArbSys *ArbSysSession) SendTxToL1(destination common.Address, data []byte) (*types.Transaction, error) {
	return _ArbSys.Contract.SendTxToL1(&_ArbSys.TransactOpts, destination, data)
}

func (_ArbSys *ArbSysTransactorSession) SendTxToL1(destination common.Address, data []byte) (*types.Transaction, error) {
	return _ArbSys.Contract.SendTxToL1(&_ArbSys.TransactOpts, destination, data)
}

func (_ArbSys *ArbSysTransactor) WithdrawEth(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _ArbSys.contract.Transact(opts, "withdrawEth", destination)
}

func (_ArbSys *ArbSysSession) WithdrawEth(destination common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawEth(&_ArbSys.TransactOpts, destination)
}

func (_ArbSys *ArbSysTransactorSession) WithdrawEth(destination common.Address) (*types.Transaction, error) {
	return _ArbSys.Contract.WithdrawEth(&_ArbSys.TransactOpts, destination)
}

type ArbSysL2ToL1TransactionIterator struct {
	Event *ArbSysL2ToL1Transaction

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbSysL2ToL1TransactionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSysL2ToL1Transaction)
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

	select {
	case log := <-it.logs:
		it.Event = new(ArbSysL2ToL1Transaction)
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

func (it *ArbSysL2ToL1TransactionIterator) Error() error {
	return it.fail
}

func (it *ArbSysL2ToL1TransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbSysL2ToL1Transaction struct {
	Caller       common.Address
	Destination  common.Address
	UniqueId     *big.Int
	BatchNumber  *big.Int
	IndexInBatch *big.Int
	ArbBlockNum  *big.Int
	EthBlockNum  *big.Int
	Timestamp    *big.Int
	Callvalue    *big.Int
	Data         []byte
	Raw          types.Log
}

func (_ArbSys *ArbSysFilterer) FilterL2ToL1Transaction(opts *bind.FilterOpts, destination []common.Address, uniqueId []*big.Int, batchNumber []*big.Int) (*ArbSysL2ToL1TransactionIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var uniqueIdRule []interface{}
	for _, uniqueIdItem := range uniqueId {
		uniqueIdRule = append(uniqueIdRule, uniqueIdItem)
	}
	var batchNumberRule []interface{}
	for _, batchNumberItem := range batchNumber {
		batchNumberRule = append(batchNumberRule, batchNumberItem)
	}

	logs, sub, err := _ArbSys.contract.FilterLogs(opts, "L2ToL1Transaction", destinationRule, uniqueIdRule, batchNumberRule)
	if err != nil {
		return nil, err
	}
	return &ArbSysL2ToL1TransactionIterator{contract: _ArbSys.contract, event: "L2ToL1Transaction", logs: logs, sub: sub}, nil
}

func (_ArbSys *ArbSysFilterer) WatchL2ToL1Transaction(opts *bind.WatchOpts, sink chan<- *ArbSysL2ToL1Transaction, destination []common.Address, uniqueId []*big.Int, batchNumber []*big.Int) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var uniqueIdRule []interface{}
	for _, uniqueIdItem := range uniqueId {
		uniqueIdRule = append(uniqueIdRule, uniqueIdItem)
	}
	var batchNumberRule []interface{}
	for _, batchNumberItem := range batchNumber {
		batchNumberRule = append(batchNumberRule, batchNumberItem)
	}

	logs, sub, err := _ArbSys.contract.WatchLogs(opts, "L2ToL1Transaction", destinationRule, uniqueIdRule, batchNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbSysL2ToL1Transaction)
				if err := _ArbSys.contract.UnpackLog(event, "L2ToL1Transaction", log); err != nil {
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

func (_ArbSys *ArbSysFilterer) ParseL2ToL1Transaction(log types.Log) (*ArbSysL2ToL1Transaction, error) {
	event := new(ArbSysL2ToL1Transaction)
	if err := _ArbSys.contract.UnpackLog(event, "L2ToL1Transaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbSysL2ToL1TxIterator struct {
	Event *ArbSysL2ToL1Tx

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbSysL2ToL1TxIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSysL2ToL1Tx)
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

	select {
	case log := <-it.logs:
		it.Event = new(ArbSysL2ToL1Tx)
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

func (it *ArbSysL2ToL1TxIterator) Error() error {
	return it.fail
}

func (it *ArbSysL2ToL1TxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbSysL2ToL1Tx struct {
	Caller      common.Address
	Destination common.Address
	Hash        *big.Int
	Position    *big.Int
	ArbBlockNum *big.Int
	EthBlockNum *big.Int
	Timestamp   *big.Int
	Callvalue   *big.Int
	Data        []byte
	Raw         types.Log
}

func (_ArbSys *ArbSysFilterer) FilterL2ToL1Tx(opts *bind.FilterOpts, destination []common.Address, hash []*big.Int, position []*big.Int) (*ArbSysL2ToL1TxIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _ArbSys.contract.FilterLogs(opts, "L2ToL1Tx", destinationRule, hashRule, positionRule)
	if err != nil {
		return nil, err
	}
	return &ArbSysL2ToL1TxIterator{contract: _ArbSys.contract, event: "L2ToL1Tx", logs: logs, sub: sub}, nil
}

func (_ArbSys *ArbSysFilterer) WatchL2ToL1Tx(opts *bind.WatchOpts, sink chan<- *ArbSysL2ToL1Tx, destination []common.Address, hash []*big.Int, position []*big.Int) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _ArbSys.contract.WatchLogs(opts, "L2ToL1Tx", destinationRule, hashRule, positionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbSysL2ToL1Tx)
				if err := _ArbSys.contract.UnpackLog(event, "L2ToL1Tx", log); err != nil {
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

func (_ArbSys *ArbSysFilterer) ParseL2ToL1Tx(log types.Log) (*ArbSysL2ToL1Tx, error) {
	event := new(ArbSysL2ToL1Tx)
	if err := _ArbSys.contract.UnpackLog(event, "L2ToL1Tx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbSysSendMerkleUpdateIterator struct {
	Event *ArbSysSendMerkleUpdate

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbSysSendMerkleUpdateIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSysSendMerkleUpdate)
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

	select {
	case log := <-it.logs:
		it.Event = new(ArbSysSendMerkleUpdate)
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

func (it *ArbSysSendMerkleUpdateIterator) Error() error {
	return it.fail
}

func (it *ArbSysSendMerkleUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbSysSendMerkleUpdate struct {
	Reserved *big.Int
	Hash     [32]byte
	Position *big.Int
	Raw      types.Log
}

func (_ArbSys *ArbSysFilterer) FilterSendMerkleUpdate(opts *bind.FilterOpts, reserved []*big.Int, hash [][32]byte, position []*big.Int) (*ArbSysSendMerkleUpdateIterator, error) {

	var reservedRule []interface{}
	for _, reservedItem := range reserved {
		reservedRule = append(reservedRule, reservedItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _ArbSys.contract.FilterLogs(opts, "SendMerkleUpdate", reservedRule, hashRule, positionRule)
	if err != nil {
		return nil, err
	}
	return &ArbSysSendMerkleUpdateIterator{contract: _ArbSys.contract, event: "SendMerkleUpdate", logs: logs, sub: sub}, nil
}

func (_ArbSys *ArbSysFilterer) WatchSendMerkleUpdate(opts *bind.WatchOpts, sink chan<- *ArbSysSendMerkleUpdate, reserved []*big.Int, hash [][32]byte, position []*big.Int) (event.Subscription, error) {

	var reservedRule []interface{}
	for _, reservedItem := range reserved {
		reservedRule = append(reservedRule, reservedItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var positionRule []interface{}
	for _, positionItem := range position {
		positionRule = append(positionRule, positionItem)
	}

	logs, sub, err := _ArbSys.contract.WatchLogs(opts, "SendMerkleUpdate", reservedRule, hashRule, positionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbSysSendMerkleUpdate)
				if err := _ArbSys.contract.UnpackLog(event, "SendMerkleUpdate", log); err != nil {
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

func (_ArbSys *ArbSysFilterer) ParseSendMerkleUpdate(log types.Log) (*ArbSysSendMerkleUpdate, error) {
	event := new(ArbSysSendMerkleUpdate)
	if err := _ArbSys.contract.UnpackLog(event, "SendMerkleUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var BeaconVRFConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"shouldFail\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fail\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_ReceivedRandomnessByRequestID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_arguments\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_gasAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"name\":\"s_myBeaconRequests\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.SlotNumber\",\"name\":\"slotNumber\",\"type\":\"uint32\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_randomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"s_requestsIDs\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_subId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"shouldFail\",\"type\":\"bool\"}],\"name\":\"setFail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"reqId\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"delay\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"}],\"name\":\"storeBeaconRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"testRedeemRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"testRequestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"testRequestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161129638038061129683398101604081905261002f9161006c565b6001600160a01b03929092166080819052600580546001600160a01b03191690911790556008805460ff19169115159190911790556009556100c0565b60008060006060848603121561008157600080fd5b83516001600160a01b038116811461009857600080fd5b602085015190935080151581146100ae57600080fd5b80925050604084015190509250925092565b6080516111bb6100db600039600061034801526111bb6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063a9cc471811610097578063ea7502ab11610066578063ea7502ab14610304578063f08c5daa14610317578063f339c79414610320578063f6eaffc81461033357600080fd5b8063a9cc4718146102aa578063c6d61301146102c7578063cd0593df146102da578063d4fab62c146102f157600080fd5b80636d162a3e116100d35780636d162a3e1461021c578063706da1ca1461022f5780637716cdaa146102745780639d7694021461028957600080fd5b806319a5fa22146101055780632f7527cc146101a35780635a47dd71146101bd5780635f15cccc146101d2575b600080fd5b610161610113366004610a38565b60026020526000908152604090205463ffffffff811690640100000000810462ffffff1690670100000000000000810461ffff1690690100000000000000000090046001600160a01b031684565b6040805163ffffffff909516855262ffffff909316602085015261ffff909116918301919091526001600160a01b031660608201526080015b60405180910390f35b6101ab600881565b60405160ff909116815260200161019a565b6101d06101cb366004610b37565b610346565b005b6102056101e0366004610c21565b600160209081526000928352604080842090915290825290205465ffffffffffff1681565b60405165ffffffffffff909116815260200161019a565b6101d061022a366004610c5f565b6103d3565b60065461025b9074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161019a565b61027c610506565b60405161019a9190610cfb565b6101d0610297366004610d0e565b6008805460ff1916911515919091179055565b6008546102b79060ff1681565b604051901515815260200161019a565b6102056102d5366004610d30565b610594565b6102e360095481565b60405190815260200161019a565b6101d06102ff366004610d6c565b61068c565b610205610312366004610d9c565b61075f565b6102e360075481565b6102e361032e366004610e20565b61085d565b6102e3610341366004610e4c565b61088e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633146103c35760405162461bcd60e51b815260206004820152601760248201527f6f6e6c7920726f757465722063616e2066756c66696c6c00000000000000000060448201526064015b60405180910390fd5b6103ce8383836108af565b505050565b600083815260016020908152604080832062ffffff861684529091528120805465ffffffffffff191665ffffffffffff87161790556009546104159085610e91565b6040805160808101825263ffffffff928316815262ffffff958616602080830191825261ffff968716838501908152306060850190815265ffffffffffff909b1660009081526002909252939020915182549151935199516001600160a01b03166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff9a90971667010000000000000002999099167fffffff00000000000000000000000000000000000000000000ffffffffffffff939097166401000000000266ffffffffffffff199091169890931697909717919091171692909217179092555050565b6004805461051390610ea5565b80601f016020809104026020016040519081016040528092919081815260200182805461053f90610ea5565b801561058c5780601f106105615761010080835404028352916020019161058c565b820191906000526020600020905b81548152906001019060200180831161056f57829003601f168201915b505050505081565b6000806009546105a2610937565b6105ac9190610edf565b90506000816009546105bc610937565b6105c69190610ef3565b6105d09190610f0b565b60055460408051602081018252600080825291517f4ffac83a00000000000000000000000000000000000000000000000000000000815293945090926001600160a01b0390921691634ffac83a91610631918a918c918b9190600401610f22565b6020604051808303816000875af1158015610650573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106749190610f50565b90506106828183878a6103d3565b9695505050505050565b60055460408051602081018252600080825291517fa8a254ed00000000000000000000000000000000000000000000000000000000815291926001600160a01b03169163a8a254ed916106e59187918791600401610f6d565b6000604051808303816000875af1158015610704573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261072c9190810190610f9d565b65ffffffffffff8316600090815260036020908152604090912082519293506107599290918401906109c1565b50505050565b60008060095461076d610937565b6107779190610edf565b9050600081600954610787610937565b6107919190610ef3565b61079b9190610f0b565b60055460408051602081018252600080825291517fdb972c8b00000000000000000000000000000000000000000000000000000000815293945090926001600160a01b039092169163db972c8b91610800918d918d918d918d918d919060040161102e565b6020604051808303816000875af115801561081f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108439190610f50565b90506108518183898b6103d3565b98975050505050505050565b6003602052816000526040600020818154811061087957600080fd5b90600052602060002001600091509150505481565b6000818154811061089e57600080fd5b600091825260209091200154905081565b60085460ff16156109025760405162461bcd60e51b815260206004820152601d60248201527f206661696c656420696e2066756c66696c6c52616e646f6d576f72647300000060448201526064016103ba565b65ffffffffffff83166000908152600360209081526040909120835161092a928501906109c1565b50600461075982826110d5565b60004661a4b181148061094c575062066eed81145b156109ba5760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610990573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b49190611195565b91505090565b4391505090565b8280548282559060005260206000209081019282156109fc579160200282015b828111156109fc5782518255916020019190600101906109e1565b50610a08929150610a0c565b5090565b5b80821115610a085760008155600101610a0d565b65ffffffffffff81168114610a3557600080fd5b50565b600060208284031215610a4a57600080fd5b8135610a5581610a21565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610a9b57610a9b610a5c565b604052919050565b600067ffffffffffffffff821115610abd57610abd610a5c565b5060051b60200190565b600082601f830112610ad857600080fd5b813567ffffffffffffffff811115610af257610af2610a5c565b610b05601f8201601f1916602001610a72565b818152846020838601011115610b1a57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215610b4c57600080fd5b8335610b5781610a21565b925060208481013567ffffffffffffffff80821115610b7557600080fd5b818701915087601f830112610b8957600080fd5b8135610b9c610b9782610aa3565b610a72565b81815260059190911b8301840190848101908a831115610bbb57600080fd5b938501935b82851015610bd957843582529385019390850190610bc0565b965050506040870135925080831115610bf157600080fd5b5050610bff86828701610ac7565b9150509250925092565b803562ffffff81168114610c1c57600080fd5b919050565b60008060408385031215610c3457600080fd5b82359150610c4460208401610c09565b90509250929050565b803561ffff81168114610c1c57600080fd5b60008060008060808587031215610c7557600080fd5b8435610c8081610a21565b935060208501359250610c9560408601610c09565b9150610ca360608601610c4d565b905092959194509250565b6000815180845260005b81811015610cd457602081850181015186830182015201610cb8565b81811115610ce6576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610a556020830184610cae565b600060208284031215610d2057600080fd5b81358015158114610a5557600080fd5b600080600060608486031215610d4557600080fd5b610d4e84610c4d565b925060208401359150610d6360408501610c09565b90509250925092565b60008060408385031215610d7f57600080fd5b823591506020830135610d9181610a21565b809150509250929050565b600080600080600060a08688031215610db457600080fd5b85359450610dc460208701610c4d565b9350610dd260408701610c09565b9250606086013563ffffffff81168114610deb57600080fd5b9150608086013567ffffffffffffffff811115610e0757600080fd5b610e1388828901610ac7565b9150509295509295909350565b60008060408385031215610e3357600080fd5b8235610e3e81610a21565b946020939093013593505050565b600060208284031215610e5e57600080fd5b5035919050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082610ea057610ea0610e65565b500490565b600181811c90821680610eb957607f821691505b602082108103610ed957634e487b7160e01b600052602260045260246000fd5b50919050565b600082610eee57610eee610e65565b500690565b60008219821115610f0657610f06610e7b565b500190565b600082821015610f1d57610f1d610e7b565b500390565b84815261ffff8416602082015262ffffff831660408201526080606082015260006106826080830184610cae565b600060208284031215610f6257600080fd5b8151610a5581610a21565b83815265ffffffffffff83166020820152606060408201526000610f946060830184610cae565b95945050505050565b60006020808385031215610fb057600080fd5b825167ffffffffffffffff811115610fc757600080fd5b8301601f81018513610fd857600080fd5b8051610fe6610b9782610aa3565b81815260059190911b8201830190838101908783111561100557600080fd5b928401925b828410156110235783518252928401929084019061100a565b979650505050505050565b86815261ffff8616602082015262ffffff8516604082015263ffffffff8416606082015260c06080820152600061106860c0830185610cae565b82810360a084015261107a8185610cae565b9998505050505050505050565b601f8211156103ce57600081815260208120601f850160051c810160208610156110ae5750805b601f850160051c820191505b818110156110cd578281556001016110ba565b505050505050565b815167ffffffffffffffff8111156110ef576110ef610a5c565b611103816110fd8454610ea5565b84611087565b602080601f83116001811461113857600084156111205750858301515b600019600386901b1c1916600185901b1785556110cd565b600085815260208120601f198616915b8281101561116757888601518255948401946001909101908401611148565b50858210156111855787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000602082840312156111a757600080fd5b505191905056fea164736f6c634300080f000a",
}

var BeaconVRFConsumerABI = BeaconVRFConsumerMetaData.ABI

var BeaconVRFConsumerBin = BeaconVRFConsumerMetaData.Bin

func DeployBeaconVRFConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address, shouldFail bool, beaconPeriodBlocks *big.Int) (common.Address, *types.Transaction, *BeaconVRFConsumer, error) {
	parsed, err := BeaconVRFConsumerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeaconVRFConsumerBin), backend, router, shouldFail, beaconPeriodBlocks)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeaconVRFConsumer{BeaconVRFConsumerCaller: BeaconVRFConsumerCaller{contract: contract}, BeaconVRFConsumerTransactor: BeaconVRFConsumerTransactor{contract: contract}, BeaconVRFConsumerFilterer: BeaconVRFConsumerFilterer{contract: contract}}, nil
}

type BeaconVRFConsumer struct {
	BeaconVRFConsumerCaller
	BeaconVRFConsumerTransactor
	BeaconVRFConsumerFilterer
}

type BeaconVRFConsumerCaller struct {
	contract *bind.BoundContract
}

type BeaconVRFConsumerTransactor struct {
	contract *bind.BoundContract
}

type BeaconVRFConsumerFilterer struct {
	contract *bind.BoundContract
}

type BeaconVRFConsumerSession struct {
	Contract     *BeaconVRFConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BeaconVRFConsumerCallerSession struct {
	Contract *BeaconVRFConsumerCaller
	CallOpts bind.CallOpts
}

type BeaconVRFConsumerTransactorSession struct {
	Contract     *BeaconVRFConsumerTransactor
	TransactOpts bind.TransactOpts
}

type BeaconVRFConsumerRaw struct {
	Contract *BeaconVRFConsumer
}

type BeaconVRFConsumerCallerRaw struct {
	Contract *BeaconVRFConsumerCaller
}

type BeaconVRFConsumerTransactorRaw struct {
	Contract *BeaconVRFConsumerTransactor
}

func NewBeaconVRFConsumer(address common.Address, backend bind.ContractBackend) (*BeaconVRFConsumer, error) {
	contract, err := bindBeaconVRFConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumer{BeaconVRFConsumerCaller: BeaconVRFConsumerCaller{contract: contract}, BeaconVRFConsumerTransactor: BeaconVRFConsumerTransactor{contract: contract}, BeaconVRFConsumerFilterer: BeaconVRFConsumerFilterer{contract: contract}}, nil
}

func NewBeaconVRFConsumerCaller(address common.Address, caller bind.ContractCaller) (*BeaconVRFConsumerCaller, error) {
	contract, err := bindBeaconVRFConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerCaller{contract: contract}, nil
}

func NewBeaconVRFConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconVRFConsumerTransactor, error) {
	contract, err := bindBeaconVRFConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerTransactor{contract: contract}, nil
}

func NewBeaconVRFConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconVRFConsumerFilterer, error) {
	contract, err := bindBeaconVRFConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerFilterer{contract: contract}, nil
}

func bindBeaconVRFConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BeaconVRFConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconVRFConsumer.Contract.BeaconVRFConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.BeaconVRFConsumerTransactor.contract.Transfer(opts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.BeaconVRFConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconVRFConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.contract.Transfer(opts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) NUMCONFDELAYS() (uint8, error) {
	return _BeaconVRFConsumer.Contract.NUMCONFDELAYS(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _BeaconVRFConsumer.Contract.NUMCONFDELAYS(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) Fail(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "fail")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) Fail() (bool, error) {
	return _BeaconVRFConsumer.Contract.Fail(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) Fail() (bool, error) {
	return _BeaconVRFConsumer.Contract.Fail(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.IBeaconPeriodBlocks(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.IBeaconPeriodBlocks(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SReceivedRandomnessByRequestID(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_ReceivedRandomnessByRequestID", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SReceivedRandomnessByRequestID(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SReceivedRandomnessByRequestID(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SReceivedRandomnessByRequestID(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SReceivedRandomnessByRequestID(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SArguments(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_arguments")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SArguments() ([]byte, error) {
	return _BeaconVRFConsumer.Contract.SArguments(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SArguments() ([]byte, error) {
	return _BeaconVRFConsumer.Contract.SArguments(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SGasAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_gasAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SGasAvailable() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SGasAvailable(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SGasAvailable() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SGasAvailable(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SMyBeaconRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_myBeaconRequests", arg0)

	outstruct := new(struct {
		SlotNumber        uint32
		ConfirmationDelay *big.Int
		NumWords          uint16
		Requester         common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlotNumber = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ConfirmationDelay = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NumWords = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Requester = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SMyBeaconRequests(arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	return _BeaconVRFConsumer.Contract.SMyBeaconRequests(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SMyBeaconRequests(arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	return _BeaconVRFConsumer.Contract.SMyBeaconRequests(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_randomWords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRandomWords(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRandomWords(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SRequestsIDs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_requestsIDs", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SRequestsIDs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRequestsIDs(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SRequestsIDs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRequestsIDs(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SSubId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_subId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SSubId() (uint64, error) {
	return _BeaconVRFConsumer.Contract.SSubId(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SSubId() (uint64, error) {
	return _BeaconVRFConsumer.Contract.SSubId(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.RawFulfillRandomWords(&_BeaconVRFConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.RawFulfillRandomWords(&_BeaconVRFConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) SetFail(opts *bind.TransactOpts, shouldFail bool) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "setFail", shouldFail)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SetFail(shouldFail bool) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.SetFail(&_BeaconVRFConsumer.TransactOpts, shouldFail)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) SetFail(shouldFail bool) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.SetFail(&_BeaconVRFConsumer.TransactOpts, shouldFail)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) StoreBeaconRequest(opts *bind.TransactOpts, reqId *big.Int, height *big.Int, delay *big.Int, numWords uint16) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "storeBeaconRequest", reqId, height, delay, numWords)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) StoreBeaconRequest(reqId *big.Int, height *big.Int, delay *big.Int, numWords uint16) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.StoreBeaconRequest(&_BeaconVRFConsumer.TransactOpts, reqId, height, delay, numWords)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) StoreBeaconRequest(reqId *big.Int, height *big.Int, delay *big.Int, numWords uint16) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.StoreBeaconRequest(&_BeaconVRFConsumer.TransactOpts, reqId, height, delay, numWords)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) TestRedeemRandomness(opts *bind.TransactOpts, subID *big.Int, requestID *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "testRedeemRandomness", subID, requestID)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) TestRedeemRandomness(subID *big.Int, requestID *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRedeemRandomness(&_BeaconVRFConsumer.TransactOpts, subID, requestID)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) TestRedeemRandomness(subID *big.Int, requestID *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRedeemRandomness(&_BeaconVRFConsumer.TransactOpts, subID, requestID)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) TestRequestRandomness(opts *bind.TransactOpts, numWords uint16, subID *big.Int, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "testRequestRandomness", numWords, subID, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) TestRequestRandomness(numWords uint16, subID *big.Int, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomness(&_BeaconVRFConsumer.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) TestRequestRandomness(numWords uint16, subID *big.Int, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomness(&_BeaconVRFConsumer.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) TestRequestRandomnessFulfillment(opts *bind.TransactOpts, subID *big.Int, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "testRequestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) TestRequestRandomnessFulfillment(subID *big.Int, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomnessFulfillment(&_BeaconVRFConsumer.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) TestRequestRandomnessFulfillment(subID *big.Int, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomnessFulfillment(&_BeaconVRFConsumer.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

type BeaconVRFConsumerConfigSetIterator struct {
	Event *BeaconVRFConsumerConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BeaconVRFConsumerConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconVRFConsumerConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(BeaconVRFConsumerConfigSet)
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

func (it *BeaconVRFConsumerConfigSetIterator) Error() error {
	return it.fail
}

func (it *BeaconVRFConsumerConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BeaconVRFConsumerConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) FilterConfigSet(opts *bind.FilterOpts) (*BeaconVRFConsumerConfigSetIterator, error) {

	logs, sub, err := _BeaconVRFConsumer.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerConfigSetIterator{contract: _BeaconVRFConsumer.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *BeaconVRFConsumerConfigSet) (event.Subscription, error) {

	logs, sub, err := _BeaconVRFConsumer.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BeaconVRFConsumerConfigSet)
				if err := _BeaconVRFConsumer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) ParseConfigSet(log types.Log) (*BeaconVRFConsumerConfigSet, error) {
	event := new(BeaconVRFConsumerConfigSet)
	if err := _BeaconVRFConsumer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BeaconVRFConsumerNewTransmissionIterator struct {
	Event *BeaconVRFConsumerNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BeaconVRFConsumerNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconVRFConsumerNewTransmission)
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

	select {
	case log := <-it.logs:
		it.Event = new(BeaconVRFConsumerNewTransmission)
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

func (it *BeaconVRFConsumerNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *BeaconVRFConsumerNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BeaconVRFConsumerNewTransmission struct {
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*BeaconVRFConsumerNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _BeaconVRFConsumer.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerNewTransmissionIterator{contract: _BeaconVRFConsumer.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *BeaconVRFConsumerNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _BeaconVRFConsumer.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BeaconVRFConsumerNewTransmission)
				if err := _BeaconVRFConsumer.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) ParseNewTransmission(log types.Log) (*BeaconVRFConsumerNewTransmission, error) {
	event := new(BeaconVRFConsumerNewTransmission)
	if err := _BeaconVRFConsumer.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BeaconVRFConsumerOutputsServedIterator struct {
	Event *BeaconVRFConsumerOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BeaconVRFConsumerOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconVRFConsumerOutputsServed)
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

	select {
	case log := <-it.logs:
		it.Event = new(BeaconVRFConsumerOutputsServed)
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

func (it *BeaconVRFConsumerOutputsServedIterator) Error() error {
	return it.fail
}

func (it *BeaconVRFConsumerOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BeaconVRFConsumerOutputsServed struct {
	RecentBlockHeight  uint64
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*BeaconVRFConsumerOutputsServedIterator, error) {

	logs, sub, err := _BeaconVRFConsumer.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerOutputsServedIterator{contract: _BeaconVRFConsumer.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *BeaconVRFConsumerOutputsServed) (event.Subscription, error) {

	logs, sub, err := _BeaconVRFConsumer.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BeaconVRFConsumerOutputsServed)
				if err := _BeaconVRFConsumer.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) ParseOutputsServed(log types.Log) (*BeaconVRFConsumerOutputsServed, error) {
	event := new(BeaconVRFConsumerOutputsServed)
	if err := _BeaconVRFConsumer.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BeaconVRFConsumerRandomWordsFulfilledIterator struct {
	Event *BeaconVRFConsumerRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BeaconVRFConsumerRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconVRFConsumerRandomWordsFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(BeaconVRFConsumerRandomWordsFulfilled)
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

func (it *BeaconVRFConsumerRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *BeaconVRFConsumerRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BeaconVRFConsumerRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*BeaconVRFConsumerRandomWordsFulfilledIterator, error) {

	logs, sub, err := _BeaconVRFConsumer.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerRandomWordsFulfilledIterator{contract: _BeaconVRFConsumer.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *BeaconVRFConsumerRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _BeaconVRFConsumer.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BeaconVRFConsumerRandomWordsFulfilled)
				if err := _BeaconVRFConsumer.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) ParseRandomWordsFulfilled(log types.Log) (*BeaconVRFConsumerRandomWordsFulfilled, error) {
	event := new(BeaconVRFConsumerRandomWordsFulfilled)
	if err := _BeaconVRFConsumer.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BeaconVRFConsumerRandomnessFulfillmentRequestedIterator struct {
	Event *BeaconVRFConsumerRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BeaconVRFConsumerRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconVRFConsumerRandomnessFulfillmentRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(BeaconVRFConsumerRandomnessFulfillmentRequested)
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

func (it *BeaconVRFConsumerRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *BeaconVRFConsumerRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BeaconVRFConsumerRandomnessFulfillmentRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	GasAllowance           uint32
	GasPrice               *big.Int
	WeiPerUnitLink         *big.Int
	Arguments              []byte
	Raw                    types.Log
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*BeaconVRFConsumerRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BeaconVRFConsumer.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerRandomnessFulfillmentRequestedIterator{contract: _BeaconVRFConsumer.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *BeaconVRFConsumerRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BeaconVRFConsumer.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BeaconVRFConsumerRandomnessFulfillmentRequested)
				if err := _BeaconVRFConsumer.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*BeaconVRFConsumerRandomnessFulfillmentRequested, error) {
	event := new(BeaconVRFConsumerRandomnessFulfillmentRequested)
	if err := _BeaconVRFConsumer.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BeaconVRFConsumerRandomnessRequestedIterator struct {
	Event *BeaconVRFConsumerRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BeaconVRFConsumerRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconVRFConsumerRandomnessRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(BeaconVRFConsumerRandomnessRequested)
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

func (it *BeaconVRFConsumerRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *BeaconVRFConsumerRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BeaconVRFConsumerRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	Raw                    types.Log
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*BeaconVRFConsumerRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BeaconVRFConsumer.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerRandomnessRequestedIterator{contract: _BeaconVRFConsumer.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *BeaconVRFConsumerRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BeaconVRFConsumer.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BeaconVRFConsumerRandomnessRequested)
				if err := _BeaconVRFConsumer.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_BeaconVRFConsumer *BeaconVRFConsumerFilterer) ParseRandomnessRequested(log types.Log) (*BeaconVRFConsumerRandomnessRequested, error) {
	event := new(BeaconVRFConsumerRandomnessRequested)
	if err := _BeaconVRFConsumer.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var ChainSpecificUtilMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x602d6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea164736f6c634300080f000a",
}

var ChainSpecificUtilABI = ChainSpecificUtilMetaData.ABI

var ChainSpecificUtilBin = ChainSpecificUtilMetaData.Bin

func DeployChainSpecificUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChainSpecificUtil, error) {
	parsed, err := ChainSpecificUtilMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChainSpecificUtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainSpecificUtil{ChainSpecificUtilCaller: ChainSpecificUtilCaller{contract: contract}, ChainSpecificUtilTransactor: ChainSpecificUtilTransactor{contract: contract}, ChainSpecificUtilFilterer: ChainSpecificUtilFilterer{contract: contract}}, nil
}

type ChainSpecificUtil struct {
	ChainSpecificUtilCaller
	ChainSpecificUtilTransactor
	ChainSpecificUtilFilterer
}

type ChainSpecificUtilCaller struct {
	contract *bind.BoundContract
}

type ChainSpecificUtilTransactor struct {
	contract *bind.BoundContract
}

type ChainSpecificUtilFilterer struct {
	contract *bind.BoundContract
}

type ChainSpecificUtilSession struct {
	Contract     *ChainSpecificUtil
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ChainSpecificUtilCallerSession struct {
	Contract *ChainSpecificUtilCaller
	CallOpts bind.CallOpts
}

type ChainSpecificUtilTransactorSession struct {
	Contract     *ChainSpecificUtilTransactor
	TransactOpts bind.TransactOpts
}

type ChainSpecificUtilRaw struct {
	Contract *ChainSpecificUtil
}

type ChainSpecificUtilCallerRaw struct {
	Contract *ChainSpecificUtilCaller
}

type ChainSpecificUtilTransactorRaw struct {
	Contract *ChainSpecificUtilTransactor
}

func NewChainSpecificUtil(address common.Address, backend bind.ContractBackend) (*ChainSpecificUtil, error) {
	contract, err := bindChainSpecificUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainSpecificUtil{ChainSpecificUtilCaller: ChainSpecificUtilCaller{contract: contract}, ChainSpecificUtilTransactor: ChainSpecificUtilTransactor{contract: contract}, ChainSpecificUtilFilterer: ChainSpecificUtilFilterer{contract: contract}}, nil
}

func NewChainSpecificUtilCaller(address common.Address, caller bind.ContractCaller) (*ChainSpecificUtilCaller, error) {
	contract, err := bindChainSpecificUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainSpecificUtilCaller{contract: contract}, nil
}

func NewChainSpecificUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainSpecificUtilTransactor, error) {
	contract, err := bindChainSpecificUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainSpecificUtilTransactor{contract: contract}, nil
}

func NewChainSpecificUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainSpecificUtilFilterer, error) {
	contract, err := bindChainSpecificUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainSpecificUtilFilterer{contract: contract}, nil
}

func bindChainSpecificUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainSpecificUtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ChainSpecificUtil *ChainSpecificUtilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainSpecificUtil.Contract.ChainSpecificUtilCaller.contract.Call(opts, result, method, params...)
}

func (_ChainSpecificUtil *ChainSpecificUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainSpecificUtil.Contract.ChainSpecificUtilTransactor.contract.Transfer(opts)
}

func (_ChainSpecificUtil *ChainSpecificUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainSpecificUtil.Contract.ChainSpecificUtilTransactor.contract.Transact(opts, method, params...)
}

func (_ChainSpecificUtil *ChainSpecificUtilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainSpecificUtil.Contract.contract.Call(opts, result, method, params...)
}

func (_ChainSpecificUtil *ChainSpecificUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainSpecificUtil.Contract.contract.Transfer(opts)
}

func (_ChainSpecificUtil *ChainSpecificUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainSpecificUtil.Contract.contract.Transact(opts, method, params...)
}

var ConfirmedOwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161045538038061045583398101604081905261002f9161016e565b8060006001600160a01b03821661008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100bd576100bd816100c5565b50505061019e565b336001600160a01b0382160361011d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561018057600080fd5b81516001600160a01b038116811461019757600080fd5b9392505050565b6102a8806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c634300080f000a",
}

var ConfirmedOwnerABI = ConfirmedOwnerMetaData.ABI

var ConfirmedOwnerBin = ConfirmedOwnerMetaData.Bin

func DeployConfirmedOwner(auth *bind.TransactOpts, backend bind.ContractBackend, newOwner common.Address) (common.Address, *types.Transaction, *ConfirmedOwner, error) {
	parsed, err := ConfirmedOwnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConfirmedOwnerBin), backend, newOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConfirmedOwner{ConfirmedOwnerCaller: ConfirmedOwnerCaller{contract: contract}, ConfirmedOwnerTransactor: ConfirmedOwnerTransactor{contract: contract}, ConfirmedOwnerFilterer: ConfirmedOwnerFilterer{contract: contract}}, nil
}

type ConfirmedOwner struct {
	ConfirmedOwnerCaller
	ConfirmedOwnerTransactor
	ConfirmedOwnerFilterer
}

type ConfirmedOwnerCaller struct {
	contract *bind.BoundContract
}

type ConfirmedOwnerTransactor struct {
	contract *bind.BoundContract
}

type ConfirmedOwnerFilterer struct {
	contract *bind.BoundContract
}

type ConfirmedOwnerSession struct {
	Contract     *ConfirmedOwner
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ConfirmedOwnerCallerSession struct {
	Contract *ConfirmedOwnerCaller
	CallOpts bind.CallOpts
}

type ConfirmedOwnerTransactorSession struct {
	Contract     *ConfirmedOwnerTransactor
	TransactOpts bind.TransactOpts
}

type ConfirmedOwnerRaw struct {
	Contract *ConfirmedOwner
}

type ConfirmedOwnerCallerRaw struct {
	Contract *ConfirmedOwnerCaller
}

type ConfirmedOwnerTransactorRaw struct {
	Contract *ConfirmedOwnerTransactor
}

func NewConfirmedOwner(address common.Address, backend bind.ContractBackend) (*ConfirmedOwner, error) {
	contract, err := bindConfirmedOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwner{ConfirmedOwnerCaller: ConfirmedOwnerCaller{contract: contract}, ConfirmedOwnerTransactor: ConfirmedOwnerTransactor{contract: contract}, ConfirmedOwnerFilterer: ConfirmedOwnerFilterer{contract: contract}}, nil
}

func NewConfirmedOwnerCaller(address common.Address, caller bind.ContractCaller) (*ConfirmedOwnerCaller, error) {
	contract, err := bindConfirmedOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerCaller{contract: contract}, nil
}

func NewConfirmedOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfirmedOwnerTransactor, error) {
	contract, err := bindConfirmedOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerTransactor{contract: contract}, nil
}

func NewConfirmedOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfirmedOwnerFilterer, error) {
	contract, err := bindConfirmedOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerFilterer{contract: contract}, nil
}

func bindConfirmedOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmedOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ConfirmedOwner *ConfirmedOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConfirmedOwner.Contract.ConfirmedOwnerCaller.contract.Call(opts, result, method, params...)
}

func (_ConfirmedOwner *ConfirmedOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.ConfirmedOwnerTransactor.contract.Transfer(opts)
}

func (_ConfirmedOwner *ConfirmedOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.ConfirmedOwnerTransactor.contract.Transact(opts, method, params...)
}

func (_ConfirmedOwner *ConfirmedOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConfirmedOwner.Contract.contract.Call(opts, result, method, params...)
}

func (_ConfirmedOwner *ConfirmedOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.contract.Transfer(opts)
}

func (_ConfirmedOwner *ConfirmedOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.contract.Transact(opts, method, params...)
}

func (_ConfirmedOwner *ConfirmedOwnerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConfirmedOwner.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ConfirmedOwner *ConfirmedOwnerSession) Owner() (common.Address, error) {
	return _ConfirmedOwner.Contract.Owner(&_ConfirmedOwner.CallOpts)
}

func (_ConfirmedOwner *ConfirmedOwnerCallerSession) Owner() (common.Address, error) {
	return _ConfirmedOwner.Contract.Owner(&_ConfirmedOwner.CallOpts)
}

func (_ConfirmedOwner *ConfirmedOwnerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfirmedOwner.contract.Transact(opts, "acceptOwnership")
}

func (_ConfirmedOwner *ConfirmedOwnerSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.AcceptOwnership(&_ConfirmedOwner.TransactOpts)
}

func (_ConfirmedOwner *ConfirmedOwnerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.AcceptOwnership(&_ConfirmedOwner.TransactOpts)
}

func (_ConfirmedOwner *ConfirmedOwnerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ConfirmedOwner.contract.Transact(opts, "transferOwnership", to)
}

func (_ConfirmedOwner *ConfirmedOwnerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.TransferOwnership(&_ConfirmedOwner.TransactOpts, to)
}

func (_ConfirmedOwner *ConfirmedOwnerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ConfirmedOwner.Contract.TransferOwnership(&_ConfirmedOwner.TransactOpts, to)
}

type ConfirmedOwnerOwnershipTransferRequestedIterator struct {
	Event *ConfirmedOwnerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfirmedOwnerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmedOwnerOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(ConfirmedOwnerOwnershipTransferRequested)
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

func (it *ConfirmedOwnerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ConfirmedOwnerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfirmedOwnerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ConfirmedOwner *ConfirmedOwnerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfirmedOwnerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwner.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerOwnershipTransferRequestedIterator{contract: _ConfirmedOwner.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ConfirmedOwner *ConfirmedOwnerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ConfirmedOwnerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwner.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfirmedOwnerOwnershipTransferRequested)
				if err := _ConfirmedOwner.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ConfirmedOwner *ConfirmedOwnerFilterer) ParseOwnershipTransferRequested(log types.Log) (*ConfirmedOwnerOwnershipTransferRequested, error) {
	event := new(ConfirmedOwnerOwnershipTransferRequested)
	if err := _ConfirmedOwner.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConfirmedOwnerOwnershipTransferredIterator struct {
	Event *ConfirmedOwnerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfirmedOwnerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmedOwnerOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(ConfirmedOwnerOwnershipTransferred)
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

func (it *ConfirmedOwnerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ConfirmedOwnerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfirmedOwnerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ConfirmedOwner *ConfirmedOwnerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfirmedOwnerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwner.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerOwnershipTransferredIterator{contract: _ConfirmedOwner.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ConfirmedOwner *ConfirmedOwnerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConfirmedOwnerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwner.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfirmedOwnerOwnershipTransferred)
				if err := _ConfirmedOwner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ConfirmedOwner *ConfirmedOwnerFilterer) ParseOwnershipTransferred(log types.Log) (*ConfirmedOwnerOwnershipTransferred, error) {
	event := new(ConfirmedOwnerOwnershipTransferred)
	if err := _ConfirmedOwner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var ConfirmedOwnerWithProposalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161047038038061047083398101604081905261002f91610186565b6001600160a01b03821661008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100ba576100ba816100c1565b50506101b9565b336001600160a01b038216036101195760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610081565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b038116811461018157600080fd5b919050565b6000806040838503121561019957600080fd5b6101a28361016a565b91506101b06020840161016a565b90509250929050565b6102a8806101c86000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c634300080f000a",
}

var ConfirmedOwnerWithProposalABI = ConfirmedOwnerWithProposalMetaData.ABI

var ConfirmedOwnerWithProposalBin = ConfirmedOwnerWithProposalMetaData.Bin

func DeployConfirmedOwnerWithProposal(auth *bind.TransactOpts, backend bind.ContractBackend, newOwner common.Address, pendingOwner common.Address) (common.Address, *types.Transaction, *ConfirmedOwnerWithProposal, error) {
	parsed, err := ConfirmedOwnerWithProposalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConfirmedOwnerWithProposalBin), backend, newOwner, pendingOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConfirmedOwnerWithProposal{ConfirmedOwnerWithProposalCaller: ConfirmedOwnerWithProposalCaller{contract: contract}, ConfirmedOwnerWithProposalTransactor: ConfirmedOwnerWithProposalTransactor{contract: contract}, ConfirmedOwnerWithProposalFilterer: ConfirmedOwnerWithProposalFilterer{contract: contract}}, nil
}

type ConfirmedOwnerWithProposal struct {
	ConfirmedOwnerWithProposalCaller
	ConfirmedOwnerWithProposalTransactor
	ConfirmedOwnerWithProposalFilterer
}

type ConfirmedOwnerWithProposalCaller struct {
	contract *bind.BoundContract
}

type ConfirmedOwnerWithProposalTransactor struct {
	contract *bind.BoundContract
}

type ConfirmedOwnerWithProposalFilterer struct {
	contract *bind.BoundContract
}

type ConfirmedOwnerWithProposalSession struct {
	Contract     *ConfirmedOwnerWithProposal
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ConfirmedOwnerWithProposalCallerSession struct {
	Contract *ConfirmedOwnerWithProposalCaller
	CallOpts bind.CallOpts
}

type ConfirmedOwnerWithProposalTransactorSession struct {
	Contract     *ConfirmedOwnerWithProposalTransactor
	TransactOpts bind.TransactOpts
}

type ConfirmedOwnerWithProposalRaw struct {
	Contract *ConfirmedOwnerWithProposal
}

type ConfirmedOwnerWithProposalCallerRaw struct {
	Contract *ConfirmedOwnerWithProposalCaller
}

type ConfirmedOwnerWithProposalTransactorRaw struct {
	Contract *ConfirmedOwnerWithProposalTransactor
}

func NewConfirmedOwnerWithProposal(address common.Address, backend bind.ContractBackend) (*ConfirmedOwnerWithProposal, error) {
	contract, err := bindConfirmedOwnerWithProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerWithProposal{ConfirmedOwnerWithProposalCaller: ConfirmedOwnerWithProposalCaller{contract: contract}, ConfirmedOwnerWithProposalTransactor: ConfirmedOwnerWithProposalTransactor{contract: contract}, ConfirmedOwnerWithProposalFilterer: ConfirmedOwnerWithProposalFilterer{contract: contract}}, nil
}

func NewConfirmedOwnerWithProposalCaller(address common.Address, caller bind.ContractCaller) (*ConfirmedOwnerWithProposalCaller, error) {
	contract, err := bindConfirmedOwnerWithProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerWithProposalCaller{contract: contract}, nil
}

func NewConfirmedOwnerWithProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfirmedOwnerWithProposalTransactor, error) {
	contract, err := bindConfirmedOwnerWithProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerWithProposalTransactor{contract: contract}, nil
}

func NewConfirmedOwnerWithProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfirmedOwnerWithProposalFilterer, error) {
	contract, err := bindConfirmedOwnerWithProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerWithProposalFilterer{contract: contract}, nil
}

func bindConfirmedOwnerWithProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmedOwnerWithProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConfirmedOwnerWithProposal.Contract.ConfirmedOwnerWithProposalCaller.contract.Call(opts, result, method, params...)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.ConfirmedOwnerWithProposalTransactor.contract.Transfer(opts)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.ConfirmedOwnerWithProposalTransactor.contract.Transact(opts, method, params...)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConfirmedOwnerWithProposal.Contract.contract.Call(opts, result, method, params...)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.contract.Transfer(opts)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.contract.Transact(opts, method, params...)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConfirmedOwnerWithProposal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalSession) Owner() (common.Address, error) {
	return _ConfirmedOwnerWithProposal.Contract.Owner(&_ConfirmedOwnerWithProposal.CallOpts)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalCallerSession) Owner() (common.Address, error) {
	return _ConfirmedOwnerWithProposal.Contract.Owner(&_ConfirmedOwnerWithProposal.CallOpts)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.contract.Transact(opts, "acceptOwnership")
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.AcceptOwnership(&_ConfirmedOwnerWithProposal.TransactOpts)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.AcceptOwnership(&_ConfirmedOwnerWithProposal.TransactOpts)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.contract.Transact(opts, "transferOwnership", to)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.TransferOwnership(&_ConfirmedOwnerWithProposal.TransactOpts, to)
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ConfirmedOwnerWithProposal.Contract.TransferOwnership(&_ConfirmedOwnerWithProposal.TransactOpts, to)
}

type ConfirmedOwnerWithProposalOwnershipTransferRequestedIterator struct {
	Event *ConfirmedOwnerWithProposalOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfirmedOwnerWithProposalOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmedOwnerWithProposalOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(ConfirmedOwnerWithProposalOwnershipTransferRequested)
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

func (it *ConfirmedOwnerWithProposalOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ConfirmedOwnerWithProposalOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfirmedOwnerWithProposalOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfirmedOwnerWithProposalOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwnerWithProposal.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerWithProposalOwnershipTransferRequestedIterator{contract: _ConfirmedOwnerWithProposal.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ConfirmedOwnerWithProposalOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwnerWithProposal.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfirmedOwnerWithProposalOwnershipTransferRequested)
				if err := _ConfirmedOwnerWithProposal.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalFilterer) ParseOwnershipTransferRequested(log types.Log) (*ConfirmedOwnerWithProposalOwnershipTransferRequested, error) {
	event := new(ConfirmedOwnerWithProposalOwnershipTransferRequested)
	if err := _ConfirmedOwnerWithProposal.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConfirmedOwnerWithProposalOwnershipTransferredIterator struct {
	Event *ConfirmedOwnerWithProposalOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfirmedOwnerWithProposalOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmedOwnerWithProposalOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(ConfirmedOwnerWithProposalOwnershipTransferred)
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

func (it *ConfirmedOwnerWithProposalOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ConfirmedOwnerWithProposalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfirmedOwnerWithProposalOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfirmedOwnerWithProposalOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwnerWithProposal.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConfirmedOwnerWithProposalOwnershipTransferredIterator{contract: _ConfirmedOwnerWithProposal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConfirmedOwnerWithProposalOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConfirmedOwnerWithProposal.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfirmedOwnerWithProposalOwnershipTransferred)
				if err := _ConfirmedOwnerWithProposal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ConfirmedOwnerWithProposal *ConfirmedOwnerWithProposalFilterer) ParseOwnershipTransferred(log types.Log) (*ConfirmedOwnerWithProposalOwnershipTransferred, error) {
	event := new(ConfirmedOwnerWithProposalOwnershipTransferred)
	if err := _ConfirmedOwnerWithProposal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var ECCArithmeticMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50601680601d6000396000f3fe6080604052600080fdfea164736f6c634300080f000a",
}

var ECCArithmeticABI = ECCArithmeticMetaData.ABI

var ECCArithmeticBin = ECCArithmeticMetaData.Bin

func DeployECCArithmetic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECCArithmetic, error) {
	parsed, err := ECCArithmeticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECCArithmeticBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECCArithmetic{ECCArithmeticCaller: ECCArithmeticCaller{contract: contract}, ECCArithmeticTransactor: ECCArithmeticTransactor{contract: contract}, ECCArithmeticFilterer: ECCArithmeticFilterer{contract: contract}}, nil
}

type ECCArithmetic struct {
	ECCArithmeticCaller
	ECCArithmeticTransactor
	ECCArithmeticFilterer
}

type ECCArithmeticCaller struct {
	contract *bind.BoundContract
}

type ECCArithmeticTransactor struct {
	contract *bind.BoundContract
}

type ECCArithmeticFilterer struct {
	contract *bind.BoundContract
}

type ECCArithmeticSession struct {
	Contract     *ECCArithmetic
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ECCArithmeticCallerSession struct {
	Contract *ECCArithmeticCaller
	CallOpts bind.CallOpts
}

type ECCArithmeticTransactorSession struct {
	Contract     *ECCArithmeticTransactor
	TransactOpts bind.TransactOpts
}

type ECCArithmeticRaw struct {
	Contract *ECCArithmetic
}

type ECCArithmeticCallerRaw struct {
	Contract *ECCArithmeticCaller
}

type ECCArithmeticTransactorRaw struct {
	Contract *ECCArithmeticTransactor
}

func NewECCArithmetic(address common.Address, backend bind.ContractBackend) (*ECCArithmetic, error) {
	contract, err := bindECCArithmetic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECCArithmetic{ECCArithmeticCaller: ECCArithmeticCaller{contract: contract}, ECCArithmeticTransactor: ECCArithmeticTransactor{contract: contract}, ECCArithmeticFilterer: ECCArithmeticFilterer{contract: contract}}, nil
}

func NewECCArithmeticCaller(address common.Address, caller bind.ContractCaller) (*ECCArithmeticCaller, error) {
	contract, err := bindECCArithmetic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECCArithmeticCaller{contract: contract}, nil
}

func NewECCArithmeticTransactor(address common.Address, transactor bind.ContractTransactor) (*ECCArithmeticTransactor, error) {
	contract, err := bindECCArithmetic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECCArithmeticTransactor{contract: contract}, nil
}

func NewECCArithmeticFilterer(address common.Address, filterer bind.ContractFilterer) (*ECCArithmeticFilterer, error) {
	contract, err := bindECCArithmetic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECCArithmeticFilterer{contract: contract}, nil
}

func bindECCArithmetic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCArithmeticABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ECCArithmetic *ECCArithmeticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECCArithmetic.Contract.ECCArithmeticCaller.contract.Call(opts, result, method, params...)
}

func (_ECCArithmetic *ECCArithmeticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCArithmetic.Contract.ECCArithmeticTransactor.contract.Transfer(opts)
}

func (_ECCArithmetic *ECCArithmeticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCArithmetic.Contract.ECCArithmeticTransactor.contract.Transact(opts, method, params...)
}

func (_ECCArithmetic *ECCArithmeticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECCArithmetic.Contract.contract.Call(opts, result, method, params...)
}

func (_ECCArithmetic *ECCArithmeticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCArithmetic.Contract.contract.Transfer(opts)
}

func (_ECCArithmetic *ECCArithmeticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCArithmetic.Contract.contract.Transact(opts, method, params...)
}

var ERC677ReceiverInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var ERC677ReceiverInterfaceABI = ERC677ReceiverInterfaceMetaData.ABI

type ERC677ReceiverInterface struct {
	ERC677ReceiverInterfaceCaller
	ERC677ReceiverInterfaceTransactor
	ERC677ReceiverInterfaceFilterer
}

type ERC677ReceiverInterfaceCaller struct {
	contract *bind.BoundContract
}

type ERC677ReceiverInterfaceTransactor struct {
	contract *bind.BoundContract
}

type ERC677ReceiverInterfaceFilterer struct {
	contract *bind.BoundContract
}

type ERC677ReceiverInterfaceSession struct {
	Contract     *ERC677ReceiverInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ERC677ReceiverInterfaceCallerSession struct {
	Contract *ERC677ReceiverInterfaceCaller
	CallOpts bind.CallOpts
}

type ERC677ReceiverInterfaceTransactorSession struct {
	Contract     *ERC677ReceiverInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type ERC677ReceiverInterfaceRaw struct {
	Contract *ERC677ReceiverInterface
}

type ERC677ReceiverInterfaceCallerRaw struct {
	Contract *ERC677ReceiverInterfaceCaller
}

type ERC677ReceiverInterfaceTransactorRaw struct {
	Contract *ERC677ReceiverInterfaceTransactor
}

func NewERC677ReceiverInterface(address common.Address, backend bind.ContractBackend) (*ERC677ReceiverInterface, error) {
	contract, err := bindERC677ReceiverInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC677ReceiverInterface{ERC677ReceiverInterfaceCaller: ERC677ReceiverInterfaceCaller{contract: contract}, ERC677ReceiverInterfaceTransactor: ERC677ReceiverInterfaceTransactor{contract: contract}, ERC677ReceiverInterfaceFilterer: ERC677ReceiverInterfaceFilterer{contract: contract}}, nil
}

func NewERC677ReceiverInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC677ReceiverInterfaceCaller, error) {
	contract, err := bindERC677ReceiverInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC677ReceiverInterfaceCaller{contract: contract}, nil
}

func NewERC677ReceiverInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC677ReceiverInterfaceTransactor, error) {
	contract, err := bindERC677ReceiverInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC677ReceiverInterfaceTransactor{contract: contract}, nil
}

func NewERC677ReceiverInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC677ReceiverInterfaceFilterer, error) {
	contract, err := bindERC677ReceiverInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC677ReceiverInterfaceFilterer{contract: contract}, nil
}

func bindERC677ReceiverInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC677ReceiverInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC677ReceiverInterface.Contract.ERC677ReceiverInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC677ReceiverInterface.Contract.ERC677ReceiverInterfaceTransactor.contract.Transfer(opts)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC677ReceiverInterface.Contract.ERC677ReceiverInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC677ReceiverInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC677ReceiverInterface.Contract.contract.Transfer(opts)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC677ReceiverInterface.Contract.contract.Transact(opts, method, params...)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceTransactor) OnTokenTransfer(opts *bind.TransactOpts, sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC677ReceiverInterface.contract.Transact(opts, "onTokenTransfer", sender, amount, data)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC677ReceiverInterface.Contract.OnTokenTransfer(&_ERC677ReceiverInterface.TransactOpts, sender, amount, data)
}

func (_ERC677ReceiverInterface *ERC677ReceiverInterfaceTransactorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC677ReceiverInterface.Contract.OnTokenTransfer(&_ERC677ReceiverInterface.TransactOpts, sender, amount, data)
}

var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x602d6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea164736f6c634300080f000a",
}

var EnumerableSetABI = EnumerableSetMetaData.ABI

var EnumerableSetBin = EnumerableSetMetaData.Bin

func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

type EnumerableSet struct {
	EnumerableSetCaller
	EnumerableSetTransactor
	EnumerableSetFilterer
}

type EnumerableSetCaller struct {
	contract *bind.BoundContract
}

type EnumerableSetTransactor struct {
	contract *bind.BoundContract
}

type EnumerableSetFilterer struct {
	contract *bind.BoundContract
}

type EnumerableSetSession struct {
	Contract     *EnumerableSet
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller
	CallOpts bind.CallOpts
}

type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor
	TransactOpts bind.TransactOpts
}

type EnumerableSetRaw struct {
	Contract *EnumerableSet
}

type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller
}

type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor
}

func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

var IVRFRouterConsumerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IVRFRouterConsumerABI = IVRFRouterConsumerMetaData.ABI

type IVRFRouterConsumer struct {
	IVRFRouterConsumerCaller
	IVRFRouterConsumerTransactor
	IVRFRouterConsumerFilterer
}

type IVRFRouterConsumerCaller struct {
	contract *bind.BoundContract
}

type IVRFRouterConsumerTransactor struct {
	contract *bind.BoundContract
}

type IVRFRouterConsumerFilterer struct {
	contract *bind.BoundContract
}

type IVRFRouterConsumerSession struct {
	Contract     *IVRFRouterConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFRouterConsumerCallerSession struct {
	Contract *IVRFRouterConsumerCaller
	CallOpts bind.CallOpts
}

type IVRFRouterConsumerTransactorSession struct {
	Contract     *IVRFRouterConsumerTransactor
	TransactOpts bind.TransactOpts
}

type IVRFRouterConsumerRaw struct {
	Contract *IVRFRouterConsumer
}

type IVRFRouterConsumerCallerRaw struct {
	Contract *IVRFRouterConsumerCaller
}

type IVRFRouterConsumerTransactorRaw struct {
	Contract *IVRFRouterConsumerTransactor
}

func NewIVRFRouterConsumer(address common.Address, backend bind.ContractBackend) (*IVRFRouterConsumer, error) {
	contract, err := bindIVRFRouterConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumer{IVRFRouterConsumerCaller: IVRFRouterConsumerCaller{contract: contract}, IVRFRouterConsumerTransactor: IVRFRouterConsumerTransactor{contract: contract}, IVRFRouterConsumerFilterer: IVRFRouterConsumerFilterer{contract: contract}}, nil
}

func NewIVRFRouterConsumerCaller(address common.Address, caller bind.ContractCaller) (*IVRFRouterConsumerCaller, error) {
	contract, err := bindIVRFRouterConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerCaller{contract: contract}, nil
}

func NewIVRFRouterConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFRouterConsumerTransactor, error) {
	contract, err := bindIVRFRouterConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerTransactor{contract: contract}, nil
}

func NewIVRFRouterConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFRouterConsumerFilterer, error) {
	contract, err := bindIVRFRouterConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerFilterer{contract: contract}, nil
}

func bindIVRFRouterConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFRouterConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFRouterConsumer *IVRFRouterConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFRouterConsumer.Contract.IVRFRouterConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFRouterConsumer.Contract.IVRFRouterConsumerTransactor.contract.Transfer(opts)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFRouterConsumer.Contract.IVRFRouterConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFRouterConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFRouterConsumer.Contract.contract.Transfer(opts)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFRouterConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IVRFRouterConsumer.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_IVRFRouterConsumer *IVRFRouterConsumerSession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFRouterConsumer.Contract.NUMCONFDELAYS(&_IVRFRouterConsumer.CallOpts)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFRouterConsumer.Contract.NUMCONFDELAYS(&_IVRFRouterConsumer.CallOpts)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFRouterConsumer.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords, arguments)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFRouterConsumer.Contract.RawFulfillRandomWords(&_IVRFRouterConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_IVRFRouterConsumer *IVRFRouterConsumerTransactorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFRouterConsumer.Contract.RawFulfillRandomWords(&_IVRFRouterConsumer.TransactOpts, requestID, randomWords, arguments)
}

type IVRFRouterConsumerConfigSetIterator struct {
	Event *IVRFRouterConsumerConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFRouterConsumerConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFRouterConsumerConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(IVRFRouterConsumerConfigSet)
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

func (it *IVRFRouterConsumerConfigSetIterator) Error() error {
	return it.fail
}

func (it *IVRFRouterConsumerConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFRouterConsumerConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) FilterConfigSet(opts *bind.FilterOpts) (*IVRFRouterConsumerConfigSetIterator, error) {

	logs, sub, err := _IVRFRouterConsumer.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerConfigSetIterator{contract: _IVRFRouterConsumer.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *IVRFRouterConsumerConfigSet) (event.Subscription, error) {

	logs, sub, err := _IVRFRouterConsumer.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFRouterConsumerConfigSet)
				if err := _IVRFRouterConsumer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) ParseConfigSet(log types.Log) (*IVRFRouterConsumerConfigSet, error) {
	event := new(IVRFRouterConsumerConfigSet)
	if err := _IVRFRouterConsumer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFRouterConsumerNewTransmissionIterator struct {
	Event *IVRFRouterConsumerNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFRouterConsumerNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFRouterConsumerNewTransmission)
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

	select {
	case log := <-it.logs:
		it.Event = new(IVRFRouterConsumerNewTransmission)
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

func (it *IVRFRouterConsumerNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *IVRFRouterConsumerNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFRouterConsumerNewTransmission struct {
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*IVRFRouterConsumerNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _IVRFRouterConsumer.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerNewTransmissionIterator{contract: _IVRFRouterConsumer.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *IVRFRouterConsumerNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _IVRFRouterConsumer.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFRouterConsumerNewTransmission)
				if err := _IVRFRouterConsumer.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) ParseNewTransmission(log types.Log) (*IVRFRouterConsumerNewTransmission, error) {
	event := new(IVRFRouterConsumerNewTransmission)
	if err := _IVRFRouterConsumer.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFRouterConsumerOutputsServedIterator struct {
	Event *IVRFRouterConsumerOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFRouterConsumerOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFRouterConsumerOutputsServed)
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

	select {
	case log := <-it.logs:
		it.Event = new(IVRFRouterConsumerOutputsServed)
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

func (it *IVRFRouterConsumerOutputsServedIterator) Error() error {
	return it.fail
}

func (it *IVRFRouterConsumerOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFRouterConsumerOutputsServed struct {
	RecentBlockHeight  uint64
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*IVRFRouterConsumerOutputsServedIterator, error) {

	logs, sub, err := _IVRFRouterConsumer.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerOutputsServedIterator{contract: _IVRFRouterConsumer.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *IVRFRouterConsumerOutputsServed) (event.Subscription, error) {

	logs, sub, err := _IVRFRouterConsumer.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFRouterConsumerOutputsServed)
				if err := _IVRFRouterConsumer.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) ParseOutputsServed(log types.Log) (*IVRFRouterConsumerOutputsServed, error) {
	event := new(IVRFRouterConsumerOutputsServed)
	if err := _IVRFRouterConsumer.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFRouterConsumerRandomWordsFulfilledIterator struct {
	Event *IVRFRouterConsumerRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFRouterConsumerRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFRouterConsumerRandomWordsFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(IVRFRouterConsumerRandomWordsFulfilled)
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

func (it *IVRFRouterConsumerRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *IVRFRouterConsumerRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFRouterConsumerRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*IVRFRouterConsumerRandomWordsFulfilledIterator, error) {

	logs, sub, err := _IVRFRouterConsumer.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerRandomWordsFulfilledIterator{contract: _IVRFRouterConsumer.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *IVRFRouterConsumerRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _IVRFRouterConsumer.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFRouterConsumerRandomWordsFulfilled)
				if err := _IVRFRouterConsumer.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) ParseRandomWordsFulfilled(log types.Log) (*IVRFRouterConsumerRandomWordsFulfilled, error) {
	event := new(IVRFRouterConsumerRandomWordsFulfilled)
	if err := _IVRFRouterConsumer.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFRouterConsumerRandomnessFulfillmentRequestedIterator struct {
	Event *IVRFRouterConsumerRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFRouterConsumerRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFRouterConsumerRandomnessFulfillmentRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(IVRFRouterConsumerRandomnessFulfillmentRequested)
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

func (it *IVRFRouterConsumerRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *IVRFRouterConsumerRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFRouterConsumerRandomnessFulfillmentRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	GasAllowance           uint32
	GasPrice               *big.Int
	WeiPerUnitLink         *big.Int
	Arguments              []byte
	Raw                    types.Log
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*IVRFRouterConsumerRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFRouterConsumer.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerRandomnessFulfillmentRequestedIterator{contract: _IVRFRouterConsumer.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *IVRFRouterConsumerRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFRouterConsumer.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFRouterConsumerRandomnessFulfillmentRequested)
				if err := _IVRFRouterConsumer.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*IVRFRouterConsumerRandomnessFulfillmentRequested, error) {
	event := new(IVRFRouterConsumerRandomnessFulfillmentRequested)
	if err := _IVRFRouterConsumer.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFRouterConsumerRandomnessRequestedIterator struct {
	Event *IVRFRouterConsumerRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFRouterConsumerRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFRouterConsumerRandomnessRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(IVRFRouterConsumerRandomnessRequested)
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

func (it *IVRFRouterConsumerRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *IVRFRouterConsumerRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFRouterConsumerRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	Raw                    types.Log
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*IVRFRouterConsumerRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFRouterConsumer.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &IVRFRouterConsumerRandomnessRequestedIterator{contract: _IVRFRouterConsumer.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *IVRFRouterConsumerRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFRouterConsumer.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFRouterConsumerRandomnessRequested)
				if err := _IVRFRouterConsumer.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_IVRFRouterConsumer *IVRFRouterConsumerFilterer) ParseRandomnessRequested(log types.Log) (*IVRFRouterConsumerRandomnessRequested, error) {
	event := new(IVRFRouterConsumerRandomnessRequested)
	if err := _IVRFRouterConsumer.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var LinkTokenInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimalPlaces\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTokensIssued\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var LinkTokenInterfaceABI = LinkTokenInterfaceMetaData.ABI

type LinkTokenInterface struct {
	LinkTokenInterfaceCaller
	LinkTokenInterfaceTransactor
	LinkTokenInterfaceFilterer
}

type LinkTokenInterfaceCaller struct {
	contract *bind.BoundContract
}

type LinkTokenInterfaceTransactor struct {
	contract *bind.BoundContract
}

type LinkTokenInterfaceFilterer struct {
	contract *bind.BoundContract
}

type LinkTokenInterfaceSession struct {
	Contract     *LinkTokenInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LinkTokenInterfaceCallerSession struct {
	Contract *LinkTokenInterfaceCaller
	CallOpts bind.CallOpts
}

type LinkTokenInterfaceTransactorSession struct {
	Contract     *LinkTokenInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type LinkTokenInterfaceRaw struct {
	Contract *LinkTokenInterface
}

type LinkTokenInterfaceCallerRaw struct {
	Contract *LinkTokenInterfaceCaller
}

type LinkTokenInterfaceTransactorRaw struct {
	Contract *LinkTokenInterfaceTransactor
}

func NewLinkTokenInterface(address common.Address, backend bind.ContractBackend) (*LinkTokenInterface, error) {
	contract, err := bindLinkTokenInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterface{LinkTokenInterfaceCaller: LinkTokenInterfaceCaller{contract: contract}, LinkTokenInterfaceTransactor: LinkTokenInterfaceTransactor{contract: contract}, LinkTokenInterfaceFilterer: LinkTokenInterfaceFilterer{contract: contract}}, nil
}

func NewLinkTokenInterfaceCaller(address common.Address, caller bind.ContractCaller) (*LinkTokenInterfaceCaller, error) {
	contract, err := bindLinkTokenInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceCaller{contract: contract}, nil
}

func NewLinkTokenInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkTokenInterfaceTransactor, error) {
	contract, err := bindLinkTokenInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceTransactor{contract: contract}, nil
}

func NewLinkTokenInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkTokenInterfaceFilterer, error) {
	contract, err := bindLinkTokenInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkTokenInterfaceFilterer{contract: contract}, nil
}

func bindLinkTokenInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkTokenInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_LinkTokenInterface *LinkTokenInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceTransactor.contract.Transfer(opts)
}

func (_LinkTokenInterface *LinkTokenInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.LinkTokenInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkTokenInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.contract.Transfer(opts)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.contract.Transact(opts, method, params...)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.Allowance(&_LinkTokenInterface.CallOpts, owner, spender)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.Allowance(&_LinkTokenInterface.CallOpts, owner, spender)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.BalanceOf(&_LinkTokenInterface.CallOpts, owner)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkTokenInterface.Contract.BalanceOf(&_LinkTokenInterface.CallOpts, owner)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Decimals() (uint8, error) {
	return _LinkTokenInterface.Contract.Decimals(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Decimals() (uint8, error) {
	return _LinkTokenInterface.Contract.Decimals(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Name() (string, error) {
	return _LinkTokenInterface.Contract.Name(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Name() (string, error) {
	return _LinkTokenInterface.Contract.Name(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Symbol() (string, error) {
	return _LinkTokenInterface.Contract.Symbol(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) Symbol() (string, error) {
	return _LinkTokenInterface.Contract.Symbol(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinkTokenInterface.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LinkTokenInterface *LinkTokenInterfaceSession) TotalSupply() (*big.Int, error) {
	return _LinkTokenInterface.Contract.TotalSupply(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _LinkTokenInterface.Contract.TotalSupply(&_LinkTokenInterface.CallOpts)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "approve", spender, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Approve(&_LinkTokenInterface.TransactOpts, spender, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Approve(&_LinkTokenInterface.TransactOpts, spender, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "decreaseApproval", spender, addedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) DecreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.DecreaseApproval(&_LinkTokenInterface.TransactOpts, spender, addedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) DecreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.DecreaseApproval(&_LinkTokenInterface.TransactOpts, spender, addedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "increaseApproval", spender, subtractedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) IncreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.IncreaseApproval(&_LinkTokenInterface.TransactOpts, spender, subtractedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) IncreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.IncreaseApproval(&_LinkTokenInterface.TransactOpts, spender, subtractedValue)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transfer", to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Transfer(&_LinkTokenInterface.TransactOpts, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.Transfer(&_LinkTokenInterface.TransactOpts, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transferAndCall", to, value, data)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferAndCall(&_LinkTokenInterface.TransactOpts, to, value, data)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferAndCall(&_LinkTokenInterface.TransactOpts, to, value, data)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferFrom(&_LinkTokenInterface.TransactOpts, from, to, value)
}

func (_LinkTokenInterface *LinkTokenInterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _LinkTokenInterface.Contract.TransferFrom(&_LinkTokenInterface.TransactOpts, from, to, value)
}

var OwnableInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var OwnableInterfaceABI = OwnableInterfaceMetaData.ABI

type OwnableInterface struct {
	OwnableInterfaceCaller
	OwnableInterfaceTransactor
	OwnableInterfaceFilterer
}

type OwnableInterfaceCaller struct {
	contract *bind.BoundContract
}

type OwnableInterfaceTransactor struct {
	contract *bind.BoundContract
}

type OwnableInterfaceFilterer struct {
	contract *bind.BoundContract
}

type OwnableInterfaceSession struct {
	Contract     *OwnableInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OwnableInterfaceCallerSession struct {
	Contract *OwnableInterfaceCaller
	CallOpts bind.CallOpts
}

type OwnableInterfaceTransactorSession struct {
	Contract     *OwnableInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type OwnableInterfaceRaw struct {
	Contract *OwnableInterface
}

type OwnableInterfaceCallerRaw struct {
	Contract *OwnableInterfaceCaller
}

type OwnableInterfaceTransactorRaw struct {
	Contract *OwnableInterfaceTransactor
}

func NewOwnableInterface(address common.Address, backend bind.ContractBackend) (*OwnableInterface, error) {
	contract, err := bindOwnableInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableInterface{OwnableInterfaceCaller: OwnableInterfaceCaller{contract: contract}, OwnableInterfaceTransactor: OwnableInterfaceTransactor{contract: contract}, OwnableInterfaceFilterer: OwnableInterfaceFilterer{contract: contract}}, nil
}

func NewOwnableInterfaceCaller(address common.Address, caller bind.ContractCaller) (*OwnableInterfaceCaller, error) {
	contract, err := bindOwnableInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableInterfaceCaller{contract: contract}, nil
}

func NewOwnableInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableInterfaceTransactor, error) {
	contract, err := bindOwnableInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableInterfaceTransactor{contract: contract}, nil
}

func NewOwnableInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableInterfaceFilterer, error) {
	contract, err := bindOwnableInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableInterfaceFilterer{contract: contract}, nil
}

func bindOwnableInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OwnableInterface *OwnableInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableInterface.Contract.OwnableInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_OwnableInterface *OwnableInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableInterface.Contract.OwnableInterfaceTransactor.contract.Transfer(opts)
}

func (_OwnableInterface *OwnableInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableInterface.Contract.OwnableInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_OwnableInterface *OwnableInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_OwnableInterface *OwnableInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableInterface.Contract.contract.Transfer(opts)
}

func (_OwnableInterface *OwnableInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableInterface.Contract.contract.Transact(opts, method, params...)
}

func (_OwnableInterface *OwnableInterfaceTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableInterface.contract.Transact(opts, "acceptOwnership")
}

func (_OwnableInterface *OwnableInterfaceSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableInterface.Contract.AcceptOwnership(&_OwnableInterface.TransactOpts)
}

func (_OwnableInterface *OwnableInterfaceTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableInterface.Contract.AcceptOwnership(&_OwnableInterface.TransactOpts)
}

func (_OwnableInterface *OwnableInterfaceTransactor) Owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableInterface.contract.Transact(opts, "owner")
}

func (_OwnableInterface *OwnableInterfaceSession) Owner() (*types.Transaction, error) {
	return _OwnableInterface.Contract.Owner(&_OwnableInterface.TransactOpts)
}

func (_OwnableInterface *OwnableInterfaceTransactorSession) Owner() (*types.Transaction, error) {
	return _OwnableInterface.Contract.Owner(&_OwnableInterface.TransactOpts)
}

func (_OwnableInterface *OwnableInterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _OwnableInterface.contract.Transact(opts, "transferOwnership", recipient)
}

func (_OwnableInterface *OwnableInterfaceSession) TransferOwnership(recipient common.Address) (*types.Transaction, error) {
	return _OwnableInterface.Contract.TransferOwnership(&_OwnableInterface.TransactOpts, recipient)
}

func (_OwnableInterface *OwnableInterfaceTransactorSession) TransferOwnership(recipient common.Address) (*types.Transaction, error) {
	return _OwnableInterface.Contract.TransferOwnership(&_OwnableInterface.TransactOpts, recipient)
}

var OwnerIsCreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6102a8806101576000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c634300080f000a",
}

var OwnerIsCreatorABI = OwnerIsCreatorMetaData.ABI

var OwnerIsCreatorBin = OwnerIsCreatorMetaData.Bin

func DeployOwnerIsCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OwnerIsCreator, error) {
	parsed, err := OwnerIsCreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnerIsCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnerIsCreator{OwnerIsCreatorCaller: OwnerIsCreatorCaller{contract: contract}, OwnerIsCreatorTransactor: OwnerIsCreatorTransactor{contract: contract}, OwnerIsCreatorFilterer: OwnerIsCreatorFilterer{contract: contract}}, nil
}

type OwnerIsCreator struct {
	OwnerIsCreatorCaller
	OwnerIsCreatorTransactor
	OwnerIsCreatorFilterer
}

type OwnerIsCreatorCaller struct {
	contract *bind.BoundContract
}

type OwnerIsCreatorTransactor struct {
	contract *bind.BoundContract
}

type OwnerIsCreatorFilterer struct {
	contract *bind.BoundContract
}

type OwnerIsCreatorSession struct {
	Contract     *OwnerIsCreator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OwnerIsCreatorCallerSession struct {
	Contract *OwnerIsCreatorCaller
	CallOpts bind.CallOpts
}

type OwnerIsCreatorTransactorSession struct {
	Contract     *OwnerIsCreatorTransactor
	TransactOpts bind.TransactOpts
}

type OwnerIsCreatorRaw struct {
	Contract *OwnerIsCreator
}

type OwnerIsCreatorCallerRaw struct {
	Contract *OwnerIsCreatorCaller
}

type OwnerIsCreatorTransactorRaw struct {
	Contract *OwnerIsCreatorTransactor
}

func NewOwnerIsCreator(address common.Address, backend bind.ContractBackend) (*OwnerIsCreator, error) {
	contract, err := bindOwnerIsCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnerIsCreator{OwnerIsCreatorCaller: OwnerIsCreatorCaller{contract: contract}, OwnerIsCreatorTransactor: OwnerIsCreatorTransactor{contract: contract}, OwnerIsCreatorFilterer: OwnerIsCreatorFilterer{contract: contract}}, nil
}

func NewOwnerIsCreatorCaller(address common.Address, caller bind.ContractCaller) (*OwnerIsCreatorCaller, error) {
	contract, err := bindOwnerIsCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerIsCreatorCaller{contract: contract}, nil
}

func NewOwnerIsCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerIsCreatorTransactor, error) {
	contract, err := bindOwnerIsCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerIsCreatorTransactor{contract: contract}, nil
}

func NewOwnerIsCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerIsCreatorFilterer, error) {
	contract, err := bindOwnerIsCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerIsCreatorFilterer{contract: contract}, nil
}

func bindOwnerIsCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerIsCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OwnerIsCreator *OwnerIsCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerIsCreator.Contract.OwnerIsCreatorCaller.contract.Call(opts, result, method, params...)
}

func (_OwnerIsCreator *OwnerIsCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.OwnerIsCreatorTransactor.contract.Transfer(opts)
}

func (_OwnerIsCreator *OwnerIsCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.OwnerIsCreatorTransactor.contract.Transact(opts, method, params...)
}

func (_OwnerIsCreator *OwnerIsCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerIsCreator.Contract.contract.Call(opts, result, method, params...)
}

func (_OwnerIsCreator *OwnerIsCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.contract.Transfer(opts)
}

func (_OwnerIsCreator *OwnerIsCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.contract.Transact(opts, method, params...)
}

func (_OwnerIsCreator *OwnerIsCreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnerIsCreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OwnerIsCreator *OwnerIsCreatorSession) Owner() (common.Address, error) {
	return _OwnerIsCreator.Contract.Owner(&_OwnerIsCreator.CallOpts)
}

func (_OwnerIsCreator *OwnerIsCreatorCallerSession) Owner() (common.Address, error) {
	return _OwnerIsCreator.Contract.Owner(&_OwnerIsCreator.CallOpts)
}

func (_OwnerIsCreator *OwnerIsCreatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerIsCreator.contract.Transact(opts, "acceptOwnership")
}

func (_OwnerIsCreator *OwnerIsCreatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.AcceptOwnership(&_OwnerIsCreator.TransactOpts)
}

func (_OwnerIsCreator *OwnerIsCreatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.AcceptOwnership(&_OwnerIsCreator.TransactOpts)
}

func (_OwnerIsCreator *OwnerIsCreatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OwnerIsCreator.contract.Transact(opts, "transferOwnership", to)
}

func (_OwnerIsCreator *OwnerIsCreatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.TransferOwnership(&_OwnerIsCreator.TransactOpts, to)
}

func (_OwnerIsCreator *OwnerIsCreatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OwnerIsCreator.Contract.TransferOwnership(&_OwnerIsCreator.TransactOpts, to)
}

type OwnerIsCreatorOwnershipTransferRequestedIterator struct {
	Event *OwnerIsCreatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnerIsCreatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerIsCreatorOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(OwnerIsCreatorOwnershipTransferRequested)
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

func (it *OwnerIsCreatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OwnerIsCreatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnerIsCreatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OwnerIsCreator *OwnerIsCreatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnerIsCreatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnerIsCreator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OwnerIsCreatorOwnershipTransferRequestedIterator{contract: _OwnerIsCreator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OwnerIsCreator *OwnerIsCreatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OwnerIsCreatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnerIsCreator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnerIsCreatorOwnershipTransferRequested)
				if err := _OwnerIsCreator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OwnerIsCreator *OwnerIsCreatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*OwnerIsCreatorOwnershipTransferRequested, error) {
	event := new(OwnerIsCreatorOwnershipTransferRequested)
	if err := _OwnerIsCreator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OwnerIsCreatorOwnershipTransferredIterator struct {
	Event *OwnerIsCreatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnerIsCreatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerIsCreatorOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(OwnerIsCreatorOwnershipTransferred)
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

func (it *OwnerIsCreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OwnerIsCreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnerIsCreatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OwnerIsCreator *OwnerIsCreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnerIsCreatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnerIsCreator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OwnerIsCreatorOwnershipTransferredIterator{contract: _OwnerIsCreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OwnerIsCreator *OwnerIsCreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnerIsCreatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnerIsCreator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnerIsCreatorOwnershipTransferred)
				if err := _OwnerIsCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OwnerIsCreator *OwnerIsCreatorFilterer) ParseOwnershipTransferred(log types.Log) (*OwnerIsCreatorOwnershipTransferred, error) {
	event := new(OwnerIsCreatorOwnershipTransferred)
	if err := _OwnerIsCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var TypeAndVersionInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

var TypeAndVersionInterfaceABI = TypeAndVersionInterfaceMetaData.ABI

type TypeAndVersionInterface struct {
	TypeAndVersionInterfaceCaller
	TypeAndVersionInterfaceTransactor
	TypeAndVersionInterfaceFilterer
}

type TypeAndVersionInterfaceCaller struct {
	contract *bind.BoundContract
}

type TypeAndVersionInterfaceTransactor struct {
	contract *bind.BoundContract
}

type TypeAndVersionInterfaceFilterer struct {
	contract *bind.BoundContract
}

type TypeAndVersionInterfaceSession struct {
	Contract     *TypeAndVersionInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TypeAndVersionInterfaceCallerSession struct {
	Contract *TypeAndVersionInterfaceCaller
	CallOpts bind.CallOpts
}

type TypeAndVersionInterfaceTransactorSession struct {
	Contract     *TypeAndVersionInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type TypeAndVersionInterfaceRaw struct {
	Contract *TypeAndVersionInterface
}

type TypeAndVersionInterfaceCallerRaw struct {
	Contract *TypeAndVersionInterfaceCaller
}

type TypeAndVersionInterfaceTransactorRaw struct {
	Contract *TypeAndVersionInterfaceTransactor
}

func NewTypeAndVersionInterface(address common.Address, backend bind.ContractBackend) (*TypeAndVersionInterface, error) {
	contract, err := bindTypeAndVersionInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TypeAndVersionInterface{TypeAndVersionInterfaceCaller: TypeAndVersionInterfaceCaller{contract: contract}, TypeAndVersionInterfaceTransactor: TypeAndVersionInterfaceTransactor{contract: contract}, TypeAndVersionInterfaceFilterer: TypeAndVersionInterfaceFilterer{contract: contract}}, nil
}

func NewTypeAndVersionInterfaceCaller(address common.Address, caller bind.ContractCaller) (*TypeAndVersionInterfaceCaller, error) {
	contract, err := bindTypeAndVersionInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypeAndVersionInterfaceCaller{contract: contract}, nil
}

func NewTypeAndVersionInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TypeAndVersionInterfaceTransactor, error) {
	contract, err := bindTypeAndVersionInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypeAndVersionInterfaceTransactor{contract: contract}, nil
}

func NewTypeAndVersionInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TypeAndVersionInterfaceFilterer, error) {
	contract, err := bindTypeAndVersionInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypeAndVersionInterfaceFilterer{contract: contract}, nil
}

func bindTypeAndVersionInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypeAndVersionInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypeAndVersionInterface.Contract.TypeAndVersionInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypeAndVersionInterface.Contract.TypeAndVersionInterfaceTransactor.contract.Transfer(opts)
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypeAndVersionInterface.Contract.TypeAndVersionInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypeAndVersionInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypeAndVersionInterface.Contract.contract.Transfer(opts)
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypeAndVersionInterface.Contract.contract.Transact(opts, method, params...)
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TypeAndVersionInterface.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceSession) TypeAndVersion() (string, error) {
	return _TypeAndVersionInterface.Contract.TypeAndVersion(&_TypeAndVersionInterface.CallOpts)
}

func (_TypeAndVersionInterface *TypeAndVersionInterfaceCallerSession) TypeAndVersion() (string, error) {
	return _TypeAndVersionInterface.Contract.TypeAndVersion(&_TypeAndVersionInterface.CallOpts)
}

var VRFBeaconTypesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50605780601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80632f7527cc14602d575b600080fd5b6034600881565b60405160ff909116815260200160405180910390f3fea164736f6c634300080f000a",
}

var VRFBeaconTypesABI = VRFBeaconTypesMetaData.ABI

var VRFBeaconTypesBin = VRFBeaconTypesMetaData.Bin

func DeployVRFBeaconTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFBeaconTypes, error) {
	parsed, err := VRFBeaconTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeaconTypes{VRFBeaconTypesCaller: VRFBeaconTypesCaller{contract: contract}, VRFBeaconTypesTransactor: VRFBeaconTypesTransactor{contract: contract}, VRFBeaconTypesFilterer: VRFBeaconTypesFilterer{contract: contract}}, nil
}

type VRFBeaconTypes struct {
	VRFBeaconTypesCaller
	VRFBeaconTypesTransactor
	VRFBeaconTypesFilterer
}

type VRFBeaconTypesCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconTypesTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconTypesFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconTypesSession struct {
	Contract     *VRFBeaconTypes
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconTypesCallerSession struct {
	Contract *VRFBeaconTypesCaller
	CallOpts bind.CallOpts
}

type VRFBeaconTypesTransactorSession struct {
	Contract     *VRFBeaconTypesTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconTypesRaw struct {
	Contract *VRFBeaconTypes
}

type VRFBeaconTypesCallerRaw struct {
	Contract *VRFBeaconTypesCaller
}

type VRFBeaconTypesTransactorRaw struct {
	Contract *VRFBeaconTypesTransactor
}

func NewVRFBeaconTypes(address common.Address, backend bind.ContractBackend) (*VRFBeaconTypes, error) {
	contract, err := bindVRFBeaconTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypes{VRFBeaconTypesCaller: VRFBeaconTypesCaller{contract: contract}, VRFBeaconTypesTransactor: VRFBeaconTypesTransactor{contract: contract}, VRFBeaconTypesFilterer: VRFBeaconTypesFilterer{contract: contract}}, nil
}

func NewVRFBeaconTypesCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconTypesCaller, error) {
	contract, err := bindVRFBeaconTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesCaller{contract: contract}, nil
}

func NewVRFBeaconTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconTypesTransactor, error) {
	contract, err := bindVRFBeaconTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesTransactor{contract: contract}, nil
}

func NewVRFBeaconTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconTypesFilterer, error) {
	contract, err := bindVRFBeaconTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesFilterer{contract: contract}, nil
}

func bindVRFBeaconTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeaconTypes *VRFBeaconTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconTypes.Contract.VRFBeaconTypesCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconTypes *VRFBeaconTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconTypes.Contract.VRFBeaconTypesTransactor.contract.Transfer(opts)
}

func (_VRFBeaconTypes *VRFBeaconTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconTypes.Contract.VRFBeaconTypesTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeaconTypes *VRFBeaconTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconTypes.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconTypes *VRFBeaconTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconTypes.Contract.contract.Transfer(opts)
}

func (_VRFBeaconTypes *VRFBeaconTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconTypes.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeaconTypes *VRFBeaconTypesCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFBeaconTypes.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFBeaconTypes *VRFBeaconTypesSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconTypes.Contract.NUMCONFDELAYS(&_VRFBeaconTypes.CallOpts)
}

func (_VRFBeaconTypes *VRFBeaconTypesCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconTypes.Contract.NUMCONFDELAYS(&_VRFBeaconTypes.CallOpts)
}

type VRFBeaconTypesConfigSetIterator struct {
	Event *VRFBeaconTypesConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconTypesConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconTypesConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconTypesConfigSet)
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

func (it *VRFBeaconTypesConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconTypesConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconTypesConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFBeaconTypesConfigSetIterator, error) {

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesConfigSetIterator{contract: _VRFBeaconTypes.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconTypesConfigSet)
				if err := _VRFBeaconTypes.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) ParseConfigSet(log types.Log) (*VRFBeaconTypesConfigSet, error) {
	event := new(VRFBeaconTypesConfigSet)
	if err := _VRFBeaconTypes.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconTypesNewTransmissionIterator struct {
	Event *VRFBeaconTypesNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconTypesNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconTypesNewTransmission)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconTypesNewTransmission)
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

func (it *VRFBeaconTypesNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconTypesNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconTypesNewTransmission struct {
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*VRFBeaconTypesNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesNewTransmissionIterator{contract: _VRFBeaconTypes.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconTypesNewTransmission)
				if err := _VRFBeaconTypes.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) ParseNewTransmission(log types.Log) (*VRFBeaconTypesNewTransmission, error) {
	event := new(VRFBeaconTypesNewTransmission)
	if err := _VRFBeaconTypes.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconTypesOutputsServedIterator struct {
	Event *VRFBeaconTypesOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconTypesOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconTypesOutputsServed)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconTypesOutputsServed)
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

func (it *VRFBeaconTypesOutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconTypesOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconTypesOutputsServed struct {
	RecentBlockHeight  uint64
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFBeaconTypesOutputsServedIterator, error) {

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesOutputsServedIterator{contract: _VRFBeaconTypes.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesOutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconTypesOutputsServed)
				if err := _VRFBeaconTypes.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) ParseOutputsServed(log types.Log) (*VRFBeaconTypesOutputsServed, error) {
	event := new(VRFBeaconTypesOutputsServed)
	if err := _VRFBeaconTypes.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconTypesRandomWordsFulfilledIterator struct {
	Event *VRFBeaconTypesRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconTypesRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconTypesRandomWordsFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconTypesRandomWordsFulfilled)
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

func (it *VRFBeaconTypesRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconTypesRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconTypesRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFBeaconTypesRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesRandomWordsFulfilledIterator{contract: _VRFBeaconTypes.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconTypesRandomWordsFulfilled)
				if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFBeaconTypesRandomWordsFulfilled, error) {
	event := new(VRFBeaconTypesRandomWordsFulfilled)
	if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconTypesRandomnessFulfillmentRequestedIterator struct {
	Event *VRFBeaconTypesRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconTypesRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconTypesRandomnessFulfillmentRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconTypesRandomnessFulfillmentRequested)
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

func (it *VRFBeaconTypesRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconTypesRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconTypesRandomnessFulfillmentRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	GasAllowance           uint32
	GasPrice               *big.Int
	WeiPerUnitLink         *big.Int
	Arguments              []byte
	Raw                    types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconTypesRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconTypes.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconTypesRandomnessFulfillmentRequested)
				if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconTypesRandomnessFulfillmentRequested, error) {
	event := new(VRFBeaconTypesRandomnessFulfillmentRequested)
	if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconTypesRandomnessRequestedIterator struct {
	Event *VRFBeaconTypesRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconTypesRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconTypesRandomnessRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconTypesRandomnessRequested)
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

func (it *VRFBeaconTypesRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconTypesRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconTypesRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	Raw                    types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconTypesRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesRandomnessRequestedIterator{contract: _VRFBeaconTypes.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconTypesRandomnessRequested)
				if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) ParseRandomnessRequested(log types.Log) (*VRFBeaconTypesRandomnessRequested, error) {
	event := new(VRFBeaconTypesRandomnessRequested)
	if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFMigratableCoordinatorInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var VRFMigratableCoordinatorInterfaceABI = VRFMigratableCoordinatorInterfaceMetaData.ABI

type VRFMigratableCoordinatorInterface struct {
	VRFMigratableCoordinatorInterfaceCaller
	VRFMigratableCoordinatorInterfaceTransactor
	VRFMigratableCoordinatorInterfaceFilterer
}

type VRFMigratableCoordinatorInterfaceCaller struct {
	contract *bind.BoundContract
}

type VRFMigratableCoordinatorInterfaceTransactor struct {
	contract *bind.BoundContract
}

type VRFMigratableCoordinatorInterfaceFilterer struct {
	contract *bind.BoundContract
}

type VRFMigratableCoordinatorInterfaceSession struct {
	Contract     *VRFMigratableCoordinatorInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFMigratableCoordinatorInterfaceCallerSession struct {
	Contract *VRFMigratableCoordinatorInterfaceCaller
	CallOpts bind.CallOpts
}

type VRFMigratableCoordinatorInterfaceTransactorSession struct {
	Contract     *VRFMigratableCoordinatorInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type VRFMigratableCoordinatorInterfaceRaw struct {
	Contract *VRFMigratableCoordinatorInterface
}

type VRFMigratableCoordinatorInterfaceCallerRaw struct {
	Contract *VRFMigratableCoordinatorInterfaceCaller
}

type VRFMigratableCoordinatorInterfaceTransactorRaw struct {
	Contract *VRFMigratableCoordinatorInterfaceTransactor
}

func NewVRFMigratableCoordinatorInterface(address common.Address, backend bind.ContractBackend) (*VRFMigratableCoordinatorInterface, error) {
	contract, err := bindVRFMigratableCoordinatorInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterface{VRFMigratableCoordinatorInterfaceCaller: VRFMigratableCoordinatorInterfaceCaller{contract: contract}, VRFMigratableCoordinatorInterfaceTransactor: VRFMigratableCoordinatorInterfaceTransactor{contract: contract}, VRFMigratableCoordinatorInterfaceFilterer: VRFMigratableCoordinatorInterfaceFilterer{contract: contract}}, nil
}

func NewVRFMigratableCoordinatorInterfaceCaller(address common.Address, caller bind.ContractCaller) (*VRFMigratableCoordinatorInterfaceCaller, error) {
	contract, err := bindVRFMigratableCoordinatorInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceCaller{contract: contract}, nil
}

func NewVRFMigratableCoordinatorInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFMigratableCoordinatorInterfaceTransactor, error) {
	contract, err := bindVRFMigratableCoordinatorInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceTransactor{contract: contract}, nil
}

func NewVRFMigratableCoordinatorInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFMigratableCoordinatorInterfaceFilterer, error) {
	contract, err := bindVRFMigratableCoordinatorInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceFilterer{contract: contract}, nil
}

func bindVRFMigratableCoordinatorInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFMigratableCoordinatorInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFMigratableCoordinatorInterface.Contract.VRFMigratableCoordinatorInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.VRFMigratableCoordinatorInterfaceTransactor.contract.Transfer(opts)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.VRFMigratableCoordinatorInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFMigratableCoordinatorInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.contract.Transfer(opts)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.contract.Transact(opts, method, params...)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFMigratableCoordinatorInterface.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFMigratableCoordinatorInterface.Contract.NUMCONFDELAYS(&_VRFMigratableCoordinatorInterface.CallOpts)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFMigratableCoordinatorInterface.Contract.NUMCONFDELAYS(&_VRFMigratableCoordinatorInterface.CallOpts)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactor) RedeemRandomness(opts *bind.TransactOpts, requester common.Address, subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.contract.Transact(opts, "redeemRandomness", requester, subID, requestID, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceSession) RedeemRandomness(requester common.Address, subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.RedeemRandomness(&_VRFMigratableCoordinatorInterface.TransactOpts, requester, subID, requestID, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactorSession) RedeemRandomness(requester common.Address, subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.RedeemRandomness(&_VRFMigratableCoordinatorInterface.TransactOpts, requester, subID, requestID, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactor) RequestRandomness(opts *bind.TransactOpts, requester common.Address, subID *big.Int, numWords uint16, confirmationDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.contract.Transact(opts, "requestRandomness", requester, subID, numWords, confirmationDelay, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceSession) RequestRandomness(requester common.Address, subID *big.Int, numWords uint16, confirmationDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.RequestRandomness(&_VRFMigratableCoordinatorInterface.TransactOpts, requester, subID, numWords, confirmationDelay, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactorSession) RequestRandomness(requester common.Address, subID *big.Int, numWords uint16, confirmationDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.RequestRandomness(&_VRFMigratableCoordinatorInterface.TransactOpts, requester, subID, numWords, confirmationDelay, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, requester common.Address, subID *big.Int, numWords uint16, confirmationDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.contract.Transact(opts, "requestRandomnessFulfillment", requester, subID, numWords, confirmationDelay, callbackGasLimit, arguments, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceSession) RequestRandomnessFulfillment(requester common.Address, subID *big.Int, numWords uint16, confirmationDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.RequestRandomnessFulfillment(&_VRFMigratableCoordinatorInterface.TransactOpts, requester, subID, numWords, confirmationDelay, callbackGasLimit, arguments, extraArgs)
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceTransactorSession) RequestRandomnessFulfillment(requester common.Address, subID *big.Int, numWords uint16, confirmationDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _VRFMigratableCoordinatorInterface.Contract.RequestRandomnessFulfillment(&_VRFMigratableCoordinatorInterface.TransactOpts, requester, subID, numWords, confirmationDelay, callbackGasLimit, arguments, extraArgs)
}

type VRFMigratableCoordinatorInterfaceConfigSetIterator struct {
	Event *VRFMigratableCoordinatorInterfaceConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMigratableCoordinatorInterfaceConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMigratableCoordinatorInterfaceConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFMigratableCoordinatorInterfaceConfigSet)
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

func (it *VRFMigratableCoordinatorInterfaceConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFMigratableCoordinatorInterfaceConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMigratableCoordinatorInterfaceConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFMigratableCoordinatorInterfaceConfigSetIterator, error) {

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceConfigSetIterator{contract: _VRFMigratableCoordinatorInterface.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFMigratableCoordinatorInterfaceConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMigratableCoordinatorInterfaceConfigSet)
				if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) ParseConfigSet(log types.Log) (*VRFMigratableCoordinatorInterfaceConfigSet, error) {
	event := new(VRFMigratableCoordinatorInterfaceConfigSet)
	if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFMigratableCoordinatorInterfaceNewTransmissionIterator struct {
	Event *VRFMigratableCoordinatorInterfaceNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMigratableCoordinatorInterfaceNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMigratableCoordinatorInterfaceNewTransmission)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFMigratableCoordinatorInterfaceNewTransmission)
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

func (it *VRFMigratableCoordinatorInterfaceNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFMigratableCoordinatorInterfaceNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMigratableCoordinatorInterfaceNewTransmission struct {
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*VRFMigratableCoordinatorInterfaceNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceNewTransmissionIterator{contract: _VRFMigratableCoordinatorInterface.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFMigratableCoordinatorInterfaceNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMigratableCoordinatorInterfaceNewTransmission)
				if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) ParseNewTransmission(log types.Log) (*VRFMigratableCoordinatorInterfaceNewTransmission, error) {
	event := new(VRFMigratableCoordinatorInterfaceNewTransmission)
	if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFMigratableCoordinatorInterfaceOutputsServedIterator struct {
	Event *VRFMigratableCoordinatorInterfaceOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMigratableCoordinatorInterfaceOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMigratableCoordinatorInterfaceOutputsServed)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFMigratableCoordinatorInterfaceOutputsServed)
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

func (it *VRFMigratableCoordinatorInterfaceOutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFMigratableCoordinatorInterfaceOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMigratableCoordinatorInterfaceOutputsServed struct {
	RecentBlockHeight  uint64
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFMigratableCoordinatorInterfaceOutputsServedIterator, error) {

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceOutputsServedIterator{contract: _VRFMigratableCoordinatorInterface.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFMigratableCoordinatorInterfaceOutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMigratableCoordinatorInterfaceOutputsServed)
				if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) ParseOutputsServed(log types.Log) (*VRFMigratableCoordinatorInterfaceOutputsServed, error) {
	event := new(VRFMigratableCoordinatorInterfaceOutputsServed)
	if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFMigratableCoordinatorInterfaceRandomWordsFulfilledIterator struct {
	Event *VRFMigratableCoordinatorInterfaceRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMigratableCoordinatorInterfaceRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMigratableCoordinatorInterfaceRandomWordsFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFMigratableCoordinatorInterfaceRandomWordsFulfilled)
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

func (it *VRFMigratableCoordinatorInterfaceRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFMigratableCoordinatorInterfaceRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMigratableCoordinatorInterfaceRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFMigratableCoordinatorInterfaceRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceRandomWordsFulfilledIterator{contract: _VRFMigratableCoordinatorInterface.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFMigratableCoordinatorInterfaceRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMigratableCoordinatorInterfaceRandomWordsFulfilled)
				if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFMigratableCoordinatorInterfaceRandomWordsFulfilled, error) {
	event := new(VRFMigratableCoordinatorInterfaceRandomWordsFulfilled)
	if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequestedIterator struct {
	Event *VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested)
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

func (it *VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	GasAllowance           uint32
	GasPrice               *big.Int
	WeiPerUnitLink         *big.Int
	Arguments              []byte
	Raw                    types.Log
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequestedIterator{contract: _VRFMigratableCoordinatorInterface.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested)
				if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested, error) {
	event := new(VRFMigratableCoordinatorInterfaceRandomnessFulfillmentRequested)
	if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFMigratableCoordinatorInterfaceRandomnessRequestedIterator struct {
	Event *VRFMigratableCoordinatorInterfaceRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFMigratableCoordinatorInterfaceRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFMigratableCoordinatorInterfaceRandomnessRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFMigratableCoordinatorInterfaceRandomnessRequested)
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

func (it *VRFMigratableCoordinatorInterfaceRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFMigratableCoordinatorInterfaceRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFMigratableCoordinatorInterfaceRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	Raw                    types.Log
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFMigratableCoordinatorInterfaceRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFMigratableCoordinatorInterfaceRandomnessRequestedIterator{contract: _VRFMigratableCoordinatorInterface.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFMigratableCoordinatorInterfaceRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFMigratableCoordinatorInterface.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFMigratableCoordinatorInterfaceRandomnessRequested)
				if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFMigratableCoordinatorInterface *VRFMigratableCoordinatorInterfaceFilterer) ParseRandomnessRequested(log types.Log) (*VRFMigratableCoordinatorInterfaceRandomnessRequested, error) {
	event := new(VRFMigratableCoordinatorInterfaceRandomnessRequested)
	if err := _VRFMigratableCoordinatorInterface.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFMigrationInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractVRFMigrationInterface\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedRequest\",\"type\":\"bytes\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedData\",\"type\":\"bytes\"}],\"name\":\"onMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var VRFMigrationInterfaceABI = VRFMigrationInterfaceMetaData.ABI

type VRFMigrationInterface struct {
	VRFMigrationInterfaceCaller
	VRFMigrationInterfaceTransactor
	VRFMigrationInterfaceFilterer
}

type VRFMigrationInterfaceCaller struct {
	contract *bind.BoundContract
}

type VRFMigrationInterfaceTransactor struct {
	contract *bind.BoundContract
}

type VRFMigrationInterfaceFilterer struct {
	contract *bind.BoundContract
}

type VRFMigrationInterfaceSession struct {
	Contract     *VRFMigrationInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFMigrationInterfaceCallerSession struct {
	Contract *VRFMigrationInterfaceCaller
	CallOpts bind.CallOpts
}

type VRFMigrationInterfaceTransactorSession struct {
	Contract     *VRFMigrationInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type VRFMigrationInterfaceRaw struct {
	Contract *VRFMigrationInterface
}

type VRFMigrationInterfaceCallerRaw struct {
	Contract *VRFMigrationInterfaceCaller
}

type VRFMigrationInterfaceTransactorRaw struct {
	Contract *VRFMigrationInterfaceTransactor
}

func NewVRFMigrationInterface(address common.Address, backend bind.ContractBackend) (*VRFMigrationInterface, error) {
	contract, err := bindVRFMigrationInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFMigrationInterface{VRFMigrationInterfaceCaller: VRFMigrationInterfaceCaller{contract: contract}, VRFMigrationInterfaceTransactor: VRFMigrationInterfaceTransactor{contract: contract}, VRFMigrationInterfaceFilterer: VRFMigrationInterfaceFilterer{contract: contract}}, nil
}

func NewVRFMigrationInterfaceCaller(address common.Address, caller bind.ContractCaller) (*VRFMigrationInterfaceCaller, error) {
	contract, err := bindVRFMigrationInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFMigrationInterfaceCaller{contract: contract}, nil
}

func NewVRFMigrationInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFMigrationInterfaceTransactor, error) {
	contract, err := bindVRFMigrationInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFMigrationInterfaceTransactor{contract: contract}, nil
}

func NewVRFMigrationInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFMigrationInterfaceFilterer, error) {
	contract, err := bindVRFMigrationInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFMigrationInterfaceFilterer{contract: contract}, nil
}

func bindVRFMigrationInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFMigrationInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFMigrationInterface *VRFMigrationInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFMigrationInterface.Contract.VRFMigrationInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.VRFMigrationInterfaceTransactor.contract.Transfer(opts)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.VRFMigrationInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFMigrationInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.contract.Transfer(opts)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.contract.Transact(opts, method, params...)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceCaller) MigrationVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFMigrationInterface.contract.Call(opts, &out, "migrationVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFMigrationInterface *VRFMigrationInterfaceSession) MigrationVersion() (uint8, error) {
	return _VRFMigrationInterface.Contract.MigrationVersion(&_VRFMigrationInterface.CallOpts)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceCallerSession) MigrationVersion() (uint8, error) {
	return _VRFMigrationInterface.Contract.MigrationVersion(&_VRFMigrationInterface.CallOpts)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceTransactor) Migrate(opts *bind.TransactOpts, newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _VRFMigrationInterface.contract.Transact(opts, "migrate", newCoordinator, encodedRequest)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceSession) Migrate(newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.Migrate(&_VRFMigrationInterface.TransactOpts, newCoordinator, encodedRequest)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceTransactorSession) Migrate(newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.Migrate(&_VRFMigrationInterface.TransactOpts, newCoordinator, encodedRequest)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceTransactor) OnMigration(opts *bind.TransactOpts, encodedData []byte) (*types.Transaction, error) {
	return _VRFMigrationInterface.contract.Transact(opts, "onMigration", encodedData)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceSession) OnMigration(encodedData []byte) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.OnMigration(&_VRFMigrationInterface.TransactOpts, encodedData)
}

func (_VRFMigrationInterface *VRFMigrationInterfaceTransactorSession) OnMigration(encodedData []byte) (*types.Transaction, error) {
	return _VRFMigrationInterface.Contract.OnMigration(&_VRFMigrationInterface.TransactOpts, encodedData)
}

var VRFRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CoordinatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CoordinatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedMigrationVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"RouteSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callWithExactGasEvenIfTargetIsNoContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"sufficientGas\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"deregisterCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCoordinators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"getRoute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"isCoordinatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"registerCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"resetRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"setRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b611245806101576000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80639479b74e11610097578063d9c4a44d11610066578063d9c4a44d1461027b578063db972c8b1461028e578063f2fde38b146102a1578063fa66358a146102b457600080fd5b80639479b74e14610212578063a8a254ed14610225578063bb9b2b3414610245578063cef40aa61461025857600080fd5b806355fe9763116100d357806355fe9763146101a65780635d5d8d19146101d057806379ba5097146101e55780638da5cb5b146101ed57600080fd5b8063181f5a77146101055780632f7527cc1461014d5780634b2407d4146101675780634ffac83a1461017c575b600080fd5b604080518082018252600f81527f565246526f7574657220312e302e300000000000000000000000000000000000602082015290516101449190610cbf565b60405180910390f35b610155600881565b60405160ff9091168152602001610144565b61017a610175366004610cee565b6102c7565b005b61018f61018a366004610de5565b610401565b60405165ffffffffffff9091168152602001610144565b6101b96101b4366004610e4d565b6104aa565b604080519215158352901515602083015201610144565b6101d8610513565b6040516101449190610ea4565b61017a610524565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610144565b61017a610220366004610ef1565b6105e7565b610238610233366004610f1e565b610679565b6040516101449190610f45565b61017a610253366004610ef1565b610724565b61026b610266366004610cee565b6107a7565b6040519015158152602001610144565b61017a610289366004610cee565b6107c9565b61018f61029c366004610f7d565b61083f565b61017a6102af366004610cee565b6108ee565b6101fa6102c2366004610ef1565b610902565b6102cf610966565b6102da6003826109c2565b15610311576040517fdcecb7bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000819050806001600160a01b031663294daa496040518163ffffffff1660e01b8152600401602060405180830381865afa158015610354573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103789190611026565b60ff166000036103b4576040517f6e81684800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103bf6003836109e7565b506040516001600160a01b03831681527fb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625906020015b60405180910390a15050565b60008061040d86610902565b6040517f62f8b6200000000000000000000000000000000000000000000000000000000081529091506001600160a01b038216906362f8b6209061045d9033908a908a908a908a90600401611049565b6020604051808303816000875af115801561047c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a09190611091565b9695505050505050565b600080336104b96003826109c2565b6104d6576040516301fd70a160e51b815260040160405180910390fd5b5a611388811061050957611388810390508660408204820311156105095760008086516020880160008a8cf19350600192505b5050935093915050565b606061051f60036109fc565b905090565b6001546001600160a01b031633146105835760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b336105f36003826109c2565b610610576040516301fd70a160e51b815260040160405180910390fd5b600082815260026020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff191633908117909155915191825283917fd8bb190f51a471c0cc88b5df464ee10c29138cce0d70fb36c472d60b414f3b3a91015b60405180910390a25050565b6060600061068685610902565b6040517fabbf1c9b0000000000000000000000000000000000000000000000000000000081529091506001600160a01b0382169063abbf1c9b906106d49033908990899089906004016110ae565b6000604051808303816000875af11580156106f3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261071b91908101906110e4565b95945050505050565b336107306003826109c2565b61074d576040516301fd70a160e51b815260040160405180910390fd5b6000828152600260209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191690555191825283917fd8bb190f51a471c0cc88b5df464ee10c29138cce0d70fb36c472d60b414f3b3a910161066d565b60006107b46003836109c2565b156107c157506001919050565b506000919050565b6107d1610966565b806107dd6003826109c2565b6107fa576040516301fd70a160e51b815260040160405180910390fd5b610805600383610a09565b506040516001600160a01b03831681527ff80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af37906020016103f5565b60008061084b88610902565b6040517f95009f080000000000000000000000000000000000000000000000000000000081529091506001600160a01b038216906395009f089061089f9033908c908c908c908c908c908c9060040161117e565b6020604051808303816000875af11580156108be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e29190611091565b98975050505050505050565b6108f6610966565b6108ff81610a1e565b50565b6000818152600260205260408120546001600160a01b031680610938576040516302d703d160e61b815260040160405180910390fd5b6109436003826109c2565b610960576040516302d703d160e61b815260040160405180910390fd5b92915050565b6000546001600160a01b031633146109c05760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161057a565b565b6001600160a01b038116600090815260018301602052604081205415155b9392505050565b60006109e0836001600160a01b038416610ad4565b606060006109e083610b23565b60006109e0836001600160a01b038416610b7f565b336001600160a01b03821603610a765760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161057a565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600183016020526040812054610b1b57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610960565b506000610960565b606081600001805480602002602001604051908101604052809291908181526020018280548015610b7357602002820191906000526020600020905b815481526020019060010190808311610b5f575b50505050509050919050565b60008181526001830160205260408120548015610c68576000610ba36001836111e7565b8554909150600090610bb7906001906111e7565b9050818114610c1c576000866000018281548110610bd757610bd761120c565b9060005260206000200154905080876000018481548110610bfa57610bfa61120c565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610c2d57610c2d611222565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610960565b6000915050610960565b6000815180845260005b81811015610c9857602081850181015186830182015201610c7c565b81811115610caa576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006109e06020830184610c72565b80356001600160a01b0381168114610ce957600080fd5b919050565b600060208284031215610d0057600080fd5b6109e082610cd2565b803561ffff81168114610ce957600080fd5b803562ffffff81168114610ce957600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610d6d57610d6d610d2e565b604052919050565b600082601f830112610d8657600080fd5b813567ffffffffffffffff811115610da057610da0610d2e565b610db3601f8201601f1916602001610d44565b818152846020838601011115610dc857600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215610dfb57600080fd5b84359350610e0b60208601610d09565b9250610e1960408601610d1b565b9150606085013567ffffffffffffffff811115610e3557600080fd5b610e4187828801610d75565b91505092959194509250565b600080600060608486031215610e6257600080fd5b83359250610e7260208501610cd2565b9150604084013567ffffffffffffffff811115610e8e57600080fd5b610e9a86828701610d75565b9150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015610ee55783516001600160a01b031683529284019291840191600101610ec0565b50909695505050505050565b600060208284031215610f0357600080fd5b5035919050565b65ffffffffffff811681146108ff57600080fd5b600080600060608486031215610f3357600080fd5b833592506020840135610e7281610f0a565b6020808252825182820181905260009190848201906040850190845b81811015610ee557835183529284019291840191600101610f61565b60008060008060008060c08789031215610f9657600080fd5b86359550610fa660208801610d09565b9450610fb460408801610d1b565b9350606087013563ffffffff81168114610fcd57600080fd5b9250608087013567ffffffffffffffff80821115610fea57600080fd5b610ff68a838b01610d75565b935060a089013591508082111561100c57600080fd5b5061101989828a01610d75565b9150509295509295509295565b60006020828403121561103857600080fd5b815160ff811681146109e057600080fd5b6001600160a01b038616815284602082015261ffff8416604082015262ffffff8316606082015260a06080820152600061108660a0830184610c72565b979650505050505050565b6000602082840312156110a357600080fd5b81516109e081610f0a565b6001600160a01b038516815283602082015265ffffffffffff831660408201526080606082015260006104a06080830184610c72565b600060208083850312156110f757600080fd5b825167ffffffffffffffff8082111561110f57600080fd5b818501915085601f83011261112357600080fd5b81518181111561113557611135610d2e565b8060051b9150611146848301610d44565b818152918301840191848101908884111561116057600080fd5b938501935b838510156108e257845182529385019390850190611165565b6001600160a01b038816815286602082015261ffff8616604082015262ffffff8516606082015263ffffffff8416608082015260e060a082015260006111c760e0830185610c72565b82810360c08401526111d98185610c72565b9a9950505050505050505050565b60008282101561120757634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fdfea164736f6c634300080f000a",
}

var VRFRouterABI = VRFRouterMetaData.ABI

var VRFRouterBin = VRFRouterMetaData.Bin

func DeployVRFRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFRouter, error) {
	parsed, err := VRFRouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFRouter{VRFRouterCaller: VRFRouterCaller{contract: contract}, VRFRouterTransactor: VRFRouterTransactor{contract: contract}, VRFRouterFilterer: VRFRouterFilterer{contract: contract}}, nil
}

type VRFRouter struct {
	VRFRouterCaller
	VRFRouterTransactor
	VRFRouterFilterer
}

type VRFRouterCaller struct {
	contract *bind.BoundContract
}

type VRFRouterTransactor struct {
	contract *bind.BoundContract
}

type VRFRouterFilterer struct {
	contract *bind.BoundContract
}

type VRFRouterSession struct {
	Contract     *VRFRouter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFRouterCallerSession struct {
	Contract *VRFRouterCaller
	CallOpts bind.CallOpts
}

type VRFRouterTransactorSession struct {
	Contract     *VRFRouterTransactor
	TransactOpts bind.TransactOpts
}

type VRFRouterRaw struct {
	Contract *VRFRouter
}

type VRFRouterCallerRaw struct {
	Contract *VRFRouterCaller
}

type VRFRouterTransactorRaw struct {
	Contract *VRFRouterTransactor
}

func NewVRFRouter(address common.Address, backend bind.ContractBackend) (*VRFRouter, error) {
	contract, err := bindVRFRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFRouter{VRFRouterCaller: VRFRouterCaller{contract: contract}, VRFRouterTransactor: VRFRouterTransactor{contract: contract}, VRFRouterFilterer: VRFRouterFilterer{contract: contract}}, nil
}

func NewVRFRouterCaller(address common.Address, caller bind.ContractCaller) (*VRFRouterCaller, error) {
	contract, err := bindVRFRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFRouterCaller{contract: contract}, nil
}

func NewVRFRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFRouterTransactor, error) {
	contract, err := bindVRFRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFRouterTransactor{contract: contract}, nil
}

func NewVRFRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFRouterFilterer, error) {
	contract, err := bindVRFRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFRouterFilterer{contract: contract}, nil
}

func bindVRFRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFRouter *VRFRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFRouter.Contract.VRFRouterCaller.contract.Call(opts, result, method, params...)
}

func (_VRFRouter *VRFRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFRouter.Contract.VRFRouterTransactor.contract.Transfer(opts)
}

func (_VRFRouter *VRFRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFRouter.Contract.VRFRouterTransactor.contract.Transact(opts, method, params...)
}

func (_VRFRouter *VRFRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFRouter.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFRouter *VRFRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFRouter.Contract.contract.Transfer(opts)
}

func (_VRFRouter *VRFRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFRouter.Contract.contract.Transact(opts, method, params...)
}

func (_VRFRouter *VRFRouterCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFRouter.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFRouter *VRFRouterSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFRouter.Contract.NUMCONFDELAYS(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFRouter.Contract.NUMCONFDELAYS(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterCaller) GetCoordinators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VRFRouter.contract.Call(opts, &out, "getCoordinators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_VRFRouter *VRFRouterSession) GetCoordinators() ([]common.Address, error) {
	return _VRFRouter.Contract.GetCoordinators(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterCallerSession) GetCoordinators() ([]common.Address, error) {
	return _VRFRouter.Contract.GetCoordinators(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterCaller) GetRoute(opts *bind.CallOpts, subID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VRFRouter.contract.Call(opts, &out, "getRoute", subID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFRouter *VRFRouterSession) GetRoute(subID *big.Int) (common.Address, error) {
	return _VRFRouter.Contract.GetRoute(&_VRFRouter.CallOpts, subID)
}

func (_VRFRouter *VRFRouterCallerSession) GetRoute(subID *big.Int) (common.Address, error) {
	return _VRFRouter.Contract.GetRoute(&_VRFRouter.CallOpts, subID)
}

func (_VRFRouter *VRFRouterCaller) IsCoordinatorRegistered(opts *bind.CallOpts, coordinatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _VRFRouter.contract.Call(opts, &out, "isCoordinatorRegistered", coordinatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFRouter *VRFRouterSession) IsCoordinatorRegistered(coordinatorAddress common.Address) (bool, error) {
	return _VRFRouter.Contract.IsCoordinatorRegistered(&_VRFRouter.CallOpts, coordinatorAddress)
}

func (_VRFRouter *VRFRouterCallerSession) IsCoordinatorRegistered(coordinatorAddress common.Address) (bool, error) {
	return _VRFRouter.Contract.IsCoordinatorRegistered(&_VRFRouter.CallOpts, coordinatorAddress)
}

func (_VRFRouter *VRFRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFRouter *VRFRouterSession) Owner() (common.Address, error) {
	return _VRFRouter.Contract.Owner(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterCallerSession) Owner() (common.Address, error) {
	return _VRFRouter.Contract.Owner(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFRouter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFRouter *VRFRouterSession) TypeAndVersion() (string, error) {
	return _VRFRouter.Contract.TypeAndVersion(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterCallerSession) TypeAndVersion() (string, error) {
	return _VRFRouter.Contract.TypeAndVersion(&_VRFRouter.CallOpts)
}

func (_VRFRouter *VRFRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "acceptOwnership")
}

func (_VRFRouter *VRFRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFRouter.Contract.AcceptOwnership(&_VRFRouter.TransactOpts)
}

func (_VRFRouter *VRFRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFRouter.Contract.AcceptOwnership(&_VRFRouter.TransactOpts)
}

func (_VRFRouter *VRFRouterTransactor) CallWithExactGasEvenIfTargetIsNoContract(opts *bind.TransactOpts, gasAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "callWithExactGasEvenIfTargetIsNoContract", gasAmount, target, data)
}

func (_VRFRouter *VRFRouterSession) CallWithExactGasEvenIfTargetIsNoContract(gasAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.CallWithExactGasEvenIfTargetIsNoContract(&_VRFRouter.TransactOpts, gasAmount, target, data)
}

func (_VRFRouter *VRFRouterTransactorSession) CallWithExactGasEvenIfTargetIsNoContract(gasAmount *big.Int, target common.Address, data []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.CallWithExactGasEvenIfTargetIsNoContract(&_VRFRouter.TransactOpts, gasAmount, target, data)
}

func (_VRFRouter *VRFRouterTransactor) DeregisterCoordinator(opts *bind.TransactOpts, coordinatorAddress common.Address) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "deregisterCoordinator", coordinatorAddress)
}

func (_VRFRouter *VRFRouterSession) DeregisterCoordinator(coordinatorAddress common.Address) (*types.Transaction, error) {
	return _VRFRouter.Contract.DeregisterCoordinator(&_VRFRouter.TransactOpts, coordinatorAddress)
}

func (_VRFRouter *VRFRouterTransactorSession) DeregisterCoordinator(coordinatorAddress common.Address) (*types.Transaction, error) {
	return _VRFRouter.Contract.DeregisterCoordinator(&_VRFRouter.TransactOpts, coordinatorAddress)
}

func (_VRFRouter *VRFRouterTransactor) RedeemRandomness(opts *bind.TransactOpts, subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "redeemRandomness", subID, requestID, extraArgs)
}

func (_VRFRouter *VRFRouterSession) RedeemRandomness(subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.RedeemRandomness(&_VRFRouter.TransactOpts, subID, requestID, extraArgs)
}

func (_VRFRouter *VRFRouterTransactorSession) RedeemRandomness(subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.RedeemRandomness(&_VRFRouter.TransactOpts, subID, requestID, extraArgs)
}

func (_VRFRouter *VRFRouterTransactor) RegisterCoordinator(opts *bind.TransactOpts, coordinatorAddress common.Address) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "registerCoordinator", coordinatorAddress)
}

func (_VRFRouter *VRFRouterSession) RegisterCoordinator(coordinatorAddress common.Address) (*types.Transaction, error) {
	return _VRFRouter.Contract.RegisterCoordinator(&_VRFRouter.TransactOpts, coordinatorAddress)
}

func (_VRFRouter *VRFRouterTransactorSession) RegisterCoordinator(coordinatorAddress common.Address) (*types.Transaction, error) {
	return _VRFRouter.Contract.RegisterCoordinator(&_VRFRouter.TransactOpts, coordinatorAddress)
}

func (_VRFRouter *VRFRouterTransactor) RequestRandomness(opts *bind.TransactOpts, subID *big.Int, numWords uint16, confDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "requestRandomness", subID, numWords, confDelay, extraArgs)
}

func (_VRFRouter *VRFRouterSession) RequestRandomness(subID *big.Int, numWords uint16, confDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.RequestRandomness(&_VRFRouter.TransactOpts, subID, numWords, confDelay, extraArgs)
}

func (_VRFRouter *VRFRouterTransactorSession) RequestRandomness(subID *big.Int, numWords uint16, confDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.RequestRandomness(&_VRFRouter.TransactOpts, subID, numWords, confDelay, extraArgs)
}

func (_VRFRouter *VRFRouterTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID *big.Int, numWords uint16, confDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confDelay, callbackGasLimit, arguments, extraArgs)
}

func (_VRFRouter *VRFRouterSession) RequestRandomnessFulfillment(subID *big.Int, numWords uint16, confDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.RequestRandomnessFulfillment(&_VRFRouter.TransactOpts, subID, numWords, confDelay, callbackGasLimit, arguments, extraArgs)
}

func (_VRFRouter *VRFRouterTransactorSession) RequestRandomnessFulfillment(subID *big.Int, numWords uint16, confDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _VRFRouter.Contract.RequestRandomnessFulfillment(&_VRFRouter.TransactOpts, subID, numWords, confDelay, callbackGasLimit, arguments, extraArgs)
}

func (_VRFRouter *VRFRouterTransactor) ResetRoute(opts *bind.TransactOpts, subID *big.Int) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "resetRoute", subID)
}

func (_VRFRouter *VRFRouterSession) ResetRoute(subID *big.Int) (*types.Transaction, error) {
	return _VRFRouter.Contract.ResetRoute(&_VRFRouter.TransactOpts, subID)
}

func (_VRFRouter *VRFRouterTransactorSession) ResetRoute(subID *big.Int) (*types.Transaction, error) {
	return _VRFRouter.Contract.ResetRoute(&_VRFRouter.TransactOpts, subID)
}

func (_VRFRouter *VRFRouterTransactor) SetRoute(opts *bind.TransactOpts, subID *big.Int) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "setRoute", subID)
}

func (_VRFRouter *VRFRouterSession) SetRoute(subID *big.Int) (*types.Transaction, error) {
	return _VRFRouter.Contract.SetRoute(&_VRFRouter.TransactOpts, subID)
}

func (_VRFRouter *VRFRouterTransactorSession) SetRoute(subID *big.Int) (*types.Transaction, error) {
	return _VRFRouter.Contract.SetRoute(&_VRFRouter.TransactOpts, subID)
}

func (_VRFRouter *VRFRouterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFRouter.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFRouter *VRFRouterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFRouter.Contract.TransferOwnership(&_VRFRouter.TransactOpts, to)
}

func (_VRFRouter *VRFRouterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFRouter.Contract.TransferOwnership(&_VRFRouter.TransactOpts, to)
}

type VRFRouterConfigSetIterator struct {
	Event *VRFRouterConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterConfigSet)
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

func (it *VRFRouterConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFRouterConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFRouterConfigSetIterator, error) {

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFRouterConfigSetIterator{contract: _VRFRouter.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFRouterConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterConfigSet)
				if err := _VRFRouter.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseConfigSet(log types.Log) (*VRFRouterConfigSet, error) {
	event := new(VRFRouterConfigSet)
	if err := _VRFRouter.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterCoordinatorDeregisteredIterator struct {
	Event *VRFRouterCoordinatorDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterCoordinatorDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterCoordinatorDeregistered)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterCoordinatorDeregistered)
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

func (it *VRFRouterCoordinatorDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFRouterCoordinatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterCoordinatorDeregistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterCoordinatorDeregistered(opts *bind.FilterOpts) (*VRFRouterCoordinatorDeregisteredIterator, error) {

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return &VRFRouterCoordinatorDeregisteredIterator{contract: _VRFRouter.contract, event: "CoordinatorDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchCoordinatorDeregistered(opts *bind.WatchOpts, sink chan<- *VRFRouterCoordinatorDeregistered) (event.Subscription, error) {

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterCoordinatorDeregistered)
				if err := _VRFRouter.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseCoordinatorDeregistered(log types.Log) (*VRFRouterCoordinatorDeregistered, error) {
	event := new(VRFRouterCoordinatorDeregistered)
	if err := _VRFRouter.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterCoordinatorRegisteredIterator struct {
	Event *VRFRouterCoordinatorRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterCoordinatorRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterCoordinatorRegistered)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterCoordinatorRegistered)
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

func (it *VRFRouterCoordinatorRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFRouterCoordinatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterCoordinatorRegistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFRouterCoordinatorRegisteredIterator, error) {

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return &VRFRouterCoordinatorRegisteredIterator{contract: _VRFRouter.contract, event: "CoordinatorRegistered", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFRouterCoordinatorRegistered) (event.Subscription, error) {

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterCoordinatorRegistered)
				if err := _VRFRouter.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseCoordinatorRegistered(log types.Log) (*VRFRouterCoordinatorRegistered, error) {
	event := new(VRFRouterCoordinatorRegistered)
	if err := _VRFRouter.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterNewTransmissionIterator struct {
	Event *VRFRouterNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterNewTransmission)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterNewTransmission)
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

func (it *VRFRouterNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFRouterNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterNewTransmission struct {
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*VRFRouterNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFRouterNewTransmissionIterator{contract: _VRFRouter.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFRouterNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterNewTransmission)
				if err := _VRFRouter.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseNewTransmission(log types.Log) (*VRFRouterNewTransmission, error) {
	event := new(VRFRouterNewTransmission)
	if err := _VRFRouter.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterOutputsServedIterator struct {
	Event *VRFRouterOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterOutputsServed)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterOutputsServed)
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

func (it *VRFRouterOutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFRouterOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterOutputsServed struct {
	RecentBlockHeight  uint64
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFRouterOutputsServedIterator, error) {

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFRouterOutputsServedIterator{contract: _VRFRouter.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFRouterOutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterOutputsServed)
				if err := _VRFRouter.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseOutputsServed(log types.Log) (*VRFRouterOutputsServed, error) {
	event := new(VRFRouterOutputsServed)
	if err := _VRFRouter.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterOwnershipTransferRequestedIterator struct {
	Event *VRFRouterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterOwnershipTransferRequested)
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

func (it *VRFRouterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFRouterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFRouterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFRouterOwnershipTransferRequestedIterator{contract: _VRFRouter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFRouterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterOwnershipTransferRequested)
				if err := _VRFRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFRouterOwnershipTransferRequested, error) {
	event := new(VRFRouterOwnershipTransferRequested)
	if err := _VRFRouter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterOwnershipTransferredIterator struct {
	Event *VRFRouterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterOwnershipTransferred)
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

func (it *VRFRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFRouterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFRouterOwnershipTransferredIterator{contract: _VRFRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFRouterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterOwnershipTransferred)
				if err := _VRFRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseOwnershipTransferred(log types.Log) (*VRFRouterOwnershipTransferred, error) {
	event := new(VRFRouterOwnershipTransferred)
	if err := _VRFRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterRandomWordsFulfilledIterator struct {
	Event *VRFRouterRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterRandomWordsFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterRandomWordsFulfilled)
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

func (it *VRFRouterRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFRouterRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFRouterRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFRouterRandomWordsFulfilledIterator{contract: _VRFRouter.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFRouterRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterRandomWordsFulfilled)
				if err := _VRFRouter.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFRouterRandomWordsFulfilled, error) {
	event := new(VRFRouterRandomWordsFulfilled)
	if err := _VRFRouter.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterRandomnessFulfillmentRequestedIterator struct {
	Event *VRFRouterRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterRandomnessFulfillmentRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterRandomnessFulfillmentRequested)
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

func (it *VRFRouterRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFRouterRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterRandomnessFulfillmentRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	GasAllowance           uint32
	GasPrice               *big.Int
	WeiPerUnitLink         *big.Int
	Arguments              []byte
	Raw                    types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFRouterRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFRouterRandomnessFulfillmentRequestedIterator{contract: _VRFRouter.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFRouterRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterRandomnessFulfillmentRequested)
				if err := _VRFRouter.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFRouterRandomnessFulfillmentRequested, error) {
	event := new(VRFRouterRandomnessFulfillmentRequested)
	if err := _VRFRouter.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterRandomnessRequestedIterator struct {
	Event *VRFRouterRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterRandomnessRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterRandomnessRequested)
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

func (it *VRFRouterRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFRouterRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	Raw                    types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFRouterRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFRouterRandomnessRequestedIterator{contract: _VRFRouter.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFRouterRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterRandomnessRequested)
				if err := _VRFRouter.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseRandomnessRequested(log types.Log) (*VRFRouterRandomnessRequested, error) {
	event := new(VRFRouterRandomnessRequested)
	if err := _VRFRouter.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRouterRouteSetIterator struct {
	Event *VRFRouterRouteSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRouterRouteSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRouterRouteSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFRouterRouteSet)
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

func (it *VRFRouterRouteSetIterator) Error() error {
	return it.fail
}

func (it *VRFRouterRouteSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRouterRouteSet struct {
	SubID              *big.Int
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFRouter *VRFRouterFilterer) FilterRouteSet(opts *bind.FilterOpts, subID []*big.Int) (*VRFRouterRouteSetIterator, error) {

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFRouter.contract.FilterLogs(opts, "RouteSet", subIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFRouterRouteSetIterator{contract: _VRFRouter.contract, event: "RouteSet", logs: logs, sub: sub}, nil
}

func (_VRFRouter *VRFRouterFilterer) WatchRouteSet(opts *bind.WatchOpts, sink chan<- *VRFRouterRouteSet, subID []*big.Int) (event.Subscription, error) {

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFRouter.contract.WatchLogs(opts, "RouteSet", subIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRouterRouteSet)
				if err := _VRFRouter.contract.UnpackLog(event, "RouteSet", log); err != nil {
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

func (_VRFRouter *VRFRouterFilterer) ParseRouteSet(log types.Log) (*VRFRouterRouteSet, error) {
	event := new(VRFRouterRouteSet)
	if err := _VRFRouter.contract.UnpackLog(event, "RouteSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

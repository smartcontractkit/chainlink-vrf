package vrf

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

type ECCArithmeticG1Point struct {
	P [2]*big.Int
}

type ECCArithmeticG2Point struct {
	P [4]*big.Int
}

type HashToCurveFProof struct {
	DenomInv    *big.Int
	TInvSquared *big.Int
	Y1          *big.Int
	Y2          *big.Int
	Y3          *big.Int
}

type KeyDataStructKeyData struct {
	PublicKey []byte
	Hashes    [][32]byte
}

type VRFProof struct {
	PubKey ECCArithmeticG2Point
	Output ECCArithmeticG1Point
	F1     HashToCurveFProof
	F2     HashToCurveFProof
}

type VRFRequest struct {
	RequestID [32]byte
	Seed      *big.Int
	NumWords  uint32
	Sender    common.Address
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

var ChainSpecificUtilMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x602d6037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea164736f6c6343000813000a",
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
	Bin: "0x608060405234801561001057600080fd5b5060405161045538038061045583398101604081905261002f9161016e565b8060006001600160a01b03821661008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100bd576100bd816100c5565b50505061019e565b336001600160a01b0382160361011d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561018057600080fd5b81516001600160a01b038116811461019757600080fd5b9392505050565b6102a8806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c6343000813000a",
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
	Bin: "0x608060405234801561001057600080fd5b5060405161047038038061047083398101604081905261002f91610186565b6001600160a01b03821661008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100ba576100ba816100c1565b50506101b9565b336001600160a01b038216036101195760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610081565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b038116811461018157600080fd5b919050565b6000806040838503121561019957600080fd5b6101a28361016a565b91506101b06020840161016a565b90509250929050565b6102a8806101c86000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c6343000813000a",
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

var DKGMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"CalldataLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfFaultyOracles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedNumSignatures\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ssLength\",\"type\":\"uint256\"}],\"name\":\"IncorrectNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expectedLength\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"InvalidOnchainConfigLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"InvalidTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyIDCopyFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numFaultyOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"}],\"name\":\"NumberOfFaultyOraclesTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedSignerAddress\",\"type\":\"address\"}],\"name\":\"RepeatedSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedTransmitterAddress\",\"type\":\"address\"}],\"name\":\"RepeatedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numTransmitters\",\"type\":\"uint256\"}],\"name\":\"SignersTransmittersMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maxOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"providedOracles\",\"type\":\"uint256\"}],\"name\":\"TooManyOracles\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractDKGClient\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"DKGClientError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"indexed\":false,\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"KeyGenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"addClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_configDigest\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"removeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000a5565b50505062000150565b336001600160a01b03821603620000ff5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61256d80620001606000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063c3105a6b1161005b578063c3105a6b146101ce578063e3d0e712146101ee578063f2fde38b1461020157600080fd5b80638da5cb5b14610176578063afcb95d714610191578063b1dc65a4146101bb57600080fd5b806379ba5097116100b257806379ba50971461012b5780637bf1ffc51461013357806381ff70481461014657600080fd5b8063181f5a77146100ce5780635429a79e14610116575b600080fd5b604080518082018252600981527f444b4720302e302e3100000000000000000000000000000000000000000000006020820152905161010d9190611b84565b60405180910390f35b610129610124366004611bb3565b610214565b005b610129610455565b610129610141366004611bb3565b61050b565b6007546005546040805163ffffffff8085168252640100000000909404909316602084015282015260600161010d565b6000546040516001600160a01b03909116815260200161010d565b6005546004546040805160008152602081019390935263ffffffff9091169082015260600161010d565b6101296101c9366004611c2f565b610552565b6101e16101dc366004611d14565b610695565b60405161010d9190611d36565b6101296101fc366004611f19565b6107be565b61012961020f366004611fe6565b610f1d565b61021c610f31565b60008281526002602090815260408083208054825181850281018501909352808352919290919083018282801561027c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161025e575b505050505090506000815167ffffffffffffffff81111561029f5761029f611d9e565b6040519080825280602002602001820160405280156102c8578160200160208202803683370190505b5090506000805b835181101561036b57846001600160a01b03168482815181106102f4576102f4612003565b60200260200101516001600160a01b03161461034b578483610316848461202f565b8151811061032657610326612003565b60200260200101906001600160a01b031690816001600160a01b031681525050610359565b8161035581612042565b9250505b8061036381612042565b9150506102cf565b50600081845161037b919061202f565b67ffffffffffffffff81111561039357610393611d9e565b6040519080825280602002602001820160405280156103bc578160200160208202803683370190505b50905060005b8285516103cf919061202f565b81101561042c578381815181106103e8576103e8612003565b602002602001015182828151811061040257610402612003565b6001600160a01b03909216602092830291909101909101528061042481612042565b9150506103c2565b506000868152600260209081526040909120825161044c92840190611a60565b50505050505050565b6001546001600160a01b031633146104b45760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610513610f31565b600091825260026020908152604083208054600181018255908452922090910180546001600160a01b0319166001600160a01b03909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916105a29184918491908e908e9081908401838280828437600092019190915250610f8d92505050565b6040805183815263ffffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260055480825260065460ff808216602085015261010090910416928201929092529083146106555780516040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810191909152602481018490526044016104ab565b6106638b8b8b8b8b8b6111e7565b6106748c8c8c8c8c8c8c8c8961126e565b50505063ffffffff811061068a5761068a61205b565b505050505050505050565b604080518082019091526060808252602082015260008381526003602090815260408083208584529091529081902081518083019092528054829082906106db90612071565b80601f016020809104026020016040519081016040528092919081815260200182805461070790612071565b80156107545780601f1061072957610100808354040283529160200191610754565b820191906000526020600020905b81548152906001019060200180831161073757829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156107ac57602002820191906000526020600020905b815481526020019060010190808311610798575b50505050508152505090505b92915050565b8551855185601f831115610808576040517f809fc428000000000000000000000000000000000000000000000000000000008152601f6004820152602481018490526044016104ab565b81831461084b576040517f988a080400000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044016104ab565b6108568160036120ab565b60ff16831161089d576040517ffda9db7800000000000000000000000000000000000000000000000000000000815260ff82166004820152602481018490526044016104ab565b8060ff166000036108da576040517fe77dba5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108e2610f31565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60095415610a325760095460009061093a9060019061202f565b905060006009828154811061095157610951612003565b6000918252602082200154600a80546001600160a01b039092169350908490811061097e5761097e612003565b60009182526020808320909101546001600160a01b03858116845260089092526040808420805461ffff19908116909155929091168084529220805490911690556009805491925090806109d4576109d46120ce565b600082815260209020810160001990810180546001600160a01b0319169055019055600a805480610a0757610a076120ce565b600082815260209020810160001990810180546001600160a01b031916905501905550610920915050565b60005b815151811015610d9f5760006008600084600001518481518110610a5b57610a5b612003565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610a9857610a986120e4565b14610afb578151805182908110610ab157610ab1612003565b60200260200101516040517f7451f83e0000000000000000000000000000000000000000000000000000000081526004016104ab91906001600160a01b0391909116815260200190565b6040805180820190915260ff82168152600160208201528251805160089160009185908110610b2c57610b2c612003565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610b8557610b856120e4565b021790555060009150610b959050565b6008600084602001518481518110610baf57610baf612003565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610bec57610bec6120e4565b14610c515781602001518181518110610c0757610c07612003565b60200260200101516040517fe8d298990000000000000000000000000000000000000000000000000000000081526004016104ab91906001600160a01b0391909116815260200190565b6040805180820190915260ff821681526020810160028152506008600084602001518481518110610c8457610c84612003565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610cdd57610cdd6120e4565b021790555050825180516009925083908110610cfb57610cfb612003565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b03909316929092179091558201518051600a919083908110610d5257610d52612003565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b0390921691909117905580610d9781612042565b915050610a35565b5060408101516006805460ff191660ff909216919091179055600754640100000000900463ffffffff16610dd16116bf565b6007805463ffffffff9283166401000000000267ffffffff0000000019821681178355600193600092610e0b9286929081169116176120fa565b92506101000a81548163ffffffff021916908363ffffffff1602179055506000610e6c4630600760009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a00151611749565b6005819055835180516006805460ff9092166101000261ff00199092169190911790556007546020860151604080880151606089015160808a015160a08b015193519798507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0597610ef4978b978b9763ffffffff90911696919590949093909290919061215b565b60405180910390a1610f0f83604001518460600151836117d6565b505050505050505050505050565b610f25610f31565b610f2e816119b7565b50565b6000546001600160a01b03163314610f8b5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016104ab565b565b600060608083806020019051810190610fa691906121f1565b60408051808201825283815260208082018490526000868152600282528381208054855181850281018501909652808652979a509598509396509094929391929083018282801561102057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611002575b5050505050905060005b815181101561113d5781818151811061104557611045612003565b60200260200101516001600160a01b031663bf2732c7846040518263ffffffff1660e01b81526004016110789190611d36565b600060405180830381600087803b15801561109257600080fd5b505af19250505080156110a3575060015b61112b573d8080156110d1576040519150601f19603f3d011682016040523d82523d6000602084013e6110d6565b606091505b507f116391732f5df106193bda7cedf1728f3b07b62f6cdcdd611c9eeec44efcae5483838151811061110a5761110a612003565b6020026020010151826040516111219291906122ef565b60405180910390a1505b8061113581612042565b91505061102a565b5060008581526003602090815260408083208b84529091529020825183919081906111689082612368565b5060208281015180516111819260018501920190611ac5565b5090505084887fc8db841f5b2231ccf7190311f440aa197b161e369f3b40b023508160cc555656846040516111b69190611d36565b60405180910390a350506004805460089690961c63ffffffff1663ffffffff19909616959095179094555050505050565b60006111f4826020612428565b6111ff856020612428565b61120b8861014461243f565b611215919061243f565b61121f919061243f565b61122a90600061243f565b905036811461044c576040517ff7b94f0a000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016104ab565b60006002826020015183604001516112869190612452565b611290919061246b565b61129b906001612452565b60408051600180825281830190925260ff929092169250600091906020820181803683370190505090508160f81b816000815181106112dc576112dc612003565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350868214158061131a5750868514155b15611362576040517fe307bd570000000000000000000000000000000000000000000000000000000081526004810183905260248101889052604481018690526064016104ab565b3360009081526008602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156113a5576113a56120e4565b60028111156113b6576113b66120e4565b90525090506002816020015160028111156113d3576113d36120e4565b14158061140e5750600a816000015160ff16815481106113f5576113f5612003565b6000918252602090912001546001600160a01b03163314155b15611447576040517f2d0f0c0f0000000000000000000000000000000000000000000000000000000081523360048201526024016104ab565b5050506000888860405161145c92919061249b565b604051908190038120611473918c906020016124ab565b604051602081830303815290604052805190602001209050611493611b00565b604080518082019091526000808252602082015260005b888110156116b05760006001858884602081106114c9576114c9612003565b6114d691901a601b612452565b8d8d868181106114e8576114e8612003565b905060200201358c8c8781811061150157611501612003565b905060200201356040516000815260200160405260405161153e949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611560573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526008602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156115b5576115b56120e4565b60028111156115c6576115c66120e4565b90525092506001836020015160028111156115e3576115e36120e4565b14611625576040517fbf18af430000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016104ab565b8251849060ff16601f811061163c5761163c612003565b602002015115611678576040517f21cf3b4400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600184846000015160ff16601f811061169357611693612003565b9115156020909202015250806116a881612042565b9150506114aa565b50505050505050505050505050565b60004661a4b18114806116d4575062066eed81145b156117425760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611718573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061173c91906124bf565b91505090565b4391505090565b6000808a8a8a8a8a8a8a8a8a60405160200161176d999897969594939291906124d8565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60008083516020146118215783516040517f1625adfe0000000000000000000000000000000000000000000000000000000081526020600482015260248101919091526044016104ab565b60208401519150808203611861576040517faf5e77d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260608082526020820152600083815260036020908152604080832087845290915290208151829190819061189f9082612368565b5060208281015180516118b89260018501920190611ac5565b50505060008381526002602090815260408083208054825181850281018501909352808352919290919083018282801561191b57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116118fd575b5050505050905060005b81518110156119ad5781818151811061194057611940612003565b60200260200101516001600160a01b03166355e487496040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561198257600080fd5b505af1158015611996573d6000803e3d6000fd5b5050505080806119a590612042565b915050611925565b5050505050505050565b336001600160a01b03821603611a0f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016104ab565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215611ab5579160200282015b82811115611ab557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611a80565b50611ac1929150611b1f565b5090565b828054828255906000526020600020908101928215611ab5579160200282015b82811115611ab5578251825591602001919060010190611ae5565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115611ac15760008155600101611b20565b60005b83811015611b4f578181015183820152602001611b37565b50506000910152565b60008151808452611b70816020860160208601611b34565b601f01601f19169290920160200192915050565b602081526000611b976020830184611b58565b9392505050565b6001600160a01b0381168114610f2e57600080fd5b60008060408385031215611bc657600080fd5b823591506020830135611bd881611b9e565b809150509250929050565b60008083601f840112611bf557600080fd5b50813567ffffffffffffffff811115611c0d57600080fd5b6020830191508360208260051b8501011115611c2857600080fd5b9250929050565b60008060008060008060008060e0898b031215611c4b57600080fd5b606089018a811115611c5c57600080fd5b8998503567ffffffffffffffff80821115611c7657600080fd5b818b0191508b601f830112611c8a57600080fd5b813581811115611c9957600080fd5b8c6020828501011115611cab57600080fd5b6020830199508098505060808b0135915080821115611cc957600080fd5b611cd58c838d01611be3565b909750955060a08b0135915080821115611cee57600080fd5b50611cfb8b828c01611be3565b999c989b50969995989497949560c00135949350505050565b60008060408385031215611d2757600080fd5b50508035926020909101359150565b600060208083528351604082850152611d526060850182611b58565b85830151858203601f19016040870152805180835290840192506000918401905b80831015611d935783518252928401926001929092019190840190611d73565b509695505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611ddd57611ddd611d9e565b604052919050565b600067ffffffffffffffff821115611dff57611dff611d9e565b5060051b60200190565b600082601f830112611e1a57600080fd5b81356020611e2f611e2a83611de5565b611db4565b82815260059290921b84018101918181019086841115611e4e57600080fd5b8286015b84811015611d93578035611e6581611b9e565b8352918301918301611e52565b803560ff81168114611e8357600080fd5b919050565b600067ffffffffffffffff821115611ea257611ea2611d9e565b50601f01601f191660200190565b600082601f830112611ec157600080fd5b8135611ecf611e2a82611e88565b818152846020838601011115611ee457600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff81168114611e8357600080fd5b60008060008060008060c08789031215611f3257600080fd5b863567ffffffffffffffff80821115611f4a57600080fd5b611f568a838b01611e09565b97506020890135915080821115611f6c57600080fd5b611f788a838b01611e09565b9650611f8660408a01611e72565b95506060890135915080821115611f9c57600080fd5b611fa88a838b01611eb0565b9450611fb660808a01611f01565b935060a0890135915080821115611fcc57600080fd5b50611fd989828a01611eb0565b9150509295509295509295565b600060208284031215611ff857600080fd5b8135611b9781611b9e565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b818103818111156107b8576107b8612019565b60006001820161205457612054612019565b5060010190565b634e487b7160e01b600052600160045260246000fd5b600181811c9082168061208557607f821691505b6020821081036120a557634e487b7160e01b600052602260045260246000fd5b50919050565b60ff81811683821602908116908181146120c7576120c7612019565b5092915050565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b63ffffffff8181168382160190808211156120c7576120c7612019565b600081518084526020808501945080840160005b838110156121505781516001600160a01b03168752958201959082019060010161212b565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261218b8184018a612117565b9050828103608084015261219f8189612117565b905060ff871660a084015282810360c08401526121bc8187611b58565b905067ffffffffffffffff851660e08401528281036101008401526121e18185611b58565b9c9b505050505050505050505050565b60008060006060848603121561220657600080fd5b8351925060208085015167ffffffffffffffff8082111561222657600080fd5b818701915087601f83011261223a57600080fd5b8151612248611e2a82611e88565b818152898583860101111561225c57600080fd5b61226b82868301878701611b34565b60408901519096509250508082111561228357600080fd5b508501601f8101871361229557600080fd5b80516122a3611e2a82611de5565b81815260059190911b820183019083810190898311156122c257600080fd5b928401925b828410156122e0578351825292840192908401906122c7565b80955050505050509250925092565b6001600160a01b03831681526040602082015260006123116040830184611b58565b949350505050565b601f82111561236357600081815260208120601f850160051c810160208610156123405750805b601f850160051c820191505b8181101561235f5782815560010161234c565b5050505b505050565b815167ffffffffffffffff81111561238257612382611d9e565b612396816123908454612071565b84612319565b602080601f8311600181146123cb57600084156123b35750858301515b600019600386901b1c1916600185901b17855561235f565b600085815260208120601f198616915b828110156123fa578886015182559484019460019091019084016123db565b50858210156124185787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b80820281158282048414176107b8576107b8612019565b808201808211156107b8576107b8612019565b60ff81811683821601908111156107b8576107b8612019565b600060ff83168061248c57634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b8183823760009101908152919050565b828152606082602083013760800192915050565b6000602082840312156124d157600080fd5b5051919050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526125128285018b612117565b91508382036080850152612526828a612117565b915060ff881660a085015283820360c08501526125438288611b58565b90861660e085015283810361010085015290506121e18185611b5856fea164736f6c6343000813000a",
}

var DKGABI = DKGMetaData.ABI

var DKGBin = DKGMetaData.Bin

func DeployDKG(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DKG, error) {
	parsed, err := DKGMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DKG{DKGCaller: DKGCaller{contract: contract}, DKGTransactor: DKGTransactor{contract: contract}, DKGFilterer: DKGFilterer{contract: contract}}, nil
}

type DKG struct {
	DKGCaller
	DKGTransactor
	DKGFilterer
}

type DKGCaller struct {
	contract *bind.BoundContract
}

type DKGTransactor struct {
	contract *bind.BoundContract
}

type DKGFilterer struct {
	contract *bind.BoundContract
}

type DKGSession struct {
	Contract     *DKG
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type DKGCallerSession struct {
	Contract *DKGCaller
	CallOpts bind.CallOpts
}

type DKGTransactorSession struct {
	Contract     *DKGTransactor
	TransactOpts bind.TransactOpts
}

type DKGRaw struct {
	Contract *DKG
}

type DKGCallerRaw struct {
	Contract *DKGCaller
}

type DKGTransactorRaw struct {
	Contract *DKGTransactor
}

func NewDKG(address common.Address, backend bind.ContractBackend) (*DKG, error) {
	contract, err := bindDKG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKG{DKGCaller: DKGCaller{contract: contract}, DKGTransactor: DKGTransactor{contract: contract}, DKGFilterer: DKGFilterer{contract: contract}}, nil
}

func NewDKGCaller(address common.Address, caller bind.ContractCaller) (*DKGCaller, error) {
	contract, err := bindDKG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGCaller{contract: contract}, nil
}

func NewDKGTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGTransactor, error) {
	contract, err := bindDKG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGTransactor{contract: contract}, nil
}

func NewDKGFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGFilterer, error) {
	contract, err := bindDKG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGFilterer{contract: contract}, nil
}

func bindDKG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DKGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_DKG *DKGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.DKGCaller.contract.Call(opts, result, method, params...)
}

func (_DKG *DKGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transfer(opts)
}

func (_DKG *DKGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transact(opts, method, params...)
}

func (_DKG *DKGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.contract.Call(opts, result, method, params...)
}

func (_DKG *DKGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transfer(opts)
}

func (_DKG *DKGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transact(opts, method, params...)
}

func (_DKG *DKGCaller) GetKey(opts *bind.CallOpts, _keyID [32]byte, _configDigest [32]byte) (KeyDataStructKeyData, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getKey", _keyID, _configDigest)

	if err != nil {
		return *new(KeyDataStructKeyData), err
	}

	out0 := *abi.ConvertType(out[0], new(KeyDataStructKeyData)).(*KeyDataStructKeyData)

	return out0, err

}

func (_DKG *DKGSession) GetKey(_keyID [32]byte, _configDigest [32]byte) (KeyDataStructKeyData, error) {
	return _DKG.Contract.GetKey(&_DKG.CallOpts, _keyID, _configDigest)
}

func (_DKG *DKGCallerSession) GetKey(_keyID [32]byte, _configDigest [32]byte) (KeyDataStructKeyData, error) {
	return _DKG.Contract.GetKey(&_DKG.CallOpts, _keyID, _configDigest)
}

func (_DKG *DKGCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(struct {
		ConfigCount  uint32
		BlockNumber  uint32
		ConfigDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_DKG *DKGSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _DKG.Contract.LatestConfigDetails(&_DKG.CallOpts)
}

func (_DKG *DKGCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _DKG.Contract.LatestConfigDetails(&_DKG.CallOpts)
}

func (_DKG *DKGCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(struct {
		ScanLogs     bool
		ConfigDigest [32]byte
		Epoch        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_DKG *DKGSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _DKG.Contract.LatestConfigDigestAndEpoch(&_DKG.CallOpts)
}

func (_DKG *DKGCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _DKG.Contract.LatestConfigDigestAndEpoch(&_DKG.CallOpts)
}

func (_DKG *DKGCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_DKG *DKGSession) Owner() (common.Address, error) {
	return _DKG.Contract.Owner(&_DKG.CallOpts)
}

func (_DKG *DKGCallerSession) Owner() (common.Address, error) {
	return _DKG.Contract.Owner(&_DKG.CallOpts)
}

func (_DKG *DKGCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_DKG *DKGSession) TypeAndVersion() (string, error) {
	return _DKG.Contract.TypeAndVersion(&_DKG.CallOpts)
}

func (_DKG *DKGCallerSession) TypeAndVersion() (string, error) {
	return _DKG.Contract.TypeAndVersion(&_DKG.CallOpts)
}

func (_DKG *DKGTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "acceptOwnership")
}

func (_DKG *DKGSession) AcceptOwnership() (*types.Transaction, error) {
	return _DKG.Contract.AcceptOwnership(&_DKG.TransactOpts)
}

func (_DKG *DKGTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DKG.Contract.AcceptOwnership(&_DKG.TransactOpts)
}

func (_DKG *DKGTransactor) AddClient(opts *bind.TransactOpts, keyID [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "addClient", keyID, clientAddress)
}

func (_DKG *DKGSession) AddClient(keyID [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AddClient(&_DKG.TransactOpts, keyID, clientAddress)
}

func (_DKG *DKGTransactorSession) AddClient(keyID [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AddClient(&_DKG.TransactOpts, keyID, clientAddress)
}

func (_DKG *DKGTransactor) RemoveClient(opts *bind.TransactOpts, keyID [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "removeClient", keyID, clientAddress)
}

func (_DKG *DKGSession) RemoveClient(keyID [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _DKG.Contract.RemoveClient(&_DKG.TransactOpts, keyID, clientAddress)
}

func (_DKG *DKGTransactorSession) RemoveClient(keyID [32]byte, clientAddress common.Address) (*types.Transaction, error) {
	return _DKG.Contract.RemoveClient(&_DKG.TransactOpts, keyID, clientAddress)
}

func (_DKG *DKGTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_DKG *DKGSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _DKG.Contract.SetConfig(&_DKG.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_DKG *DKGTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _DKG.Contract.SetConfig(&_DKG.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_DKG *DKGTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "transferOwnership", to)
}

func (_DKG *DKGSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DKG.Contract.TransferOwnership(&_DKG.TransactOpts, to)
}

func (_DKG *DKGTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DKG.Contract.TransferOwnership(&_DKG.TransactOpts, to)
}

func (_DKG *DKGTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_DKG *DKGSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _DKG.Contract.Transmit(&_DKG.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_DKG *DKGTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _DKG.Contract.Transmit(&_DKG.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type DKGConfigSetIterator struct {
	Event *DKGConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DKGConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGConfigSet)
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
		it.Event = new(DKGConfigSet)
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

func (it *DKGConfigSetIterator) Error() error {
	return it.fail
}

func (it *DKGConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DKGConfigSet struct {
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

func (_DKG *DKGFilterer) FilterConfigSet(opts *bind.FilterOpts) (*DKGConfigSetIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &DKGConfigSetIterator{contract: _DKG.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_DKG *DKGFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *DKGConfigSet) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DKGConfigSet)
				if err := _DKG.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_DKG *DKGFilterer) ParseConfigSet(log types.Log) (*DKGConfigSet, error) {
	event := new(DKGConfigSet)
	if err := _DKG.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DKGDKGClientErrorIterator struct {
	Event *DKGDKGClientError

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DKGDKGClientErrorIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGDKGClientError)
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
		it.Event = new(DKGDKGClientError)
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

func (it *DKGDKGClientErrorIterator) Error() error {
	return it.fail
}

func (it *DKGDKGClientErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DKGDKGClientError struct {
	Client    common.Address
	ErrorData []byte
	Raw       types.Log
}

func (_DKG *DKGFilterer) FilterDKGClientError(opts *bind.FilterOpts) (*DKGDKGClientErrorIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "DKGClientError")
	if err != nil {
		return nil, err
	}
	return &DKGDKGClientErrorIterator{contract: _DKG.contract, event: "DKGClientError", logs: logs, sub: sub}, nil
}

func (_DKG *DKGFilterer) WatchDKGClientError(opts *bind.WatchOpts, sink chan<- *DKGDKGClientError) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "DKGClientError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DKGDKGClientError)
				if err := _DKG.contract.UnpackLog(event, "DKGClientError", log); err != nil {
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

func (_DKG *DKGFilterer) ParseDKGClientError(log types.Log) (*DKGDKGClientError, error) {
	event := new(DKGDKGClientError)
	if err := _DKG.contract.UnpackLog(event, "DKGClientError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DKGKeyGeneratedIterator struct {
	Event *DKGKeyGenerated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DKGKeyGeneratedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGKeyGenerated)
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
		it.Event = new(DKGKeyGenerated)
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

func (it *DKGKeyGeneratedIterator) Error() error {
	return it.fail
}

func (it *DKGKeyGeneratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DKGKeyGenerated struct {
	ConfigDigest [32]byte
	KeyID        [32]byte
	Key          KeyDataStructKeyData
	Raw          types.Log
}

func (_DKG *DKGFilterer) FilterKeyGenerated(opts *bind.FilterOpts, configDigest [][32]byte, keyID [][32]byte) (*DKGKeyGeneratedIterator, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}
	var keyIDRule []interface{}
	for _, keyIDItem := range keyID {
		keyIDRule = append(keyIDRule, keyIDItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "KeyGenerated", configDigestRule, keyIDRule)
	if err != nil {
		return nil, err
	}
	return &DKGKeyGeneratedIterator{contract: _DKG.contract, event: "KeyGenerated", logs: logs, sub: sub}, nil
}

func (_DKG *DKGFilterer) WatchKeyGenerated(opts *bind.WatchOpts, sink chan<- *DKGKeyGenerated, configDigest [][32]byte, keyID [][32]byte) (event.Subscription, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}
	var keyIDRule []interface{}
	for _, keyIDItem := range keyID {
		keyIDRule = append(keyIDRule, keyIDItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "KeyGenerated", configDigestRule, keyIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DKGKeyGenerated)
				if err := _DKG.contract.UnpackLog(event, "KeyGenerated", log); err != nil {
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

func (_DKG *DKGFilterer) ParseKeyGenerated(log types.Log) (*DKGKeyGenerated, error) {
	event := new(DKGKeyGenerated)
	if err := _DKG.contract.UnpackLog(event, "KeyGenerated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DKGOwnershipTransferRequestedIterator struct {
	Event *DKGOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DKGOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGOwnershipTransferRequested)
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
		it.Event = new(DKGOwnershipTransferRequested)
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

func (it *DKGOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *DKGOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DKGOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DKG *DKGFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DKGOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DKGOwnershipTransferRequestedIterator{contract: _DKG.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_DKG *DKGFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DKGOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DKGOwnershipTransferRequested)
				if err := _DKG.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_DKG *DKGFilterer) ParseOwnershipTransferRequested(log types.Log) (*DKGOwnershipTransferRequested, error) {
	event := new(DKGOwnershipTransferRequested)
	if err := _DKG.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DKGOwnershipTransferredIterator struct {
	Event *DKGOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DKGOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGOwnershipTransferred)
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
		it.Event = new(DKGOwnershipTransferred)
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

func (it *DKGOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *DKGOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DKGOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_DKG *DKGFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DKGOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DKGOwnershipTransferredIterator{contract: _DKG.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_DKG *DKGFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DKGOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DKGOwnershipTransferred)
				if err := _DKG.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_DKG *DKGFilterer) ParseOwnershipTransferred(log types.Log) (*DKGOwnershipTransferred, error) {
	event := new(DKGOwnershipTransferred)
	if err := _DKG.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type DKGTransmittedIterator struct {
	Event *DKGTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *DKGTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGTransmitted)
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
		it.Event = new(DKGTransmitted)
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

func (it *DKGTransmittedIterator) Error() error {
	return it.fail
}

func (it *DKGTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type DKGTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_DKG *DKGFilterer) FilterTransmitted(opts *bind.FilterOpts) (*DKGTransmittedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &DKGTransmittedIterator{contract: _DKG.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_DKG *DKGFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *DKGTransmitted) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(DKGTransmitted)
				if err := _DKG.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_DKG *DKGFilterer) ParseTransmitted(log types.Log) (*DKGTransmitted, error) {
	event := new(DKGTransmitted)
	if err := _DKG.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var DKGClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var DKGClientABI = DKGClientMetaData.ABI

type DKGClient struct {
	DKGClientCaller
	DKGClientTransactor
	DKGClientFilterer
}

type DKGClientCaller struct {
	contract *bind.BoundContract
}

type DKGClientTransactor struct {
	contract *bind.BoundContract
}

type DKGClientFilterer struct {
	contract *bind.BoundContract
}

type DKGClientSession struct {
	Contract     *DKGClient
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type DKGClientCallerSession struct {
	Contract *DKGClientCaller
	CallOpts bind.CallOpts
}

type DKGClientTransactorSession struct {
	Contract     *DKGClientTransactor
	TransactOpts bind.TransactOpts
}

type DKGClientRaw struct {
	Contract *DKGClient
}

type DKGClientCallerRaw struct {
	Contract *DKGClientCaller
}

type DKGClientTransactorRaw struct {
	Contract *DKGClientTransactor
}

func NewDKGClient(address common.Address, backend bind.ContractBackend) (*DKGClient, error) {
	contract, err := bindDKGClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKGClient{DKGClientCaller: DKGClientCaller{contract: contract}, DKGClientTransactor: DKGClientTransactor{contract: contract}, DKGClientFilterer: DKGClientFilterer{contract: contract}}, nil
}

func NewDKGClientCaller(address common.Address, caller bind.ContractCaller) (*DKGClientCaller, error) {
	contract, err := bindDKGClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGClientCaller{contract: contract}, nil
}

func NewDKGClientTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGClientTransactor, error) {
	contract, err := bindDKGClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGClientTransactor{contract: contract}, nil
}

func NewDKGClientFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGClientFilterer, error) {
	contract, err := bindDKGClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGClientFilterer{contract: contract}, nil
}

func bindDKGClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DKGClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_DKGClient *DKGClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGClient.Contract.DKGClientCaller.contract.Call(opts, result, method, params...)
}

func (_DKGClient *DKGClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGClient.Contract.DKGClientTransactor.contract.Transfer(opts)
}

func (_DKGClient *DKGClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGClient.Contract.DKGClientTransactor.contract.Transact(opts, method, params...)
}

func (_DKGClient *DKGClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGClient.Contract.contract.Call(opts, result, method, params...)
}

func (_DKGClient *DKGClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGClient.Contract.contract.Transfer(opts)
}

func (_DKGClient *DKGClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGClient.Contract.contract.Transact(opts, method, params...)
}

func (_DKGClient *DKGClientTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _DKGClient.contract.Transact(opts, "keyGenerated", kd)
}

func (_DKGClient *DKGClientSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _DKGClient.Contract.KeyGenerated(&_DKGClient.TransactOpts, kd)
}

func (_DKGClient *DKGClientTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _DKGClient.Contract.KeyGenerated(&_DKGClient.TransactOpts, kd)
}

func (_DKGClient *DKGClientTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGClient.contract.Transact(opts, "newKeyRequested")
}

func (_DKGClient *DKGClientSession) NewKeyRequested() (*types.Transaction, error) {
	return _DKGClient.Contract.NewKeyRequested(&_DKGClient.TransactOpts)
}

func (_DKGClient *DKGClientTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _DKGClient.Contract.NewKeyRequested(&_DKGClient.TransactOpts)
}

var DebugMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addressToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061063e806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630bc643e81461005157806339614e4f1461007b5780635e57966d1461009b5780639201de55146100ae575b600080fd5b61006461005f3660046103cd565b6100c1565b60405160ff90911681526020015b60405180910390f35b61008e610089366004610406565b6100eb565b60405161007291906104b7565b61008e6100a9366004610505565b61026d565b61008e6100bc36600461053b565b610341565b6000600a8260ff1610156100e0576100da82603061056a565b92915050565b6100da82605761056a565b6060600080835160026100fe9190610583565b67ffffffffffffffff811115610116576101166103f0565b6040519080825280601f01601f191660200182016040528015610140576020820181803683370190505b509050600091505b80518260ff1610156102665760008461016260028561059a565b60ff1681518110610175576101756105ca565b60209101015160f81c600f169050600060048661019360028761059a565b60ff16815181106101a6576101a66105ca565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c90506101dd816100c1565b60f81b838560ff16815181106101f5576101f56105ca565b60200101906001600160f81b031916908160001a90535061021784600161056a565b9350610222826100c1565b60f81b838560ff168151811061023a5761023a6105ca565b60200101906001600160f81b031916908160001a9053505050818061025e906105e0565b925050610148565b9392505050565b604080516014808252818301909252606091600091906020820181803683370190505090508260005b60148160ff16101561032f577fff0000000000000000000000000000000000000000000000000000000000000060f883901b16836102d58360136105ff565b60ff16815181106102e8576102e86105ca565b60200101906001600160f81b031916908160001a90535060088273ffffffffffffffffffffffffffffffffffffffff16901c91508080610327906105e0565b915050610296565b50610339826100eb565b949350505050565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156103c357838160208110610384576103846105ca565b1a60f81b82828151811061039a5761039a6105ca565b60200101906001600160f81b031916908160001a905350806103bb81610618565b915050610369565b50610266816100eb565b6000602082840312156103df57600080fd5b813560ff8116811461026657600080fd5b634e487b7160e01b600052604160045260246000fd5b60006020828403121561041857600080fd5b813567ffffffffffffffff8082111561043057600080fd5b818401915084601f83011261044457600080fd5b813581811115610456576104566103f0565b604051601f8201601f19908116603f0116810190838211818310171561047e5761047e6103f0565b8160405282815287602084870101111561049757600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156104e4578581018301518582016040015282016104c8565b506000604082860101526040601f19601f8301168501019250505092915050565b60006020828403121561051757600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461026657600080fd5b60006020828403121561054d57600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b60ff81811683821601908111156100da576100da610554565b80820281158282048414176100da576100da610554565b600060ff8316806105bb57634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036105f6576105f6610554565b60010192915050565b60ff82811682821603908111156100da576100da610554565b60006001820161062a5761062a610554565b506001019056fea164736f6c6343000813000a",
}

var DebugABI = DebugMetaData.ABI

var DebugBin = DebugMetaData.Bin

func DeployDebug(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Debug, error) {
	parsed, err := DebugMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DebugBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Debug{DebugCaller: DebugCaller{contract: contract}, DebugTransactor: DebugTransactor{contract: contract}, DebugFilterer: DebugFilterer{contract: contract}}, nil
}

type Debug struct {
	DebugCaller
	DebugTransactor
	DebugFilterer
}

type DebugCaller struct {
	contract *bind.BoundContract
}

type DebugTransactor struct {
	contract *bind.BoundContract
}

type DebugFilterer struct {
	contract *bind.BoundContract
}

type DebugSession struct {
	Contract     *Debug
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type DebugCallerSession struct {
	Contract *DebugCaller
	CallOpts bind.CallOpts
}

type DebugTransactorSession struct {
	Contract     *DebugTransactor
	TransactOpts bind.TransactOpts
}

type DebugRaw struct {
	Contract *Debug
}

type DebugCallerRaw struct {
	Contract *DebugCaller
}

type DebugTransactorRaw struct {
	Contract *DebugTransactor
}

func NewDebug(address common.Address, backend bind.ContractBackend) (*Debug, error) {
	contract, err := bindDebug(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Debug{DebugCaller: DebugCaller{contract: contract}, DebugTransactor: DebugTransactor{contract: contract}, DebugFilterer: DebugFilterer{contract: contract}}, nil
}

func NewDebugCaller(address common.Address, caller bind.ContractCaller) (*DebugCaller, error) {
	contract, err := bindDebug(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugCaller{contract: contract}, nil
}

func NewDebugTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugTransactor, error) {
	contract, err := bindDebug(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugTransactor{contract: contract}, nil
}

func NewDebugFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugFilterer, error) {
	contract, err := bindDebug(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugFilterer{contract: contract}, nil
}

func bindDebug(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_Debug *DebugRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Debug.Contract.DebugCaller.contract.Call(opts, result, method, params...)
}

func (_Debug *DebugRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Debug.Contract.DebugTransactor.contract.Transfer(opts)
}

func (_Debug *DebugRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Debug.Contract.DebugTransactor.contract.Transact(opts, method, params...)
}

func (_Debug *DebugCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Debug.Contract.contract.Call(opts, result, method, params...)
}

func (_Debug *DebugTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Debug.Contract.contract.Transfer(opts)
}

func (_Debug *DebugTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Debug.Contract.contract.Transact(opts, method, params...)
}

func (_Debug *DebugCaller) AddressToString(opts *bind.CallOpts, a common.Address) (string, error) {
	var out []interface{}
	err := _Debug.contract.Call(opts, &out, "addressToString", a)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Debug *DebugSession) AddressToString(a common.Address) (string, error) {
	return _Debug.Contract.AddressToString(&_Debug.CallOpts, a)
}

func (_Debug *DebugCallerSession) AddressToString(a common.Address) (string, error) {
	return _Debug.Contract.AddressToString(&_Debug.CallOpts, a)
}

func (_Debug *DebugCaller) Bytes32ToString(opts *bind.CallOpts, s [32]byte) (string, error) {
	var out []interface{}
	err := _Debug.contract.Call(opts, &out, "bytes32ToString", s)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Debug *DebugSession) Bytes32ToString(s [32]byte) (string, error) {
	return _Debug.Contract.Bytes32ToString(&_Debug.CallOpts, s)
}

func (_Debug *DebugCallerSession) Bytes32ToString(s [32]byte) (string, error) {
	return _Debug.Contract.Bytes32ToString(&_Debug.CallOpts, s)
}

func (_Debug *DebugCaller) BytesToString(opts *bind.CallOpts, _bytes []byte) (string, error) {
	var out []interface{}
	err := _Debug.contract.Call(opts, &out, "bytesToString", _bytes)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Debug *DebugSession) BytesToString(_bytes []byte) (string, error) {
	return _Debug.Contract.BytesToString(&_Debug.CallOpts, _bytes)
}

func (_Debug *DebugCallerSession) BytesToString(_bytes []byte) (string, error) {
	return _Debug.Contract.BytesToString(&_Debug.CallOpts, _bytes)
}

func (_Debug *DebugCaller) ToASCII(opts *bind.CallOpts, _uint8 uint8) (uint8, error) {
	var out []interface{}
	err := _Debug.contract.Call(opts, &out, "toASCII", _uint8)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Debug *DebugSession) ToASCII(_uint8 uint8) (uint8, error) {
	return _Debug.Contract.ToASCII(&_Debug.CallOpts, _uint8)
}

func (_Debug *DebugCallerSession) ToASCII(_uint8 uint8) (uint8, error) {
	return _Debug.Contract.ToASCII(&_Debug.CallOpts, _uint8)
}

var ECCArithmeticMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50601680601d6000396000f3fe6080604052600080fdfea164736f6c6343000813000a",
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

var HashToCurveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"m\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f2\",\"type\":\"tuple\"}],\"name\":\"hashToCurve\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"hashPoint\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610de6806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e95d957714610030575b600080fd5b61004361003e366004610c8d565b610059565b6040516100509190610ccc565b60405180910390f35b610061610bbb565b600061006c856100a7565b9050600061008182825b6020020151866101b0565b90506000610090836001610076565b905061009c8282610928565b979650505050505050565b6100af610bd3565b6000805b60028210156101a95760408051602081018690520160408051601f1981840301815291905280516020909101209350836100fc600080516020610dba8339815191526005610d2b565b8110156101405761011b600080516020610dba83398151915282610d58565b84846002811061012d5761012d610cff565b60200201528261013c81610d7a565b9350505b8161014a81610d7a565b92505060208211156101a35760405162461bcd60e51b815260206004820152601960248201527f617474656d7074656420746f6f206d616e79206861736865730000000000000060448201526064015b60405180910390fd5b506100b3565b5050919050565b6101b8610bbb565b826000036101fc5750604080516060810182527759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe60208201908152600292820192909252908152610922565b610215600080516020610dba8339815191526005610d2b565b83106102635760405162461bcd60e51b815260206004820152601660248201527f74206e6f74206120756e69666f726d2073616d706c6500000000000000000000604482015260640161019a565b6000600080516020610dba833981519152848509905060008161028860036001610d93565b6102929190610d93565b9050600080516020610dba833981519152845182096001146102f65760405162461bcd60e51b815260206004820152601760248201527f77726f6e6720696e766572736520666f722064656e6f6d000000000000000000604482015260640161019a565b6000600080516020610dba8339815191528551600080516020610dba8339815191528577b3c4d79d41a91759a9e4c7e359b6b89eaec68e62effffffd090990506000600080516020610dba83398151915261035f83600080516020610dba833981519152610da6565b7759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe0890506000610388600289610d58565b9050816000600080516020610dba8339815191526003600080516020610dba8339815191528085860985090890506000600080516020610dba83398151915260408b0151800990508181036104b957600080516020610dba8339815191528a60400151106104385760405162461bcd60e51b815260206004820152600c60248201527f793120746f6f206c617267650000000000000000000000000000000000000000604482015260640161019a565b8360028b6040015161044a9190610d58565b146104975760405162461bcd60e51b815260206004820152601860248201527f793120706172697479206d757374206d61746368207427730000000000000000604482015260640161019a565b885183905260408a0151895160015b6020020152506109229650505050505050565b6104d182600080516020610dba833981519152610da6565b811461051f5760405162461bcd60e51b815260206004820152601860248201527f7931213d70736575646f20737172206f662078315e332b420000000000000000604482015260640161019a565b5050506000600080516020610dba8339815191528061054057610540610d42565b8360010861055c90600080516020610dba833981519152610da6565b90506000600080516020610dba8339815191526003600080516020610dba8339815191528085860985090890506000600080516020610dba83398151915260608b01518009905081810361066c57600080516020610dba8339815191528a60600151106105fa5760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b604482015260640161019a565b8360028b6060015161060c9190610d58565b146106595760405162461bcd60e51b815260206004820152601860248201527f793220706172697479206d757374206d61746368207427730000000000000000604482015260640161019a565b885183905260608a0151895160016104a6565b61068482600080516020610dba833981519152610da6565b81146106d25760405162461bcd60e51b815260206004820152601860248201527f7932213d70736575646f20737172206f662078325e332b420000000000000000604482015260640161019a565b505050600080516020610dba833981519152806106f1576106f1610d42565b876020015186096001146107475760405162461bcd60e51b815260206004820152601c60248201527f74496e76537175617265642a742a2a3220213d3d2031206d6f64205000000000604482015260640161019a565b600080516020610dba8339815191527f2042def740cbc01bd03583cf0100e593ba56470b9af68708d2c05d6490535385600080516020610dba83398151915260208a0151600080516020610dba833981519152888909090992506000600080516020610dba8339815191526107ca85600080516020610dba833981519152610da6565b60010890506000600080516020610dba8339815191526003600080516020610dba83398151915280858609850908905080600080516020610dba83398151915260808b015180091461085e5760405162461bcd60e51b815260206004820152601c60248201527f646964206e6f74206f627461696e206120637572766520706f696e7400000000604482015260640161019a565b600080516020610dba8339815191528960800151106108ae5760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b604482015260640161019a565b8260028a608001516108c09190610d58565b1461090d5760405162461bcd60e51b815260206004820152601860248201527f793320706172697479206d757374206d61746368207427730000000000000000604482015260640161019a565b50865152505050506080830151825160200152505b92915050565b610930610bbb565b600061093c84846109a7565b80515190915015801590610954575080516020015115155b6109a05760405162461bcd60e51b815260206004820152601b60248201527f6164646731206661696c65643a207a65726f206f7264696e6174650000000000604482015260640161019a565b9392505050565b6109af610bbb565b6109b883610a6b565b6109c182610a6b565b6109c9610bf1565b835151815283516020908101518282015283515160408301528351015160608201526109f3610bd3565b600060408260808560066096fa905080600003610a525760405162461bcd60e51b815260206004820152601160248201527f61646467312063616c6c206661696c6564000000000000000000000000000000604482015260640161019a565b5080518351526020908101518351909101525092915050565b805151600080516020610dba83398151915211610aca5760405162461bcd60e51b815260206004820152600c60248201527f78206e6f7420696e20465f500000000000000000000000000000000000000000604482015260640161019a565b805160200151600080516020610dba83398151915211610b2c5760405162461bcd60e51b815260206004820152600c60248201527f79206e6f7420696e20465f500000000000000000000000000000000000000000604482015260640161019a565b805151600090600080516020610dba8339815191529060039082908181800909088251602001519091508190600080516020610dba83398151915290800914610bb75760405162461bcd60e51b815260206004820152601260248201527f706f696e74206e6f74206f6e2063757276650000000000000000000000000000604482015260640161019a565b5050565b6040518060200160405280610bce610bd3565b905290565b60405180604001604052806002906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b600060a08284031215610c2157600080fd5b60405160a0810181811067ffffffffffffffff82111715610c5257634e487b7160e01b600052604160045260246000fd5b806040525080915082358152602083013560208201526040830135604082015260608301356060820152608083013560808201525092915050565b60008060006101608486031215610ca357600080fd5b83359250610cb48560208601610c0f565b9150610cc38560c08601610c0f565b90509250925092565b815160408201908260005b6002811015610cf6578251825260209283019290910190600101610cd7565b50505092915050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808202811582820484141761092257610922610d15565b634e487b7160e01b600052601260045260246000fd5b600082610d7557634e487b7160e01b600052601260045260246000fd5b500690565b600060018201610d8c57610d8c610d15565b5060010190565b8082018082111561092257610922610d15565b8181038181111561092257610922610d1556fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a164736f6c6343000813000a",
}

var HashToCurveABI = HashToCurveMetaData.ABI

var HashToCurveBin = HashToCurveMetaData.Bin

func DeployHashToCurve(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashToCurve, error) {
	parsed, err := HashToCurveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashToCurveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashToCurve{HashToCurveCaller: HashToCurveCaller{contract: contract}, HashToCurveTransactor: HashToCurveTransactor{contract: contract}, HashToCurveFilterer: HashToCurveFilterer{contract: contract}}, nil
}

type HashToCurve struct {
	HashToCurveCaller
	HashToCurveTransactor
	HashToCurveFilterer
}

type HashToCurveCaller struct {
	contract *bind.BoundContract
}

type HashToCurveTransactor struct {
	contract *bind.BoundContract
}

type HashToCurveFilterer struct {
	contract *bind.BoundContract
}

type HashToCurveSession struct {
	Contract     *HashToCurve
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type HashToCurveCallerSession struct {
	Contract *HashToCurveCaller
	CallOpts bind.CallOpts
}

type HashToCurveTransactorSession struct {
	Contract     *HashToCurveTransactor
	TransactOpts bind.TransactOpts
}

type HashToCurveRaw struct {
	Contract *HashToCurve
}

type HashToCurveCallerRaw struct {
	Contract *HashToCurveCaller
}

type HashToCurveTransactorRaw struct {
	Contract *HashToCurveTransactor
}

func NewHashToCurve(address common.Address, backend bind.ContractBackend) (*HashToCurve, error) {
	contract, err := bindHashToCurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashToCurve{HashToCurveCaller: HashToCurveCaller{contract: contract}, HashToCurveTransactor: HashToCurveTransactor{contract: contract}, HashToCurveFilterer: HashToCurveFilterer{contract: contract}}, nil
}

func NewHashToCurveCaller(address common.Address, caller bind.ContractCaller) (*HashToCurveCaller, error) {
	contract, err := bindHashToCurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashToCurveCaller{contract: contract}, nil
}

func NewHashToCurveTransactor(address common.Address, transactor bind.ContractTransactor) (*HashToCurveTransactor, error) {
	contract, err := bindHashToCurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashToCurveTransactor{contract: contract}, nil
}

func NewHashToCurveFilterer(address common.Address, filterer bind.ContractFilterer) (*HashToCurveFilterer, error) {
	contract, err := bindHashToCurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashToCurveFilterer{contract: contract}, nil
}

func bindHashToCurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashToCurveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_HashToCurve *HashToCurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashToCurve.Contract.HashToCurveCaller.contract.Call(opts, result, method, params...)
}

func (_HashToCurve *HashToCurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashToCurve.Contract.HashToCurveTransactor.contract.Transfer(opts)
}

func (_HashToCurve *HashToCurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashToCurve.Contract.HashToCurveTransactor.contract.Transact(opts, method, params...)
}

func (_HashToCurve *HashToCurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashToCurve.Contract.contract.Call(opts, result, method, params...)
}

func (_HashToCurve *HashToCurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashToCurve.Contract.contract.Transfer(opts)
}

func (_HashToCurve *HashToCurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashToCurve.Contract.contract.Transact(opts, method, params...)
}

func (_HashToCurve *HashToCurveCaller) HashToCurve(opts *bind.CallOpts, m [32]byte, f1 HashToCurveFProof, f2 HashToCurveFProof) (ECCArithmeticG1Point, error) {
	var out []interface{}
	err := _HashToCurve.contract.Call(opts, &out, "hashToCurve", m, f1, f2)

	if err != nil {
		return *new(ECCArithmeticG1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(ECCArithmeticG1Point)).(*ECCArithmeticG1Point)

	return out0, err

}

func (_HashToCurve *HashToCurveSession) HashToCurve(m [32]byte, f1 HashToCurveFProof, f2 HashToCurveFProof) (ECCArithmeticG1Point, error) {
	return _HashToCurve.Contract.HashToCurve(&_HashToCurve.CallOpts, m, f1, f2)
}

func (_HashToCurve *HashToCurveCallerSession) HashToCurve(m [32]byte, f1 HashToCurveFProof, f2 HashToCurveFProof) (ECCArithmeticG1Point, error) {
	return _HashToCurve.Contract.HashToCurve(&_HashToCurve.CallOpts, m, f1, f2)
}

var IVRFConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IVRFConsumerABI = IVRFConsumerMetaData.ABI

type IVRFConsumer struct {
	IVRFConsumerCaller
	IVRFConsumerTransactor
	IVRFConsumerFilterer
}

type IVRFConsumerCaller struct {
	contract *bind.BoundContract
}

type IVRFConsumerTransactor struct {
	contract *bind.BoundContract
}

type IVRFConsumerFilterer struct {
	contract *bind.BoundContract
}

type IVRFConsumerSession struct {
	Contract     *IVRFConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFConsumerCallerSession struct {
	Contract *IVRFConsumerCaller
	CallOpts bind.CallOpts
}

type IVRFConsumerTransactorSession struct {
	Contract     *IVRFConsumerTransactor
	TransactOpts bind.TransactOpts
}

type IVRFConsumerRaw struct {
	Contract *IVRFConsumer
}

type IVRFConsumerCallerRaw struct {
	Contract *IVRFConsumerCaller
}

type IVRFConsumerTransactorRaw struct {
	Contract *IVRFConsumerTransactor
}

func NewIVRFConsumer(address common.Address, backend bind.ContractBackend) (*IVRFConsumer, error) {
	contract, err := bindIVRFConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFConsumer{IVRFConsumerCaller: IVRFConsumerCaller{contract: contract}, IVRFConsumerTransactor: IVRFConsumerTransactor{contract: contract}, IVRFConsumerFilterer: IVRFConsumerFilterer{contract: contract}}, nil
}

func NewIVRFConsumerCaller(address common.Address, caller bind.ContractCaller) (*IVRFConsumerCaller, error) {
	contract, err := bindIVRFConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFConsumerCaller{contract: contract}, nil
}

func NewIVRFConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFConsumerTransactor, error) {
	contract, err := bindIVRFConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFConsumerTransactor{contract: contract}, nil
}

func NewIVRFConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFConsumerFilterer, error) {
	contract, err := bindIVRFConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFConsumerFilterer{contract: contract}, nil
}

func bindIVRFConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFConsumer *IVRFConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFConsumer.Contract.IVRFConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_IVRFConsumer *IVRFConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFConsumer.Contract.IVRFConsumerTransactor.contract.Transfer(opts)
}

func (_IVRFConsumer *IVRFConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFConsumer.Contract.IVRFConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_IVRFConsumer *IVRFConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFConsumer *IVRFConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFConsumer.Contract.contract.Transfer(opts)
}

func (_IVRFConsumer *IVRFConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFConsumer *IVRFConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID [32]byte, randomWords []*big.Int) (*types.Transaction, error) {
	return _IVRFConsumer.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords)
}

func (_IVRFConsumer *IVRFConsumerSession) RawFulfillRandomWords(requestID [32]byte, randomWords []*big.Int) (*types.Transaction, error) {
	return _IVRFConsumer.Contract.RawFulfillRandomWords(&_IVRFConsumer.TransactOpts, requestID, randomWords)
}

func (_IVRFConsumer *IVRFConsumerTransactorSession) RawFulfillRandomWords(requestID [32]byte, randomWords []*big.Int) (*types.Transaction, error) {
	return _IVRFConsumer.Contract.RawFulfillRandomWords(&_IVRFConsumer.TransactOpts, requestID, randomWords)
}

var KeyDataStructMetaData = &bind.MetaData{
	ABI: "[]",
}

var KeyDataStructABI = KeyDataStructMetaData.ABI

type KeyDataStruct struct {
	KeyDataStructCaller
	KeyDataStructTransactor
	KeyDataStructFilterer
}

type KeyDataStructCaller struct {
	contract *bind.BoundContract
}

type KeyDataStructTransactor struct {
	contract *bind.BoundContract
}

type KeyDataStructFilterer struct {
	contract *bind.BoundContract
}

type KeyDataStructSession struct {
	Contract     *KeyDataStruct
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type KeyDataStructCallerSession struct {
	Contract *KeyDataStructCaller
	CallOpts bind.CallOpts
}

type KeyDataStructTransactorSession struct {
	Contract     *KeyDataStructTransactor
	TransactOpts bind.TransactOpts
}

type KeyDataStructRaw struct {
	Contract *KeyDataStruct
}

type KeyDataStructCallerRaw struct {
	Contract *KeyDataStructCaller
}

type KeyDataStructTransactorRaw struct {
	Contract *KeyDataStructTransactor
}

func NewKeyDataStruct(address common.Address, backend bind.ContractBackend) (*KeyDataStruct, error) {
	contract, err := bindKeyDataStruct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyDataStruct{KeyDataStructCaller: KeyDataStructCaller{contract: contract}, KeyDataStructTransactor: KeyDataStructTransactor{contract: contract}, KeyDataStructFilterer: KeyDataStructFilterer{contract: contract}}, nil
}

func NewKeyDataStructCaller(address common.Address, caller bind.ContractCaller) (*KeyDataStructCaller, error) {
	contract, err := bindKeyDataStruct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyDataStructCaller{contract: contract}, nil
}

func NewKeyDataStructTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyDataStructTransactor, error) {
	contract, err := bindKeyDataStruct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyDataStructTransactor{contract: contract}, nil
}

func NewKeyDataStructFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyDataStructFilterer, error) {
	contract, err := bindKeyDataStruct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyDataStructFilterer{contract: contract}, nil
}

func bindKeyDataStruct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyDataStructABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_KeyDataStruct *KeyDataStructRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyDataStruct.Contract.KeyDataStructCaller.contract.Call(opts, result, method, params...)
}

func (_KeyDataStruct *KeyDataStructRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyDataStruct.Contract.KeyDataStructTransactor.contract.Transfer(opts)
}

func (_KeyDataStruct *KeyDataStructRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyDataStruct.Contract.KeyDataStructTransactor.contract.Transact(opts, method, params...)
}

func (_KeyDataStruct *KeyDataStructCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyDataStruct.Contract.contract.Call(opts, result, method, params...)
}

func (_KeyDataStruct *KeyDataStructTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyDataStruct.Contract.contract.Transfer(opts)
}

func (_KeyDataStruct *KeyDataStructTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyDataStruct.Contract.contract.Transact(opts, method, params...)
}

var OCR2AbstractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

var OCR2AbstractABI = OCR2AbstractMetaData.ABI

type OCR2Abstract struct {
	OCR2AbstractCaller
	OCR2AbstractTransactor
	OCR2AbstractFilterer
}

type OCR2AbstractCaller struct {
	contract *bind.BoundContract
}

type OCR2AbstractTransactor struct {
	contract *bind.BoundContract
}

type OCR2AbstractFilterer struct {
	contract *bind.BoundContract
}

type OCR2AbstractSession struct {
	Contract     *OCR2Abstract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OCR2AbstractCallerSession struct {
	Contract *OCR2AbstractCaller
	CallOpts bind.CallOpts
}

type OCR2AbstractTransactorSession struct {
	Contract     *OCR2AbstractTransactor
	TransactOpts bind.TransactOpts
}

type OCR2AbstractRaw struct {
	Contract *OCR2Abstract
}

type OCR2AbstractCallerRaw struct {
	Contract *OCR2AbstractCaller
}

type OCR2AbstractTransactorRaw struct {
	Contract *OCR2AbstractTransactor
}

func NewOCR2Abstract(address common.Address, backend bind.ContractBackend) (*OCR2Abstract, error) {
	contract, err := bindOCR2Abstract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OCR2Abstract{OCR2AbstractCaller: OCR2AbstractCaller{contract: contract}, OCR2AbstractTransactor: OCR2AbstractTransactor{contract: contract}, OCR2AbstractFilterer: OCR2AbstractFilterer{contract: contract}}, nil
}

func NewOCR2AbstractCaller(address common.Address, caller bind.ContractCaller) (*OCR2AbstractCaller, error) {
	contract, err := bindOCR2Abstract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OCR2AbstractCaller{contract: contract}, nil
}

func NewOCR2AbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*OCR2AbstractTransactor, error) {
	contract, err := bindOCR2Abstract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OCR2AbstractTransactor{contract: contract}, nil
}

func NewOCR2AbstractFilterer(address common.Address, filterer bind.ContractFilterer) (*OCR2AbstractFilterer, error) {
	contract, err := bindOCR2Abstract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OCR2AbstractFilterer{contract: contract}, nil
}

func bindOCR2Abstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OCR2AbstractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_OCR2Abstract *OCR2AbstractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OCR2Abstract.Contract.OCR2AbstractCaller.contract.Call(opts, result, method, params...)
}

func (_OCR2Abstract *OCR2AbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.OCR2AbstractTransactor.contract.Transfer(opts)
}

func (_OCR2Abstract *OCR2AbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.OCR2AbstractTransactor.contract.Transact(opts, method, params...)
}

func (_OCR2Abstract *OCR2AbstractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OCR2Abstract.Contract.contract.Call(opts, result, method, params...)
}

func (_OCR2Abstract *OCR2AbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.contract.Transfer(opts)
}

func (_OCR2Abstract *OCR2AbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.contract.Transact(opts, method, params...)
}

func (_OCR2Abstract *OCR2AbstractCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _OCR2Abstract.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(struct {
		ConfigCount  uint32
		BlockNumber  uint32
		ConfigDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_OCR2Abstract *OCR2AbstractSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _OCR2Abstract.Contract.LatestConfigDetails(&_OCR2Abstract.CallOpts)
}

func (_OCR2Abstract *OCR2AbstractCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _OCR2Abstract.Contract.LatestConfigDetails(&_OCR2Abstract.CallOpts)
}

func (_OCR2Abstract *OCR2AbstractCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _OCR2Abstract.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(struct {
		ScanLogs     bool
		ConfigDigest [32]byte
		Epoch        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_OCR2Abstract *OCR2AbstractSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _OCR2Abstract.Contract.LatestConfigDigestAndEpoch(&_OCR2Abstract.CallOpts)
}

func (_OCR2Abstract *OCR2AbstractCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _OCR2Abstract.Contract.LatestConfigDigestAndEpoch(&_OCR2Abstract.CallOpts)
}

func (_OCR2Abstract *OCR2AbstractCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OCR2Abstract.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OCR2Abstract *OCR2AbstractSession) TypeAndVersion() (string, error) {
	return _OCR2Abstract.Contract.TypeAndVersion(&_OCR2Abstract.CallOpts)
}

func (_OCR2Abstract *OCR2AbstractCallerSession) TypeAndVersion() (string, error) {
	return _OCR2Abstract.Contract.TypeAndVersion(&_OCR2Abstract.CallOpts)
}

func (_OCR2Abstract *OCR2AbstractTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OCR2Abstract.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_OCR2Abstract *OCR2AbstractSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.SetConfig(&_OCR2Abstract.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_OCR2Abstract *OCR2AbstractTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.SetConfig(&_OCR2Abstract.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_OCR2Abstract *OCR2AbstractTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OCR2Abstract.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_OCR2Abstract *OCR2AbstractSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.Transmit(&_OCR2Abstract.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OCR2Abstract *OCR2AbstractTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OCR2Abstract.Contract.Transmit(&_OCR2Abstract.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type OCR2AbstractConfigSetIterator struct {
	Event *OCR2AbstractConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AbstractConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AbstractConfigSet)
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
		it.Event = new(OCR2AbstractConfigSet)
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

func (it *OCR2AbstractConfigSetIterator) Error() error {
	return it.fail
}

func (it *OCR2AbstractConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AbstractConfigSet struct {
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

func (_OCR2Abstract *OCR2AbstractFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OCR2AbstractConfigSetIterator, error) {

	logs, sub, err := _OCR2Abstract.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OCR2AbstractConfigSetIterator{contract: _OCR2Abstract.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OCR2Abstract *OCR2AbstractFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OCR2AbstractConfigSet) (event.Subscription, error) {

	logs, sub, err := _OCR2Abstract.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AbstractConfigSet)
				if err := _OCR2Abstract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OCR2Abstract *OCR2AbstractFilterer) ParseConfigSet(log types.Log) (*OCR2AbstractConfigSet, error) {
	event := new(OCR2AbstractConfigSet)
	if err := _OCR2Abstract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OCR2AbstractTransmittedIterator struct {
	Event *OCR2AbstractTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OCR2AbstractTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OCR2AbstractTransmitted)
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
		it.Event = new(OCR2AbstractTransmitted)
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

func (it *OCR2AbstractTransmittedIterator) Error() error {
	return it.fail
}

func (it *OCR2AbstractTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OCR2AbstractTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_OCR2Abstract *OCR2AbstractFilterer) FilterTransmitted(opts *bind.FilterOpts) (*OCR2AbstractTransmittedIterator, error) {

	logs, sub, err := _OCR2Abstract.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &OCR2AbstractTransmittedIterator{contract: _OCR2Abstract.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_OCR2Abstract *OCR2AbstractFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OCR2AbstractTransmitted) (event.Subscription, error) {

	logs, sub, err := _OCR2Abstract.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OCR2AbstractTransmitted)
				if err := _OCR2Abstract.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_OCR2Abstract *OCR2AbstractFilterer) ParseTransmitted(log types.Log) (*OCR2AbstractTransmitted, error) {
	event := new(OCR2AbstractTransmitted)
	if err := _OCR2Abstract.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6102a8806101576000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c6343000813000a",
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

var VRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDKG\",\"name\":\"_keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"output\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addressToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structVRF.Request\",\"name\":\"r\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256[4]\",\"name\":\"p\",\"type\":\"uint256[4]\"}],\"internalType\":\"structECCArithmetic.G2Point\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f2\",\"type\":\"tuple\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"m\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f2\",\"type\":\"tuple\"}],\"name\":\"hashToCurve\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"hashPoint\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"s_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"transmitVRFResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"input\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256[4]\",\"name\":\"p\",\"type\":\"uint256[4]\"}],\"internalType\":\"structECCArithmetic.G2Point\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f2\",\"type\":\"tuple\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"vrfOutput\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"output\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162003a7338038062003a73833981016040819052620000349162000196565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000eb565b5050600280546001600160a01b0319166001600160a01b03949094169390931790925560035550620001d2565b336001600160a01b03821603620001455760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008060408385031215620001aa57600080fd5b82516001600160a01b0381168114620001c257600080fd5b6020939093015192949293505050565b61389180620001e26000396000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c80639201de55116100d8578063bf2732c71161008c578063e3d0e71211610066578063e3d0e71214610367578063e95d95771461037a578063f2fde38b1461039a57600080fd5b8063bf2732c714610342578063cc31f7dd14610355578063d57fc45a1461035e57600080fd5b8063afcb95d7116100bd578063afcb95d7146102ed578063b1dc65a414610317578063bf0e15c61461032f57600080fd5b80639201de55146102c7578063a954b4ef146102da57600080fd5b806339614e4f1161013a57806379ba50971161011457806379ba50971461027757806381ff70481461027f5780638da5cb5b146102ac57600080fd5b806339614e4f1461024957806355e487491461025c5780635e57966d1461026457600080fd5b80630bc643e81161016b5780630bc643e8146101c25780630e3ca2a7146101e7578063181f5a771461020757600080fd5b8063012cfe8614610187578063092576f71461019c575b600080fd5b61019a6101953660046129fb565b6103ad565b005b6101af6101aa366004612a61565b610523565b6040519081526020015b60405180910390f35b6101d56101d0366004612aa2565b61066d565b60405160ff90911681526020016101b9565b6101af6101f5366004612ad2565b60056020526000908152604090205481565b60408051808201909152600981527f56524620302e302e31000000000000000000000000000000000000000000000060208201525b6040516101b99190612b3f565b61023c610257366004612c9b565b610696565b61019a610818565b61023c610272366004612ad2565b610879565b61019a610940565b600a54600854604080516000815264010000000090930463ffffffff1660208401528201526060016101b9565b6000546040516001600160a01b0390911681526020016101b9565b61023c6102d5366004612cd0565b6109f1565b61019a6102e8366004612e50565b610a7d565b6008546007546040805160008152602081019390935263ffffffff909116908201526060016101b9565b61019a610325366004612f0d565b5050505050505050565b6101af61033d366004612fc0565b610d11565b61019a610350366004613009565b610eba565b6101af60035481565b6101af60045481565b61019a61037536600461317e565b610f44565b61038d61038836600461324b565b6116a5565b6040516101b9919061328a565b61019a6103a8366004612ad2565b6116f3565b6004546000906104045760405162461bcd60e51b815260206004820152601060248201527f6e6f206b657920617661696c61626c650000000000000000000000000000000060448201526064015b60405180910390fd5b60005a604080516020601f870181900481028201810190925285815291925086359181880135916104549184918491908a908a908190840183828082843760009201919091525061170792505050565b6040805183815263ffffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260085480825260095460ff808216602085015261010090910416928201929092529083146105105760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016103fb565b63ffffffff8410610325576103256132ca565b6040805160808101825260008082526020808301828152838501838152336060808701828152918652600585528786205460035492518951921b6bffffffffffffffffffffffff191682870152603482018390526054808301829052895180840390910181526074909201909852805194019390932085529087905263ffffffff861690529092906105b483611756565b83516000908152600660209081526040918290209290925584519185015181860151606087015192517fda543d8fd5d52cb865899d85adee45422306c16f47e6e4394f043006ff5cde30946106319490939291938452602084019290925263ffffffff1660408301526001600160a01b0316606082015260800190565b60405180910390a16106448260016132f6565b60608401516001600160a01b031660009081526005602052604090205550505190505b92915050565b6000600a8260ff16101561068657610667826030613309565b610667826057613309565b919050565b6060600080835160026106a99190613322565b67ffffffffffffffff8111156106c1576106c1612b52565b6040519080825280601f01601f1916602001820160405280156106eb576020820181803683370190505b509050600091505b80518260ff1610156108115760008461070d60028561334f565b60ff1681518110610720576107206132b4565b60209101015160f81c600f169050600060048661073e60028761334f565b60ff1681518110610751576107516132b4565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c90506107888161066d565b60f81b838560ff16815181106107a0576107a06132b4565b60200101906001600160f81b031916908160001a9053506107c2846001613309565b93506107cd8261066d565b60f81b838560ff16815181106107e5576107e56132b4565b60200101906001600160f81b031916908160001a9053505050818061080990613371565b9250506106f3565b9392505050565b6002546001600160a01b031633146108725760405162461bcd60e51b815260206004820181905260248201527f6b657920696e666f206d75737420636f6d652066726f6d2070726f766964657260448201526064016103fb565b6000600455565b604080516014808252818301909252606091600091906020820181803683370190505090508260005b60148160ff16101561092e577fff0000000000000000000000000000000000000000000000000000000000000060f883901b16836108e1836013613390565b60ff16815181106108f4576108f46132b4565b60200101906001600160f81b031916908160001a9053506008826001600160a01b0316901c9150808061092690613371565b9150506108a2565b5061093882610696565b949350505050565b6001546001600160a01b0316331461099a5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103fb565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b6020811015610a7357838160208110610a3457610a346132b4565b1a60f81b828281518110610a4a57610a4a6132b4565b60200101906001600160f81b031916908160001a90535080610a6b816133a9565b915050610a19565b5061081181610696565b815160009081526006602052604081205490610a9884611756565b9050808214610ae95760405162461bcd60e51b815260206004820152601560248201527f72657175657374206c6f6f6b7570206661696c6564000000000000000000000060448201526064016103fb565b6000610af58385610d11565b85516000908152600660205260408082208290558701519192509063ffffffff1667ffffffffffffffff811115610b2e57610b2e612b52565b604051908082528060200260200182016040528015610b57578160200160208202803683370190505b50905060005b866040015163ffffffff16811015610bcb5760408051602081018590529081018290526060016040516020818303038152906040528051906020012060001c828281518110610bae57610bae6132b4565b602090810291909101015280610bc3816133a9565b915050610b5d565b50606086015186516040517f75bf929b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916375bf929b91610c199185906004016133fd565b600060405180830381600087803b158015610c3357600080fd5b505af1925050508015610c44575060015b610cbd573d808015610c72576040519150601f19603f3d011682016040523d82523d6000602084013e610c77565b606091505b5086516040517fa7231a311a37fec0b9b631c5b7d4c4aa3effe5304f25dbfeaf0de676cdd715ba90610caf9085906000908690613416565b60405180910390a250610d09565b8551604080516000815260208101918290527fa7231a311a37fec0b9b631c5b7d4c4aa3effe5304f25dbfeaf0de676cdd715ba91610d0091859160019190613416565b60405180910390a25b505050505050565b60045481515160405160009291610d2a91602001613443565b6040516020818303038152906040528051906020012014610d8d5760405162461bcd60e51b815260206004820152601060248201527f77726f6e67207075626c6963206b65790000000000000000000000000000000060448201526064016103fb565b6000610da284846040015185606001516116a5565b9050610db781846020015185600001516117ec565b610e035760405162461bcd60e51b815260206004820152600d60248201527f626164205652462070726f6f660000000000000000000000000000000000000060448201526064016103fb565b60208301515151600080516020613865833981519152118015610e3b5750602083810151510151600080516020613865833981519152115b610e875760405162461bcd60e51b815260206004820152601f60248201527f62616420726570726573656e746174696f6e206f66206f75747075742070740060448201526064016103fb565b60208084015151604051610e9b9201613477565b6040516020818303038152906040528051906020012091505092915050565b6002546001600160a01b03163314610f145760405162461bcd60e51b815260206004820181905260248201527f6b657920696e666f206d75737420636f6d652066726f6d2070726f766964657260448201526064016103fb565b8051604051610f2691906020016134ab565b60408051601f19818403018152919052805160209091012060045550565b855185518560ff16601f831115610f9d5760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016103fb565b60008111610fed5760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016103fb565b8183146110615760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016103fb565b61106c816003613322565b83116110ba5760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016103fb565b6110c26119d5565b60006040518060c001604052808b81526020018a81526020018960ff1681526020018881526020018767ffffffffffffffff1681526020018681525090505b600c541561121357600c5460009061111b906001906134c7565b90506000600c8281548110611132576111326132b4565b6000918252602082200154600d80546001600160a01b039092169350908490811061115f5761115f6132b4565b60009182526020808320909101546001600160a01b038581168452600b9092526040808420805461ffff1990811690915592909116808452922080549091169055600c805491925090806111b5576111b56134da565b600082815260209020810160001990810180546001600160a01b0319169055019055600d8054806111e8576111e86134da565b600082815260209020810160001990810180546001600160a01b031916905501905550611101915050565b60005b815151811015611552576000600b60008460000151848151811061123c5761123c6132b4565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115611279576112796134f0565b146112c65760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016103fb565b6040805180820190915260ff821681526001602082015282518051600b91600091859081106112f7576112f76132b4565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115611350576113506134f0565b0217905550600091506113609050565b600b60008460200151848151811061137a5761137a6132b4565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156113b7576113b76134f0565b146114045760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016103fb565b6040805180820190915260ff82168152602081016002815250600b600084602001518481518110611437576114376132b4565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115611490576114906134f0565b02179055505082518051600c9250839081106114ae576114ae6132b4565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b03909316929092179091558201518051600d919083908110611505576115056132b4565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b039092169190911790558061154a816133a9565b915050611216565b5060408101516009805460ff191660ff909216919091179055600a805467ffffffff0000000019811664010000000063ffffffff4381168202928317855590830481169360019390926000926115af928692908216911617613506565b92506101000a81548163ffffffff021916908363ffffffff16021790555061160e4630600a60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151611a31565b6008819055825180516009805460ff9092166101000261ff0019909216919091179055600a5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598611690988b98919763ffffffff909216969095919491939192613563565b60405180910390a15050505050505050505050565b6116ad61291c565b60006116b885611abe565b905060006116cd82825b602002015186611bc2565b905060006116dc8360016116c2565b90506116e8828261233a565b979650505050505050565b6116fb6119d5565b611704816123b2565b50565b6000808280602001905181019061171e919061369d565b9150915061172c8282610a7d565b50506007805460089390931c63ffffffff1663ffffffff19909316929092179091555050565b5050565b60008082600001518360200151846040015185606001516040516020016117cd9493929190938452602084019290925260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016604083015260601b6bffffffffffffffffffffffff1916604482015260580190565b60408051601f1981840301815291905280516020909101209392505050565b60408051600280825260608201909252600091829190816020015b61180f61291c565b8152602001906001900390816118075750506040805160028082526060820190925291925060009190602082015b611845612934565b81526020019060019003908161183d5750506040805160608101825288515160208083019182528a51015193945090928392830190611892906000805160206138658339815191526134c7565b815250815250826000815181106118ab576118ab6132b4565b602002602001018190525084826001815181106118ca576118ca6132b4565b602002602001018190525083816000815181106118e9576118e96132b4565b6020026020010181905250604051806020016040528060405180608001604052807f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281526020017f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed81526020017f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b81526020017f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa815250815250816001815181106119b6576119b66132b4565b60200260200101819052506119cb828261245b565b9695505050505050565b6000546001600160a01b03163314611a2f5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103fb565b565b6000808a8a8a8a8a8a8a8a8a604051602001611a55999897969594939291906137c8565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b611ac6612947565b6000805b6002821015611bbb5760408051602081018690520160408051601f198184030181529190528051602090910120935083611b136000805160206138658339815191526005613322565b811015611b5757611b3260008051602061386583398151915282613850565b848460028110611b4457611b446132b4565b602002015282611b53816133a9565b9350505b81611b61816133a9565b9250506020821115611bb55760405162461bcd60e51b815260206004820152601960248201527f617474656d7074656420746f6f206d616e79206861736865730000000000000060448201526064016103fb565b50611aca565b5050919050565b611bca61291c565b82600003611c0e5750604080516060810182527759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe60208201908152600292820192909252908152610667565b611c276000805160206138658339815191526005613322565b8310611c755760405162461bcd60e51b815260206004820152601660248201527f74206e6f74206120756e69666f726d2073616d706c650000000000000000000060448201526064016103fb565b60006000805160206138658339815191528485099050600081611c9a600360016132f6565b611ca491906132f6565b905060008051602061386583398151915284518209600114611d085760405162461bcd60e51b815260206004820152601760248201527f77726f6e6720696e766572736520666f722064656e6f6d00000000000000000060448201526064016103fb565b600060008051602061386583398151915285516000805160206138658339815191528577b3c4d79d41a91759a9e4c7e359b6b89eaec68e62effffffd090990506000600080516020613865833981519152611d71836000805160206138658339815191526134c7565b7759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe0890506000611d9a600289613850565b90508160006000805160206138658339815191526003600080516020613865833981519152808586098509089050600060008051602061386583398151915260408b015180099050818103611ecb576000805160206138658339815191528a6040015110611e4a5760405162461bcd60e51b815260206004820152600c60248201527f793120746f6f206c61726765000000000000000000000000000000000000000060448201526064016103fb565b8360028b60400151611e5c9190613850565b14611ea95760405162461bcd60e51b815260206004820152601860248201527f793120706172697479206d757374206d6174636820742773000000000000000060448201526064016103fb565b885183905260408a0151895160015b6020020152506106679650505050505050565b611ee3826000805160206138658339815191526134c7565b8114611f315760405162461bcd60e51b815260206004820152601860248201527f7931213d70736575646f20737172206f662078315e332b42000000000000000060448201526064016103fb565b505050600060008051602061386583398151915280611f5257611f52613339565b83600108611f6e906000805160206138658339815191526134c7565b905060006000805160206138658339815191526003600080516020613865833981519152808586098509089050600060008051602061386583398151915260608b01518009905081810361207e576000805160206138658339815191528a606001511061200c5760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b60448201526064016103fb565b8360028b6060015161201e9190613850565b1461206b5760405162461bcd60e51b815260206004820152601860248201527f793220706172697479206d757374206d6174636820742773000000000000000060448201526064016103fb565b885183905260608a015189516001611eb8565b612096826000805160206138658339815191526134c7565b81146120e45760405162461bcd60e51b815260206004820152601860248201527f7932213d70736575646f20737172206f662078325e332b42000000000000000060448201526064016103fb565b5050506000805160206138658339815191528061210357612103613339565b876020015186096001146121595760405162461bcd60e51b815260206004820152601c60248201527f74496e76537175617265642a742a2a3220213d3d2031206d6f6420500000000060448201526064016103fb565b6000805160206138658339815191527f2042def740cbc01bd03583cf0100e593ba56470b9af68708d2c05d649053538560008051602061386583398151915260208a01516000805160206138658339815191528889090909925060006000805160206138658339815191526121dc856000805160206138658339815191526134c7565b6001089050600060008051602061386583398151915260036000805160206138658339815191528085860985090890508060008051602061386583398151915260808b01518009146122705760405162461bcd60e51b815260206004820152601c60248201527f646964206e6f74206f627461696e206120637572766520706f696e740000000060448201526064016103fb565b6000805160206138658339815191528960800151106122c05760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b60448201526064016103fb565b8260028a608001516122d29190613850565b1461231f5760405162461bcd60e51b815260206004820152601860248201527f793320706172697479206d757374206d6174636820742773000000000000000060448201526064016103fb565b50865152505050506080929092015181516020015292915050565b61234261291c565b600061234e848461270c565b80515190915015801590612366575080516020015115155b6108115760405162461bcd60e51b815260206004820152601b60248201527f6164646731206661696c65643a207a65726f206f7264696e617465000000000060448201526064016103fb565b336001600160a01b0382160361240a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103fb565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000815183511461246b57600080fd5b8251600061247a826006613322565b905060008167ffffffffffffffff81111561249757612497612b52565b6040519080825280602002602001820160405280156124c0578160200160208202803683370190505b50905060005b838110156126d5578681815181106124e0576124e06132b4565b60209081029190910101515151826124f9836006613322565b6125049060006132f6565b81518110612514576125146132b4565b602002602001018181525050868181518110612532576125326132b4565b6020908102919091018101515101518261254d836006613322565b6125589060016132f6565b81518110612568576125686132b4565b602002602001018181525050858181518110612586576125866132b4565b602090810291909101015151518261259f836006613322565b6125aa9060026132f6565b815181106125ba576125ba6132b4565b6020026020010181815250508581815181106125d8576125d86132b4565b602090810291909101810151510151826125f3836006613322565b6125fe9060036132f6565b8151811061260e5761260e6132b4565b60200260200101818152505085818151811061262c5761262c6132b4565b6020908102919091010151516040015182612648836006613322565b6126539060046132f6565b81518110612663576126636132b4565b602002602001018181525050858181518110612681576126816132b4565b602090810291909101015151606001518261269d836006613322565b6126a89060056132f6565b815181106126b8576126b86132b4565b6020908102919091010152806126cd816133a9565b9150506124c6565b506126de612965565b6000602082602086026020860160086201b968fa9050806126fe57600080fd5b505115159695505050505050565b61271461291c565b61271d836127d0565b612726826127d0565b61272e612983565b83515181528351602090810151828201528351516040830152835101516060820152612758612947565b600060408260808560066096fa9050806000036127b75760405162461bcd60e51b815260206004820152601160248201527f61646467312063616c6c206661696c656400000000000000000000000000000060448201526064016103fb565b5080518351526020908101518351909101525092915050565b8051516000805160206138658339815191521161282f5760405162461bcd60e51b815260206004820152600c60248201527f78206e6f7420696e20465f50000000000000000000000000000000000000000060448201526064016103fb565b805160200151600080516020613865833981519152116128915760405162461bcd60e51b815260206004820152600c60248201527f79206e6f7420696e20465f50000000000000000000000000000000000000000060448201526064016103fb565b8051516000906000805160206138658339815191529060039082908181800909088251602001519091508190600080516020613865833981519152908009146117525760405162461bcd60e51b815260206004820152601260248201527f706f696e74206e6f74206f6e206375727665000000000000000000000000000060448201526064016103fb565b604051806020016040528061292f612947565b905290565b604051806020016040528061292f612983565b60405180604001604052806002906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b806060810183101561066757600080fd5b60008083601f8401126129c457600080fd5b50813567ffffffffffffffff8111156129dc57600080fd5b6020830191508360208285010111156129f457600080fd5b9250929050565b600080600060808486031215612a1057600080fd5b612a1a85856129a1565b9250606084013567ffffffffffffffff811115612a3657600080fd5b612a42868287016129b2565b9497909650939450505050565b63ffffffff8116811461170457600080fd5b60008060408385031215612a7457600080fd5b823591506020830135612a8681612a4f565b809150509250929050565b803560ff8116811461069157600080fd5b600060208284031215612ab457600080fd5b61081182612a91565b6001600160a01b038116811461170457600080fd5b600060208284031215612ae457600080fd5b813561081181612abd565b60005b83811015612b0a578181015183820152602001612af2565b50506000910152565b60008151808452612b2b816020860160208601612aef565b601f01601f19169290920160200192915050565b6020815260006108116020830184612b13565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715612b8b57612b8b612b52565b60405290565b6040516080810167ffffffffffffffff81118282101715612b8b57612b8b612b52565b6040516020810167ffffffffffffffff81118282101715612b8b57612b8b612b52565b6040805190810167ffffffffffffffff81118282101715612b8b57612b8b612b52565b604051601f8201601f1916810167ffffffffffffffff81118282101715612c2357612c23612b52565b604052919050565b600082601f830112612c3c57600080fd5b813567ffffffffffffffff811115612c5657612c56612b52565b612c69601f8201601f1916602001612bfa565b818152846020838601011115612c7e57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215612cad57600080fd5b813567ffffffffffffffff811115612cc457600080fd5b61093884828501612c2b565b600060208284031215612ce257600080fd5b5035919050565b6000612cf3612b91565b9050806080830184811115612d0757600080fd5b835b81811015612d21578035835260209283019201612d09565b50505092915050565b6000612d34612bd7565b9050806040830184811115612d0757600080fd5b600060a08284031215612d5a57600080fd5b612d62612b68565b9050813581526020820135602082015260408201356040820152606082013560608201526080820135608082015292915050565b6000818303610200811215612daa57600080fd5b612db2612b91565b91506080811215612dc257600080fd5b612dca612bb4565b84601f850112612dd957600080fd5b612de38585612ce9565b815282526040607f1982011215612df957600080fd5b50612e02612bb4565b83609f840112612e1157600080fd5b612e1e8460808501612d2a565b81526020820152612e328360c08401612d48565b6040820152612e45836101608401612d48565b606082015292915050565b600080828403610280811215612e6557600080fd5b6080811215612e7357600080fd5b50612e7c612b91565b83358152602084013560208201526040840135612e9881612a4f565b60408201526060840135612eab81612abd565b60608201529150612ebf8460808501612d96565b90509250929050565b60008083601f840112612eda57600080fd5b50813567ffffffffffffffff811115612ef257600080fd5b6020830191508360208260051b85010111156129f457600080fd5b60008060008060008060008060e0898b031215612f2957600080fd5b612f338a8a6129a1565b9750606089013567ffffffffffffffff80821115612f5057600080fd5b612f5c8c838d016129b2565b909950975060808b0135915080821115612f7557600080fd5b612f818c838d01612ec8565b909750955060a08b0135915080821115612f9a57600080fd5b50612fa78b828c01612ec8565b999c989b50969995989497949560c00135949350505050565b6000806102208385031215612fd457600080fd5b82359150612ebf8460208501612d96565b600067ffffffffffffffff821115612fff57612fff612b52565b5060051b60200190565b6000602080838503121561301c57600080fd5b823567ffffffffffffffff8082111561303457600080fd5b908401906040828703121561304857600080fd5b613050612bd7565b82358281111561305f57600080fd5b61306b88828601612c2b565b825250838301358281111561307f57600080fd5b80840193505086601f84011261309457600080fd5b823591506130a96130a483612fe5565b612bfa565b82815260059290921b830184019184810190888411156130c857600080fd5b938501935b838510156130e6578435825293850193908501906130cd565b948201949094529695505050505050565b600082601f83011261310857600080fd5b813560206131186130a483612fe5565b82815260059290921b8401810191818101908684111561313757600080fd5b8286015b8481101561315b57803561314e81612abd565b835291830191830161313b565b509695505050505050565b803567ffffffffffffffff8116811461069157600080fd5b60008060008060008060c0878903121561319757600080fd5b863567ffffffffffffffff808211156131af57600080fd5b6131bb8a838b016130f7565b975060208901359150808211156131d157600080fd5b6131dd8a838b016130f7565b96506131eb60408a01612a91565b9550606089013591508082111561320157600080fd5b61320d8a838b01612c2b565b945061321b60808a01613166565b935060a089013591508082111561323157600080fd5b5061323e89828a01612c2b565b9150509295509295509295565b6000806000610160848603121561326157600080fd5b833592506132728560208601612d48565b91506132818560c08601612d48565b90509250925092565b815160408201908260005b6002811015612d21578251825260209283019290910190600101613295565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610667576106676132e0565b60ff8181168382160190811115610667576106676132e0565b8082028115828204841417610667576106676132e0565b634e487b7160e01b600052601260045260246000fd5b600060ff83168061336257613362613339565b8060ff84160491505092915050565b600060ff821660ff8103613387576133876132e0565b60010192915050565b60ff8281168282160390811115610667576106676132e0565b6000600182016133bb576133bb6132e0565b5060010190565b600081518084526020808501945080840160005b838110156133f2578151875295820195908201906001016133d6565b509495945050505050565b82815260406020820152600061093860408301846133c2565b60608152600061342960608301866133c2565b841515602084015282810360408401526119cb8185612b13565b60008183825b6004811015613468578151835260209283019290910190600101613449565b50505060808201905092915050565b60008183825b600281101561349c57815183526020928301929091019060010161347d565b50505060408201905092915050565b600082516134bd818460208701612aef565b9190910192915050565b81810381811115610667576106676132e0565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b63ffffffff818116838216019080821115613523576135236132e0565b5092915050565b600081518084526020808501945080840160005b838110156133f25781516001600160a01b03168752958201959082019060010161353e565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526135938184018a61352a565b905082810360808401526135a7818961352a565b905060ff871660a084015282810360c08401526135c48187612b13565b905067ffffffffffffffff851660e08401528281036101008401526135e98185612b13565b9c9b505050505050505050505050565b6000613603612b91565b905080608083018481111561361757600080fd5b835b81811015612d21578051835260209283019201613619565b600061363b612bd7565b905080604083018481111561361757600080fd5b600060a0828403121561366157600080fd5b613669612b68565b9050815181526020820151602082015260408201516040820152606082015160608201526080820151608082015292915050565b6000808284036102808112156136b257600080fd5b60808112156136c057600080fd5b6136c8612b91565b845181526020850151602082015260408501516136e481612a4f565b604082015260608501516136f781612abd565b60608201529250607f19810161020081121561371257600080fd5b61371a612b91565b608082121561372857600080fd5b613730612bb4565b915086609f87011261374157600080fd5b61374e87608088016135f9565b8252818152604060ff198401121561376557600080fd5b61376d612bb4565b92508661011f87011261377f57600080fd5b61378d876101008801613631565b83528260208201526137a387610140880161364f565b60408201526137b6876101e0880161364f565b60608201528093505050509250929050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526138028285018b61352a565b91508382036080850152613816828a61352a565b915060ff881660a085015283820360c08501526138338288612b13565b90861660e085015283810361010085015290506135e98185612b13565b60008261385f5761385f613339565b50069056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a164736f6c6343000813000a",
}

var VRFABI = VRFMetaData.ABI

var VRFBin = VRFMetaData.Bin

func DeployVRF(auth *bind.TransactOpts, backend bind.ContractBackend, _keyProvider common.Address, _keyID [32]byte) (common.Address, *types.Transaction, *VRF, error) {
	parsed, err := VRFMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBin), backend, _keyProvider, _keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

type VRF struct {
	VRFCaller
	VRFTransactor
	VRFFilterer
}

type VRFCaller struct {
	contract *bind.BoundContract
}

type VRFTransactor struct {
	contract *bind.BoundContract
}

type VRFFilterer struct {
	contract *bind.BoundContract
}

type VRFSession struct {
	Contract     *VRF
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCallerSession struct {
	Contract *VRFCaller
	CallOpts bind.CallOpts
}

type VRFTransactorSession struct {
	Contract     *VRFTransactor
	TransactOpts bind.TransactOpts
}

type VRFRaw struct {
	Contract *VRF
}

type VRFCallerRaw struct {
	Contract *VRFCaller
}

type VRFTransactorRaw struct {
	Contract *VRFTransactor
}

func NewVRF(address common.Address, backend bind.ContractBackend) (*VRF, error) {
	contract, err := bindVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

func NewVRFCaller(address common.Address, caller bind.ContractCaller) (*VRFCaller, error) {
	contract, err := bindVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCaller{contract: contract}, nil
}

func NewVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFTransactor, error) {
	contract, err := bindVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFTransactor{contract: contract}, nil
}

func NewVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFFilterer, error) {
	contract, err := bindVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFFilterer{contract: contract}, nil
}

func bindVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRF *VRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.VRFCaller.contract.Call(opts, result, method, params...)
}

func (_VRF *VRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transfer(opts)
}

func (_VRF *VRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transact(opts, method, params...)
}

func (_VRF *VRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.contract.Call(opts, result, method, params...)
}

func (_VRF *VRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transfer(opts)
}

func (_VRF *VRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transact(opts, method, params...)
}

func (_VRF *VRFCaller) AddressToString(opts *bind.CallOpts, a common.Address) (string, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "addressToString", a)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRF *VRFSession) AddressToString(a common.Address) (string, error) {
	return _VRF.Contract.AddressToString(&_VRF.CallOpts, a)
}

func (_VRF *VRFCallerSession) AddressToString(a common.Address) (string, error) {
	return _VRF.Contract.AddressToString(&_VRF.CallOpts, a)
}

func (_VRF *VRFCaller) Bytes32ToString(opts *bind.CallOpts, s [32]byte) (string, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "bytes32ToString", s)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRF *VRFSession) Bytes32ToString(s [32]byte) (string, error) {
	return _VRF.Contract.Bytes32ToString(&_VRF.CallOpts, s)
}

func (_VRF *VRFCallerSession) Bytes32ToString(s [32]byte) (string, error) {
	return _VRF.Contract.Bytes32ToString(&_VRF.CallOpts, s)
}

func (_VRF *VRFCaller) BytesToString(opts *bind.CallOpts, _bytes []byte) (string, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "bytesToString", _bytes)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRF *VRFSession) BytesToString(_bytes []byte) (string, error) {
	return _VRF.Contract.BytesToString(&_VRF.CallOpts, _bytes)
}

func (_VRF *VRFCallerSession) BytesToString(_bytes []byte) (string, error) {
	return _VRF.Contract.BytesToString(&_VRF.CallOpts, _bytes)
}

func (_VRF *VRFCaller) HashToCurve(opts *bind.CallOpts, m [32]byte, f1 HashToCurveFProof, f2 HashToCurveFProof) (ECCArithmeticG1Point, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "hashToCurve", m, f1, f2)

	if err != nil {
		return *new(ECCArithmeticG1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(ECCArithmeticG1Point)).(*ECCArithmeticG1Point)

	return out0, err

}

func (_VRF *VRFSession) HashToCurve(m [32]byte, f1 HashToCurveFProof, f2 HashToCurveFProof) (ECCArithmeticG1Point, error) {
	return _VRF.Contract.HashToCurve(&_VRF.CallOpts, m, f1, f2)
}

func (_VRF *VRFCallerSession) HashToCurve(m [32]byte, f1 HashToCurveFProof, f2 HashToCurveFProof) (ECCArithmeticG1Point, error) {
	return _VRF.Contract.HashToCurve(&_VRF.CallOpts, m, f1, f2)
}

func (_VRF *VRFCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(struct {
		ConfigCount  uint32
		BlockNumber  uint32
		ConfigDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_VRF *VRFSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRF.Contract.LatestConfigDetails(&_VRF.CallOpts)
}

func (_VRF *VRFCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRF.Contract.LatestConfigDetails(&_VRF.CallOpts)
}

func (_VRF *VRFCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(struct {
		ScanLogs     bool
		ConfigDigest [32]byte
		Epoch        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRF *VRFSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRF.Contract.LatestConfigDigestAndEpoch(&_VRF.CallOpts)
}

func (_VRF *VRFCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRF.Contract.LatestConfigDigestAndEpoch(&_VRF.CallOpts)
}

func (_VRF *VRFCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRF *VRFSession) Owner() (common.Address, error) {
	return _VRF.Contract.Owner(&_VRF.CallOpts)
}

func (_VRF *VRFCallerSession) Owner() (common.Address, error) {
	return _VRF.Contract.Owner(&_VRF.CallOpts)
}

func (_VRF *VRFCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRF *VRFSession) SKeyID() ([32]byte, error) {
	return _VRF.Contract.SKeyID(&_VRF.CallOpts)
}

func (_VRF *VRFCallerSession) SKeyID() ([32]byte, error) {
	return _VRF.Contract.SKeyID(&_VRF.CallOpts)
}

func (_VRF *VRFCaller) SNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "s_nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRF *VRFSession) SNonce(arg0 common.Address) (*big.Int, error) {
	return _VRF.Contract.SNonce(&_VRF.CallOpts, arg0)
}

func (_VRF *VRFCallerSession) SNonce(arg0 common.Address) (*big.Int, error) {
	return _VRF.Contract.SNonce(&_VRF.CallOpts, arg0)
}

func (_VRF *VRFCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRF *VRFSession) SProvingKeyHash() ([32]byte, error) {
	return _VRF.Contract.SProvingKeyHash(&_VRF.CallOpts)
}

func (_VRF *VRFCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _VRF.Contract.SProvingKeyHash(&_VRF.CallOpts)
}

func (_VRF *VRFCaller) ToASCII(opts *bind.CallOpts, _uint8 uint8) (uint8, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "toASCII", _uint8)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRF *VRFSession) ToASCII(_uint8 uint8) (uint8, error) {
	return _VRF.Contract.ToASCII(&_VRF.CallOpts, _uint8)
}

func (_VRF *VRFCallerSession) ToASCII(_uint8 uint8) (uint8, error) {
	return _VRF.Contract.ToASCII(&_VRF.CallOpts, _uint8)
}

func (_VRF *VRFCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRF *VRFSession) TypeAndVersion() (string, error) {
	return _VRF.Contract.TypeAndVersion(&_VRF.CallOpts)
}

func (_VRF *VRFCallerSession) TypeAndVersion() (string, error) {
	return _VRF.Contract.TypeAndVersion(&_VRF.CallOpts)
}

func (_VRF *VRFCaller) VrfOutput(opts *bind.CallOpts, input [32]byte, p VRFProof) ([32]byte, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "vrfOutput", input, p)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRF *VRFSession) VrfOutput(input [32]byte, p VRFProof) ([32]byte, error) {
	return _VRF.Contract.VrfOutput(&_VRF.CallOpts, input, p)
}

func (_VRF *VRFCallerSession) VrfOutput(input [32]byte, p VRFProof) ([32]byte, error) {
	return _VRF.Contract.VrfOutput(&_VRF.CallOpts, input, p)
}

func (_VRF *VRFTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "acceptOwnership")
}

func (_VRF *VRFSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRF.Contract.AcceptOwnership(&_VRF.TransactOpts)
}

func (_VRF *VRFTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRF.Contract.AcceptOwnership(&_VRF.TransactOpts)
}

func (_VRF *VRFTransactor) FulfillRandomWords(opts *bind.TransactOpts, r VRFRequest, p VRFProof) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "fulfillRandomWords", r, p)
}

func (_VRF *VRFSession) FulfillRandomWords(r VRFRequest, p VRFProof) (*types.Transaction, error) {
	return _VRF.Contract.FulfillRandomWords(&_VRF.TransactOpts, r, p)
}

func (_VRF *VRFTransactorSession) FulfillRandomWords(r VRFRequest, p VRFProof) (*types.Transaction, error) {
	return _VRF.Contract.FulfillRandomWords(&_VRF.TransactOpts, r, p)
}

func (_VRF *VRFTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "keyGenerated", kd)
}

func (_VRF *VRFSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRF.Contract.KeyGenerated(&_VRF.TransactOpts, kd)
}

func (_VRF *VRFTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRF.Contract.KeyGenerated(&_VRF.TransactOpts, kd)
}

func (_VRF *VRFTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "newKeyRequested")
}

func (_VRF *VRFSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRF.Contract.NewKeyRequested(&_VRF.TransactOpts)
}

func (_VRF *VRFTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRF.Contract.NewKeyRequested(&_VRF.TransactOpts)
}

func (_VRF *VRFTransactor) RequestRandomWords(opts *bind.TransactOpts, seed *big.Int, numWords uint32) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "requestRandomWords", seed, numWords)
}

func (_VRF *VRFSession) RequestRandomWords(seed *big.Int, numWords uint32) (*types.Transaction, error) {
	return _VRF.Contract.RequestRandomWords(&_VRF.TransactOpts, seed, numWords)
}

func (_VRF *VRFTransactorSession) RequestRandomWords(seed *big.Int, numWords uint32) (*types.Transaction, error) {
	return _VRF.Contract.RequestRandomWords(&_VRF.TransactOpts, seed, numWords)
}

func (_VRF *VRFTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_VRF *VRFSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _VRF.Contract.SetConfig(&_VRF.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_VRF *VRFTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _VRF.Contract.SetConfig(&_VRF.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_VRF *VRFTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "transferOwnership", to)
}

func (_VRF *VRFSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRF.Contract.TransferOwnership(&_VRF.TransactOpts, to)
}

func (_VRF *VRFTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRF.Contract.TransferOwnership(&_VRF.TransactOpts, to)
}

func (_VRF *VRFTransactor) Transmit(opts *bind.TransactOpts, arg0 [3][32]byte, arg1 []byte, arg2 [][32]byte, arg3 [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "transmit", arg0, arg1, arg2, arg3, arg4)
}

func (_VRF *VRFSession) Transmit(arg0 [3][32]byte, arg1 []byte, arg2 [][32]byte, arg3 [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _VRF.Contract.Transmit(&_VRF.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

func (_VRF *VRFTransactorSession) Transmit(arg0 [3][32]byte, arg1 []byte, arg2 [][32]byte, arg3 [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _VRF.Contract.Transmit(&_VRF.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

func (_VRF *VRFTransactor) TransmitVRFResponse(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "transmitVRFResponse", reportContext, report)
}

func (_VRF *VRFSession) TransmitVRFResponse(reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _VRF.Contract.TransmitVRFResponse(&_VRF.TransactOpts, reportContext, report)
}

func (_VRF *VRFTransactorSession) TransmitVRFResponse(reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _VRF.Contract.TransmitVRFResponse(&_VRF.TransactOpts, reportContext, report)
}

type VRFConfigSetIterator struct {
	Event *VRFConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFConfigSet)
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
		it.Event = new(VRFConfigSet)
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

func (it *VRFConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFConfigSet struct {
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

func (_VRF *VRFFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFConfigSetIterator, error) {

	logs, sub, err := _VRF.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFConfigSetIterator{contract: _VRF.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRF *VRFFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRF.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFConfigSet)
				if err := _VRF.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRF *VRFFilterer) ParseConfigSet(log types.Log) (*VRFConfigSet, error) {
	event := new(VRFConfigSet)
	if err := _VRF.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFOwnershipTransferRequestedIterator struct {
	Event *VRFOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFOwnershipTransferRequested)
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
		it.Event = new(VRFOwnershipTransferRequested)
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

func (it *VRFOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRF *VRFFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFOwnershipTransferRequestedIterator{contract: _VRF.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRF *VRFFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFOwnershipTransferRequested)
				if err := _VRF.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRF *VRFFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFOwnershipTransferRequested, error) {
	event := new(VRFOwnershipTransferRequested)
	if err := _VRF.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFOwnershipTransferredIterator struct {
	Event *VRFOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFOwnershipTransferred)
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
		it.Event = new(VRFOwnershipTransferred)
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

func (it *VRFOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRF *VRFFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFOwnershipTransferredIterator{contract: _VRF.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRF *VRFFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFOwnershipTransferred)
				if err := _VRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRF *VRFFilterer) ParseOwnershipTransferred(log types.Log) (*VRFOwnershipTransferred, error) {
	event := new(VRFOwnershipTransferred)
	if err := _VRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRandomWordsFulfilledIterator struct {
	Event *VRFRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRandomWordsFulfilled)
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
		it.Event = new(VRFRandomWordsFulfilled)
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

func (it *VRFRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRandomWordsFulfilled struct {
	RequestID [32]byte
	Output    []*big.Int
	Success   bool
	ErrorData []byte
	Raw       types.Log
}

func (_VRF *VRFFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestID [][32]byte) (*VRFRandomWordsFulfilledIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFRandomWordsFulfilledIterator{contract: _VRF.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRF *VRFFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFRandomWordsFulfilled, requestID [][32]byte) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRandomWordsFulfilled)
				if err := _VRF.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRF *VRFFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFRandomWordsFulfilled, error) {
	event := new(VRFRandomWordsFulfilled)
	if err := _VRF.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFRandomWordsRequestedIterator struct {
	Event *VRFRandomWordsRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFRandomWordsRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRandomWordsRequested)
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
		it.Event = new(VRFRandomWordsRequested)
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

func (it *VRFRandomWordsRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFRandomWordsRequested struct {
	RequestID [32]byte
	Seed      *big.Int
	NumWords  uint32
	Sender    common.Address
	Raw       types.Log
}

func (_VRF *VRFFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts) (*VRFRandomWordsRequestedIterator, error) {

	logs, sub, err := _VRF.contract.FilterLogs(opts, "RandomWordsRequested")
	if err != nil {
		return nil, err
	}
	return &VRFRandomWordsRequestedIterator{contract: _VRF.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

func (_VRF *VRFFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFRandomWordsRequested) (event.Subscription, error) {

	logs, sub, err := _VRF.contract.WatchLogs(opts, "RandomWordsRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFRandomWordsRequested)
				if err := _VRF.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
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

func (_VRF *VRFFilterer) ParseRandomWordsRequested(log types.Log) (*VRFRandomWordsRequested, error) {
	event := new(VRFRandomWordsRequested)
	if err := _VRF.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFTransmittedIterator struct {
	Event *VRFTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFTransmitted)
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
		it.Event = new(VRFTransmitted)
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

func (it *VRFTransmittedIterator) Error() error {
	return it.fail
}

func (it *VRFTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_VRF *VRFFilterer) FilterTransmitted(opts *bind.FilterOpts) (*VRFTransmittedIterator, error) {

	logs, sub, err := _VRF.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &VRFTransmittedIterator{contract: _VRF.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_VRF *VRFFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *VRFTransmitted) (event.Subscription, error) {

	logs, sub, err := _VRF.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFTransmitted)
				if err := _VRF.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_VRF *VRFFilterer) ParseTransmitted(log types.Log) (*VRFTransmitted, error) {
	event := new(VRFTransmitted)
	if err := _VRF.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

package vrfcoordinator

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

type VRFBeaconTypesCallback struct {
	RequestID      *big.Int
	NumWords       uint16
	Requester      common.Address
	Arguments      []byte
	GasAllowance   *big.Int
	SubID          *big.Int
	GasPrice       *big.Int
	WeiPerUnitLink *big.Int
}

type VRFBeaconTypesCoordinatorConfig struct {
	UseReasonableGasPrice             bool
	ReentrancyLock                    bool
	Paused                            bool
	PremiumPercentage                 uint8
	UnusedGasPenaltyPercent           uint8
	StalenessSeconds                  uint32
	RedeemableRequestGasOverhead      uint32
	CallbackRequestGasOverhead        uint32
	ReasonableGasPriceStalenessBlocks uint32
	FallbackWeiPerUnitLink            *big.Int
}

type VRFBeaconTypesCostedCallback struct {
	Callback VRFBeaconTypesCallback
	Price    *big.Int
}

type VRFBeaconTypesOutputServed struct {
	Height            uint64
	ConfirmationDelay *big.Int
	ProofG1X          *big.Int
	ProofG1Y          *big.Int
}

type VRFBeaconTypesVRFOutput struct {
	BlockHeight       uint64
	ConfirmationDelay *big.Int
	VrfOutput         ECCArithmeticG1Point
	Callbacks         []VRFBeaconTypesCostedCallback
	ShouldStore       bool
}

type VRFCoordinatorCallbackConfig struct {
	MaxCallbackGasLimit        uint32
	MaxCallbackArgumentsLength uint32
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
	Bin: "0x608060405234801561001057600080fd5b5060405161042738038061042783398101604081905261002f9161016e565b8060006001600160a01b03821661008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100bd576100bd816100c5565b50505061019e565b336001600160a01b0382160361011d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561018057600080fd5b81516001600160a01b038116811461019757600080fd5b9392505050565b61027a806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461023d565b610131565b6001546001600160a01b031633146100da5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610139610145565b6101428161019a565b50565b6000546001600160a01b031633146101985760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b60448201526064016100d1565b565b336001600160a01b038216036101ec5760405162461bcd60e51b815260206004820152601760248201527621b0b73737ba103a3930b739b332b9103a379039b2b63360491b60448201526064016100d1565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561024f57600080fd5b81356001600160a01b038116811461026657600080fd5b939250505056fea164736f6c6343000813000a",
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
	Bin: "0x608060405234801561001057600080fd5b5060405161044238038061044283398101604081905261002f91610186565b6001600160a01b03821661008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100ba576100ba816100c1565b50506101b9565b336001600160a01b038216036101195760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610081565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b038116811461018157600080fd5b919050565b6000806040838503121561019957600080fd5b6101a28361016a565b91506101b06020840161016a565b90509250929050565b61027a806101c86000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461023d565b610131565b6001546001600160a01b031633146100da5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610139610145565b6101428161019a565b50565b6000546001600160a01b031633146101985760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b60448201526064016100d1565b565b336001600160a01b038216036101ec5760405162461bcd60e51b815260206004820152601760248201527621b0b73737ba103a3930b739b332b9103a379039b2b63360491b60448201526064016100d1565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561024f57600080fd5b81356001600160a01b038116811461026657600080fd5b939250505056fea164736f6c6343000813000a",
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

var IERC677ReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IERC677ReceiverABI = IERC677ReceiverMetaData.ABI

type IERC677Receiver struct {
	IERC677ReceiverCaller
	IERC677ReceiverTransactor
	IERC677ReceiverFilterer
}

type IERC677ReceiverCaller struct {
	contract *bind.BoundContract
}

type IERC677ReceiverTransactor struct {
	contract *bind.BoundContract
}

type IERC677ReceiverFilterer struct {
	contract *bind.BoundContract
}

type IERC677ReceiverSession struct {
	Contract     *IERC677Receiver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IERC677ReceiverCallerSession struct {
	Contract *IERC677ReceiverCaller
	CallOpts bind.CallOpts
}

type IERC677ReceiverTransactorSession struct {
	Contract     *IERC677ReceiverTransactor
	TransactOpts bind.TransactOpts
}

type IERC677ReceiverRaw struct {
	Contract *IERC677Receiver
}

type IERC677ReceiverCallerRaw struct {
	Contract *IERC677ReceiverCaller
}

type IERC677ReceiverTransactorRaw struct {
	Contract *IERC677ReceiverTransactor
}

func NewIERC677Receiver(address common.Address, backend bind.ContractBackend) (*IERC677Receiver, error) {
	contract, err := bindIERC677Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC677Receiver{IERC677ReceiverCaller: IERC677ReceiverCaller{contract: contract}, IERC677ReceiverTransactor: IERC677ReceiverTransactor{contract: contract}, IERC677ReceiverFilterer: IERC677ReceiverFilterer{contract: contract}}, nil
}

func NewIERC677ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC677ReceiverCaller, error) {
	contract, err := bindIERC677Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC677ReceiverCaller{contract: contract}, nil
}

func NewIERC677ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC677ReceiverTransactor, error) {
	contract, err := bindIERC677Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC677ReceiverTransactor{contract: contract}, nil
}

func NewIERC677ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC677ReceiverFilterer, error) {
	contract, err := bindIERC677Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC677ReceiverFilterer{contract: contract}, nil
}

func bindIERC677Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC677ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IERC677Receiver *IERC677ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC677Receiver.Contract.IERC677ReceiverCaller.contract.Call(opts, result, method, params...)
}

func (_IERC677Receiver *IERC677ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.IERC677ReceiverTransactor.contract.Transfer(opts)
}

func (_IERC677Receiver *IERC677ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.IERC677ReceiverTransactor.contract.Transact(opts, method, params...)
}

func (_IERC677Receiver *IERC677ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC677Receiver.Contract.contract.Call(opts, result, method, params...)
}

func (_IERC677Receiver *IERC677ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.contract.Transfer(opts)
}

func (_IERC677Receiver *IERC677ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.contract.Transact(opts, method, params...)
}

func (_IERC677Receiver *IERC677ReceiverTransactor) OnTokenTransfer(opts *bind.TransactOpts, sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677Receiver.contract.Transact(opts, "onTokenTransfer", sender, amount, data)
}

func (_IERC677Receiver *IERC677ReceiverSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.OnTokenTransfer(&_IERC677Receiver.TransactOpts, sender, amount, data)
}

func (_IERC677Receiver *IERC677ReceiverTransactorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC677Receiver.Contract.OnTokenTransfer(&_IERC677Receiver.TransactOpts, sender, amount, data)
}

var ISubscriptionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"pendingFulfillments\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var ISubscriptionABI = ISubscriptionMetaData.ABI

type ISubscription struct {
	ISubscriptionCaller
	ISubscriptionTransactor
	ISubscriptionFilterer
}

type ISubscriptionCaller struct {
	contract *bind.BoundContract
}

type ISubscriptionTransactor struct {
	contract *bind.BoundContract
}

type ISubscriptionFilterer struct {
	contract *bind.BoundContract
}

type ISubscriptionSession struct {
	Contract     *ISubscription
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ISubscriptionCallerSession struct {
	Contract *ISubscriptionCaller
	CallOpts bind.CallOpts
}

type ISubscriptionTransactorSession struct {
	Contract     *ISubscriptionTransactor
	TransactOpts bind.TransactOpts
}

type ISubscriptionRaw struct {
	Contract *ISubscription
}

type ISubscriptionCallerRaw struct {
	Contract *ISubscriptionCaller
}

type ISubscriptionTransactorRaw struct {
	Contract *ISubscriptionTransactor
}

func NewISubscription(address common.Address, backend bind.ContractBackend) (*ISubscription, error) {
	contract, err := bindISubscription(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISubscription{ISubscriptionCaller: ISubscriptionCaller{contract: contract}, ISubscriptionTransactor: ISubscriptionTransactor{contract: contract}, ISubscriptionFilterer: ISubscriptionFilterer{contract: contract}}, nil
}

func NewISubscriptionCaller(address common.Address, caller bind.ContractCaller) (*ISubscriptionCaller, error) {
	contract, err := bindISubscription(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionCaller{contract: contract}, nil
}

func NewISubscriptionTransactor(address common.Address, transactor bind.ContractTransactor) (*ISubscriptionTransactor, error) {
	contract, err := bindISubscription(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionTransactor{contract: contract}, nil
}

func NewISubscriptionFilterer(address common.Address, filterer bind.ContractFilterer) (*ISubscriptionFilterer, error) {
	contract, err := bindISubscription(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISubscriptionFilterer{contract: contract}, nil
}

func bindISubscription(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISubscriptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_ISubscription *ISubscriptionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISubscription.Contract.ISubscriptionCaller.contract.Call(opts, result, method, params...)
}

func (_ISubscription *ISubscriptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscription.Contract.ISubscriptionTransactor.contract.Transfer(opts)
}

func (_ISubscription *ISubscriptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscription.Contract.ISubscriptionTransactor.contract.Transact(opts, method, params...)
}

func (_ISubscription *ISubscriptionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISubscription.Contract.contract.Call(opts, result, method, params...)
}

func (_ISubscription *ISubscriptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscription.Contract.contract.Transfer(opts)
}

func (_ISubscription *ISubscriptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscription.Contract.contract.Transact(opts, method, params...)
}

func (_ISubscription *ISubscriptionCaller) GetSubscription(opts *bind.CallOpts, subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	var out []interface{}
	err := _ISubscription.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(struct {
		Balance             *big.Int
		PendingFulfillments uint64
		Owner               common.Address
		Consumers           []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingFulfillments = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_ISubscription *ISubscriptionSession) GetSubscription(subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	return _ISubscription.Contract.GetSubscription(&_ISubscription.CallOpts, subId)
}

func (_ISubscription *ISubscriptionCallerSession) GetSubscription(subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	return _ISubscription.Contract.GetSubscription(&_ISubscription.CallOpts, subId)
}

func (_ISubscription *ISubscriptionTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_ISubscription *ISubscriptionSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _ISubscription.Contract.AcceptSubscriptionOwnerTransfer(&_ISubscription.TransactOpts, subId)
}

func (_ISubscription *ISubscriptionTransactorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _ISubscription.Contract.AcceptSubscriptionOwnerTransfer(&_ISubscription.TransactOpts, subId)
}

func (_ISubscription *ISubscriptionTransactor) AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_ISubscription *ISubscriptionSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.AddConsumer(&_ISubscription.TransactOpts, subId, consumer)
}

func (_ISubscription *ISubscriptionTransactorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.AddConsumer(&_ISubscription.TransactOpts, subId, consumer)
}

func (_ISubscription *ISubscriptionTransactor) CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_ISubscription *ISubscriptionSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.CancelSubscription(&_ISubscription.TransactOpts, subId, to)
}

func (_ISubscription *ISubscriptionTransactorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.CancelSubscription(&_ISubscription.TransactOpts, subId, to)
}

func (_ISubscription *ISubscriptionTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "createSubscription")
}

func (_ISubscription *ISubscriptionSession) CreateSubscription() (*types.Transaction, error) {
	return _ISubscription.Contract.CreateSubscription(&_ISubscription.TransactOpts)
}

func (_ISubscription *ISubscriptionTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _ISubscription.Contract.CreateSubscription(&_ISubscription.TransactOpts)
}

func (_ISubscription *ISubscriptionTransactor) RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_ISubscription *ISubscriptionSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.RemoveConsumer(&_ISubscription.TransactOpts, subId, consumer)
}

func (_ISubscription *ISubscriptionTransactorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.RemoveConsumer(&_ISubscription.TransactOpts, subId, consumer)
}

func (_ISubscription *ISubscriptionTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _ISubscription.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_ISubscription *ISubscriptionSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.RequestSubscriptionOwnerTransfer(&_ISubscription.TransactOpts, subId, newOwner)
}

func (_ISubscription *ISubscriptionTransactorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _ISubscription.Contract.RequestSubscriptionOwnerTransfer(&_ISubscription.TransactOpts, subId, newOwner)
}

var IVRFCoordinatorConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MustBeCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeOwnerOrCoordinator\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"CoordinatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IVRFCoordinatorConsumerABI = IVRFCoordinatorConsumerMetaData.ABI

type IVRFCoordinatorConsumer struct {
	IVRFCoordinatorConsumerCaller
	IVRFCoordinatorConsumerTransactor
	IVRFCoordinatorConsumerFilterer
}

type IVRFCoordinatorConsumerCaller struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorConsumerTransactor struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorConsumerFilterer struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorConsumerSession struct {
	Contract     *IVRFCoordinatorConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFCoordinatorConsumerCallerSession struct {
	Contract *IVRFCoordinatorConsumerCaller
	CallOpts bind.CallOpts
}

type IVRFCoordinatorConsumerTransactorSession struct {
	Contract     *IVRFCoordinatorConsumerTransactor
	TransactOpts bind.TransactOpts
}

type IVRFCoordinatorConsumerRaw struct {
	Contract *IVRFCoordinatorConsumer
}

type IVRFCoordinatorConsumerCallerRaw struct {
	Contract *IVRFCoordinatorConsumerCaller
}

type IVRFCoordinatorConsumerTransactorRaw struct {
	Contract *IVRFCoordinatorConsumerTransactor
}

func NewIVRFCoordinatorConsumer(address common.Address, backend bind.ContractBackend) (*IVRFCoordinatorConsumer, error) {
	contract, err := bindIVRFCoordinatorConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorConsumer{IVRFCoordinatorConsumerCaller: IVRFCoordinatorConsumerCaller{contract: contract}, IVRFCoordinatorConsumerTransactor: IVRFCoordinatorConsumerTransactor{contract: contract}, IVRFCoordinatorConsumerFilterer: IVRFCoordinatorConsumerFilterer{contract: contract}}, nil
}

func NewIVRFCoordinatorConsumerCaller(address common.Address, caller bind.ContractCaller) (*IVRFCoordinatorConsumerCaller, error) {
	contract, err := bindIVRFCoordinatorConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorConsumerCaller{contract: contract}, nil
}

func NewIVRFCoordinatorConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFCoordinatorConsumerTransactor, error) {
	contract, err := bindIVRFCoordinatorConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorConsumerTransactor{contract: contract}, nil
}

func NewIVRFCoordinatorConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFCoordinatorConsumerFilterer, error) {
	contract, err := bindIVRFCoordinatorConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorConsumerFilterer{contract: contract}, nil
}

func bindIVRFCoordinatorConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFCoordinatorConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFCoordinatorConsumer.Contract.IVRFCoordinatorConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.IVRFCoordinatorConsumerTransactor.contract.Transfer(opts)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.IVRFCoordinatorConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFCoordinatorConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.contract.Transfer(opts)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVRFCoordinatorConsumer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerSession) Owner() (common.Address, error) {
	return _IVRFCoordinatorConsumer.Contract.Owner(&_IVRFCoordinatorConsumer.CallOpts)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerCallerSession) Owner() (common.Address, error) {
	return _IVRFCoordinatorConsumer.Contract.Owner(&_IVRFCoordinatorConsumer.CallOpts)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.contract.Transact(opts, "acceptOwnership")
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerSession) AcceptOwnership() (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.AcceptOwnership(&_IVRFCoordinatorConsumer.TransactOpts)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.AcceptOwnership(&_IVRFCoordinatorConsumer.TransactOpts)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords, arguments)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.RawFulfillRandomWords(&_IVRFCoordinatorConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.RawFulfillRandomWords(&_IVRFCoordinatorConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactor) SetCoordinator(opts *bind.TransactOpts, coordinator common.Address) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.contract.Transact(opts, "setCoordinator", coordinator)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerSession) SetCoordinator(coordinator common.Address) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.SetCoordinator(&_IVRFCoordinatorConsumer.TransactOpts, coordinator)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactorSession) SetCoordinator(coordinator common.Address) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.SetCoordinator(&_IVRFCoordinatorConsumer.TransactOpts, coordinator)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.contract.Transact(opts, "transferOwnership", to)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.TransferOwnership(&_IVRFCoordinatorConsumer.TransactOpts, to)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _IVRFCoordinatorConsumer.Contract.TransferOwnership(&_IVRFCoordinatorConsumer.TransactOpts, to)
}

type IVRFCoordinatorConsumerCoordinatorUpdatedIterator struct {
	Event *IVRFCoordinatorConsumerCoordinatorUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorConsumerCoordinatorUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorConsumerCoordinatorUpdated)
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
		it.Event = new(IVRFCoordinatorConsumerCoordinatorUpdated)
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

func (it *IVRFCoordinatorConsumerCoordinatorUpdatedIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorConsumerCoordinatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorConsumerCoordinatorUpdated struct {
	Coordinator common.Address
	Raw         types.Log
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) FilterCoordinatorUpdated(opts *bind.FilterOpts, coordinator []common.Address) (*IVRFCoordinatorConsumerCoordinatorUpdatedIterator, error) {

	var coordinatorRule []interface{}
	for _, coordinatorItem := range coordinator {
		coordinatorRule = append(coordinatorRule, coordinatorItem)
	}

	logs, sub, err := _IVRFCoordinatorConsumer.contract.FilterLogs(opts, "CoordinatorUpdated", coordinatorRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorConsumerCoordinatorUpdatedIterator{contract: _IVRFCoordinatorConsumer.contract, event: "CoordinatorUpdated", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) WatchCoordinatorUpdated(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorConsumerCoordinatorUpdated, coordinator []common.Address) (event.Subscription, error) {

	var coordinatorRule []interface{}
	for _, coordinatorItem := range coordinator {
		coordinatorRule = append(coordinatorRule, coordinatorItem)
	}

	logs, sub, err := _IVRFCoordinatorConsumer.contract.WatchLogs(opts, "CoordinatorUpdated", coordinatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorConsumerCoordinatorUpdated)
				if err := _IVRFCoordinatorConsumer.contract.UnpackLog(event, "CoordinatorUpdated", log); err != nil {
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

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) ParseCoordinatorUpdated(log types.Log) (*IVRFCoordinatorConsumerCoordinatorUpdated, error) {
	event := new(IVRFCoordinatorConsumerCoordinatorUpdated)
	if err := _IVRFCoordinatorConsumer.contract.UnpackLog(event, "CoordinatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFCoordinatorConsumerOwnershipTransferRequestedIterator struct {
	Event *IVRFCoordinatorConsumerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorConsumerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorConsumerOwnershipTransferRequested)
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
		it.Event = new(IVRFCoordinatorConsumerOwnershipTransferRequested)
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

func (it *IVRFCoordinatorConsumerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorConsumerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorConsumerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IVRFCoordinatorConsumerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IVRFCoordinatorConsumer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorConsumerOwnershipTransferRequestedIterator{contract: _IVRFCoordinatorConsumer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IVRFCoordinatorConsumer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorConsumerOwnershipTransferRequested)
				if err := _IVRFCoordinatorConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) ParseOwnershipTransferRequested(log types.Log) (*IVRFCoordinatorConsumerOwnershipTransferRequested, error) {
	event := new(IVRFCoordinatorConsumerOwnershipTransferRequested)
	if err := _IVRFCoordinatorConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFCoordinatorConsumerOwnershipTransferredIterator struct {
	Event *IVRFCoordinatorConsumerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorConsumerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorConsumerOwnershipTransferred)
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
		it.Event = new(IVRFCoordinatorConsumerOwnershipTransferred)
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

func (it *IVRFCoordinatorConsumerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorConsumerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorConsumerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IVRFCoordinatorConsumerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IVRFCoordinatorConsumer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorConsumerOwnershipTransferredIterator{contract: _IVRFCoordinatorConsumer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IVRFCoordinatorConsumer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorConsumerOwnershipTransferred)
				if err := _IVRFCoordinatorConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerFilterer) ParseOwnershipTransferred(log types.Log) (*IVRFCoordinatorConsumerOwnershipTransferred, error) {
	event := new(IVRFCoordinatorConsumerOwnershipTransferred)
	if err := _IVRFCoordinatorConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var IVRFCoordinatorProducerAPIMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint96[]\",\"name\":\"subBalances\",\"type\":\"uint96[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"subIDs\",\"type\":\"uint256[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"paymentsInJuels\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubscriptionLinkBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"shouldStore\",\"type\":\"bool\"}],\"internalType\":\"structVRFBeaconTypes.VRFOutput[]\",\"name\":\"vrfOutputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"processVRFOutputs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24[8]\",\"name\":\"confDelays\",\"type\":\"uint24[8]\"}],\"name\":\"setConfirmationDelays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"juelsAmount\",\"type\":\"uint256\"}],\"name\":\"transferLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IVRFCoordinatorProducerAPIABI = IVRFCoordinatorProducerAPIMetaData.ABI

type IVRFCoordinatorProducerAPI struct {
	IVRFCoordinatorProducerAPICaller
	IVRFCoordinatorProducerAPITransactor
	IVRFCoordinatorProducerAPIFilterer
}

type IVRFCoordinatorProducerAPICaller struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorProducerAPITransactor struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorProducerAPIFilterer struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorProducerAPISession struct {
	Contract     *IVRFCoordinatorProducerAPI
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFCoordinatorProducerAPICallerSession struct {
	Contract *IVRFCoordinatorProducerAPICaller
	CallOpts bind.CallOpts
}

type IVRFCoordinatorProducerAPITransactorSession struct {
	Contract     *IVRFCoordinatorProducerAPITransactor
	TransactOpts bind.TransactOpts
}

type IVRFCoordinatorProducerAPIRaw struct {
	Contract *IVRFCoordinatorProducerAPI
}

type IVRFCoordinatorProducerAPICallerRaw struct {
	Contract *IVRFCoordinatorProducerAPICaller
}

type IVRFCoordinatorProducerAPITransactorRaw struct {
	Contract *IVRFCoordinatorProducerAPITransactor
}

func NewIVRFCoordinatorProducerAPI(address common.Address, backend bind.ContractBackend) (*IVRFCoordinatorProducerAPI, error) {
	contract, err := bindIVRFCoordinatorProducerAPI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPI{IVRFCoordinatorProducerAPICaller: IVRFCoordinatorProducerAPICaller{contract: contract}, IVRFCoordinatorProducerAPITransactor: IVRFCoordinatorProducerAPITransactor{contract: contract}, IVRFCoordinatorProducerAPIFilterer: IVRFCoordinatorProducerAPIFilterer{contract: contract}}, nil
}

func NewIVRFCoordinatorProducerAPICaller(address common.Address, caller bind.ContractCaller) (*IVRFCoordinatorProducerAPICaller, error) {
	contract, err := bindIVRFCoordinatorProducerAPI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPICaller{contract: contract}, nil
}

func NewIVRFCoordinatorProducerAPITransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFCoordinatorProducerAPITransactor, error) {
	contract, err := bindIVRFCoordinatorProducerAPI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPITransactor{contract: contract}, nil
}

func NewIVRFCoordinatorProducerAPIFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFCoordinatorProducerAPIFilterer, error) {
	contract, err := bindIVRFCoordinatorProducerAPI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIFilterer{contract: contract}, nil
}

func bindIVRFCoordinatorProducerAPI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFCoordinatorProducerAPIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFCoordinatorProducerAPI.Contract.IVRFCoordinatorProducerAPICaller.contract.Call(opts, result, method, params...)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.IVRFCoordinatorProducerAPITransactor.contract.Transfer(opts)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.IVRFCoordinatorProducerAPITransactor.contract.Transact(opts, method, params...)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFCoordinatorProducerAPI.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.contract.Transfer(opts)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPICaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IVRFCoordinatorProducerAPI.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFCoordinatorProducerAPI.Contract.NUMCONFDELAYS(&_IVRFCoordinatorProducerAPI.CallOpts)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPICallerSession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFCoordinatorProducerAPI.Contract.NUMCONFDELAYS(&_IVRFCoordinatorProducerAPI.CallOpts)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPICaller) GetSubscriptionLinkBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVRFCoordinatorProducerAPI.contract.Call(opts, &out, "getSubscriptionLinkBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) GetSubscriptionLinkBalance() (*big.Int, error) {
	return _IVRFCoordinatorProducerAPI.Contract.GetSubscriptionLinkBalance(&_IVRFCoordinatorProducerAPI.CallOpts)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPICallerSession) GetSubscriptionLinkBalance() (*big.Int, error) {
	return _IVRFCoordinatorProducerAPI.Contract.GetSubscriptionLinkBalance(&_IVRFCoordinatorProducerAPI.CallOpts)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactor) BatchTransferLink(opts *bind.TransactOpts, recipients []common.Address, paymentsInJuels []*big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.contract.Transact(opts, "batchTransferLink", recipients, paymentsInJuels)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) BatchTransferLink(recipients []common.Address, paymentsInJuels []*big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.BatchTransferLink(&_IVRFCoordinatorProducerAPI.TransactOpts, recipients, paymentsInJuels)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorSession) BatchTransferLink(recipients []common.Address, paymentsInJuels []*big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.BatchTransferLink(&_IVRFCoordinatorProducerAPI.TransactOpts, recipients, paymentsInJuels)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactor) ProcessVRFOutputs(opts *bind.TransactOpts, vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.contract.Transact(opts, "processVRFOutputs", vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) ProcessVRFOutputs(vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.ProcessVRFOutputs(&_IVRFCoordinatorProducerAPI.TransactOpts, vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorSession) ProcessVRFOutputs(vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.ProcessVRFOutputs(&_IVRFCoordinatorProducerAPI.TransactOpts, vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactor) SetConfirmationDelays(opts *bind.TransactOpts, confDelays [8]*big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.contract.Transact(opts, "setConfirmationDelays", confDelays)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) SetConfirmationDelays(confDelays [8]*big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.SetConfirmationDelays(&_IVRFCoordinatorProducerAPI.TransactOpts, confDelays)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorSession) SetConfirmationDelays(confDelays [8]*big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.SetConfirmationDelays(&_IVRFCoordinatorProducerAPI.TransactOpts, confDelays)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactor) TransferLink(opts *bind.TransactOpts, recipient common.Address, juelsAmount *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.contract.Transact(opts, "transferLink", recipient, juelsAmount)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) TransferLink(recipient common.Address, juelsAmount *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.TransferLink(&_IVRFCoordinatorProducerAPI.TransactOpts, recipient, juelsAmount)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorSession) TransferLink(recipient common.Address, juelsAmount *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.TransferLink(&_IVRFCoordinatorProducerAPI.TransactOpts, recipient, juelsAmount)
}

type IVRFCoordinatorProducerAPIOutputsServedIterator struct {
	Event *IVRFCoordinatorProducerAPIOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorProducerAPIOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorProducerAPIOutputsServed)
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
		it.Event = new(IVRFCoordinatorProducerAPIOutputsServed)
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

func (it *IVRFCoordinatorProducerAPIOutputsServedIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorProducerAPIOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorProducerAPIOutputsServed struct {
	RecentBlockHeight  uint64
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*IVRFCoordinatorProducerAPIOutputsServedIterator, error) {

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIOutputsServedIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIOutputsServed) (event.Subscription, error) {

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorProducerAPIOutputsServed)
				if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) ParseOutputsServed(log types.Log) (*IVRFCoordinatorProducerAPIOutputsServed, error) {
	event := new(IVRFCoordinatorProducerAPIOutputsServed)
	if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFCoordinatorProducerAPIRandomWordsFulfilledIterator struct {
	Event *IVRFCoordinatorProducerAPIRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorProducerAPIRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorProducerAPIRandomWordsFulfilled)
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
		it.Event = new(IVRFCoordinatorProducerAPIRandomWordsFulfilled)
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

func (it *IVRFCoordinatorProducerAPIRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorProducerAPIRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorProducerAPIRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	SubBalances           []*big.Int
	SubIDs                []*big.Int
	Raw                   types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*IVRFCoordinatorProducerAPIRandomWordsFulfilledIterator, error) {

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIRandomWordsFulfilledIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorProducerAPIRandomWordsFulfilled)
				if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) ParseRandomWordsFulfilled(log types.Log) (*IVRFCoordinatorProducerAPIRandomWordsFulfilled, error) {
	event := new(IVRFCoordinatorProducerAPIRandomWordsFulfilled)
	if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator struct {
	Event *IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested)
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
		it.Event = new(IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested)
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

func (it *IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested struct {
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
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int) (*IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested)
				if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested, error) {
	event := new(IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested)
	if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFCoordinatorProducerAPIRandomnessRedeemedIterator struct {
	Event *IVRFCoordinatorProducerAPIRandomnessRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorProducerAPIRandomnessRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorProducerAPIRandomnessRedeemed)
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
		it.Event = new(IVRFCoordinatorProducerAPIRandomnessRedeemed)
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

func (it *IVRFCoordinatorProducerAPIRandomnessRedeemedIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorProducerAPIRandomnessRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorProducerAPIRandomnessRedeemed struct {
	RequestID *big.Int
	Requester common.Address
	SubID     *big.Int
	Raw       types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*IVRFCoordinatorProducerAPIRandomnessRedeemedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIRandomnessRedeemedIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "RandomnessRedeemed", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorProducerAPIRandomnessRedeemed)
				if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) ParseRandomnessRedeemed(log types.Log) (*IVRFCoordinatorProducerAPIRandomnessRedeemed, error) {
	event := new(IVRFCoordinatorProducerAPIRandomnessRedeemed)
	if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFCoordinatorProducerAPIRandomnessRequestedIterator struct {
	Event *IVRFCoordinatorProducerAPIRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorProducerAPIRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorProducerAPIRandomnessRequested)
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
		it.Event = new(IVRFCoordinatorProducerAPIRandomnessRequested)
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

func (it *IVRFCoordinatorProducerAPIRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorProducerAPIRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorProducerAPIRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int) (*IVRFCoordinatorProducerAPIRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIRandomnessRequestedIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIRandomnessRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorProducerAPIRandomnessRequested)
				if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) ParseRandomnessRequested(log types.Log) (*IVRFCoordinatorProducerAPIRandomnessRequested, error) {
	event := new(IVRFCoordinatorProducerAPIRandomnessRequested)
	if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var IVRFMigratableCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"getFulfillmentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

var IVRFMigratableCoordinatorABI = IVRFMigratableCoordinatorMetaData.ABI

type IVRFMigratableCoordinator struct {
	IVRFMigratableCoordinatorCaller
	IVRFMigratableCoordinatorTransactor
	IVRFMigratableCoordinatorFilterer
}

type IVRFMigratableCoordinatorCaller struct {
	contract *bind.BoundContract
}

type IVRFMigratableCoordinatorTransactor struct {
	contract *bind.BoundContract
}

type IVRFMigratableCoordinatorFilterer struct {
	contract *bind.BoundContract
}

type IVRFMigratableCoordinatorSession struct {
	Contract     *IVRFMigratableCoordinator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFMigratableCoordinatorCallerSession struct {
	Contract *IVRFMigratableCoordinatorCaller
	CallOpts bind.CallOpts
}

type IVRFMigratableCoordinatorTransactorSession struct {
	Contract     *IVRFMigratableCoordinatorTransactor
	TransactOpts bind.TransactOpts
}

type IVRFMigratableCoordinatorRaw struct {
	Contract *IVRFMigratableCoordinator
}

type IVRFMigratableCoordinatorCallerRaw struct {
	Contract *IVRFMigratableCoordinatorCaller
}

type IVRFMigratableCoordinatorTransactorRaw struct {
	Contract *IVRFMigratableCoordinatorTransactor
}

func NewIVRFMigratableCoordinator(address common.Address, backend bind.ContractBackend) (*IVRFMigratableCoordinator, error) {
	contract, err := bindIVRFMigratableCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFMigratableCoordinator{IVRFMigratableCoordinatorCaller: IVRFMigratableCoordinatorCaller{contract: contract}, IVRFMigratableCoordinatorTransactor: IVRFMigratableCoordinatorTransactor{contract: contract}, IVRFMigratableCoordinatorFilterer: IVRFMigratableCoordinatorFilterer{contract: contract}}, nil
}

func NewIVRFMigratableCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*IVRFMigratableCoordinatorCaller, error) {
	contract, err := bindIVRFMigratableCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFMigratableCoordinatorCaller{contract: contract}, nil
}

func NewIVRFMigratableCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFMigratableCoordinatorTransactor, error) {
	contract, err := bindIVRFMigratableCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFMigratableCoordinatorTransactor{contract: contract}, nil
}

func NewIVRFMigratableCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFMigratableCoordinatorFilterer, error) {
	contract, err := bindIVRFMigratableCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFMigratableCoordinatorFilterer{contract: contract}, nil
}

func bindIVRFMigratableCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFMigratableCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFMigratableCoordinator.Contract.IVRFMigratableCoordinatorCaller.contract.Call(opts, result, method, params...)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.IVRFMigratableCoordinatorTransactor.contract.Transfer(opts)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.IVRFMigratableCoordinatorTransactor.contract.Transact(opts, method, params...)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFMigratableCoordinator.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.contract.Transfer(opts)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorCaller) GetFee(opts *bind.CallOpts, subID *big.Int, extraArgs []byte) (*big.Int, error) {
	var out []interface{}
	err := _IVRFMigratableCoordinator.contract.Call(opts, &out, "getFee", subID, extraArgs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorSession) GetFee(subID *big.Int, extraArgs []byte) (*big.Int, error) {
	return _IVRFMigratableCoordinator.Contract.GetFee(&_IVRFMigratableCoordinator.CallOpts, subID, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorCallerSession) GetFee(subID *big.Int, extraArgs []byte) (*big.Int, error) {
	return _IVRFMigratableCoordinator.Contract.GetFee(&_IVRFMigratableCoordinator.CallOpts, subID, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorCaller) GetFulfillmentFee(opts *bind.CallOpts, subID *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*big.Int, error) {
	var out []interface{}
	err := _IVRFMigratableCoordinator.contract.Call(opts, &out, "getFulfillmentFee", subID, callbackGasLimit, arguments, extraArgs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorSession) GetFulfillmentFee(subID *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*big.Int, error) {
	return _IVRFMigratableCoordinator.Contract.GetFulfillmentFee(&_IVRFMigratableCoordinator.CallOpts, subID, callbackGasLimit, arguments, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorCallerSession) GetFulfillmentFee(subID *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*big.Int, error) {
	return _IVRFMigratableCoordinator.Contract.GetFulfillmentFee(&_IVRFMigratableCoordinator.CallOpts, subID, callbackGasLimit, arguments, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactor) RedeemRandomness(opts *bind.TransactOpts, subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.contract.Transact(opts, "redeemRandomness", subID, requestID, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorSession) RedeemRandomness(subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.RedeemRandomness(&_IVRFMigratableCoordinator.TransactOpts, subID, requestID, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactorSession) RedeemRandomness(subID *big.Int, requestID *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.RedeemRandomness(&_IVRFMigratableCoordinator.TransactOpts, subID, requestID, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactor) RequestRandomness(opts *bind.TransactOpts, subID *big.Int, numWords uint16, confirmationDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.contract.Transact(opts, "requestRandomness", subID, numWords, confirmationDelay, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorSession) RequestRandomness(subID *big.Int, numWords uint16, confirmationDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.RequestRandomness(&_IVRFMigratableCoordinator.TransactOpts, subID, numWords, confirmationDelay, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactorSession) RequestRandomness(subID *big.Int, numWords uint16, confirmationDelay *big.Int, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.RequestRandomness(&_IVRFMigratableCoordinator.TransactOpts, subID, numWords, confirmationDelay, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID *big.Int, numWords uint16, confirmationDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confirmationDelay, callbackGasLimit, arguments, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorSession) RequestRandomnessFulfillment(subID *big.Int, numWords uint16, confirmationDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.RequestRandomnessFulfillment(&_IVRFMigratableCoordinator.TransactOpts, subID, numWords, confirmationDelay, callbackGasLimit, arguments, extraArgs)
}

func (_IVRFMigratableCoordinator *IVRFMigratableCoordinatorTransactorSession) RequestRandomnessFulfillment(subID *big.Int, numWords uint16, confirmationDelay *big.Int, callbackGasLimit uint32, arguments []byte, extraArgs []byte) (*types.Transaction, error) {
	return _IVRFMigratableCoordinator.Contract.RequestRandomnessFulfillment(&_IVRFMigratableCoordinator.TransactOpts, subID, numWords, confirmationDelay, callbackGasLimit, arguments, extraArgs)
}

var IVRFMigrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVRFMigration\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedRequest\",\"type\":\"bytes\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedData\",\"type\":\"bytes\"}],\"name\":\"onMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IVRFMigrationABI = IVRFMigrationMetaData.ABI

type IVRFMigration struct {
	IVRFMigrationCaller
	IVRFMigrationTransactor
	IVRFMigrationFilterer
}

type IVRFMigrationCaller struct {
	contract *bind.BoundContract
}

type IVRFMigrationTransactor struct {
	contract *bind.BoundContract
}

type IVRFMigrationFilterer struct {
	contract *bind.BoundContract
}

type IVRFMigrationSession struct {
	Contract     *IVRFMigration
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFMigrationCallerSession struct {
	Contract *IVRFMigrationCaller
	CallOpts bind.CallOpts
}

type IVRFMigrationTransactorSession struct {
	Contract     *IVRFMigrationTransactor
	TransactOpts bind.TransactOpts
}

type IVRFMigrationRaw struct {
	Contract *IVRFMigration
}

type IVRFMigrationCallerRaw struct {
	Contract *IVRFMigrationCaller
}

type IVRFMigrationTransactorRaw struct {
	Contract *IVRFMigrationTransactor
}

func NewIVRFMigration(address common.Address, backend bind.ContractBackend) (*IVRFMigration, error) {
	contract, err := bindIVRFMigration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFMigration{IVRFMigrationCaller: IVRFMigrationCaller{contract: contract}, IVRFMigrationTransactor: IVRFMigrationTransactor{contract: contract}, IVRFMigrationFilterer: IVRFMigrationFilterer{contract: contract}}, nil
}

func NewIVRFMigrationCaller(address common.Address, caller bind.ContractCaller) (*IVRFMigrationCaller, error) {
	contract, err := bindIVRFMigration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFMigrationCaller{contract: contract}, nil
}

func NewIVRFMigrationTransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFMigrationTransactor, error) {
	contract, err := bindIVRFMigration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFMigrationTransactor{contract: contract}, nil
}

func NewIVRFMigrationFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFMigrationFilterer, error) {
	contract, err := bindIVRFMigration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFMigrationFilterer{contract: contract}, nil
}

func bindIVRFMigration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFMigrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFMigration *IVRFMigrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFMigration.Contract.IVRFMigrationCaller.contract.Call(opts, result, method, params...)
}

func (_IVRFMigration *IVRFMigrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFMigration.Contract.IVRFMigrationTransactor.contract.Transfer(opts)
}

func (_IVRFMigration *IVRFMigrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFMigration.Contract.IVRFMigrationTransactor.contract.Transact(opts, method, params...)
}

func (_IVRFMigration *IVRFMigrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFMigration.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFMigration *IVRFMigrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFMigration.Contract.contract.Transfer(opts)
}

func (_IVRFMigration *IVRFMigrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFMigration.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFMigration *IVRFMigrationCaller) MigrationVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IVRFMigration.contract.Call(opts, &out, "migrationVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_IVRFMigration *IVRFMigrationSession) MigrationVersion() (uint8, error) {
	return _IVRFMigration.Contract.MigrationVersion(&_IVRFMigration.CallOpts)
}

func (_IVRFMigration *IVRFMigrationCallerSession) MigrationVersion() (uint8, error) {
	return _IVRFMigration.Contract.MigrationVersion(&_IVRFMigration.CallOpts)
}

func (_IVRFMigration *IVRFMigrationTransactor) Migrate(opts *bind.TransactOpts, newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _IVRFMigration.contract.Transact(opts, "migrate", newCoordinator, encodedRequest)
}

func (_IVRFMigration *IVRFMigrationSession) Migrate(newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _IVRFMigration.Contract.Migrate(&_IVRFMigration.TransactOpts, newCoordinator, encodedRequest)
}

func (_IVRFMigration *IVRFMigrationTransactorSession) Migrate(newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _IVRFMigration.Contract.Migrate(&_IVRFMigration.TransactOpts, newCoordinator, encodedRequest)
}

func (_IVRFMigration *IVRFMigrationTransactor) OnMigration(opts *bind.TransactOpts, encodedData []byte) (*types.Transaction, error) {
	return _IVRFMigration.contract.Transact(opts, "onMigration", encodedData)
}

func (_IVRFMigration *IVRFMigrationSession) OnMigration(encodedData []byte) (*types.Transaction, error) {
	return _IVRFMigration.Contract.OnMigration(&_IVRFMigration.TransactOpts, encodedData)
}

func (_IVRFMigration *IVRFMigrationTransactorSession) OnMigration(encodedData []byte) (*types.Transaction, error) {
	return _IVRFMigration.Contract.OnMigration(&_IVRFMigration.TransactOpts, encodedData)
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
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61027a806101576000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461023d565b610131565b6001546001600160a01b031633146100da5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610139610145565b6101428161019a565b50565b6000546001600160a01b031633146101985760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b60448201526064016100d1565b565b336001600160a01b038216036101ec5760405162461bcd60e51b815260206004820152601760248201527621b0b73737ba103a3930b739b332b9103a379039b2b63360491b60448201526064016100d1565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561024f57600080fd5b81356001600160a01b038116811461026657600080fd5b939250505056fea164736f6c6343000813000a",
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

var VRFBeaconBillingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ContractPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"expectedLength\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCoordinatorConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJuelsConversion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestedSubID\",\"type\":\"uint256\"}],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"useReasonableGasPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"unusedGasPenaltyPercent\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redeemableRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callbackRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceStalenessBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.CoordinatorConfig\",\"name\":\"coordinatorConfig\",\"type\":\"tuple\"}],\"name\":\"CoordinatorConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PauseFlagChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"pendingFulfillments\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_coordinatorConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"useReasonableGasPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"unusedGasPenaltyPercent\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redeemableRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callbackRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceStalenessBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"useReasonableGasPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"unusedGasPenaltyPercent\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redeemableRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callbackRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceStalenessBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CoordinatorConfig\",\"name\":\"coordinatorConfig\",\"type\":\"tuple\"}],\"name\":\"setCoordinatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var VRFBeaconBillingABI = VRFBeaconBillingMetaData.ABI

type VRFBeaconBilling struct {
	VRFBeaconBillingCaller
	VRFBeaconBillingTransactor
	VRFBeaconBillingFilterer
}

type VRFBeaconBillingCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconBillingTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconBillingFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconBillingSession struct {
	Contract     *VRFBeaconBilling
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconBillingCallerSession struct {
	Contract *VRFBeaconBillingCaller
	CallOpts bind.CallOpts
}

type VRFBeaconBillingTransactorSession struct {
	Contract     *VRFBeaconBillingTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconBillingRaw struct {
	Contract *VRFBeaconBilling
}

type VRFBeaconBillingCallerRaw struct {
	Contract *VRFBeaconBillingCaller
}

type VRFBeaconBillingTransactorRaw struct {
	Contract *VRFBeaconBillingTransactor
}

func NewVRFBeaconBilling(address common.Address, backend bind.ContractBackend) (*VRFBeaconBilling, error) {
	contract, err := bindVRFBeaconBilling(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBilling{VRFBeaconBillingCaller: VRFBeaconBillingCaller{contract: contract}, VRFBeaconBillingTransactor: VRFBeaconBillingTransactor{contract: contract}, VRFBeaconBillingFilterer: VRFBeaconBillingFilterer{contract: contract}}, nil
}

func NewVRFBeaconBillingCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconBillingCaller, error) {
	contract, err := bindVRFBeaconBilling(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingCaller{contract: contract}, nil
}

func NewVRFBeaconBillingTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconBillingTransactor, error) {
	contract, err := bindVRFBeaconBilling(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingTransactor{contract: contract}, nil
}

func NewVRFBeaconBillingFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconBillingFilterer, error) {
	contract, err := bindVRFBeaconBilling(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingFilterer{contract: contract}, nil
}

func bindVRFBeaconBilling(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconBillingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeaconBilling *VRFBeaconBillingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconBilling.Contract.VRFBeaconBillingCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconBilling *VRFBeaconBillingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.VRFBeaconBillingTransactor.contract.Transfer(opts)
}

func (_VRFBeaconBilling *VRFBeaconBillingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.VRFBeaconBillingTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeaconBilling *VRFBeaconBillingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconBilling.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.contract.Transfer(opts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeaconBilling *VRFBeaconBillingCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFBeaconBilling.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFBeaconBilling *VRFBeaconBillingSession) MAXCONSUMERS() (uint16, error) {
	return _VRFBeaconBilling.Contract.MAXCONSUMERS(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFBeaconBilling.Contract.MAXCONSUMERS(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCaller) GetSubscription(opts *bind.CallOpts, subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	var out []interface{}
	err := _VRFBeaconBilling.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(struct {
		Balance             *big.Int
		PendingFulfillments uint64
		Owner               common.Address
		Consumers           []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingFulfillments = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VRFBeaconBilling *VRFBeaconBillingSession) GetSubscription(subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	return _VRFBeaconBilling.Contract.GetSubscription(&_VRFBeaconBilling.CallOpts, subId)
}

func (_VRFBeaconBilling *VRFBeaconBillingCallerSession) GetSubscription(subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	return _VRFBeaconBilling.Contract.GetSubscription(&_VRFBeaconBilling.CallOpts, subId)
}

func (_VRFBeaconBilling *VRFBeaconBillingCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconBilling.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconBilling *VRFBeaconBillingSession) ILink() (common.Address, error) {
	return _VRFBeaconBilling.Contract.ILink(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCallerSession) ILink() (common.Address, error) {
	return _VRFBeaconBilling.Contract.ILink(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconBilling.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconBilling *VRFBeaconBillingSession) Owner() (common.Address, error) {
	return _VRFBeaconBilling.Contract.Owner(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCallerSession) Owner() (common.Address, error) {
	return _VRFBeaconBilling.Contract.Owner(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCaller) SCoordinatorConfig(opts *bind.CallOpts) (struct {
	UseReasonableGasPrice             bool
	ReentrancyLock                    bool
	Paused                            bool
	PremiumPercentage                 uint8
	UnusedGasPenaltyPercent           uint8
	StalenessSeconds                  uint32
	RedeemableRequestGasOverhead      uint32
	CallbackRequestGasOverhead        uint32
	ReasonableGasPriceStalenessBlocks uint32
	FallbackWeiPerUnitLink            *big.Int
}, error) {
	var out []interface{}
	err := _VRFBeaconBilling.contract.Call(opts, &out, "s_coordinatorConfig")

	outstruct := new(struct {
		UseReasonableGasPrice             bool
		ReentrancyLock                    bool
		Paused                            bool
		PremiumPercentage                 uint8
		UnusedGasPenaltyPercent           uint8
		StalenessSeconds                  uint32
		RedeemableRequestGasOverhead      uint32
		CallbackRequestGasOverhead        uint32
		ReasonableGasPriceStalenessBlocks uint32
		FallbackWeiPerUnitLink            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UseReasonableGasPrice = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ReentrancyLock = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.PremiumPercentage = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.UnusedGasPenaltyPercent = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.StalenessSeconds = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.RedeemableRequestGasOverhead = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.CallbackRequestGasOverhead = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.ReasonableGasPriceStalenessBlocks = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.FallbackWeiPerUnitLink = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFBeaconBilling *VRFBeaconBillingSession) SCoordinatorConfig() (struct {
	UseReasonableGasPrice             bool
	ReentrancyLock                    bool
	Paused                            bool
	PremiumPercentage                 uint8
	UnusedGasPenaltyPercent           uint8
	StalenessSeconds                  uint32
	RedeemableRequestGasOverhead      uint32
	CallbackRequestGasOverhead        uint32
	ReasonableGasPriceStalenessBlocks uint32
	FallbackWeiPerUnitLink            *big.Int
}, error) {
	return _VRFBeaconBilling.Contract.SCoordinatorConfig(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCallerSession) SCoordinatorConfig() (struct {
	UseReasonableGasPrice             bool
	ReentrancyLock                    bool
	Paused                            bool
	PremiumPercentage                 uint8
	UnusedGasPenaltyPercent           uint8
	StalenessSeconds                  uint32
	RedeemableRequestGasOverhead      uint32
	CallbackRequestGasOverhead        uint32
	ReasonableGasPriceStalenessBlocks uint32
	FallbackWeiPerUnitLink            *big.Int
}, error) {
	return _VRFBeaconBilling.Contract.SCoordinatorConfig(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AcceptOwnership(&_VRFBeaconBilling.TransactOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AcceptOwnership(&_VRFBeaconBilling.TransactOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AcceptSubscriptionOwnerTransfer(&_VRFBeaconBilling.TransactOpts, subId)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AcceptSubscriptionOwnerTransfer(&_VRFBeaconBilling.TransactOpts, subId)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AddConsumer(&_VRFBeaconBilling.TransactOpts, subId, consumer)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AddConsumer(&_VRFBeaconBilling.TransactOpts, subId, consumer)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.CancelSubscription(&_VRFBeaconBilling.TransactOpts, subId, to)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.CancelSubscription(&_VRFBeaconBilling.TransactOpts, subId, to)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "createSubscription")
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.CreateSubscription(&_VRFBeaconBilling.TransactOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.CreateSubscription(&_VRFBeaconBilling.TransactOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.OnTokenTransfer(&_VRFBeaconBilling.TransactOpts, arg0, amount, data)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.OnTokenTransfer(&_VRFBeaconBilling.TransactOpts, arg0, amount, data)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.RemoveConsumer(&_VRFBeaconBilling.TransactOpts, subId, consumer)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.RemoveConsumer(&_VRFBeaconBilling.TransactOpts, subId, consumer)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.RequestSubscriptionOwnerTransfer(&_VRFBeaconBilling.TransactOpts, subId, newOwner)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.RequestSubscriptionOwnerTransfer(&_VRFBeaconBilling.TransactOpts, subId, newOwner)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) SetCoordinatorConfig(opts *bind.TransactOpts, coordinatorConfig VRFBeaconTypesCoordinatorConfig) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "setCoordinatorConfig", coordinatorConfig)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) SetCoordinatorConfig(coordinatorConfig VRFBeaconTypesCoordinatorConfig) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.SetCoordinatorConfig(&_VRFBeaconBilling.TransactOpts, coordinatorConfig)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) SetCoordinatorConfig(coordinatorConfig VRFBeaconTypesCoordinatorConfig) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.SetCoordinatorConfig(&_VRFBeaconBilling.TransactOpts, coordinatorConfig)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) SetPauseFlag(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "setPauseFlag", pause)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) SetPauseFlag(pause bool) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.SetPauseFlag(&_VRFBeaconBilling.TransactOpts, pause)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) SetPauseFlag(pause bool) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.SetPauseFlag(&_VRFBeaconBilling.TransactOpts, pause)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.TransferOwnership(&_VRFBeaconBilling.TransactOpts, to)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.TransferOwnership(&_VRFBeaconBilling.TransactOpts, to)
}

type VRFBeaconBillingCoordinatorConfigSetIterator struct {
	Event *VRFBeaconBillingCoordinatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingCoordinatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingCoordinatorConfigSet)
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
		it.Event = new(VRFBeaconBillingCoordinatorConfigSet)
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

func (it *VRFBeaconBillingCoordinatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingCoordinatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingCoordinatorConfigSet struct {
	CoordinatorConfig VRFBeaconTypesCoordinatorConfig
	Raw               types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterCoordinatorConfigSet(opts *bind.FilterOpts) (*VRFBeaconBillingCoordinatorConfigSetIterator, error) {

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "CoordinatorConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingCoordinatorConfigSetIterator{contract: _VRFBeaconBilling.contract, event: "CoordinatorConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchCoordinatorConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingCoordinatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "CoordinatorConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingCoordinatorConfigSet)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "CoordinatorConfigSet", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseCoordinatorConfigSet(log types.Log) (*VRFBeaconBillingCoordinatorConfigSet, error) {
	event := new(VRFBeaconBillingCoordinatorConfigSet)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "CoordinatorConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingOwnershipTransferRequestedIterator struct {
	Event *VRFBeaconBillingOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingOwnershipTransferRequested)
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
		it.Event = new(VRFBeaconBillingOwnershipTransferRequested)
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

func (it *VRFBeaconBillingOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconBillingOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingOwnershipTransferRequestedIterator{contract: _VRFBeaconBilling.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingOwnershipTransferRequested)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconBillingOwnershipTransferRequested, error) {
	event := new(VRFBeaconBillingOwnershipTransferRequested)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingOwnershipTransferredIterator struct {
	Event *VRFBeaconBillingOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingOwnershipTransferred)
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
		it.Event = new(VRFBeaconBillingOwnershipTransferred)
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

func (it *VRFBeaconBillingOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconBillingOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingOwnershipTransferredIterator{contract: _VRFBeaconBilling.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingOwnershipTransferred)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseOwnershipTransferred(log types.Log) (*VRFBeaconBillingOwnershipTransferred, error) {
	event := new(VRFBeaconBillingOwnershipTransferred)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingPauseFlagChangedIterator struct {
	Event *VRFBeaconBillingPauseFlagChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingPauseFlagChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingPauseFlagChanged)
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
		it.Event = new(VRFBeaconBillingPauseFlagChanged)
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

func (it *VRFBeaconBillingPauseFlagChangedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingPauseFlagChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingPauseFlagChanged struct {
	Paused bool
	Raw    types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterPauseFlagChanged(opts *bind.FilterOpts) (*VRFBeaconBillingPauseFlagChangedIterator, error) {

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "PauseFlagChanged")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingPauseFlagChangedIterator{contract: _VRFBeaconBilling.contract, event: "PauseFlagChanged", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchPauseFlagChanged(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingPauseFlagChanged) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "PauseFlagChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingPauseFlagChanged)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "PauseFlagChanged", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParsePauseFlagChanged(log types.Log) (*VRFBeaconBillingPauseFlagChanged, error) {
	event := new(VRFBeaconBillingPauseFlagChanged)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "PauseFlagChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSubscriptionCanceledIterator struct {
	Event *VRFBeaconBillingSubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSubscriptionCanceled)
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
		it.Event = new(VRFBeaconBillingSubscriptionCanceled)
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

func (it *VRFBeaconBillingSubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSubscriptionCanceled struct {
	SubId  *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFBeaconBillingSubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSubscriptionCanceledIterator{contract: _VRFBeaconBilling.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSubscriptionCanceled, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSubscriptionCanceled)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseSubscriptionCanceled(log types.Log) (*VRFBeaconBillingSubscriptionCanceled, error) {
	event := new(VRFBeaconBillingSubscriptionCanceled)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSubscriptionConsumerAddedIterator struct {
	Event *VRFBeaconBillingSubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSubscriptionConsumerAdded)
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
		it.Event = new(VRFBeaconBillingSubscriptionConsumerAdded)
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

func (it *VRFBeaconBillingSubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSubscriptionConsumerAdded struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFBeaconBillingSubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSubscriptionConsumerAddedIterator{contract: _VRFBeaconBilling.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSubscriptionConsumerAdded)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFBeaconBillingSubscriptionConsumerAdded, error) {
	event := new(VRFBeaconBillingSubscriptionConsumerAdded)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSubscriptionConsumerRemovedIterator struct {
	Event *VRFBeaconBillingSubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSubscriptionConsumerRemoved)
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
		it.Event = new(VRFBeaconBillingSubscriptionConsumerRemoved)
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

func (it *VRFBeaconBillingSubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSubscriptionConsumerRemoved struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFBeaconBillingSubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSubscriptionConsumerRemovedIterator{contract: _VRFBeaconBilling.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSubscriptionConsumerRemoved)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFBeaconBillingSubscriptionConsumerRemoved, error) {
	event := new(VRFBeaconBillingSubscriptionConsumerRemoved)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSubscriptionCreatedIterator struct {
	Event *VRFBeaconBillingSubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSubscriptionCreated)
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
		it.Event = new(VRFBeaconBillingSubscriptionCreated)
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

func (it *VRFBeaconBillingSubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSubscriptionCreated struct {
	SubId *big.Int
	Owner common.Address
	Raw   types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int, owner []common.Address) (*VRFBeaconBillingSubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSubscriptionCreatedIterator{contract: _VRFBeaconBilling.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSubscriptionCreated, subId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSubscriptionCreated)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseSubscriptionCreated(log types.Log) (*VRFBeaconBillingSubscriptionCreated, error) {
	event := new(VRFBeaconBillingSubscriptionCreated)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSubscriptionFundedIterator struct {
	Event *VRFBeaconBillingSubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSubscriptionFunded)
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
		it.Event = new(VRFBeaconBillingSubscriptionFunded)
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

func (it *VRFBeaconBillingSubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSubscriptionFunded struct {
	SubId      *big.Int
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFBeaconBillingSubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSubscriptionFundedIterator{contract: _VRFBeaconBilling.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSubscriptionFunded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSubscriptionFunded)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseSubscriptionFunded(log types.Log) (*VRFBeaconBillingSubscriptionFunded, error) {
	event := new(VRFBeaconBillingSubscriptionFunded)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFBeaconBillingSubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSubscriptionOwnerTransferRequested)
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
		it.Event = new(VRFBeaconBillingSubscriptionOwnerTransferRequested)
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

func (it *VRFBeaconBillingSubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSubscriptionOwnerTransferRequested struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFBeaconBillingSubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSubscriptionOwnerTransferRequestedIterator{contract: _VRFBeaconBilling.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSubscriptionOwnerTransferRequested)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFBeaconBillingSubscriptionOwnerTransferRequested, error) {
	event := new(VRFBeaconBillingSubscriptionOwnerTransferRequested)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSubscriptionOwnerTransferredIterator struct {
	Event *VRFBeaconBillingSubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSubscriptionOwnerTransferred)
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
		it.Event = new(VRFBeaconBillingSubscriptionOwnerTransferred)
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

func (it *VRFBeaconBillingSubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSubscriptionOwnerTransferred struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFBeaconBillingSubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSubscriptionOwnerTransferredIterator{contract: _VRFBeaconBilling.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFBeaconBilling.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSubscriptionOwnerTransferred)
				if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
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

func (_VRFBeaconBilling *VRFBeaconBillingFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFBeaconBillingSubscriptionOwnerTransferred, error) {
	event := new(VRFBeaconBillingSubscriptionOwnerTransferred)
	if err := _VRFBeaconBilling.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFBeaconTypesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint96[]\",\"name\":\"subBalances\",\"type\":\"uint96[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"subIDs\",\"type\":\"uint256[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50605780601d6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80632f7527cc14602d575b600080fd5b6034600881565b60405160ff909116815260200160405180910390f3fea164736f6c6343000813000a",
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
	SubBalances           []*big.Int
	SubIDs                []*big.Int
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
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFBeaconTypesRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconTypes.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesRandomnessFulfillmentRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
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

type VRFBeaconTypesRandomnessRedeemedIterator struct {
	Event *VRFBeaconTypesRandomnessRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconTypesRandomnessRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconTypesRandomnessRedeemed)
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
		it.Event = new(VRFBeaconTypesRandomnessRedeemed)
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

func (it *VRFBeaconTypesRandomnessRedeemedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconTypesRandomnessRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconTypesRandomnessRedeemed struct {
	RequestID *big.Int
	Requester common.Address
	SubID     *big.Int
	Raw       types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconTypesRandomnessRedeemedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesRandomnessRedeemedIterator{contract: _VRFBeaconTypes.contract, event: "RandomnessRedeemed", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconTypesRandomnessRedeemed)
				if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) ParseRandomnessRedeemed(log types.Log) (*VRFBeaconTypesRandomnessRedeemed, error) {
	event := new(VRFBeaconTypesRandomnessRedeemed)
	if err := _VRFBeaconTypes.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFBeaconTypesRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTypesRandomnessRequestedIterator{contract: _VRFBeaconTypes.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconTypes *VRFBeaconTypesFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconTypesRandomnessRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFBeaconTypes.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule)
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

var VRFCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"}],\"name\":\"GasAllowanceExceedsGasLeft\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"expectedLength\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCoordinatorConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidJuelsConversion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numRecipients\",\"type\":\"uint256\"}],\"name\":\"InvalidNumberOfRecipients\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestedSubID\",\"type\":\"uint256\"}],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"requestedVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"coordinatorVersion\",\"type\":\"uint8\"}],\"name\":\"MigrationVersionMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeProducer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativePaymentGiven\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnMigrationNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"ProducerAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confDelay\",\"type\":\"uint256\"}],\"name\":\"RandomnessSeedNotFoundForCallbacks\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numRecipients\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numPayments\",\"type\":\"uint256\"}],\"name\":\"RecipientsPaymentsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxCallbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxCallbackArgumentsLength\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structVRFCoordinator.CallbackConfig\",\"name\":\"newConfig\",\"type\":\"tuple\"}],\"name\":\"CallbackConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"useReasonableGasPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"unusedGasPenaltyPercent\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redeemableRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callbackRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceStalenessBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.CoordinatorConfig\",\"name\":\"coordinatorConfig\",\"type\":\"tuple\"}],\"name\":\"CoordinatorConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"newVersion\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"MigrationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PauseFlagChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint96[]\",\"name\":\"subBalances\",\"type\":\"uint96[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"subIDs\",\"type\":\"uint256[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubBalance\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"paymentsInJuels\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"deregisterMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getCallbackMemo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfirmationDelays\",\"outputs\":[{\"internalType\":\"uint24[8]\",\"name\":\"\",\"type\":\"uint24[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"getFulfillmentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"pendingFulfillments\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubscriptionLinkBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVRFMigration\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"encodedRequest\",\"type\":\"bytes\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onMigration\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"shouldStore\",\"type\":\"bool\"}],\"internalType\":\"structVRFBeaconTypes.VRFOutput[]\",\"name\":\"vrfOutputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"processVRFOutputs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"containsNewOutputs\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"registerMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_callbackConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maxCallbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxCallbackArgumentsLength\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_coordinatorConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"useReasonableGasPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"unusedGasPenaltyPercent\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redeemableRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callbackRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceStalenessBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_pendingRequests\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"slotNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_producer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxCallbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxCallbackArgumentsLength\",\"type\":\"uint32\"}],\"internalType\":\"structVRFCoordinator.CallbackConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setCallbackConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24[8]\",\"name\":\"confDelays\",\"type\":\"uint24[8]\"}],\"name\":\"setConfirmationDelays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"useReasonableGasPrice\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"premiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"unusedGasPenaltyPercent\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"redeemableRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callbackRequestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceStalenessBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CoordinatorConfig\",\"name\":\"coordinatorConfig\",\"type\":\"tuple\"}],\"name\":\"setCoordinatorConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPauseFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"producer\",\"type\":\"address\"}],\"name\":\"setProducer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"juelsAmount\",\"type\":\"uint256\"}],\"name\":\"transferLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620062f5380380620062f5833981016040819052620000349162000239565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf816200018e565b5050506001600160a01b03166080908152604080519182018152600080835260208301819052908201819052662386f26fc10000606090920191909152642386f26fc160b01b6006556004805463ffffffff60281b191668ffffffff00000000001790558290036200014457604051632abc297960e01b815260040160405180910390fd5b60a0829052600e805465ffffffffffff16906000620001638362000278565b91906101000a81548165ffffffffffff021916908365ffffffffffff160217905550505050620002ac565b336001600160a01b03821603620001e85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600080604083850312156200024d57600080fd5b825160208401519092506001600160a01b03811681146200026d57600080fd5b809150509250929050565b600065ffffffffffff808316818103620002a257634e487b7160e01b600052601160045260246000fd5b6001019392505050565b60805160a051615fdd620003186000396000818161075f01528181611ca301528181613cb301528181613ce201528181613d1a015261430e01526000818161047501528181610bb501528181611a410152818161237b01528181612d610152612df60152615fdd6000f3fe6080604052600436106101d65760003560e01c806304104edb146101db5780630ae09540146101fd57806316f6ee9a1461021d578063294daa491461025d5780632b38bafc1461027f5780632f7527cc1461029f5780633e79167f146102b457806340d6bb82146102d457806347c3e2cb146102ea5780634ffac83a14610385578063597d2f3c146103985780635d06b4ab146103b657806364d51a2a146103d657806373433a2f146103fe57806376f2e3f41461041e57806379ba50971461044e5780637d253aff1461046357806385c64e11146104a45780638c7cba66146104c65780638da5cb5b146104e65780638da92e71146105045780638eef585f146105245780639e20103614610544578063a21a23e414610564578063a4c0ed3614610579578063acfc6cdd14610599578063b2a7cac5146105c6578063b79fa6f7146105e6578063bd58017f146106cd578063bec4c08c146106ed578063c3fbb6fd1461070d578063cb6317971461072d578063cd0593df1461074d578063ce3f471914610781578063dac83d29146107a1578063db972c8b146107c1578063dc311dd3146107d4578063e30afa4a14610804578063f2fde38b14610849578063f99b1d6814610869578063f9c45ced14610889575b600080fd5b3480156101e757600080fd5b506101fb6101f6366004614abd565b6108a9565b005b34801561020957600080fd5b506101fb610218366004614ae1565b610a5e565b34801561022957600080fd5b5061024a610238366004614b11565b6000908152600c602052604090205490565b6040519081526020015b60405180910390f35b34801561026957600080fd5b5060015b60405160ff9091168152602001610254565b34801561028b57600080fd5b506101fb61029a366004614abd565b610cc0565b3480156102ab57600080fd5b5061026d600881565b3480156102c057600080fd5b506101fb6102cf366004614b2a565b610d21565b3480156102e057600080fd5b5061024a6103e881565b3480156102f657600080fd5b50610348610305366004614b11565b60106020526000908152604090205463ffffffff811690600160201b810462ffffff1690600160381b810461ffff1690600160481b90046001600160a01b031684565b6040805163ffffffff909516855262ffffff909316602085015261ffff909116918301919091526001600160a01b03166060820152608001610254565b61024a610393366004614cac565b610dd2565b3480156103a457600080fd5b506002546001600160601b031661024a565b3480156103c257600080fd5b506101fb6103d1366004614abd565b610f7e565b3480156103e257600080fd5b506103eb606481565b60405161ffff9091168152602001610254565b34801561040a57600080fd5b506101fb610419366004614d57565b61102a565b34801561042a57600080fd5b5061043e610439366004614dd9565b611115565b6040519015158152602001610254565b34801561045a57600080fd5b506101fb611453565b34801561046f57600080fd5b506104977f000000000000000000000000000000000000000000000000000000000000000081565b6040516102549190614e5b565b3480156104b057600080fd5b506104b96114fd565b6040516102549190614e6f565b3480156104d257600080fd5b506101fb6104e1366004614ec3565b611562565b3480156104f257600080fd5b506000546001600160a01b0316610497565b34801561051057600080fd5b506101fb61051f366004614f1d565b6115d6565b34801561053057600080fd5b506101fb61053f366004614f3a565b611640565b34801561055057600080fd5b5061024a61055f366004614f65565b61167c565b34801561057057600080fd5b5061024a61179d565b34801561058557600080fd5b506101fb610594366004615019565b6119e3565b3480156105a557600080fd5b506105b96105b4366004615068565b611bce565b60405161025491906150f2565b3480156105d257600080fd5b506101fb6105e1366004614b11565b611dd2565b3480156105f257600080fd5b506004546005546106629160ff80821692610100830482169262010000810483169263010000008204811692600160201b83049091169163ffffffff600160281b8204811692600160481b8304821692600160681b8104831692600160881b90910416906001600160601b03168a565b604080519a15158b5298151560208b01529615159789019790975260ff948516606089015292909316608087015263ffffffff90811660a087015291821660c0860152811660e08501529091166101008301526001600160601b031661012082015261014001610254565b3480156106d957600080fd5b50600a54610497906001600160a01b031681565b3480156106f957600080fd5b506101fb610708366004614ae1565b611f03565b34801561071957600080fd5b506101fb610728366004615105565b6120bf565b34801561073957600080fd5b506101fb610748366004614ae1565b6125ac565b34801561075957600080fd5b5061024a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561078d57600080fd5b506101fb61079c366004615159565b612899565b3480156107ad57600080fd5b506101fb6107bc366004614ae1565b6128b2565b61024a6107cf36600461519a565b6129c3565b3480156107e057600080fd5b506107f46107ef366004614b11565b612c21565b6040516102549493929190615272565b34801561081057600080fd5b50600b5461082c9063ffffffff80821691600160201b90041682565b6040805163ffffffff938416815292909116602083015201610254565b34801561085557600080fd5b506101fb610864366004614abd565b612d0e565b34801561087557600080fd5b506101fb6108843660046152be565b612d1f565b34801561089557600080fd5b5061024a6108a43660046152ea565b612e88565b6108b1612f9f565b60095460005b81811015610a3657826001600160a01b0316600982815481106108dc576108dc615330565b6000918252602090912001546001600160a01b031603610a2457600961090360018461535c565b8154811061091357610913615330565b600091825260209091200154600980546001600160a01b03909216918390811061093f5761093f615330565b600091825260209091200180546001600160a01b0319166001600160a01b039290921691909117905582600961097660018561535c565b8154811061098657610986615330565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060098054806109c5576109c561536f565b600082815260209020810160001990810180546001600160a01b03191690550190556040517ff80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af3790610a17908590614e5b565b60405180910390a1505050565b80610a2e81615385565b9150506108b7565b5081604051635428d44960e01b8152600401610a529190614e5b565b60405180910390fd5b50565b60008281526007602052604090205482906001600160a01b031680610a995760405163c5171ee960e01b815260048101839052602401610a52565b336001600160a01b03821614610ac45780604051636c51fda960e11b8152600401610a529190614e5b565b600454610100900460ff1615610aed5760405163769dd35360e11b815260040160405180910390fd5b600084815260086020526040902054600160601b90046001600160401b031615610b2a57604051631685ecdd60e31b815260040160405180910390fd5b6000848152600860209081526040918290208251808401909352546001600160601b038116808452600160601b9091046001600160401b031691830191909152610b7386612ff4565b600280546001600160601b03169082906000610b8f838561539e565b92506101000a8154816001600160601b0302191690836001600160601b031602179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb87846001600160601b03166040518363ffffffff1660e01b8152600401610c0a9291906153c5565b6020604051808303816000875af1158015610c29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4d91906153de565b610c7d5760405163cf47918160e01b81526001600160601b03808316600483015283166024820152604401610a52565b867f3784f77e8e883de95b5d47cd713ced01229fa74d118c0a462224bcb0516d43f18784604051610caf9291906153fb565b60405180910390a250505050505050565b610cc8612f9f565b600a546001600160a01b031615610cff57600a5460405163ea6d390560e01b8152610a52916001600160a01b031690600401614e5b565b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b610d29612f9f565b6064610d3b60a0830160808401615437565b60ff161180610d555750610d556040820160208301614f1d565b80610d6b5750610d6b6060820160408301614f1d565b15610d895760405163b0e7bd8360e01b815260040160405180910390fd5b806004610d96828261549d565b9050507e28d3a46e95e67def989d41c66eb331add9809460b95b5fb4eb006157728fc581604051610dc79190615670565b60405180910390a150565b60045460009062010000900460ff1615610dff5760405163ab35696f60e01b815260040160405180910390fd5b600454610100900460ff1615610e285760405163769dd35360e11b815260040160405180910390fd5b3415610e4957604051630b829bad60e21b8152346004820152602401610a52565b6000806000610e5a88338989613143565b925092509250600080610e6d338b61328e565b600087815260106020908152604091829020885181548a8401518b8601516060808e015163ffffffff90951666ffffffffffffff1990941693909317600160201b62ffffff9384160217600160381b600160e81b031916600160381b61ffff92831602600160481b600160e81b03191617600160481b6001600160a01b03909516949094029390931790935584513381526001600160401b038b1694810194909452918e169383019390935281018e9052908c16608082015260a081018390526001600160601b03821660c0820152919350915085907fb7933fba96b6b452eb44f99fdc08052a45dff82363d59abaff0456931c3d24599060e00160405180910390a2509298975050505050505050565b610f86612f9f565b610f8f8161346a565b15610faf578060405163ac8a27ef60e01b8152600401610a529190614e5b565b600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0180546001600160a01b0319166001600160a01b0383161790556040517fb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af0162590610dc7908390614e5b565b600a546001600160a01b0316331461105557604051634bea32db60e11b815260040160405180910390fd5b828015806110635750601f81115b1561108457604051634ecc4fef60e01b815260048101829052602401610a52565b8082146110a85760405163339f8a9d60e01b8152610a529082908490600401615758565b60005b8181101561110d576110fb8686838181106110c8576110c8615330565b90506020020160208101906110dd9190614abd565b8585848181106110ef576110ef615330565b90506020020135612d1f565b8061110581615385565b9150506110ab565b505050505050565b600a546000906001600160a01b0316331461114357604051634bea32db60e11b815260040160405180910390fd5b60045462010000900460ff161561116d5760405163ab35696f60e01b815260040160405180910390fd5b6001600160c01b038416156111aa57600680546001600160601b038616600160a01b02600160201b600160a01b0390911663ffffffff4216171790555b6001600160401b038316156111ff5760068054436001600160401b03908116600160201b02600160201b600160601b0319918716600160601b0291909116600160201b600160a01b0319909216919091171790555b600080866001600160401b0381111561121a5761121a614b68565b60405190808252806020026020018201604052801561125357816020015b6112406148f6565b8152602001906001900390816112385790505b50905060005b8781101561135557600089898381811061127557611275615330565b90506020028101906112879190615766565b611290906158ea565b9050600061129f82888b6134d3565b905085806112aa5750805b604083015151519096501515806112c957506040820151516020015115155b15611340576040805160808101825283516001600160401b0316815260208085015162ffffff168183015284830180515151938301939093529151519091015160608201528451859061ffff881690811061132657611326615330565b6020026020010181905250848061133c906159d0565b9550505b5050808061134d90615385565b915050611259565b5060008261ffff166001600160401b0381111561137457611374614b68565b6040519080825280602002602001820160405280156113ad57816020015b61139a6148f6565b8152602001906001900390816113925790505b50905060005b8361ffff16811015611409578281815181106113d1576113d1615330565b60200260200101518282815181106113eb576113eb615330565b6020026020010181905250808061140190615385565b9150506113b3565b507ff10ea936d00579b4c52035ee33bf46929646b3aa87554c565d8fb2c7aa549c448588888460405161143f94939291906159f1565b60405180910390a150505095945050505050565b6001546001600160a01b031633146114a65760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b6044820152606401610a52565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61150561492c565b6040805161010081019182905290600f90600890826000855b82829054906101000a900462ffffff1662ffffff168152602001906003019060208260020104928301926001038202915080841161151e5790505050505050905090565b61156a612f9f565b8051600b80546020808501805163ffffffff908116600160201b026001600160401b031990941695811695861793909317909355604080519485529251909116908301527f0cc54509a45ab33cd67614d4a2892c083ecf8fb43b9d29f6ea8130b9023e51df9101610dc7565b6115de612f9f565b60045460ff6201000090910416151581151514610a5b5760048054821515620100000262ff0000199091161790556040517f49ba7c1de2d8853088b6270e43df2118516b217f38b917dd2b80dea360860fbe90610dc790831515815260200190565b600a546001600160a01b0316331461166b57604051634bea32db60e11b815260040160405180910390fd5b611678600f82600861494b565b5050565b604080516101408101825260045460ff80821615158352610100808304821615156020808601919091526201000084048316151585870152630100000084048316606080870191909152600160201b808604909416608080880191909152600160281b860463ffffffff90811660a0890152600160481b8704811660c0890152600160681b8704811660e0890152600160881b9096048616938701939093526005546001600160601b039081166101208801528751938401885260065480871685526001600160401b03958104861693850193909352600160601b830490941696830196909652600160a01b90049091169381019390935260009283926117889288169187919061366b565b50506001600160601b03169695505050505050565b600454600090610100900460ff16156117c95760405163769dd35360e11b815260040160405180910390fd5b60045462010000900460ff16156117f35760405163ab35696f60e01b815260040160405180910390fd5b60003361180160014361535c565b6001546040516001600160601b0319606094851b81166020830152924060348201523090931b90911660548301526001600160c01b0319600160a01b90910460c01b16606882015260700160408051808303601f19018152919052805160209091012060018054919250600160a01b9091046001600160401b031690601461188883615a86565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550506000806001600160401b038111156118c7576118c7614b68565b6040519080825280602002602001820160405280156118f0578160200160208202803683370190505b5060408051808201825260008082526020808301828152878352600882528483209351845491516001600160601b039091166001600160a01b031992831617600160601b6001600160401b039092169190910217909355835160608101855233815280820183815281860187815289855260078452959093208151815486166001600160a01b03918216178255935160018201805490961694169390931790935592518051949550919390926119ad9260028501929101906149e9565b505060405133915083907f1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d90600090a350905090565b600454610100900460ff1615611a0c5760405163769dd35360e11b815260040160405180910390fd5b60045462010000900460ff1615611a365760405163ab35696f60e01b815260040160405180910390fd5b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611a7f576040516344b0e3c360e01b815260040160405180910390fd5b60208114611aa557604051636865567560e01b8152610a52906020908390600401615aaa565b6000611ab382840184614b11565b6000818152600760205260409020549091506001600160a01b0316611aee5760405163c5171ee960e01b815260048101829052602401610a52565b600081815260086020526040812080546001600160601b031691869190611b158385615abe565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555084600260008282829054906101000a90046001600160601b0316611b5d9190615abe565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a828784611bb09190615ade565b604051611bbe929190615758565b60405180910390a2505050505050565b600454606090610100900460ff1615611bfa5760405163769dd35360e11b815260040160405180910390fd5b60008381526010602081815260408084208151608081018352815463ffffffff8116825262ffffff600160201b8204168286015261ffff600160381b820416938201939093526001600160a01b03600160481b8404811660608301908152968a9052949093526001600160e81b031990911690559151163314611c9857806060015133604051638e30e82360e01b8152600401610a52929190615af1565b8051600090611cce907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16615b0b565b90506000611cda613715565b90506000836020015162ffffff1682611cf3919061535c565b9050808310611d385782846020015162ffffff1684611d129190615ade565b611d1d906001615ade565b6040516315ad27c360e01b8152600401610a52929190615758565b6001600160401b03831115611d63576040516302c6ef8160e11b815260048101849052602401610a52565b604051888152339088907f16f3f633197fafab10a5df69e6f3f2f7f20092f08d8d47de0a91c0f4b96a1a259060200160405180910390a3611dc68785600d6000611db1888a6020015161379f565b815260200190815260200160002054866137ae565b98975050505050505050565b600454610100900460ff1615611dfb5760405163769dd35360e11b815260040160405180910390fd5b6000818152600760205260409020546001600160a01b0316611e335760405163c5171ee960e01b815260048101829052602401610a52565b6000818152600760205260409020600101546001600160a01b03163314611e8a576000818152600760205260409081902060010154905163d084e97560e01b8152610a52916001600160a01b031690600401614e5b565b6000818152600760205260409081902080546001600160a01b031980821633908117845560019093018054909116905591516001600160a01b039092169183917fd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c938691611ef7918591615af1565b60405180910390a25050565b60008281526007602052604090205482906001600160a01b031680611f3e5760405163c5171ee960e01b815260048101839052602401610a52565b336001600160a01b03821614611f695780604051636c51fda960e11b8152600401610a529190614e5b565b600454610100900460ff1615611f925760405163769dd35360e11b815260040160405180910390fd5b60045462010000900460ff1615611fbc5760405163ab35696f60e01b815260040160405180910390fd5b60008481526007602052604090206002015460631901611fef576040516305a48e0f60e01b815260040160405180910390fd5b60036000611ffd8587613967565b815260208101919091526040016000205460ff166120b9576001600360006120258688613967565b815260208082019290925260409081016000908120805460ff191694151594909417909355868352600782528083206002018054600181018255908452919092200180546001600160a01b0319166001600160a01b0386161790555184907f1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1906120b0908690614e5b565b60405180910390a25b50505050565b600454610100900460ff16156120e85760405163769dd35360e11b815260040160405180910390fd5b6120f18361346a565b6121105782604051635428d44960e01b8152600401610a529190614e5b565b604081146121355760408051636865567560e01b8152610a5291908390600401615aaa565b600061214382840184615b22565b90506000806000806121588560200151612c21565b9350935093509350816001600160a01b0316336001600160a01b0316146121945781604051636c51fda960e11b8152600401610a529190614e5b565b876001600160a01b031663294daa496040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f69190615b5c565b60ff16856000015160ff1614612293578460000151886001600160a01b031663294daa496040518163ffffffff1660e01b8152600401602060405180830381865afa158015612249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061226d9190615b5c565b60405163e7aada9560e01b815260ff928316600482015291166024820152604401610a52565b6001600160401b038316156122bb57604051631685ecdd60e31b815260040160405180910390fd5b60006040518060a001604052806122d0600190565b60ff16815260200187602001518152602001846001600160a01b03168152602001838152602001866001600160601b031681525090506000816040516020016123199190615b79565b60405160208183030381529060405290506123378760200151612ff4565b600280548791906000906123559084906001600160601b031661539e565b92506101000a8154816001600160601b0302191690836001600160601b031602179055507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb8b886040518363ffffffff1660e01b81526004016123c79291906153fb565b6020604051808303816000875af11580156123e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061240a91906153de565b61244b5760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b6044820152606401610a52565b60405163ce3f471960e01b81526001600160a01b038b169063ce3f471990612477908490600401615c25565b600060405180830381600087803b15801561249157600080fd5b505af11580156124a5573d6000803e3d6000fd5b50506004805461ff00191661010017905550600090505b835181101561254f578381815181106124d7576124d7615330565b60200260200101516001600160a01b0316638ea981178c6040518263ffffffff1660e01b815260040161250a9190614e5b565b600060405180830381600087803b15801561252457600080fd5b505af1158015612538573d6000803e3d6000fd5b50505050808061254790615385565b9150506124bc565b506004805461ff00191690556020870151875160405160ff909116907fbd89b747474d3fc04664dfbd1d56ae7ffbe46ee097cdb9979c13916bb76269ce90612598908e90614e5b565b60405180910390a350505050505050505050565b60008281526007602052604090205482906001600160a01b0316806125e75760405163c5171ee960e01b815260048101839052602401610a52565b336001600160a01b038216146126125780604051636c51fda960e11b8152600401610a529190614e5b565b600454610100900460ff161561263b5760405163769dd35360e11b815260040160405180910390fd5b600084815260086020526040902054600160601b90046001600160401b03161561267857604051631685ecdd60e31b815260040160405180910390fd5b600360006126868587613967565b815260208101919091526040016000205460ff166126bb5783836040516379bfd40160e01b8152600401610a52929190615c38565b60008481526007602090815260408083206002018054825181850281018501909352808352919290919083018282801561271e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612700575b50505050509050600060018251612735919061535c565b905060005b825181101561284057856001600160a01b031683828151811061275f5761275f615330565b60200260200101516001600160a01b03160361282e57600083838151811061278957612789615330565b6020026020010151905080600760008a815260200190815260200160002060020183815481106127bb576127bb615330565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092558981526007909152604090206002018054806128065761280661536f565b600082815260209020810160001990810180546001600160a01b031916905501905550612840565b8061283881615385565b91505061273a565b506003600061284f8789613967565b815260208101919091526040908101600020805460ff191690555186907f32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a790611bbe908890614e5b565b604051632cb6686f60e01b815260040160405180910390fd5b60008281526007602052604090205482906001600160a01b0316806128ed5760405163c5171ee960e01b815260048101839052602401610a52565b336001600160a01b038216146129185780604051636c51fda960e11b8152600401610a529190614e5b565b600454610100900460ff16156129415760405163769dd35360e11b815260040160405180910390fd5b6000848152600760205260409020600101546001600160a01b038481169116146120b9576000848152600760205260409081902060010180546001600160a01b0319166001600160a01b0386161790555184907f21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1906120b09033908790615af1565b60045460009062010000900460ff16156129f05760405163ab35696f60e01b815260040160405180910390fd5b600454610100900460ff1615612a195760405163769dd35360e11b815260040160405180910390fd5b3415612a3a57604051630b829bad60e21b8152346004820152602401610a52565b600080612a4989338a8a613143565b925050915060006040518061010001604052808481526020018a61ffff168152602001336001600160a01b031681526020018781526020018863ffffffff166001600160601b031681526020018b81526020016000815260200160008152509050600080612ab68361397d565b60c087019190915260e08601919091526040519193509150612ae29085908c908f908790602001615c4f565b60405160208183030381529060405280519060200120600c6000878152602001908152602001600020819055506000604051806101600160405280878152602001336001600160a01b03168152602001866001600160401b031681526020018c62ffffff1681526020018e81526020018d61ffff1681526020018b63ffffffff1681526020018581526020018a8152602001848152602001836001600160601b0316815250905080600001517f01872fb9c7d6d68af06a17347935e04412da302a377224c205e672c26e18c37f82602001518360400151846060015185608001518660a001518760c001518860e0015160c001518960e0015160e001518a61010001518b61012001518c6101400151604051612c089b9a99989796959493929190615d02565b60405180910390a250939b9a5050505050505050505050565b600081815260076020526040812054819081906060906001600160a01b0316612c605760405163c5171ee960e01b815260048101869052602401610a52565b60008581526008602090815260408083205460078352928190208054600290910180548351818602810186019094528084526001600160601b03861695600160601b90046001600160401b0316946001600160a01b03909316939192839190830182828015612cf857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612cda575b5050505050905093509350935093509193509193565b612d16612f9f565b610a5b81613bb4565b600a546001600160a01b03163314612d4a57604051634bea32db60e11b815260040160405180910390fd5b60405163a9059cbb60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb90612d9890859085906004016153c5565b6020604051808303816000875af1158015612db7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ddb91906153de565b611678576040516370a0823160e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190612e2b903090600401614e5b565b602060405180830381865afa158015612e48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e6c9190615d98565b8160405163cf47918160e01b8152600401610a52929190615758565b604080516101408101825260045460ff80821615158352610100808304821615156020808601919091526201000084048316151585870152630100000084048316606080870191909152600160201b80860490941660808088019190915263ffffffff600160281b8704811660a0890152600160481b8704811660c0890152600160681b8704811660e0890152600160881b9096048616938701939093526005546001600160601b039081166101208801528751938401885260065495861684526001600160401b03948604851692840192909252600160601b850490931695820195909552600160a01b90920490931692810192909252600091612f8d9190613c57565b6001600160601b031690505b92915050565b6000546001600160a01b03163314612ff25760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b6044820152606401610a52565b565b6000818152600760209081526040808320815160608101835281546001600160a01b0390811682526001830154168185015260028201805484518187028101870186528181529295939486019383018282801561307a57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161305c575b505050505081525050905060005b8160400151518110156130ea57600360006130c0846040015184815181106130b2576130b2615330565b602002602001015186613967565b81526020810191909152604001600020805460ff19169055806130e281615385565b915050613088565b50600082815260076020526040812080546001600160a01b031990811682556001820180549091169055906131226002830182614a3e565b505050600090815260086020526040902080546001600160a01b0319169055565b600061314d614a5c565b60006103e88561ffff16111561317c57846103e8604051634a90778560e01b8152600401610a52929190615aaa565b8461ffff166000036131a1576040516308fad2a760e01b815260040160405180910390fd5b6000806131ac613c9d565b600e54919350915065ffffffffffff1660006132178b8b84604080513060208201529081018490526001600160a01b038316606082015265ffffffffffff8216608082015260009060a00160408051601f198184030181529190528051602090910120949350505050565b9050613224826001615db1565b600e805465ffffffffffff9290921665ffffffffffff199092169190911790556040805160808101825263ffffffff909416845262ffffff8916602085015261ffff8a16908401526001600160a01b038a1660608401529550909350909150509450945094915050565b604080516080808201835260065463ffffffff8082168452600160201b8083046001600160401b03908116602080880191909152600160601b850490911686880152600160a01b9093046001600160601b0390811660608088019190915287516101408101895260045460ff808216151583526101008083048216151598840198909852620100008204811615159a83019a909a52630100000081048a169282019290925292810490971694820194909452600160281b8604821660a0820152600160481b8604821660c0820152600160681b8604821660e0820152600160881b9095041690840152600554166101208301526000918291906003836133948888613967565b815260208101919091526040016000205460ff166133c95784866040516379bfd40160e01b8152600401610a52929190615c38565b60006133d58284613c57565b600087815260086020526040902080546001600160601b0392831693509091168281101561342457815460405163cf47918160e01b8152610a52916001600160601b0316908590600401615dd0565b81546001600160601b0319908116918490036001600160601b038181169390931790935560028054918216918316859003909216179055909450925050505b9250929050565b6000805b6009548110156134ca57826001600160a01b03166009828154811061349557613495615330565b6000918252602090912001546001600160a01b0316036134b85750600192915050565b806134c281615385565b91505061346e565b50600092915050565b6000826001600160401b031684600001516001600160401b0316111561352257835160405163012d824d60e01b81526001600160401b0380861660048301529091166024820152604401610a52565b606084015151604080860151905160009161353f91602001615de9565b60405160208183030381529060405280519060200120905085604001516000015160006002811061357257613572615330565b602002015115801561358b575060408601515160200151155b156135c557600d60006135af88600001516001600160401b0316896020015161379f565b8152602001908152602001600020549050613648565b856080015115613648576000600d60006135f089600001516001600160401b03168a6020015161379f565b81526020810191909152604001600020549050806136425781600d60006136288a600001516001600160401b03168b6020015161379f565b815260208101919091526040016000205560019350613646565b8091505b505b6000613655838389613d70565b905083806136605750805b979650505050505050565b60008060008061367b8686614189565b6001600160401b031690506000613693826010615b0b565b9050600060146136a4836015615b0b565b6136ae9190615e29565b89516136ba9190615b0b565b838960e0015163ffffffff168c6136d19190615abe565b6001600160601b03166136e49190615b0b565b6136ee9190615ade565b90506000806137008360008c8c614202565b909d909c50949a509398505050505050505050565b60004661a4b181148061372a575062066eed81145b156137985760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561376e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137929190615d98565b91505090565b4391505090565b62ffffff1660189190911b1790565b6060826137e05760405163220a34e960e11b8152600481018690526001600160401b0383166024820152604401610a52565b604080516020808201889052865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff1611156138825784604001516103e8604051634a90778560e01b8152600401610a52929190615aaa565b6000856040015161ffff166001600160401b038111156138a4576138a4614b68565b6040519080825280602002602001820160405280156138cd578160200160208202803683370190505b50905060005b866040015161ffff168161ffff16101561395c57828160405160200161391092919091825260f01b6001600160f01b031916602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff168151811061393f5761393f615330565b602090810291909101015280613954816159d0565b9150506138d3565b509695505050505050565b60a081901b6001600160a01b0383161792915050565b6000806000806003600061399987604001518860a00151613967565b815260208101919091526040016000205460ff166139d6578460a0015185604001516040516379bfd40160e01b8152600401610a52929190615c38565b604080516080808201835260065463ffffffff80821684526001600160401b03600160201b8084048216602080880191909152600160601b8504909216868801526001600160601b03600160a01b909404841660608088019190915287516101408101895260045460ff808216151583526101008083048216151596840196909652620100008204811615159a83019a909a52630100000081048a168284015292830490981688870152600160281b8204841660a0890152600160481b8204841660c0890152600160681b8204841660e0890152600160881b90910490921690860152600554909116610120850152908801519088015191929160009182918291613ae291868861366b565b60a08d0151600090815260086020526040902080546001600160601b0394851697509295509093509116841115613b3a57805460405163cf47918160e01b8152610a52916001600160601b0316908690600401615dd0565b80546001600160601b0360016001600160401b03600160601b8085048216929092011602818116828416178790038083166001600160601b03199283166001600160a01b03199095169490941793909317909355600280548083168890039092169190931617909155929a91995097509095509350505050565b336001600160a01b03821603613c065760405162461bcd60e51b815260206004820152601760248201527621b0b73737ba103a3930b739b332b9103a379039b2b63360491b6044820152606401610a52565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600080613c648484614189565b8460c0015163ffffffff16613c799190615e3d565b6001600160401b031690506000613c938260008787614202565b5095945050505050565b6000806000613caa613715565b90506000613cd87f000000000000000000000000000000000000000000000000000000000000000083615e60565b9050600081613d077f000000000000000000000000000000000000000000000000000000000000000085615ade565b613d11919061535c565b90506000613d3f7f000000000000000000000000000000000000000000000000000000000000000083615e29565b905063ffffffff8110613d65576040516307b2a52360e41b815260040160405180910390fd5b909590945092505050565b6000806040518060c00160405280866001600160401b03811115613d9657613d96614b68565b604051908082528060200260200182016040528015613dbf578160200160208202803683370190505b508152602001866001600160401b03811115613ddd57613ddd614b68565b6040519080825280601f01601f191660200182016040528015613e07576020820181803683370190505b508152602001866001600160401b03811115613e2557613e25614b68565b604051908082528060200260200182016040528015613e5857816020015b6060815260200190600190039081613e435790505b50815260006020820152604001866001600160401b03811115613e7d57613e7d614b68565b604051908082528060200260200182016040528015613ea6578160200160208202803683370190505b508152602001866001600160401b03811115613ec457613ec4614b68565b604051908082528060200260200182016040528015613eed578160200160208202803683370190505b509052905060005b8581101561406c57600084606001518281518110613f1557613f15615330565b60200260200101519050600080600080613f3989600001518a602001518c8861424d565b93509350935093508315613f8d57828760400151886060015161ffff1681518110613f6657613f66615330565b602090810291909101015260608701805190613f81826159d0565b61ffff16905250613fc0565b600160f81b87602001518781518110613fa857613fa8615330565b60200101906001600160f81b031916908160001a9053505b8780613fca575080155b85515188518051929a50909188908110613fe657613fe6615330565b602002602001018181525050818760800151878151811061400957614009615330565b60200260200101906001600160601b031690816001600160601b031681525050846000015160a001518760a00151878151811061404857614048615330565b6020026020010181815250505050505050808061406490615385565b915050613ef5565b5060608301515115614181576000816060015161ffff166001600160401b0381111561409a5761409a614b68565b6040519080825280602002602001820160405280156140cd57816020015b60608152602001906001900390816140b85790505b50905060005b826060015161ffff1681101561413157826040015181815181106140f9576140f9615330565b602002602001015182828151811061411357614113615330565b6020026020010181905250808061412990615385565b9150506140d3565b5081516020830151608084015160a08501516040517f8f79f730779e875ce76c428039cc2052b5b5918c2a55c598fab251c1198aec549461417794909390928792615ead565b60405180910390a1505b509392505050565b815160009080156141a6575060408201516001600160401b031615155b156141fa5761010083015163ffffffff16431080806141e757506101008401516141d69063ffffffff164361535c565b83602001516001600160401b031610155b156141f85750506040810151612f99565b505b503a92915050565b600080600060648560600151606461421a9190615f50565b6142279060ff1689615b0b565b6142319190615e29565b905061423f818787876145bc565b925092505094509492505050565b805160a09081015160009081526008602090815260408083208551948501519151939460609486948594859261428b928e928e929091879101615c4f565b60408051601f19818403018152918152815160209283012084516000908152600c90935291205490915081146142fe5750505460408051808201909152601081526f756e6b6e6f776e2063616c6c6261636b60801b60208201526001955093506001600160601b031691508390506145b1565b50614307614a5c565b600061433c7f00000000000000000000000000000000000000000000000000000000000000006001600160401b038e16615e29565b6040805160808101825263ffffffff909216825262ffffff8d1660208084019190915285015161ffff16828201528401516001600160a01b0316606082015291508990506143d0575050604080518082019091526016815275756e617661696c61626c652072616e646f6d6e65737360501b60208201529054600195509093506001600160601b03169150600090506145b1565b60006143e28360000151838c8f6137ae565b606080840151855191860151604051939450909260009263d21ea8fd60e01b9261441192879190602401615f69565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526004805461ff00191661010017905590506000805a9050600061447c8e60000151608001516001600160601b0316896040015186614630565b9093509050806144b1578d516080015160405163aad1598360e01b8152610a52916001600160601b0316908490600401615758565b506000610bb85a6144c29190615ade565b6004805461ff00191690559050818110156144eb576144eb6144e4828461535c565b8f5161466f565b8854600160601b90046001600160401b031689600c61450983615f94565b82546001600160401b039182166101009390930a92830291909202199091161790555087516000908152600c60205260408120558261457f5760408051808201909152601081526f195e1958dd5d1a5bdb8819985a5b195960821b60208201528954600191906001600160601b0316600061459f565b604080516020810190915260008082528a549091906001600160601b0316825b9c509c509c509c505050505050505050505b945094509450949050565b6000808085156145cc57856145d6565b6145d6858561489b565b90506000816145ed89670de0b6b3a7640000615b0b565b6145f79190615e29565b9050676765c793fa10079d601b1b8111156146245760405162de437160e81b815260040160405180910390fd5b97909650945050505050565b6000805a610bb8811061466657610bb881039050856040820482031115614666576000808551602087016000898bf19250600191505b50935093915050565b80608001516001600160601b0316821115614688575050565b6004546000906064906146a590600160201b900460ff1682615fb7565b60ff168360c001518585608001516001600160601b03166146c6919061535c565b6146d09190615b0b565b6146da9190615b0b565b6146e49190615e29565b60e080840151604080516101408101825260045460ff80821615158352610100808304821615156020808601919091526201000084048316151585870152630100000084048316606080870191909152600160201b80860490941660808088019190915263ffffffff600160281b8704811660a0890152600160481b8704811660c0890152600160681b870481169a88019a909a52600160881b9095048916928601929092526005546001600160601b039081166101208701528651948501875260065498891685526001600160401b03938904841691850191909152600160601b880490921694830194909452600160a01b909504909416918401919091529293506000926147f792859291906145bc565b5060a08401516000908152600860205260408120805492935083929091906148299084906001600160601b0316615abe565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600260008282829054906101000a90046001600160601b03166148719190615abe565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555050505050565b60a0820151606082015160009190600163ffffffff831611908180156148d7575084516148ce9063ffffffff164261535c565b8363ffffffff16105b156148e457506101208501515b6001600160601b031695945050505050565b604051806080016040528060006001600160401b03168152602001600062ffffff16815260200160008152602001600081525090565b6040518061010001604052806008906020820280368337509192915050565b6001830191839082156149d95791602002820160005b838211156149a857833562ffffff1683826101000a81548162ffffff021916908362ffffff1602179055509260200192600301602081600201049283019260010302614961565b80156149d75782816101000a81549062ffffff02191690556003016020816002010492830192600103026149a8565b505b506149e5929150614a83565b5090565b8280548282559060005260206000209081019282156149d9579160200282015b828111156149d957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614a09565b5080546000825590600052602060002090810190610a5b9190614a83565b60408051608081018252600080825260208201819052918101829052606081019190915290565b5b808211156149e55760008155600101614a84565b6001600160a01b0381168114610a5b57600080fd5b8035614ab881614a98565b919050565b600060208284031215614acf57600080fd5b8135614ada81614a98565b9392505050565b60008060408385031215614af457600080fd5b823591506020830135614b0681614a98565b809150509250929050565b600060208284031215614b2357600080fd5b5035919050565b60006101408284031215614b3d57600080fd5b50919050565b803561ffff81168114614ab857600080fd5b803562ffffff81168114614ab857600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614ba057614ba0614b68565b60405290565b60405161010081016001600160401b0381118282101715614ba057614ba0614b68565b60405160a081016001600160401b0381118282101715614ba057614ba0614b68565b604051602081016001600160401b0381118282101715614ba057614ba0614b68565b604051601f8201601f191681016001600160401b0381118282101715614c3557614c35614b68565b604052919050565b600082601f830112614c4e57600080fd5b81356001600160401b03811115614c6757614c67614b68565b614c7a601f8201601f1916602001614c0d565b818152846020838601011115614c8f57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215614cc257600080fd5b84359350614cd260208601614b43565b9250614ce060408601614b55565b915060608501356001600160401b03811115614cfb57600080fd5b614d0787828801614c3d565b91505092959194509250565b60008083601f840112614d2557600080fd5b5081356001600160401b03811115614d3c57600080fd5b6020830191508360208260051b850101111561346357600080fd5b60008060008060408587031215614d6d57600080fd5b84356001600160401b0380821115614d8457600080fd5b614d9088838901614d13565b90965094506020870135915080821115614da957600080fd5b50614db687828801614d13565b95989497509550505050565b80356001600160401b0381168114614ab857600080fd5b600080600080600060808688031215614df157600080fd5b85356001600160401b03811115614e0757600080fd5b614e1388828901614d13565b90965094505060208601356001600160c01b0381168114614e3357600080fd5b9250614e4160408701614dc2565b9150614e4f60608701614dc2565b90509295509295909350565b6001600160a01b0391909116815260200190565b6101008101818360005b6008811015614e9d57815162ffffff16835260209283019290910190600101614e79565b50505092915050565b63ffffffff81168114610a5b57600080fd5b8035614ab881614ea6565b600060408284031215614ed557600080fd5b614edd614b7e565b8235614ee881614ea6565b81526020830135614ef881614ea6565b60208201529392505050565b8015158114610a5b57600080fd5b8035614ab881614f04565b600060208284031215614f2f57600080fd5b8135614ada81614f04565b6000610100808385031215614f4e57600080fd5b838184011115614f5d57600080fd5b509092915050565b60008060008060808587031215614f7b57600080fd5b843593506020850135614f8d81614ea6565b925060408501356001600160401b0380821115614fa957600080fd5b614fb588838901614c3d565b93506060870135915080821115614fcb57600080fd5b50614d0787828801614c3d565b60008083601f840112614fea57600080fd5b5081356001600160401b0381111561500157600080fd5b60208301915083602082850101111561346357600080fd5b6000806000806060858703121561502f57600080fd5b843561503a81614a98565b93506020850135925060408501356001600160401b0381111561505c57600080fd5b614db687828801614fd8565b60008060006060848603121561507d57600080fd5b833592506020840135915060408401356001600160401b038111156150a157600080fd5b6150ad86828701614c3d565b9150509250925092565b600081518084526020808501945080840160005b838110156150e7578151875295820195908201906001016150cb565b509495945050505050565b602081526000614ada60208301846150b7565b60008060006040848603121561511a57600080fd5b833561512581614a98565b925060208401356001600160401b0381111561514057600080fd5b61514c86828701614fd8565b9497909650939450505050565b6000806020838503121561516c57600080fd5b82356001600160401b0381111561518257600080fd5b61518e85828601614fd8565b90969095509350505050565b60008060008060008060c087890312156151b357600080fd5b863595506151c360208801614b43565b94506151d160408801614b55565b935060608701356151e181614ea6565b925060808701356001600160401b03808211156151fd57600080fd5b6152098a838b01614c3d565b935060a089013591508082111561521f57600080fd5b5061522c89828a01614c3d565b9150509295509295509295565b600081518084526020808501945080840160005b838110156150e75781516001600160a01b03168752958201959082019060010161524d565b6001600160601b03851681526001600160401b03841660208201526001600160a01b03831660408201526080606082018190526000906152b490830184615239565b9695505050505050565b600080604083850312156152d157600080fd5b82356152dc81614a98565b946020939093013593505050565b600080604083850312156152fd57600080fd5b8235915060208301356001600160401b0381111561531a57600080fd5b61532685828601614c3d565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115612f9957612f99615346565b634e487b7160e01b600052603160045260246000fd5b60006001820161539757615397615346565b5060010190565b6001600160601b038281168282160390808211156153be576153be615346565b5092915050565b6001600160a01b03929092168252602082015260400190565b6000602082840312156153f057600080fd5b8151614ada81614f04565b6001600160a01b039290921682526001600160601b0316602082015260400190565b60ff81168114610a5b57600080fd5b8035614ab88161541d565b60006020828403121561544957600080fd5b8135614ada8161541d565b60008135612f9981614f04565b60008135612f998161541d565b60008135612f9981614ea6565b6001600160601b0381168114610a5b57600080fd5b60008135612f998161547b565b81356154a881614f04565b815490151560ff1660ff19919091161781556154e36154c960208401615454565b82805461ff00191691151560081b61ff0016919091179055565b61550e6154f260408401615454565b82805462ff0000191691151560101b62ff000016919091179055565b61553761551d60608401615461565b825463ff000000191660189190911b63ff00000016178255565b61556461554660808401615461565b82805460ff60201b191660209290921b60ff60201b16919091179055565b61559761557360a0840161546e565b82805463ffffffff60281b191660289290921b63ffffffff60281b16919091179055565b6155ca6155a660c0840161546e565b82805463ffffffff60481b191660489290921b63ffffffff60481b16919091179055565b6155fd6155d960e0840161546e565b82805463ffffffff60681b191660689290921b63ffffffff60681b16919091179055565b61563161560d610100840161546e565b82805463ffffffff60881b191660889290921b63ffffffff60881b16919091179055565b6116786156416101208401615490565b6001830180546001600160601b0319166001600160601b0392909216919091179055565b8035614ab88161547b565b61014081016156888261568285614f12565b15159052565b61569460208401614f12565b151560208301526156a760408401614f12565b151560408301526156ba6060840161542c565b60ff1660608301526156ce6080840161542c565b60ff1660808301526156e260a08401614eb8565b63ffffffff1660a08301526156f960c08401614eb8565b63ffffffff1660c083015261571060e08401614eb8565b63ffffffff1660e0830152610100615729848201614eb8565b63ffffffff1690830152610120615741848201615665565b6001600160601b038116848301525b505092915050565b918252602082015260400190565b6000823560be1983360301811261577c57600080fd5b9190910192915050565b600082601f83011261579757600080fd5b813560206001600160401b03808311156157b3576157b3614b68565b8260051b6157c2838201614c0d565b93845285810183019383810190888611156157dc57600080fd5b84880192505b85831015611dc6578235848111156157f957600080fd5b8801601f196040828c038201121561581057600080fd5b615818614b7e565b878301358781111561582957600080fd5b8301610100818e038401121561583e57600080fd5b615846614ba6565b925088810135835261585a60408201614b43565b8984015261586a60608201614aad565b604084015260808101358881111561588157600080fd5b61588f8e8b83850101614c3d565b6060850152506158a160a08201615665565b608084015260c081013560a084015260e081013560c084015261010081013560e0840152508181526158d560408401615665565b818901528452505091840191908401906157e2565b600081360360c08112156158fd57600080fd5b615905614bc9565b61590e84614dc2565b8152602061591d818601614b55565b828201526040603f198401121561593357600080fd5b61593b614beb565b925036605f86011261594c57600080fd5b615954614b7e565b80608087013681111561596657600080fd5b604088015b81811015615982578035845292840192840161596b565b50908552604084019490945250509035906001600160401b038211156159a757600080fd5b6159b336838601615786565b60608201526159c460a08501614f12565b60808201529392505050565b600061ffff8083168181036159e7576159e7615346565b6001019392505050565b6000608080830160018060401b038089168552602060018060c01b038916818701526040828916818801526060858189015284895180875260a08a019150848b01965060005b81811015615a735787518051881684528681015162ffffff16878501528581015186850152840151848401529685019691880191600101615a37565b50909d9c50505050505050505050505050565b60006001600160401b038281166002600160401b031981016159e7576159e7615346565b61ffff929092168252602082015260400190565b6001600160601b038181168382160190808211156153be576153be615346565b80820180821115612f9957612f99615346565b6001600160a01b0392831681529116602082015260400190565b8082028115828204841417612f9957612f99615346565b600060408284031215615b3457600080fd5b615b3c614b7e565b8235615b478161541d565b81526020928301359281019290925250919050565b600060208284031215615b6e57600080fd5b8151614ada8161541d565b6020815260ff82511660208201526020820151604082015260018060a01b0360408301511660608201526000606083015160a06080840152615bbe60c0840182615239565b608094909401516001600160601b031660a093909301929092525090919050565b6000815180845260005b81811015615c0557602081850181015186830182015201615be9565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000614ada6020830184615bdf565b9182526001600160a01b0316602082015260400190565b60018060401b038516815262ffffff84166020820152826040820152608060608201528151608082015261ffff60208301511660a082015260018060a01b0360408301511660c0820152600060608301516101008060e0850152615cb7610180850183615bdf565b91506080850151615cd2828601826001600160601b03169052565b505060a084015161012084015260c084015161014084015260e08401516101608401528091505095945050505050565b6001600160a01b038c1681526001600160401b038b16602082015262ffffff8a1660408201526060810189905261ffff8816608082015263ffffffff871660a082015260c0810186905260e081018590526101606101008201819052600090615d6d83820187615bdf565b61012084019590955250506001600160601b0391909116610140909101529998505050505050505050565b600060208284031215615daa57600080fd5b5051919050565b65ffffffffffff8181168382160190808211156153be576153be615346565b6001600160601b03929092168252602082015260400190565b815160408201908260005b6002811015614e9d578251825260209283019290910190600101615df4565b634e487b7160e01b600052601260045260246000fd5b600082615e3857615e38615e13565b500490565b6001600160401b0381811683821602808216919082811461575057615750615346565b600082615e6f57615e6f615e13565b500690565b600081518084526020808501945080840160005b838110156150e75781516001600160601b031687529582019590820190600101615e88565b60a081526000615ec060a08301886150b7565b602083820381850152615ed38289615bdf565b915083820360408501528187518084528284019150828160051b850101838a0160005b83811015615f2457601f19878403018552615f12838351615bdf565b94860194925090850190600101615ef6565b50508681036060880152615f38818a615e74565b9450505050508281036080840152611dc681856150b7565b60ff8181168382160190811115612f9957612f99615346565b838152606060208201526000615f8260608301856150b7565b82810360408401526152b48185615bdf565b60006001600160401b03821680615fad57615fad615346565b6000190192915050565b60ff8281168282160390811115612f9957612f9961534656fea164736f6c6343000813000a",
}

var VRFCoordinatorABI = VRFCoordinatorMetaData.ABI

var VRFCoordinatorBin = VRFCoordinatorMetaData.Bin

func DeployVRFCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, beaconPeriodBlocksArg *big.Int, linkToken common.Address) (common.Address, *types.Transaction, *VRFCoordinator, error) {
	parsed, err := VRFCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorBin), backend, beaconPeriodBlocksArg, linkToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinator{VRFCoordinatorCaller: VRFCoordinatorCaller{contract: contract}, VRFCoordinatorTransactor: VRFCoordinatorTransactor{contract: contract}, VRFCoordinatorFilterer: VRFCoordinatorFilterer{contract: contract}}, nil
}

type VRFCoordinator struct {
	VRFCoordinatorCaller
	VRFCoordinatorTransactor
	VRFCoordinatorFilterer
}

type VRFCoordinatorCaller struct {
	contract *bind.BoundContract
}

type VRFCoordinatorTransactor struct {
	contract *bind.BoundContract
}

type VRFCoordinatorFilterer struct {
	contract *bind.BoundContract
}

type VRFCoordinatorSession struct {
	Contract     *VRFCoordinator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorCallerSession struct {
	Contract *VRFCoordinatorCaller
	CallOpts bind.CallOpts
}

type VRFCoordinatorTransactorSession struct {
	Contract     *VRFCoordinatorTransactor
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorRaw struct {
	Contract *VRFCoordinator
}

type VRFCoordinatorCallerRaw struct {
	Contract *VRFCoordinatorCaller
}

type VRFCoordinatorTransactorRaw struct {
	Contract *VRFCoordinatorTransactor
}

func NewVRFCoordinator(address common.Address, backend bind.ContractBackend) (*VRFCoordinator, error) {
	contract, err := bindVRFCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinator{VRFCoordinatorCaller: VRFCoordinatorCaller{contract: contract}, VRFCoordinatorTransactor: VRFCoordinatorTransactor{contract: contract}, VRFCoordinatorFilterer: VRFCoordinatorFilterer{contract: contract}}, nil
}

func NewVRFCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorCaller, error) {
	contract, err := bindVRFCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorCaller{contract: contract}, nil
}

func NewVRFCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorTransactor, error) {
	contract, err := bindVRFCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorTransactor{contract: contract}, nil
}

func NewVRFCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorFilterer, error) {
	contract, err := bindVRFCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorFilterer{contract: contract}, nil
}

func bindVRFCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFCoordinator *VRFCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinator.Contract.VRFCoordinatorCaller.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinator *VRFCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.VRFCoordinatorTransactor.contract.Transfer(opts)
}

func (_VRFCoordinator *VRFCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.VRFCoordinatorTransactor.contract.Transact(opts, method, params...)
}

func (_VRFCoordinator *VRFCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinator.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinator *VRFCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.contract.Transfer(opts)
}

func (_VRFCoordinator *VRFCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.contract.Transact(opts, method, params...)
}

func (_VRFCoordinator *VRFCoordinatorCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinator.Contract.MAXCONSUMERS(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinator.Contract.MAXCONSUMERS(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) MAXNUMWORDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) MAXNUMWORDS() (*big.Int, error) {
	return _VRFCoordinator.Contract.MAXNUMWORDS(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) MAXNUMWORDS() (*big.Int, error) {
	return _VRFCoordinator.Contract.MAXNUMWORDS(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFCoordinator.Contract.NUMCONFDELAYS(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFCoordinator.Contract.NUMCONFDELAYS(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) GetCallbackMemo(opts *bind.CallOpts, requestId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "getCallbackMemo", requestId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) GetCallbackMemo(requestId *big.Int) ([32]byte, error) {
	return _VRFCoordinator.Contract.GetCallbackMemo(&_VRFCoordinator.CallOpts, requestId)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) GetCallbackMemo(requestId *big.Int) ([32]byte, error) {
	return _VRFCoordinator.Contract.GetCallbackMemo(&_VRFCoordinator.CallOpts, requestId)
}

func (_VRFCoordinator *VRFCoordinatorCaller) GetConfirmationDelays(opts *bind.CallOpts) ([8]*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "getConfirmationDelays")

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) GetConfirmationDelays() ([8]*big.Int, error) {
	return _VRFCoordinator.Contract.GetConfirmationDelays(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) GetConfirmationDelays() ([8]*big.Int, error) {
	return _VRFCoordinator.Contract.GetConfirmationDelays(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) GetFee(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "getFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) GetFee(arg0 *big.Int, arg1 []byte) (*big.Int, error) {
	return _VRFCoordinator.Contract.GetFee(&_VRFCoordinator.CallOpts, arg0, arg1)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) GetFee(arg0 *big.Int, arg1 []byte) (*big.Int, error) {
	return _VRFCoordinator.Contract.GetFee(&_VRFCoordinator.CallOpts, arg0, arg1)
}

func (_VRFCoordinator *VRFCoordinatorCaller) GetFulfillmentFee(opts *bind.CallOpts, arg0 *big.Int, callbackGasLimit uint32, arguments []byte, arg3 []byte) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "getFulfillmentFee", arg0, callbackGasLimit, arguments, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) GetFulfillmentFee(arg0 *big.Int, callbackGasLimit uint32, arguments []byte, arg3 []byte) (*big.Int, error) {
	return _VRFCoordinator.Contract.GetFulfillmentFee(&_VRFCoordinator.CallOpts, arg0, callbackGasLimit, arguments, arg3)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) GetFulfillmentFee(arg0 *big.Int, callbackGasLimit uint32, arguments []byte, arg3 []byte) (*big.Int, error) {
	return _VRFCoordinator.Contract.GetFulfillmentFee(&_VRFCoordinator.CallOpts, arg0, callbackGasLimit, arguments, arg3)
}

func (_VRFCoordinator *VRFCoordinatorCaller) GetSubscription(opts *bind.CallOpts, subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(struct {
		Balance             *big.Int
		PendingFulfillments uint64
		Owner               common.Address
		Consumers           []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PendingFulfillments = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VRFCoordinator *VRFCoordinatorSession) GetSubscription(subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	return _VRFCoordinator.Contract.GetSubscription(&_VRFCoordinator.CallOpts, subId)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) GetSubscription(subId *big.Int) (struct {
	Balance             *big.Int
	PendingFulfillments uint64
	Owner               common.Address
	Consumers           []common.Address
}, error) {
	return _VRFCoordinator.Contract.GetSubscription(&_VRFCoordinator.CallOpts, subId)
}

func (_VRFCoordinator *VRFCoordinatorCaller) GetSubscriptionLinkBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "getSubscriptionLinkBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) GetSubscriptionLinkBalance() (*big.Int, error) {
	return _VRFCoordinator.Contract.GetSubscriptionLinkBalance(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) GetSubscriptionLinkBalance() (*big.Int, error) {
	return _VRFCoordinator.Contract.GetSubscriptionLinkBalance(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFCoordinator.Contract.IBeaconPeriodBlocks(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFCoordinator.Contract.IBeaconPeriodBlocks(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) ILink() (common.Address, error) {
	return _VRFCoordinator.Contract.ILink(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) ILink() (common.Address, error) {
	return _VRFCoordinator.Contract.ILink(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) MigrationVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "migrationVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) MigrationVersion() (uint8, error) {
	return _VRFCoordinator.Contract.MigrationVersion(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) MigrationVersion() (uint8, error) {
	return _VRFCoordinator.Contract.MigrationVersion(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) OnMigration(opts *bind.CallOpts, arg0 []byte) error {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "onMigration", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_VRFCoordinator *VRFCoordinatorSession) OnMigration(arg0 []byte) error {
	return _VRFCoordinator.Contract.OnMigration(&_VRFCoordinator.CallOpts, arg0)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) OnMigration(arg0 []byte) error {
	return _VRFCoordinator.Contract.OnMigration(&_VRFCoordinator.CallOpts, arg0)
}

func (_VRFCoordinator *VRFCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) Owner() (common.Address, error) {
	return _VRFCoordinator.Contract.Owner(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) Owner() (common.Address, error) {
	return _VRFCoordinator.Contract.Owner(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) SCallbackConfig(opts *bind.CallOpts) (struct {
	MaxCallbackGasLimit        uint32
	MaxCallbackArgumentsLength uint32
}, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "s_callbackConfig")

	outstruct := new(struct {
		MaxCallbackGasLimit        uint32
		MaxCallbackArgumentsLength uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxCallbackGasLimit = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.MaxCallbackArgumentsLength = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinator *VRFCoordinatorSession) SCallbackConfig() (struct {
	MaxCallbackGasLimit        uint32
	MaxCallbackArgumentsLength uint32
}, error) {
	return _VRFCoordinator.Contract.SCallbackConfig(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) SCallbackConfig() (struct {
	MaxCallbackGasLimit        uint32
	MaxCallbackArgumentsLength uint32
}, error) {
	return _VRFCoordinator.Contract.SCallbackConfig(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) SCoordinatorConfig(opts *bind.CallOpts) (struct {
	UseReasonableGasPrice             bool
	ReentrancyLock                    bool
	Paused                            bool
	PremiumPercentage                 uint8
	UnusedGasPenaltyPercent           uint8
	StalenessSeconds                  uint32
	RedeemableRequestGasOverhead      uint32
	CallbackRequestGasOverhead        uint32
	ReasonableGasPriceStalenessBlocks uint32
	FallbackWeiPerUnitLink            *big.Int
}, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "s_coordinatorConfig")

	outstruct := new(struct {
		UseReasonableGasPrice             bool
		ReentrancyLock                    bool
		Paused                            bool
		PremiumPercentage                 uint8
		UnusedGasPenaltyPercent           uint8
		StalenessSeconds                  uint32
		RedeemableRequestGasOverhead      uint32
		CallbackRequestGasOverhead        uint32
		ReasonableGasPriceStalenessBlocks uint32
		FallbackWeiPerUnitLink            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UseReasonableGasPrice = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ReentrancyLock = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.PremiumPercentage = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.UnusedGasPenaltyPercent = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.StalenessSeconds = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.RedeemableRequestGasOverhead = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.CallbackRequestGasOverhead = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.ReasonableGasPriceStalenessBlocks = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.FallbackWeiPerUnitLink = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFCoordinator *VRFCoordinatorSession) SCoordinatorConfig() (struct {
	UseReasonableGasPrice             bool
	ReentrancyLock                    bool
	Paused                            bool
	PremiumPercentage                 uint8
	UnusedGasPenaltyPercent           uint8
	StalenessSeconds                  uint32
	RedeemableRequestGasOverhead      uint32
	CallbackRequestGasOverhead        uint32
	ReasonableGasPriceStalenessBlocks uint32
	FallbackWeiPerUnitLink            *big.Int
}, error) {
	return _VRFCoordinator.Contract.SCoordinatorConfig(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) SCoordinatorConfig() (struct {
	UseReasonableGasPrice             bool
	ReentrancyLock                    bool
	Paused                            bool
	PremiumPercentage                 uint8
	UnusedGasPenaltyPercent           uint8
	StalenessSeconds                  uint32
	RedeemableRequestGasOverhead      uint32
	CallbackRequestGasOverhead        uint32
	ReasonableGasPriceStalenessBlocks uint32
	FallbackWeiPerUnitLink            *big.Int
}, error) {
	return _VRFCoordinator.Contract.SCoordinatorConfig(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCaller) SPendingRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "s_pendingRequests", arg0)

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

func (_VRFCoordinator *VRFCoordinatorSession) SPendingRequests(arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	return _VRFCoordinator.Contract.SPendingRequests(&_VRFCoordinator.CallOpts, arg0)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) SPendingRequests(arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	return _VRFCoordinator.Contract.SPendingRequests(&_VRFCoordinator.CallOpts, arg0)
}

func (_VRFCoordinator *VRFCoordinatorCaller) SProducer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "s_producer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinator *VRFCoordinatorSession) SProducer() (common.Address, error) {
	return _VRFCoordinator.Contract.SProducer(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorCallerSession) SProducer() (common.Address, error) {
	return _VRFCoordinator.Contract.SProducer(&_VRFCoordinator.CallOpts)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "acceptOwnership")
}

func (_VRFCoordinator *VRFCoordinatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinator.Contract.AcceptOwnership(&_VRFCoordinator.TransactOpts)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinator.Contract.AcceptOwnership(&_VRFCoordinator.TransactOpts)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFCoordinator *VRFCoordinatorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinator.TransactOpts, subId)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinator.TransactOpts, subId)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFCoordinator *VRFCoordinatorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.AddConsumer(&_VRFCoordinator.TransactOpts, subId, consumer)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.AddConsumer(&_VRFCoordinator.TransactOpts, subId, consumer)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) BatchTransferLink(opts *bind.TransactOpts, recipients []common.Address, paymentsInJuels []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "batchTransferLink", recipients, paymentsInJuels)
}

func (_VRFCoordinator *VRFCoordinatorSession) BatchTransferLink(recipients []common.Address, paymentsInJuels []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.BatchTransferLink(&_VRFCoordinator.TransactOpts, recipients, paymentsInJuels)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) BatchTransferLink(recipients []common.Address, paymentsInJuels []*big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.BatchTransferLink(&_VRFCoordinator.TransactOpts, recipients, paymentsInJuels)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFCoordinator *VRFCoordinatorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.CancelSubscription(&_VRFCoordinator.TransactOpts, subId, to)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.CancelSubscription(&_VRFCoordinator.TransactOpts, subId, to)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "createSubscription")
}

func (_VRFCoordinator *VRFCoordinatorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinator.Contract.CreateSubscription(&_VRFCoordinator.TransactOpts)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinator.Contract.CreateSubscription(&_VRFCoordinator.TransactOpts)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) DeregisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "deregisterMigratableCoordinator", target)
}

func (_VRFCoordinator *VRFCoordinatorSession) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.DeregisterMigratableCoordinator(&_VRFCoordinator.TransactOpts, target)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.DeregisterMigratableCoordinator(&_VRFCoordinator.TransactOpts, target)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) Migrate(opts *bind.TransactOpts, newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "migrate", newCoordinator, encodedRequest)
}

func (_VRFCoordinator *VRFCoordinatorSession) Migrate(newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.Migrate(&_VRFCoordinator.TransactOpts, newCoordinator, encodedRequest)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) Migrate(newCoordinator common.Address, encodedRequest []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.Migrate(&_VRFCoordinator.TransactOpts, newCoordinator, encodedRequest)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFCoordinator *VRFCoordinatorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.OnTokenTransfer(&_VRFCoordinator.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.OnTokenTransfer(&_VRFCoordinator.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) ProcessVRFOutputs(opts *bind.TransactOpts, vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "processVRFOutputs", vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight)
}

func (_VRFCoordinator *VRFCoordinatorSession) ProcessVRFOutputs(vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.ProcessVRFOutputs(&_VRFCoordinator.TransactOpts, vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) ProcessVRFOutputs(vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.ProcessVRFOutputs(&_VRFCoordinator.TransactOpts, vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) RedeemRandomness(opts *bind.TransactOpts, subID *big.Int, requestID *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "redeemRandomness", subID, requestID, arg2)
}

func (_VRFCoordinator *VRFCoordinatorSession) RedeemRandomness(subID *big.Int, requestID *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RedeemRandomness(&_VRFCoordinator.TransactOpts, subID, requestID, arg2)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) RedeemRandomness(subID *big.Int, requestID *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RedeemRandomness(&_VRFCoordinator.TransactOpts, subID, requestID, arg2)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "registerMigratableCoordinator", target)
}

func (_VRFCoordinator *VRFCoordinatorSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RegisterMigratableCoordinator(&_VRFCoordinator.TransactOpts, target)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RegisterMigratableCoordinator(&_VRFCoordinator.TransactOpts, target)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFCoordinator *VRFCoordinatorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RemoveConsumer(&_VRFCoordinator.TransactOpts, subId, consumer)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RemoveConsumer(&_VRFCoordinator.TransactOpts, subId, consumer)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) RequestRandomness(opts *bind.TransactOpts, subID *big.Int, numWords uint16, confDelay *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "requestRandomness", subID, numWords, confDelay, arg3)
}

func (_VRFCoordinator *VRFCoordinatorSession) RequestRandomness(subID *big.Int, numWords uint16, confDelay *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestRandomness(&_VRFCoordinator.TransactOpts, subID, numWords, confDelay, arg3)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) RequestRandomness(subID *big.Int, numWords uint16, confDelay *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestRandomness(&_VRFCoordinator.TransactOpts, subID, numWords, confDelay, arg3)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID *big.Int, numWords uint16, confDelay *big.Int, callbackGasLimit uint32, arguments []byte, arg5 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confDelay, callbackGasLimit, arguments, arg5)
}

func (_VRFCoordinator *VRFCoordinatorSession) RequestRandomnessFulfillment(subID *big.Int, numWords uint16, confDelay *big.Int, callbackGasLimit uint32, arguments []byte, arg5 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestRandomnessFulfillment(&_VRFCoordinator.TransactOpts, subID, numWords, confDelay, callbackGasLimit, arguments, arg5)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) RequestRandomnessFulfillment(subID *big.Int, numWords uint16, confDelay *big.Int, callbackGasLimit uint32, arguments []byte, arg5 []byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestRandomnessFulfillment(&_VRFCoordinator.TransactOpts, subID, numWords, confDelay, callbackGasLimit, arguments, arg5)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFCoordinator *VRFCoordinatorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinator.TransactOpts, subId, newOwner)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinator.TransactOpts, subId, newOwner)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) SetCallbackConfig(opts *bind.TransactOpts, config VRFCoordinatorCallbackConfig) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "setCallbackConfig", config)
}

func (_VRFCoordinator *VRFCoordinatorSession) SetCallbackConfig(config VRFCoordinatorCallbackConfig) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetCallbackConfig(&_VRFCoordinator.TransactOpts, config)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) SetCallbackConfig(config VRFCoordinatorCallbackConfig) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetCallbackConfig(&_VRFCoordinator.TransactOpts, config)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) SetConfirmationDelays(opts *bind.TransactOpts, confDelays [8]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "setConfirmationDelays", confDelays)
}

func (_VRFCoordinator *VRFCoordinatorSession) SetConfirmationDelays(confDelays [8]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetConfirmationDelays(&_VRFCoordinator.TransactOpts, confDelays)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) SetConfirmationDelays(confDelays [8]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetConfirmationDelays(&_VRFCoordinator.TransactOpts, confDelays)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) SetCoordinatorConfig(opts *bind.TransactOpts, coordinatorConfig VRFBeaconTypesCoordinatorConfig) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "setCoordinatorConfig", coordinatorConfig)
}

func (_VRFCoordinator *VRFCoordinatorSession) SetCoordinatorConfig(coordinatorConfig VRFBeaconTypesCoordinatorConfig) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetCoordinatorConfig(&_VRFCoordinator.TransactOpts, coordinatorConfig)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) SetCoordinatorConfig(coordinatorConfig VRFBeaconTypesCoordinatorConfig) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetCoordinatorConfig(&_VRFCoordinator.TransactOpts, coordinatorConfig)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) SetPauseFlag(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "setPauseFlag", pause)
}

func (_VRFCoordinator *VRFCoordinatorSession) SetPauseFlag(pause bool) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetPauseFlag(&_VRFCoordinator.TransactOpts, pause)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) SetPauseFlag(pause bool) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetPauseFlag(&_VRFCoordinator.TransactOpts, pause)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) SetProducer(opts *bind.TransactOpts, producer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "setProducer", producer)
}

func (_VRFCoordinator *VRFCoordinatorSession) SetProducer(producer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetProducer(&_VRFCoordinator.TransactOpts, producer)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) SetProducer(producer common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.SetProducer(&_VRFCoordinator.TransactOpts, producer)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) TransferLink(opts *bind.TransactOpts, recipient common.Address, juelsAmount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "transferLink", recipient, juelsAmount)
}

func (_VRFCoordinator *VRFCoordinatorSession) TransferLink(recipient common.Address, juelsAmount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.TransferLink(&_VRFCoordinator.TransactOpts, recipient, juelsAmount)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) TransferLink(recipient common.Address, juelsAmount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.TransferLink(&_VRFCoordinator.TransactOpts, recipient, juelsAmount)
}

func (_VRFCoordinator *VRFCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFCoordinator *VRFCoordinatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.TransferOwnership(&_VRFCoordinator.TransactOpts, to)
}

func (_VRFCoordinator *VRFCoordinatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.TransferOwnership(&_VRFCoordinator.TransactOpts, to)
}

type VRFCoordinatorCallbackConfigSetIterator struct {
	Event *VRFCoordinatorCallbackConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorCallbackConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorCallbackConfigSet)
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
		it.Event = new(VRFCoordinatorCallbackConfigSet)
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

func (it *VRFCoordinatorCallbackConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorCallbackConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorCallbackConfigSet struct {
	NewConfig VRFCoordinatorCallbackConfig
	Raw       types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterCallbackConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorCallbackConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "CallbackConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorCallbackConfigSetIterator{contract: _VRFCoordinator.contract, event: "CallbackConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchCallbackConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorCallbackConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "CallbackConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorCallbackConfigSet)
				if err := _VRFCoordinator.contract.UnpackLog(event, "CallbackConfigSet", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseCallbackConfigSet(log types.Log) (*VRFCoordinatorCallbackConfigSet, error) {
	event := new(VRFCoordinatorCallbackConfigSet)
	if err := _VRFCoordinator.contract.UnpackLog(event, "CallbackConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorCoordinatorConfigSetIterator struct {
	Event *VRFCoordinatorCoordinatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorCoordinatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorCoordinatorConfigSet)
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
		it.Event = new(VRFCoordinatorCoordinatorConfigSet)
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

func (it *VRFCoordinatorCoordinatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorCoordinatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorCoordinatorConfigSet struct {
	CoordinatorConfig VRFBeaconTypesCoordinatorConfig
	Raw               types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterCoordinatorConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorCoordinatorConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "CoordinatorConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorCoordinatorConfigSetIterator{contract: _VRFCoordinator.contract, event: "CoordinatorConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchCoordinatorConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorCoordinatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "CoordinatorConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorCoordinatorConfigSet)
				if err := _VRFCoordinator.contract.UnpackLog(event, "CoordinatorConfigSet", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseCoordinatorConfigSet(log types.Log) (*VRFCoordinatorCoordinatorConfigSet, error) {
	event := new(VRFCoordinatorCoordinatorConfigSet)
	if err := _VRFCoordinator.contract.UnpackLog(event, "CoordinatorConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorCoordinatorDeregisteredIterator struct {
	Event *VRFCoordinatorCoordinatorDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorCoordinatorDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorCoordinatorDeregistered)
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
		it.Event = new(VRFCoordinatorCoordinatorDeregistered)
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

func (it *VRFCoordinatorCoordinatorDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorCoordinatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorCoordinatorDeregistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterCoordinatorDeregistered(opts *bind.FilterOpts) (*VRFCoordinatorCoordinatorDeregisteredIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorCoordinatorDeregisteredIterator{contract: _VRFCoordinator.contract, event: "CoordinatorDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchCoordinatorDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorCoordinatorDeregistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorCoordinatorDeregistered)
				if err := _VRFCoordinator.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseCoordinatorDeregistered(log types.Log) (*VRFCoordinatorCoordinatorDeregistered, error) {
	event := new(VRFCoordinatorCoordinatorDeregistered)
	if err := _VRFCoordinator.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorCoordinatorRegisteredIterator struct {
	Event *VRFCoordinatorCoordinatorRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorCoordinatorRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorCoordinatorRegistered)
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
		it.Event = new(VRFCoordinatorCoordinatorRegistered)
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

func (it *VRFCoordinatorCoordinatorRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorCoordinatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorCoordinatorRegistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VRFCoordinatorCoordinatorRegisteredIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorCoordinatorRegisteredIterator{contract: _VRFCoordinator.contract, event: "CoordinatorRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorCoordinatorRegistered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorCoordinatorRegistered)
				if err := _VRFCoordinator.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseCoordinatorRegistered(log types.Log) (*VRFCoordinatorCoordinatorRegistered, error) {
	event := new(VRFCoordinatorCoordinatorRegistered)
	if err := _VRFCoordinator.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorMigrationCompletedIterator struct {
	Event *VRFCoordinatorMigrationCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorMigrationCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorMigrationCompleted)
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
		it.Event = new(VRFCoordinatorMigrationCompleted)
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

func (it *VRFCoordinatorMigrationCompletedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorMigrationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorMigrationCompleted struct {
	NewVersion     uint8
	NewCoordinator common.Address
	SubID          *big.Int
	Raw            types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterMigrationCompleted(opts *bind.FilterOpts, newVersion []uint8, subID []*big.Int) (*VRFCoordinatorMigrationCompletedIterator, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "MigrationCompleted", newVersionRule, subIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorMigrationCompletedIterator{contract: _VRFCoordinator.contract, event: "MigrationCompleted", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorMigrationCompleted, newVersion []uint8, subID []*big.Int) (event.Subscription, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	var subIDRule []interface{}
	for _, subIDItem := range subID {
		subIDRule = append(subIDRule, subIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "MigrationCompleted", newVersionRule, subIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorMigrationCompleted)
				if err := _VRFCoordinator.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseMigrationCompleted(log types.Log) (*VRFCoordinatorMigrationCompleted, error) {
	event := new(VRFCoordinatorMigrationCompleted)
	if err := _VRFCoordinator.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorOutputsServedIterator struct {
	Event *VRFCoordinatorOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorOutputsServed)
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
		it.Event = new(VRFCoordinatorOutputsServed)
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

func (it *VRFCoordinatorOutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorOutputsServed struct {
	RecentBlockHeight  uint64
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFCoordinatorOutputsServedIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorOutputsServedIterator{contract: _VRFCoordinator.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorOutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorOutputsServed)
				if err := _VRFCoordinator.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseOutputsServed(log types.Log) (*VRFCoordinatorOutputsServed, error) {
	event := new(VRFCoordinatorOutputsServed)
	if err := _VRFCoordinator.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorOwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorOwnershipTransferRequested)
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
		it.Event = new(VRFCoordinatorOwnershipTransferRequested)
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

func (it *VRFCoordinatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorOwnershipTransferRequestedIterator{contract: _VRFCoordinator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorOwnershipTransferRequested)
				if err := _VRFCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorOwnershipTransferRequested, error) {
	event := new(VRFCoordinatorOwnershipTransferRequested)
	if err := _VRFCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorOwnershipTransferredIterator struct {
	Event *VRFCoordinatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorOwnershipTransferred)
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
		it.Event = new(VRFCoordinatorOwnershipTransferred)
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

func (it *VRFCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorOwnershipTransferredIterator{contract: _VRFCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorOwnershipTransferred)
				if err := _VRFCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorOwnershipTransferred, error) {
	event := new(VRFCoordinatorOwnershipTransferred)
	if err := _VRFCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorPauseFlagChangedIterator struct {
	Event *VRFCoordinatorPauseFlagChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorPauseFlagChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorPauseFlagChanged)
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
		it.Event = new(VRFCoordinatorPauseFlagChanged)
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

func (it *VRFCoordinatorPauseFlagChangedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorPauseFlagChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorPauseFlagChanged struct {
	Paused bool
	Raw    types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterPauseFlagChanged(opts *bind.FilterOpts) (*VRFCoordinatorPauseFlagChangedIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "PauseFlagChanged")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorPauseFlagChangedIterator{contract: _VRFCoordinator.contract, event: "PauseFlagChanged", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchPauseFlagChanged(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorPauseFlagChanged) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "PauseFlagChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorPauseFlagChanged)
				if err := _VRFCoordinator.contract.UnpackLog(event, "PauseFlagChanged", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParsePauseFlagChanged(log types.Log) (*VRFCoordinatorPauseFlagChanged, error) {
	event := new(VRFCoordinatorPauseFlagChanged)
	if err := _VRFCoordinator.contract.UnpackLog(event, "PauseFlagChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorRandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorRandomWordsFulfilled)
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
		it.Event = new(VRFCoordinatorRandomWordsFulfilled)
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

func (it *VRFCoordinatorRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	SubBalances           []*big.Int
	SubIDs                []*big.Int
	Raw                   types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFCoordinatorRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorRandomWordsFulfilledIterator{contract: _VRFCoordinator.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorRandomWordsFulfilled)
				if err := _VRFCoordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorRandomWordsFulfilled, error) {
	event := new(VRFCoordinatorRandomWordsFulfilled)
	if err := _VRFCoordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorRandomnessFulfillmentRequestedIterator struct {
	Event *VRFCoordinatorRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorRandomnessFulfillmentRequested)
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
		it.Event = new(VRFCoordinatorRandomnessFulfillmentRequested)
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

func (it *VRFCoordinatorRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorRandomnessFulfillmentRequested struct {
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
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFCoordinatorRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorRandomnessFulfillmentRequestedIterator{contract: _VRFCoordinator.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorRandomnessFulfillmentRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorRandomnessFulfillmentRequested)
				if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFCoordinatorRandomnessFulfillmentRequested, error) {
	event := new(VRFCoordinatorRandomnessFulfillmentRequested)
	if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorRandomnessRedeemedIterator struct {
	Event *VRFCoordinatorRandomnessRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorRandomnessRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorRandomnessRedeemed)
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
		it.Event = new(VRFCoordinatorRandomnessRedeemed)
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

func (it *VRFCoordinatorRandomnessRedeemedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorRandomnessRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorRandomnessRedeemed struct {
	RequestID *big.Int
	Requester common.Address
	SubID     *big.Int
	Raw       types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFCoordinatorRandomnessRedeemedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorRandomnessRedeemedIterator{contract: _VRFCoordinator.contract, event: "RandomnessRedeemed", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorRandomnessRedeemed)
				if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseRandomnessRedeemed(log types.Log) (*VRFCoordinatorRandomnessRedeemed, error) {
	event := new(VRFCoordinatorRandomnessRedeemed)
	if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorRandomnessRequestedIterator struct {
	Event *VRFCoordinatorRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorRandomnessRequested)
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
		it.Event = new(VRFCoordinatorRandomnessRequested)
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

func (it *VRFCoordinatorRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	CostJuels              *big.Int
	NewSubBalance          *big.Int
	Raw                    types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int) (*VRFCoordinatorRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorRandomnessRequestedIterator{contract: _VRFCoordinator.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorRandomnessRequested, requestID []*big.Int) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorRandomnessRequested)
				if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseRandomnessRequested(log types.Log) (*VRFCoordinatorRandomnessRequested, error) {
	event := new(VRFCoordinatorRandomnessRequested)
	if err := _VRFCoordinator.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorSubscriptionCanceledIterator struct {
	Event *VRFCoordinatorSubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorSubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorSubscriptionCanceled)
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
		it.Event = new(VRFCoordinatorSubscriptionCanceled)
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

func (it *VRFCoordinatorSubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorSubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorSubscriptionCanceled struct {
	SubId  *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorSubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorSubscriptionCanceledIterator{contract: _VRFCoordinator.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorSubscriptionCanceled, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorSubscriptionCanceled)
				if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorSubscriptionCanceled, error) {
	event := new(VRFCoordinatorSubscriptionCanceled)
	if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorSubscriptionConsumerAddedIterator struct {
	Event *VRFCoordinatorSubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorSubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorSubscriptionConsumerAdded)
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
		it.Event = new(VRFCoordinatorSubscriptionConsumerAdded)
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

func (it *VRFCoordinatorSubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorSubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorSubscriptionConsumerAdded struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorSubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorSubscriptionConsumerAddedIterator{contract: _VRFCoordinator.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorSubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorSubscriptionConsumerAdded)
				if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorSubscriptionConsumerAdded, error) {
	event := new(VRFCoordinatorSubscriptionConsumerAdded)
	if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorSubscriptionConsumerRemovedIterator struct {
	Event *VRFCoordinatorSubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorSubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorSubscriptionConsumerRemoved)
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
		it.Event = new(VRFCoordinatorSubscriptionConsumerRemoved)
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

func (it *VRFCoordinatorSubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorSubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorSubscriptionConsumerRemoved struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorSubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorSubscriptionConsumerRemovedIterator{contract: _VRFCoordinator.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorSubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorSubscriptionConsumerRemoved)
				if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorSubscriptionConsumerRemoved, error) {
	event := new(VRFCoordinatorSubscriptionConsumerRemoved)
	if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorSubscriptionCreatedIterator struct {
	Event *VRFCoordinatorSubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorSubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorSubscriptionCreated)
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
		it.Event = new(VRFCoordinatorSubscriptionCreated)
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

func (it *VRFCoordinatorSubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorSubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorSubscriptionCreated struct {
	SubId *big.Int
	Owner common.Address
	Raw   types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int, owner []common.Address) (*VRFCoordinatorSubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorSubscriptionCreatedIterator{contract: _VRFCoordinator.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorSubscriptionCreated, subId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorSubscriptionCreated)
				if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorSubscriptionCreated, error) {
	event := new(VRFCoordinatorSubscriptionCreated)
	if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorSubscriptionFundedIterator struct {
	Event *VRFCoordinatorSubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorSubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorSubscriptionFunded)
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
		it.Event = new(VRFCoordinatorSubscriptionFunded)
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

func (it *VRFCoordinatorSubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorSubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorSubscriptionFunded struct {
	SubId      *big.Int
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorSubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorSubscriptionFundedIterator{contract: _VRFCoordinator.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorSubscriptionFunded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorSubscriptionFunded)
				if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorSubscriptionFunded, error) {
	event := new(VRFCoordinatorSubscriptionFunded)
	if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorSubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFCoordinatorSubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorSubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorSubscriptionOwnerTransferRequested)
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
		it.Event = new(VRFCoordinatorSubscriptionOwnerTransferRequested)
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

func (it *VRFCoordinatorSubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorSubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorSubscriptionOwnerTransferRequested struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorSubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorSubscriptionOwnerTransferRequestedIterator{contract: _VRFCoordinator.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorSubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorSubscriptionOwnerTransferRequested)
				if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorSubscriptionOwnerTransferRequested, error) {
	event := new(VRFCoordinatorSubscriptionOwnerTransferRequested)
	if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorSubscriptionOwnerTransferredIterator struct {
	Event *VRFCoordinatorSubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorSubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorSubscriptionOwnerTransferred)
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
		it.Event = new(VRFCoordinatorSubscriptionOwnerTransferred)
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

func (it *VRFCoordinatorSubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorSubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorSubscriptionOwnerTransferred struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinator *VRFCoordinatorFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VRFCoordinatorSubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorSubscriptionOwnerTransferredIterator{contract: _VRFCoordinator.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinator *VRFCoordinatorFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorSubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorSubscriptionOwnerTransferred)
				if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
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

func (_VRFCoordinator *VRFCoordinatorFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorSubscriptionOwnerTransferred, error) {
	event := new(VRFCoordinatorSubscriptionOwnerTransferred)
	if err := _VRFCoordinator.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

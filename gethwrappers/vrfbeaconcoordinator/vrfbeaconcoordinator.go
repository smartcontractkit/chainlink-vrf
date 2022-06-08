package vrfbeaconcoordinator

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

type KeyDataStructKeyData struct {
	PublicKey []byte
	Hashes    [][32]byte
}

type VRFBeaconReportHotVars struct {
	F                         uint8
	LatestEpochAndRound       *big.Int
	LatestAggregatorRoundId   uint32
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}

type VRFBeaconReportOutputServed struct {
	Height            uint64
	ConfirmationDelay *big.Int
}

type VRFBeaconReportReport struct {
	Outputs           []VRFBeaconReportVRFOutput
	JuelsPerFeeCoin   *big.Int
	RecentBlockHeight uint64
	RecentBlockHash   [32]byte
}

type VRFBeaconReportVRFOutput struct {
	BlockHeight       uint64
	ConfirmationDelay *big.Int
	VrfOutput         ECCArithmeticG1Point
	Callbacks         []VRFBeaconTypesCostedCallback
}

type VRFBeaconTypesCallback struct {
	RequestID    *big.Int
	NumWords     uint16
	Requester    common.Address
	Arguments    []byte
	SubID        uint64
	GasAllowance *big.Int
}

type VRFBeaconTypesCostedCallback struct {
	Callback VRFBeaconTypesCallback
	Price    *big.Int
}

var AccessControllerInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6b14daf8": "hasAccess(address,bytes)",
	},
}

var AccessControllerInterfaceABI = AccessControllerInterfaceMetaData.ABI

var AccessControllerInterfaceFuncSigs = AccessControllerInterfaceMetaData.Sigs

type AccessControllerInterface struct {
	AccessControllerInterfaceCaller
	AccessControllerInterfaceTransactor
	AccessControllerInterfaceFilterer
}

type AccessControllerInterfaceCaller struct {
	contract *bind.BoundContract
}

type AccessControllerInterfaceTransactor struct {
	contract *bind.BoundContract
}

type AccessControllerInterfaceFilterer struct {
	contract *bind.BoundContract
}

type AccessControllerInterfaceSession struct {
	Contract     *AccessControllerInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AccessControllerInterfaceCallerSession struct {
	Contract *AccessControllerInterfaceCaller
	CallOpts bind.CallOpts
}

type AccessControllerInterfaceTransactorSession struct {
	Contract     *AccessControllerInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type AccessControllerInterfaceRaw struct {
	Contract *AccessControllerInterface
}

type AccessControllerInterfaceCallerRaw struct {
	Contract *AccessControllerInterfaceCaller
}

type AccessControllerInterfaceTransactorRaw struct {
	Contract *AccessControllerInterfaceTransactor
}

func NewAccessControllerInterface(address common.Address, backend bind.ContractBackend) (*AccessControllerInterface, error) {
	contract, err := bindAccessControllerInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControllerInterface{AccessControllerInterfaceCaller: AccessControllerInterfaceCaller{contract: contract}, AccessControllerInterfaceTransactor: AccessControllerInterfaceTransactor{contract: contract}, AccessControllerInterfaceFilterer: AccessControllerInterfaceFilterer{contract: contract}}, nil
}

func NewAccessControllerInterfaceCaller(address common.Address, caller bind.ContractCaller) (*AccessControllerInterfaceCaller, error) {
	contract, err := bindAccessControllerInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControllerInterfaceCaller{contract: contract}, nil
}

func NewAccessControllerInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControllerInterfaceTransactor, error) {
	contract, err := bindAccessControllerInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControllerInterfaceTransactor{contract: contract}, nil
}

func NewAccessControllerInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControllerInterfaceFilterer, error) {
	contract, err := bindAccessControllerInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControllerInterfaceFilterer{contract: contract}, nil
}

func bindAccessControllerInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControllerInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AccessControllerInterface *AccessControllerInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControllerInterface.Contract.AccessControllerInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_AccessControllerInterface *AccessControllerInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControllerInterface.Contract.AccessControllerInterfaceTransactor.contract.Transfer(opts)
}

func (_AccessControllerInterface *AccessControllerInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControllerInterface.Contract.AccessControllerInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_AccessControllerInterface *AccessControllerInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControllerInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_AccessControllerInterface *AccessControllerInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControllerInterface.Contract.contract.Transfer(opts)
}

func (_AccessControllerInterface *AccessControllerInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControllerInterface.Contract.contract.Transact(opts, method, params...)
}

func (_AccessControllerInterface *AccessControllerInterfaceCaller) HasAccess(opts *bind.CallOpts, user common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _AccessControllerInterface.contract.Call(opts, &out, "hasAccess", user, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AccessControllerInterface *AccessControllerInterfaceSession) HasAccess(user common.Address, data []byte) (bool, error) {
	return _AccessControllerInterface.Contract.HasAccess(&_AccessControllerInterface.CallOpts, user, data)
}

func (_AccessControllerInterface *AccessControllerInterfaceCallerSession) HasAccess(user common.Address, data []byte) (bool, error) {
	return _AccessControllerInterface.Contract.HasAccess(&_AccessControllerInterface.CallOpts, user, data)
}

var ConfirmedOwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161045638038061045683398101604081905261002f9161016e565b8060006001600160a01b03821661008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100bd576100bd816100c5565b50505061019e565b336001600160a01b0382160361011d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561018057600080fd5b81516001600160a01b038116811461019757600080fd5b9392505050565b6102a9806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d366004610243565b610131565b6001546001600160a01b031633146100da5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610139610145565b6101428161019a565b50565b6000546001600160a01b031633146101985760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b60448201526064016100d1565b565b336001600160a01b038216036101f25760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d1565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561025557600080fd5b81356001600160a01b038116811461026c57600080fd5b939250505056fea2646970667358221220b4da426ae50c83c1e19e64812d16826708de27ebc2ebbaf27675229b792b452064736f6c634300080d0033",
}

var ConfirmedOwnerABI = ConfirmedOwnerMetaData.ABI

var ConfirmedOwnerFuncSigs = ConfirmedOwnerMetaData.Sigs

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
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161047138038061047183398101604081905261002f91610186565b6001600160a01b03821661008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100ba576100ba816100c1565b50506101b9565b336001600160a01b038216036101195760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610081565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b038116811461018157600080fd5b919050565b6000806040838503121561019957600080fd5b6101a28361016a565b91506101b06020840161016a565b90509250929050565b6102a9806101c86000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d366004610243565b610131565b6001546001600160a01b031633146100da5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610139610145565b6101428161019a565b50565b6000546001600160a01b031633146101985760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b60448201526064016100d1565b565b336001600160a01b038216036101f25760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d1565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561025557600080fd5b81356001600160a01b038116811461026c57600080fd5b939250505056fea264697066735822122019ccdf4d2d94f8d033673c2e6491c10a72ef7626a6472b0e22c7a6004551d3d864736f6c634300080d0033",
}

var ConfirmedOwnerWithProposalABI = ConfirmedOwnerWithProposalMetaData.ABI

var ConfirmedOwnerWithProposalFuncSigs = ConfirmedOwnerWithProposalMetaData.Sigs

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractDKGClient\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"DKGClientError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"indexed\":false,\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"KeyGenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"addClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_configDigest\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"removeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"7bf1ffc5": "addClient(bytes32,address)",
		"9201de55": "bytes32ToString(bytes32)",
		"39614e4f": "bytesToString(bytes)",
		"c3105a6b": "getKey(bytes32,bytes32)",
		"81ff7048": "latestConfigDetails()",
		"afcb95d7": "latestConfigDigestAndEpoch()",
		"8da5cb5b": "owner()",
		"5429a79e": "removeClient(bytes32,address)",
		"e3d0e712": "setConfig(address[],address[],uint8,bytes,uint64,bytes)",
		"0bc643e8": "toASCII(uint8)",
		"f2fde38b": "transferOwnership(address)",
		"b1dc65a4": "transmit(bytes32[3],bytes,bytes32[],bytes32[],bytes32)",
		"181f5a77": "typeAndVersion()",
	},
	Bin: "0x60806040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000a5565b50505062000150565b336001600160a01b03821603620000ff5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61276480620001606000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063b1dc65a411610066578063b1dc65a41461020f578063c3105a6b14610222578063e3d0e71214610242578063f2fde38b1461025557600080fd5b80638da5cb5b146101b75780639201de55146101d2578063afcb95d7146101e557600080fd5b80635429a79e116100c85780635429a79e1461015a57806379ba50971461016f5780637bf1ffc51461017757806381ff70481461018a57600080fd5b80630bc643e8146100ef578063181f5a771461011957806339614e4f14610147575b600080fd5b6101026100fd366004611dad565b610268565b60405160ff90911681526020015b60405180910390f35b604080518082019091526009815268444b4720302e302e3160b81b60208201525b6040516101109190611e24565b61013a610155366004611efa565b610297565b61016d610168366004611f4b565b610400565b005b61016d61063f565b61016d610185366004611f4b565b6106ee565b600754600554604080516000815264010000000090930463ffffffff166020840152820152606001610110565b6000546040516001600160a01b039091168152602001610110565b61013a6101e0366004611f7b565b610735565b6005546004546040805160008152602081019390935263ffffffff90911690820152606001610110565b61016d61021d366004611fdf565b6107c1565b6102356102303660046120c3565b610905565b60405161011091906120e5565b61016d6102503660046121eb565b610a2d565b61016d6102633660046122b7565b611177565b6000600a8260ff161015610287576102818260306122ea565b92915050565b6102818260576122ea565b919050565b6060600080835160026102aa919061230f565b6001600160401b038111156102c1576102c1611e37565b6040519080825280601f01601f1916602001820160405280156102eb576020820181803683370190505b509050600091505b80518260ff1610156103f95760008461030d60028561232e565b60ff16815181106103205761032061235e565b60209101015160f81c600f169050600060048661033e60028761232e565b60ff16815181106103515761035161235e565b01602001516001600160f81b031916901c60f81c905061037081610268565b60f81b838560ff16815181106103885761038861235e565b60200101906001600160f81b031916908160001a9053506103aa8460016122ea565b93506103b582610268565b60f81b838560ff16815181106103cd576103cd61235e565b60200101906001600160f81b031916908160001a905350505081806103f190612374565b9250506102f3565b9392505050565b61040861118b565b60008281526002602090815260408083208054825181850281018501909352808352919290919083018282801561046857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161044a575b50505050509050600081516001600160401b0381111561048a5761048a611e37565b6040519080825280602002602001820160405280156104b3578160200160208202803683370190505b5090506000805b835181101561055657846001600160a01b03168482815181106104df576104df61235e565b60200260200101516001600160a01b0316146105365784836105018484612393565b815181106105115761051161235e565b60200260200101906001600160a01b031690816001600160a01b031681525050610544565b81610540816123aa565b9250505b8061054e816123aa565b9150506104ba565b5060008184516105669190612393565b6001600160401b0381111561057d5761057d611e37565b6040519080825280602002602001820160405280156105a6578160200160208202803683370190505b50905060005b8285516105b99190612393565b811015610616578381815181106105d2576105d261235e565b60200260200101518282815181106105ec576105ec61235e565b6001600160a01b03909216602092830291909101909101528061060e816123aa565b9150506105ac565b506000868152600260209081526040909120825161063692840190611c55565b50505050505050565b6001546001600160a01b031633146106975760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6106f661118b565b600091825260026020908152604083208054600181018255908452922090910180546001600160a01b0319166001600160a01b03909216919091179055565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156107b7578381602081106107785761077861235e565b1a60f81b82828151811061078e5761078e61235e565b60200101906001600160f81b031916908160001a905350806107af816123aa565b91505061075d565b506103f981610297565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916108119184918491908e908e90819084018382808284376000920191909152506111e092505050565b6040805183815263ffffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260055480825260065460ff808216602085015261010090910416928201929092529083146108c55760405162461bcd60e51b81526020600482015260156024820152740c6dedcccd2ce88d2cecae6e840dad2e6dac2e8c6d605b1b604482015260640161068e565b6108d38b8b8b8b8b8b61143d565b6108e48c8c8c8c8c8c8c8c896114d1565b50505063ffffffff81106108fa576108fa6123c3565b505050505050505050565b6040805180820190915260608082526020820152600083815260036020908152604080832085845290915290819020815180830190925280548290829061094b906123d9565b80601f0160208091040260200160405190810160405280929190818152602001828054610977906123d9565b80156109c45780601f10610999576101008083540402835291602001916109c4565b820191906000526020600020905b8154815290600101906020018083116109a757829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610a1c57602002820191906000526020600020905b815481526020019060010190808311610a08575b505050505081525050905092915050565b855185518560ff16601f831115610a795760405162461bcd60e51b815260206004820152601060248201526f746f6f206d616e79207369676e65727360801b604482015260640161068e565b60008111610abe5760405162461bcd60e51b815260206004820152601260248201527166206d75737420626520706f73697469766560701b604482015260640161068e565b818314610b195760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f662072656769737472616044820152633a34b7b760e11b606482015260840161068e565b610b2481600361230f565b8311610b725760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161068e565b610b7a61118b565b6040805160c0810182528a8152602081018a905260ff891691810191909152606081018790526001600160401b038616608082015260a081018590525b60095415610cc957600954600090610bd190600190612393565b9050600060098281548110610be857610be861235e565b6000918252602082200154600a80546001600160a01b0390921693509084908110610c1557610c1561235e565b60009182526020808320909101546001600160a01b03858116845260089092526040808420805461ffff1990811690915592909116808452922080549091169055600980549192509080610c6b57610c6b612413565b600082815260209020810160001990810180546001600160a01b0319169055019055600a805480610c9e57610c9e612413565b600082815260209020810160001990810180546001600160a01b031916905501905550610bb7915050565b60005b8151518110156110085760006008600084600001518481518110610cf257610cf261235e565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610d2f57610d2f612429565b14610d7c5760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161068e565b6040805180820190915260ff82168152600160208201528251805160089160009185908110610dad57610dad61235e565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610e0657610e06612429565b021790555060009150610e169050565b6008600084602001518481518110610e3057610e3061235e565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610e6d57610e6d612429565b14610eba5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161068e565b6040805180820190915260ff821681526020810160028152506008600084602001518481518110610eed57610eed61235e565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610f4657610f46612429565b021790555050825180516009925083908110610f6457610f6461235e565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b03909316929092179091558201518051600a919083908110610fbb57610fbb61235e565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b0390921691909117905580611000816123aa565b915050610ccc565b5060408101516006805460ff191660ff9092169190911790556007805467ffffffff0000000019811664010000000063ffffffff43811682029283178555908304811693600193909260009261106592869290821691161761243f565b92506101000a81548163ffffffff021916908363ffffffff16021790555060006110c64630600760009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a00151611957565b6005819055835180516006805460ff9092166101000261ff00199092169190911790556007546020860151604080880151606089015160808a015160a08b015193519798507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e059761114e978b978b9763ffffffff9091169691959094909390929091906124ab565b60405180910390a161116983604001518460600151836119b2565b505050505050505050505050565b61117f61118b565b61118881611bac565b50565b6000546001600160a01b031633146111de5760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015260640161068e565b565b6000606080838060200190518101906111f99190612540565b60408051808201825283815260208082018490526000868152600282528381208054855181850281018501909652808652979a509598509396509094929391929083018282801561127357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611255575b5050505050905060005b8151811015611390578181815181106112985761129861235e565b60200260200101516001600160a01b031663bf2732c7846040518263ffffffff1660e01b81526004016112cb91906120e5565b600060405180830381600087803b1580156112e557600080fd5b505af19250505080156112f6575060015b61137e573d808015611324576040519150601f19603f3d011682016040523d82523d6000602084013e611329565b606091505b507f116391732f5df106193bda7cedf1728f3b07b62f6cdcdd611c9eeec44efcae5483838151811061135d5761135d61235e565b60200260200101518260405161137492919061263d565b60405180910390a1505b80611388816123aa565b91505061127d565b5060008581526003602090815260408083208b845282529091208351805185936113be928492910190611cba565b5060208281015180516113d79260018501920190611d2e565b5090505084887fc8db841f5b2231ccf7190311f440aa197b161e369f3b40b023508160cc5556568460405161140c91906120e5565b60405180910390a350506004805460089690961c63ffffffff1663ffffffff19909616959095179094555050505050565b600061144a82602061230f565b61145585602061230f565b61146188610144612661565b61146b9190612661565b6114759190612661565b611480906000612661565b90503681146106365760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604482015260640161068e565b60006002826020015183604001516114e991906122ea565b6114f3919061232e565b6114fe9060016122ea565b60408051600180825281830190925260ff929092169250600091906020820181803683370190505090508160f81b8160008151811061153f5761153f61235e565b60200101906001600160f81b031916908160001a90535086821461156282610297565b906115805760405162461bcd60e51b815260040161068e9190611e24565b508685146115d05760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161068e565b3360009081526008602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561161357611613612429565b600281111561162457611624612429565b905250905060028160200151600281111561164157611641612429565b14801561167b5750600a816000015160ff16815481106116635761166361235e565b6000918252602090912001546001600160a01b031633145b6116c75760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161068e565b505050600088886040516116dc929190612679565b6040519081900381206116f3918c90602001612689565b604051602081830303815290604052805190602001209050611713611d68565b604080518082019091526000808252602082015260005b888110156119485760006001858884602081106117495761174961235e565b61175691901a601b6122ea565b8d8d868181106117685761176861235e565b905060200201358c8c878181106117815761178161235e565b90506020020135604051600081526020016040526040516117be949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156117e0573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526008602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561183557611835612429565b600281111561184657611846612429565b905250925060018360200151600281111561186357611863612429565b146118b05760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161068e565b8251849060ff16601f81106118c7576118c761235e565b6020020151156119105760405162461bcd60e51b81526020600482015260146024820152736e6f6e2d756e69717565207369676e617475726560601b604482015260640161068e565b600184846000015160ff16601f811061192b5761192b61235e565b911515602090920201525080611940816123aa565b91505061172a565b50505050505050505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161197b999897969594939291906126a5565b60408051601f1981840301815291905280516020909101206001600160f01b0316600160f01b179150509998505050505050505050565b6000808351602014611a065760405162461bcd60e51b815260206004820152601e60248201527f77726f6e67206c656e67746820666f72206f6e636861696e436f6e6669670000604482015260640161068e565b60208401519150808203611a535760405162461bcd60e51b815260206004820152601460248201527319985a5b1959081d1bc818dbdc1e481ad95e525160621b604482015260640161068e565b604080518082019091526060808252602082015260008381526003602090815260408083208784528252909120825180518493611a94928492910190611cba565b506020828101518051611aad9260018501920190611d2e565b505050600083815260026020908152604080832080548251818502810185019093528083529192909190830182828015611b1057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611af2575b5050505050905060005b8151811015611ba257818181518110611b3557611b3561235e565b60200260200101516001600160a01b03166355e487496040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611b7757600080fd5b505af1158015611b8b573d6000803e3d6000fd5b505050508080611b9a906123aa565b915050611b1a565b5050505050505050565b336001600160a01b03821603611c045760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161068e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215611caa579160200282015b82811115611caa57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611c75565b50611cb6929150611d87565b5090565b828054611cc6906123d9565b90600052602060002090601f016020900481019282611ce85760008555611caa565b82601f10611d0157805160ff1916838001178555611caa565b82800160010185558215611caa579182015b82811115611caa578251825591602001919060010190611d13565b828054828255906000526020600020908101928215611caa5791602002820182811115611caa578251825591602001919060010190611d13565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115611cb65760008155600101611d88565b803560ff8116811461029257600080fd5b600060208284031215611dbf57600080fd5b6103f982611d9c565b60005b83811015611de3578181015183820152602001611dcb565b83811115611df2576000848401525b50505050565b60008151808452611e10816020860160208601611dc8565b601f01601f19169290920160200192915050565b6020815260006103f96020830184611df8565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715611e7557611e75611e37565b604052919050565b60006001600160401b03821115611e9657611e96611e37565b50601f01601f191660200190565b600082601f830112611eb557600080fd5b8135611ec8611ec382611e7d565b611e4d565b818152846020838601011115611edd57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611f0c57600080fd5b81356001600160401b03811115611f2257600080fd5b611f2e84828501611ea4565b949350505050565b6001600160a01b038116811461118857600080fd5b60008060408385031215611f5e57600080fd5b823591506020830135611f7081611f36565b809150509250929050565b600060208284031215611f8d57600080fd5b5035919050565b60008083601f840112611fa657600080fd5b5081356001600160401b03811115611fbd57600080fd5b6020830191508360208260051b8501011115611fd857600080fd5b9250929050565b60008060008060008060008060e0898b031215611ffb57600080fd5b606089018a81111561200c57600080fd5b899850356001600160401b038082111561202557600080fd5b818b0191508b601f83011261203957600080fd5b81358181111561204857600080fd5b8c602082850101111561205a57600080fd5b6020830199508098505060808b013591508082111561207857600080fd5b6120848c838d01611f94565b909750955060a08b013591508082111561209d57600080fd5b506120aa8b828c01611f94565b999c989b50969995989497949560c00135949350505050565b600080604083850312156120d657600080fd5b50508035926020909101359150565b6000602080835283516040828501526121016060850182611df8565b85830151858203601f19016040870152805180835290840192506000918401905b808310156121425783518252928401926001929092019190840190612122565b509695505050505050565b60006001600160401b0382111561216657612166611e37565b5060051b60200190565b600082601f83011261218157600080fd5b81356020612191611ec38361214d565b82815260059290921b840181019181810190868411156121b057600080fd5b8286015b848110156121425780356121c781611f36565b83529183019183016121b4565b80356001600160401b038116811461029257600080fd5b60008060008060008060c0878903121561220457600080fd5b86356001600160401b038082111561221b57600080fd5b6122278a838b01612170565b9750602089013591508082111561223d57600080fd5b6122498a838b01612170565b965061225760408a01611d9c565b9550606089013591508082111561226d57600080fd5b6122798a838b01611ea4565b945061228760808a016121d4565b935060a089013591508082111561229d57600080fd5b506122aa89828a01611ea4565b9150509295509295509295565b6000602082840312156122c957600080fd5b81356103f981611f36565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff03821115612307576123076122d4565b019392505050565b6000816000190483118215151615612329576123296122d4565b500290565b600060ff83168061234f57634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff810361238a5761238a6122d4565b60010192915050565b6000828210156123a5576123a56122d4565b500390565b6000600182016123bc576123bc6122d4565b5060010190565b634e487b7160e01b600052600160045260246000fd5b600181811c908216806123ed57607f821691505b60208210810361240d57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600063ffffffff80831681851680830382111561245e5761245e6122d4565b01949350505050565b600081518084526020808501945080840160005b838110156124a05781516001600160a01b03168752958201959082019060010161247b565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526124db8184018a612467565b905082810360808401526124ef8189612467565b905060ff871660a084015282810360c084015261250c8187611df8565b90506001600160401b03851660e08401528281036101008401526125308185611df8565b9c9b505050505050505050505050565b60008060006060848603121561255557600080fd5b835192506020808501516001600160401b038082111561257457600080fd5b818701915087601f83011261258857600080fd5b8151612596611ec382611e7d565b81815289858386010111156125aa57600080fd5b6125b982868301878701611dc8565b6040890151909650925050808211156125d157600080fd5b508501601f810187136125e357600080fd5b80516125f1611ec38261214d565b81815260059190911b8201830190838101908983111561261057600080fd5b928401925b8284101561262e57835182529284019290840190612615565b80955050505050509250925092565b6001600160a01b0383168152604060208201819052600090611f2e90830184611df8565b60008219821115612674576126746122d4565b500190565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b8981526001600160a01b03891660208201526001600160401b038881166040830152610120606083018190526000916126e08483018b612467565b915083820360808501526126f4828a612467565b915060ff881660a085015283820360c08501526127118288611df8565b90861660e085015283810361010085015290506125308185611df856fea26469706673582212204332a72e5be8559471af0760b7b6870e79270be418de72c17ec48948498eec3064736f6c634300080d0033",
}

var DKGABI = DKGMetaData.ABI

var DKGFuncSigs = DKGMetaData.Sigs

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

func (_DKG *DKGCaller) Bytes32ToString(opts *bind.CallOpts, s [32]byte) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "bytes32ToString", s)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_DKG *DKGSession) Bytes32ToString(s [32]byte) (string, error) {
	return _DKG.Contract.Bytes32ToString(&_DKG.CallOpts, s)
}

func (_DKG *DKGCallerSession) Bytes32ToString(s [32]byte) (string, error) {
	return _DKG.Contract.Bytes32ToString(&_DKG.CallOpts, s)
}

func (_DKG *DKGCaller) BytesToString(opts *bind.CallOpts, _bytes []byte) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "bytesToString", _bytes)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_DKG *DKGSession) BytesToString(_bytes []byte) (string, error) {
	return _DKG.Contract.BytesToString(&_DKG.CallOpts, _bytes)
}

func (_DKG *DKGCallerSession) BytesToString(_bytes []byte) (string, error) {
	return _DKG.Contract.BytesToString(&_DKG.CallOpts, _bytes)
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

func (_DKG *DKGCaller) ToASCII(opts *bind.CallOpts, _uint8 uint8) (uint8, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "toASCII", _uint8)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_DKG *DKGSession) ToASCII(_uint8 uint8) (uint8, error) {
	return _DKG.Contract.ToASCII(&_DKG.CallOpts, _uint8)
}

func (_DKG *DKGCallerSession) ToASCII(_uint8 uint8) (uint8, error) {
	return _DKG.Contract.ToASCII(&_DKG.CallOpts, _uint8)
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
	Sigs: map[string]string{
		"bf2732c7": "keyGenerated((bytes,bytes32[]))",
		"55e48749": "newKeyRequested()",
	},
}

var DKGClientABI = DKGClientMetaData.ABI

var DKGClientFuncSigs = DKGClientMetaData.Sigs

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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9201de55": "bytes32ToString(bytes32)",
		"39614e4f": "bytesToString(bytes)",
		"0bc643e8": "toASCII(uint8)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610529806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630bc643e81461004657806339614e4f146100705780639201de5514610090575b600080fd5b6100596100543660046102c3565b6100a3565b60405160ff90911681526020015b60405180910390f35b61008361007e3660046102fc565b6100cd565b60405161006791906103ad565b61008361009e366004610402565b610237565b6000600a8260ff1610156100c2576100bc826030610431565b92915050565b6100bc826057610431565b6060600080835160026100e09190610456565b67ffffffffffffffff8111156100f8576100f86102e6565b6040519080825280601f01601f191660200182016040528015610122576020820181803683370190505b509050600091505b80518260ff16101561023057600084610144600285610475565b60ff1681518110610157576101576104a5565b60209101015160f81c600f1690506000600486610175600287610475565b60ff1681518110610188576101886104a5565b01602001516001600160f81b031916901c60f81c90506101a7816100a3565b60f81b838560ff16815181106101bf576101bf6104a5565b60200101906001600160f81b031916908160001a9053506101e1846001610431565b93506101ec826100a3565b60f81b838560ff1681518110610204576102046104a5565b60200101906001600160f81b031916908160001a90535050508180610228906104bb565b92505061012a565b9392505050565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156102b95783816020811061027a5761027a6104a5565b1a60f81b828281518110610290576102906104a5565b60200101906001600160f81b031916908160001a905350806102b1816104da565b91505061025f565b50610230816100cd565b6000602082840312156102d557600080fd5b813560ff8116811461023057600080fd5b634e487b7160e01b600052604160045260246000fd5b60006020828403121561030e57600080fd5b813567ffffffffffffffff8082111561032657600080fd5b818401915084601f83011261033a57600080fd5b81358181111561034c5761034c6102e6565b604051601f8201601f19908116603f01168101908382118183101715610374576103746102e6565b8160405282815287602084870101111561038d57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156103da578581018301518582016040015282016103be565b818111156103ec576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561041457600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff0382111561044e5761044e61041b565b019392505050565b60008160001904831182151516156104705761047061041b565b500290565b600060ff83168061049657634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036104d1576104d161041b565b60010192915050565b6000600182016104ec576104ec61041b565b506001019056fea26469706673582212202c35cc92ba701e5bc53cb5db8927b173d0a40e96ae53f40ba7486078f43c492d64736f6c634300080d0033",
}

var DebugABI = DebugMetaData.ABI

var DebugFuncSigs = DebugMetaData.Sigs

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
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212207855ac9f9daa967983394f3af34267e2999fe11445f155d97f1a43874bc6848864736f6c634300080d0033",
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

var IVRFBeaconConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5a47dd71": "rawFulfillRandomWords(uint48,uint256[],bytes)",
	},
}

var IVRFBeaconConsumerABI = IVRFBeaconConsumerMetaData.ABI

var IVRFBeaconConsumerFuncSigs = IVRFBeaconConsumerMetaData.Sigs

type IVRFBeaconConsumer struct {
	IVRFBeaconConsumerCaller
	IVRFBeaconConsumerTransactor
	IVRFBeaconConsumerFilterer
}

type IVRFBeaconConsumerCaller struct {
	contract *bind.BoundContract
}

type IVRFBeaconConsumerTransactor struct {
	contract *bind.BoundContract
}

type IVRFBeaconConsumerFilterer struct {
	contract *bind.BoundContract
}

type IVRFBeaconConsumerSession struct {
	Contract     *IVRFBeaconConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFBeaconConsumerCallerSession struct {
	Contract *IVRFBeaconConsumerCaller
	CallOpts bind.CallOpts
}

type IVRFBeaconConsumerTransactorSession struct {
	Contract     *IVRFBeaconConsumerTransactor
	TransactOpts bind.TransactOpts
}

type IVRFBeaconConsumerRaw struct {
	Contract *IVRFBeaconConsumer
}

type IVRFBeaconConsumerCallerRaw struct {
	Contract *IVRFBeaconConsumerCaller
}

type IVRFBeaconConsumerTransactorRaw struct {
	Contract *IVRFBeaconConsumerTransactor
}

func NewIVRFBeaconConsumer(address common.Address, backend bind.ContractBackend) (*IVRFBeaconConsumer, error) {
	contract, err := bindIVRFBeaconConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFBeaconConsumer{IVRFBeaconConsumerCaller: IVRFBeaconConsumerCaller{contract: contract}, IVRFBeaconConsumerTransactor: IVRFBeaconConsumerTransactor{contract: contract}, IVRFBeaconConsumerFilterer: IVRFBeaconConsumerFilterer{contract: contract}}, nil
}

func NewIVRFBeaconConsumerCaller(address common.Address, caller bind.ContractCaller) (*IVRFBeaconConsumerCaller, error) {
	contract, err := bindIVRFBeaconConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFBeaconConsumerCaller{contract: contract}, nil
}

func NewIVRFBeaconConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFBeaconConsumerTransactor, error) {
	contract, err := bindIVRFBeaconConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFBeaconConsumerTransactor{contract: contract}, nil
}

func NewIVRFBeaconConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFBeaconConsumerFilterer, error) {
	contract, err := bindIVRFBeaconConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFBeaconConsumerFilterer{contract: contract}, nil
}

func bindIVRFBeaconConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFBeaconConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFBeaconConsumer.Contract.IVRFBeaconConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFBeaconConsumer.Contract.IVRFBeaconConsumerTransactor.contract.Transfer(opts)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFBeaconConsumer.Contract.IVRFBeaconConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFBeaconConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFBeaconConsumer.Contract.contract.Transfer(opts)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFBeaconConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFBeaconConsumer.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords, arguments)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFBeaconConsumer.Contract.RawFulfillRandomWords(&_IVRFBeaconConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_IVRFBeaconConsumer *IVRFBeaconConsumerTransactorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _IVRFBeaconConsumer.Contract.RawFulfillRandomWords(&_IVRFBeaconConsumer.TransactOpts, requestID, randomWords, arguments)
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

var LinkTokenInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"decimalPlaces\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTokensIssued\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"66188463": "decreaseApproval(address,uint256)",
		"d73dd623": "increaseApproval(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"4000aea0": "transferAndCall(address,uint256,bytes)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

var LinkTokenInterfaceABI = LinkTokenInterfaceMetaData.ABI

var LinkTokenInterfaceFuncSigs = LinkTokenInterfaceMetaData.Sigs

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

var OCR2AbstractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"81ff7048": "latestConfigDetails()",
		"afcb95d7": "latestConfigDigestAndEpoch()",
		"e3d0e712": "setConfig(address[],address[],uint8,bytes,uint64,bytes)",
		"b1dc65a4": "transmit(bytes32[3],bytes,bytes32[],bytes32[],bytes32)",
		"181f5a77": "typeAndVersion()",
	},
}

var OCR2AbstractABI = OCR2AbstractMetaData.ABI

var OCR2AbstractFuncSigs = OCR2AbstractMetaData.Sigs

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
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"f2fde38b": "transferOwnership(address)",
	},
}

var OwnableInterfaceABI = OwnableInterfaceMetaData.ABI

var OwnableInterfaceFuncSigs = OwnableInterfaceMetaData.Sigs

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
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6102a9806101576000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d366004610243565b610131565b6001546001600160a01b031633146100da5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610139610145565b6101428161019a565b50565b6000546001600160a01b031633146101985760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b60448201526064016100d1565b565b336001600160a01b038216036101f25760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d1565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561025557600080fd5b81356001600160a01b038116811461026c57600080fd5b939250505056fea2646970667358221220602cbac6f8505d81ac263d3de7989905a3fb63f6073b0bf6b7550f85fba2d8e264736f6c634300080d0033",
}

var OwnerIsCreatorABI = OwnerIsCreatorMetaData.ABI

var OwnerIsCreatorFuncSigs = OwnerIsCreatorMetaData.Sigs

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
	Sigs: map[string]string{
		"181f5a77": "typeAndVersion()",
	},
}

var TypeAndVersionInterfaceABI = TypeAndVersionInterfaceMetaData.ABI

var TypeAndVersionInterfaceFuncSigs = TypeAndVersionInterfaceMetaData.Sigs

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

var VRFBeaconBillingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212202708d3073522ae7e79ee5cf2fba8ff4f07085475dbc17a9761ed461eea65bf0064736f6c634300080d0033",
}

var VRFBeaconBillingABI = VRFBeaconBillingMetaData.ABI

var VRFBeaconBillingBin = VRFBeaconBillingMetaData.Bin

func DeployVRFBeaconBilling(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFBeaconBilling, error) {
	parsed, err := VRFBeaconBillingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconBillingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeaconBilling{VRFBeaconBillingCaller: VRFBeaconBillingCaller{contract: contract}, VRFBeaconBillingTransactor: VRFBeaconBillingTransactor{contract: contract}, VRFBeaconBillingFilterer: VRFBeaconBillingFilterer{contract: contract}}, nil
}

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

var VRFBeaconCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"contractDKG\",\"name\":\"keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OffchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"occVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconReport.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconReport.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxErrorMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2f7527cc": "NUM_CONF_DELAYS()",
		"79ba5097": "acceptOwnership()",
		"b121e147": "acceptPayeeship(address)",
		"c278e5b7": "exposeType(((uint64,uint24,(uint256[2]),((uint48,uint16,address,bytes,uint64,uint96),uint96)[])[],uint192,uint64,bytes32))",
		"29937268": "getBilling()",
		"c4c92b37": "getBillingAccessController()",
		"0b93e168": "getRandomness(uint48)",
		"cf7e754a": "i_StartSlot()",
		"cd0593df": "i_beaconPeriodBlocks()",
		"bf2732c7": "keyGenerated((bytes,bytes32[]))",
		"81ff7048": "latestConfigDetails()",
		"afcb95d7": "latestConfigDigestAndEpoch()",
		"d09dc339": "linkAvailableForPayment()",
		"7a464944": "maxErrorMsgLength()",
		"bbcdd0d8": "maxNumWords()",
		"c63c4e9b": "minDelay()",
		"55e48749": "newKeyRequested()",
		"e4902f82": "oracleObservationCount(address)",
		"0eafb25b": "owedPayment(address)",
		"8da5cb5b": "owner()",
		"dc92accf": "requestRandomness(uint16,uint64,uint24)",
		"f645dcb1": "requestRandomnessFulfillment(uint64,uint16,uint24,uint32,bytes)",
		"cc31f7dd": "s_keyID()",
		"d57fc45a": "s_provingKeyHash()",
		"643dc105": "setBilling(uint32,uint32,uint32,uint32,uint24)",
		"fbffd2c1": "setBillingAccessController(address)",
		"e3d0e712": "setConfig(address[],address[],uint8,bytes,uint64,bytes)",
		"9c849b30": "setPayees(address[],address[])",
		"f2fde38b": "transferOwnership(address)",
		"eb5dcd6c": "transferPayeeship(address,address)",
		"b1dc65a4": "transmit(bytes32[3],bytes,bytes32[],bytes32[],bytes32)",
		"181f5a77": "typeAndVersion()",
		"c1075329": "withdrawFunds(address,uint256)",
		"8ac28d5a": "withdrawPayment(address)",
	},
	Bin: "0x60c06040523480156200001157600080fd5b506040516200553938038062005539833981016040819052620000349162000232565b81818486338060008480806000036200006057604051632abc297960e01b815260040160405180910390fd5b6080819052600062000073824362000280565b9050600081608051620000879190620002b9565b9050620000958143620002d3565b60a0525050506001600160a01b0383169050620000f95760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600680546001600160a01b0319166001600160a01b03848116919091179091558116156200012c576200012c816200016d565b5050601780546001600160a01b039384166001600160a01b031991821617909155601880549690931695169490941790555060195550620002ee9350505050565b336001600160a01b03821603620001c75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000f0565b600780546001600160a01b0319166001600160a01b03838116918217909255600654604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6001600160a01b03811681146200022f57600080fd5b50565b600080600080608085870312156200024957600080fd5b8451620002568162000219565b602086015160408701519195509350620002708162000219565b6060959095015193969295505050565b6000826200029e57634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b600082821015620002ce57620002ce620002a3565b500390565b60008219821115620002e957620002e9620002a3565b500190565b60805160a0516152026200033760003960006104bc01526000818161049501528181610668015281816130b3015281816130e20152818161311a01526139df01526152026000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c8063bf2732c71161011a578063d09dc339116100ad578063e4902f821161007c578063e4902f821461052c578063eb5dcd6c14610554578063f2fde38b14610567578063f645dcb11461057a578063fbffd2c11461058d57600080fd5b8063d09dc339146104de578063d57fc45a146104e6578063dc92accf146104ef578063e3d0e7121461051957600080fd5b8063c63c4e9b116100e9578063c63c4e9b1461046c578063cc31f7dd14610487578063cd0593df14610490578063cf7e754a146104b757600080fd5b8063bf2732c714610424578063c107532914610437578063c278e5b71461044a578063c4c92b371461045b57600080fd5b80637a4649441161019d5780639c849b301161016c5780639c849b30146103b8578063afcb95d7146103cb578063b121e147146103f5578063b1dc65a414610408578063bbcdd0d81461041b57600080fd5b80637a4649441461034c57806381ff7048146103545780638ac28d5a146103805780638da5cb5b1461039357600080fd5b80632f7527cc116101d95780632f7527cc1461030b57806355e4874914610325578063643dc1051461033157806379ba50971461034457600080fd5b80630b93e1681461020b5780630eafb25b14610234578063181f5a7714610255578063299372681461028c575b600080fd5b61021e610219366004613e19565b6105a0565b60405161022b9190613e71565b60405180910390f35b610247610242366004613e99565b61073b565b60405190815260200161022b565b6040805180820182526015815274565246426561636f6e20312e302e302d616c70686160581b6020820152905161022b9190613f0e565b6102cf600554600160501b810463ffffffff90811692600160701b8304821692600160901b8104831692600160b01b82041691600160d01b90910462ffffff1690565b6040805163ffffffff9687168152948616602086015292851692840192909252909216606082015262ffffff909116608082015260a00161022b565b610313600881565b60405160ff909116815260200161022b565b61032f6000601a55565b005b61032f61033f366004613f4b565b610840565b61032f610a26565b610247608081565b6007546009546040805160008152600160c01b90930463ffffffff16602084015282015260600161022b565b61032f61038e366004613e99565b610ad4565b6006546001600160a01b03165b6040516001600160a01b03909116815260200161022b565b61032f6103c6366004613fff565b610b46565b600954600b546040805160008152602081019390935263ffffffff9091169082015260600161022b565b61032f610403366004613e99565b610d18565b61032f6104163660046140ab565b610df4565b6102476103e881565b61032f6104323660046142d5565b611277565b61032f6104453660046143bd565b6112b7565b61032f6104583660046143e9565b50565b6016546001600160a01b03166103a0565b610474600381565b60405161ffff909116815260200161022b565b61024760195481565b6102477f000000000000000000000000000000000000000000000000000000000000000081565b6102477f000000000000000000000000000000000000000000000000000000000000000081565b610247611508565b610247601a5481565b6105026104fd366004614453565b611598565b60405165ffffffffffff909116815260200161022b565b61032f6105273660046144af565b6116b5565b61053f61053a366004613e99565b611de1565b60405163ffffffff909116815260200161022b565b61032f61056236600461459c565b611e90565b61032f610575366004613e99565b611fc9565b6105026105883660046145d5565b611fda565b61032f61059b366004613e99565b6120d9565b65ffffffffffff811660008181526004602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff600160381b820416938201939093526001600160a01b03600160481b84048116606083810191825298909752949093526001600160e81b03199091169055915116331461065d576060810151604051638e30e82360e01b81526001600160a01b0390911660048201523360248201526044015b60405180910390fd5b8051600090610693907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16614670565b90506000826020015162ffffff16436106ac919061468f565b90508082106106d7576040516315ad27c360e01b815260048101839052436024820152604401610654565b6001600160401b03821115610702576040516302c6ef8160e11b815260048101839052602401610654565b60008281526001602090815260408083208287015162ffffff1684529091529020546107329086908590856120ea565b95945050505050565b6001600160a01b0381166000908152600c602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b0316918101919091529061079d5750600092915050565b6005546020820151600091600160901b900463ffffffff169060109060ff16601f81106107cc576107cc6146a6565b6008810491909101546005546107ff926007166004026101000a90910463ffffffff90811691600160301b9004166146bc565b63ffffffff1661080f9190614670565b61081d90633b9aca00614670565b905081604001516001600160601b03168161083891906146e1565b949350505050565b6016546001600160a01b031661085e6006546001600160a01b031690565b6001600160a01b0316336001600160a01b031614806108ea5750604051630d629b5f60e31b81526001600160a01b03821690636b14daf8906108a99033906000903690600401614722565b602060405180830381865afa1580156108c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ea9190614747565b6109365760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c6044820152606401610654565b61093e6122be565b6005805467ffffffffffffffff60501b1916600160501b63ffffffff89811691820263ffffffff60701b191692909217600160701b8984169081029190911767ffffffffffffffff60901b1916600160901b89851690810263ffffffff60b01b191691909117600160b01b9489169485021762ffffff60d01b1916600160d01b62ffffff89169081029190911790955560408051938452602084019290925290820152606081019190915260808101919091527f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f9060a00160405180910390a1505050505050565b6007546001600160a01b03163314610a795760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b6044820152606401610654565b600680546001600160a01b0319808216339081179093556007805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6001600160a01b03818116600090815260146020526040902054163314610b3d5760405162461bcd60e51b815260206004820152601760248201527f4f6e6c792070617965652063616e2077697468647261770000000000000000006044820152606401610654565b61045881612633565b610b4e61281b565b828114610b9d5760405162461bcd60e51b815260206004820181905260248201527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a656044820152606401610654565b60005b83811015610d11576000858583818110610bbc57610bbc6146a6565b9050602002016020810190610bd19190613e99565b90506000848484818110610be757610be76146a6565b9050602002016020810190610bfc9190613e99565b6001600160a01b038084166000908152601460205260409020549192501680158080610c395750826001600160a01b0316826001600160a01b0316145b610c795760405162461bcd60e51b81526020600482015260116024820152701c185e595948185b1c9958591e481cd95d607a1b6044820152606401610654565b6001600160a01b03848116600090815260146020526040902080546001600160a01b03191685831690811790915590831614610cfa57826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b505050508080610d0990614769565b915050610ba0565b5050505050565b6001600160a01b03818116600090815260156020526040902054163314610d815760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792070726f706f736564207061796565732063616e20616363657074006044820152606401610654565b6001600160a01b0381811660008181526014602090815260408083208054336001600160a01b031980831682179093556015909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a60408051610100808201835260055460ff808216845291810464ffffffffff16602080850191909152600160301b820463ffffffff90811685870152600160501b830481166060860152600160701b830481166080860152600160901b8304811660a0860152600160b01b83041660c0850152600160d01b90910462ffffff1660e0840152336000908152600c825293909320549394509092918c01359116610ee25760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610654565b6009548b3514610f2c5760405162461bcd60e51b81526020600482015260156024820152740c6dedcccd2ce88d2cecae6e840dad2e6dac2e8c6d605b1b6044820152606401610654565b610f3a8a8a8a8a8a8a612870565b8151610f47906001614782565b60ff168714610f985760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610654565b868514610fe75760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610654565b60008a8a604051610ff99291906147a7565b604051908190038120611010918e906020016147b7565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a8110156111a85760006001858a846020811061105f5761105f6146a6565b61106c91901a601b614782565b8f8f8681811061107e5761107e6146a6565b905060200201358e8e87818110611097576110976146a6565b90506020020135604051600081526020016040526040516110d4949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156110f6573d6000803e3d6000fd5b505060408051601f198101516001600160a01b0381166000908152600d602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506111815760405162461bcd60e51b815260206004820152600f60248201526e39b4b3b730ba3ab9329032b93937b960891b6044820152606401610654565b826020015160080260ff166001901b840193505080806111a090614769565b915050611040565b5081827e01010101010101010101010101010101010101010101010101010101010101161461120c5760405162461bcd60e51b815260206004820152601060248201526f323ab83634b1b0ba329039b4b3b732b960811b6044820152606401610654565b506000915061125b9050838d836020020135848e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061290d92505050565b905061126983828633612d3a565b505050505050505050505050565b60185481516040516001600160a01b039092169161129891906020016147d3565b60408051601f198184030181529190528051602090910120601a555050565b6006546001600160a01b03163314806113415750601654604051630d629b5f60e31b81526001600160a01b0390911690636b14daf8906113009033906000903690600401614722565b602060405180830381865afa15801561131d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113419190614747565b61138d5760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c6044820152606401610654565b6000611397612e49565b6017546040516370a0823160e01b81523060048201529192506000916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156113e5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061140991906147ef565b9050818110156114525760405162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e742062616c616e636560601b6044820152606401610654565b6017546001600160a01b031663a9059cbb85611477611471868661468f565b87613013565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156114c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e69190614747565b6115025760405162461bcd60e51b815260040161065490614808565b50505050565b6017546040516370a0823160e01b815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa158015611555573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061157991906147ef565b90506000611585612e49565b90506115918183614834565b9250505090565b6000806000806115a8878661302d565b92509250925065ffffffffffff831660009081526004602090815260409182902084518154928601518487015160608801516001600160a01b0316600160481b027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff909216600160381b0291909116670100000000000000600160e81b031962ffffff9093166401000000000266ffffffffffffff1990961663ffffffff90941693909317949094171617919091179055516001600160401b038216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc5906116a090889062ffffff91909116815260200190565b60405180910390a250909150505b9392505050565b6116bd61281b565b601f8911156117015760405162461bcd60e51b815260206004820152601060248201526f746f6f206d616e79206f7261636c657360801b6044820152606401610654565b8887146117495760405162461bcd60e51b81526020600482015260166024820152750dee4c2c6d8ca40d8cadccee8d040dad2e6dac2e8c6d60531b6044820152606401610654565b88611755876003614873565b60ff16106117a55760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610654565b6117b18660ff166132e9565b6040805160e060208c02808301820190935260c082018c815260009383928f918f918291908601908490808284376000920191909152505050908252506040805160208c810282810182019093528c82529283019290918d918d91829185019084908082843760009201919091525050509082525060ff891660208083019190915260408051601f8a0183900483028101830182528981529201919089908990819084018382808284376000920191909152505050908252506001600160401b03861660208083019190915260408051601f8701839004830281018301825286815292019190869086908190840183828082843760009201919091525050509152506005805465ffffffffff001916905590506118cc6122be565b600e5460005b8181101561197d576000600e82815481106118ef576118ef6146a6565b6000918252602082200154600f80546001600160a01b039092169350908490811061191c5761191c6146a6565b60009182526020808320909101546001600160a01b039485168352600d82526040808420805461ffff1916905594168252600c90529190912080546dffffffffffffffffffffffffffff19169055508061197581614769565b9150506118d2565b5061198a600e6000613c30565b611996600f6000613c30565b60005b825151811015611c0f57600d6000846000015183815181106119bd576119bd6146a6565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611a315760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610654565b604080518082019091526001815260ff8216602082015283518051600d9160009185908110611a6257611a626146a6565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051600c92919084908110611acd57611acd6146a6565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611b415760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610654565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b0316815250600c600085602001518481518110611b8657611b866146a6565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516001600160601b0316620100000262010000600160701b031960ff959095166101000261ff00199315159390931661ffff1990941693909317919091179290921617905580611c0781614769565b915050611999565b5081518051611c2691600e91602090910190613c4e565b506020808301518051611c3d92600f920190613c4e565b5060408201516005805460ff191660ff9092169190911790556007805463ffffffff60c01b198116600160c01b63ffffffff43811682029290921793849055909104811691600091611c9891600160a01b900416600161489c565b905080600760146101000a81548163ffffffff021916908363ffffffff1602179055506000611cec46308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a0015161332e565b9050806009600001819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05838284886000015189602001518a604001518b606001518c608001518d60a00151604051611d4f999897969594939291906148fd565b60405180910390a1600554600160301b900463ffffffff1660005b865151811015611dc45781601082601f8110611d8857611d886146a6565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff1602179055508080611dbc90614769565b915050611d6a565b50611dcf8b8b613389565b50505050505050505050505050505050565b6001600160a01b0381166000908152600c602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b03169181019190915290611e435750600092915050565b6010816020015160ff16601f8110611e5d57611e5d6146a6565b6008810491909101546005546116ae926007166004026101000a90910463ffffffff90811691600160301b9004166146bc565b6001600160a01b03828116600090815260146020526040902054163314611ef95760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792063757272656e742070617965652063616e207570646174650000006044820152606401610654565b6001600160a01b0381163303611f515760405162461bcd60e51b815260206004820152601760248201527f63616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610654565b6001600160a01b03808316600090815260156020526040902080548383166001600160a01b031982168117909255909116908114611fc4576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a45b505050565b611fd161281b565b61045881613397565b6000806000611fe9878761302d565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a6001600160401b031681526020018763ffffffff166001600160601b0316815250905081878a836040516020016120639493929190614992565b60408051601f19818403018152828252805160209182012065ffffffffffff8716600090815291829052919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906120c49084908a908d908690614992565b60405180910390a15090979650505050505050565b6120e161281b565b61045881613441565b6060826121235760405163c7d41b1b60e01b815265ffffffffffff861660048201526001600160401b0383166024820152604401610654565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff1611156121d9576040808601519051634a90778560e01b815261ffff90911660048201526103e86024820152604401610654565b6000856040015161ffff166001600160401b038111156121fb576121fb614161565b604051908082528060200260200182016040528015612224578160200160208202803683370190505b50905060005b866040015161ffff168161ffff1610156122b357828160405160200161226792919091825260f01b6001600160f01b031916602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff1681518110612296576122966146a6565b6020908102919091010152806122ab81614a34565b91505061222a565b509695505050505050565b601754600554604080516103e08101918290526001600160a01b0390931692600160301b90920463ffffffff1691600091601090601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116122f9579050505050505090506000600f80548060200260200160405190810160405280929190818152602001828054801561239457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612376575b5050505050905060005b8151811015612625576000600c60008484815181106123bf576123bf6146a6565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b031690506000600c6000858581518110612421576124216146a6565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008483601f8110612484576124846146a6565b602002015160055490870363ffffffff9081169250600160901b909104168102633b9aca00028201801561261a576000601460008787815181106124ca576124ca6146a6565b6020908102919091018101516001600160a01b03908116835290820192909252604090810160002054905163a9059cbb60e01b815290821660048201819052602482018590529250908a169063a9059cbb906044016020604051808303816000875af115801561253e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125629190614747565b61257e5760405162461bcd60e51b815260040161065490614808565b878786601f8110612591576125916146a6565b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b03168787815181106125ce576125ce6146a6565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8560405161261091815260200190565b60405180910390a4505b50505060010161239e565b50610d11601083601f613cb3565b6001600160a01b0381166000908152600c60209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b031692810192909252612690575050565b600061269b8361073b565b90508015611fc4576001600160a01b038381166000908152601460205260409081902054601754915163a9059cbb60e01b8152908316600482018190526024820185905292919091169063a9059cbb906044016020604051808303816000875af115801561270d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127319190614747565b61274d5760405162461bcd60e51b815260040161065490614808565b600560000160069054906101000a900463ffffffff166010846020015160ff16601f811061277d5761277d6146a6565b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b038481166000818152600c6020908152604091829020805462010000600160701b0319169055601754915186815291841693851692917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c910160405180910390a450505050565b6006546001600160a01b0316331461286e5760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b6044820152606401610654565b565b600061287d826020614670565b612888856020614670565b612894886101446146e1565b61289e91906146e1565b6128a891906146e1565b6128b39060006146e1565b90503681146129045760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610654565b50505050505050565b600080828060200190518101906129249190614c26565b64ffffffffff8516602088015260408701805191925061294382614dfa565b63ffffffff1663ffffffff168152505085600560008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050600081604001516001600160401b031640905080826060015114612ae5576060820151604080840151905163aed0afe560e01b81526004810192909252602482018390526001600160401b03166044820152606401610654565b6000808360000151516001600160401b03811115612b0557612b05614161565b604051908082528060200260200182016040528015612b4a57816020015b6040805180820190915260008082526020820152815260200190600190039081612b235790505b50905060005b845151811015612c1a57600085600001518281518110612b7257612b726146a6565b60200260200101519050612b8f81876040015188602001516134b7565b60408101515151151580612bab57506040810151516020015115155b15612c0757604051806040016040528082600001516001600160401b03168152602001826020015162ffffff16815250838381518110612bed57612bed6146a6565b60200260200101819052508380612c0390614a34565b9450505b5080612c1281614769565b915050612b50565b5060008261ffff166001600160401b03811115612c3957612c39614161565b604051908082528060200260200182016040528015612c7e57816020015b6040805180820190915260008082526020820152815260200190600190039081612c575790505b50905060005b8361ffff16811015612cda57828181518110612ca257612ca26146a6565b6020026020010151828281518110612cbc57612cbc6146a6565b60200260200101819052508080612cd290614769565b915050612c84565b50896040015163ffffffff167f7484067466b4f2452757769a8dc9a8b41497154367515673c79386f9f0b74f163387602001518c8c86604051612d21959493929190614e13565b60405180910390a2505050506020015195945050505050565b6000612d61633b9aca003a04866080015163ffffffff16876060015163ffffffff16613890565b90506010360260005a90506000612d8a8663ffffffff1685858b60e0015162ffffff16866138ad565b90506000670de0b6b3a76400006001600160c01b03891683026001600160a01b0388166000908152600c602052604090205460c08c01519290910492506201000090046001600160601b039081169163ffffffff16633b9aca000282840101908116821115612dff5750505050505050611502565b6001600160a01b0388166000908152600c6020526040902080546001600160601b03909216620100000262010000600160701b031990921691909117905550505050505050505050565b600080600f805480602002602001604051908101604052809291908181526020018280548015612ea257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612e84575b50508351600554604080516103e08101918290529697509195600160301b90910463ffffffff169450600093509150601090601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612edb5790505050505050905060005b83811015612f6e578181601f8110612f3b57612f3b6146a6565b6020020151612f4a90846146bc565b612f5a9063ffffffff16876146e1565b955080612f6681614769565b915050612f21565b50600554612f8d90600160901b900463ffffffff16633b9aca00614670565b612f979086614670565b945060005b8381101561300b57600c6000868381518110612fba57612fba6146a6565b6020908102919091018101516001600160a01b0316825281019190915260400160002054612ff7906201000090046001600160601b0316876146e1565b95508061300381614769565b915050612f9c565b505050505090565b600081831015613024575081613027565b50805b92915050565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff16111561308757604051634a90778560e01b815261ffff861660048201526103e86024820152604401610654565b8461ffff166000036130ac576040516308fad2a760e01b815260040160405180910390fd5b60006130d87f000000000000000000000000000000000000000000000000000000000000000043614ec2565b90506000816131077f0000000000000000000000000000000000000000000000000000000000000000436146e1565b613111919061468f565b9050600061313f7f000000000000000000000000000000000000000000000000000000000000000083614ed6565b905063ffffffff8110613165576040516307b2a52360e41b815260040160405180910390fd5b6040805180820182526002805465ffffffffffff16825282516101008101938490528493600093929160208401916003906008908288855b82829054906101000a900462ffffff1662ffffff168152602001906003019060208260020104928301926001038202915080841161319d57905050505091909252505081519192505065ffffffffffff8082161061320e57604051630568cab760e31b815260040160405180910390fd5b613219816001614eea565b6002805465ffffffffffff191665ffffffffffff9290921691909117905560005b6008811015613280578a62ffffff168360200151826008811061325f5761325f6146a6565b602002015162ffffff1614613280578061327881614769565b91505061323a565b600881106132a8576020830151604051630c4f769b60e41b8152610654918d91600401614f33565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b806000106104585760405162461bcd60e51b815260206004820152601260248201527166206d75737420626520706f73697469766560701b6044820152606401610654565b6000808a8a8a8a8a8a8a8a8a60405160200161335299989796959493929190614f4d565b60408051601f1981840301815291905280516020909101206001600160f01b0316600160f01b179150509998505050505050505050565b6133938282613911565b5050565b336001600160a01b038216036133ef5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610654565b600780546001600160a01b0319166001600160a01b03838116918217909255600654604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6016546001600160a01b03908116908216811461339357601680546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b82516001600160401b03808416911611156134fb57825160405163012d824d60e01b81526001600160401b0380851660048301529091166024820152604401610654565b60408301515151600090158015613519575060408401515160200151155b15613551575082516001600160401b031660009081526001602090815260408083208287015162ffffff1684529091529020546135ab565b83604001516040516020016135669190614fd6565b60408051601f19818403018152918152815160209283012086516001600160401b03166000908152600184528281208885015162ffffff168252909352912081905590505b6060840151516000816001600160401b038111156135cb576135cb614161565b6040519080825280602002602001820160405280156135f4578160200160208202803683370190505b5090506000826001600160401b0381111561361157613611614161565b6040519080825280601f01601f19166020018201604052801561363b576020820181803683370190505b5090506000836001600160401b0381111561365857613658614161565b60405190808252806020026020018201604052801561368b57816020015b60608152602001906001900390816136765790505b5090506000805b8581101561378e5760008a6060015182815181106136b2576136b26146a6565b602090810291909101015190506000806136d68d600001518e602001518c866139d5565b9150915081156137155780868661ffff16815181106136f7576136f76146a6565b6020026020010181905250848061370d90614a34565b955050613744565b600160f81b87858151811061372c5761372c6146a6565b60200101906001600160f81b031916908160001a9053505b825151885189908690811061375b5761375b6146a6565b602002602001019065ffffffffffff16908165ffffffffffff16815250505050508061378681614769565b915050613692565b50606089015151156138855760008161ffff166001600160401b038111156137b8576137b8614161565b6040519080825280602002602001820160405280156137eb57816020015b60608152602001906001900390816137d65790505b50905060005b8261ffff168110156138475783818151811061380f5761380f6146a6565b6020026020010151828281518110613829576138296146a6565b6020026020010181905250808061383f90614769565b9150506137f1565b507f47ddf7bb0cbd94c1b43c5097f1352a80db0ceb3696f029d32b24f32cd631d2b785858360405161387b93929190615009565b60405180910390a1505b505050505050505050565b600083838110156138a357600285850304015b6107328184613013565b6000818610156138ff5760405162461bcd60e51b815260206004820181905260248201527f6c6566744761732063616e6e6f742065786365656420696e697469616c4761736044820152606401610654565b50633b9aca0094039190910101020290565b61010081811461393a57828282604051635c9d52ef60e11b8152600401610654939291906150bf565b613942613d4a565b818160405160200161395491906150e3565b6040516020818303038152906040525114613971576139716150f2565b6040805180820190915260025465ffffffffffff1681526020810161399885870187615108565b905280516002805465ffffffffffff191665ffffffffffff90921691909117815560208201516139cc906003906008613d69565b50611502915050565b6000606081613a0d7f00000000000000000000000000000000000000000000000000000000000000006001600160401b038916614ed6565b845160808101516040519293509091600091613a31918b918b918690602001614992565b60408051601f198184030181529181528151602092830120845165ffffffffffff166000908152928390529120549091508114613a9f5760016040518060400160405280601081526020016f756e6b6e6f776e2063616c6c6261636b60801b81525094509450505050613c27565b815165ffffffffffff16600090815260208181526040808320839055805160808101825263ffffffff8716815262ffffff8c16818401529185015161ffff16828201528401516001600160a01b031660608201528351909190613b0490838b8e6120ea565b6060808401518a5160a00151875192880151604051635a47dd7160e01b815294955091936001600160a01b03851693635a47dd71936001600160601b0390931692613b549288919060040161518f565b600060405180830381600088803b158015613b6e57600080fd5b5087f193505050508015613b80575060015b613c0a573d808015613bae576040519150601f19603f3d011682016040523d82523d6000602084013e613bb3565b606091505b50608081511015613bd057600198509650613c2795505050505050565b60016040518060400160405280600f81526020016e6572726d736720746f6f206c6f6e6760881b8152509850985050505050505050613c27565b600060405180602001604052806000815250975097505050505050505b94509492505050565b50805460008255906000526020600020908101906104589190613df0565b828054828255906000526020600020908101928215613ca3579160200282015b82811115613ca357825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613c6e565b50613caf929150613df0565b5090565b600483019183908215613ca35791602002820160005b83821115613d0d57835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302613cc9565b8015613d3d5782816101000a81549063ffffffff0219169055600401602081600301049283019260010302613d0d565b5050613caf929150613df0565b6040518061010001604052806008906020820280368337509192915050565b600183019183908215613ca35791602002820160005b83821115613dc157835183826101000a81548162ffffff021916908362ffffff1602179055509260200192600301602081600201049283019260010302613d7f565b8015613d3d5782816101000a81549062ffffff0219169055600301602081600201049283019260010302613dc1565b5b80821115613caf5760008155600101613df1565b65ffffffffffff8116811461045857600080fd5b600060208284031215613e2b57600080fd5b81356116ae81613e05565b600081518084526020808501945080840160005b83811015613e6657815187529582019590820190600101613e4a565b509495945050505050565b6020815260006116ae6020830184613e36565b6001600160a01b038116811461045857600080fd5b600060208284031215613eab57600080fd5b81356116ae81613e84565b60005b83811015613ed1578181015183820152602001613eb9565b838111156115025750506000910152565b60008151808452613efa816020860160208601613eb6565b601f01601f19169290920160200192915050565b6020815260006116ae6020830184613ee2565b803563ffffffff81168114613f3557600080fd5b919050565b62ffffff8116811461045857600080fd5b600080600080600060a08688031215613f6357600080fd5b613f6c86613f21565b9450613f7a60208701613f21565b9350613f8860408701613f21565b9250613f9660608701613f21565b91506080860135613fa681613f3a565b809150509295509295909350565b60008083601f840112613fc657600080fd5b5081356001600160401b03811115613fdd57600080fd5b6020830191508360208260051b8501011115613ff857600080fd5b9250929050565b6000806000806040858703121561401557600080fd5b84356001600160401b038082111561402c57600080fd5b61403888838901613fb4565b9096509450602087013591508082111561405157600080fd5b5061405e87828801613fb4565b95989497509550505050565b60008083601f84011261407c57600080fd5b5081356001600160401b0381111561409357600080fd5b602083019150836020828501011115613ff857600080fd5b60008060008060008060008060e0898b0312156140c757600080fd5b606089018a8111156140d857600080fd5b899850356001600160401b03808211156140f157600080fd5b6140fd8c838d0161406a565b909950975060808b013591508082111561411657600080fd5b6141228c838d01613fb4565b909750955060a08b013591508082111561413b57600080fd5b506141488b828c01613fb4565b999c989b50969995989497949560c00135949350505050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561419957614199614161565b60405290565b60405160c081016001600160401b038111828210171561419957614199614161565b604051608081016001600160401b038111828210171561419957614199614161565b604051602081016001600160401b038111828210171561419957614199614161565b604051601f8201601f191681016001600160401b038111828210171561422d5761422d614161565b604052919050565b60006001600160401b0382111561424e5761424e614161565b50601f01601f191660200190565b600082601f83011261426d57600080fd5b813561428061427b82614235565b614205565b81815284602083860101111561429557600080fd5b816020850160208301376000918101602001919091529392505050565b60006001600160401b038211156142cb576142cb614161565b5060051b60200190565b600060208083850312156142e857600080fd5b82356001600160401b03808211156142ff57600080fd5b908401906040828703121561431357600080fd5b61431b614177565b82358281111561432a57600080fd5b6143368882860161425c565b825250838301358281111561434a57600080fd5b80840193505086601f84011261435f57600080fd5b8235915061436f61427b836142b2565b82815260059290921b8301840191848101908884111561438e57600080fd5b938501935b838510156143ac57843582529385019390850190614393565b948201949094529695505050505050565b600080604083850312156143d057600080fd5b82356143db81613e84565b946020939093013593505050565b6000602082840312156143fb57600080fd5b81356001600160401b0381111561441157600080fd5b8201608081850312156116ae57600080fd5b61ffff8116811461045857600080fd5b6001600160401b038116811461045857600080fd5b8035613f3581614433565b60008060006060848603121561446857600080fd5b833561447381614423565b9250602084013561448381614433565b9150604084013561449381613f3a565b809150509250925092565b803560ff81168114613f3557600080fd5b60008060008060008060008060008060c08b8d0312156144ce57600080fd5b8a356001600160401b03808211156144e557600080fd5b6144f18e838f01613fb4565b909c509a5060208d013591508082111561450a57600080fd5b6145168e838f01613fb4565b909a50985088915061452a60408e0161449e565b975060608d013591508082111561454057600080fd5b61454c8e838f0161406a565b909750955085915061456060808e01614448565b945060a08d013591508082111561457657600080fd5b506145838d828e0161406a565b915080935050809150509295989b9194979a5092959850565b600080604083850312156145af57600080fd5b82356145ba81613e84565b915060208301356145ca81613e84565b809150509250929050565b600080600080600060a086880312156145ed57600080fd5b85356145f881614433565b9450602086013561460881614423565b9350604086013561461881613f3a565b925061462660608701613f21565b915060808601356001600160401b0381111561464157600080fd5b61464d8882890161425c565b9150509295509295909350565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561468a5761468a61465a565b500290565b6000828210156146a1576146a161465a565b500390565b634e487b7160e01b600052603260045260246000fd5b600063ffffffff838116908316818110156146d9576146d961465a565b039392505050565b600082198211156146f4576146f461465a565b500190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b038416815260406020820181905260009061073290830184866146f9565b60006020828403121561475957600080fd5b815180151581146116ae57600080fd5b60006001820161477b5761477b61465a565b5060010190565b600060ff821660ff84168060ff0382111561479f5761479f61465a565b019392505050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b600082516147e5818460208701613eb6565b9190910192915050565b60006020828403121561480157600080fd5b5051919050565b602080825260129082015271696e73756666696369656e742066756e647360701b604082015260600190565b60008083128015600160ff1b8501841216156148525761485261465a565b6001600160ff1b038401831381161561486d5761486d61465a565b50500390565b600060ff821660ff84168160ff04811182151516156148945761489461465a565b029392505050565b600063ffffffff8083168185168083038211156148bb576148bb61465a565b01949350505050565b600081518084526020808501945080840160005b83811015613e665781516001600160a01b0316875295820195908201906001016148d8565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261492d8184018a6148c4565b9050828103608084015261494181896148c4565b905060ff871660a084015282810360c084015261495e8187613ee2565b90506001600160401b03851660e08401528281036101008401526149828185613ee2565b9c9b505050505050505050505050565b60006001600160401b03808716835262ffffff8616602084015280851660408401526080606084015265ffffffffffff845116608084015261ffff60208501511660a084015260018060a01b0360408501511660c0840152606084015160c060e0850152614a04610140850182613ee2565b60808601519092166101008501525060a0909301516001600160601b031661012090920191909152509392505050565b600061ffff808316818103614a4b57614a4b61465a565b6001019392505050565b8051613f3581614433565b600082601f830112614a7157600080fd5b8151614a7f61427b82614235565b818152846020838601011115614a9457600080fd5b610838826020830160208701613eb6565b80516001600160601b0381168114613f3557600080fd5b600082601f830112614acd57600080fd5b81516020614add61427b836142b2565b82815260059290921b84018101918181019086841115614afc57600080fd5b8286015b848110156122b35780516001600160401b0380821115614b1f57600080fd5b90880190601f196040838c0382011215614b3857600080fd5b614b40614177565b8784015183811115614b5157600080fd5b840160c0818e0384011215614b6557600080fd5b614b6d61419f565b925088810151614b7c81613e05565b83526040810151614b8c81614423565b838a01526060810151614b9e81613e84565b6040840152608081015184811115614bb557600080fd5b614bc38e8b83850101614a60565b606085015250614bd560a08201614a55565b6080840152614be660c08201614aa5565b60a084015250818152614bfb60408501614aa5565b818901528652505050918301918301614b00565b80516001600160c01b0381168114613f3557600080fd5b600060208284031215614c3857600080fd5b81516001600160401b0380821115614c4f57600080fd5b9083019060808286031215614c6357600080fd5b614c6b6141c1565b825182811115614c7a57600080fd5b8301601f81018713614c8b57600080fd5b8051614c9961427b826142b2565b8082825260208201915060208360051b850101925089831115614cbb57600080fd5b602084015b83811015614dbb57805187811115614cd757600080fd5b850160a0818d03601f19011215614ced57600080fd5b614cf56141c1565b6020820151614d0381614433565b81526040820151614d1381613f3a565b60208201526040828e03605f19011215614d2c57600080fd5b614d346141e3565b8d607f840112614d4357600080fd5b614d4b614177565b808f60a086011115614d5c57600080fd5b606085015b60a08601811015614d7c578051835260209283019201614d61565b50825250604082015260a082015189811115614d9757600080fd5b614da68e602083860101614abc565b60608301525084525060209283019201614cc0565b50845250614dce91505060208401614c0f565b6020820152614ddf60408401614a55565b60408201526060830151606082015280935050505092915050565b600063ffffffff808316818103614a4b57614a4b61465a565b6001600160a01b03861681526001600160c01b038516602080830191909152604080830186905264ffffffffff8516606084015260a060808401819052845190840181905260009285810192909160c0860190855b81811015614e9b57855180516001600160401b0316845285015162ffffff16858401529484019491830191600101614e68565b50909b9a5050505050505050505050565b634e487b7160e01b600052601260045260246000fd5b600082614ed157614ed1614eac565b500690565b600082614ee557614ee5614eac565b500490565b600065ffffffffffff8083168185168083038211156148bb576148bb61465a565b8060005b600881101561150257815162ffffff16845260209384019390910190600101614f0f565b62ffffff8316815261012081016116ae6020830184614f0b565b8981526001600160a01b03891660208201526001600160401b03888116604083015261012060608301819052600091614f888483018b6148c4565b91508382036080850152614f9c828a6148c4565b915060ff881660a085015283820360c0850152614fb98288613ee2565b90861660e085015283810361010085015290506149828185613ee2565b815160408201908260005b6002811015615000578251825260209283019290910190600101614fe1565b50505092915050565b606080825284519082018190526000906020906080840190828801845b8281101561504a57815165ffffffffffff1684529284019290840190600101615026565b5050508381038285015261505e8187613ee2565b905083810360408501528085518083528383019150838160051b84010184880160005b838110156150af57601f1986840301855261509d838351613ee2565b94870194925090860190600101615081565b50909a9950505050505050505050565b6040815260006150d36040830185876146f9565b9050826020830152949350505050565b61010081016130278284614f0b565b634e487b7160e01b600052600160045260246000fd5b600061010080838503121561511c57600080fd5b83601f84011261512b57600080fd5b6040518181018181106001600160401b038211171561514c5761514c614161565b60405290830190808583111561516157600080fd5b845b8381101561518457803561517681613f3a565b825260209182019101615163565b509095945050505050565b65ffffffffffff841681526060602082015260006151b06060830185613e36565b82810360408401526151c28185613ee2565b969550505050505056fea264697066735822122014333acc19b4b7a5552e9a66577ce0592363abc1eb9a27ae74c44499738d8b2e64736f6c634300080d0033",
}

var VRFBeaconCoordinatorABI = VRFBeaconCoordinatorMetaData.ABI

var VRFBeaconCoordinatorFuncSigs = VRFBeaconCoordinatorMetaData.Sigs

var VRFBeaconCoordinatorBin = VRFBeaconCoordinatorMetaData.Bin

func DeployVRFBeaconCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, beaconPeriodBlocksArg *big.Int, keyProvider common.Address, keyID [32]byte) (common.Address, *types.Transaction, *VRFBeaconCoordinator, error) {
	parsed, err := VRFBeaconCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconCoordinatorBin), backend, link, beaconPeriodBlocksArg, keyProvider, keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeaconCoordinator{VRFBeaconCoordinatorCaller: VRFBeaconCoordinatorCaller{contract: contract}, VRFBeaconCoordinatorTransactor: VRFBeaconCoordinatorTransactor{contract: contract}, VRFBeaconCoordinatorFilterer: VRFBeaconCoordinatorFilterer{contract: contract}}, nil
}

type VRFBeaconCoordinator struct {
	VRFBeaconCoordinatorCaller
	VRFBeaconCoordinatorTransactor
	VRFBeaconCoordinatorFilterer
}

type VRFBeaconCoordinatorCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconCoordinatorTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconCoordinatorFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconCoordinatorSession struct {
	Contract     *VRFBeaconCoordinator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconCoordinatorCallerSession struct {
	Contract *VRFBeaconCoordinatorCaller
	CallOpts bind.CallOpts
}

type VRFBeaconCoordinatorTransactorSession struct {
	Contract     *VRFBeaconCoordinatorTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconCoordinatorRaw struct {
	Contract *VRFBeaconCoordinator
}

type VRFBeaconCoordinatorCallerRaw struct {
	Contract *VRFBeaconCoordinatorCaller
}

type VRFBeaconCoordinatorTransactorRaw struct {
	Contract *VRFBeaconCoordinatorTransactor
}

func NewVRFBeaconCoordinator(address common.Address, backend bind.ContractBackend) (*VRFBeaconCoordinator, error) {
	contract, err := bindVRFBeaconCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinator{VRFBeaconCoordinatorCaller: VRFBeaconCoordinatorCaller{contract: contract}, VRFBeaconCoordinatorTransactor: VRFBeaconCoordinatorTransactor{contract: contract}, VRFBeaconCoordinatorFilterer: VRFBeaconCoordinatorFilterer{contract: contract}}, nil
}

func NewVRFBeaconCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconCoordinatorCaller, error) {
	contract, err := bindVRFBeaconCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorCaller{contract: contract}, nil
}

func NewVRFBeaconCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconCoordinatorTransactor, error) {
	contract, err := bindVRFBeaconCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorTransactor{contract: contract}, nil
}

func NewVRFBeaconCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconCoordinatorFilterer, error) {
	contract, err := bindVRFBeaconCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorFilterer{contract: contract}, nil
}

func bindVRFBeaconCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconCoordinator.Contract.VRFBeaconCoordinatorCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.VRFBeaconCoordinatorTransactor.contract.Transfer(opts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.VRFBeaconCoordinatorTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconCoordinator.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.contract.Transfer(opts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconCoordinator.Contract.NUMCONFDELAYS(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconCoordinator.Contract.NUMCONFDELAYS(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) GetBilling(opts *bind.CallOpts) (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "getBilling")

	outstruct := new(struct {
		MaximumGasPriceGwei       uint32
		ReasonableGasPriceGwei    uint32
		ObservationPaymentGjuels  uint32
		TransmissionPaymentGjuels uint32
		AccountingGas             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPriceGwei = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ReasonableGasPriceGwei = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ObservationPaymentGjuels = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TransmissionPaymentGjuels = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _VRFBeaconCoordinator.Contract.GetBilling(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _VRFBeaconCoordinator.Contract.GetBilling(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeaconCoordinator.Contract.GetBillingAccessController(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeaconCoordinator.Contract.GetBillingAccessController(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) IStartSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "i_StartSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.IStartSlot(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.IStartSlot(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.IBeaconPeriodBlocks(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.IBeaconPeriodBlocks(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "latestConfigDetails")

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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRFBeaconCoordinator.Contract.LatestConfigDetails(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRFBeaconCoordinator.Contract.LatestConfigDetails(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRFBeaconCoordinator.Contract.LatestConfigDigestAndEpoch(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRFBeaconCoordinator.Contract.LatestConfigDigestAndEpoch(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.LinkAvailableForPayment(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.LinkAvailableForPayment(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) MaxErrorMsgLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "maxErrorMsgLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) MaxErrorMsgLength() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.MaxErrorMsgLength(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) MaxErrorMsgLength() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.MaxErrorMsgLength(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) MaxNumWords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "maxNumWords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.MaxNumWords(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.MaxNumWords(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) MinDelay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "minDelay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) MinDelay() (uint16, error) {
	return _VRFBeaconCoordinator.Contract.MinDelay(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) MinDelay() (uint16, error) {
	return _VRFBeaconCoordinator.Contract.MinDelay(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _VRFBeaconCoordinator.Contract.OracleObservationCount(&_VRFBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _VRFBeaconCoordinator.Contract.OracleObservationCount(&_VRFBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.OwedPayment(&_VRFBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.OwedPayment(&_VRFBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) Owner() (common.Address, error) {
	return _VRFBeaconCoordinator.Contract.Owner(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) Owner() (common.Address, error) {
	return _VRFBeaconCoordinator.Contract.Owner(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) SKeyID() ([32]byte, error) {
	return _VRFBeaconCoordinator.Contract.SKeyID(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) SKeyID() ([32]byte, error) {
	return _VRFBeaconCoordinator.Contract.SKeyID(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeaconCoordinator.Contract.SProvingKeyHash(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeaconCoordinator.Contract.SProvingKeyHash(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) TypeAndVersion() (string, error) {
	return _VRFBeaconCoordinator.Contract.TypeAndVersion(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) TypeAndVersion() (string, error) {
	return _VRFBeaconCoordinator.Contract.TypeAndVersion(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.AcceptOwnership(&_VRFBeaconCoordinator.TransactOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.AcceptOwnership(&_VRFBeaconCoordinator.TransactOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.AcceptPayeeship(&_VRFBeaconCoordinator.TransactOpts, transmitter)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.AcceptPayeeship(&_VRFBeaconCoordinator.TransactOpts, transmitter)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) ExposeType(opts *bind.TransactOpts, arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "exposeType", arg0)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.ExposeType(&_VRFBeaconCoordinator.TransactOpts, arg0)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.ExposeType(&_VRFBeaconCoordinator.TransactOpts, arg0)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) GetRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "getRandomness", requestID)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.GetRandomness(&_VRFBeaconCoordinator.TransactOpts, requestID)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.GetRandomness(&_VRFBeaconCoordinator.TransactOpts, requestID)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "keyGenerated", kd)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.KeyGenerated(&_VRFBeaconCoordinator.TransactOpts, kd)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.KeyGenerated(&_VRFBeaconCoordinator.TransactOpts, kd)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "newKeyRequested")
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.NewKeyRequested(&_VRFBeaconCoordinator.TransactOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.NewKeyRequested(&_VRFBeaconCoordinator.TransactOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) RequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "requestRandomness", numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.RequestRandomness(&_VRFBeaconCoordinator.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.RequestRandomness(&_VRFBeaconCoordinator.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.RequestRandomnessFulfillment(&_VRFBeaconCoordinator.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.RequestRandomnessFulfillment(&_VRFBeaconCoordinator.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetBilling(&_VRFBeaconCoordinator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetBilling(&_VRFBeaconCoordinator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetBillingAccessController(&_VRFBeaconCoordinator.TransactOpts, _billingAccessController)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetBillingAccessController(&_VRFBeaconCoordinator.TransactOpts, _billingAccessController)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetConfig(&_VRFBeaconCoordinator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetConfig(&_VRFBeaconCoordinator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetPayees(&_VRFBeaconCoordinator.TransactOpts, transmitters, payees)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.SetPayees(&_VRFBeaconCoordinator.TransactOpts, transmitters, payees)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.TransferOwnership(&_VRFBeaconCoordinator.TransactOpts, to)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.TransferOwnership(&_VRFBeaconCoordinator.TransactOpts, to)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.TransferPayeeship(&_VRFBeaconCoordinator.TransactOpts, transmitter, proposed)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.TransferPayeeship(&_VRFBeaconCoordinator.TransactOpts, transmitter, proposed)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.Transmit(&_VRFBeaconCoordinator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.Transmit(&_VRFBeaconCoordinator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.WithdrawFunds(&_VRFBeaconCoordinator.TransactOpts, recipient, amount)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.WithdrawFunds(&_VRFBeaconCoordinator.TransactOpts, recipient, amount)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.WithdrawPayment(&_VRFBeaconCoordinator.TransactOpts, transmitter)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.WithdrawPayment(&_VRFBeaconCoordinator.TransactOpts, transmitter)
}

type VRFBeaconCoordinatorBillingAccessControllerSetIterator struct {
	Event *VRFBeaconCoordinatorBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorBillingAccessControllerSet)
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
		it.Event = new(VRFBeaconCoordinatorBillingAccessControllerSet)
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

func (it *VRFBeaconCoordinatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*VRFBeaconCoordinatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorBillingAccessControllerSetIterator{contract: _VRFBeaconCoordinator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorBillingAccessControllerSet)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*VRFBeaconCoordinatorBillingAccessControllerSet, error) {
	event := new(VRFBeaconCoordinatorBillingAccessControllerSet)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorBillingSetIterator struct {
	Event *VRFBeaconCoordinatorBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorBillingSet)
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
		it.Event = new(VRFBeaconCoordinatorBillingSet)
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

func (it *VRFBeaconCoordinatorBillingSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*VRFBeaconCoordinatorBillingSetIterator, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorBillingSetIterator{contract: _VRFBeaconCoordinator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorBillingSet)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseBillingSet(log types.Log) (*VRFBeaconCoordinatorBillingSet, error) {
	event := new(VRFBeaconCoordinatorBillingSet)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorConfigSetIterator struct {
	Event *VRFBeaconCoordinatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorConfigSet)
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
		it.Event = new(VRFBeaconCoordinatorConfigSet)
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

func (it *VRFBeaconCoordinatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorConfigSet struct {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFBeaconCoordinatorConfigSetIterator, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorConfigSetIterator{contract: _VRFBeaconCoordinator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorConfigSet)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseConfigSet(log types.Log) (*VRFBeaconCoordinatorConfigSet, error) {
	event := new(VRFBeaconCoordinatorConfigSet)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorNewTransmissionIterator struct {
	Event *VRFBeaconCoordinatorNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorNewTransmission)
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
		it.Event = new(VRFBeaconCoordinatorNewTransmission)
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

func (it *VRFBeaconCoordinatorNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorNewTransmission struct {
	AggregatorRoundId uint32
	Transmitter       common.Address
	JuelsPerFeeCoin   *big.Int
	ConfigDigest      [32]byte
	EpochAndRound     *big.Int
	OutputsServed     []VRFBeaconReportOutputServed
	Raw               types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*VRFBeaconCoordinatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorNewTransmissionIterator{contract: _VRFBeaconCoordinator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorNewTransmission)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseNewTransmission(log types.Log) (*VRFBeaconCoordinatorNewTransmission, error) {
	event := new(VRFBeaconCoordinatorNewTransmission)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorOraclePaidIterator struct {
	Event *VRFBeaconCoordinatorOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorOraclePaid)
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
		it.Event = new(VRFBeaconCoordinatorOraclePaid)
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

func (it *VRFBeaconCoordinatorOraclePaidIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*VRFBeaconCoordinatorOraclePaidIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorOraclePaidIterator{contract: _VRFBeaconCoordinator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorOraclePaid)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseOraclePaid(log types.Log) (*VRFBeaconCoordinatorOraclePaid, error) {
	event := new(VRFBeaconCoordinatorOraclePaid)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorOwnershipTransferRequestedIterator struct {
	Event *VRFBeaconCoordinatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorOwnershipTransferRequested)
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
		it.Event = new(VRFBeaconCoordinatorOwnershipTransferRequested)
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

func (it *VRFBeaconCoordinatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconCoordinatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorOwnershipTransferRequestedIterator{contract: _VRFBeaconCoordinator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorOwnershipTransferRequested)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconCoordinatorOwnershipTransferRequested, error) {
	event := new(VRFBeaconCoordinatorOwnershipTransferRequested)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorOwnershipTransferredIterator struct {
	Event *VRFBeaconCoordinatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorOwnershipTransferred)
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
		it.Event = new(VRFBeaconCoordinatorOwnershipTransferred)
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

func (it *VRFBeaconCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconCoordinatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorOwnershipTransferredIterator{contract: _VRFBeaconCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorOwnershipTransferred)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*VRFBeaconCoordinatorOwnershipTransferred, error) {
	event := new(VRFBeaconCoordinatorOwnershipTransferred)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorPayeeshipTransferRequestedIterator struct {
	Event *VRFBeaconCoordinatorPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorPayeeshipTransferRequested)
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
		it.Event = new(VRFBeaconCoordinatorPayeeshipTransferRequested)
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

func (it *VRFBeaconCoordinatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*VRFBeaconCoordinatorPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorPayeeshipTransferRequestedIterator{contract: _VRFBeaconCoordinator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorPayeeshipTransferRequested)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*VRFBeaconCoordinatorPayeeshipTransferRequested, error) {
	event := new(VRFBeaconCoordinatorPayeeshipTransferRequested)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorPayeeshipTransferredIterator struct {
	Event *VRFBeaconCoordinatorPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorPayeeshipTransferred)
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
		it.Event = new(VRFBeaconCoordinatorPayeeshipTransferred)
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

func (it *VRFBeaconCoordinatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*VRFBeaconCoordinatorPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorPayeeshipTransferredIterator{contract: _VRFBeaconCoordinator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorPayeeshipTransferred)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParsePayeeshipTransferred(log types.Log) (*VRFBeaconCoordinatorPayeeshipTransferred, error) {
	event := new(VRFBeaconCoordinatorPayeeshipTransferred)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorRandomWordsFulfilledIterator struct {
	Event *VRFBeaconCoordinatorRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorRandomWordsFulfilled)
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
		it.Event = new(VRFBeaconCoordinatorRandomWordsFulfilled)
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

func (it *VRFBeaconCoordinatorRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFBeaconCoordinatorRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorRandomWordsFulfilledIterator{contract: _VRFBeaconCoordinator.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorRandomWordsFulfilled)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFBeaconCoordinatorRandomWordsFulfilled, error) {
	event := new(VRFBeaconCoordinatorRandomWordsFulfilled)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorRandomnessFulfillmentRequestedIterator struct {
	Event *VRFBeaconCoordinatorRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorRandomnessFulfillmentRequested)
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
		it.Event = new(VRFBeaconCoordinatorRandomnessFulfillmentRequested)
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

func (it *VRFBeaconCoordinatorRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorRandomnessFulfillmentRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  uint64
	Callback               VRFBeaconTypesCallback
	Raw                    types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts) (*VRFBeaconCoordinatorRandomnessFulfillmentRequestedIterator, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconCoordinator.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorRandomnessFulfillmentRequested) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorRandomnessFulfillmentRequested)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconCoordinatorRandomnessFulfillmentRequested, error) {
	event := new(VRFBeaconCoordinatorRandomnessFulfillmentRequested)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconCoordinatorRandomnessRequestedIterator struct {
	Event *VRFBeaconCoordinatorRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconCoordinatorRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconCoordinatorRandomnessRequested)
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
		it.Event = new(VRFBeaconCoordinatorRandomnessRequested)
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

func (it *VRFBeaconCoordinatorRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconCoordinatorRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconCoordinatorRandomnessRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	Raw                    types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, nextBeaconOutputHeight []uint64) (*VRFBeaconCoordinatorRandomnessRequestedIterator, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorRandomnessRequestedIterator{contract: _VRFBeaconCoordinator.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorRandomnessRequested, nextBeaconOutputHeight []uint64) (event.Subscription, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconCoordinatorRandomnessRequested)
				if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) ParseRandomnessRequested(log types.Log) (*VRFBeaconCoordinatorRandomnessRequested, error) {
	event := new(VRFBeaconCoordinatorRandomnessRequested)
	if err := _VRFBeaconCoordinator.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFBeaconDKGClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDKG\",\"name\":\"_keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf2732c7": "keyGenerated((bytes,bytes32[]))",
		"55e48749": "newKeyRequested()",
		"cc31f7dd": "s_keyID()",
		"d57fc45a": "s_provingKeyHash()",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516103b93803806103b983398101604081905261002f91610058565b600080546001600160a01b0319166001600160a01b039390931692909217909155600155610092565b6000806040838503121561006b57600080fd5b82516001600160a01b038116811461008257600080fd5b6020939093015192949293505050565b610318806100a16000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806355e4874914610051578063bf2732c71461005d578063cc31f7dd14610070578063d57fc45a1461008b575b600080fd5b61005b6000600255565b005b61005b61006b3660046101c4565b610094565b61007960015481565b60405190815260200160405180910390f35b61007960025481565b60005481516040516001600160a01b03909216916100b591906020016102a7565b60408051601f1981840301815291905280516020909101206002555050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561010d5761010d6100d4565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561013c5761013c6100d4565b604052919050565b600082601f83011261015557600080fd5b8135602067ffffffffffffffff821115610171576101716100d4565b8160051b610180828201610113565b928352848101820192828101908785111561019a57600080fd5b83870192505b848310156101b9578235825291830191908301906101a0565b979650505050505050565b600060208083850312156101d757600080fd5b823567ffffffffffffffff808211156101ef57600080fd5b908401906040828703121561020357600080fd5b61020b6100ea565b82358281111561021a57600080fd5b8301601f8101881361022b57600080fd5b80358381111561023d5761023d6100d4565b61024f601f8201601f19168701610113565b818152898783850101111561026357600080fd5b81878401888301376000878383010152808452505050838301358281111561028a57600080fd5b61029688828601610144565b948201949094529695505050505050565b6000825160005b818110156102c857602081860181015185830152016102ae565b818111156102d7576000828501525b50919091019291505056fea2646970667358221220623a3e4b9910b90387ed5d47b1b013d28a10c0f56f800310c82b0d92ec1988dc64736f6c634300080d0033",
}

var VRFBeaconDKGClientABI = VRFBeaconDKGClientMetaData.ABI

var VRFBeaconDKGClientFuncSigs = VRFBeaconDKGClientMetaData.Sigs

var VRFBeaconDKGClientBin = VRFBeaconDKGClientMetaData.Bin

func DeployVRFBeaconDKGClient(auth *bind.TransactOpts, backend bind.ContractBackend, _keyProvider common.Address, _keyID [32]byte) (common.Address, *types.Transaction, *VRFBeaconDKGClient, error) {
	parsed, err := VRFBeaconDKGClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconDKGClientBin), backend, _keyProvider, _keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeaconDKGClient{VRFBeaconDKGClientCaller: VRFBeaconDKGClientCaller{contract: contract}, VRFBeaconDKGClientTransactor: VRFBeaconDKGClientTransactor{contract: contract}, VRFBeaconDKGClientFilterer: VRFBeaconDKGClientFilterer{contract: contract}}, nil
}

type VRFBeaconDKGClient struct {
	VRFBeaconDKGClientCaller
	VRFBeaconDKGClientTransactor
	VRFBeaconDKGClientFilterer
}

type VRFBeaconDKGClientCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconDKGClientTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconDKGClientFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconDKGClientSession struct {
	Contract     *VRFBeaconDKGClient
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconDKGClientCallerSession struct {
	Contract *VRFBeaconDKGClientCaller
	CallOpts bind.CallOpts
}

type VRFBeaconDKGClientTransactorSession struct {
	Contract     *VRFBeaconDKGClientTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconDKGClientRaw struct {
	Contract *VRFBeaconDKGClient
}

type VRFBeaconDKGClientCallerRaw struct {
	Contract *VRFBeaconDKGClientCaller
}

type VRFBeaconDKGClientTransactorRaw struct {
	Contract *VRFBeaconDKGClientTransactor
}

func NewVRFBeaconDKGClient(address common.Address, backend bind.ContractBackend) (*VRFBeaconDKGClient, error) {
	contract, err := bindVRFBeaconDKGClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconDKGClient{VRFBeaconDKGClientCaller: VRFBeaconDKGClientCaller{contract: contract}, VRFBeaconDKGClientTransactor: VRFBeaconDKGClientTransactor{contract: contract}, VRFBeaconDKGClientFilterer: VRFBeaconDKGClientFilterer{contract: contract}}, nil
}

func NewVRFBeaconDKGClientCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconDKGClientCaller, error) {
	contract, err := bindVRFBeaconDKGClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconDKGClientCaller{contract: contract}, nil
}

func NewVRFBeaconDKGClientTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconDKGClientTransactor, error) {
	contract, err := bindVRFBeaconDKGClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconDKGClientTransactor{contract: contract}, nil
}

func NewVRFBeaconDKGClientFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconDKGClientFilterer, error) {
	contract, err := bindVRFBeaconDKGClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconDKGClientFilterer{contract: contract}, nil
}

func bindVRFBeaconDKGClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconDKGClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconDKGClient.Contract.VRFBeaconDKGClientCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.VRFBeaconDKGClientTransactor.contract.Transfer(opts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.VRFBeaconDKGClientTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconDKGClient.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.contract.Transfer(opts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeaconDKGClient.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientSession) SKeyID() ([32]byte, error) {
	return _VRFBeaconDKGClient.Contract.SKeyID(&_VRFBeaconDKGClient.CallOpts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientCallerSession) SKeyID() ([32]byte, error) {
	return _VRFBeaconDKGClient.Contract.SKeyID(&_VRFBeaconDKGClient.CallOpts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeaconDKGClient.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeaconDKGClient.Contract.SProvingKeyHash(&_VRFBeaconDKGClient.CallOpts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeaconDKGClient.Contract.SProvingKeyHash(&_VRFBeaconDKGClient.CallOpts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.contract.Transact(opts, "keyGenerated", kd)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.KeyGenerated(&_VRFBeaconDKGClient.TransactOpts, kd)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.KeyGenerated(&_VRFBeaconDKGClient.TransactOpts, kd)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconDKGClient.contract.Transact(opts, "newKeyRequested")
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.NewKeyRequested(&_VRFBeaconDKGClient.TransactOpts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeaconDKGClient.Contract.NewKeyRequested(&_VRFBeaconDKGClient.TransactOpts)
}

var VRFBeaconExternalAPIMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2f7527cc": "NUM_CONF_DELAYS()",
		"0b93e168": "getRandomness(uint48)",
		"cf7e754a": "i_StartSlot()",
		"cd0593df": "i_beaconPeriodBlocks()",
		"bbcdd0d8": "maxNumWords()",
		"c63c4e9b": "minDelay()",
		"dc92accf": "requestRandomness(uint16,uint64,uint24)",
		"f645dcb1": "requestRandomnessFulfillment(uint64,uint16,uint24,uint32,bytes)",
	},
	Bin: "0x60c06040523480156200001157600080fd5b5060405162000fcb38038062000fcb833981016040819052620000349162000098565b806000036200005657604051632abc297960e01b815260040160405180910390fd5b60808190526000620000698243620000b2565b90506000816080516200007d9190620000eb565b90506200008b814362000105565b60a0525062000120915050565b600060208284031215620000ab57600080fd5b5051919050565b600082620000d057634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b600082821015620001005762000100620000d5565b500390565b600082198211156200011b576200011b620000d5565b500190565b60805160a051610e6962000162600039600061012e01526000818161010701528181610255015281816107a6015281816107d5015261080d0152610e696000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063cd0593df1161005b578063cd0593df14610102578063cf7e754a14610129578063dc92accf14610150578063f645dcb11461017a57600080fd5b80630b93e1681461008d5780632f7527cc146100b6578063bbcdd0d8146100d0578063c63c4e9b146100e7575b600080fd5b6100a061009b3660046109dc565b61018d565b6040516100ad9190610a0b565b60405180910390f35b6100be600881565b60405160ff90911681526020016100ad565b6100d96103e881565b6040519081526020016100ad565b6100ef600381565b60405161ffff90911681526020016100ad565b6100d97f000000000000000000000000000000000000000000000000000000000000000081565b6100d97f000000000000000000000000000000000000000000000000000000000000000081565b61016361015e366004610a91565b610329565b60405165ffffffffffff90911681526020016100ad565b610163610188366004610aea565b610445565b65ffffffffffff811660008181526004602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff600160381b820416938201939093526001600160a01b03600160481b84048116606083810191825298909752949093526001600160e81b03199091169055915116331461024a576060810151604051638e30e82360e01b81526001600160a01b0390911660048201523360248201526044015b60405180910390fd5b8051600090610280907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16610bff565b90506000826020015162ffffff16436102999190610c1e565b90508082106102c4576040516315ad27c360e01b815260048101839052436024820152604401610241565b67ffffffffffffffff8211156102f0576040516302c6ef8160e11b815260048101839052602401610241565b60008281526001602090815260408083208287015162ffffff16845290915290205461032090869085908561054a565b95945050505050565b6000806000806103398786610720565b92509250925065ffffffffffff831660009081526004602090815260409182902084518154928601518487015160608801516001600160a01b0316600160481b027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff909216600160381b0291909116670100000000000000600160e81b031962ffffff9093166401000000000266ffffffffffffff1990961663ffffffff909416939093179490941716179190911790555167ffffffffffffffff8216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc59061043290889062ffffff91909116815260200190565b60405180910390a2509095945050505050565b60008060006104548787610720565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a67ffffffffffffffff1681526020018763ffffffff166bffffffffffffffffffffffff16815250905081878a836040516020016104d49493929190610c35565b60408051601f19818403018152828252805160209182012065ffffffffffff8716600090815291829052919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906105359084908a908d908690610c35565b60405180910390a15090979650505050505050565b6060826105845760405163c7d41b1b60e01b815265ffffffffffff8616600482015267ffffffffffffffff83166024820152604401610241565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff16111561063a576040808601519051634a90778560e01b815261ffff90911660048201526103e86024820152604401610241565b6000856040015161ffff1667ffffffffffffffff81111561065d5761065d610ad4565b604051908082528060200260200182016040528015610686578160200160208202803683370190505b50905060005b866040015161ffff168161ffff1610156107155782816040516020016106c992919091825260f01b6001600160f01b031916602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff16815181106106f8576106f8610d20565b60209081029190910101528061070d81610d36565b91505061068c565b509695505050505050565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff16111561077a57604051634a90778560e01b815261ffff861660048201526103e86024820152604401610241565b8461ffff1660000361079f576040516308fad2a760e01b815260040160405180910390fd5b60006107cb7f000000000000000000000000000000000000000000000000000000000000000043610d6d565b90506000816107fa7f000000000000000000000000000000000000000000000000000000000000000043610d81565b6108049190610c1e565b905060006108327f000000000000000000000000000000000000000000000000000000000000000083610d99565b905063ffffffff8110610858576040516307b2a52360e41b815260040160405180910390fd5b6040805180820182526002805465ffffffffffff16825282516101008101938490528493600093929160208401916003906008908288855b82829054906101000a900462ffffff1662ffffff168152602001906003019060208260020104928301926001038202915080841161089057905050505091909252505081519192505065ffffffffffff8082161061090157604051630568cab760e31b815260040160405180910390fd5b61090c816001610dad565b6002805465ffffffffffff191665ffffffffffff9290921691909117905560005b6008811015610973578a62ffffff168360200151826008811061095257610952610d20565b602002015162ffffff1614610973578061096b81610dd7565b91505061092d565b6008811061099b576020830151604051630c4f769b60e41b8152610241918d91600401610df0565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b6000602082840312156109ee57600080fd5b813565ffffffffffff81168114610a0457600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610a4357835183529284019291840191600101610a27565b50909695505050505050565b803561ffff81168114610a6157600080fd5b919050565b803567ffffffffffffffff81168114610a6157600080fd5b803562ffffff81168114610a6157600080fd5b600080600060608486031215610aa657600080fd5b610aaf84610a4f565b9250610abd60208501610a66565b9150610acb60408501610a7e565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b600080600080600060a08688031215610b0257600080fd5b610b0b86610a66565b9450610b1960208701610a4f565b9350610b2760408701610a7e565b9250606086013563ffffffff81168114610b4057600080fd5b9150608086013567ffffffffffffffff80821115610b5d57600080fd5b818801915088601f830112610b7157600080fd5b813581811115610b8357610b83610ad4565b604051601f8201601f19908116603f01168101908382118183101715610bab57610bab610ad4565b816040528281528b6020848701011115610bc457600080fd5b8260208601602083013760006020848301015280955050505050509295509295909350565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610c1957610c19610be9565b500290565b600082821015610c3057610c30610be9565b500390565b600067ffffffffffffffff8087168352602062ffffff87168185015281861660408501526080606085015265ffffffffffff855116608085015261ffff818601511660a085015260018060a01b0360408601511660c08501526060850151915060c060e085015281518061014086015260005b81811015610cc55783810183015186820161016001528201610ca8565b81811115610cd857600061016083880101525b50608086015167ffffffffffffffff1661010086015260a0909501516bffffffffffffffffffffffff16610120850152505050601f909101601f191601610160019392505050565b634e487b7160e01b600052603260045260246000fd5b600061ffff808316818103610d4d57610d4d610be9565b6001019392505050565b634e487b7160e01b600052601260045260246000fd5b600082610d7c57610d7c610d57565b500690565b60008219821115610d9457610d94610be9565b500190565b600082610da857610da8610d57565b500490565b600065ffffffffffff808316818516808303821115610dce57610dce610be9565b01949350505050565b600060018201610de957610de9610be9565b5060010190565b62ffffff838116825261012082019060208084018560005b6008811015610e27578151851683529183019190830190600101610e08565b5050505050939250505056fea264697066735822122054425f95f8e36f7e09b2b7818d3e828c46a2a42e85ba10e13bcb6145117ccbd064736f6c634300080d0033",
}

var VRFBeaconExternalAPIABI = VRFBeaconExternalAPIMetaData.ABI

var VRFBeaconExternalAPIFuncSigs = VRFBeaconExternalAPIMetaData.Sigs

var VRFBeaconExternalAPIBin = VRFBeaconExternalAPIMetaData.Bin

func DeployVRFBeaconExternalAPI(auth *bind.TransactOpts, backend bind.ContractBackend, beaconPeriodBlocksArg *big.Int) (common.Address, *types.Transaction, *VRFBeaconExternalAPI, error) {
	parsed, err := VRFBeaconExternalAPIMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconExternalAPIBin), backend, beaconPeriodBlocksArg)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeaconExternalAPI{VRFBeaconExternalAPICaller: VRFBeaconExternalAPICaller{contract: contract}, VRFBeaconExternalAPITransactor: VRFBeaconExternalAPITransactor{contract: contract}, VRFBeaconExternalAPIFilterer: VRFBeaconExternalAPIFilterer{contract: contract}}, nil
}

type VRFBeaconExternalAPI struct {
	VRFBeaconExternalAPICaller
	VRFBeaconExternalAPITransactor
	VRFBeaconExternalAPIFilterer
}

type VRFBeaconExternalAPICaller struct {
	contract *bind.BoundContract
}

type VRFBeaconExternalAPITransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconExternalAPIFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconExternalAPISession struct {
	Contract     *VRFBeaconExternalAPI
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconExternalAPICallerSession struct {
	Contract *VRFBeaconExternalAPICaller
	CallOpts bind.CallOpts
}

type VRFBeaconExternalAPITransactorSession struct {
	Contract     *VRFBeaconExternalAPITransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconExternalAPIRaw struct {
	Contract *VRFBeaconExternalAPI
}

type VRFBeaconExternalAPICallerRaw struct {
	Contract *VRFBeaconExternalAPICaller
}

type VRFBeaconExternalAPITransactorRaw struct {
	Contract *VRFBeaconExternalAPITransactor
}

func NewVRFBeaconExternalAPI(address common.Address, backend bind.ContractBackend) (*VRFBeaconExternalAPI, error) {
	contract, err := bindVRFBeaconExternalAPI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPI{VRFBeaconExternalAPICaller: VRFBeaconExternalAPICaller{contract: contract}, VRFBeaconExternalAPITransactor: VRFBeaconExternalAPITransactor{contract: contract}, VRFBeaconExternalAPIFilterer: VRFBeaconExternalAPIFilterer{contract: contract}}, nil
}

func NewVRFBeaconExternalAPICaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconExternalAPICaller, error) {
	contract, err := bindVRFBeaconExternalAPI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPICaller{contract: contract}, nil
}

func NewVRFBeaconExternalAPITransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconExternalAPITransactor, error) {
	contract, err := bindVRFBeaconExternalAPI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPITransactor{contract: contract}, nil
}

func NewVRFBeaconExternalAPIFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconExternalAPIFilterer, error) {
	contract, err := bindVRFBeaconExternalAPI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPIFilterer{contract: contract}, nil
}

func bindVRFBeaconExternalAPI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconExternalAPIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconExternalAPI.Contract.VRFBeaconExternalAPICaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.VRFBeaconExternalAPITransactor.contract.Transfer(opts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.VRFBeaconExternalAPITransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconExternalAPI.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.contract.Transfer(opts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFBeaconExternalAPI.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconExternalAPI.Contract.NUMCONFDELAYS(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconExternalAPI.Contract.NUMCONFDELAYS(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICaller) IStartSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconExternalAPI.contract.Call(opts, &out, "i_StartSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconExternalAPI.Contract.IStartSlot(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerSession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconExternalAPI.Contract.IStartSlot(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconExternalAPI.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconExternalAPI.Contract.IBeaconPeriodBlocks(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconExternalAPI.Contract.IBeaconPeriodBlocks(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICaller) MaxNumWords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconExternalAPI.contract.Call(opts, &out, "maxNumWords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconExternalAPI.Contract.MaxNumWords(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerSession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconExternalAPI.Contract.MaxNumWords(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICaller) MinDelay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFBeaconExternalAPI.contract.Call(opts, &out, "minDelay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) MinDelay() (uint16, error) {
	return _VRFBeaconExternalAPI.Contract.MinDelay(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerSession) MinDelay() (uint16, error) {
	return _VRFBeaconExternalAPI.Contract.MinDelay(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactor) GetRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.contract.Transact(opts, "getRandomness", requestID)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.GetRandomness(&_VRFBeaconExternalAPI.TransactOpts, requestID)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.GetRandomness(&_VRFBeaconExternalAPI.TransactOpts, requestID)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactor) RequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.contract.Transact(opts, "requestRandomness", numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.RequestRandomness(&_VRFBeaconExternalAPI.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.RequestRandomness(&_VRFBeaconExternalAPI.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.RequestRandomnessFulfillment(&_VRFBeaconExternalAPI.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.RequestRandomnessFulfillment(&_VRFBeaconExternalAPI.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

type VRFBeaconExternalAPIRandomnessFulfillmentRequestedIterator struct {
	Event *VRFBeaconExternalAPIRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconExternalAPIRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconExternalAPIRandomnessFulfillmentRequested)
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
		it.Event = new(VRFBeaconExternalAPIRandomnessFulfillmentRequested)
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

func (it *VRFBeaconExternalAPIRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconExternalAPIRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconExternalAPIRandomnessFulfillmentRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  uint64
	Callback               VRFBeaconTypesCallback
	Raw                    types.Log
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts) (*VRFBeaconExternalAPIRandomnessFulfillmentRequestedIterator, error) {

	logs, sub, err := _VRFBeaconExternalAPI.contract.FilterLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPIRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconExternalAPI.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconExternalAPIRandomnessFulfillmentRequested) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconExternalAPI.contract.WatchLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconExternalAPIRandomnessFulfillmentRequested)
				if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconExternalAPIRandomnessFulfillmentRequested, error) {
	event := new(VRFBeaconExternalAPIRandomnessFulfillmentRequested)
	if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconExternalAPIRandomnessRequestedIterator struct {
	Event *VRFBeaconExternalAPIRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconExternalAPIRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconExternalAPIRandomnessRequested)
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
		it.Event = new(VRFBeaconExternalAPIRandomnessRequested)
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

func (it *VRFBeaconExternalAPIRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconExternalAPIRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconExternalAPIRandomnessRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	Raw                    types.Log
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, nextBeaconOutputHeight []uint64) (*VRFBeaconExternalAPIRandomnessRequestedIterator, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconExternalAPI.contract.FilterLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPIRandomnessRequestedIterator{contract: _VRFBeaconExternalAPI.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconExternalAPIRandomnessRequested, nextBeaconOutputHeight []uint64) (event.Subscription, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconExternalAPI.contract.WatchLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconExternalAPIRandomnessRequested)
				if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) ParseRandomnessRequested(log types.Log) (*VRFBeaconExternalAPIRandomnessRequested, error) {
	event := new(VRFBeaconExternalAPIRandomnessRequested)
	if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFBeaconOCRMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OffchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"occVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconReport.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconReport.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxErrorMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2f7527cc": "NUM_CONF_DELAYS()",
		"79ba5097": "acceptOwnership()",
		"b121e147": "acceptPayeeship(address)",
		"c278e5b7": "exposeType(((uint64,uint24,(uint256[2]),((uint48,uint16,address,bytes,uint64,uint96),uint96)[])[],uint192,uint64,bytes32))",
		"29937268": "getBilling()",
		"c4c92b37": "getBillingAccessController()",
		"0b93e168": "getRandomness(uint48)",
		"cf7e754a": "i_StartSlot()",
		"cd0593df": "i_beaconPeriodBlocks()",
		"81ff7048": "latestConfigDetails()",
		"afcb95d7": "latestConfigDigestAndEpoch()",
		"d09dc339": "linkAvailableForPayment()",
		"7a464944": "maxErrorMsgLength()",
		"bbcdd0d8": "maxNumWords()",
		"c63c4e9b": "minDelay()",
		"e4902f82": "oracleObservationCount(address)",
		"0eafb25b": "owedPayment(address)",
		"8da5cb5b": "owner()",
		"dc92accf": "requestRandomness(uint16,uint64,uint24)",
		"f645dcb1": "requestRandomnessFulfillment(uint64,uint16,uint24,uint32,bytes)",
		"643dc105": "setBilling(uint32,uint32,uint32,uint32,uint24)",
		"fbffd2c1": "setBillingAccessController(address)",
		"e3d0e712": "setConfig(address[],address[],uint8,bytes,uint64,bytes)",
		"9c849b30": "setPayees(address[],address[])",
		"f2fde38b": "transferOwnership(address)",
		"eb5dcd6c": "transferPayeeship(address,address)",
		"b1dc65a4": "transmit(bytes32[3],bytes,bytes32[],bytes32[],bytes32)",
		"181f5a77": "typeAndVersion()",
		"c1075329": "withdrawFunds(address,uint256)",
		"8ac28d5a": "withdrawPayment(address)",
	},
	Bin: "0x60c06040523480156200001157600080fd5b506040516200534638038062005346833981016040819052620000349162000200565b338060008480806000036200005c57604051632abc297960e01b815260040160405180910390fd5b608081905260006200006f82436200023f565b905060008160805162000083919062000278565b905062000091814362000292565b60a0525050506001600160a01b0383169050620000f55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600680546001600160a01b0319166001600160a01b03848116919091179091558116156200012857620001288162000154565b5050601780546001600160a01b0319166001600160a01b03939093169290921790915550620002ad9050565b336001600160a01b03821603620001ae5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ec565b600780546001600160a01b0319166001600160a01b03838116918217909255600654604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b600080604083850312156200021457600080fd5b825160208401519092506001600160a01b03811681146200023457600080fd5b809150509250929050565b6000826200025d57634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b6000828210156200028d576200028d62000262565b500390565b60008219821115620002a857620002a862000262565b500190565b60805160a051615050620002f6600039600061046a0152600081816104430152818161060d01528181613018015281816130470152818161307f015261394401526150506000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c8063bbcdd0d811610104578063d09dc339116100a2578063eb5dcd6c11610071578063eb5dcd6c146104f9578063f2fde38b1461050c578063f645dcb11461051f578063fbffd2c11461053257600080fd5b8063d09dc3391461048c578063dc92accf14610494578063e3d0e712146104be578063e4902f82146104d157600080fd5b8063c4c92b37116100de578063c4c92b3714610412578063c63c4e9b14610423578063cd0593df1461043e578063cf7e754a1461046557600080fd5b8063bbcdd0d8146103e5578063c1075329146103ee578063c278e5b71461040157600080fd5b80637a4649441161017c5780639c849b301161014b5780639c849b3014610382578063afcb95d714610395578063b121e147146103bf578063b1dc65a4146103d257600080fd5b80637a4649441461031657806381ff70481461031e5780638ac28d5a1461034a5780638da5cb5b1461035d57600080fd5b806329937268116101b857806329937268146102605780632f7527cc146102df578063643dc105146102f957806379ba50971461030e57600080fd5b80630b93e168146101df5780630eafb25b14610208578063181f5a7714610229575b600080fd5b6101f26101ed366004613d7e565b610545565b6040516101ff9190613dd6565b60405180910390f35b61021b610216366004613dfe565b6106e0565b6040519081526020016101ff565b6040805180820182526015815274565246426561636f6e20312e302e302d616c70686160581b602082015290516101ff9190613e73565b6102a3600554600160501b810463ffffffff90811692600160701b8304821692600160901b8104831692600160b01b82041691600160d01b90910462ffffff1690565b6040805163ffffffff9687168152948616602086015292851692840192909252909216606082015262ffffff909116608082015260a0016101ff565b6102e7600881565b60405160ff90911681526020016101ff565b61030c610307366004613eb0565b6107e5565b005b61030c6109cb565b61021b608081565b6007546009546040805160008152600160c01b90930463ffffffff1660208401528201526060016101ff565b61030c610358366004613dfe565b610a79565b6006546001600160a01b03165b6040516001600160a01b0390911681526020016101ff565b61030c610390366004613f64565b610aeb565b600954600b546040805160008152602081019390935263ffffffff909116908201526060016101ff565b61030c6103cd366004613dfe565b610cbd565b61030c6103e0366004614010565b610d99565b61021b6103e881565b61030c6103fc3660046140c6565b61121c565b61030c61040f3660046140f2565b50565b6016546001600160a01b031661036a565b61042b600381565b60405161ffff90911681526020016101ff565b61021b7f000000000000000000000000000000000000000000000000000000000000000081565b61021b7f000000000000000000000000000000000000000000000000000000000000000081565b61021b61146d565b6104a76104a236600461415c565b6114fd565b60405165ffffffffffff90911681526020016101ff565b61030c6104cc3660046141b8565b61161a565b6104e46104df366004613dfe565b611d46565b60405163ffffffff90911681526020016101ff565b61030c6105073660046142a5565b611df5565b61030c61051a366004613dfe565b611f2e565b6104a761052d3660046143d9565b611f3f565b61030c610540366004613dfe565b61203e565b65ffffffffffff811660008181526004602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff600160381b820416938201939093526001600160a01b03600160481b84048116606083810191825298909752949093526001600160e81b031990911690559151163314610602576060810151604051638e30e82360e01b81526001600160a01b0390911660048201523360248201526044015b60405180910390fd5b8051600090610638907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff166144b7565b90506000826020015162ffffff164361065191906144d6565b905080821061067c576040516315ad27c360e01b8152600481018390524360248201526044016105f9565b6001600160401b038211156106a7576040516302c6ef8160e11b8152600481018390526024016105f9565b60008281526001602090815260408083208287015162ffffff1684529091529020546106d790869085908561204f565b95945050505050565b6001600160a01b0381166000908152600c602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906107425750600092915050565b6005546020820151600091600160901b900463ffffffff169060109060ff16601f8110610771576107716144ed565b6008810491909101546005546107a4926007166004026101000a90910463ffffffff90811691600160301b900416614503565b63ffffffff166107b491906144b7565b6107c290633b9aca006144b7565b905081604001516001600160601b0316816107dd9190614528565b949350505050565b6016546001600160a01b03166108036006546001600160a01b031690565b6001600160a01b0316336001600160a01b0316148061088f5750604051630d629b5f60e31b81526001600160a01b03821690636b14daf89061084e9033906000903690600401614569565b602060405180830381865afa15801561086b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088f919061458e565b6108db5760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c60448201526064016105f9565b6108e3612223565b6005805467ffffffffffffffff60501b1916600160501b63ffffffff89811691820263ffffffff60701b191692909217600160701b8984169081029190911767ffffffffffffffff60901b1916600160901b89851690810263ffffffff60b01b191691909117600160b01b9489169485021762ffffff60d01b1916600160d01b62ffffff89169081029190911790955560408051938452602084019290925290820152606081019190915260808101919091527f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f9060a00160405180910390a1505050505050565b6007546001600160a01b03163314610a1e5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b60448201526064016105f9565b600680546001600160a01b0319808216339081179093556007805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6001600160a01b03818116600090815260146020526040902054163314610ae25760405162461bcd60e51b815260206004820152601760248201527f4f6e6c792070617965652063616e20776974686472617700000000000000000060448201526064016105f9565b61040f81612598565b610af3612780565b828114610b425760405162461bcd60e51b815260206004820181905260248201527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a6560448201526064016105f9565b60005b83811015610cb6576000858583818110610b6157610b616144ed565b9050602002016020810190610b769190613dfe565b90506000848484818110610b8c57610b8c6144ed565b9050602002016020810190610ba19190613dfe565b6001600160a01b038084166000908152601460205260409020549192501680158080610bde5750826001600160a01b0316826001600160a01b0316145b610c1e5760405162461bcd60e51b81526020600482015260116024820152701c185e595948185b1c9958591e481cd95d607a1b60448201526064016105f9565b6001600160a01b03848116600090815260146020526040902080546001600160a01b03191685831690811790915590831614610c9f57826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b505050508080610cae906145b0565b915050610b45565b5050505050565b6001600160a01b03818116600090815260156020526040902054163314610d265760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792070726f706f736564207061796565732063616e206163636570740060448201526064016105f9565b6001600160a01b0381811660008181526014602090815260408083208054336001600160a01b031980831682179093556015909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a60408051610100808201835260055460ff808216845291810464ffffffffff16602080850191909152600160301b820463ffffffff90811685870152600160501b830481166060860152600160701b830481166080860152600160901b8304811660a0860152600160b01b83041660c0850152600160d01b90910462ffffff1660e0840152336000908152600c825293909320549394509092918c01359116610e875760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016105f9565b6009548b3514610ed15760405162461bcd60e51b81526020600482015260156024820152740c6dedcccd2ce88d2cecae6e840dad2e6dac2e8c6d605b1b60448201526064016105f9565b610edf8a8a8a8a8a8a6127d5565b8151610eec9060016145c9565b60ff168714610f3d5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016105f9565b868514610f8c5760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016105f9565b60008a8a604051610f9e9291906145ee565b604051908190038120610fb5918e906020016145fe565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a81101561114d5760006001858a8460208110611004576110046144ed565b61101191901a601b6145c9565b8f8f86818110611023576110236144ed565b905060200201358e8e8781811061103c5761103c6144ed565b9050602002013560405160008152602001604052604051611079949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561109b573d6000803e3d6000fd5b505060408051601f198101516001600160a01b0381166000908152600d602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506111265760405162461bcd60e51b815260206004820152600f60248201526e39b4b3b730ba3ab9329032b93937b960891b60448201526064016105f9565b826020015160080260ff166001901b84019350508080611145906145b0565b915050610fe5565b5081827e0101010101010101010101010101010101010101010101010101010101010116146111b15760405162461bcd60e51b815260206004820152601060248201526f323ab83634b1b0ba329039b4b3b732b960811b60448201526064016105f9565b50600091506112009050838d836020020135848e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061287292505050565b905061120e83828633612c9f565b505050505050505050505050565b6006546001600160a01b03163314806112a65750601654604051630d629b5f60e31b81526001600160a01b0390911690636b14daf8906112659033906000903690600401614569565b602060405180830381865afa158015611282573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a6919061458e565b6112f25760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c60448201526064016105f9565b60006112fc612dae565b6017546040516370a0823160e01b81523060048201529192506000916001600160a01b03909116906370a0823190602401602060405180830381865afa15801561134a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061136e919061461a565b9050818110156113b75760405162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e742062616c616e636560601b60448201526064016105f9565b6017546001600160a01b031663a9059cbb856113dc6113d686866144d6565b87612f78565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015611427573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144b919061458e565b6114675760405162461bcd60e51b81526004016105f990614633565b50505050565b6017546040516370a0823160e01b815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156114ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114de919061461a565b905060006114ea612dae565b90506114f6818361465f565b9250505090565b60008060008061150d8786612f92565b92509250925065ffffffffffff831660009081526004602090815260409182902084518154928601518487015160608801516001600160a01b0316600160481b027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff909216600160381b0291909116670100000000000000600160e81b031962ffffff9093166401000000000266ffffffffffffff1990961663ffffffff90941693909317949094171617919091179055516001600160401b038216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc59061160590889062ffffff91909116815260200190565b60405180910390a250909150505b9392505050565b611622612780565b601f8911156116665760405162461bcd60e51b815260206004820152601060248201526f746f6f206d616e79206f7261636c657360801b60448201526064016105f9565b8887146116ae5760405162461bcd60e51b81526020600482015260166024820152750dee4c2c6d8ca40d8cadccee8d040dad2e6dac2e8c6d60531b60448201526064016105f9565b886116ba87600361469e565b60ff161061170a5760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016105f9565b6117168660ff1661324e565b6040805160e060208c02808301820190935260c082018c815260009383928f918f918291908601908490808284376000920191909152505050908252506040805160208c810282810182019093528c82529283019290918d918d91829185019084908082843760009201919091525050509082525060ff891660208083019190915260408051601f8a0183900483028101830182528981529201919089908990819084018382808284376000920191909152505050908252506001600160401b03861660208083019190915260408051601f8701839004830281018301825286815292019190869086908190840183828082843760009201919091525050509152506005805465ffffffffff00191690559050611831612223565b600e5460005b818110156118e2576000600e8281548110611854576118546144ed565b6000918252602082200154600f80546001600160a01b0390921693509084908110611881576118816144ed565b60009182526020808320909101546001600160a01b039485168352600d82526040808420805461ffff1916905594168252600c90529190912080546dffffffffffffffffffffffffffff1916905550806118da816145b0565b915050611837565b506118ef600e6000613b95565b6118fb600f6000613b95565b60005b825151811015611b7457600d600084600001518381518110611922576119226144ed565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156119965760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016105f9565b604080518082019091526001815260ff8216602082015283518051600d91600091859081106119c7576119c76144ed565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051600c92919084908110611a3257611a326144ed565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611aa65760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016105f9565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b0316815250600c600085602001518481518110611aeb57611aeb6144ed565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516001600160601b0316620100000262010000600160701b031960ff959095166101000261ff00199315159390931661ffff1990941693909317919091179290921617905580611b6c816145b0565b9150506118fe565b5081518051611b8b91600e91602090910190613bb3565b506020808301518051611ba292600f920190613bb3565b5060408201516005805460ff191660ff9092169190911790556007805463ffffffff60c01b198116600160c01b63ffffffff43811682029290921793849055909104811691600091611bfd91600160a01b90041660016146c7565b905080600760146101000a81548163ffffffff021916908363ffffffff1602179055506000611c5146308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a00151613293565b9050806009600001819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05838284886000015189602001518a604001518b606001518c608001518d60a00151604051611cb499989796959493929190614728565b60405180910390a1600554600160301b900463ffffffff1660005b865151811015611d295781601082601f8110611ced57611ced6144ed565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff1602179055508080611d21906145b0565b915050611ccf565b50611d348b8b6132ee565b50505050505050505050505050505050565b6001600160a01b0381166000908152600c602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b03169181019190915290611da85750600092915050565b6010816020015160ff16601f8110611dc257611dc26144ed565b600881049190910154600554611613926007166004026101000a90910463ffffffff90811691600160301b900416614503565b6001600160a01b03828116600090815260146020526040902054163314611e5e5760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792063757272656e742070617965652063616e2075706461746500000060448201526064016105f9565b6001600160a01b0381163303611eb65760405162461bcd60e51b815260206004820152601760248201527f63616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105f9565b6001600160a01b03808316600090815260156020526040902080548383166001600160a01b031982168117909255909116908114611f29576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a45b505050565b611f36612780565b61040f816132fc565b6000806000611f4e8787612f92565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a6001600160401b031681526020018763ffffffff166001600160601b0316815250905081878a83604051602001611fc894939291906147bd565b60408051601f19818403018152828252805160209182012065ffffffffffff8716600090815291829052919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906120299084908a908d9086906147bd565b60405180910390a15090979650505050505050565b612046612780565b61040f816133a6565b6060826120885760405163c7d41b1b60e01b815265ffffffffffff861660048201526001600160401b03831660248201526044016105f9565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff16111561213e576040808601519051634a90778560e01b815261ffff90911660048201526103e860248201526044016105f9565b6000856040015161ffff166001600160401b03811115612160576121606142de565b604051908082528060200260200182016040528015612189578160200160208202803683370190505b50905060005b866040015161ffff168161ffff1610156122185782816040516020016121cc92919091825260f01b6001600160f01b031916602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff16815181106121fb576121fb6144ed565b6020908102919091010152806122108161485f565b91505061218f565b509695505050505050565b601754600554604080516103e08101918290526001600160a01b0390931692600160301b90920463ffffffff1691600091601090601f908285855b82829054906101000a900463ffffffff1663ffffffff168152602001906004019060208260030104928301926001038202915080841161225e579050505050505090506000600f8054806020026020016040519081016040528092919081815260200182805480156122f957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116122db575b5050505050905060005b815181101561258a576000600c6000848481518110612324576123246144ed565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b031690506000600c6000858581518110612386576123866144ed565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008483601f81106123e9576123e96144ed565b602002015160055490870363ffffffff9081169250600160901b909104168102633b9aca00028201801561257f5760006014600087878151811061242f5761242f6144ed565b6020908102919091018101516001600160a01b03908116835290820192909252604090810160002054905163a9059cbb60e01b815290821660048201819052602482018590529250908a169063a9059cbb906044016020604051808303816000875af11580156124a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124c7919061458e565b6124e35760405162461bcd60e51b81526004016105f990614633565b878786601f81106124f6576124f66144ed565b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b0316878781518110612533576125336144ed565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8560405161257591815260200190565b60405180910390a4505b505050600101612303565b50610cb6601083601f613c18565b6001600160a01b0381166000908152600c60209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b0316928101929092526125f5575050565b6000612600836106e0565b90508015611f29576001600160a01b038381166000908152601460205260409081902054601754915163a9059cbb60e01b8152908316600482018190526024820185905292919091169063a9059cbb906044016020604051808303816000875af1158015612672573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612696919061458e565b6126b25760405162461bcd60e51b81526004016105f990614633565b600560000160069054906101000a900463ffffffff166010846020015160ff16601f81106126e2576126e26144ed565b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b038481166000818152600c6020908152604091829020805462010000600160701b0319169055601754915186815291841693851692917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c910160405180910390a450505050565b6006546001600160a01b031633146127d35760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b60448201526064016105f9565b565b60006127e28260206144b7565b6127ed8560206144b7565b6127f988610144614528565b6128039190614528565b61280d9190614528565b612818906000614528565b90503681146128695760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016105f9565b50505050505050565b600080828060200190518101906128899190614a74565b64ffffffffff851660208801526040870180519192506128a882614c48565b63ffffffff1663ffffffff168152505085600560008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050600081604001516001600160401b031640905080826060015114612a4a576060820151604080840151905163aed0afe560e01b81526004810192909252602482018390526001600160401b031660448201526064016105f9565b6000808360000151516001600160401b03811115612a6a57612a6a6142de565b604051908082528060200260200182016040528015612aaf57816020015b6040805180820190915260008082526020820152815260200190600190039081612a885790505b50905060005b845151811015612b7f57600085600001518281518110612ad757612ad76144ed565b60200260200101519050612af4818760400151886020015161341c565b60408101515151151580612b1057506040810151516020015115155b15612b6c57604051806040016040528082600001516001600160401b03168152602001826020015162ffffff16815250838381518110612b5257612b526144ed565b60200260200101819052508380612b689061485f565b9450505b5080612b77816145b0565b915050612ab5565b5060008261ffff166001600160401b03811115612b9e57612b9e6142de565b604051908082528060200260200182016040528015612be357816020015b6040805180820190915260008082526020820152815260200190600190039081612bbc5790505b50905060005b8361ffff16811015612c3f57828181518110612c0757612c076144ed565b6020026020010151828281518110612c2157612c216144ed565b60200260200101819052508080612c37906145b0565b915050612be9565b50896040015163ffffffff167f7484067466b4f2452757769a8dc9a8b41497154367515673c79386f9f0b74f163387602001518c8c86604051612c86959493929190614c61565b60405180910390a2505050506020015195945050505050565b6000612cc6633b9aca003a04866080015163ffffffff16876060015163ffffffff166137f5565b90506010360260005a90506000612cef8663ffffffff1685858b60e0015162ffffff1686613812565b90506000670de0b6b3a76400006001600160c01b03891683026001600160a01b0388166000908152600c602052604090205460c08c01519290910492506201000090046001600160601b039081169163ffffffff16633b9aca000282840101908116821115612d645750505050505050611467565b6001600160a01b0388166000908152600c6020526040902080546001600160601b03909216620100000262010000600160701b031990921691909117905550505050505050505050565b600080600f805480602002602001604051908101604052809291908181526020018280548015612e0757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612de9575b50508351600554604080516103e08101918290529697509195600160301b90910463ffffffff169450600093509150601090601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612e405790505050505050905060005b83811015612ed3578181601f8110612ea057612ea06144ed565b6020020151612eaf9084614503565b612ebf9063ffffffff1687614528565b955080612ecb816145b0565b915050612e86565b50600554612ef290600160901b900463ffffffff16633b9aca006144b7565b612efc90866144b7565b945060005b83811015612f7057600c6000868381518110612f1f57612f1f6144ed565b6020908102919091018101516001600160a01b0316825281019190915260400160002054612f5c906201000090046001600160601b031687614528565b955080612f68816145b0565b915050612f01565b505050505090565b600081831015612f89575081612f8c565b50805b92915050565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff161115612fec57604051634a90778560e01b815261ffff861660048201526103e860248201526044016105f9565b8461ffff16600003613011576040516308fad2a760e01b815260040160405180910390fd5b600061303d7f000000000000000000000000000000000000000000000000000000000000000043614d10565b905060008161306c7f000000000000000000000000000000000000000000000000000000000000000043614528565b61307691906144d6565b905060006130a47f000000000000000000000000000000000000000000000000000000000000000083614d24565b905063ffffffff81106130ca576040516307b2a52360e41b815260040160405180910390fd5b6040805180820182526002805465ffffffffffff16825282516101008101938490528493600093929160208401916003906008908288855b82829054906101000a900462ffffff1662ffffff168152602001906003019060208260020104928301926001038202915080841161310257905050505091909252505081519192505065ffffffffffff8082161061317357604051630568cab760e31b815260040160405180910390fd5b61317e816001614d38565b6002805465ffffffffffff191665ffffffffffff9290921691909117905560005b60088110156131e5578a62ffffff16836020015182600881106131c4576131c46144ed565b602002015162ffffff16146131e557806131dd816145b0565b91505061319f565b6008811061320d576020830151604051630c4f769b60e41b81526105f9918d91600401614d81565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b8060001061040f5760405162461bcd60e51b815260206004820152601260248201527166206d75737420626520706f73697469766560701b60448201526064016105f9565b6000808a8a8a8a8a8a8a8a8a6040516020016132b799989796959493929190614d9b565b60408051601f1981840301815291905280516020909101206001600160f01b0316600160f01b179150509998505050505050505050565b6132f88282613876565b5050565b336001600160a01b038216036133545760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105f9565b600780546001600160a01b0319166001600160a01b03838116918217909255600654604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6016546001600160a01b0390811690821681146132f857601680546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b82516001600160401b038084169116111561346057825160405163012d824d60e01b81526001600160401b03808516600483015290911660248201526044016105f9565b6040830151515160009015801561347e575060408401515160200151155b156134b6575082516001600160401b031660009081526001602090815260408083208287015162ffffff168452909152902054613510565b83604001516040516020016134cb9190614e24565b60408051601f19818403018152918152815160209283012086516001600160401b03166000908152600184528281208885015162ffffff168252909352912081905590505b6060840151516000816001600160401b03811115613530576135306142de565b604051908082528060200260200182016040528015613559578160200160208202803683370190505b5090506000826001600160401b03811115613576576135766142de565b6040519080825280601f01601f1916602001820160405280156135a0576020820181803683370190505b5090506000836001600160401b038111156135bd576135bd6142de565b6040519080825280602002602001820160405280156135f057816020015b60608152602001906001900390816135db5790505b5090506000805b858110156136f35760008a606001518281518110613617576136176144ed565b6020908102919091010151905060008061363b8d600001518e602001518c8661393a565b91509150811561367a5780868661ffff168151811061365c5761365c6144ed565b602002602001018190525084806136729061485f565b9550506136a9565b600160f81b878581518110613691576136916144ed565b60200101906001600160f81b031916908160001a9053505b82515188518990869081106136c0576136c06144ed565b602002602001019065ffffffffffff16908165ffffffffffff1681525050505050806136eb816145b0565b9150506135f7565b50606089015151156137ea5760008161ffff166001600160401b0381111561371d5761371d6142de565b60405190808252806020026020018201604052801561375057816020015b606081526020019060019003908161373b5790505b50905060005b8261ffff168110156137ac57838181518110613774576137746144ed565b602002602001015182828151811061378e5761378e6144ed565b602002602001018190525080806137a4906145b0565b915050613756565b507f47ddf7bb0cbd94c1b43c5097f1352a80db0ceb3696f029d32b24f32cd631d2b78585836040516137e093929190614e57565b60405180910390a1505b505050505050505050565b6000838381101561380857600285850304015b6106d78184612f78565b6000818610156138645760405162461bcd60e51b815260206004820181905260248201527f6c6566744761732063616e6e6f742065786365656420696e697469616c47617360448201526064016105f9565b50633b9aca0094039190910101020290565b61010081811461389f57828282604051635c9d52ef60e11b81526004016105f993929190614f0d565b6138a7613caf565b81816040516020016138b99190614f31565b60405160208183030381529060405251146138d6576138d6614f40565b6040805180820190915260025465ffffffffffff168152602081016138fd85870187614f56565b905280516002805465ffffffffffff191665ffffffffffff9092169190911781556020820151613931906003906008613cce565b50611467915050565b60006060816139727f00000000000000000000000000000000000000000000000000000000000000006001600160401b038916614d24565b845160808101516040519293509091600091613996918b918b9186906020016147bd565b60408051601f198184030181529181528151602092830120845165ffffffffffff166000908152928390529120549091508114613a045760016040518060400160405280601081526020016f756e6b6e6f776e2063616c6c6261636b60801b81525094509450505050613b8c565b815165ffffffffffff16600090815260208181526040808320839055805160808101825263ffffffff8716815262ffffff8c16818401529185015161ffff16828201528401516001600160a01b031660608201528351909190613a6990838b8e61204f565b6060808401518a5160a00151875192880151604051635a47dd7160e01b815294955091936001600160a01b03851693635a47dd71936001600160601b0390931692613ab992889190600401614fdd565b600060405180830381600088803b158015613ad357600080fd5b5087f193505050508015613ae5575060015b613b6f573d808015613b13576040519150601f19603f3d011682016040523d82523d6000602084013e613b18565b606091505b50608081511015613b3557600198509650613b8c95505050505050565b60016040518060400160405280600f81526020016e6572726d736720746f6f206c6f6e6760881b8152509850985050505050505050613b8c565b600060405180602001604052806000815250975097505050505050505b94509492505050565b508054600082559060005260206000209081019061040f9190613d55565b828054828255906000526020600020908101928215613c08579160200282015b82811115613c0857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613bd3565b50613c14929150613d55565b5090565b600483019183908215613c085791602002820160005b83821115613c7257835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302613c2e565b8015613ca25782816101000a81549063ffffffff0219169055600401602081600301049283019260010302613c72565b5050613c14929150613d55565b6040518061010001604052806008906020820280368337509192915050565b600183019183908215613c085791602002820160005b83821115613d2657835183826101000a81548162ffffff021916908362ffffff1602179055509260200192600301602081600201049283019260010302613ce4565b8015613ca25782816101000a81549062ffffff0219169055600301602081600201049283019260010302613d26565b5b80821115613c145760008155600101613d56565b65ffffffffffff8116811461040f57600080fd5b600060208284031215613d9057600080fd5b813561161381613d6a565b600081518084526020808501945080840160005b83811015613dcb57815187529582019590820190600101613daf565b509495945050505050565b6020815260006116136020830184613d9b565b6001600160a01b038116811461040f57600080fd5b600060208284031215613e1057600080fd5b813561161381613de9565b60005b83811015613e36578181015183820152602001613e1e565b838111156114675750506000910152565b60008151808452613e5f816020860160208601613e1b565b601f01601f19169290920160200192915050565b6020815260006116136020830184613e47565b803563ffffffff81168114613e9a57600080fd5b919050565b62ffffff8116811461040f57600080fd5b600080600080600060a08688031215613ec857600080fd5b613ed186613e86565b9450613edf60208701613e86565b9350613eed60408701613e86565b9250613efb60608701613e86565b91506080860135613f0b81613e9f565b809150509295509295909350565b60008083601f840112613f2b57600080fd5b5081356001600160401b03811115613f4257600080fd5b6020830191508360208260051b8501011115613f5d57600080fd5b9250929050565b60008060008060408587031215613f7a57600080fd5b84356001600160401b0380821115613f9157600080fd5b613f9d88838901613f19565b90965094506020870135915080821115613fb657600080fd5b50613fc387828801613f19565b95989497509550505050565b60008083601f840112613fe157600080fd5b5081356001600160401b03811115613ff857600080fd5b602083019150836020828501011115613f5d57600080fd5b60008060008060008060008060e0898b03121561402c57600080fd5b606089018a81111561403d57600080fd5b899850356001600160401b038082111561405657600080fd5b6140628c838d01613fcf565b909950975060808b013591508082111561407b57600080fd5b6140878c838d01613f19565b909750955060a08b01359150808211156140a057600080fd5b506140ad8b828c01613f19565b999c989b50969995989497949560c00135949350505050565b600080604083850312156140d957600080fd5b82356140e481613de9565b946020939093013593505050565b60006020828403121561410457600080fd5b81356001600160401b0381111561411a57600080fd5b82016080818503121561161357600080fd5b61ffff8116811461040f57600080fd5b6001600160401b038116811461040f57600080fd5b8035613e9a8161413c565b60008060006060848603121561417157600080fd5b833561417c8161412c565b9250602084013561418c8161413c565b9150604084013561419c81613e9f565b809150509250925092565b803560ff81168114613e9a57600080fd5b60008060008060008060008060008060c08b8d0312156141d757600080fd5b8a356001600160401b03808211156141ee57600080fd5b6141fa8e838f01613f19565b909c509a5060208d013591508082111561421357600080fd5b61421f8e838f01613f19565b909a50985088915061423360408e016141a7565b975060608d013591508082111561424957600080fd5b6142558e838f01613fcf565b909750955085915061426960808e01614151565b945060a08d013591508082111561427f57600080fd5b5061428c8d828e01613fcf565b915080935050809150509295989b9194979a5092959850565b600080604083850312156142b857600080fd5b82356142c381613de9565b915060208301356142d381613de9565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715614316576143166142de565b60405290565b60405160c081016001600160401b0381118282101715614316576143166142de565b604051608081016001600160401b0381118282101715614316576143166142de565b604051602081016001600160401b0381118282101715614316576143166142de565b604051601f8201601f191681016001600160401b03811182821017156143aa576143aa6142de565b604052919050565b60006001600160401b038211156143cb576143cb6142de565b50601f01601f191660200190565b600080600080600060a086880312156143f157600080fd5b85356143fc8161413c565b9450602086013561440c8161412c565b9350604086013561441c81613e9f565b925061442a60608701613e86565b915060808601356001600160401b0381111561444557600080fd5b8601601f8101881361445657600080fd5b8035614469614464826143b2565b614382565b81815289602083850101111561447e57600080fd5b816020840160208301376000602083830101528093505050509295509295909350565b634e487b7160e01b600052601160045260246000fd5b60008160001904831182151516156144d1576144d16144a1565b500290565b6000828210156144e8576144e86144a1565b500390565b634e487b7160e01b600052603260045260246000fd5b600063ffffffff83811690831681811015614520576145206144a1565b039392505050565b6000821982111561453b5761453b6144a1565b500190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03841681526040602082018190526000906106d79083018486614540565b6000602082840312156145a057600080fd5b8151801515811461161357600080fd5b6000600182016145c2576145c26144a1565b5060010190565b600060ff821660ff84168060ff038211156145e6576145e66144a1565b019392505050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60006020828403121561462c57600080fd5b5051919050565b602080825260129082015271696e73756666696369656e742066756e647360701b604082015260600190565b60008083128015600160ff1b85018412161561467d5761467d6144a1565b6001600160ff1b0384018313811615614698576146986144a1565b50500390565b600060ff821660ff84168160ff04811182151516156146bf576146bf6144a1565b029392505050565b600063ffffffff8083168185168083038211156146e6576146e66144a1565b01949350505050565b600081518084526020808501945080840160005b83811015613dcb5781516001600160a01b031687529582019590820190600101614703565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526147588184018a6146ef565b9050828103608084015261476c81896146ef565b905060ff871660a084015282810360c08401526147898187613e47565b90506001600160401b03851660e08401528281036101008401526147ad8185613e47565b9c9b505050505050505050505050565b60006001600160401b03808716835262ffffff8616602084015280851660408401526080606084015265ffffffffffff845116608084015261ffff60208501511660a084015260018060a01b0360408501511660c0840152606084015160c060e085015261482f610140850182613e47565b60808601519092166101008501525060a0909301516001600160601b031661012090920191909152509392505050565b600061ffff808316818103614876576148766144a1565b6001019392505050565b60006001600160401b03821115614899576148996142de565b5060051b60200190565b8051613e9a8161413c565b600082601f8301126148bf57600080fd5b81516148cd614464826143b2565b8181528460208386010111156148e257600080fd5b6107dd826020830160208701613e1b565b80516001600160601b0381168114613e9a57600080fd5b600082601f83011261491b57600080fd5b8151602061492b61446483614880565b82815260059290921b8401810191818101908684111561494a57600080fd5b8286015b848110156122185780516001600160401b038082111561496d57600080fd5b90880190601f196040838c038201121561498657600080fd5b61498e6142f4565b878401518381111561499f57600080fd5b840160c0818e03840112156149b357600080fd5b6149bb61431c565b9250888101516149ca81613d6a565b835260408101516149da8161412c565b838a015260608101516149ec81613de9565b6040840152608081015184811115614a0357600080fd5b614a118e8b838501016148ae565b606085015250614a2360a082016148a3565b6080840152614a3460c082016148f3565b60a084015250818152614a49604085016148f3565b81890152865250505091830191830161494e565b80516001600160c01b0381168114613e9a57600080fd5b600060208284031215614a8657600080fd5b81516001600160401b0380821115614a9d57600080fd5b9083019060808286031215614ab157600080fd5b614ab961433e565b825182811115614ac857600080fd5b8301601f81018713614ad957600080fd5b8051614ae761446482614880565b8082825260208201915060208360051b850101925089831115614b0957600080fd5b602084015b83811015614c0957805187811115614b2557600080fd5b850160a0818d03601f19011215614b3b57600080fd5b614b4361433e565b6020820151614b518161413c565b81526040820151614b6181613e9f565b60208201526040828e03605f19011215614b7a57600080fd5b614b82614360565b8d607f840112614b9157600080fd5b614b996142f4565b808f60a086011115614baa57600080fd5b606085015b60a08601811015614bca578051835260209283019201614baf565b50825250604082015260a082015189811115614be557600080fd5b614bf48e60208386010161490a565b60608301525084525060209283019201614b0e565b50845250614c1c91505060208401614a5d565b6020820152614c2d604084016148a3565b60408201526060830151606082015280935050505092915050565b600063ffffffff808316818103614876576148766144a1565b6001600160a01b03861681526001600160c01b038516602080830191909152604080830186905264ffffffffff8516606084015260a060808401819052845190840181905260009285810192909160c0860190855b81811015614ce957855180516001600160401b0316845285015162ffffff16858401529484019491830191600101614cb6565b50909b9a5050505050505050505050565b634e487b7160e01b600052601260045260246000fd5b600082614d1f57614d1f614cfa565b500690565b600082614d3357614d33614cfa565b500490565b600065ffffffffffff8083168185168083038211156146e6576146e66144a1565b8060005b600881101561146757815162ffffff16845260209384019390910190600101614d5d565b62ffffff8316815261012081016116136020830184614d59565b8981526001600160a01b03891660208201526001600160401b03888116604083015261012060608301819052600091614dd68483018b6146ef565b91508382036080850152614dea828a6146ef565b915060ff881660a085015283820360c0850152614e078288613e47565b90861660e085015283810361010085015290506147ad8185613e47565b815160408201908260005b6002811015614e4e578251825260209283019290910190600101614e2f565b50505092915050565b606080825284519082018190526000906020906080840190828801845b82811015614e9857815165ffffffffffff1684529284019290840190600101614e74565b50505083810382850152614eac8187613e47565b905083810360408501528085518083528383019150838160051b84010184880160005b83811015614efd57601f19868403018552614eeb838351613e47565b94870194925090860190600101614ecf565b50909a9950505050505050505050565b604081526000614f21604083018587614540565b9050826020830152949350505050565b6101008101612f8c8284614d59565b634e487b7160e01b600052600160045260246000fd5b6000610100808385031215614f6a57600080fd5b83601f840112614f7957600080fd5b6040518181018181106001600160401b0382111715614f9a57614f9a6142de565b604052908301908085831115614faf57600080fd5b845b83811015614fd2578035614fc481613e9f565b825260209182019101614fb1565b509095945050505050565b65ffffffffffff84168152606060208201526000614ffe6060830185613d9b565b82810360408401526150108185613e47565b969550505050505056fea26469706673582212209600b356a3fe8b9d0180789f3bf7ac2795a8256101472161ec2a7d4844fb1eb064736f6c634300080d0033",
}

var VRFBeaconOCRABI = VRFBeaconOCRMetaData.ABI

var VRFBeaconOCRFuncSigs = VRFBeaconOCRMetaData.Sigs

var VRFBeaconOCRBin = VRFBeaconOCRMetaData.Bin

func DeployVRFBeaconOCR(auth *bind.TransactOpts, backend bind.ContractBackend, beaconPeriodBlocksArg *big.Int, link common.Address) (common.Address, *types.Transaction, *VRFBeaconOCR, error) {
	parsed, err := VRFBeaconOCRMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconOCRBin), backend, beaconPeriodBlocksArg, link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeaconOCR{VRFBeaconOCRCaller: VRFBeaconOCRCaller{contract: contract}, VRFBeaconOCRTransactor: VRFBeaconOCRTransactor{contract: contract}, VRFBeaconOCRFilterer: VRFBeaconOCRFilterer{contract: contract}}, nil
}

type VRFBeaconOCR struct {
	VRFBeaconOCRCaller
	VRFBeaconOCRTransactor
	VRFBeaconOCRFilterer
}

type VRFBeaconOCRCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconOCRTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconOCRFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconOCRSession struct {
	Contract     *VRFBeaconOCR
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconOCRCallerSession struct {
	Contract *VRFBeaconOCRCaller
	CallOpts bind.CallOpts
}

type VRFBeaconOCRTransactorSession struct {
	Contract     *VRFBeaconOCRTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconOCRRaw struct {
	Contract *VRFBeaconOCR
}

type VRFBeaconOCRCallerRaw struct {
	Contract *VRFBeaconOCRCaller
}

type VRFBeaconOCRTransactorRaw struct {
	Contract *VRFBeaconOCRTransactor
}

func NewVRFBeaconOCR(address common.Address, backend bind.ContractBackend) (*VRFBeaconOCR, error) {
	contract, err := bindVRFBeaconOCR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCR{VRFBeaconOCRCaller: VRFBeaconOCRCaller{contract: contract}, VRFBeaconOCRTransactor: VRFBeaconOCRTransactor{contract: contract}, VRFBeaconOCRFilterer: VRFBeaconOCRFilterer{contract: contract}}, nil
}

func NewVRFBeaconOCRCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconOCRCaller, error) {
	contract, err := bindVRFBeaconOCR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRCaller{contract: contract}, nil
}

func NewVRFBeaconOCRTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconOCRTransactor, error) {
	contract, err := bindVRFBeaconOCR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRTransactor{contract: contract}, nil
}

func NewVRFBeaconOCRFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconOCRFilterer, error) {
	contract, err := bindVRFBeaconOCR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRFilterer{contract: contract}, nil
}

func bindVRFBeaconOCR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconOCRABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeaconOCR *VRFBeaconOCRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconOCR.Contract.VRFBeaconOCRCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconOCR *VRFBeaconOCRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.VRFBeaconOCRTransactor.contract.Transfer(opts)
}

func (_VRFBeaconOCR *VRFBeaconOCRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.VRFBeaconOCRTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconOCR.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.contract.Transfer(opts)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconOCR.Contract.NUMCONFDELAYS(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconOCR.Contract.NUMCONFDELAYS(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) GetBilling(opts *bind.CallOpts) (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "getBilling")

	outstruct := new(struct {
		MaximumGasPriceGwei       uint32
		ReasonableGasPriceGwei    uint32
		ObservationPaymentGjuels  uint32
		TransmissionPaymentGjuels uint32
		AccountingGas             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPriceGwei = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ReasonableGasPriceGwei = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ObservationPaymentGjuels = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TransmissionPaymentGjuels = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _VRFBeaconOCR.Contract.GetBilling(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _VRFBeaconOCR.Contract.GetBilling(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeaconOCR.Contract.GetBillingAccessController(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeaconOCR.Contract.GetBillingAccessController(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) IStartSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "i_StartSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.IStartSlot(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.IStartSlot(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.IBeaconPeriodBlocks(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.IBeaconPeriodBlocks(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "latestConfigDetails")

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

func (_VRFBeaconOCR *VRFBeaconOCRSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRFBeaconOCR.Contract.LatestConfigDetails(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRFBeaconOCR.Contract.LatestConfigDetails(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

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

func (_VRFBeaconOCR *VRFBeaconOCRSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRFBeaconOCR.Contract.LatestConfigDigestAndEpoch(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRFBeaconOCR.Contract.LatestConfigDigestAndEpoch(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.LinkAvailableForPayment(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.LinkAvailableForPayment(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) MaxErrorMsgLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "maxErrorMsgLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) MaxErrorMsgLength() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.MaxErrorMsgLength(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) MaxErrorMsgLength() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.MaxErrorMsgLength(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) MaxNumWords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "maxNumWords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.MaxNumWords(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconOCR.Contract.MaxNumWords(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) MinDelay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "minDelay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) MinDelay() (uint16, error) {
	return _VRFBeaconOCR.Contract.MinDelay(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) MinDelay() (uint16, error) {
	return _VRFBeaconOCR.Contract.MinDelay(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _VRFBeaconOCR.Contract.OracleObservationCount(&_VRFBeaconOCR.CallOpts, transmitterAddress)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _VRFBeaconOCR.Contract.OracleObservationCount(&_VRFBeaconOCR.CallOpts, transmitterAddress)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeaconOCR.Contract.OwedPayment(&_VRFBeaconOCR.CallOpts, transmitterAddress)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeaconOCR.Contract.OwedPayment(&_VRFBeaconOCR.CallOpts, transmitterAddress)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) Owner() (common.Address, error) {
	return _VRFBeaconOCR.Contract.Owner(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) Owner() (common.Address, error) {
	return _VRFBeaconOCR.Contract.Owner(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) TypeAndVersion() (string, error) {
	return _VRFBeaconOCR.Contract.TypeAndVersion(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) TypeAndVersion() (string, error) {
	return _VRFBeaconOCR.Contract.TypeAndVersion(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.AcceptOwnership(&_VRFBeaconOCR.TransactOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.AcceptOwnership(&_VRFBeaconOCR.TransactOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.AcceptPayeeship(&_VRFBeaconOCR.TransactOpts, transmitter)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.AcceptPayeeship(&_VRFBeaconOCR.TransactOpts, transmitter)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) ExposeType(opts *bind.TransactOpts, arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "exposeType", arg0)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.ExposeType(&_VRFBeaconOCR.TransactOpts, arg0)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.ExposeType(&_VRFBeaconOCR.TransactOpts, arg0)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) GetRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "getRandomness", requestID)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.GetRandomness(&_VRFBeaconOCR.TransactOpts, requestID)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.GetRandomness(&_VRFBeaconOCR.TransactOpts, requestID)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) RequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "requestRandomness", numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.RequestRandomness(&_VRFBeaconOCR.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.RequestRandomness(&_VRFBeaconOCR.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.RequestRandomnessFulfillment(&_VRFBeaconOCR.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.RequestRandomnessFulfillment(&_VRFBeaconOCR.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetBilling(&_VRFBeaconOCR.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetBilling(&_VRFBeaconOCR.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetBillingAccessController(&_VRFBeaconOCR.TransactOpts, _billingAccessController)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetBillingAccessController(&_VRFBeaconOCR.TransactOpts, _billingAccessController)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetConfig(&_VRFBeaconOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetConfig(&_VRFBeaconOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetPayees(&_VRFBeaconOCR.TransactOpts, transmitters, payees)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetPayees(&_VRFBeaconOCR.TransactOpts, transmitters, payees)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.TransferOwnership(&_VRFBeaconOCR.TransactOpts, to)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.TransferOwnership(&_VRFBeaconOCR.TransactOpts, to)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.TransferPayeeship(&_VRFBeaconOCR.TransactOpts, transmitter, proposed)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.TransferPayeeship(&_VRFBeaconOCR.TransactOpts, transmitter, proposed)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.Transmit(&_VRFBeaconOCR.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.Transmit(&_VRFBeaconOCR.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.WithdrawFunds(&_VRFBeaconOCR.TransactOpts, recipient, amount)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.WithdrawFunds(&_VRFBeaconOCR.TransactOpts, recipient, amount)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.WithdrawPayment(&_VRFBeaconOCR.TransactOpts, transmitter)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.WithdrawPayment(&_VRFBeaconOCR.TransactOpts, transmitter)
}

type VRFBeaconOCRBillingAccessControllerSetIterator struct {
	Event *VRFBeaconOCRBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRBillingAccessControllerSet)
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
		it.Event = new(VRFBeaconOCRBillingAccessControllerSet)
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

func (it *VRFBeaconOCRBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*VRFBeaconOCRBillingAccessControllerSetIterator, error) {

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRBillingAccessControllerSetIterator{contract: _VRFBeaconOCR.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRBillingAccessControllerSet)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseBillingAccessControllerSet(log types.Log) (*VRFBeaconOCRBillingAccessControllerSet, error) {
	event := new(VRFBeaconOCRBillingAccessControllerSet)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRBillingSetIterator struct {
	Event *VRFBeaconOCRBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRBillingSet)
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
		it.Event = new(VRFBeaconOCRBillingSet)
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

func (it *VRFBeaconOCRBillingSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterBillingSet(opts *bind.FilterOpts) (*VRFBeaconOCRBillingSetIterator, error) {

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRBillingSetIterator{contract: _VRFBeaconOCR.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRBillingSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRBillingSet)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseBillingSet(log types.Log) (*VRFBeaconOCRBillingSet, error) {
	event := new(VRFBeaconOCRBillingSet)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRConfigSetIterator struct {
	Event *VRFBeaconOCRConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRConfigSet)
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
		it.Event = new(VRFBeaconOCRConfigSet)
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

func (it *VRFBeaconOCRConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRConfigSet struct {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFBeaconOCRConfigSetIterator, error) {

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRConfigSetIterator{contract: _VRFBeaconOCR.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRConfigSet)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseConfigSet(log types.Log) (*VRFBeaconOCRConfigSet, error) {
	event := new(VRFBeaconOCRConfigSet)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRNewTransmissionIterator struct {
	Event *VRFBeaconOCRNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRNewTransmission)
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
		it.Event = new(VRFBeaconOCRNewTransmission)
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

func (it *VRFBeaconOCRNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRNewTransmission struct {
	AggregatorRoundId uint32
	Transmitter       common.Address
	JuelsPerFeeCoin   *big.Int
	ConfigDigest      [32]byte
	EpochAndRound     *big.Int
	OutputsServed     []VRFBeaconReportOutputServed
	Raw               types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*VRFBeaconOCRNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRNewTransmissionIterator{contract: _VRFBeaconOCR.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRNewTransmission)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseNewTransmission(log types.Log) (*VRFBeaconOCRNewTransmission, error) {
	event := new(VRFBeaconOCRNewTransmission)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCROraclePaidIterator struct {
	Event *VRFBeaconOCROraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCROraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCROraclePaid)
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
		it.Event = new(VRFBeaconOCROraclePaid)
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

func (it *VRFBeaconOCROraclePaidIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCROraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCROraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*VRFBeaconOCROraclePaidIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCROraclePaidIterator{contract: _VRFBeaconOCR.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCROraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCROraclePaid)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseOraclePaid(log types.Log) (*VRFBeaconOCROraclePaid, error) {
	event := new(VRFBeaconOCROraclePaid)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCROwnershipTransferRequestedIterator struct {
	Event *VRFBeaconOCROwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCROwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCROwnershipTransferRequested)
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
		it.Event = new(VRFBeaconOCROwnershipTransferRequested)
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

func (it *VRFBeaconOCROwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCROwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCROwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOCROwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCROwnershipTransferRequestedIterator{contract: _VRFBeaconOCR.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCROwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCROwnershipTransferRequested)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconOCROwnershipTransferRequested, error) {
	event := new(VRFBeaconOCROwnershipTransferRequested)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCROwnershipTransferredIterator struct {
	Event *VRFBeaconOCROwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCROwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCROwnershipTransferred)
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
		it.Event = new(VRFBeaconOCROwnershipTransferred)
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

func (it *VRFBeaconOCROwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCROwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCROwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOCROwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCROwnershipTransferredIterator{contract: _VRFBeaconOCR.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCROwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCROwnershipTransferred)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseOwnershipTransferred(log types.Log) (*VRFBeaconOCROwnershipTransferred, error) {
	event := new(VRFBeaconOCROwnershipTransferred)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRPayeeshipTransferRequestedIterator struct {
	Event *VRFBeaconOCRPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRPayeeshipTransferRequested)
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
		it.Event = new(VRFBeaconOCRPayeeshipTransferRequested)
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

func (it *VRFBeaconOCRPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*VRFBeaconOCRPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRPayeeshipTransferRequestedIterator{contract: _VRFBeaconOCR.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRPayeeshipTransferRequested)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParsePayeeshipTransferRequested(log types.Log) (*VRFBeaconOCRPayeeshipTransferRequested, error) {
	event := new(VRFBeaconOCRPayeeshipTransferRequested)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRPayeeshipTransferredIterator struct {
	Event *VRFBeaconOCRPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRPayeeshipTransferred)
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
		it.Event = new(VRFBeaconOCRPayeeshipTransferred)
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

func (it *VRFBeaconOCRPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*VRFBeaconOCRPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRPayeeshipTransferredIterator{contract: _VRFBeaconOCR.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRPayeeshipTransferred)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParsePayeeshipTransferred(log types.Log) (*VRFBeaconOCRPayeeshipTransferred, error) {
	event := new(VRFBeaconOCRPayeeshipTransferred)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRRandomWordsFulfilledIterator struct {
	Event *VRFBeaconOCRRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRRandomWordsFulfilled)
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
		it.Event = new(VRFBeaconOCRRandomWordsFulfilled)
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

func (it *VRFBeaconOCRRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFBeaconOCRRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRRandomWordsFulfilledIterator{contract: _VRFBeaconOCR.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRRandomWordsFulfilled)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFBeaconOCRRandomWordsFulfilled, error) {
	event := new(VRFBeaconOCRRandomWordsFulfilled)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRRandomnessFulfillmentRequestedIterator struct {
	Event *VRFBeaconOCRRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRRandomnessFulfillmentRequested)
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
		it.Event = new(VRFBeaconOCRRandomnessFulfillmentRequested)
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

func (it *VRFBeaconOCRRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRRandomnessFulfillmentRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  uint64
	Callback               VRFBeaconTypesCallback
	Raw                    types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts) (*VRFBeaconOCRRandomnessFulfillmentRequestedIterator, error) {

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconOCR.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRRandomnessFulfillmentRequested) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRRandomnessFulfillmentRequested)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconOCRRandomnessFulfillmentRequested, error) {
	event := new(VRFBeaconOCRRandomnessFulfillmentRequested)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOCRRandomnessRequestedIterator struct {
	Event *VRFBeaconOCRRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRRandomnessRequested)
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
		it.Event = new(VRFBeaconOCRRandomnessRequested)
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

func (it *VRFBeaconOCRRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRRandomnessRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	Raw                    types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, nextBeaconOutputHeight []uint64) (*VRFBeaconOCRRandomnessRequestedIterator, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRRandomnessRequestedIterator{contract: _VRFBeaconOCR.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRRandomnessRequested, nextBeaconOutputHeight []uint64) (event.Subscription, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRRandomnessRequested)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseRandomnessRequested(log types.Log) (*VRFBeaconOCRRandomnessRequested, error) {
	event := new(VRFBeaconOCRRandomnessRequested)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFBeaconReportMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconReport.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconReport.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxErrorMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2f7527cc": "NUM_CONF_DELAYS()",
		"c278e5b7": "exposeType(((uint64,uint24,(uint256[2]),((uint48,uint16,address,bytes,uint64,uint96),uint96)[])[],uint192,uint64,bytes32))",
		"0b93e168": "getRandomness(uint48)",
		"cf7e754a": "i_StartSlot()",
		"cd0593df": "i_beaconPeriodBlocks()",
		"7a464944": "maxErrorMsgLength()",
		"bbcdd0d8": "maxNumWords()",
		"c63c4e9b": "minDelay()",
		"dc92accf": "requestRandomness(uint16,uint64,uint24)",
		"f645dcb1": "requestRandomnessFulfillment(uint64,uint16,uint24,uint32,bytes)",
	},
	Bin: "0x60c06040523480156200001157600080fd5b50604051620010393803806200103983398101604081905262000034916200009a565b80806000036200005757604051632abc297960e01b815260040160405180910390fd5b608081905260006200006a8243620000b4565b90506000816080516200007e9190620000ed565b90506200008c814362000107565b60a052506200012292505050565b600060208284031215620000ad57600080fd5b5051919050565b600082620000d257634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b600082821015620001025762000102620000d7565b500390565b600082198211156200011d576200011d620000d7565b500190565b60805160a051610ed562000164600039600061015f01526000818161013801528181610286015281816107d701528181610806015261083e0152610ed56000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063c63c4e9b11610066578063c63c4e9b14610118578063cd0593df14610133578063cf7e754a1461015a578063dc92accf14610181578063f645dcb1146101ab57600080fd5b80630b93e168146100a35780632f7527cc146100cc5780637a464944146100e6578063bbcdd0d8146100fc578063c278e5b714610105575b600080fd5b6100b66100b1366004610a0d565b6101be565b6040516100c39190610a3c565b60405180910390f35b6100d4600881565b60405160ff90911681526020016100c3565b6100ee608081565b6040519081526020016100c3565b6100ee6103e881565b610116610113366004610a80565b50565b005b610120600381565b60405161ffff90911681526020016100c3565b6100ee7f000000000000000000000000000000000000000000000000000000000000000081565b6100ee7f000000000000000000000000000000000000000000000000000000000000000081565b61019461018f366004610afd565b61035a565b60405165ffffffffffff90911681526020016100c3565b6101946101b9366004610b56565b610476565b65ffffffffffff811660008181526004602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff600160381b820416938201939093526001600160a01b03600160481b84048116606083810191825298909752949093526001600160e81b03199091169055915116331461027b576060810151604051638e30e82360e01b81526001600160a01b0390911660048201523360248201526044015b60405180910390fd5b80516000906102b1907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16610c6b565b90506000826020015162ffffff16436102ca9190610c8a565b90508082106102f5576040516315ad27c360e01b815260048101839052436024820152604401610272565b67ffffffffffffffff821115610321576040516302c6ef8160e11b815260048101839052602401610272565b60008281526001602090815260408083208287015162ffffff16845290915290205461035190869085908561057b565b95945050505050565b60008060008061036a8786610751565b92509250925065ffffffffffff831660009081526004602090815260409182902084518154928601518487015160608801516001600160a01b0316600160481b027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff909216600160381b0291909116670100000000000000600160e81b031962ffffff9093166401000000000266ffffffffffffff1990961663ffffffff909416939093179490941716179190911790555167ffffffffffffffff8216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc59061046390889062ffffff91909116815260200190565b60405180910390a2509095945050505050565b60008060006104858787610751565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a67ffffffffffffffff1681526020018763ffffffff166bffffffffffffffffffffffff16815250905081878a836040516020016105059493929190610ca1565b60408051601f19818403018152828252805160209182012065ffffffffffff8716600090815291829052919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906105669084908a908d908690610ca1565b60405180910390a15090979650505050505050565b6060826105b55760405163c7d41b1b60e01b815265ffffffffffff8616600482015267ffffffffffffffff83166024820152604401610272565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff16111561066b576040808601519051634a90778560e01b815261ffff90911660048201526103e86024820152604401610272565b6000856040015161ffff1667ffffffffffffffff81111561068e5761068e610b40565b6040519080825280602002602001820160405280156106b7578160200160208202803683370190505b50905060005b866040015161ffff168161ffff1610156107465782816040516020016106fa92919091825260f01b6001600160f01b031916602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff168151811061072957610729610d8c565b60209081029190910101528061073e81610da2565b9150506106bd565b509695505050505050565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff1611156107ab57604051634a90778560e01b815261ffff861660048201526103e86024820152604401610272565b8461ffff166000036107d0576040516308fad2a760e01b815260040160405180910390fd5b60006107fc7f000000000000000000000000000000000000000000000000000000000000000043610dd9565b905060008161082b7f000000000000000000000000000000000000000000000000000000000000000043610ded565b6108359190610c8a565b905060006108637f000000000000000000000000000000000000000000000000000000000000000083610e05565b905063ffffffff8110610889576040516307b2a52360e41b815260040160405180910390fd5b6040805180820182526002805465ffffffffffff16825282516101008101938490528493600093929160208401916003906008908288855b82829054906101000a900462ffffff1662ffffff16815260200190600301906020826002010492830192600103820291508084116108c157905050505091909252505081519192505065ffffffffffff8082161061093257604051630568cab760e31b815260040160405180910390fd5b61093d816001610e19565b6002805465ffffffffffff191665ffffffffffff9290921691909117905560005b60088110156109a4578a62ffffff168360200151826008811061098357610983610d8c565b602002015162ffffff16146109a4578061099c81610e43565b91505061095e565b600881106109cc576020830151604051630c4f769b60e41b8152610272918d91600401610e5c565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b600060208284031215610a1f57600080fd5b813565ffffffffffff81168114610a3557600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610a7457835183529284019291840191600101610a58565b50909695505050505050565b600060208284031215610a9257600080fd5b813567ffffffffffffffff811115610aa957600080fd5b820160808185031215610a3557600080fd5b803561ffff81168114610acd57600080fd5b919050565b803567ffffffffffffffff81168114610acd57600080fd5b803562ffffff81168114610acd57600080fd5b600080600060608486031215610b1257600080fd5b610b1b84610abb565b9250610b2960208501610ad2565b9150610b3760408501610aea565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b600080600080600060a08688031215610b6e57600080fd5b610b7786610ad2565b9450610b8560208701610abb565b9350610b9360408701610aea565b9250606086013563ffffffff81168114610bac57600080fd5b9150608086013567ffffffffffffffff80821115610bc957600080fd5b818801915088601f830112610bdd57600080fd5b813581811115610bef57610bef610b40565b604051601f8201601f19908116603f01168101908382118183101715610c1757610c17610b40565b816040528281528b6020848701011115610c3057600080fd5b8260208601602083013760006020848301015280955050505050509295509295909350565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610c8557610c85610c55565b500290565b600082821015610c9c57610c9c610c55565b500390565b600067ffffffffffffffff8087168352602062ffffff87168185015281861660408501526080606085015265ffffffffffff855116608085015261ffff818601511660a085015260018060a01b0360408601511660c08501526060850151915060c060e085015281518061014086015260005b81811015610d315783810183015186820161016001528201610d14565b81811115610d4457600061016083880101525b50608086015167ffffffffffffffff1661010086015260a0909501516bffffffffffffffffffffffff16610120850152505050601f909101601f191601610160019392505050565b634e487b7160e01b600052603260045260246000fd5b600061ffff808316818103610db957610db9610c55565b6001019392505050565b634e487b7160e01b600052601260045260246000fd5b600082610de857610de8610dc3565b500690565b60008219821115610e0057610e00610c55565b500190565b600082610e1457610e14610dc3565b500490565b600065ffffffffffff808316818516808303821115610e3a57610e3a610c55565b01949350505050565b600060018201610e5557610e55610c55565b5060010190565b62ffffff838116825261012082019060208084018560005b6008811015610e93578151851683529183019190830190600101610e74565b5050505050939250505056fea2646970667358221220d5b9d0785a7ba7fcaa0806e595512680161f97af3fd0f496315b848d39dbda9c64736f6c634300080d0033",
}

var VRFBeaconReportABI = VRFBeaconReportMetaData.ABI

var VRFBeaconReportFuncSigs = VRFBeaconReportMetaData.Sigs

var VRFBeaconReportBin = VRFBeaconReportMetaData.Bin

func DeployVRFBeaconReport(auth *bind.TransactOpts, backend bind.ContractBackend, beaconPeriodBlocksArg *big.Int) (common.Address, *types.Transaction, *VRFBeaconReport, error) {
	parsed, err := VRFBeaconReportMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconReportBin), backend, beaconPeriodBlocksArg)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeaconReport{VRFBeaconReportCaller: VRFBeaconReportCaller{contract: contract}, VRFBeaconReportTransactor: VRFBeaconReportTransactor{contract: contract}, VRFBeaconReportFilterer: VRFBeaconReportFilterer{contract: contract}}, nil
}

type VRFBeaconReport struct {
	VRFBeaconReportCaller
	VRFBeaconReportTransactor
	VRFBeaconReportFilterer
}

type VRFBeaconReportCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconReportTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconReportFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconReportSession struct {
	Contract     *VRFBeaconReport
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconReportCallerSession struct {
	Contract *VRFBeaconReportCaller
	CallOpts bind.CallOpts
}

type VRFBeaconReportTransactorSession struct {
	Contract     *VRFBeaconReportTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconReportRaw struct {
	Contract *VRFBeaconReport
}

type VRFBeaconReportCallerRaw struct {
	Contract *VRFBeaconReportCaller
}

type VRFBeaconReportTransactorRaw struct {
	Contract *VRFBeaconReportTransactor
}

func NewVRFBeaconReport(address common.Address, backend bind.ContractBackend) (*VRFBeaconReport, error) {
	contract, err := bindVRFBeaconReport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReport{VRFBeaconReportCaller: VRFBeaconReportCaller{contract: contract}, VRFBeaconReportTransactor: VRFBeaconReportTransactor{contract: contract}, VRFBeaconReportFilterer: VRFBeaconReportFilterer{contract: contract}}, nil
}

func NewVRFBeaconReportCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconReportCaller, error) {
	contract, err := bindVRFBeaconReport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportCaller{contract: contract}, nil
}

func NewVRFBeaconReportTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconReportTransactor, error) {
	contract, err := bindVRFBeaconReport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportTransactor{contract: contract}, nil
}

func NewVRFBeaconReportFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconReportFilterer, error) {
	contract, err := bindVRFBeaconReport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportFilterer{contract: contract}, nil
}

func bindVRFBeaconReport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconReportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeaconReport *VRFBeaconReportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconReport.Contract.VRFBeaconReportCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconReport *VRFBeaconReportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.VRFBeaconReportTransactor.contract.Transfer(opts)
}

func (_VRFBeaconReport *VRFBeaconReportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.VRFBeaconReportTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeaconReport *VRFBeaconReportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeaconReport.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.contract.Transfer(opts)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconReport.Contract.NUMCONFDELAYS(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeaconReport.Contract.NUMCONFDELAYS(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) IStartSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "i_StartSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconReport.Contract.IStartSlot(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) IStartSlot() (*big.Int, error) {
	return _VRFBeaconReport.Contract.IStartSlot(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconReport.Contract.IBeaconPeriodBlocks(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _VRFBeaconReport.Contract.IBeaconPeriodBlocks(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) MaxErrorMsgLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "maxErrorMsgLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) MaxErrorMsgLength() (*big.Int, error) {
	return _VRFBeaconReport.Contract.MaxErrorMsgLength(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) MaxErrorMsgLength() (*big.Int, error) {
	return _VRFBeaconReport.Contract.MaxErrorMsgLength(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) MaxNumWords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "maxNumWords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconReport.Contract.MaxNumWords(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) MaxNumWords() (*big.Int, error) {
	return _VRFBeaconReport.Contract.MaxNumWords(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) MinDelay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "minDelay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) MinDelay() (uint16, error) {
	return _VRFBeaconReport.Contract.MinDelay(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) MinDelay() (uint16, error) {
	return _VRFBeaconReport.Contract.MinDelay(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportTransactor) ExposeType(opts *bind.TransactOpts, arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "exposeType", arg0)
}

func (_VRFBeaconReport *VRFBeaconReportSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.ExposeType(&_VRFBeaconReport.TransactOpts, arg0)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.ExposeType(&_VRFBeaconReport.TransactOpts, arg0)
}

func (_VRFBeaconReport *VRFBeaconReportTransactor) GetRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "getRandomness", requestID)
}

func (_VRFBeaconReport *VRFBeaconReportSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.GetRandomness(&_VRFBeaconReport.TransactOpts, requestID)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.GetRandomness(&_VRFBeaconReport.TransactOpts, requestID)
}

func (_VRFBeaconReport *VRFBeaconReportTransactor) RequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "requestRandomness", numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconReport *VRFBeaconReportSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.RequestRandomness(&_VRFBeaconReport.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.RequestRandomness(&_VRFBeaconReport.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_VRFBeaconReport *VRFBeaconReportTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconReport *VRFBeaconReportSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.RequestRandomnessFulfillment(&_VRFBeaconReport.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.RequestRandomnessFulfillment(&_VRFBeaconReport.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

type VRFBeaconReportNewTransmissionIterator struct {
	Event *VRFBeaconReportNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportNewTransmission)
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
		it.Event = new(VRFBeaconReportNewTransmission)
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

func (it *VRFBeaconReportNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportNewTransmission struct {
	AggregatorRoundId uint32
	Transmitter       common.Address
	JuelsPerFeeCoin   *big.Int
	ConfigDigest      [32]byte
	EpochAndRound     *big.Int
	OutputsServed     []VRFBeaconReportOutputServed
	Raw               types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*VRFBeaconReportNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportNewTransmissionIterator{contract: _VRFBeaconReport.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportNewTransmission)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseNewTransmission(log types.Log) (*VRFBeaconReportNewTransmission, error) {
	event := new(VRFBeaconReportNewTransmission)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconReportRandomWordsFulfilledIterator struct {
	Event *VRFBeaconReportRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportRandomWordsFulfilled)
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
		it.Event = new(VRFBeaconReportRandomWordsFulfilled)
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

func (it *VRFBeaconReportRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFBeaconReportRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportRandomWordsFulfilledIterator{contract: _VRFBeaconReport.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportRandomWordsFulfilled)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFBeaconReportRandomWordsFulfilled, error) {
	event := new(VRFBeaconReportRandomWordsFulfilled)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconReportRandomnessFulfillmentRequestedIterator struct {
	Event *VRFBeaconReportRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportRandomnessFulfillmentRequested)
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
		it.Event = new(VRFBeaconReportRandomnessFulfillmentRequested)
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

func (it *VRFBeaconReportRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportRandomnessFulfillmentRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  uint64
	Callback               VRFBeaconTypesCallback
	Raw                    types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts) (*VRFBeaconReportRandomnessFulfillmentRequestedIterator, error) {

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconReport.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportRandomnessFulfillmentRequested) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportRandomnessFulfillmentRequested)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconReportRandomnessFulfillmentRequested, error) {
	event := new(VRFBeaconReportRandomnessFulfillmentRequested)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconReportRandomnessRequestedIterator struct {
	Event *VRFBeaconReportRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportRandomnessRequested)
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
		it.Event = new(VRFBeaconReportRandomnessRequested)
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

func (it *VRFBeaconReportRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportRandomnessRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	Raw                    types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, nextBeaconOutputHeight []uint64) (*VRFBeaconReportRandomnessRequestedIterator, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportRandomnessRequestedIterator{contract: _VRFBeaconReport.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportRandomnessRequested, nextBeaconOutputHeight []uint64) (event.Subscription, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportRandomnessRequested)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseRandomnessRequested(log types.Log) (*VRFBeaconReportRandomnessRequested, error) {
	event := new(VRFBeaconReportRandomnessRequested)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFBeaconTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122010116c0e3c2d41a6ba5c0bb8e289278b725501c402f39fee23ff3a1eed03c11164736f6c634300080d0033",
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

var WrapperBeaconCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"contractDKG\",\"name\":\"keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OffchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"occVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconReport.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconReport.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxErrorMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"latestEpochAndRound\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"latestAggregatorRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"internalType\":\"structVRFBeaconReport.HotVars\",\"name\":\"hotVars\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"rawReport\",\"type\":\"bytes\"}],\"name\":\"testReport\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2f7527cc": "NUM_CONF_DELAYS()",
		"79ba5097": "acceptOwnership()",
		"b121e147": "acceptPayeeship(address)",
		"06661abd": "count()",
		"c278e5b7": "exposeType(((uint64,uint24,(uint256[2]),((uint48,uint16,address,bytes,uint64,uint96),uint96)[])[],uint192,uint64,bytes32))",
		"29937268": "getBilling()",
		"c4c92b37": "getBillingAccessController()",
		"0b93e168": "getRandomness(uint48)",
		"cf7e754a": "i_StartSlot()",
		"cd0593df": "i_beaconPeriodBlocks()",
		"bf2732c7": "keyGenerated((bytes,bytes32[]))",
		"81ff7048": "latestConfigDetails()",
		"afcb95d7": "latestConfigDigestAndEpoch()",
		"d09dc339": "linkAvailableForPayment()",
		"7a464944": "maxErrorMsgLength()",
		"bbcdd0d8": "maxNumWords()",
		"c63c4e9b": "minDelay()",
		"55e48749": "newKeyRequested()",
		"e4902f82": "oracleObservationCount(address)",
		"0eafb25b": "owedPayment(address)",
		"8da5cb5b": "owner()",
		"5a47dd71": "rawFulfillRandomWords(uint48,uint256[],bytes)",
		"dc92accf": "requestRandomness(uint16,uint64,uint24)",
		"f645dcb1": "requestRandomnessFulfillment(uint64,uint16,uint24,uint32,bytes)",
		"cc31f7dd": "s_keyID()",
		"d57fc45a": "s_provingKeyHash()",
		"643dc105": "setBilling(uint32,uint32,uint32,uint32,uint24)",
		"fbffd2c1": "setBillingAccessController(address)",
		"e3d0e712": "setConfig(address[],address[],uint8,bytes,uint64,bytes)",
		"9c849b30": "setPayees(address[],address[])",
		"3eaddf3f": "testReport((uint8,uint40,uint32,uint32,uint32,uint32,uint32,uint24),bytes32,uint40,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"eb5dcd6c": "transferPayeeship(address,address)",
		"b1dc65a4": "transmit(bytes32[3],bytes,bytes32[],bytes32[],bytes32)",
		"181f5a77": "typeAndVersion()",
		"c1075329": "withdrawFunds(address,uint256)",
		"8ac28d5a": "withdrawPayment(address)",
	},
	Bin: "0x60e06040526000601b553480156200001657600080fd5b5060405162005a3938038062005a3983398101604081905262000039916200024b565b338484848481818486338060008480806000036200006a57604051632abc297960e01b815260040160405180910390fd5b608081905260006200007d824362000299565b9050600081608051620000919190620002d2565b90506200009f8143620002ec565b60a0525050506001600160a01b0383169050620001035760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600680546001600160a01b0319166001600160a01b03848116919091179091558116156200013657620001368162000186565b5050601780546001600160a01b03199081166001600160a01b03948516179091556018805490911695831695909517909455601992909255509590951660c0525062000307975050505050505050565b336001600160a01b03821603620001e05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000fa565b600780546001600160a01b0319166001600160a01b03838116918217909255600654604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6001600160a01b03811681146200024857600080fd5b50565b600080600080608085870312156200026257600080fd5b84516200026f8162000232565b602086015160408701519195509350620002898162000232565b6060959095015193969295505050565b600082620002b757634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b600082821015620002e757620002e7620002bc565b500390565b60008219821115620003025762000302620002bc565b500190565b60805160a05160c0516156df6200035a6000396000610a7b0152600061067201526000818161063e015281816108930152818161338d015281816133bc015281816133f40152613cb901526156df6000f3fe60806040526004361061021a5760003560e01c8063b1dc65a411610123578063cf7e754a116100ab578063e4902f821161006f578063e4902f8214610716578063eb5dcd6c1461074b578063f2fde38b1461076b578063f645dcb11461078b578063fbffd2c1146107ab57600080fd5b8063cf7e754a14610660578063d09dc33914610694578063d57fc45a146106a9578063dc92accf146106bf578063e3d0e712146106f657600080fd5b8063c278e5b7116100f2578063c278e5b7146105b2578063c4c92b37146105d0578063c63c4e9b146105ee578063cc31f7dd14610616578063cd0593df1461062c57600080fd5b8063b1dc65a41461053c578063bbcdd0d81461055c578063bf2732c714610572578063c10753291461059257600080fd5b8063643dc105116101a65780638ac28d5a116101755780638ac28d5a146104735780638da5cb5b146104935780639c849b30146104c5578063afcb95d7146104e5578063b121e1471461051c57600080fd5b8063643dc105146103f057806379ba5097146104105780637a4649441461042557806381ff70481461043a57600080fd5b806329937268116101ed57806329937268146102d95780632f7527cc146103655780633eaddf3f1461038c57806355e48749146103b75780635a47dd71146103d057600080fd5b806306661abd1461021f5780630b93e168146102485780630eafb25b14610275578063181f5a7714610295575b600080fd5b34801561022b57600080fd5b50610235601b5481565b6040519081526020015b60405180910390f35b34801561025457600080fd5b506102686102633660046140f3565b6107cb565b60405161023f919061414b565b34801561028157600080fd5b50610235610290366004614173565b610966565b3480156102a157600080fd5b506040805180820182526015815274565246426561636f6e20312e302e302d616c70686160581b6020820152905161023f91906141e8565b3480156102e557600080fd5b50610329600554600160501b810463ffffffff90811692600160701b8304821692600160901b8104831692600160b01b82041691600160d01b90910462ffffff1690565b6040805163ffffffff9687168152948616602086015292851692840192909252909216606082015262ffffff909116608082015260a00161023f565b34801561037157600080fd5b5061037a600881565b60405160ff909116815260200161023f565b61039f61039a3660046143bf565b610a6b565b6040516001600160c01b03909116815260200161023f565b3480156103c357600080fd5b506103ce6000601a55565b005b3480156103dc57600080fd5b506103ce6103eb3660046144e1565b610a79565b3480156103fc57600080fd5b506103ce61040b3660046145ad565b610b01565b34801561041c57600080fd5b506103ce610ce7565b34801561043157600080fd5b50610235608081565b34801561044657600080fd5b506007546009546040805160008152600160c01b90930463ffffffff16602084015282015260600161023f565b34801561047f57600080fd5b506103ce61048e366004614173565b610d95565b34801561049f57600080fd5b506006546001600160a01b03165b6040516001600160a01b03909116815260200161023f565b3480156104d157600080fd5b506103ce6104e0366004614661565b610e07565b3480156104f157600080fd5b50600954600b546040805160008152602081019390935263ffffffff9091169082015260600161023f565b34801561052857600080fd5b506103ce610537366004614173565b610fd9565b34801561054857600080fd5b506103ce61055736600461470d565b6110b5565b34801561056857600080fd5b506102356103e881565b34801561057e57600080fd5b506103ce61058d3660046147c3565b611538565b34801561059e57600080fd5b506103ce6105ad3660046148ab565b611578565b3480156105be57600080fd5b506103ce6105cd3660046148d7565b50565b3480156105dc57600080fd5b506016546001600160a01b03166104ad565b3480156105fa57600080fd5b50610603600381565b60405161ffff909116815260200161023f565b34801561062257600080fd5b5061023560195481565b34801561063857600080fd5b506102357f000000000000000000000000000000000000000000000000000000000000000081565b34801561066c57600080fd5b506102357f000000000000000000000000000000000000000000000000000000000000000081565b3480156106a057600080fd5b506102356117c9565b3480156106b557600080fd5b50610235601a5481565b3480156106cb57600080fd5b506106df6106da366004614941565b611859565b60405165ffffffffffff909116815260200161023f565b34801561070257600080fd5b506103ce61071136600461498c565b611976565b34801561072257600080fd5b50610736610731366004614173565b6120a2565b60405163ffffffff909116815260200161023f565b34801561075757600080fd5b506103ce610766366004614a79565b612151565b34801561077757600080fd5b506103ce610786366004614173565b612289565b34801561079757600080fd5b506106df6107a6366004614ab2565b61229a565b3480156107b757600080fd5b506103ce6107c6366004614173565b612399565b65ffffffffffff811660008181526004602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff600160381b820416938201939093526001600160a01b03600160481b84048116606083810191825298909752949093526001600160e81b031990911690559151163314610888576060810151604051638e30e82360e01b81526001600160a01b0390911660048201523360248201526044015b60405180910390fd5b80516000906108be907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16614b4d565b90506000826020015162ffffff16436108d79190614b6c565b9050808210610902576040516315ad27c360e01b81526004810183905243602482015260440161087f565b6001600160401b0382111561092d576040516302c6ef8160e11b81526004810183905260240161087f565b60008281526001602090815260408083208287015162ffffff16845290915290205461095d9086908590856123aa565b95945050505050565b6001600160a01b0381166000908152600c602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906109c85750600092915050565b6005546020820151600091600160901b900463ffffffff169060109060ff16601f81106109f7576109f7614b83565b600881049190910154600554610a2a926007166004026101000a90910463ffffffff90811691600160301b900416614b99565b63ffffffff16610a3a9190614b4d565b610a4890633b9aca00614b4d565b905081604001516001600160601b031681610a639190614bbe565b949350505050565b600061095d8585858561257e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314610af15760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c7920636f6f7264696e61746f722063616e2066756c66696c6c00000000604482015260640161087f565b610afc8383836129ab565b505050565b6016546001600160a01b0316610b1f6006546001600160a01b031690565b6001600160a01b0316336001600160a01b03161480610bab5750604051630d629b5f60e31b81526001600160a01b03821690636b14daf890610b6a9033906000903690600401614bff565b602060405180830381865afa158015610b87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bab9190614c24565b610bf75760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c604482015260640161087f565b610bff6129c5565b6005805467ffffffffffffffff60501b1916600160501b63ffffffff89811691820263ffffffff60701b191692909217600160701b8984169081029190911767ffffffffffffffff60901b1916600160901b89851690810263ffffffff60b01b191691909117600160b01b9489169485021762ffffff60d01b1916600160d01b62ffffff89169081029190911790955560408051938452602084019290925290820152606081019190915260808101919091527f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f9060a00160405180910390a1505050505050565b6007546001600160a01b03163314610d3a5760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b604482015260640161087f565b600680546001600160a01b0319808216339081179093556007805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6001600160a01b03818116600090815260146020526040902054163314610dfe5760405162461bcd60e51b815260206004820152601760248201527f4f6e6c792070617965652063616e207769746864726177000000000000000000604482015260640161087f565b6105cd81612d3a565b610e0f612f22565b828114610e5e5760405162461bcd60e51b815260206004820181905260248201527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a65604482015260640161087f565b60005b83811015610fd2576000858583818110610e7d57610e7d614b83565b9050602002016020810190610e929190614173565b90506000848484818110610ea857610ea8614b83565b9050602002016020810190610ebd9190614173565b6001600160a01b038084166000908152601460205260409020549192501680158080610efa5750826001600160a01b0316826001600160a01b0316145b610f3a5760405162461bcd60e51b81526020600482015260116024820152701c185e595948185b1c9958591e481cd95d607a1b604482015260640161087f565b6001600160a01b03848116600090815260146020526040902080546001600160a01b03191685831690811790915590831614610fbb57826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b505050508080610fca90614c46565b915050610e61565b5050505050565b6001600160a01b038181166000908152601560205260409020541633146110425760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792070726f706f736564207061796565732063616e2061636365707400604482015260640161087f565b6001600160a01b0381811660008181526014602090815260408083208054336001600160a01b031980831682179093556015909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a60408051610100808201835260055460ff808216845291810464ffffffffff16602080850191909152600160301b820463ffffffff90811685870152600160501b830481166060860152600160701b830481166080860152600160901b8304811660a0860152600160b01b83041660c0850152600160d01b90910462ffffff1660e0840152336000908152600c825293909320549394509092918c013591166111a35760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161087f565b6009548b35146111ed5760405162461bcd60e51b81526020600482015260156024820152740c6dedcccd2ce88d2cecae6e840dad2e6dac2e8c6d605b1b604482015260640161087f565b6111fb8a8a8a8a8a8a612f77565b8151611208906001614c5f565b60ff1687146112595760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161087f565b8685146112a85760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161087f565b60008a8a6040516112ba929190614c84565b6040519081900381206112d1918e90602001614c94565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a8110156114695760006001858a846020811061132057611320614b83565b61132d91901a601b614c5f565b8f8f8681811061133f5761133f614b83565b905060200201358e8e8781811061135857611358614b83565b9050602002013560405160008152602001604052604051611395949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156113b7573d6000803e3d6000fd5b505060408051601f198101516001600160a01b0381166000908152600d602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506114425760405162461bcd60e51b815260206004820152600f60248201526e39b4b3b730ba3ab9329032b93937b960891b604482015260640161087f565b826020015160080260ff166001901b8401935050808061146190614c46565b915050611301565b5081827e0101010101010101010101010101010101010101010101010101010101010116146114cd5760405162461bcd60e51b815260206004820152601060248201526f323ab83634b1b0ba329039b4b3b732b960811b604482015260640161087f565b506000915061151c9050838d836020020135848e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061257e92505050565b905061152a83828633613014565b505050505050505050505050565b60185481516040516001600160a01b03909216916115599190602001614cb0565b60408051601f198184030181529190528051602090910120601a555050565b6006546001600160a01b03163314806116025750601654604051630d629b5f60e31b81526001600160a01b0390911690636b14daf8906115c19033906000903690600401614bff565b602060405180830381865afa1580156115de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116029190614c24565b61164e5760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c604482015260640161087f565b6000611658613123565b6017546040516370a0823160e01b81523060048201529192506000916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156116a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116ca9190614ccc565b9050818110156117135760405162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e742062616c616e636560601b604482015260640161087f565b6017546001600160a01b031663a9059cbb856117386117328686614b6c565b876132ed565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015611783573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117a79190614c24565b6117c35760405162461bcd60e51b815260040161087f90614ce5565b50505050565b6017546040516370a0823160e01b815230600482015260009182916001600160a01b03909116906370a0823190602401602060405180830381865afa158015611816573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061183a9190614ccc565b90506000611846613123565b90506118528183614d11565b9250505090565b6000806000806118698786613307565b92509250925065ffffffffffff831660009081526004602090815260409182902084518154928601518487015160608801516001600160a01b0316600160481b027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff909216600160381b0291909116670100000000000000600160e81b031962ffffff9093166401000000000266ffffffffffffff1990961663ffffffff90941693909317949094171617919091179055516001600160401b038216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc59061196190889062ffffff91909116815260200190565b60405180910390a250909150505b9392505050565b61197e612f22565b601f8911156119c25760405162461bcd60e51b815260206004820152601060248201526f746f6f206d616e79206f7261636c657360801b604482015260640161087f565b888714611a0a5760405162461bcd60e51b81526020600482015260166024820152750dee4c2c6d8ca40d8cadccee8d040dad2e6dac2e8c6d60531b604482015260640161087f565b88611a16876003614d50565b60ff1610611a665760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161087f565b611a728660ff166135c3565b6040805160e060208c02808301820190935260c082018c815260009383928f918f918291908601908490808284376000920191909152505050908252506040805160208c810282810182019093528c82529283019290918d918d91829185019084908082843760009201919091525050509082525060ff891660208083019190915260408051601f8a0183900483028101830182528981529201919089908990819084018382808284376000920191909152505050908252506001600160401b03861660208083019190915260408051601f8701839004830281018301825286815292019190869086908190840183828082843760009201919091525050509152506005805465ffffffffff00191690559050611b8d6129c5565b600e5460005b81811015611c3e576000600e8281548110611bb057611bb0614b83565b6000918252602082200154600f80546001600160a01b0390921693509084908110611bdd57611bdd614b83565b60009182526020808320909101546001600160a01b039485168352600d82526040808420805461ffff1916905594168252600c90529190912080546dffffffffffffffffffffffffffff191690555080611c3681614c46565b915050611b93565b50611c4b600e6000613f0a565b611c57600f6000613f0a565b60005b825151811015611ed057600d600084600001518381518110611c7e57611c7e614b83565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611cf25760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161087f565b604080518082019091526001815260ff8216602082015283518051600d9160009185908110611d2357611d23614b83565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051600c92919084908110611d8e57611d8e614b83565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611e025760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161087f565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b0316815250600c600085602001518481518110611e4757611e47614b83565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516001600160601b0316620100000262010000600160701b031960ff959095166101000261ff00199315159390931661ffff1990941693909317919091179290921617905580611ec881614c46565b915050611c5a565b5081518051611ee791600e91602090910190613f28565b506020808301518051611efe92600f920190613f28565b5060408201516005805460ff191660ff9092169190911790556007805463ffffffff60c01b198116600160c01b63ffffffff43811682029290921793849055909104811691600091611f5991600160a01b9004166001614d79565b905080600760146101000a81548163ffffffff021916908363ffffffff1602179055506000611fad46308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a00151613608565b9050806009600001819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05838284886000015189602001518a604001518b606001518c608001518d60a0015160405161201099989796959493929190614dda565b60405180910390a1600554600160301b900463ffffffff1660005b8651518110156120855781601082601f811061204957612049614b83565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550808061207d90614c46565b91505061202b565b506120908b8b613663565b50505050505050505050505050505050565b6001600160a01b0381166000908152600c602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906121045750600092915050565b6010816020015160ff16601f811061211e5761211e614b83565b60088104919091015460055461196f926007166004026101000a90910463ffffffff90811691600160301b900416614b99565b6001600160a01b038281166000908152601460205260409020541633146121ba5760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792063757272656e742070617965652063616e20757064617465000000604482015260640161087f565b6001600160a01b03811633036122125760405162461bcd60e51b815260206004820152601760248201527f63616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161087f565b6001600160a01b03808316600090815260156020526040902080548383166001600160a01b031982168117909255909116908114610afc576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b612291612f22565b6105cd81613671565b60008060006122a98787613307565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a6001600160401b031681526020018763ffffffff166001600160601b0316815250905081878a836040516020016123239493929190614e6f565b60408051601f19818403018152828252805160209182012065ffffffffffff8716600090815291829052919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906123849084908a908d908690614e6f565b60405180910390a15090979650505050505050565b6123a1612f22565b6105cd8161371b565b6060826123e35760405163c7d41b1b60e01b815265ffffffffffff861660048201526001600160401b038316602482015260440161087f565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff161115612499576040808601519051634a90778560e01b815261ffff90911660048201526103e8602482015260440161087f565b6000856040015161ffff166001600160401b038111156124bb576124bb6141fb565b6040519080825280602002602001820160405280156124e4578160200160208202803683370190505b50905060005b866040015161ffff168161ffff16101561257357828160405160200161252792919091825260f01b6001600160f01b031916602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff168151811061255657612556614b83565b60209081029190910101528061256b81614f11565b9150506124ea565b509695505050505050565b600080828060200190518101906125959190615103565b64ffffffffff851660208801526040870180519192506125b4826152d7565b63ffffffff1663ffffffff168152505085600560008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff160217905550905050600081604001516001600160401b031640905080826060015114612756576060820151604080840151905163aed0afe560e01b81526004810192909252602482018390526001600160401b0316604482015260640161087f565b6000808360000151516001600160401b03811115612776576127766141fb565b6040519080825280602002602001820160405280156127bb57816020015b60408051808201909152600080825260208201528152602001906001900390816127945790505b50905060005b84515181101561288b576000856000015182815181106127e3576127e3614b83565b602002602001015190506128008187604001518860200151613791565b6040810151515115158061281c57506040810151516020015115155b1561287857604051806040016040528082600001516001600160401b03168152602001826020015162ffffff1681525083838151811061285e5761285e614b83565b6020026020010181905250838061287490614f11565b9450505b508061288381614c46565b9150506127c1565b5060008261ffff166001600160401b038111156128aa576128aa6141fb565b6040519080825280602002602001820160405280156128ef57816020015b60408051808201909152600080825260208201528152602001906001900390816128c85790505b50905060005b8361ffff1681101561294b5782818151811061291357612913614b83565b602002602001015182828151811061292d5761292d614b83565b6020026020010181905250808061294390614c46565b9150506128f5565b50896040015163ffffffff167f7484067466b4f2452757769a8dc9a8b41497154367515673c79386f9f0b74f163387602001518c8c866040516129929594939291906152f0565b60405180910390a2505050506020015195945050505050565b601b80549060006129bb83614c46565b9190505550505050565b601754600554604080516103e08101918290526001600160a01b0390931692600160301b90920463ffffffff1691600091601090601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612a00579050505050505090506000600f805480602002602001604051908101604052809291908181526020018280548015612a9b57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612a7d575b5050505050905060005b8151811015612d2c576000600c6000848481518110612ac657612ac6614b83565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b031690506000600c6000858581518110612b2857612b28614b83565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008483601f8110612b8b57612b8b614b83565b602002015160055490870363ffffffff9081169250600160901b909104168102633b9aca000282018015612d2157600060146000878781518110612bd157612bd1614b83565b6020908102919091018101516001600160a01b03908116835290820192909252604090810160002054905163a9059cbb60e01b815290821660048201819052602482018590529250908a169063a9059cbb906044016020604051808303816000875af1158015612c45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c699190614c24565b612c855760405162461bcd60e51b815260040161087f90614ce5565b878786601f8110612c9857612c98614b83565b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b0316878781518110612cd557612cd5614b83565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c85604051612d1791815260200190565b60405180910390a4505b505050600101612aa5565b50610fd2601083601f613f8d565b6001600160a01b0381166000908152600c60209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b031692810192909252612d97575050565b6000612da283610966565b90508015610afc576001600160a01b038381166000908152601460205260409081902054601754915163a9059cbb60e01b8152908316600482018190526024820185905292919091169063a9059cbb906044016020604051808303816000875af1158015612e14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e389190614c24565b612e545760405162461bcd60e51b815260040161087f90614ce5565b600560000160069054906101000a900463ffffffff166010846020015160ff16601f8110612e8457612e84614b83565b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b038481166000818152600c6020908152604091829020805462010000600160701b0319169055601754915186815291841693851692917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c910160405180910390a450505050565b6006546001600160a01b03163314612f755760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015260640161087f565b565b6000612f84826020614b4d565b612f8f856020614b4d565b612f9b88610144614bbe565b612fa59190614bbe565b612faf9190614bbe565b612fba906000614bbe565b905036811461300b5760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d617463680000000000000000604482015260640161087f565b50505050505050565b600061303b633b9aca003a04866080015163ffffffff16876060015163ffffffff16613b6a565b90506010360260005a905060006130648663ffffffff1685858b60e0015162ffffff1686613b87565b90506000670de0b6b3a76400006001600160c01b03891683026001600160a01b0388166000908152600c602052604090205460c08c01519290910492506201000090046001600160601b039081169163ffffffff16633b9aca0002828401019081168211156130d957505050505050506117c3565b6001600160a01b0388166000908152600c6020526040902080546001600160601b03909216620100000262010000600160701b031990921691909117905550505050505050505050565b600080600f80548060200260200160405190810160405280929190818152602001828054801561317c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161315e575b50508351600554604080516103e08101918290529697509195600160301b90910463ffffffff169450600093509150601090601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116131b55790505050505050905060005b83811015613248578181601f811061321557613215614b83565b60200201516132249084614b99565b6132349063ffffffff1687614bbe565b95508061324081614c46565b9150506131fb565b5060055461326790600160901b900463ffffffff16633b9aca00614b4d565b6132719086614b4d565b945060005b838110156132e557600c600086838151811061329457613294614b83565b6020908102919091018101516001600160a01b03168252810191909152604001600020546132d1906201000090046001600160601b031687614bbe565b9550806132dd81614c46565b915050613276565b505050505090565b6000818310156132fe575081613301565b50805b92915050565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff16111561336157604051634a90778560e01b815261ffff861660048201526103e8602482015260440161087f565b8461ffff16600003613386576040516308fad2a760e01b815260040160405180910390fd5b60006133b27f00000000000000000000000000000000000000000000000000000000000000004361539f565b90506000816133e17f000000000000000000000000000000000000000000000000000000000000000043614bbe565b6133eb9190614b6c565b905060006134197f0000000000000000000000000000000000000000000000000000000000000000836153b3565b905063ffffffff811061343f576040516307b2a52360e41b815260040160405180910390fd5b6040805180820182526002805465ffffffffffff16825282516101008101938490528493600093929160208401916003906008908288855b82829054906101000a900462ffffff1662ffffff168152602001906003019060208260020104928301926001038202915080841161347757905050505091909252505081519192505065ffffffffffff808216106134e857604051630568cab760e31b815260040160405180910390fd5b6134f38160016153c7565b6002805465ffffffffffff191665ffffffffffff9290921691909117905560005b600881101561355a578a62ffffff168360200151826008811061353957613539614b83565b602002015162ffffff161461355a578061355281614c46565b915050613514565b60088110613582576020830151604051630c4f769b60e41b815261087f918d91600401615410565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b806000106105cd5760405162461bcd60e51b815260206004820152601260248201527166206d75737420626520706f73697469766560701b604482015260640161087f565b6000808a8a8a8a8a8a8a8a8a60405160200161362c9998979695949392919061542a565b60408051601f1981840301815291905280516020909101206001600160f01b0316600160f01b179150509998505050505050505050565b61366d8282613beb565b5050565b336001600160a01b038216036136c95760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161087f565b600780546001600160a01b0319166001600160a01b03838116918217909255600654604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6016546001600160a01b03908116908216811461366d57601680546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b82516001600160401b03808416911611156137d557825160405163012d824d60e01b81526001600160401b038085166004830152909116602482015260440161087f565b604083015151516000901580156137f3575060408401515160200151155b1561382b575082516001600160401b031660009081526001602090815260408083208287015162ffffff168452909152902054613885565b836040015160405160200161384091906154b3565b60408051601f19818403018152918152815160209283012086516001600160401b03166000908152600184528281208885015162ffffff168252909352912081905590505b6060840151516000816001600160401b038111156138a5576138a56141fb565b6040519080825280602002602001820160405280156138ce578160200160208202803683370190505b5090506000826001600160401b038111156138eb576138eb6141fb565b6040519080825280601f01601f191660200182016040528015613915576020820181803683370190505b5090506000836001600160401b03811115613932576139326141fb565b60405190808252806020026020018201604052801561396557816020015b60608152602001906001900390816139505790505b5090506000805b85811015613a685760008a60600151828151811061398c5761398c614b83565b602090810291909101015190506000806139b08d600001518e602001518c86613caf565b9150915081156139ef5780868661ffff16815181106139d1576139d1614b83565b602002602001018190525084806139e790614f11565b955050613a1e565b600160f81b878581518110613a0657613a06614b83565b60200101906001600160f81b031916908160001a9053505b8251518851899086908110613a3557613a35614b83565b602002602001019065ffffffffffff16908165ffffffffffff168152505050505080613a6081614c46565b91505061396c565b5060608901515115613b5f5760008161ffff166001600160401b03811115613a9257613a926141fb565b604051908082528060200260200182016040528015613ac557816020015b6060815260200190600190039081613ab05790505b50905060005b8261ffff16811015613b2157838181518110613ae957613ae9614b83565b6020026020010151828281518110613b0357613b03614b83565b60200260200101819052508080613b1990614c46565b915050613acb565b507f47ddf7bb0cbd94c1b43c5097f1352a80db0ceb3696f029d32b24f32cd631d2b7858583604051613b55939291906154e6565b60405180910390a1505b505050505050505050565b60008383811015613b7d57600285850304015b61095d81846132ed565b600081861015613bd95760405162461bcd60e51b815260206004820181905260248201527f6c6566744761732063616e6e6f742065786365656420696e697469616c476173604482015260640161087f565b50633b9aca0094039190910101020290565b610100818114613c1457828282604051635c9d52ef60e11b815260040161087f9392919061559c565b613c1c614024565b8181604051602001613c2e91906155c0565b6040516020818303038152906040525114613c4b57613c4b6155cf565b6040805180820190915260025465ffffffffffff16815260208101613c72858701876155e5565b905280516002805465ffffffffffff191665ffffffffffff9092169190911781556020820151613ca6906003906008614043565b506117c3915050565b6000606081613ce77f00000000000000000000000000000000000000000000000000000000000000006001600160401b0389166153b3565b845160808101516040519293509091600091613d0b918b918b918690602001614e6f565b60408051601f198184030181529181528151602092830120845165ffffffffffff166000908152928390529120549091508114613d795760016040518060400160405280601081526020016f756e6b6e6f776e2063616c6c6261636b60801b81525094509450505050613f01565b815165ffffffffffff16600090815260208181526040808320839055805160808101825263ffffffff8716815262ffffff8c16818401529185015161ffff16828201528401516001600160a01b031660608201528351909190613dde90838b8e6123aa565b6060808401518a5160a00151875192880151604051635a47dd7160e01b815294955091936001600160a01b03851693635a47dd71936001600160601b0390931692613e2e9288919060040161566c565b600060405180830381600088803b158015613e4857600080fd5b5087f193505050508015613e5a575060015b613ee4573d808015613e88576040519150601f19603f3d011682016040523d82523d6000602084013e613e8d565b606091505b50608081511015613eaa57600198509650613f0195505050505050565b60016040518060400160405280600f81526020016e6572726d736720746f6f206c6f6e6760881b8152509850985050505050505050613f01565b600060405180602001604052806000815250975097505050505050505b94509492505050565b50805460008255906000526020600020908101906105cd91906140ca565b828054828255906000526020600020908101928215613f7d579160200282015b82811115613f7d57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613f48565b50613f899291506140ca565b5090565b600483019183908215613f7d5791602002820160005b83821115613fe757835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302613fa3565b80156140175782816101000a81549063ffffffff0219169055600401602081600301049283019260010302613fe7565b5050613f899291506140ca565b6040518061010001604052806008906020820280368337509192915050565b600183019183908215613f7d5791602002820160005b8382111561409b57835183826101000a81548162ffffff021916908362ffffff1602179055509260200192600301602081600201049283019260010302614059565b80156140175782816101000a81549062ffffff021916905560030160208160020104928301926001030261409b565b5b80821115613f8957600081556001016140cb565b65ffffffffffff811681146105cd57600080fd5b60006020828403121561410557600080fd5b813561196f816140df565b600081518084526020808501945080840160005b8381101561414057815187529582019590820190600101614124565b509495945050505050565b60208152600061196f6020830184614110565b6001600160a01b03811681146105cd57600080fd5b60006020828403121561418557600080fd5b813561196f8161415e565b60005b838110156141ab578181015183820152602001614193565b838111156117c35750506000910152565b600081518084526141d4816020860160208601614190565b601f01601f19169290920160200192915050565b60208152600061196f60208301846141bc565b634e487b7160e01b600052604160045260246000fd5b60405161010081016001600160401b0381118282101715614234576142346141fb565b60405290565b604080519081016001600160401b0381118282101715614234576142346141fb565b60405160c081016001600160401b0381118282101715614234576142346141fb565b604051608081016001600160401b0381118282101715614234576142346141fb565b604051602081016001600160401b0381118282101715614234576142346141fb565b604051601f8201601f191681016001600160401b03811182821017156142ea576142ea6141fb565b604052919050565b803560ff8116811461430357600080fd5b919050565b803564ffffffffff8116811461430357600080fd5b803563ffffffff8116811461430357600080fd5b62ffffff811681146105cd57600080fd5b60006001600160401b0382111561435b5761435b6141fb565b50601f01601f191660200190565b600082601f83011261437a57600080fd5b813561438d61438882614342565b6142c2565b8181528460208386010111156143a257600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000808486036101608112156143d757600080fd5b610100808212156143e757600080fd5b6143ef614211565b91506143fa876142f2565b825261440860208801614308565b60208301526144196040880161431d565b604083015261442a6060880161431d565b606083015261443b6080880161431d565b608083015261444c60a0880161431d565b60a083015261445d60c0880161431d565b60c083015260e087013561447081614331565b60e0830152909450850135925061448a6101208601614308565b91506101408501356001600160401b038111156144a657600080fd5b6144b287828801614369565b91505092959194509250565b60006001600160401b038211156144d7576144d76141fb565b5060051b60200190565b6000806000606084860312156144f657600080fd5b8335614501816140df565b92506020848101356001600160401b038082111561451e57600080fd5b818701915087601f83011261453257600080fd5b8135614540614388826144be565b81815260059190911b8301840190848101908a83111561455f57600080fd5b938501935b8285101561457d57843582529385019390850190614564565b96505050604087013592508083111561459557600080fd5b50506145a386828701614369565b9150509250925092565b600080600080600060a086880312156145c557600080fd5b6145ce8661431d565b94506145dc6020870161431d565b93506145ea6040870161431d565b92506145f86060870161431d565b9150608086013561460881614331565b809150509295509295909350565b60008083601f84011261462857600080fd5b5081356001600160401b0381111561463f57600080fd5b6020830191508360208260051b850101111561465a57600080fd5b9250929050565b6000806000806040858703121561467757600080fd5b84356001600160401b038082111561468e57600080fd5b61469a88838901614616565b909650945060208701359150808211156146b357600080fd5b506146c087828801614616565b95989497509550505050565b60008083601f8401126146de57600080fd5b5081356001600160401b038111156146f557600080fd5b60208301915083602082850101111561465a57600080fd5b60008060008060008060008060e0898b03121561472957600080fd5b606089018a81111561473a57600080fd5b899850356001600160401b038082111561475357600080fd5b61475f8c838d016146cc565b909950975060808b013591508082111561477857600080fd5b6147848c838d01614616565b909750955060a08b013591508082111561479d57600080fd5b506147aa8b828c01614616565b999c989b50969995989497949560c00135949350505050565b600060208083850312156147d657600080fd5b82356001600160401b03808211156147ed57600080fd5b908401906040828703121561480157600080fd5b61480961423a565b82358281111561481857600080fd5b61482488828601614369565b825250838301358281111561483857600080fd5b80840193505086601f84011261484d57600080fd5b8235915061485d614388836144be565b82815260059290921b8301840191848101908884111561487c57600080fd5b938501935b8385101561489a57843582529385019390850190614881565b948201949094529695505050505050565b600080604083850312156148be57600080fd5b82356148c98161415e565b946020939093013593505050565b6000602082840312156148e957600080fd5b81356001600160401b038111156148ff57600080fd5b82016080818503121561196f57600080fd5b61ffff811681146105cd57600080fd5b6001600160401b03811681146105cd57600080fd5b803561430381614921565b60008060006060848603121561495657600080fd5b833561496181614911565b9250602084013561497181614921565b9150604084013561498181614331565b809150509250925092565b60008060008060008060008060008060c08b8d0312156149ab57600080fd5b8a356001600160401b03808211156149c257600080fd5b6149ce8e838f01614616565b909c509a5060208d01359150808211156149e757600080fd5b6149f38e838f01614616565b909a509850889150614a0760408e016142f2565b975060608d0135915080821115614a1d57600080fd5b614a298e838f016146cc565b9097509550859150614a3d60808e01614936565b945060a08d0135915080821115614a5357600080fd5b50614a608d828e016146cc565b915080935050809150509295989b9194979a5092959850565b60008060408385031215614a8c57600080fd5b8235614a978161415e565b91506020830135614aa78161415e565b809150509250929050565b600080600080600060a08688031215614aca57600080fd5b8535614ad581614921565b94506020860135614ae581614911565b93506040860135614af581614331565b9250614b036060870161431d565b915060808601356001600160401b03811115614b1e57600080fd5b614b2a88828901614369565b9150509295509295909350565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614b6757614b67614b37565b500290565b600082821015614b7e57614b7e614b37565b500390565b634e487b7160e01b600052603260045260246000fd5b600063ffffffff83811690831681811015614bb657614bb6614b37565b039392505050565b60008219821115614bd157614bd1614b37565b500190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b038416815260406020820181905260009061095d9083018486614bd6565b600060208284031215614c3657600080fd5b8151801515811461196f57600080fd5b600060018201614c5857614c58614b37565b5060010190565b600060ff821660ff84168060ff03821115614c7c57614c7c614b37565b019392505050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60008251614cc2818460208701614190565b9190910192915050565b600060208284031215614cde57600080fd5b5051919050565b602080825260129082015271696e73756666696369656e742066756e647360701b604082015260600190565b60008083128015600160ff1b850184121615614d2f57614d2f614b37565b6001600160ff1b0384018313811615614d4a57614d4a614b37565b50500390565b600060ff821660ff84168160ff0481118215151615614d7157614d71614b37565b029392505050565b600063ffffffff808316818516808303821115614d9857614d98614b37565b01949350505050565b600081518084526020808501945080840160005b838110156141405781516001600160a01b031687529582019590820190600101614db5565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614e0a8184018a614da1565b90508281036080840152614e1e8189614da1565b905060ff871660a084015282810360c0840152614e3b81876141bc565b90506001600160401b03851660e0840152828103610100840152614e5f81856141bc565b9c9b505050505050505050505050565b60006001600160401b03808716835262ffffff8616602084015280851660408401526080606084015265ffffffffffff845116608084015261ffff60208501511660a084015260018060a01b0360408501511660c0840152606084015160c060e0850152614ee16101408501826141bc565b60808601519092166101008501525060a0909301516001600160601b031661012090920191909152509392505050565b600061ffff808316818103614f2857614f28614b37565b6001019392505050565b805161430381614921565b600082601f830112614f4e57600080fd5b8151614f5c61438882614342565b818152846020838601011115614f7157600080fd5b610a63826020830160208701614190565b80516001600160601b038116811461430357600080fd5b600082601f830112614faa57600080fd5b81516020614fba614388836144be565b82815260059290921b84018101918181019086841115614fd957600080fd5b8286015b848110156125735780516001600160401b0380821115614ffc57600080fd5b90880190601f196040838c038201121561501557600080fd5b61501d61423a565b878401518381111561502e57600080fd5b840160c0818e038401121561504257600080fd5b61504a61425c565b925088810151615059816140df565b8352604081015161506981614911565b838a0152606081015161507b8161415e565b604084015260808101518481111561509257600080fd5b6150a08e8b83850101614f3d565b6060850152506150b260a08201614f32565b60808401526150c360c08201614f82565b60a0840152508181526150d860408501614f82565b818901528652505050918301918301614fdd565b80516001600160c01b038116811461430357600080fd5b60006020828403121561511557600080fd5b81516001600160401b038082111561512c57600080fd5b908301906080828603121561514057600080fd5b61514861427e565b82518281111561515757600080fd5b8301601f8101871361516857600080fd5b8051615176614388826144be565b8082825260208201915060208360051b85010192508983111561519857600080fd5b602084015b83811015615298578051878111156151b457600080fd5b850160a0818d03601f190112156151ca57600080fd5b6151d261427e565b60208201516151e081614921565b815260408201516151f081614331565b60208201526040828e03605f1901121561520957600080fd5b6152116142a0565b8d607f84011261522057600080fd5b61522861423a565b808f60a08601111561523957600080fd5b606085015b60a0860181101561525957805183526020928301920161523e565b50825250604082015260a08201518981111561527457600080fd5b6152838e602083860101614f99565b6060830152508452506020928301920161519d565b508452506152ab915050602084016150ec565b60208201526152bc60408401614f32565b60408201526060830151606082015280935050505092915050565b600063ffffffff808316818103614f2857614f28614b37565b6001600160a01b03861681526001600160c01b038516602080830191909152604080830186905264ffffffffff8516606084015260a060808401819052845190840181905260009285810192909160c0860190855b8181101561537857855180516001600160401b0316845285015162ffffff16858401529484019491830191600101615345565b50909b9a5050505050505050505050565b634e487b7160e01b600052601260045260246000fd5b6000826153ae576153ae615389565b500690565b6000826153c2576153c2615389565b500490565b600065ffffffffffff808316818516808303821115614d9857614d98614b37565b8060005b60088110156117c357815162ffffff168452602093840193909101906001016153ec565b62ffffff83168152610120810161196f60208301846153e8565b8981526001600160a01b03891660208201526001600160401b038881166040830152610120606083018190526000916154658483018b614da1565b91508382036080850152615479828a614da1565b915060ff881660a085015283820360c085015261549682886141bc565b90861660e08501528381036101008501529050614e5f81856141bc565b815160408201908260005b60028110156154dd5782518252602092830192909101906001016154be565b50505092915050565b606080825284519082018190526000906020906080840190828801845b8281101561552757815165ffffffffffff1684529284019290840190600101615503565b5050508381038285015261553b81876141bc565b905083810360408501528085518083528383019150838160051b84010184880160005b8381101561558c57601f1986840301855261557a8383516141bc565b9487019492509086019060010161555e565b50909a9950505050505050505050565b6040815260006155b0604083018587614bd6565b9050826020830152949350505050565b610100810161330182846153e8565b634e487b7160e01b600052600160045260246000fd5b60006101008083850312156155f957600080fd5b83601f84011261560857600080fd5b6040518181018181106001600160401b0382111715615629576156296141fb565b60405290830190808583111561563e57600080fd5b845b8381101561566157803561565381614331565b825260209182019101615640565b509095945050505050565b65ffffffffffff8416815260606020820152600061568d6060830185614110565b828103604084015261569f81856141bc565b969550505050505056fea2646970667358221220243615cb3404f188bca412843601f844debdaf6f9fef20e369eb4975f77e4cce64736f6c634300080d0033",
}

var WrapperBeaconCoordinatorABI = WrapperBeaconCoordinatorMetaData.ABI

var WrapperBeaconCoordinatorFuncSigs = WrapperBeaconCoordinatorMetaData.Sigs

var WrapperBeaconCoordinatorBin = WrapperBeaconCoordinatorMetaData.Bin

func DeployWrapperBeaconCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, beaconPeriodBlocksArg *big.Int, keyProvider common.Address, keyID [32]byte) (common.Address, *types.Transaction, *WrapperBeaconCoordinator, error) {
	parsed, err := WrapperBeaconCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrapperBeaconCoordinatorBin), backend, link, beaconPeriodBlocksArg, keyProvider, keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrapperBeaconCoordinator{WrapperBeaconCoordinatorCaller: WrapperBeaconCoordinatorCaller{contract: contract}, WrapperBeaconCoordinatorTransactor: WrapperBeaconCoordinatorTransactor{contract: contract}, WrapperBeaconCoordinatorFilterer: WrapperBeaconCoordinatorFilterer{contract: contract}}, nil
}

type WrapperBeaconCoordinator struct {
	WrapperBeaconCoordinatorCaller
	WrapperBeaconCoordinatorTransactor
	WrapperBeaconCoordinatorFilterer
}

type WrapperBeaconCoordinatorCaller struct {
	contract *bind.BoundContract
}

type WrapperBeaconCoordinatorTransactor struct {
	contract *bind.BoundContract
}

type WrapperBeaconCoordinatorFilterer struct {
	contract *bind.BoundContract
}

type WrapperBeaconCoordinatorSession struct {
	Contract     *WrapperBeaconCoordinator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type WrapperBeaconCoordinatorCallerSession struct {
	Contract *WrapperBeaconCoordinatorCaller
	CallOpts bind.CallOpts
}

type WrapperBeaconCoordinatorTransactorSession struct {
	Contract     *WrapperBeaconCoordinatorTransactor
	TransactOpts bind.TransactOpts
}

type WrapperBeaconCoordinatorRaw struct {
	Contract *WrapperBeaconCoordinator
}

type WrapperBeaconCoordinatorCallerRaw struct {
	Contract *WrapperBeaconCoordinatorCaller
}

type WrapperBeaconCoordinatorTransactorRaw struct {
	Contract *WrapperBeaconCoordinatorTransactor
}

func NewWrapperBeaconCoordinator(address common.Address, backend bind.ContractBackend) (*WrapperBeaconCoordinator, error) {
	contract, err := bindWrapperBeaconCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinator{WrapperBeaconCoordinatorCaller: WrapperBeaconCoordinatorCaller{contract: contract}, WrapperBeaconCoordinatorTransactor: WrapperBeaconCoordinatorTransactor{contract: contract}, WrapperBeaconCoordinatorFilterer: WrapperBeaconCoordinatorFilterer{contract: contract}}, nil
}

func NewWrapperBeaconCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*WrapperBeaconCoordinatorCaller, error) {
	contract, err := bindWrapperBeaconCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorCaller{contract: contract}, nil
}

func NewWrapperBeaconCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapperBeaconCoordinatorTransactor, error) {
	contract, err := bindWrapperBeaconCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorTransactor{contract: contract}, nil
}

func NewWrapperBeaconCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WrapperBeaconCoordinatorFilterer, error) {
	contract, err := bindWrapperBeaconCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorFilterer{contract: contract}, nil
}

func bindWrapperBeaconCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapperBeaconCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapperBeaconCoordinator.Contract.WrapperBeaconCoordinatorCaller.contract.Call(opts, result, method, params...)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.WrapperBeaconCoordinatorTransactor.contract.Transfer(opts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.WrapperBeaconCoordinatorTransactor.contract.Transact(opts, method, params...)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapperBeaconCoordinator.Contract.contract.Call(opts, result, method, params...)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.contract.Transfer(opts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.contract.Transact(opts, method, params...)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) NUMCONFDELAYS() (uint8, error) {
	return _WrapperBeaconCoordinator.Contract.NUMCONFDELAYS(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _WrapperBeaconCoordinator.Contract.NUMCONFDELAYS(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) Count() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.Count(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) Count() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.Count(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) GetBilling(opts *bind.CallOpts) (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "getBilling")

	outstruct := new(struct {
		MaximumGasPriceGwei       uint32
		ReasonableGasPriceGwei    uint32
		ObservationPaymentGjuels  uint32
		TransmissionPaymentGjuels uint32
		AccountingGas             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPriceGwei = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ReasonableGasPriceGwei = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ObservationPaymentGjuels = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TransmissionPaymentGjuels = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _WrapperBeaconCoordinator.Contract.GetBilling(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) GetBilling() (struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
}, error) {
	return _WrapperBeaconCoordinator.Contract.GetBilling(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) GetBillingAccessController() (common.Address, error) {
	return _WrapperBeaconCoordinator.Contract.GetBillingAccessController(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) GetBillingAccessController() (common.Address, error) {
	return _WrapperBeaconCoordinator.Contract.GetBillingAccessController(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) IStartSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "i_StartSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) IStartSlot() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.IStartSlot(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) IStartSlot() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.IStartSlot(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.IBeaconPeriodBlocks(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.IBeaconPeriodBlocks(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "latestConfigDetails")

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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _WrapperBeaconCoordinator.Contract.LatestConfigDetails(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _WrapperBeaconCoordinator.Contract.LatestConfigDetails(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _WrapperBeaconCoordinator.Contract.LatestConfigDigestAndEpoch(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _WrapperBeaconCoordinator.Contract.LatestConfigDigestAndEpoch(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) LinkAvailableForPayment() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.LinkAvailableForPayment(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.LinkAvailableForPayment(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) MaxErrorMsgLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "maxErrorMsgLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) MaxErrorMsgLength() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.MaxErrorMsgLength(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) MaxErrorMsgLength() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.MaxErrorMsgLength(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) MaxNumWords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "maxNumWords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) MaxNumWords() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.MaxNumWords(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) MaxNumWords() (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.MaxNumWords(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) MinDelay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "minDelay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) MinDelay() (uint16, error) {
	return _WrapperBeaconCoordinator.Contract.MinDelay(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) MinDelay() (uint16, error) {
	return _WrapperBeaconCoordinator.Contract.MinDelay(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _WrapperBeaconCoordinator.Contract.OracleObservationCount(&_WrapperBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _WrapperBeaconCoordinator.Contract.OracleObservationCount(&_WrapperBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.OwedPayment(&_WrapperBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _WrapperBeaconCoordinator.Contract.OwedPayment(&_WrapperBeaconCoordinator.CallOpts, transmitterAddress)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) Owner() (common.Address, error) {
	return _WrapperBeaconCoordinator.Contract.Owner(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) Owner() (common.Address, error) {
	return _WrapperBeaconCoordinator.Contract.Owner(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) SKeyID() ([32]byte, error) {
	return _WrapperBeaconCoordinator.Contract.SKeyID(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) SKeyID() ([32]byte, error) {
	return _WrapperBeaconCoordinator.Contract.SKeyID(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) SProvingKeyHash() ([32]byte, error) {
	return _WrapperBeaconCoordinator.Contract.SProvingKeyHash(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _WrapperBeaconCoordinator.Contract.SProvingKeyHash(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrapperBeaconCoordinator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) TypeAndVersion() (string, error) {
	return _WrapperBeaconCoordinator.Contract.TypeAndVersion(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorCallerSession) TypeAndVersion() (string, error) {
	return _WrapperBeaconCoordinator.Contract.TypeAndVersion(&_WrapperBeaconCoordinator.CallOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "acceptOwnership")
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.AcceptOwnership(&_WrapperBeaconCoordinator.TransactOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.AcceptOwnership(&_WrapperBeaconCoordinator.TransactOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.AcceptPayeeship(&_WrapperBeaconCoordinator.TransactOpts, transmitter)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.AcceptPayeeship(&_WrapperBeaconCoordinator.TransactOpts, transmitter)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) ExposeType(opts *bind.TransactOpts, arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "exposeType", arg0)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.ExposeType(&_WrapperBeaconCoordinator.TransactOpts, arg0)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.ExposeType(&_WrapperBeaconCoordinator.TransactOpts, arg0)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) GetRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "getRandomness", requestID)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.GetRandomness(&_WrapperBeaconCoordinator.TransactOpts, requestID)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) GetRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.GetRandomness(&_WrapperBeaconCoordinator.TransactOpts, requestID)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "keyGenerated", kd)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.KeyGenerated(&_WrapperBeaconCoordinator.TransactOpts, kd)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.KeyGenerated(&_WrapperBeaconCoordinator.TransactOpts, kd)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "newKeyRequested")
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) NewKeyRequested() (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.NewKeyRequested(&_WrapperBeaconCoordinator.TransactOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.NewKeyRequested(&_WrapperBeaconCoordinator.TransactOpts)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords, arguments)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.RawFulfillRandomWords(&_WrapperBeaconCoordinator.TransactOpts, requestID, randomWords, arguments)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.RawFulfillRandomWords(&_WrapperBeaconCoordinator.TransactOpts, requestID, randomWords, arguments)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) RequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "requestRandomness", numWords, subID, confirmationDelayArg)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.RequestRandomness(&_WrapperBeaconCoordinator.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.RequestRandomness(&_WrapperBeaconCoordinator.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.RequestRandomnessFulfillment(&_WrapperBeaconCoordinator.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.RequestRandomnessFulfillment(&_WrapperBeaconCoordinator.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "setBilling", maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetBilling(&_WrapperBeaconCoordinator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) SetBilling(maximumGasPriceGwei uint32, reasonableGasPriceGwei uint32, observationPaymentGjuels uint32, transmissionPaymentGjuels uint32, accountingGas *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetBilling(&_WrapperBeaconCoordinator.TransactOpts, maximumGasPriceGwei, reasonableGasPriceGwei, observationPaymentGjuels, transmissionPaymentGjuels, accountingGas)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetBillingAccessController(&_WrapperBeaconCoordinator.TransactOpts, _billingAccessController)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetBillingAccessController(&_WrapperBeaconCoordinator.TransactOpts, _billingAccessController)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetConfig(&_WrapperBeaconCoordinator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetConfig(&_WrapperBeaconCoordinator.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetPayees(&_WrapperBeaconCoordinator.TransactOpts, transmitters, payees)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.SetPayees(&_WrapperBeaconCoordinator.TransactOpts, transmitters, payees)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) TestReport(opts *bind.TransactOpts, hotVars VRFBeaconReportHotVars, configDigest [32]byte, epochAndRound *big.Int, rawReport []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "testReport", hotVars, configDigest, epochAndRound, rawReport)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) TestReport(hotVars VRFBeaconReportHotVars, configDigest [32]byte, epochAndRound *big.Int, rawReport []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.TestReport(&_WrapperBeaconCoordinator.TransactOpts, hotVars, configDigest, epochAndRound, rawReport)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) TestReport(hotVars VRFBeaconReportHotVars, configDigest [32]byte, epochAndRound *big.Int, rawReport []byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.TestReport(&_WrapperBeaconCoordinator.TransactOpts, hotVars, configDigest, epochAndRound, rawReport)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "transferOwnership", to)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.TransferOwnership(&_WrapperBeaconCoordinator.TransactOpts, to)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.TransferOwnership(&_WrapperBeaconCoordinator.TransactOpts, to)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.TransferPayeeship(&_WrapperBeaconCoordinator.TransactOpts, transmitter, proposed)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.TransferPayeeship(&_WrapperBeaconCoordinator.TransactOpts, transmitter, proposed)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.Transmit(&_WrapperBeaconCoordinator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.Transmit(&_WrapperBeaconCoordinator.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.WithdrawFunds(&_WrapperBeaconCoordinator.TransactOpts, recipient, amount)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.WithdrawFunds(&_WrapperBeaconCoordinator.TransactOpts, recipient, amount)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.WithdrawPayment(&_WrapperBeaconCoordinator.TransactOpts, transmitter)
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _WrapperBeaconCoordinator.Contract.WithdrawPayment(&_WrapperBeaconCoordinator.TransactOpts, transmitter)
}

type WrapperBeaconCoordinatorBillingAccessControllerSetIterator struct {
	Event *WrapperBeaconCoordinatorBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorBillingAccessControllerSet)
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
		it.Event = new(WrapperBeaconCoordinatorBillingAccessControllerSet)
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

func (it *WrapperBeaconCoordinatorBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*WrapperBeaconCoordinatorBillingAccessControllerSetIterator, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorBillingAccessControllerSetIterator{contract: _WrapperBeaconCoordinator.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorBillingAccessControllerSet)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseBillingAccessControllerSet(log types.Log) (*WrapperBeaconCoordinatorBillingAccessControllerSet, error) {
	event := new(WrapperBeaconCoordinatorBillingAccessControllerSet)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorBillingSetIterator struct {
	Event *WrapperBeaconCoordinatorBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorBillingSet)
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
		it.Event = new(WrapperBeaconCoordinatorBillingSet)
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

func (it *WrapperBeaconCoordinatorBillingSetIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorBillingSet struct {
	MaximumGasPriceGwei       uint32
	ReasonableGasPriceGwei    uint32
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
	AccountingGas             *big.Int
	Raw                       types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterBillingSet(opts *bind.FilterOpts) (*WrapperBeaconCoordinatorBillingSetIterator, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorBillingSetIterator{contract: _WrapperBeaconCoordinator.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorBillingSet) (event.Subscription, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorBillingSet)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseBillingSet(log types.Log) (*WrapperBeaconCoordinatorBillingSet, error) {
	event := new(WrapperBeaconCoordinatorBillingSet)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorConfigSetIterator struct {
	Event *WrapperBeaconCoordinatorConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorConfigSet)
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
		it.Event = new(WrapperBeaconCoordinatorConfigSet)
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

func (it *WrapperBeaconCoordinatorConfigSetIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorConfigSet struct {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterConfigSet(opts *bind.FilterOpts) (*WrapperBeaconCoordinatorConfigSetIterator, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorConfigSetIterator{contract: _WrapperBeaconCoordinator.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorConfigSet) (event.Subscription, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorConfigSet)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseConfigSet(log types.Log) (*WrapperBeaconCoordinatorConfigSet, error) {
	event := new(WrapperBeaconCoordinatorConfigSet)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorNewTransmissionIterator struct {
	Event *WrapperBeaconCoordinatorNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorNewTransmission)
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
		it.Event = new(WrapperBeaconCoordinatorNewTransmission)
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

func (it *WrapperBeaconCoordinatorNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorNewTransmission struct {
	AggregatorRoundId uint32
	Transmitter       common.Address
	JuelsPerFeeCoin   *big.Int
	ConfigDigest      [32]byte
	EpochAndRound     *big.Int
	OutputsServed     []VRFBeaconReportOutputServed
	Raw               types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32) (*WrapperBeaconCoordinatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorNewTransmissionIterator{contract: _WrapperBeaconCoordinator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorNewTransmission, aggregatorRoundId []uint32) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorNewTransmission)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseNewTransmission(log types.Log) (*WrapperBeaconCoordinatorNewTransmission, error) {
	event := new(WrapperBeaconCoordinatorNewTransmission)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorOraclePaidIterator struct {
	Event *WrapperBeaconCoordinatorOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorOraclePaid)
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
		it.Event = new(WrapperBeaconCoordinatorOraclePaid)
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

func (it *WrapperBeaconCoordinatorOraclePaidIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*WrapperBeaconCoordinatorOraclePaidIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorOraclePaidIterator{contract: _WrapperBeaconCoordinator.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	var linkTokenRule []interface{}
	for _, linkTokenItem := range linkToken {
		linkTokenRule = append(linkTokenRule, linkTokenItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorOraclePaid)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseOraclePaid(log types.Log) (*WrapperBeaconCoordinatorOraclePaid, error) {
	event := new(WrapperBeaconCoordinatorOraclePaid)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorOwnershipTransferRequestedIterator struct {
	Event *WrapperBeaconCoordinatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorOwnershipTransferRequested)
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
		it.Event = new(WrapperBeaconCoordinatorOwnershipTransferRequested)
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

func (it *WrapperBeaconCoordinatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrapperBeaconCoordinatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorOwnershipTransferRequestedIterator{contract: _WrapperBeaconCoordinator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorOwnershipTransferRequested)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*WrapperBeaconCoordinatorOwnershipTransferRequested, error) {
	event := new(WrapperBeaconCoordinatorOwnershipTransferRequested)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorOwnershipTransferredIterator struct {
	Event *WrapperBeaconCoordinatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorOwnershipTransferred)
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
		it.Event = new(WrapperBeaconCoordinatorOwnershipTransferred)
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

func (it *WrapperBeaconCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrapperBeaconCoordinatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorOwnershipTransferredIterator{contract: _WrapperBeaconCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorOwnershipTransferred)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*WrapperBeaconCoordinatorOwnershipTransferred, error) {
	event := new(WrapperBeaconCoordinatorOwnershipTransferred)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorPayeeshipTransferRequestedIterator struct {
	Event *WrapperBeaconCoordinatorPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorPayeeshipTransferRequested)
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
		it.Event = new(WrapperBeaconCoordinatorPayeeshipTransferRequested)
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

func (it *WrapperBeaconCoordinatorPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*WrapperBeaconCoordinatorPayeeshipTransferRequestedIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorPayeeshipTransferRequestedIterator{contract: _WrapperBeaconCoordinator.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var proposedRule []interface{}
	for _, proposedItem := range proposed {
		proposedRule = append(proposedRule, proposedItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorPayeeshipTransferRequested)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParsePayeeshipTransferRequested(log types.Log) (*WrapperBeaconCoordinatorPayeeshipTransferRequested, error) {
	event := new(WrapperBeaconCoordinatorPayeeshipTransferRequested)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorPayeeshipTransferredIterator struct {
	Event *WrapperBeaconCoordinatorPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorPayeeshipTransferred)
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
		it.Event = new(WrapperBeaconCoordinatorPayeeshipTransferred)
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

func (it *WrapperBeaconCoordinatorPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*WrapperBeaconCoordinatorPayeeshipTransferredIterator, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorPayeeshipTransferredIterator{contract: _WrapperBeaconCoordinator.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var transmitterRule []interface{}
	for _, transmitterItem := range transmitter {
		transmitterRule = append(transmitterRule, transmitterItem)
	}
	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorPayeeshipTransferred)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParsePayeeshipTransferred(log types.Log) (*WrapperBeaconCoordinatorPayeeshipTransferred, error) {
	event := new(WrapperBeaconCoordinatorPayeeshipTransferred)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorRandomWordsFulfilledIterator struct {
	Event *WrapperBeaconCoordinatorRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorRandomWordsFulfilled)
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
		it.Event = new(WrapperBeaconCoordinatorRandomWordsFulfilled)
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

func (it *WrapperBeaconCoordinatorRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*WrapperBeaconCoordinatorRandomWordsFulfilledIterator, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorRandomWordsFulfilledIterator{contract: _WrapperBeaconCoordinator.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorRandomWordsFulfilled)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseRandomWordsFulfilled(log types.Log) (*WrapperBeaconCoordinatorRandomWordsFulfilled, error) {
	event := new(WrapperBeaconCoordinatorRandomWordsFulfilled)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorRandomnessFulfillmentRequestedIterator struct {
	Event *WrapperBeaconCoordinatorRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorRandomnessFulfillmentRequested)
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
		it.Event = new(WrapperBeaconCoordinatorRandomnessFulfillmentRequested)
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

func (it *WrapperBeaconCoordinatorRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorRandomnessFulfillmentRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  uint64
	Callback               VRFBeaconTypesCallback
	Raw                    types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts) (*WrapperBeaconCoordinatorRandomnessFulfillmentRequestedIterator, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorRandomnessFulfillmentRequestedIterator{contract: _WrapperBeaconCoordinator.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorRandomnessFulfillmentRequested) (event.Subscription, error) {

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "RandomnessFulfillmentRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorRandomnessFulfillmentRequested)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*WrapperBeaconCoordinatorRandomnessFulfillmentRequested, error) {
	event := new(WrapperBeaconCoordinatorRandomnessFulfillmentRequested)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WrapperBeaconCoordinatorRandomnessRequestedIterator struct {
	Event *WrapperBeaconCoordinatorRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WrapperBeaconCoordinatorRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrapperBeaconCoordinatorRandomnessRequested)
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
		it.Event = new(WrapperBeaconCoordinatorRandomnessRequested)
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

func (it *WrapperBeaconCoordinatorRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *WrapperBeaconCoordinatorRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WrapperBeaconCoordinatorRandomnessRequested struct {
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	Raw                    types.Log
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, nextBeaconOutputHeight []uint64) (*WrapperBeaconCoordinatorRandomnessRequestedIterator, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.FilterLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return &WrapperBeaconCoordinatorRandomnessRequestedIterator{contract: _WrapperBeaconCoordinator.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *WrapperBeaconCoordinatorRandomnessRequested, nextBeaconOutputHeight []uint64) (event.Subscription, error) {

	var nextBeaconOutputHeightRule []interface{}
	for _, nextBeaconOutputHeightItem := range nextBeaconOutputHeight {
		nextBeaconOutputHeightRule = append(nextBeaconOutputHeightRule, nextBeaconOutputHeightItem)
	}

	logs, sub, err := _WrapperBeaconCoordinator.contract.WatchLogs(opts, "RandomnessRequested", nextBeaconOutputHeightRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WrapperBeaconCoordinatorRandomnessRequested)
				if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_WrapperBeaconCoordinator *WrapperBeaconCoordinatorFilterer) ParseRandomnessRequested(log types.Log) (*WrapperBeaconCoordinatorRandomnessRequested, error) {
	event := new(WrapperBeaconCoordinatorRandomnessRequested)
	if err := _WrapperBeaconCoordinator.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

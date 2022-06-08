package dkg

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

type KeyDataStructKeyData struct {
	PublicKey []byte
	Hashes    [][32]byte
}

var ConfirmedOwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161045538038061045583398101604081905261002f9161016e565b8060006001600160a01b03821661008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100bd576100bd816100c5565b50505061019e565b336001600160a01b0382160361011d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561018057600080fd5b81516001600160a01b038116811461019757600080fd5b9392505050565b6102a8806101ad6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c634300080d000a",
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
	Bin: "0x608060405234801561001057600080fd5b5060405161047038038061047083398101604081905261002f91610186565b6001600160a01b03821661008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100ba576100ba816100c1565b50506101b9565b336001600160a01b038216036101195760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610081565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b038116811461018157600080fd5b919050565b6000806040838503121561019957600080fd5b6101a28361016a565b91506101b06020840161016a565b90509250929050565b6102a8806101c86000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c634300080d000a",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractDKGClient\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"DKGClientError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"indexed\":false,\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"KeyGenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"addClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_configDigest\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"removeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000a5565b50505062000150565b336001600160a01b03821603620000ff5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6127fd80620001606000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063b1dc65a411610066578063b1dc65a414610223578063c3105a6b14610236578063e3d0e71214610256578063f2fde38b1461026957600080fd5b80638da5cb5b146101cb5780639201de55146101e6578063afcb95d7146101f957600080fd5b80635429a79e116100c85780635429a79e1461016e57806379ba5097146101835780637bf1ffc51461018b57806381ff70481461019e57600080fd5b80630bc643e8146100ef578063181f5a771461011957806339614e4f1461015b575b600080fd5b6101026100fd366004611e68565b61027c565b60405160ff90911681526020015b60405180910390f35b60408051808201909152600981527f444b4720302e302e31000000000000000000000000000000000000000000000060208201525b6040516101109190611edf565b61014e610169366004611fb7565b6102ab565b61018161017c366004612009565b61042d565b005b61018161066e565b610181610199366004612009565b610724565b600754600554604080516000815264010000000090930463ffffffff166020840152820152606001610110565b6000546040516001600160a01b039091168152602001610110565b61014e6101f4366004612039565b61076b565b6005546004546040805160008152602081019390935263ffffffff90911690820152606001610110565b61018161023136600461209e565b6107f7565b610249610244366004612183565b610943565b60405161011091906121a5565b6101816102643660046122ad565b610a6b565b61018161027736600461237a565b6111e7565b6000600a8260ff16101561029b576102958260306123ad565b92915050565b6102958260576123ad565b919050565b6060600080835160026102be91906123d2565b67ffffffffffffffff8111156102d6576102d6611ef2565b6040519080825280601f01601f191660200182016040528015610300576020820181803683370190505b509050600091505b80518260ff161015610426576000846103226002856123f1565b60ff168151811061033557610335612421565b60209101015160f81c600f16905060006004866103536002876123f1565b60ff168151811061036657610366612421565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c905061039d8161027c565b60f81b838560ff16815181106103b5576103b5612421565b60200101906001600160f81b031916908160001a9053506103d78460016123ad565b93506103e28261027c565b60f81b838560ff16815181106103fa576103fa612421565b60200101906001600160f81b031916908160001a9053505050818061041e90612437565b925050610308565b9392505050565b6104356111fb565b60008281526002602090815260408083208054825181850281018501909352808352919290919083018282801561049557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610477575b505050505090506000815167ffffffffffffffff8111156104b8576104b8611ef2565b6040519080825280602002602001820160405280156104e1578160200160208202803683370190505b5090506000805b835181101561058457846001600160a01b031684828151811061050d5761050d612421565b60200260200101516001600160a01b03161461056457848361052f8484612456565b8151811061053f5761053f612421565b60200260200101906001600160a01b031690816001600160a01b031681525050610572565b8161056e8161246d565b9250505b8061057c8161246d565b9150506104e8565b5060008184516105949190612456565b67ffffffffffffffff8111156105ac576105ac611ef2565b6040519080825280602002602001820160405280156105d5578160200160208202803683370190505b50905060005b8285516105e89190612456565b8110156106455783818151811061060157610601612421565b602002602001015182828151811061061b5761061b612421565b6001600160a01b03909216602092830291909101909101528061063d8161246d565b9150506105db565b506000868152600260209081526040909120825161066592840190611d10565b50505050505050565b6001546001600160a01b031633146106cd5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61072c6111fb565b600091825260026020908152604083208054600181018255908452922090910180546001600160a01b0319166001600160a01b03909216919091179055565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156107ed578381602081106107ae576107ae612421565b1a60f81b8282815181106107c4576107c4612421565b60200101906001600160f81b031916908160001a905350806107e58161246d565b915050610793565b50610426816102ab565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916108479184918491908e908e908190840183828082843760009201919091525061125792505050565b6040805183815263ffffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260055480825260065460ff808216602085015261010090910416928201929092529083146109035760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016106c4565b6109118b8b8b8b8b8b6114b4565b6109228c8c8c8c8c8c8c8c89611548565b50505063ffffffff811061093857610938612486565b505050505050505050565b604080518082019091526060808252602082015260008381526003602090815260408083208584529091529081902081518083019092528054829082906109899061249c565b80601f01602080910402602001604051908101604052809291908181526020018280546109b59061249c565b8015610a025780601f106109d757610100808354040283529160200191610a02565b820191906000526020600020905b8154815290600101906020018083116109e557829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610a5a57602002820191906000526020600020905b815481526020019060010190808311610a46575b505050505081525050905092915050565b855185518560ff16601f831115610ac45760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016106c4565b60008111610b145760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016106c4565b818314610b885760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016106c4565b610b938160036123d2565b8311610be15760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016106c4565b610be96111fb565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60095415610d3957600954600090610c4190600190612456565b9050600060098281548110610c5857610c58612421565b6000918252602082200154600a80546001600160a01b0390921693509084908110610c8557610c85612421565b60009182526020808320909101546001600160a01b03858116845260089092526040808420805461ffff1990811690915592909116808452922080549091169055600980549192509080610cdb57610cdb6124d6565b600082815260209020810160001990810180546001600160a01b0319169055019055600a805480610d0e57610d0e6124d6565b600082815260209020810160001990810180546001600160a01b031916905501905550610c27915050565b60005b8151518110156110785760006008600084600001518481518110610d6257610d62612421565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610d9f57610d9f6124ec565b14610dec5760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016106c4565b6040805180820190915260ff82168152600160208201528251805160089160009185908110610e1d57610e1d612421565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610e7657610e766124ec565b021790555060009150610e869050565b6008600084602001518481518110610ea057610ea0612421565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610edd57610edd6124ec565b14610f2a5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016106c4565b6040805180820190915260ff821681526020810160028152506008600084602001518481518110610f5d57610f5d612421565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610fb657610fb66124ec565b021790555050825180516009925083908110610fd457610fd4612421565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b03909316929092179091558201518051600a91908390811061102b5761102b612421565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b03909216919091179055806110708161246d565b915050610d3c565b5060408101516006805460ff191660ff9092169190911790556007805467ffffffff0000000019811664010000000063ffffffff4381168202928317855590830481169360019390926000926110d5928692908216911617612502565b92506101000a81548163ffffffff021916908363ffffffff16021790555060006111364630600760009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a001516119d7565b6005819055835180516006805460ff9092166101000261ff00199092169190911790556007546020860151604080880151606089015160808a015160a08b015193519798507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05976111be978b978b9763ffffffff90911696919590949093909290919061256e565b60405180910390a16111d98360400151846060015183611a64565b505050505050505050505050565b6111ef6111fb565b6111f881611c67565b50565b6000546001600160a01b031633146112555760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106c4565b565b6000606080838060200190518101906112709190612604565b60408051808201825283815260208082018490526000868152600282528381208054855181850281018501909652808652979a50959850939650909492939192908301828280156112ea57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112cc575b5050505050905060005b81518110156114075781818151811061130f5761130f612421565b60200260200101516001600160a01b031663bf2732c7846040518263ffffffff1660e01b815260040161134291906121a5565b600060405180830381600087803b15801561135c57600080fd5b505af192505050801561136d575060015b6113f5573d80801561139b576040519150601f19603f3d011682016040523d82523d6000602084013e6113a0565b606091505b507f116391732f5df106193bda7cedf1728f3b07b62f6cdcdd611c9eeec44efcae548383815181106113d4576113d4612421565b6020026020010151826040516113eb929190612702565b60405180910390a1505b806113ff8161246d565b9150506112f4565b5060008581526003602090815260408083208b84528252909120835180518593611435928492910190611d75565b50602082810151805161144e9260018501920190611de9565b5090505084887fc8db841f5b2231ccf7190311f440aa197b161e369f3b40b023508160cc5556568460405161148391906121a5565b60405180910390a350506004805460089690961c63ffffffff1663ffffffff19909616959095179094555050505050565b60006114c18260206123d2565b6114cc8560206123d2565b6114d888610144612724565b6114e29190612724565b6114ec9190612724565b6114f7906000612724565b90503681146106655760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016106c4565b600060028260200151836040015161156091906123ad565b61156a91906123f1565b6115759060016123ad565b60408051600180825281830190925260ff929092169250600091906020820181803683370190505090508160f81b816000815181106115b6576115b6612421565b60200101906001600160f81b031916908160001a9053508682146115d9826102ab565b906115f75760405162461bcd60e51b81526004016106c49190611edf565b508685146116475760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016106c4565b3360009081526008602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561168a5761168a6124ec565b600281111561169b5761169b6124ec565b90525090506002816020015160028111156116b8576116b86124ec565b1480156116f25750600a816000015160ff16815481106116da576116da612421565b6000918252602090912001546001600160a01b031633145b61173e5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016106c4565b5050506000888860405161175392919061273c565b60405190819003812061176a918c9060200161274c565b60405160208183030381529060405280519060200120905061178a611e23565b604080518082019091526000808252602082015260005b888110156119c85760006001858884602081106117c0576117c0612421565b6117cd91901a601b6123ad565b8d8d868181106117df576117df612421565b905060200201358c8c878181106117f8576117f8612421565b9050602002013560405160008152602001604052604051611835949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611857573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526008602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156118ac576118ac6124ec565b60028111156118bd576118bd6124ec565b90525092506001836020015160028111156118da576118da6124ec565b146119275760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016106c4565b8251849060ff16601f811061193e5761193e612421565b6020020151156119905760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016106c4565b600184846000015160ff16601f81106119ab576119ab612421565b9115156020909202015250806119c08161246d565b9150506117a1565b50505050505050505050505050565b6000808a8a8a8a8a8a8a8a8a6040516020016119fb99989796959493929190612768565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6000808351602014611ab85760405162461bcd60e51b815260206004820152601e60248201527f77726f6e67206c656e67746820666f72206f6e636861696e436f6e666967000060448201526064016106c4565b60208401519150808203611b0e5760405162461bcd60e51b815260206004820152601460248201527f6661696c656420746f20636f7079206b6579494400000000000000000000000060448201526064016106c4565b604080518082019091526060808252602082015260008381526003602090815260408083208784528252909120825180518493611b4f928492910190611d75565b506020828101518051611b689260018501920190611de9565b505050600083815260026020908152604080832080548251818502810185019093528083529192909190830182828015611bcb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611bad575b5050505050905060005b8151811015611c5d57818181518110611bf057611bf0612421565b60200260200101516001600160a01b03166355e487496040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611c3257600080fd5b505af1158015611c46573d6000803e3d6000fd5b505050508080611c559061246d565b915050611bd5565b5050505050505050565b336001600160a01b03821603611cbf5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106c4565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215611d65579160200282015b82811115611d6557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611d30565b50611d71929150611e42565b5090565b828054611d819061249c565b90600052602060002090601f016020900481019282611da35760008555611d65565b82601f10611dbc57805160ff1916838001178555611d65565b82800160010185558215611d65579182015b82811115611d65578251825591602001919060010190611dce565b828054828255906000526020600020908101928215611d655791602002820182811115611d65578251825591602001919060010190611dce565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115611d715760008155600101611e43565b803560ff811681146102a657600080fd5b600060208284031215611e7a57600080fd5b61042682611e57565b60005b83811015611e9e578181015183820152602001611e86565b83811115611ead576000848401525b50505050565b60008151808452611ecb816020860160208601611e83565b601f01601f19169290920160200192915050565b6020815260006104266020830184611eb3565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611f3157611f31611ef2565b604052919050565b600067ffffffffffffffff821115611f5357611f53611ef2565b50601f01601f191660200190565b600082601f830112611f7257600080fd5b8135611f85611f8082611f39565b611f08565b818152846020838601011115611f9a57600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611fc957600080fd5b813567ffffffffffffffff811115611fe057600080fd5b611fec84828501611f61565b949350505050565b6001600160a01b03811681146111f857600080fd5b6000806040838503121561201c57600080fd5b82359150602083013561202e81611ff4565b809150509250929050565b60006020828403121561204b57600080fd5b5035919050565b60008083601f84011261206457600080fd5b50813567ffffffffffffffff81111561207c57600080fd5b6020830191508360208260051b850101111561209757600080fd5b9250929050565b60008060008060008060008060e0898b0312156120ba57600080fd5b606089018a8111156120cb57600080fd5b8998503567ffffffffffffffff808211156120e557600080fd5b818b0191508b601f8301126120f957600080fd5b81358181111561210857600080fd5b8c602082850101111561211a57600080fd5b6020830199508098505060808b013591508082111561213857600080fd5b6121448c838d01612052565b909750955060a08b013591508082111561215d57600080fd5b5061216a8b828c01612052565b999c989b50969995989497949560c00135949350505050565b6000806040838503121561219657600080fd5b50508035926020909101359150565b6000602080835283516040828501526121c16060850182611eb3565b85830151858203601f19016040870152805180835290840192506000918401905b8083101561220257835182529284019260019290920191908401906121e2565b509695505050505050565b600067ffffffffffffffff82111561222757612227611ef2565b5060051b60200190565b600082601f83011261224257600080fd5b81356020612252611f808361220d565b82815260059290921b8401810191818101908684111561227157600080fd5b8286015b8481101561220257803561228881611ff4565b8352918301918301612275565b803567ffffffffffffffff811681146102a657600080fd5b60008060008060008060c087890312156122c657600080fd5b863567ffffffffffffffff808211156122de57600080fd5b6122ea8a838b01612231565b9750602089013591508082111561230057600080fd5b61230c8a838b01612231565b965061231a60408a01611e57565b9550606089013591508082111561233057600080fd5b61233c8a838b01611f61565b945061234a60808a01612295565b935060a089013591508082111561236057600080fd5b5061236d89828a01611f61565b9150509295509295509295565b60006020828403121561238c57600080fd5b813561042681611ff4565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff038211156123ca576123ca612397565b019392505050565b60008160001904831182151516156123ec576123ec612397565b500290565b600060ff83168061241257634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff810361244d5761244d612397565b60010192915050565b60008282101561246857612468612397565b500390565b60006001820161247f5761247f612397565b5060010190565b634e487b7160e01b600052600160045260246000fd5b600181811c908216806124b057607f821691505b6020821081036124d057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600063ffffffff80831681851680830382111561252157612521612397565b01949350505050565b600081518084526020808501945080840160005b838110156125635781516001600160a01b03168752958201959082019060010161253e565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261259e8184018a61252a565b905082810360808401526125b2818961252a565b905060ff871660a084015282810360c08401526125cf8187611eb3565b905067ffffffffffffffff851660e08401528281036101008401526125f48185611eb3565b9c9b505050505050505050505050565b60008060006060848603121561261957600080fd5b8351925060208085015167ffffffffffffffff8082111561263957600080fd5b818701915087601f83011261264d57600080fd5b815161265b611f8082611f39565b818152898583860101111561266f57600080fd5b61267e82868301878701611e83565b60408901519096509250508082111561269657600080fd5b508501601f810187136126a857600080fd5b80516126b6611f808261220d565b81815260059190911b820183019083810190898311156126d557600080fd5b928401925b828410156126f3578351825292840192908401906126da565b80955050505050509250925092565b6001600160a01b0383168152604060208201526000611fec6040830184611eb3565b6000821982111561273757612737612397565b500190565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526127a28285018b61252a565b915083820360808501526127b6828a61252a565b915060ff881660a085015283820360c08501526127d38288611eb3565b90861660e085015283810361010085015290506125f48185611eb356fea164736f6c634300080d000a",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610518806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630bc643e81461004657806339614e4f146100705780639201de5514610090575b600080fd5b6100596100543660046102db565b6100a3565b60405160ff90911681526020015b60405180910390f35b61008361007e366004610314565b6100cd565b60405161006791906103c5565b61008361009e36600461041a565b61024f565b6000600a8260ff1610156100c2576100bc826030610449565b92915050565b6100bc826057610449565b6060600080835160026100e0919061046e565b67ffffffffffffffff8111156100f8576100f86102fe565b6040519080825280601f01601f191660200182016040528015610122576020820181803683370190505b509050600091505b80518260ff1610156102485760008461014460028561048d565b60ff1681518110610157576101576104bd565b60209101015160f81c600f169050600060048661017560028761048d565b60ff1681518110610188576101886104bd565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c90506101bf816100a3565b60f81b838560ff16815181106101d7576101d76104bd565b60200101906001600160f81b031916908160001a9053506101f9846001610449565b9350610204826100a3565b60f81b838560ff168151811061021c5761021c6104bd565b60200101906001600160f81b031916908160001a90535050508180610240906104d3565b92505061012a565b9392505050565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156102d157838160208110610292576102926104bd565b1a60f81b8282815181106102a8576102a86104bd565b60200101906001600160f81b031916908160001a905350806102c9816104f2565b915050610277565b50610248816100cd565b6000602082840312156102ed57600080fd5b813560ff8116811461024857600080fd5b634e487b7160e01b600052604160045260246000fd5b60006020828403121561032657600080fd5b813567ffffffffffffffff8082111561033e57600080fd5b818401915084601f83011261035257600080fd5b813581811115610364576103646102fe565b604051601f8201601f19908116603f0116810190838211818310171561038c5761038c6102fe565b816040528281528760208487010111156103a557600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156103f2578581018301518582016040015282016103d6565b81811115610404576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561042c57600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff0382111561046657610466610433565b019392505050565b600081600019048311821515161561048857610488610433565b500290565b600060ff8316806104ae57634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036104e9576104e9610433565b60010192915050565b60006001820161050457610504610433565b506001019056fea164736f6c634300080d000a",
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
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6102a8806101576000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806379ba5097146100465780638da5cb5b14610050578063f2fde38b1461006f575b600080fd5b61004e610082565b005b600054604080516001600160a01b039092168252519081900360200190f35b61004e61007d36600461026b565b610145565b6001546001600160a01b031633146100e15760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61014d610159565b610156816101b5565b50565b6000546001600160a01b031633146101b35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016100d8565b565b336001600160a01b0382160361020d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016100d8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561027d57600080fd5b81356001600160a01b038116811461029457600080fd5b939250505056fea164736f6c634300080d000a",
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

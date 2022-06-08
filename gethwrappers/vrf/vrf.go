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

var ECCArithmeticMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50601680601d6000396000f3fe6080604052600080fdfea164736f6c634300080d000a",
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
	Bin: "0x608060405234801561001057600080fd5b50610df7806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e95d957714610030575b600080fd5b61004361003e366004610c8d565b610059565b6040516100509190610ccc565b60405180910390f35b610061610bbb565b600061006c856100a7565b9050600061008182825b6020020151866101b0565b90506000610090836001610076565b905061009c8282610928565b979650505050505050565b6100af610bd3565b6000805b60028210156101a95760408051602081018690520160408051601f1981840301815291905280516020909101209350836100fc600080516020610dcb8339815191526005610d2b565b8110156101405761011b600080516020610dcb83398151915282610d60565b84846002811061012d5761012d610cff565b60200201528261013c81610d82565b9350505b8161014a81610d82565b92505060208211156101a35760405162461bcd60e51b815260206004820152601960248201527f617474656d7074656420746f6f206d616e79206861736865730000000000000060448201526064015b60405180910390fd5b506100b3565b5050919050565b6101b8610bbb565b826000036101fc5750604080516060810182527759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe60208201908152600292820192909252908152610922565b610215600080516020610dcb8339815191526005610d2b565b83106102635760405162461bcd60e51b815260206004820152601660248201527f74206e6f74206120756e69666f726d2073616d706c6500000000000000000000604482015260640161019a565b6000600080516020610dcb833981519152848509905060008161028860036001610d9b565b6102929190610d9b565b9050600080516020610dcb833981519152845182096001146102f65760405162461bcd60e51b815260206004820152601760248201527f77726f6e6720696e766572736520666f722064656e6f6d000000000000000000604482015260640161019a565b6000600080516020610dcb8339815191528551600080516020610dcb8339815191528577b3c4d79d41a91759a9e4c7e359b6b89eaec68e62effffffd090990506000600080516020610dcb83398151915261035f83600080516020610dcb833981519152610db3565b7759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe0890506000610388600289610d60565b9050816000600080516020610dcb8339815191526003600080516020610dcb8339815191528085860985090890506000600080516020610dcb83398151915260408b0151800990508181036104b957600080516020610dcb8339815191528a60400151106104385760405162461bcd60e51b815260206004820152600c60248201527f793120746f6f206c617267650000000000000000000000000000000000000000604482015260640161019a565b8360028b6040015161044a9190610d60565b146104975760405162461bcd60e51b815260206004820152601860248201527f793120706172697479206d757374206d61746368207427730000000000000000604482015260640161019a565b885183905260408a0151895160015b6020020152506109229650505050505050565b6104d182600080516020610dcb833981519152610db3565b811461051f5760405162461bcd60e51b815260206004820152601860248201527f7931213d70736575646f20737172206f662078315e332b420000000000000000604482015260640161019a565b5050506000600080516020610dcb8339815191528061054057610540610d4a565b8360010861055c90600080516020610dcb833981519152610db3565b90506000600080516020610dcb8339815191526003600080516020610dcb8339815191528085860985090890506000600080516020610dcb83398151915260608b01518009905081810361066c57600080516020610dcb8339815191528a60600151106105fa5760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b604482015260640161019a565b8360028b6060015161060c9190610d60565b146106595760405162461bcd60e51b815260206004820152601860248201527f793220706172697479206d757374206d61746368207427730000000000000000604482015260640161019a565b885183905260608a0151895160016104a6565b61068482600080516020610dcb833981519152610db3565b81146106d25760405162461bcd60e51b815260206004820152601860248201527f7932213d70736575646f20737172206f662078325e332b420000000000000000604482015260640161019a565b505050600080516020610dcb833981519152806106f1576106f1610d4a565b876020015186096001146107475760405162461bcd60e51b815260206004820152601c60248201527f74496e76537175617265642a742a2a3220213d3d2031206d6f64205000000000604482015260640161019a565b600080516020610dcb8339815191527f2042def740cbc01bd03583cf0100e593ba56470b9af68708d2c05d6490535385600080516020610dcb83398151915260208a0151600080516020610dcb833981519152888909090992506000600080516020610dcb8339815191526107ca85600080516020610dcb833981519152610db3565b60010890506000600080516020610dcb8339815191526003600080516020610dcb83398151915280858609850908905080600080516020610dcb83398151915260808b015180091461085e5760405162461bcd60e51b815260206004820152601c60248201527f646964206e6f74206f627461696e206120637572766520706f696e7400000000604482015260640161019a565b600080516020610dcb8339815191528960800151106108ae5760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b604482015260640161019a565b8260028a608001516108c09190610d60565b1461090d5760405162461bcd60e51b815260206004820152601860248201527f793320706172697479206d757374206d61746368207427730000000000000000604482015260640161019a565b50865152505050506080830151825160200152505b92915050565b610930610bbb565b600061093c84846109a7565b80515190915015801590610954575080516020015115155b6109a05760405162461bcd60e51b815260206004820152601b60248201527f6164646731206661696c65643a207a65726f206f7264696e6174650000000000604482015260640161019a565b9392505050565b6109af610bbb565b6109b883610a6b565b6109c182610a6b565b6109c9610bf1565b835151815283516020908101518282015283515160408301528351015160608201526109f3610bd3565b600060408260808560066096fa905080600003610a525760405162461bcd60e51b815260206004820152601160248201527f61646467312063616c6c206661696c6564000000000000000000000000000000604482015260640161019a565b5080518351526020908101518351909101525092915050565b805151600080516020610dcb83398151915211610aca5760405162461bcd60e51b815260206004820152600c60248201527f78206e6f7420696e20465f500000000000000000000000000000000000000000604482015260640161019a565b805160200151600080516020610dcb83398151915211610b2c5760405162461bcd60e51b815260206004820152600c60248201527f79206e6f7420696e20465f500000000000000000000000000000000000000000604482015260640161019a565b805151600090600080516020610dcb8339815191529060039082908181800909088251602001519091508190600080516020610dcb83398151915290800914610bb75760405162461bcd60e51b815260206004820152601260248201527f706f696e74206e6f74206f6e2063757276650000000000000000000000000000604482015260640161019a565b5050565b6040518060200160405280610bce610bd3565b905290565b60405180604001604052806002906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b600060a08284031215610c2157600080fd5b60405160a0810181811067ffffffffffffffff82111715610c5257634e487b7160e01b600052604160045260246000fd5b806040525080915082358152602083013560208201526040830135604082015260608301356060820152608083013560808201525092915050565b60008060006101608486031215610ca357600080fd5b83359250610cb48560208601610c0f565b9150610cc38560c08601610c0f565b90509250925092565b815160408201908260005b6002811015610cf6578251825260209283019290910190600101610cd7565b50505092915050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610d4557610d45610d15565b500290565b634e487b7160e01b600052601260045260246000fd5b600082610d7d57634e487b7160e01b600052601260045260246000fd5b500690565b600060018201610d9457610d94610d15565b5060010190565b60008219821115610dae57610dae610d15565b500190565b600082821015610dc557610dc5610d15565b50039056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a164736f6c634300080d000a",
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

var VRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDKG\",\"name\":\"_keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"output\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structVRF.Request\",\"name\":\"r\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256[4]\",\"name\":\"p\",\"type\":\"uint256[4]\"}],\"internalType\":\"structECCArithmetic.G2Point\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f2\",\"type\":\"tuple\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"m\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f2\",\"type\":\"tuple\"}],\"name\":\"hashToCurve\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"hashPoint\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"s_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"transmitVRFResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"input\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256[4]\",\"name\":\"p\",\"type\":\"uint256[4]\"}],\"internalType\":\"structECCArithmetic.G2Point\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"denomInv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tInvSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y3\",\"type\":\"uint256\"}],\"internalType\":\"structHashToCurve.FProof\",\"name\":\"f2\",\"type\":\"tuple\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"vrfOutput\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"output\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620039aa380380620039aa833981016040819052620000349162000196565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000eb565b5050600280546001600160a01b0319166001600160a01b03949094169390931790925560035550620001d2565b336001600160a01b03821603620001455760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008060408385031215620001aa57600080fd5b82516001600160a01b0381168114620001c257600080fd5b6020939093015192949293505050565b6137c880620001e26000396000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80639201de55116100d8578063bf2732c71161008c578063e3d0e71211610066578063e3d0e71214610349578063e95d95771461035c578063f2fde38b1461037c57600080fd5b8063bf2732c714610324578063cc31f7dd14610337578063d57fc45a1461034057600080fd5b8063afcb95d7116100bd578063afcb95d7146102cf578063b1dc65a4146102f9578063bf0e15c61461031157600080fd5b80639201de55146102a9578063a954b4ef146102bc57600080fd5b806339614e4f1161012f57806379ba50971161011457806379ba50971461025957806381ff7048146102615780638da5cb5b1461028e57600080fd5b806339614e4f1461023e57806355e487491461025157600080fd5b80630bc643e8116101605780630bc643e8146101b75780630e3ca2a7146101dc578063181f5a77146101fc57600080fd5b8063012cfe861461017c578063092576f714610191575b600080fd5b61018f61018a366004612916565b61038f565b005b6101a461019f36600461297c565b610505565b6040519081526020015b60405180910390f35b6101ca6101c53660046129bd565b61064f565b60405160ff90911681526020016101ae565b6101a46101ea3660046129ed565b60056020526000908152604090205481565b60408051808201909152600981527f56524620302e302e31000000000000000000000000000000000000000000000060208201525b6040516101ae9190612a66565b61023161024c366004612bc2565b610678565b61018f6107fa565b61018f61085b565b600a54600854604080516000815264010000000090930463ffffffff1660208401528201526060016101ae565b6000546040516001600160a01b0390911681526020016101ae565b6102316102b7366004612bff565b61090c565b61018f6102ca366004612d7f565b610998565b6008546007546040805160008152602081019390935263ffffffff909116908201526060016101ae565b61018f610307366004612e3c565b5050505050505050565b6101a461031f366004612eef565b610c2c565b61018f610332366004612f38565b610dd5565b6101a460035481565b6101a460045481565b61018f6103573660046130ad565b610e5f565b61036f61036a36600461317a565b6115c0565b6040516101ae91906131b9565b61018f61038a3660046129ed565b61160e565b6004546000906103e65760405162461bcd60e51b815260206004820152601060248201527f6e6f206b657920617661696c61626c650000000000000000000000000000000060448201526064015b60405180910390fd5b60005a604080516020601f870181900481028201810190925285815291925086359181880135916104369184918491908a908a908190840183828082843760009201919091525061162292505050565b6040805183815263ffffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260085480825260095460ff808216602085015261010090910416928201929092529083146104f25760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016103dd565b63ffffffff8410610307576103076131f9565b6040805160808101825260008082526020808301828152838501838152336060808701828152918652600585528786205460035492518951921b6bffffffffffffffffffffffff191682870152603482018390526054808301829052895180840390910181526074909201909852805194019390932085529087905263ffffffff8616905290929061059683611671565b83516000908152600660209081526040918290209290925584519185015181860151606087015192517fda543d8fd5d52cb865899d85adee45422306c16f47e6e4394f043006ff5cde30946106139490939291938452602084019290925263ffffffff1660408301526001600160a01b0316606082015260800190565b60405180910390a1610626826001613225565b60608401516001600160a01b031660009081526005602052604090205550505190505b92915050565b6000600a8260ff1610156106685761064982603061323d565b61064982605761323d565b919050565b60606000808351600261068b9190613262565b67ffffffffffffffff8111156106a3576106a3612a79565b6040519080825280601f01601f1916602001820160405280156106cd576020820181803683370190505b509050600091505b80518260ff1610156107f3576000846106ef600285613297565b60ff1681518110610702576107026131e3565b60209101015160f81c600f1690506000600486610720600287613297565b60ff1681518110610733576107336131e3565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c905061076a8161064f565b60f81b838560ff1681518110610782576107826131e3565b60200101906001600160f81b031916908160001a9053506107a484600161323d565b93506107af8261064f565b60f81b838560ff16815181106107c7576107c76131e3565b60200101906001600160f81b031916908160001a905350505081806107eb906132b9565b9250506106d5565b9392505050565b6002546001600160a01b031633146108545760405162461bcd60e51b815260206004820181905260248201527f6b657920696e666f206d75737420636f6d652066726f6d2070726f766964657260448201526064016103dd565b6000600455565b6001546001600160a01b031633146108b55760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016103dd565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b602081101561098e5783816020811061094f5761094f6131e3565b1a60f81b828281518110610965576109656131e3565b60200101906001600160f81b031916908160001a90535080610986816132d8565b915050610934565b506107f381610678565b8151600090815260066020526040812054906109b384611671565b9050808214610a045760405162461bcd60e51b815260206004820152601560248201527f72657175657374206c6f6f6b7570206661696c6564000000000000000000000060448201526064016103dd565b6000610a108385610c2c565b85516000908152600660205260408082208290558701519192509063ffffffff1667ffffffffffffffff811115610a4957610a49612a79565b604051908082528060200260200182016040528015610a72578160200160208202803683370190505b50905060005b866040015163ffffffff16811015610ae65760408051602081018590529081018290526060016040516020818303038152906040528051906020012060001c828281518110610ac957610ac96131e3565b602090810291909101015280610ade816132d8565b915050610a78565b50606086015186516040517f75bf929b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916375bf929b91610b3491859060040161332c565b600060405180830381600087803b158015610b4e57600080fd5b505af1925050508015610b5f575060015b610bd8573d808015610b8d576040519150601f19603f3d011682016040523d82523d6000602084013e610b92565b606091505b5086516040517fa7231a311a37fec0b9b631c5b7d4c4aa3effe5304f25dbfeaf0de676cdd715ba90610bca9085906000908690613345565b60405180910390a250610c24565b8551604080516000815260208101918290527fa7231a311a37fec0b9b631c5b7d4c4aa3effe5304f25dbfeaf0de676cdd715ba91610c1b91859160019190613345565b60405180910390a25b505050505050565b60045481515160405160009291610c4591602001613372565b6040516020818303038152906040528051906020012014610ca85760405162461bcd60e51b815260206004820152601060248201527f77726f6e67207075626c6963206b65790000000000000000000000000000000060448201526064016103dd565b6000610cbd84846040015185606001516115c0565b9050610cd28184602001518560000151611707565b610d1e5760405162461bcd60e51b815260206004820152600d60248201527f626164205652462070726f6f660000000000000000000000000000000000000060448201526064016103dd565b6020830151515160008051602061379c833981519152118015610d56575060208381015151015160008051602061379c833981519152115b610da25760405162461bcd60e51b815260206004820152601f60248201527f62616420726570726573656e746174696f6e206f66206f75747075742070740060448201526064016103dd565b60208084015151604051610db692016133a6565b6040516020818303038152906040528051906020012091505092915050565b6002546001600160a01b03163314610e2f5760405162461bcd60e51b815260206004820181905260248201527f6b657920696e666f206d75737420636f6d652066726f6d2070726f766964657260448201526064016103dd565b8051604051610e4191906020016133da565b60408051601f19818403018152919052805160209091012060045550565b855185518560ff16601f831115610eb85760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016103dd565b60008111610f085760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016103dd565b818314610f7c5760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016103dd565b610f87816003613262565b8311610fd55760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016103dd565b610fdd6118f0565b60006040518060c001604052808b81526020018a81526020018960ff1681526020018881526020018767ffffffffffffffff1681526020018681525090505b600c541561112e57600c54600090611036906001906133f6565b90506000600c828154811061104d5761104d6131e3565b6000918252602082200154600d80546001600160a01b039092169350908490811061107a5761107a6131e3565b60009182526020808320909101546001600160a01b038581168452600b9092526040808420805461ffff1990811690915592909116808452922080549091169055600c805491925090806110d0576110d061340d565b600082815260209020810160001990810180546001600160a01b0319169055019055600d8054806111035761110361340d565b600082815260209020810160001990810180546001600160a01b03191690550190555061101c915050565b60005b81515181101561146d576000600b600084600001518481518110611157576111576131e3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561119457611194613423565b146111e15760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016103dd565b6040805180820190915260ff821681526001602082015282518051600b9160009185908110611212576112126131e3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561126b5761126b613423565b02179055506000915061127b9050565b600b600084602001518481518110611295576112956131e3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156112d2576112d2613423565b1461131f5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016103dd565b6040805180820190915260ff82168152602081016002815250600b600084602001518481518110611352576113526131e3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff1916176101008360028111156113ab576113ab613423565b02179055505082518051600c9250839081106113c9576113c96131e3565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b03909316929092179091558201518051600d919083908110611420576114206131e3565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b0390921691909117905580611465816132d8565b915050611131565b5060408101516009805460ff191660ff909216919091179055600a805467ffffffff0000000019811664010000000063ffffffff4381168202928317855590830481169360019390926000926114ca928692908216911617613439565b92506101000a81548163ffffffff021916908363ffffffff1602179055506115294630600a60009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161194c565b6008819055825180516009805460ff9092166101000261ff0019909216919091179055600a5460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986115ab988b98919763ffffffff90921696909591949193919261349a565b60405180910390a15050505050505050505050565b6115c8612837565b60006115d3856119d9565b905060006115e882825b602002015186611add565b905060006115f78360016115dd565b90506116038282612255565b979650505050505050565b6116166118f0565b61161f816122cd565b50565b6000808280602001905181019061163991906135d4565b915091506116478282610998565b50506007805460089390931c63ffffffff1663ffffffff19909316929092179091555050565b5050565b60008082600001518360200151846040015185606001516040516020016116e89493929190938452602084019290925260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016604083015260601b6bffffffffffffffffffffffff1916604482015260580190565b60408051601f1981840301815291905280516020909101209392505050565b60408051600280825260608201909252600091829190816020015b61172a612837565b8152602001906001900390816117225750506040805160028082526060820190925291925060009190602082015b61176061284f565b8152602001906001900390816117585750506040805160608101825288515160208083019182528a510151939450909283928301906117ad9060008051602061379c8339815191526133f6565b815250815250826000815181106117c6576117c66131e3565b602002602001018190525084826001815181106117e5576117e56131e3565b60200260200101819052508381600081518110611804576118046131e3565b6020026020010181905250604051806020016040528060405180608001604052807f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281526020017f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed81526020017f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b81526020017f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa815250815250816001815181106118d1576118d16131e3565b60200260200101819052506118e68282612376565b9695505050505050565b6000546001600160a01b0316331461194a5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016103dd565b565b6000808a8a8a8a8a8a8a8a8a604051602001611970999897969594939291906136ff565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6119e1612862565b6000805b6002821015611ad65760408051602081018690520160408051601f198184030181529190528051602090910120935083611a2e60008051602061379c8339815191526005613262565b811015611a7257611a4d60008051602061379c83398151915282613787565b848460028110611a5f57611a5f6131e3565b602002015282611a6e816132d8565b9350505b81611a7c816132d8565b9250506020821115611ad05760405162461bcd60e51b815260206004820152601960248201527f617474656d7074656420746f6f206d616e79206861736865730000000000000060448201526064016103dd565b506119e5565b5050919050565b611ae5612837565b82600003611b295750604080516060810182527759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe60208201908152600292820192909252908152610649565b611b4260008051602061379c8339815191526005613262565b8310611b905760405162461bcd60e51b815260206004820152601660248201527f74206e6f74206120756e69666f726d2073616d706c650000000000000000000060448201526064016103dd565b600060008051602061379c8339815191528485099050600081611bb560036001613225565b611bbf9190613225565b905060008051602061379c83398151915284518209600114611c235760405162461bcd60e51b815260206004820152601760248201527f77726f6e6720696e766572736520666f722064656e6f6d00000000000000000060448201526064016103dd565b600060008051602061379c833981519152855160008051602061379c8339815191528577b3c4d79d41a91759a9e4c7e359b6b89eaec68e62effffffd09099050600060008051602061379c833981519152611c8c8360008051602061379c8339815191526133f6565b7759e26bcea0d48bacd4f263f1acdb5c4f5763473177fffffe0890506000611cb5600289613787565b905081600060008051602061379c833981519152600360008051602061379c833981519152808586098509089050600060008051602061379c83398151915260408b015180099050818103611de65760008051602061379c8339815191528a6040015110611d655760405162461bcd60e51b815260206004820152600c60248201527f793120746f6f206c61726765000000000000000000000000000000000000000060448201526064016103dd565b8360028b60400151611d779190613787565b14611dc45760405162461bcd60e51b815260206004820152601860248201527f793120706172697479206d757374206d6174636820742773000000000000000060448201526064016103dd565b885183905260408a0151895160015b6020020152506106499650505050505050565b611dfe8260008051602061379c8339815191526133f6565b8114611e4c5760405162461bcd60e51b815260206004820152601860248201527f7931213d70736575646f20737172206f662078315e332b42000000000000000060448201526064016103dd565b505050600060008051602061379c83398151915280611e6d57611e6d613281565b83600108611e899060008051602061379c8339815191526133f6565b9050600060008051602061379c833981519152600360008051602061379c833981519152808586098509089050600060008051602061379c83398151915260608b015180099050818103611f995760008051602061379c8339815191528a6060015110611f275760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b60448201526064016103dd565b8360028b60600151611f399190613787565b14611f865760405162461bcd60e51b815260206004820152601860248201527f793220706172697479206d757374206d6174636820742773000000000000000060448201526064016103dd565b885183905260608a015189516001611dd3565b611fb18260008051602061379c8339815191526133f6565b8114611fff5760405162461bcd60e51b815260206004820152601860248201527f7932213d70736575646f20737172206f662078325e332b42000000000000000060448201526064016103dd565b50505060008051602061379c8339815191528061201e5761201e613281565b876020015186096001146120745760405162461bcd60e51b815260206004820152601c60248201527f74496e76537175617265642a742a2a3220213d3d2031206d6f6420500000000060448201526064016103dd565b60008051602061379c8339815191527f2042def740cbc01bd03583cf0100e593ba56470b9af68708d2c05d649053538560008051602061379c83398151915260208a015160008051602061379c83398151915288890909099250600060008051602061379c8339815191526120f78560008051602061379c8339815191526133f6565b6001089050600060008051602061379c833981519152600360008051602061379c8339815191528085860985090890508060008051602061379c83398151915260808b015180091461218b5760405162461bcd60e51b815260206004820152601c60248201527f646964206e6f74206f627461696e206120637572766520706f696e740000000060448201526064016103dd565b60008051602061379c8339815191528960800151106121db5760405162461bcd60e51b815260206004820152600c60248201526b793220746f6f206c6172676560a01b60448201526064016103dd565b8260028a608001516121ed9190613787565b1461223a5760405162461bcd60e51b815260206004820152601860248201527f793320706172697479206d757374206d6174636820742773000000000000000060448201526064016103dd565b50865152505050506080929092015181516020015292915050565b61225d612837565b60006122698484612627565b80515190915015801590612281575080516020015115155b6107f35760405162461bcd60e51b815260206004820152601b60248201527f6164646731206661696c65643a207a65726f206f7264696e617465000000000060448201526064016103dd565b336001600160a01b038216036123255760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016103dd565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000815183511461238657600080fd5b82516000612395826006613262565b905060008167ffffffffffffffff8111156123b2576123b2612a79565b6040519080825280602002602001820160405280156123db578160200160208202803683370190505b50905060005b838110156125f0578681815181106123fb576123fb6131e3565b6020908102919091010151515182612414836006613262565b61241f906000613225565b8151811061242f5761242f6131e3565b60200260200101818152505086818151811061244d5761244d6131e3565b60209081029190910181015151015182612468836006613262565b612473906001613225565b81518110612483576124836131e3565b6020026020010181815250508581815181106124a1576124a16131e3565b60209081029190910101515151826124ba836006613262565b6124c5906002613225565b815181106124d5576124d56131e3565b6020026020010181815250508581815181106124f3576124f36131e3565b6020908102919091018101515101518261250e836006613262565b612519906003613225565b81518110612529576125296131e3565b602002602001018181525050858181518110612547576125476131e3565b6020908102919091010151516040015182612563836006613262565b61256e906004613225565b8151811061257e5761257e6131e3565b60200260200101818152505085818151811061259c5761259c6131e3565b60209081029190910101515160600151826125b8836006613262565b6125c3906005613225565b815181106125d3576125d36131e3565b6020908102919091010152806125e8816132d8565b9150506123e1565b506125f9612880565b6000602082602086026020860160086201b968fa90508061261957600080fd5b505115159695505050505050565b61262f612837565b612638836126eb565b612641826126eb565b61264961289e565b83515181528351602090810151828201528351516040830152835101516060820152612673612862565b600060408260808560066096fa9050806000036126d25760405162461bcd60e51b815260206004820152601160248201527f61646467312063616c6c206661696c656400000000000000000000000000000060448201526064016103dd565b5080518351526020908101518351909101525092915050565b80515160008051602061379c8339815191521161274a5760405162461bcd60e51b815260206004820152600c60248201527f78206e6f7420696e20465f50000000000000000000000000000000000000000060448201526064016103dd565b80516020015160008051602061379c833981519152116127ac5760405162461bcd60e51b815260206004820152600c60248201527f79206e6f7420696e20465f50000000000000000000000000000000000000000060448201526064016103dd565b80515160009060008051602061379c833981519152906003908290818180090908825160200151909150819060008051602061379c8339815191529080091461166d5760405162461bcd60e51b815260206004820152601260248201527f706f696e74206e6f74206f6e206375727665000000000000000000000000000060448201526064016103dd565b604051806020016040528061284a612862565b905290565b604051806020016040528061284a61289e565b60405180604001604052806002906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b806060810183101561064957600080fd5b60008083601f8401126128df57600080fd5b50813567ffffffffffffffff8111156128f757600080fd5b60208301915083602082850101111561290f57600080fd5b9250929050565b60008060006080848603121561292b57600080fd5b61293585856128bc565b9250606084013567ffffffffffffffff81111561295157600080fd5b61295d868287016128cd565b9497909650939450505050565b63ffffffff8116811461161f57600080fd5b6000806040838503121561298f57600080fd5b8235915060208301356129a18161296a565b809150509250929050565b803560ff8116811461067357600080fd5b6000602082840312156129cf57600080fd5b6107f3826129ac565b6001600160a01b038116811461161f57600080fd5b6000602082840312156129ff57600080fd5b81356107f3816129d8565b60005b83811015612a25578181015183820152602001612a0d565b83811115612a34576000848401525b50505050565b60008151808452612a52816020860160208601612a0a565b601f01601f19169290920160200192915050565b6020815260006107f36020830184612a3a565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715612ab257612ab2612a79565b60405290565b6040516080810167ffffffffffffffff81118282101715612ab257612ab2612a79565b6040516020810167ffffffffffffffff81118282101715612ab257612ab2612a79565b6040805190810167ffffffffffffffff81118282101715612ab257612ab2612a79565b604051601f8201601f1916810167ffffffffffffffff81118282101715612b4a57612b4a612a79565b604052919050565b600082601f830112612b6357600080fd5b813567ffffffffffffffff811115612b7d57612b7d612a79565b612b90601f8201601f1916602001612b21565b818152846020838601011115612ba557600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215612bd457600080fd5b813567ffffffffffffffff811115612beb57600080fd5b612bf784828501612b52565b949350505050565b600060208284031215612c1157600080fd5b5035919050565b6000612c22612ab8565b9050806080830184811115612c3657600080fd5b835b81811015612c50578035835260209283019201612c38565b50505092915050565b6000612c63612afe565b9050806040830184811115612c3657600080fd5b600060a08284031215612c8957600080fd5b612c91612a8f565b9050813581526020820135602082015260408201356040820152606082013560608201526080820135608082015292915050565b6000818303610200811215612cd957600080fd5b612ce1612ab8565b91506080811215612cf157600080fd5b612cf9612adb565b84601f850112612d0857600080fd5b612d128585612c18565b815282526040607f1982011215612d2857600080fd5b50612d31612adb565b83609f840112612d4057600080fd5b612d4d8460808501612c59565b81526020820152612d618360c08401612c77565b6040820152612d74836101608401612c77565b606082015292915050565b600080828403610280811215612d9457600080fd5b6080811215612da257600080fd5b50612dab612ab8565b83358152602084013560208201526040840135612dc78161296a565b60408201526060840135612dda816129d8565b60608201529150612dee8460808501612cc5565b90509250929050565b60008083601f840112612e0957600080fd5b50813567ffffffffffffffff811115612e2157600080fd5b6020830191508360208260051b850101111561290f57600080fd5b60008060008060008060008060e0898b031215612e5857600080fd5b612e628a8a6128bc565b9750606089013567ffffffffffffffff80821115612e7f57600080fd5b612e8b8c838d016128cd565b909950975060808b0135915080821115612ea457600080fd5b612eb08c838d01612df7565b909750955060a08b0135915080821115612ec957600080fd5b50612ed68b828c01612df7565b999c989b50969995989497949560c00135949350505050565b6000806102208385031215612f0357600080fd5b82359150612dee8460208501612cc5565b600067ffffffffffffffff821115612f2e57612f2e612a79565b5060051b60200190565b60006020808385031215612f4b57600080fd5b823567ffffffffffffffff80821115612f6357600080fd5b9084019060408287031215612f7757600080fd5b612f7f612afe565b823582811115612f8e57600080fd5b612f9a88828601612b52565b8252508383013582811115612fae57600080fd5b80840193505086601f840112612fc357600080fd5b82359150612fd8612fd383612f14565b612b21565b82815260059290921b83018401918481019088841115612ff757600080fd5b938501935b8385101561301557843582529385019390850190612ffc565b948201949094529695505050505050565b600082601f83011261303757600080fd5b81356020613047612fd383612f14565b82815260059290921b8401810191818101908684111561306657600080fd5b8286015b8481101561308a57803561307d816129d8565b835291830191830161306a565b509695505050505050565b803567ffffffffffffffff8116811461067357600080fd5b60008060008060008060c087890312156130c657600080fd5b863567ffffffffffffffff808211156130de57600080fd5b6130ea8a838b01613026565b9750602089013591508082111561310057600080fd5b61310c8a838b01613026565b965061311a60408a016129ac565b9550606089013591508082111561313057600080fd5b61313c8a838b01612b52565b945061314a60808a01613095565b935060a089013591508082111561316057600080fd5b5061316d89828a01612b52565b9150509295509295509295565b6000806000610160848603121561319057600080fd5b833592506131a18560208601612c77565b91506131b08560c08601612c77565b90509250925092565b815160408201908260005b6002811015612c505782518252602092830192909101906001016131c4565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156132385761323861320f565b500190565b600060ff821660ff84168060ff0382111561325a5761325a61320f565b019392505050565b600081600019048311821515161561327c5761327c61320f565b500290565b634e487b7160e01b600052601260045260246000fd5b600060ff8316806132aa576132aa613281565b8060ff84160491505092915050565b600060ff821660ff81036132cf576132cf61320f565b60010192915050565b6000600182016132ea576132ea61320f565b5060010190565b600081518084526020808501945080840160005b8381101561332157815187529582019590820190600101613305565b509495945050505050565b828152604060208201526000612bf760408301846132f1565b60608152600061335860608301866132f1565b841515602084015282810360408401526118e68185612a3a565b60008183825b6004811015613397578151835260209283019290910190600101613378565b50505060808201905092915050565b60008183825b60028110156133cb5781518352602092830192909101906001016133ac565b50505060408201905092915050565b600082516133ec818460208701612a0a565b9190910192915050565b6000828210156134085761340861320f565b500390565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600063ffffffff8083168185168083038211156134585761345861320f565b01949350505050565b600081518084526020808501945080840160005b838110156133215781516001600160a01b031687529582019590820190600101613475565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526134ca8184018a613461565b905082810360808401526134de8189613461565b905060ff871660a084015282810360c08401526134fb8187612a3a565b905067ffffffffffffffff851660e08401528281036101008401526135208185612a3a565b9c9b505050505050505050505050565b600061353a612ab8565b905080608083018481111561354e57600080fd5b835b81811015612c50578051835260209283019201613550565b6000613572612afe565b905080604083018481111561354e57600080fd5b600060a0828403121561359857600080fd5b6135a0612a8f565b9050815181526020820151602082015260408201516040820152606082015160608201526080820151608082015292915050565b6000808284036102808112156135e957600080fd5b60808112156135f757600080fd5b6135ff612ab8565b8451815260208501516020820152604085015161361b8161296a565b6040820152606085015161362e816129d8565b60608201529250607f19810161020081121561364957600080fd5b613651612ab8565b608082121561365f57600080fd5b613667612adb565b915086609f87011261367857600080fd5b6136858760808801613530565b8252818152604060ff198401121561369c57600080fd5b6136a4612adb565b92508661011f8701126136b657600080fd5b6136c4876101008801613568565b83528260208201526136da876101408801613586565b60408201526136ed876101e08801613586565b60608201528093505050509250929050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526137398285018b613461565b9150838203608085015261374d828a613461565b915060ff881660a085015283820360c085015261376a8288612a3a565b90861660e085015283810361010085015290506135208185612a3a565b60008261379657613796613281565b50069056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a164736f6c634300080d000a",
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

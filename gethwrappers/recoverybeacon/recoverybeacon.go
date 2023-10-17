package recoverybeacon

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

type RecoveryBeaconReportReport struct {
	AccountToRecover common.Address
	Recoverer        common.Address
	Success          bool
}

type RecoveryBeaconTypesEnrollmentRequest struct {
	AddressToEnroll common.Address
	PublicKeyBytes  []byte
}

type RecoveryBeaconTypesEnrollmentResponse struct {
	PlayerIdx                uint8
	Threshold                uint8
	Cipher                   []byte
	EphermeralKeyPubKeyBytes []byte
	Nonce                    []byte
	DistributedPublicKey     []byte
	AccountPointBytes        []byte
	RecoveryPubKeyBytes      []byte
}

type RecoveryBeaconTypesLease struct {
	Recoverer        common.Address
	AccountToRecover common.Address
	AttemptsLeft     uint8
	BlockPurchased   *big.Int
}

type RecoveryBeaconTypesRecoveryRequest struct {
	PlayerIdx                uint8
	Nonce                    []byte
	DistributedPublicKey     []byte
	Cipher                   []byte
	EphermeralKeyPubKeyBytes []byte
	RecoveryPubKeyBytes      []byte
	Recoverer                common.Address
	AddressToRecover         common.Address
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

var DKGMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractDKGClient\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"DKGClientError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"indexed\":false,\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"KeyGenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"addClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addressToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_configDigest\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"removeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000a5565b50505062000150565b336001600160a01b03821603620000ff5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61299380620001606000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806381ff704811610097578063b1dc65a411610066578063b1dc65a414610241578063c3105a6b14610254578063e3d0e71214610274578063f2fde38b1461028757600080fd5b806381ff7048146101bc5780638da5cb5b146101e95780639201de5514610204578063afcb95d71461021757600080fd5b80635429a79e116100d35780635429a79e146101795780635e57966d1461018e57806379ba5097146101a15780637bf1ffc5146101a957600080fd5b80630bc643e8146100fa578063181f5a771461012457806339614e4f14610166575b600080fd5b61010d610108366004611ed4565b61029a565b60405160ff90911681526020015b60405180910390f35b60408051808201909152600981527f444b4720302e302e31000000000000000000000000000000000000000000000060208201525b60405161011b9190611f4b565b610159610174366004612023565b6102c9565b61018c61018736600461206d565b61044b565b005b61015961019c36600461209d565b61068c565b61018c610753565b61018c6101b736600461206d565b610809565b600754600554604080516000815264010000000090930463ffffffff16602084015282015260600161011b565b6000546040516001600160a01b03909116815260200161011b565b6101596102123660046120ba565b610850565b6005546004546040805160008152602081019390935263ffffffff9091169082015260600161011b565b61018c61024f36600461211f565b6108dc565b610267610262366004612204565b610a28565b60405161011b9190612226565b61018c61028236600461232e565b610b50565b61018c61029536600461209d565b6112cc565b6000600a8260ff1610156102b9576102b3826030612411565b92915050565b6102b3826057612411565b919050565b6060600080835160026102dc9190612436565b67ffffffffffffffff8111156102f4576102f4611f5e565b6040519080825280601f01601f19166020018201604052801561031e576020820181803683370190505b509050600091505b80518260ff16101561044457600084610340600285612455565b60ff168151811061035357610353612485565b60209101015160f81c600f1690506000600486610371600287612455565b60ff168151811061038457610384612485565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c90506103bb8161029a565b60f81b838560ff16815181106103d3576103d3612485565b60200101906001600160f81b031916908160001a9053506103f5846001612411565b93506104008261029a565b60f81b838560ff168151811061041857610418612485565b60200101906001600160f81b031916908160001a9053505050818061043c9061249b565b925050610326565b9392505050565b6104536112e0565b6000828152600260209081526040808320805482518185028101850190935280835291929091908301828280156104b357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610495575b505050505090506000815167ffffffffffffffff8111156104d6576104d6611f5e565b6040519080825280602002602001820160405280156104ff578160200160208202803683370190505b5090506000805b83518110156105a257846001600160a01b031684828151811061052b5761052b612485565b60200260200101516001600160a01b03161461058257848361054d84846124ba565b8151811061055d5761055d612485565b60200260200101906001600160a01b031690816001600160a01b031681525050610590565b8161058c816124d1565b9250505b8061059a816124d1565b915050610506565b5060008184516105b291906124ba565b67ffffffffffffffff8111156105ca576105ca611f5e565b6040519080825280602002602001820160405280156105f3578160200160208202803683370190505b50905060005b82855161060691906124ba565b8110156106635783818151811061061f5761061f612485565b602002602001015182828151811061063957610639612485565b6001600160a01b03909216602092830291909101909101528061065b816124d1565b9150506105f9565b506000868152600260209081526040909120825161068392840190611def565b50505050505050565b604080516014808252818301909252606091600091906020820181803683370190505090508260005b60148160ff161015610741577fff0000000000000000000000000000000000000000000000000000000000000060f883901b16836106f48360136124ea565b60ff168151811061070757610707612485565b60200101906001600160f81b031916908160001a9053506008826001600160a01b0316901c915080806107399061249b565b9150506106b5565b5061074b826102c9565b949350505050565b6001546001600160a01b031633146107b25760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6108116112e0565b600091825260026020908152604083208054600181018255908452922090910180546001600160a01b0319166001600160a01b03909216919091179055565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156108d25783816020811061089357610893612485565b1a60f81b8282815181106108a9576108a9612485565b60200101906001600160f81b031916908160001a905350806108ca816124d1565b915050610878565b50610444816102c9565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161092c9184918491908e908e908190840183828082843760009201919091525061133c92505050565b6040805183815263ffffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260055480825260065460ff808216602085015261010090910416928201929092529083146109e85760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016107a9565b6109f68b8b8b8b8b8b611596565b610a078c8c8c8c8c8c8c8c8961162a565b50505063ffffffff8110610a1d57610a1d61250d565b505050505050505050565b60408051808201909152606080825260208201526000838152600360209081526040808320858452909152908190208151808301909252805482908290610a6e90612523565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9a90612523565b8015610ae75780601f10610abc57610100808354040283529160200191610ae7565b820191906000526020600020905b815481529060010190602001808311610aca57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610b3f57602002820191906000526020600020905b815481526020019060010190808311610b2b575b505050505081525050905092915050565b855185518560ff16601f831115610ba95760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016107a9565b60008111610bf95760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016107a9565b818314610c6d5760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016107a9565b610c78816003612436565b8311610cc65760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016107a9565b610cce6112e0565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60095415610e1e57600954600090610d26906001906124ba565b9050600060098281548110610d3d57610d3d612485565b6000918252602082200154600a80546001600160a01b0390921693509084908110610d6a57610d6a612485565b60009182526020808320909101546001600160a01b03858116845260089092526040808420805461ffff1990811690915592909116808452922080549091169055600980549192509080610dc057610dc061255d565b600082815260209020810160001990810180546001600160a01b0319169055019055600a805480610df357610df361255d565b600082815260209020810160001990810180546001600160a01b031916905501905550610d0c915050565b60005b81515181101561115d5760006008600084600001518481518110610e4757610e47612485565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610e8457610e84612573565b14610ed15760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016107a9565b6040805180820190915260ff82168152600160208201528251805160089160009185908110610f0257610f02612485565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610f5b57610f5b612573565b021790555060009150610f6b9050565b6008600084602001518481518110610f8557610f85612485565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610fc257610fc2612573565b1461100f5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016107a9565b6040805180820190915260ff82168152602081016002815250600860008460200151848151811061104257611042612485565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561109b5761109b612573565b0217905550508251805160099250839081106110b9576110b9612485565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b03909316929092179091558201518051600a91908390811061111057611110612485565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b0390921691909117905580611155816124d1565b915050610e21565b5060408101516006805460ff191660ff9092169190911790556007805467ffffffff0000000019811664010000000063ffffffff4381168202928317855590830481169360019390926000926111ba928692908216911617612589565b92506101000a81548163ffffffff021916908363ffffffff160217905550600061121b4630600760009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a00151611ab9565b6005819055835180516006805460ff9092166101000261ff00199092169190911790556007546020860151604080880151606089015160808a015160a08b015193519798507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05976112a3978b978b9763ffffffff9091169691959094909390929091906125f5565b60405180910390a16112be8360400151846060015183611b46565b505050505050505050505050565b6112d46112e0565b6112dd81611d46565b50565b6000546001600160a01b0316331461133a5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107a9565b565b600060608083806020019051810190611355919061268b565b60408051808201825283815260208082018490526000868152600282528381208054855181850281018501909652808652979a50959850939650909492939192908301828280156113cf57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113b1575b5050505050905060005b81518110156114ec578181815181106113f4576113f4612485565b60200260200101516001600160a01b031663bf2732c7846040518263ffffffff1660e01b81526004016114279190612226565b600060405180830381600087803b15801561144157600080fd5b505af1925050508015611452575060015b6114da573d808015611480576040519150601f19603f3d011682016040523d82523d6000602084013e611485565b606091505b507f116391732f5df106193bda7cedf1728f3b07b62f6cdcdd611c9eeec44efcae548383815181106114b9576114b9612485565b6020026020010151826040516114d0929190612789565b60405180910390a1505b806114e4816124d1565b9150506113d9565b5060008581526003602090815260408083208b845290915290208251839190819061151790826127fa565b5060208281015180516115309260018501920190611e54565b5090505084887fc8db841f5b2231ccf7190311f440aa197b161e369f3b40b023508160cc555656846040516115659190612226565b60405180910390a350506004805460089690961c63ffffffff1663ffffffff19909616959095179094555050505050565b60006115a3826020612436565b6115ae856020612436565b6115ba886101446128ba565b6115c491906128ba565b6115ce91906128ba565b6115d99060006128ba565b90503681146106835760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016107a9565b60006002826020015183604001516116429190612411565b61164c9190612455565b611657906001612411565b60408051600180825281830190925260ff929092169250600091906020820181803683370190505090508160f81b8160008151811061169857611698612485565b60200101906001600160f81b031916908160001a9053508682146116bb826102c9565b906116d95760405162461bcd60e51b81526004016107a99190611f4b565b508685146117295760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016107a9565b3360009081526008602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561176c5761176c612573565b600281111561177d5761177d612573565b905250905060028160200151600281111561179a5761179a612573565b1480156117d45750600a816000015160ff16815481106117bc576117bc612485565b6000918252602090912001546001600160a01b031633145b6118205760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016107a9565b505050600088886040516118359291906128d2565b60405190819003812061184c918c906020016128e2565b60405160208183030381529060405280519060200120905061186c611e8f565b604080518082019091526000808252602082015260005b88811015611aaa5760006001858884602081106118a2576118a2612485565b6118af91901a601b612411565b8d8d868181106118c1576118c1612485565b905060200201358c8c878181106118da576118da612485565b9050602002013560405160008152602001604052604051611917949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611939573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526008602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561198e5761198e612573565b600281111561199f5761199f612573565b90525092506001836020015160028111156119bc576119bc612573565b14611a095760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016107a9565b8251849060ff16601f8110611a2057611a20612485565b602002015115611a725760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016107a9565b600184846000015160ff16601f8110611a8d57611a8d612485565b911515602090920201525080611aa2816124d1565b915050611883565b50505050505050505050505050565b6000808a8a8a8a8a8a8a8a8a604051602001611add999897969594939291906128fe565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6000808351602014611b9a5760405162461bcd60e51b815260206004820152601e60248201527f77726f6e67206c656e67746820666f72206f6e636861696e436f6e666967000060448201526064016107a9565b60208401519150808203611bf05760405162461bcd60e51b815260206004820152601460248201527f6661696c656420746f20636f7079206b6579494400000000000000000000000060448201526064016107a9565b60408051808201909152606080825260208201526000838152600360209081526040808320878452909152902081518291908190611c2e90826127fa565b506020828101518051611c479260018501920190611e54565b505050600083815260026020908152604080832080548251818502810185019093528083529192909190830182828015611caa57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611c8c575b5050505050905060005b8151811015611d3c57818181518110611ccf57611ccf612485565b60200260200101516001600160a01b03166355e487496040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611d1157600080fd5b505af1158015611d25573d6000803e3d6000fd5b505050508080611d34906124d1565b915050611cb4565b5050505050505050565b336001600160a01b03821603611d9e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107a9565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215611e44579160200282015b82811115611e4457825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611e0f565b50611e50929150611eae565b5090565b828054828255906000526020600020908101928215611e44579160200282015b82811115611e44578251825591602001919060010190611e74565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115611e505760008155600101611eaf565b803560ff811681146102c457600080fd5b600060208284031215611ee657600080fd5b61044482611ec3565b60005b83811015611f0a578181015183820152602001611ef2565b83811115611f19576000848401525b50505050565b60008151808452611f37816020860160208601611eef565b601f01601f19169290920160200192915050565b6020815260006104446020830184611f1f565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611f9d57611f9d611f5e565b604052919050565b600067ffffffffffffffff821115611fbf57611fbf611f5e565b50601f01601f191660200190565b600082601f830112611fde57600080fd5b8135611ff1611fec82611fa5565b611f74565b81815284602083860101111561200657600080fd5b816020850160208301376000918101602001919091529392505050565b60006020828403121561203557600080fd5b813567ffffffffffffffff81111561204c57600080fd5b61074b84828501611fcd565b6001600160a01b03811681146112dd57600080fd5b6000806040838503121561208057600080fd5b82359150602083013561209281612058565b809150509250929050565b6000602082840312156120af57600080fd5b813561044481612058565b6000602082840312156120cc57600080fd5b5035919050565b60008083601f8401126120e557600080fd5b50813567ffffffffffffffff8111156120fd57600080fd5b6020830191508360208260051b850101111561211857600080fd5b9250929050565b60008060008060008060008060e0898b03121561213b57600080fd5b606089018a81111561214c57600080fd5b8998503567ffffffffffffffff8082111561216657600080fd5b818b0191508b601f83011261217a57600080fd5b81358181111561218957600080fd5b8c602082850101111561219b57600080fd5b6020830199508098505060808b01359150808211156121b957600080fd5b6121c58c838d016120d3565b909750955060a08b01359150808211156121de57600080fd5b506121eb8b828c016120d3565b999c989b50969995989497949560c00135949350505050565b6000806040838503121561221757600080fd5b50508035926020909101359150565b6000602080835283516040828501526122426060850182611f1f565b85830151858203601f19016040870152805180835290840192506000918401905b808310156122835783518252928401926001929092019190840190612263565b509695505050505050565b600067ffffffffffffffff8211156122a8576122a8611f5e565b5060051b60200190565b600082601f8301126122c357600080fd5b813560206122d3611fec8361228e565b82815260059290921b840181019181810190868411156122f257600080fd5b8286015b8481101561228357803561230981612058565b83529183019183016122f6565b803567ffffffffffffffff811681146102c457600080fd5b60008060008060008060c0878903121561234757600080fd5b863567ffffffffffffffff8082111561235f57600080fd5b61236b8a838b016122b2565b9750602089013591508082111561238157600080fd5b61238d8a838b016122b2565b965061239b60408a01611ec3565b955060608901359150808211156123b157600080fd5b6123bd8a838b01611fcd565b94506123cb60808a01612316565b935060a08901359150808211156123e157600080fd5b506123ee89828a01611fcd565b9150509295509295509295565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff0382111561242e5761242e6123fb565b019392505050565b6000816000190483118215151615612450576124506123fb565b500290565b600060ff83168061247657634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036124b1576124b16123fb565b60010192915050565b6000828210156124cc576124cc6123fb565b500390565b6000600182016124e3576124e36123fb565b5060010190565b600060ff821660ff841680821015612504576125046123fb565b90039392505050565b634e487b7160e01b600052600160045260246000fd5b600181811c9082168061253757607f821691505b60208210810361255757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600063ffffffff8083168185168083038211156125a8576125a86123fb565b01949350505050565b600081518084526020808501945080840160005b838110156125ea5781516001600160a01b0316875295820195908201906001016125c5565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526126258184018a6125b1565b9050828103608084015261263981896125b1565b905060ff871660a084015282810360c08401526126568187611f1f565b905067ffffffffffffffff851660e084015282810361010084015261267b8185611f1f565b9c9b505050505050505050505050565b6000806000606084860312156126a057600080fd5b8351925060208085015167ffffffffffffffff808211156126c057600080fd5b818701915087601f8301126126d457600080fd5b81516126e2611fec82611fa5565b81815289858386010111156126f657600080fd5b61270582868301878701611eef565b60408901519096509250508082111561271d57600080fd5b508501601f8101871361272f57600080fd5b805161273d611fec8261228e565b81815260059190911b8201830190838101908983111561275c57600080fd5b928401925b8284101561277a57835182529284019290840190612761565b80955050505050509250925092565b6001600160a01b038316815260406020820152600061074b6040830184611f1f565b601f8211156127f557600081815260208120601f850160051c810160208610156127d25750805b601f850160051c820191505b818110156127f1578281556001016127de565b5050505b505050565b815167ffffffffffffffff81111561281457612814611f5e565b612828816128228454612523565b846127ab565b602080601f83116001811461285d57600084156128455750858301515b600019600386901b1c1916600185901b1785556127f1565b600085815260208120601f198616915b8281101561288c5788860151825594840194600190910190840161286d565b50858210156128aa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600082198211156128cd576128cd6123fb565b500190565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526129388285018b6125b1565b9150838203608085015261294c828a6125b1565b915060ff881660a085015283820360c08501526129698288611f1f565b90861660e0850152838103610100850152905061267b8185611f1f56fea164736f6c634300080f000a",
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

func (_DKG *DKGCaller) AddressToString(opts *bind.CallOpts, a common.Address) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "addressToString", a)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_DKG *DKGSession) AddressToString(a common.Address) (string, error) {
	return _DKG.Contract.AddressToString(&_DKG.CallOpts, a)
}

func (_DKG *DKGCallerSession) AddressToString(a common.Address) (string, error) {
	return _DKG.Contract.AddressToString(&_DKG.CallOpts, a)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addressToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610663806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630bc643e81461005157806339614e4f1461007b5780635e57966d1461009b5780639201de55146100ae575b600080fd5b61006461005f3660046103cd565b6100c1565b60405160ff90911681526020015b60405180910390f35b61008e610089366004610406565b6100eb565b60405161007291906104b7565b61008e6100a936600461050c565b61026d565b61008e6100bc366004610542565b610341565b6000600a8260ff1610156100e0576100da826030610571565b92915050565b6100da826057610571565b6060600080835160026100fe9190610596565b67ffffffffffffffff811115610116576101166103f0565b6040519080825280601f01601f191660200182016040528015610140576020820181803683370190505b509050600091505b80518260ff161015610266576000846101626002856105b5565b60ff1681518110610175576101756105e5565b60209101015160f81c600f16905060006004866101936002876105b5565b60ff16815181106101a6576101a66105e5565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c90506101dd816100c1565b60f81b838560ff16815181106101f5576101f56105e5565b60200101906001600160f81b031916908160001a905350610217846001610571565b9350610222826100c1565b60f81b838560ff168151811061023a5761023a6105e5565b60200101906001600160f81b031916908160001a9053505050818061025e906105fb565b925050610148565b9392505050565b604080516014808252818301909252606091600091906020820181803683370190505090508260005b60148160ff16101561032f577fff0000000000000000000000000000000000000000000000000000000000000060f883901b16836102d583601361061a565b60ff16815181106102e8576102e86105e5565b60200101906001600160f81b031916908160001a90535060088273ffffffffffffffffffffffffffffffffffffffff16901c91508080610327906105fb565b915050610296565b50610339826100eb565b949350505050565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156103c357838160208110610384576103846105e5565b1a60f81b82828151811061039a5761039a6105e5565b60200101906001600160f81b031916908160001a905350806103bb8161063d565b915050610369565b50610266816100eb565b6000602082840312156103df57600080fd5b813560ff8116811461026657600080fd5b634e487b7160e01b600052604160045260246000fd5b60006020828403121561041857600080fd5b813567ffffffffffffffff8082111561043057600080fd5b818401915084601f83011261044457600080fd5b813581811115610456576104566103f0565b604051601f8201601f19908116603f0116810190838211818310171561047e5761047e6103f0565b8160405282815287602084870101111561049757600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156104e4578581018301518582016040015282016104c8565b818111156104f6576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561051e57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461026657600080fd5b60006020828403121561055457600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff0382111561058e5761058e61055b565b019392505050565b60008160001904831182151516156105b0576105b061055b565b500290565b600060ff8316806105d657634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036106115761061161055b565b60010192915050565b600060ff821660ff8416808210156106345761063461055b565b90039392505050565b60006001820161064f5761064f61055b565b506001019056fea164736f6c634300080f000a",
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

var RecoveryBeaconMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDKG\",\"name\":\"keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OffchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"occVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structRecoveryBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getEnrollmentResponses\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"playerIdx\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"cipher\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ephermeralKeyPubKeyBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"distributedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountPointBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recoveryPubKeyBytes\",\"type\":\"bytes\"}],\"internalType\":\"structRecoveryBeaconTypes.EnrollmentResponse[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"}],\"name\":\"getLease\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"attemptsLeft\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockPurchased\",\"type\":\"uint256\"}],\"internalType\":\"structRecoveryBeaconTypes.Lease\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMostRecentEnrollmentRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addressToEnroll\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKeyBytes\",\"type\":\"bytes\"}],\"internalType\":\"structRecoveryBeaconTypes.EnrollmentRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"}],\"name\":\"getRecoveredAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getRecovery\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRecoveryRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"playerIdx\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"distributedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"cipher\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ephermeralKeyPubKeyBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recoveryPubKeyBytes\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addressToRecover\",\"type\":\"address\"}],\"internalType\":\"structRecoveryBeaconTypes.RecoveryRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"}],\"name\":\"leaseIsExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"plunderLease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"playerIdx\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"cipher\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ephermeralKeyPubKeyBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"distributedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"accountPointBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recoveryPubKeyBytes\",\"type\":\"bytes\"}],\"internalType\":\"structRecoveryBeaconTypes.EnrollmentResponse\",\"name\":\"enrollmentResponse\",\"type\":\"tuple\"}],\"name\":\"postCipher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"playerIdx\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"distributedPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"cipher\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ephermeralKeyPubKeyBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recoveryPubKeyBytes\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addressToRecover\",\"type\":\"address\"}],\"internalType\":\"structRecoveryBeaconTypes.RecoveryRequest[]\",\"name\":\"recoveryRequests\",\"type\":\"tuple[]\"}],\"name\":\"postRecoveryRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKeyBytes\",\"type\":\"bytes\"}],\"name\":\"requestAccountRecoveryEnrollment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"}],\"name\":\"requestRecoveryLease\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_mostRecentAccountToRecover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"answer\",\"type\":\"bytes\"}],\"name\":\"setRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162003d6638038062003d6683398101604081905262000034916200019c565b818133806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600880546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620000f0565b5050601980546001600160a01b0319166001600160a01b039490941693909317909255601a5550620001d8915050565b336001600160a01b038216036200014a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600980546001600160a01b0319166001600160a01b03838116918217909255600854604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b60008060408385031215620001b057600080fd5b82516001600160a01b0381168114620001c857600080fd5b6020939093015192949293505050565b613b7e80620001e86000396000f3fe6080604052600436106101d85760003560e01c8063acf2a15611610102578063cc31f7dd11610095578063ecaa5ca711610064578063ecaa5ca714610668578063f2fde38b1461068a578063f3fef3a3146106aa578063f8b2cb4f146106bd57600080fd5b8063cc31f7dd14610606578063d0e30db01461062a578063d57fc45a14610632578063e3d0e7121461064857600080fd5b8063b99f904b116100d1578063b99f904b14610577578063bf2732c714610597578063c5a851bb146105b7578063c946b2d4146105e457600080fd5b8063acf2a156146104ed578063ad3c3b601461050d578063afcb95d714610520578063b1dc65a41461055757600080fd5b80636b8ab97d1161017a5780638bc94ccd116101495780638bc94ccd146103ad5780638da5cb5b1461048f5780639769c3fe146104ad5780639ccc7309146104cd57600080fd5b80636b8ab97d1461031e57806379ba50971461033e57806381ff704814610353578063878a1d881461038d57600080fd5b806318f69efe116101b657806318f69efe146102a45780631fcedf21146102d457806355e48749146102f45780636285b9431461030957600080fd5b80630454a05b146101dd5780630de44ac9146101ff578063181f5a7714610255575b600080fd5b3480156101e957600080fd5b506101fd6101f8366004612b93565b6106f3565b005b34801561020b57600080fd5b5061023861021a366004612bea565b6001600160a01b039081166000908152600a60205260409020541690565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561026157600080fd5b5060408051808201909152601a81527f5265636f76657279426561636f6e20312e302e302d616c70686100000000000060208201525b60405161024c9190612c66565b3480156102b057600080fd5b506102c46102bf366004612bea565b610712565b604051901515815260200161024c565b3480156102e057600080fd5b506101fd6102ef366004612c79565b610776565b34801561030057600080fd5b506101fd6107b3565b34801561031557600080fd5b506101fd6107fd565b34801561032a57600080fd5b506101fd610339366004612bea565b61090a565b34801561034a57600080fd5b506101fd6109dd565b34801561035f57600080fd5b50600c54600e54604080516000815264010000000090930463ffffffff16602084015282015260600161024c565b34801561039957600080fd5b506101fd6103a8366004612cd1565b610906565b3480156103b957600080fd5b506104476103c8366004612bea565b604080516080810182526000808252602082018190529181018290526060810191909152506001600160a01b039081166000908152600360209081526040918290208251608081018452815485168152600182015494851692810192909252600160a01b90930460ff1691810191909152600290910154606082015290565b60405161024c919060006080820190506001600160a01b038084511683528060208501511660208401525060ff60408401511660408301526060830151606083015292915050565b34801561049b57600080fd5b506008546001600160a01b0316610238565b3480156104b957600080fd5b506102976104c8366004612bea565b610a92565b3480156104d957600080fd5b506101fd6104e8366004612b93565b610b3e565b3480156104f957600080fd5b50600454610238906001600160a01b031681565b6101fd61051b366004612bea565b610bc8565b34801561052c57600080fd5b50600e546010546040805160008152602081019390935263ffffffff9091169082015260600161024c565b34801561056357600080fd5b506101fd610572366004612d2e565b610d3b565b34801561058357600080fd5b506101fd610592366004612de5565b61120f565b3480156105a357600080fd5b506101fd6105b2366004612f0b565b6113c0565b3480156105c357600080fd5b506105d76105d2366004612bea565b61142e565b60405161024c9190612fee565b3480156105f057600080fd5b506105f961182b565b60405161024c91906130ff565b34801561061257600080fd5b5061061c601a5481565b60405190815260200161024c565b6101fd6118f3565b34801561063e57600080fd5b5061061c601b5481565b34801561065457600080fd5b506101fd610663366004613164565b611919565b34801561067457600080fd5b5061067d612058565b60405161024c9190613252565b34801561069657600080fd5b506101fd6106a5366004612bea565b6123c8565b6101fd6106b836600461335a565b6123d9565b3480156106c957600080fd5b5061061c6106d8366004612bea565b6001600160a01b031660009081526007602052604090205490565b33600090815260066020526040902061070d828483613408565b505050565b6001600160a01b038116600090815260036020526040812060010154600160a01b900460ff1615816107466103e8436134de565b6001600160a01b038516600090815260036020526040902060020154109050818061076e5750805b949350505050565b6001600160a01b03821660009081526005602090815260408220805460018101825590835291208291600702016107ad828261353c565b50505050565b6019546001600160a01b03163381146107f55760405163292f4fb560e01b81523360048201526001600160a01b03821660248201526044015b60405180910390fd5b506000601b55565b336000908152600360205260409020546001600160a01b0316156109085733600081815260036020526040812080546001600160a01b031916815560018101805474ffffffffffffffffffffffffffffffffffffffffff19169055600290810182905590919061087490662386f26fc10000613623565b604051600081818185875af1925050503d80600081146108b0576040519150601f19603f3d011682016040523d82523d6000602084013e6108b5565b606091505b50509050806109065760405162461bcd60e51b815260206004820152601c60248201527f4661696c656420746f2066696e69736820706c756e646572696e672e0000000060448201526064016107ec565b505b565b6001600160a01b038116600090815260056020526040812061092b91612947565b6001600160a01b0381166000908152600a6020908152604080832080546001600160a01b03191690556006909152812061096491612968565b6001600160a01b038116600090815260036020526040812080546001600160a01b031990811682556001808301805474ffffffffffffffffffffffffffffffffffffffffff1916905560029092018390556004805482169055825416825581906109ce9082612968565b506109069050600260006129a2565b6009546001600160a01b03163314610a375760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016107ec565b600880546001600160a01b0319808216339081179093556009805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6001600160a01b0381166000908152600660205260409020805460609190610ab990613386565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae590613386565b8015610b325780601f10610b0757610100808354040283529160200191610b32565b820191906000526020600020905b815481529060010190602001808311610b1557829003601f168201915b50505050509050919050565b6040518060400160405280336001600160a01b0316815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050825181546001600160a01b0319166001600160a01b039091161781556020830151909150600190610bc19082613645565b5050505050565b662386f26fc10000341015610c1f5760405162461bcd60e51b815260206004820152601360248201527f496e73756666696369656e742066756e64732e0000000000000000000000000060448201526064016107ec565b610c2881610712565b610c9a5760405162461bcd60e51b815260206004820152602160248201527f5468652063757272656e74206c65617365206973207374696c6c2076616c696460448201527f2e0000000000000000000000000000000000000000000000000000000000000060648201526084016107ec565b604080516080810182523381526001600160a01b039283166020808301828152600584860190815243606086019081526000858152600390945295909220935184549087166001600160a01b03199182161785559051600185018054935160ff16600160a01b0274ffffffffffffffffffffffffffffffffffffffffff19909416919097161791909117909455915160029091015560048054909216179055565b604080516101008082018352600b5460ff808216845291810464ffffffffff166020808501919091526601000000000000820463ffffffff908116858701526a01000000000000000000008304811660608601526e01000000000000000000000000000083048116608086015272010000000000000000000000000000000000008304811660a086015276010000000000000000000000000000000000000000000083041660c08501527a01000000000000000000000000000000000000000000000000000090910462ffffff1660e084015233600090815260118252939093205491928b01359116610e705760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016107ec565b600e548a3514610ec25760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016107ec565b610ed0898989898989612583565b8151610edd90600161371b565b60ff168614610f2e5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016107ec565b858414610f7d5760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016107ec565b60008989604051610f8f929190613740565b604051908190038120610fa6918d90602001613750565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8981101561114c576000600185898460208110610ff557610ff5613705565b61100291901a601b61371b565b8e8e8681811061101457611014613705565b905060200201358d8d8781811061102d5761102d613705565b905060200201356040516000815260200160405260405161106a949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561108c573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526012602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506111255760405162461bcd60e51b815260206004820152600f60248201527f7369676e6174757265206572726f72000000000000000000000000000000000060448201526064016107ec565b826020015160080260ff166001901b840193505080806111449061376c565b915050610fd6565b5081827e0101010101010101010101010101010101010101010101010101010101010116146111bd5760405162461bcd60e51b815260206004820152601060248201527f6475706c6963617465207369676e65720000000000000000000000000000000060448201526064016107ec565b5050604080516020601f8c018190048102820181019092528a815261120392508c35918491908d908d908190840183828082843760009201919091525061262092505050565b50505050505050505050565b60008282600081811061122457611224613705565b90506020028101906112369190613785565b6112489061010081019060e001612bea565b905061125381610712565b156112a05760405162461bcd60e51b815260206004820152601060248201527f4c6561736520697320657870697265640000000000000000000000000000000060448201526064016107ec565b6001600160a01b03818116600090815260036020908152604091829020825160808101845281548516808252600183015495861693820193909352600160a01b90940460ff1692840192909252600290910154606083015233146113465760405162461bcd60e51b815260206004820152601460248201527f4e6f7420746865206c65617365206f776e65722e00000000000000000000000060448201526064016107ec565b611352600260006129a2565b60005b60ff8116841115610bc157600285858360ff1681811061137757611377613705565b90506020028101906113899190613785565b8154600181018355600092835260209092209091600802016113ab82826137b8565b505080806113b8906138c2565b915050611355565b6019546001600160a01b03163381146113fd5760405163292f4fb560e01b81523360048201526001600160a01b03821660248201526044016107ec565b815160405161140f91906020016138e1565b60408051601f198184030181529190528051602090910120601b555050565b6001600160a01b0381166000908152600560209081526040808320805482518185028101850190935280835260609492939192909184015b8282101561182057600084815260209081902060408051610100808201835260078702909301805460ff808216845294900490931693810193909352600182018054918401916114b590613386565b80601f01602080910402602001604051908101604052809291908181526020018280546114e190613386565b801561152e5780601f106115035761010080835404028352916020019161152e565b820191906000526020600020905b81548152906001019060200180831161151157829003601f168201915b5050505050815260200160028201805461154790613386565b80601f016020809104026020016040519081016040528092919081815260200182805461157390613386565b80156115c05780601f10611595576101008083540402835291602001916115c0565b820191906000526020600020905b8154815290600101906020018083116115a357829003601f168201915b505050505081526020016003820180546115d990613386565b80601f016020809104026020016040519081016040528092919081815260200182805461160590613386565b80156116525780601f1061162757610100808354040283529160200191611652565b820191906000526020600020905b81548152906001019060200180831161163557829003601f168201915b5050505050815260200160048201805461166b90613386565b80601f016020809104026020016040519081016040528092919081815260200182805461169790613386565b80156116e45780601f106116b9576101008083540402835291602001916116e4565b820191906000526020600020905b8154815290600101906020018083116116c757829003601f168201915b505050505081526020016005820180546116fd90613386565b80601f016020809104026020016040519081016040528092919081815260200182805461172990613386565b80156117765780601f1061174b57610100808354040283529160200191611776565b820191906000526020600020905b81548152906001019060200180831161175957829003601f168201915b5050505050815260200160068201805461178f90613386565b80601f01602080910402602001604051908101604052809291908181526020018280546117bb90613386565b80156118085780601f106117dd57610100808354040283529160200191611808565b820191906000526020600020905b8154815290600101906020018083116117eb57829003601f168201915b50505050508152505081526020019060010190611466565b505050509050919050565b60408051808201909152600081526060602082015260408051808201909152600080546001600160a01b031682526001805460208401919061186c90613386565b80601f016020809104026020016040519081016040528092919081815260200182805461189890613386565b80156118e55780601f106118ba576101008083540402835291602001916118e5565b820191906000526020600020905b8154815290600101906020018083116118c857829003601f168201915b505050505081525050905090565b33600090815260076020526040812080543492906119129084906138f3565b9091555050565b611921612766565b601f8911156119725760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79206f7261636c65730000000000000000000000000000000060448201526064016107ec565b8887146119c15760405162461bcd60e51b815260206004820152601660248201527f6f7261636c65206c656e677468206d69736d617463680000000000000000000060448201526064016107ec565b886119cd87600361390b565b60ff1610611a1d5760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016107ec565b611a298660ff166127c0565b6040805160e060208c02808301820190935260c082018c815260009383928f918f918291908601908490808284376000920191909152505050908252506040805160208c810282810182019093528c82529283019290918d918d91829185019084908082843760009201919091525050509082525060ff891660208083019190915260408051601f8a01839004830281018301825289815292019190899089908190840183828082843760009201919091525050509082525067ffffffffffffffff861660208083019190915260408051601f870183900483028101830182528681529201919086908690819084018382808284376000920182905250939094525050600b805465ffffffffff001916905560135492935090505b81811015611bef57600060138281548110611b6157611b61613705565b6000918252602082200154601480546001600160a01b0390921693509084908110611b8e57611b8e613705565b60009182526020808320909101546001600160a01b039485168352601282526040808420805461ffff1916905594168252601190529190912080546dffffffffffffffffffffffffffff191690555080611be78161376c565b915050611b44565b50611bfc601360006129c3565b611c08601460006129c3565b60005b825151811015611e90576012600084600001518381518110611c2f57611c2f613705565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611ca35760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016107ec565b604080518082019091526001815260ff821660208201528351805160129160009185908110611cd457611cd4613705565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051601192919084908110611d3f57611d3f613705565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611db35760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016107ec565b60405180606001604052806001151581526020018260ff16815260200160006bffffffffffffffffffffffff168152506011600085602001518481518110611dfd57611dfd613705565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516bffffffffffffffffffffffff1662010000026dffffffffffffffffffffffff00001960ff959095166101000261ff00199315159390931661ffff1990941693909317919091179290921617905580611e888161376c565b915050611c0b565b5081518051611ea7916013916020909101906129e1565b506020808301518051611ebe9260149201906129e1565b506040820151600b805460ff191660ff909216919091179055600c805467ffffffff0000000019811664010000000063ffffffff43811682029283179094558204831692600092611f16929082169116176001613934565b905080600c60006101000a81548163ffffffff021916908363ffffffff1602179055506000611f6a46308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a00151612810565b905080600e600001819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05838284886000015189602001518a604001518b606001518c608001518d60a00151604051611fcd999897969594939291906139a0565b60405180910390a1600b546601000000000000900463ffffffff1660005b8651518110156120455781601582601f811061200957612009613705565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550808061203d9061376c565b915050611feb565b5050505050505050505050505050505050565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156123bf57600084815260209081902060408051610100810190915260088502909101805460ff16825260018101805492939192918401916120c090613386565b80601f01602080910402602001604051908101604052809291908181526020018280546120ec90613386565b80156121395780601f1061210e57610100808354040283529160200191612139565b820191906000526020600020905b81548152906001019060200180831161211c57829003601f168201915b5050505050815260200160028201805461215290613386565b80601f016020809104026020016040519081016040528092919081815260200182805461217e90613386565b80156121cb5780601f106121a0576101008083540402835291602001916121cb565b820191906000526020600020905b8154815290600101906020018083116121ae57829003601f168201915b505050505081526020016003820180546121e490613386565b80601f016020809104026020016040519081016040528092919081815260200182805461221090613386565b801561225d5780601f106122325761010080835404028352916020019161225d565b820191906000526020600020905b81548152906001019060200180831161224057829003601f168201915b5050505050815260200160048201805461227690613386565b80601f01602080910402602001604051908101604052809291908181526020018280546122a290613386565b80156122ef5780601f106122c4576101008083540402835291602001916122ef565b820191906000526020600020905b8154815290600101906020018083116122d257829003601f168201915b5050505050815260200160058201805461230890613386565b80601f016020809104026020016040519081016040528092919081815260200182805461233490613386565b80156123815780601f1061235657610100808354040283529160200191612381565b820191906000526020600020905b81548152906001019060200180831161236457829003601f168201915b505050918352505060068201546001600160a01b0390811660208084019190915260079093015416604090910152908252600192909201910161207c565b50505050905090565b6123d0612766565b6109068161289d565b6001600160a01b03821633148061240a5750336000908152600a60205260409020546001600160a01b038381169116145b6124565760405162461bcd60e51b815260206004820152600e60248201527f4163636573732064656e6965642e00000000000000000000000000000000000060448201526064016107ec565b6001600160a01b0382166000908152600760205260409020548111156124be5760405162461bcd60e51b815260206004820152601460248201527f496e73756666696369656e742062616c616e636500000000000000000000000060448201526064016107ec565b6001600160a01b038216600090815260076020526040812080548392906124e69084906134de565b9091555050604051600090339083908381818185875af1925050503d806000811461252d576040519150601f19603f3d011682016040523d82523d6000602084013e612532565b606091505b505090508061070d5760405162461bcd60e51b815260206004820152601560248201527f4661696c656420746f2073656e642066756e64732e000000000000000000000060448201526064016107ec565b6000612590826020613a36565b61259b856020613a36565b6125a7886101446138f3565b6125b191906138f3565b6125bb91906138f3565b6125c69060006138f3565b90503681146126175760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016107ec565b50505050505050565b60025460000361262f57505050565b6000818060200190518101906126459190613a55565b905080604001511561268b5780516020808301516001600160a01b039081166000908152600a909252604090912080546001600160a01b031916919092161790556126f0565b80516001600160a01b031660009081526003602052604090206001810154600160a01b900460ff16156126ee57600181018054600160a01b900460ff169060146126d483613acc565b91906101000a81548160ff021916908360ff160217905550505b505b6126fc600260006129a2565b8051602080830151604080513381526001600160a01b039485169381019390935292168183015260608101869052905164ffffffffff8516917ffb8260fafd97df85b3771f6532a09a8af1daeac85d6938ee302237174c91a9c9919081900360800190a250505050565b6008546001600160a01b031633146109085760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107ec565b806000106109065760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016107ec565b6000808a8a8a8a8a8a8a8a8a60405160200161283499989796959493929190613ae9565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036128f55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107ec565b600980546001600160a01b0319166001600160a01b03838116918217909255600854604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b50805460008255600702906000526020600020908101906109069190612a46565b50805461297490613386565b6000825580601f10612984575050565b601f0160209004906000526020600020908101906109069190612ab5565b50805460008255600802906000526020600020908101906109069190612aca565b50805460008255906000526020600020908101906109069190612ab5565b828054828255906000526020600020908101928215612a36579160200282015b82811115612a3657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190612a01565b50612a42929150612ab5565b5090565b80821115612a4257805461ffff191681556000612a666001830182612968565b612a74600283016000612968565b612a82600383016000612968565b612a90600483016000612968565b612a9e600583016000612968565b612aac600683016000612968565b50600701612a46565b5b80821115612a425760008155600101612ab6565b80821115612a4257805460ff191681556000612ae96001830182612968565b612af7600283016000612968565b612b05600383016000612968565b612b13600483016000612968565b612b21600583016000612968565b506006810180546001600160a01b03199081169091556007820180549091169055600801612aca565b60008083601f840112612b5c57600080fd5b50813567ffffffffffffffff811115612b7457600080fd5b602083019150836020828501011115612b8c57600080fd5b9250929050565b60008060208385031215612ba657600080fd5b823567ffffffffffffffff811115612bbd57600080fd5b612bc985828601612b4a565b90969095509350505050565b6001600160a01b038116811461090657600080fd5b600060208284031215612bfc57600080fd5b8135612c0781612bd5565b9392505050565b60005b83811015612c29578181015183820152602001612c11565b838111156107ad5750506000910152565b60008151808452612c52816020860160208601612c0e565b601f01601f19169290920160200192915050565b602081526000612c076020830184612c3a565b60008060408385031215612c8c57600080fd5b8235612c9781612bd5565b9150602083013567ffffffffffffffff811115612cb357600080fd5b83016101008186031215612cc657600080fd5b809150509250929050565b600060608284031215612ce357600080fd5b50919050565b60008083601f840112612cfb57600080fd5b50813567ffffffffffffffff811115612d1357600080fd5b6020830191508360208260051b8501011115612b8c57600080fd5b60008060008060008060008060e0898b031215612d4a57600080fd5b606089018a811115612d5b57600080fd5b8998503567ffffffffffffffff80821115612d7557600080fd5b612d818c838d01612b4a565b909950975060808b0135915080821115612d9a57600080fd5b612da68c838d01612ce9565b909750955060a08b0135915080821115612dbf57600080fd5b50612dcc8b828c01612ce9565b999c989b50969995989497949560c00135949350505050565b60008060208385031215612df857600080fd5b823567ffffffffffffffff811115612e0f57600080fd5b612bc985828601612ce9565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715612e5457612e54612e1b565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715612e8357612e83612e1b565b604052919050565b600082601f830112612e9c57600080fd5b8135602067ffffffffffffffff821115612eb857612eb8612e1b565b8160051b612ec7828201612e5a565b9283528481018201928281019087851115612ee157600080fd5b83870192505b84831015612f0057823582529183019190830190612ee7565b979650505050505050565b60006020808385031215612f1e57600080fd5b823567ffffffffffffffff80821115612f3657600080fd5b9084019060408287031215612f4a57600080fd5b612f52612e31565b823582811115612f6157600080fd5b8301601f81018813612f7257600080fd5b803583811115612f8457612f84612e1b565b612f96601f8201601f19168701612e5a565b8181528987838501011115612faa57600080fd5b818784018883013760008783830101528084525050508383013582811115612fd157600080fd5b612fdd88828601612e8b565b948201949094529695505050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156130f157888303603f190185528151805160ff1684526101008189015160ff8116868b01525087820151818987015261305482870182612c3a565b9150506060808301518683038288015261306e8382612c3a565b92505050608080830151868303828801526130898382612c3a565b9250505060a080830151868303828801526130a48382612c3a565b9250505060c080830151868303828801526130bf8382612c3a565b9250505060e080830151925085820381870152506130dd8183612c3a565b968901969450505090860190600101613015565b509098975050505050505050565b602081526001600160a01b0382511660208201526000602083015160408084015261076e6060840182612c3a565b60ff8116811461090657600080fd5b80356131478161312d565b919050565b803567ffffffffffffffff8116811461314757600080fd5b60008060008060008060008060008060c08b8d03121561318357600080fd5b8a3567ffffffffffffffff8082111561319b57600080fd5b6131a78e838f01612ce9565b909c509a5060208d01359150808211156131c057600080fd5b6131cc8e838f01612ce9565b909a5098508891506131e060408e0161313c565b975060608d01359150808211156131f657600080fd5b6132028e838f01612b4a565b909750955085915061321660808e0161314c565b945060a08d013591508082111561322c57600080fd5b506132398d828e01612b4a565b915080935050809150509295989b9194979a5092959850565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156130f157888303603f190185528151805160ff16845261010088820151818a8701526132ab82870182612c3a565b91505087820151858203898701526132c38282612c3a565b915050606080830151868303828801526132dd8382612c3a565b92505050608080830151868303828801526132f88382612c3a565b9250505060a080830151868303828801526133138382612c3a565b9250505060c080830151613331828801826001600160a01b03169052565b505060e0918201516001600160a01b031694909101939093529386019390860190600101613279565b6000806040838503121561336d57600080fd5b823561337881612bd5565b946020939093013593505050565b600181811c9082168061339a57607f821691505b602082108103612ce357634e487b7160e01b600052602260045260246000fd5b601f82111561070d57600081815260208120601f850160051c810160208610156133e15750805b601f850160051c820191505b81811015613400578281556001016133ed565b505050505050565b67ffffffffffffffff83111561342057613420612e1b565b6134348361342e8354613386565b836133ba565b6000601f84116001811461346857600085156134505750838201355b600019600387901b1c1916600186901b178355610bc1565b600083815260209020601f19861690835b828110156134995786850135825560209485019460019092019101613479565b50868210156134b65760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156134f0576134f06134c8565b500390565b6000808335601e1984360301811261350c57600080fd5b83018035915067ffffffffffffffff82111561352757600080fd5b602001915036819003821315612b8c57600080fd5b81356135478161312d565b815460ff191660ff82161782555060208201356135638161312d565b815461ff00191660089190911b61ff001617815561358460408301836134f5565b613592818360018601613408565b50506135a160608301836134f5565b6135af818360028601613408565b50506135be60808301836134f5565b6135cc818360038601613408565b50506135db60a08301836134f5565b6135e9818360048601613408565b50506135f860c08301836134f5565b613606818360058601613408565b505061361560e08301836134f5565b6107ad818360068601613408565b60008261364057634e487b7160e01b600052601260045260246000fd5b500490565b815167ffffffffffffffff81111561365f5761365f612e1b565b6136738161366d8454613386565b846133ba565b602080601f8311600181146136a857600084156136905750858301515b600019600386901b1c1916600185901b178555613400565b600085815260208120601f198616915b828110156136d7578886015182559484019460019091019084016136b8565b50858210156136f55787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff84168060ff03821115613738576137386134c8565b019392505050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60006001820161377e5761377e6134c8565b5060010190565b6000823560fe1983360301811261379b57600080fd5b9190910192915050565b600081356137b281612bd5565b92915050565b81356137c38161312d565b815460ff191660ff8216178255506137de60208301836134f5565b6137ec818360018601613408565b50506137fb60408301836134f5565b613809818360028601613408565b505061381860608301836134f5565b613826818360038601613408565b505061383560808301836134f5565b613843818360048601613408565b505061385260a08301836134f5565b613860818360058601613408565b505060c082013561387081612bd5565b6006820180546001600160a01b0319166001600160a01b038316179055506138be61389d60e084016137a5565b600783016001600160a01b0382166001600160a01b03198254161781555050565b5050565b600060ff821660ff81036138d8576138d86134c8565b60010192915050565b6000825161379b818460208701612c0e565b60008219821115613906576139066134c8565b500190565b600060ff821660ff84168160ff048111821515161561392c5761392c6134c8565b029392505050565b600063ffffffff808316818516808303821115613953576139536134c8565b01949350505050565b600081518084526020808501945080840160005b838110156139955781516001600160a01b031687529582019590820190600101613970565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526139d08184018a61395c565b905082810360808401526139e4818961395c565b905060ff871660a084015282810360c0840152613a018187612c3a565b905067ffffffffffffffff851660e0840152828103610100840152613a268185612c3a565b9c9b505050505050505050505050565b6000816000190483118215151615613a5057613a506134c8565b500290565b600060608284031215613a6757600080fd5b6040516060810181811067ffffffffffffffff82111715613a8a57613a8a612e1b565b6040528251613a9881612bd5565b81526020830151613aa881612bd5565b602082015260408301518015158114613ac057600080fd5b60408201529392505050565b600060ff821680613adf57613adf6134c8565b6000190192915050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152613b238285018b61395c565b91508382036080850152613b37828a61395c565b915060ff881660a085015283820360c0850152613b548288612c3a565b90861660e08501528381036101008501529050613a268185612c3a56fea164736f6c634300080f000a",
}

var RecoveryBeaconABI = RecoveryBeaconMetaData.ABI

var RecoveryBeaconBin = RecoveryBeaconMetaData.Bin

func DeployRecoveryBeacon(auth *bind.TransactOpts, backend bind.ContractBackend, keyProvider common.Address, keyID [32]byte) (common.Address, *types.Transaction, *RecoveryBeacon, error) {
	parsed, err := RecoveryBeaconMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecoveryBeaconBin), backend, keyProvider, keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecoveryBeacon{RecoveryBeaconCaller: RecoveryBeaconCaller{contract: contract}, RecoveryBeaconTransactor: RecoveryBeaconTransactor{contract: contract}, RecoveryBeaconFilterer: RecoveryBeaconFilterer{contract: contract}}, nil
}

type RecoveryBeacon struct {
	RecoveryBeaconCaller
	RecoveryBeaconTransactor
	RecoveryBeaconFilterer
}

type RecoveryBeaconCaller struct {
	contract *bind.BoundContract
}

type RecoveryBeaconTransactor struct {
	contract *bind.BoundContract
}

type RecoveryBeaconFilterer struct {
	contract *bind.BoundContract
}

type RecoveryBeaconSession struct {
	Contract     *RecoveryBeacon
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconCallerSession struct {
	Contract *RecoveryBeaconCaller
	CallOpts bind.CallOpts
}

type RecoveryBeaconTransactorSession struct {
	Contract     *RecoveryBeaconTransactor
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconRaw struct {
	Contract *RecoveryBeacon
}

type RecoveryBeaconCallerRaw struct {
	Contract *RecoveryBeaconCaller
}

type RecoveryBeaconTransactorRaw struct {
	Contract *RecoveryBeaconTransactor
}

func NewRecoveryBeacon(address common.Address, backend bind.ContractBackend) (*RecoveryBeacon, error) {
	contract, err := bindRecoveryBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeacon{RecoveryBeaconCaller: RecoveryBeaconCaller{contract: contract}, RecoveryBeaconTransactor: RecoveryBeaconTransactor{contract: contract}, RecoveryBeaconFilterer: RecoveryBeaconFilterer{contract: contract}}, nil
}

func NewRecoveryBeaconCaller(address common.Address, caller bind.ContractCaller) (*RecoveryBeaconCaller, error) {
	contract, err := bindRecoveryBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconCaller{contract: contract}, nil
}

func NewRecoveryBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoveryBeaconTransactor, error) {
	contract, err := bindRecoveryBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconTransactor{contract: contract}, nil
}

func NewRecoveryBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoveryBeaconFilterer, error) {
	contract, err := bindRecoveryBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconFilterer{contract: contract}, nil
}

func bindRecoveryBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoveryBeaconABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_RecoveryBeacon *RecoveryBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeacon.Contract.RecoveryBeaconCaller.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeacon *RecoveryBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.RecoveryBeaconTransactor.contract.Transfer(opts)
}

func (_RecoveryBeacon *RecoveryBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.RecoveryBeaconTransactor.contract.Transact(opts, method, params...)
}

func (_RecoveryBeacon *RecoveryBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeacon.Contract.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.contract.Transfer(opts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.contract.Transact(opts, method, params...)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _RecoveryBeacon.Contract.GetBalance(&_RecoveryBeacon.CallOpts, addr)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _RecoveryBeacon.Contract.GetBalance(&_RecoveryBeacon.CallOpts, addr)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) GetEnrollmentResponses(opts *bind.CallOpts, user common.Address) ([]RecoveryBeaconTypesEnrollmentResponse, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "getEnrollmentResponses", user)

	if err != nil {
		return *new([]RecoveryBeaconTypesEnrollmentResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]RecoveryBeaconTypesEnrollmentResponse)).(*[]RecoveryBeaconTypesEnrollmentResponse)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) GetEnrollmentResponses(user common.Address) ([]RecoveryBeaconTypesEnrollmentResponse, error) {
	return _RecoveryBeacon.Contract.GetEnrollmentResponses(&_RecoveryBeacon.CallOpts, user)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) GetEnrollmentResponses(user common.Address) ([]RecoveryBeaconTypesEnrollmentResponse, error) {
	return _RecoveryBeacon.Contract.GetEnrollmentResponses(&_RecoveryBeacon.CallOpts, user)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) GetLease(opts *bind.CallOpts, accountToRecover common.Address) (RecoveryBeaconTypesLease, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "getLease", accountToRecover)

	if err != nil {
		return *new(RecoveryBeaconTypesLease), err
	}

	out0 := *abi.ConvertType(out[0], new(RecoveryBeaconTypesLease)).(*RecoveryBeaconTypesLease)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) GetLease(accountToRecover common.Address) (RecoveryBeaconTypesLease, error) {
	return _RecoveryBeacon.Contract.GetLease(&_RecoveryBeacon.CallOpts, accountToRecover)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) GetLease(accountToRecover common.Address) (RecoveryBeaconTypesLease, error) {
	return _RecoveryBeacon.Contract.GetLease(&_RecoveryBeacon.CallOpts, accountToRecover)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) GetMostRecentEnrollmentRequest(opts *bind.CallOpts) (RecoveryBeaconTypesEnrollmentRequest, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "getMostRecentEnrollmentRequest")

	if err != nil {
		return *new(RecoveryBeaconTypesEnrollmentRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(RecoveryBeaconTypesEnrollmentRequest)).(*RecoveryBeaconTypesEnrollmentRequest)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) GetMostRecentEnrollmentRequest() (RecoveryBeaconTypesEnrollmentRequest, error) {
	return _RecoveryBeacon.Contract.GetMostRecentEnrollmentRequest(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) GetMostRecentEnrollmentRequest() (RecoveryBeaconTypesEnrollmentRequest, error) {
	return _RecoveryBeacon.Contract.GetMostRecentEnrollmentRequest(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) GetRecoveredAddress(opts *bind.CallOpts, recoverer common.Address) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "getRecoveredAddress", recoverer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) GetRecoveredAddress(recoverer common.Address) (common.Address, error) {
	return _RecoveryBeacon.Contract.GetRecoveredAddress(&_RecoveryBeacon.CallOpts, recoverer)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) GetRecoveredAddress(recoverer common.Address) (common.Address, error) {
	return _RecoveryBeacon.Contract.GetRecoveredAddress(&_RecoveryBeacon.CallOpts, recoverer)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) GetRecovery(opts *bind.CallOpts, user common.Address) ([]byte, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "getRecovery", user)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) GetRecovery(user common.Address) ([]byte, error) {
	return _RecoveryBeacon.Contract.GetRecovery(&_RecoveryBeacon.CallOpts, user)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) GetRecovery(user common.Address) ([]byte, error) {
	return _RecoveryBeacon.Contract.GetRecovery(&_RecoveryBeacon.CallOpts, user)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) GetRecoveryRequests(opts *bind.CallOpts) ([]RecoveryBeaconTypesRecoveryRequest, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "getRecoveryRequests")

	if err != nil {
		return *new([]RecoveryBeaconTypesRecoveryRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]RecoveryBeaconTypesRecoveryRequest)).(*[]RecoveryBeaconTypesRecoveryRequest)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) GetRecoveryRequests() ([]RecoveryBeaconTypesRecoveryRequest, error) {
	return _RecoveryBeacon.Contract.GetRecoveryRequests(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) GetRecoveryRequests() ([]RecoveryBeaconTypesRecoveryRequest, error) {
	return _RecoveryBeacon.Contract.GetRecoveryRequests(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "latestConfigDetails")

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

func (_RecoveryBeacon *RecoveryBeaconSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _RecoveryBeacon.Contract.LatestConfigDetails(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _RecoveryBeacon.Contract.LatestConfigDetails(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

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

func (_RecoveryBeacon *RecoveryBeaconSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _RecoveryBeacon.Contract.LatestConfigDigestAndEpoch(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _RecoveryBeacon.Contract.LatestConfigDigestAndEpoch(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) LeaseIsExpired(opts *bind.CallOpts, accountToRecover common.Address) (bool, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "leaseIsExpired", accountToRecover)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) LeaseIsExpired(accountToRecover common.Address) (bool, error) {
	return _RecoveryBeacon.Contract.LeaseIsExpired(&_RecoveryBeacon.CallOpts, accountToRecover)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) LeaseIsExpired(accountToRecover common.Address) (bool, error) {
	return _RecoveryBeacon.Contract.LeaseIsExpired(&_RecoveryBeacon.CallOpts, accountToRecover)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) Owner() (common.Address, error) {
	return _RecoveryBeacon.Contract.Owner(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) Owner() (common.Address, error) {
	return _RecoveryBeacon.Contract.Owner(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) SKeyID() ([32]byte, error) {
	return _RecoveryBeacon.Contract.SKeyID(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) SKeyID() ([32]byte, error) {
	return _RecoveryBeacon.Contract.SKeyID(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) SMostRecentAccountToRecover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "s_mostRecentAccountToRecover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeacon.Contract.SMostRecentAccountToRecover(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeacon.Contract.SMostRecentAccountToRecover(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) SProvingKeyHash() ([32]byte, error) {
	return _RecoveryBeacon.Contract.SProvingKeyHash(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _RecoveryBeacon.Contract.SProvingKeyHash(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RecoveryBeacon.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RecoveryBeacon *RecoveryBeaconSession) TypeAndVersion() (string, error) {
	return _RecoveryBeacon.Contract.TypeAndVersion(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconCallerSession) TypeAndVersion() (string, error) {
	return _RecoveryBeacon.Contract.TypeAndVersion(&_RecoveryBeacon.CallOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "acceptOwnership")
}

func (_RecoveryBeacon *RecoveryBeaconSession) AcceptOwnership() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.AcceptOwnership(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.AcceptOwnership(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "deposit")
}

func (_RecoveryBeacon *RecoveryBeaconSession) Deposit() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Deposit(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) Deposit() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Deposit(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) ExposeType(opts *bind.TransactOpts, arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "exposeType", arg0)
}

func (_RecoveryBeacon *RecoveryBeaconSession) ExposeType(arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.ExposeType(&_RecoveryBeacon.TransactOpts, arg0)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) ExposeType(arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.ExposeType(&_RecoveryBeacon.TransactOpts, arg0)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "keyGenerated", kd)
}

func (_RecoveryBeacon *RecoveryBeaconSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.KeyGenerated(&_RecoveryBeacon.TransactOpts, kd)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.KeyGenerated(&_RecoveryBeacon.TransactOpts, kd)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "newKeyRequested")
}

func (_RecoveryBeacon *RecoveryBeaconSession) NewKeyRequested() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.NewKeyRequested(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.NewKeyRequested(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) PlunderLease(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "plunderLease")
}

func (_RecoveryBeacon *RecoveryBeaconSession) PlunderLease() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.PlunderLease(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) PlunderLease() (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.PlunderLease(&_RecoveryBeacon.TransactOpts)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) PostCipher(opts *bind.TransactOpts, user common.Address, enrollmentResponse RecoveryBeaconTypesEnrollmentResponse) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "postCipher", user, enrollmentResponse)
}

func (_RecoveryBeacon *RecoveryBeaconSession) PostCipher(user common.Address, enrollmentResponse RecoveryBeaconTypesEnrollmentResponse) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.PostCipher(&_RecoveryBeacon.TransactOpts, user, enrollmentResponse)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) PostCipher(user common.Address, enrollmentResponse RecoveryBeaconTypesEnrollmentResponse) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.PostCipher(&_RecoveryBeacon.TransactOpts, user, enrollmentResponse)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) PostRecoveryRequests(opts *bind.TransactOpts, recoveryRequests []RecoveryBeaconTypesRecoveryRequest) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "postRecoveryRequests", recoveryRequests)
}

func (_RecoveryBeacon *RecoveryBeaconSession) PostRecoveryRequests(recoveryRequests []RecoveryBeaconTypesRecoveryRequest) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.PostRecoveryRequests(&_RecoveryBeacon.TransactOpts, recoveryRequests)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) PostRecoveryRequests(recoveryRequests []RecoveryBeaconTypesRecoveryRequest) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.PostRecoveryRequests(&_RecoveryBeacon.TransactOpts, recoveryRequests)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) RequestAccountRecoveryEnrollment(opts *bind.TransactOpts, publicKeyBytes []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "requestAccountRecoveryEnrollment", publicKeyBytes)
}

func (_RecoveryBeacon *RecoveryBeaconSession) RequestAccountRecoveryEnrollment(publicKeyBytes []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.RequestAccountRecoveryEnrollment(&_RecoveryBeacon.TransactOpts, publicKeyBytes)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) RequestAccountRecoveryEnrollment(publicKeyBytes []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.RequestAccountRecoveryEnrollment(&_RecoveryBeacon.TransactOpts, publicKeyBytes)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) RequestRecoveryLease(opts *bind.TransactOpts, accountToRecover common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "requestRecoveryLease", accountToRecover)
}

func (_RecoveryBeacon *RecoveryBeaconSession) RequestRecoveryLease(accountToRecover common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.RequestRecoveryLease(&_RecoveryBeacon.TransactOpts, accountToRecover)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) RequestRecoveryLease(accountToRecover common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.RequestRecoveryLease(&_RecoveryBeacon.TransactOpts, accountToRecover)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) Reset(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "reset", addr)
}

func (_RecoveryBeacon *RecoveryBeaconSession) Reset(addr common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Reset(&_RecoveryBeacon.TransactOpts, addr)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) Reset(addr common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Reset(&_RecoveryBeacon.TransactOpts, addr)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_RecoveryBeacon *RecoveryBeaconSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.SetConfig(&_RecoveryBeacon.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.SetConfig(&_RecoveryBeacon.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) SetRecovery(opts *bind.TransactOpts, answer []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "setRecovery", answer)
}

func (_RecoveryBeacon *RecoveryBeaconSession) SetRecovery(answer []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.SetRecovery(&_RecoveryBeacon.TransactOpts, answer)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) SetRecovery(answer []byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.SetRecovery(&_RecoveryBeacon.TransactOpts, answer)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "transferOwnership", to)
}

func (_RecoveryBeacon *RecoveryBeaconSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.TransferOwnership(&_RecoveryBeacon.TransactOpts, to)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.TransferOwnership(&_RecoveryBeacon.TransactOpts, to)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_RecoveryBeacon *RecoveryBeaconSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Transmit(&_RecoveryBeacon.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Transmit(&_RecoveryBeacon.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_RecoveryBeacon *RecoveryBeaconTransactor) Withdraw(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RecoveryBeacon.contract.Transact(opts, "withdraw", addr, amount)
}

func (_RecoveryBeacon *RecoveryBeaconSession) Withdraw(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Withdraw(&_RecoveryBeacon.TransactOpts, addr, amount)
}

func (_RecoveryBeacon *RecoveryBeaconTransactorSession) Withdraw(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RecoveryBeacon.Contract.Withdraw(&_RecoveryBeacon.TransactOpts, addr, amount)
}

type RecoveryBeaconConfigSetIterator struct {
	Event *RecoveryBeaconConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconConfigSet)
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
		it.Event = new(RecoveryBeaconConfigSet)
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

func (it *RecoveryBeaconConfigSetIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconConfigSet struct {
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

func (_RecoveryBeacon *RecoveryBeaconFilterer) FilterConfigSet(opts *bind.FilterOpts) (*RecoveryBeaconConfigSetIterator, error) {

	logs, sub, err := _RecoveryBeacon.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconConfigSetIterator{contract: _RecoveryBeacon.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_RecoveryBeacon *RecoveryBeaconFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconConfigSet) (event.Subscription, error) {

	logs, sub, err := _RecoveryBeacon.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconConfigSet)
				if err := _RecoveryBeacon.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_RecoveryBeacon *RecoveryBeaconFilterer) ParseConfigSet(log types.Log) (*RecoveryBeaconConfigSet, error) {
	event := new(RecoveryBeaconConfigSet)
	if err := _RecoveryBeacon.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconNewTransmissionIterator struct {
	Event *RecoveryBeaconNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconNewTransmission)
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
		it.Event = new(RecoveryBeaconNewTransmission)
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

func (it *RecoveryBeaconNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconNewTransmission struct {
	EpochAndRound    *big.Int
	Transmitter      common.Address
	AccountToRecover common.Address
	Recoverer        common.Address
	ConfigDigest     [32]byte
	Raw              types.Log
}

func (_RecoveryBeacon *RecoveryBeaconFilterer) FilterNewTransmission(opts *bind.FilterOpts, epochAndRound []*big.Int) (*RecoveryBeaconNewTransmissionIterator, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _RecoveryBeacon.contract.FilterLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconNewTransmissionIterator{contract: _RecoveryBeacon.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_RecoveryBeacon *RecoveryBeaconFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconNewTransmission, epochAndRound []*big.Int) (event.Subscription, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _RecoveryBeacon.contract.WatchLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconNewTransmission)
				if err := _RecoveryBeacon.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_RecoveryBeacon *RecoveryBeaconFilterer) ParseNewTransmission(log types.Log) (*RecoveryBeaconNewTransmission, error) {
	event := new(RecoveryBeaconNewTransmission)
	if err := _RecoveryBeacon.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconOwnershipTransferRequestedIterator struct {
	Event *RecoveryBeaconOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconOwnershipTransferRequested)
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
		it.Event = new(RecoveryBeaconOwnershipTransferRequested)
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

func (it *RecoveryBeaconOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RecoveryBeacon *RecoveryBeaconFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RecoveryBeaconOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeacon.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOwnershipTransferRequestedIterator{contract: _RecoveryBeacon.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RecoveryBeacon *RecoveryBeaconFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeacon.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconOwnershipTransferRequested)
				if err := _RecoveryBeacon.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RecoveryBeacon *RecoveryBeaconFilterer) ParseOwnershipTransferRequested(log types.Log) (*RecoveryBeaconOwnershipTransferRequested, error) {
	event := new(RecoveryBeaconOwnershipTransferRequested)
	if err := _RecoveryBeacon.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconOwnershipTransferredIterator struct {
	Event *RecoveryBeaconOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconOwnershipTransferred)
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
		it.Event = new(RecoveryBeaconOwnershipTransferred)
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

func (it *RecoveryBeaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RecoveryBeacon *RecoveryBeaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RecoveryBeaconOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeacon.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOwnershipTransferredIterator{contract: _RecoveryBeacon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RecoveryBeacon *RecoveryBeaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeacon.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconOwnershipTransferred)
				if err := _RecoveryBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RecoveryBeacon *RecoveryBeaconFilterer) ParseOwnershipTransferred(log types.Log) (*RecoveryBeaconOwnershipTransferred, error) {
	event := new(RecoveryBeaconOwnershipTransferred)
	if err := _RecoveryBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var RecoveryBeaconDKGClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDKG\",\"name\":\"_keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161045338038061045383398101604081905261002f91610058565b600080546001600160a01b0319166001600160a01b039390931692909217909155600155610092565b6000806040838503121561006b57600080fd5b82516001600160a01b038116811461008257600080fd5b6020939093015192949293505050565b6103b2806100a16000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806355e4874914610051578063bf2732c71461005b578063cc31f7dd1461006e578063d57fc45a14610089575b600080fd5b610059610092565b005b610059610069366004610287565b6100f6565b61007760015481565b60405190815260200160405180910390f35b61007760025481565b60005473ffffffffffffffffffffffffffffffffffffffff163381146100ee5760405163292f4fb560e01b815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044015b60405180910390fd5b506000600255565b60005473ffffffffffffffffffffffffffffffffffffffff1633811461014d5760405163292f4fb560e01b815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044016100e5565b815160405161015f919060200161036a565b60408051601f1981840301815291905280516020909101206002555050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156101d0576101d061017e565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156101ff576101ff61017e565b604052919050565b600082601f83011261021857600080fd5b8135602067ffffffffffffffff8211156102345761023461017e565b8160051b6102438282016101d6565b928352848101820192828101908785111561025d57600080fd5b83870192505b8483101561027c57823582529183019190830190610263565b979650505050505050565b6000602080838503121561029a57600080fd5b823567ffffffffffffffff808211156102b257600080fd5b90840190604082870312156102c657600080fd5b6102ce6101ad565b8235828111156102dd57600080fd5b8301601f810188136102ee57600080fd5b8035838111156103005761030061017e565b610312601f8201601f191687016101d6565b818152898783850101111561032657600080fd5b81878401888301376000878383010152808452505050838301358281111561034d57600080fd5b61035988828601610207565b948201949094529695505050505050565b6000825160005b8181101561038b5760208186018101518583015201610371565b8181111561039a576000828501525b50919091019291505056fea164736f6c634300080f000a",
}

var RecoveryBeaconDKGClientABI = RecoveryBeaconDKGClientMetaData.ABI

var RecoveryBeaconDKGClientBin = RecoveryBeaconDKGClientMetaData.Bin

func DeployRecoveryBeaconDKGClient(auth *bind.TransactOpts, backend bind.ContractBackend, _keyProvider common.Address, _keyID [32]byte) (common.Address, *types.Transaction, *RecoveryBeaconDKGClient, error) {
	parsed, err := RecoveryBeaconDKGClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecoveryBeaconDKGClientBin), backend, _keyProvider, _keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecoveryBeaconDKGClient{RecoveryBeaconDKGClientCaller: RecoveryBeaconDKGClientCaller{contract: contract}, RecoveryBeaconDKGClientTransactor: RecoveryBeaconDKGClientTransactor{contract: contract}, RecoveryBeaconDKGClientFilterer: RecoveryBeaconDKGClientFilterer{contract: contract}}, nil
}

type RecoveryBeaconDKGClient struct {
	RecoveryBeaconDKGClientCaller
	RecoveryBeaconDKGClientTransactor
	RecoveryBeaconDKGClientFilterer
}

type RecoveryBeaconDKGClientCaller struct {
	contract *bind.BoundContract
}

type RecoveryBeaconDKGClientTransactor struct {
	contract *bind.BoundContract
}

type RecoveryBeaconDKGClientFilterer struct {
	contract *bind.BoundContract
}

type RecoveryBeaconDKGClientSession struct {
	Contract     *RecoveryBeaconDKGClient
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconDKGClientCallerSession struct {
	Contract *RecoveryBeaconDKGClientCaller
	CallOpts bind.CallOpts
}

type RecoveryBeaconDKGClientTransactorSession struct {
	Contract     *RecoveryBeaconDKGClientTransactor
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconDKGClientRaw struct {
	Contract *RecoveryBeaconDKGClient
}

type RecoveryBeaconDKGClientCallerRaw struct {
	Contract *RecoveryBeaconDKGClientCaller
}

type RecoveryBeaconDKGClientTransactorRaw struct {
	Contract *RecoveryBeaconDKGClientTransactor
}

func NewRecoveryBeaconDKGClient(address common.Address, backend bind.ContractBackend) (*RecoveryBeaconDKGClient, error) {
	contract, err := bindRecoveryBeaconDKGClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconDKGClient{RecoveryBeaconDKGClientCaller: RecoveryBeaconDKGClientCaller{contract: contract}, RecoveryBeaconDKGClientTransactor: RecoveryBeaconDKGClientTransactor{contract: contract}, RecoveryBeaconDKGClientFilterer: RecoveryBeaconDKGClientFilterer{contract: contract}}, nil
}

func NewRecoveryBeaconDKGClientCaller(address common.Address, caller bind.ContractCaller) (*RecoveryBeaconDKGClientCaller, error) {
	contract, err := bindRecoveryBeaconDKGClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconDKGClientCaller{contract: contract}, nil
}

func NewRecoveryBeaconDKGClientTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoveryBeaconDKGClientTransactor, error) {
	contract, err := bindRecoveryBeaconDKGClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconDKGClientTransactor{contract: contract}, nil
}

func NewRecoveryBeaconDKGClientFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoveryBeaconDKGClientFilterer, error) {
	contract, err := bindRecoveryBeaconDKGClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconDKGClientFilterer{contract: contract}, nil
}

func bindRecoveryBeaconDKGClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoveryBeaconDKGClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconDKGClient.Contract.RecoveryBeaconDKGClientCaller.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.RecoveryBeaconDKGClientTransactor.contract.Transfer(opts)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.RecoveryBeaconDKGClientTransactor.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconDKGClient.Contract.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.contract.Transfer(opts)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RecoveryBeaconDKGClient.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientSession) SKeyID() ([32]byte, error) {
	return _RecoveryBeaconDKGClient.Contract.SKeyID(&_RecoveryBeaconDKGClient.CallOpts)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientCallerSession) SKeyID() ([32]byte, error) {
	return _RecoveryBeaconDKGClient.Contract.SKeyID(&_RecoveryBeaconDKGClient.CallOpts)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RecoveryBeaconDKGClient.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientSession) SProvingKeyHash() ([32]byte, error) {
	return _RecoveryBeaconDKGClient.Contract.SProvingKeyHash(&_RecoveryBeaconDKGClient.CallOpts)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _RecoveryBeaconDKGClient.Contract.SProvingKeyHash(&_RecoveryBeaconDKGClient.CallOpts)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.contract.Transact(opts, "keyGenerated", kd)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.KeyGenerated(&_RecoveryBeaconDKGClient.TransactOpts, kd)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.KeyGenerated(&_RecoveryBeaconDKGClient.TransactOpts, kd)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.contract.Transact(opts, "newKeyRequested")
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientSession) NewKeyRequested() (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.NewKeyRequested(&_RecoveryBeaconDKGClient.TransactOpts)
}

func (_RecoveryBeaconDKGClient *RecoveryBeaconDKGClientTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _RecoveryBeaconDKGClient.Contract.NewKeyRequested(&_RecoveryBeaconDKGClient.TransactOpts)
}

var RecoveryBeaconOCRMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OffchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"occVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structRecoveryBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_mostRecentAccountToRecover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600880546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610149565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600980546001600160a01b0319166001600160a01b03838116918217909255600854604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b611a7a80620001596000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063acf2a15611610076578063b1dc65a41161005b578063b1dc65a4146101b5578063e3d0e712146101c8578063f2fde38b146101db57600080fd5b8063acf2a15614610178578063afcb95d71461018b57600080fd5b806381ff7048116100a757806381ff704814610115578063878a1d88146101425780638da5cb5b1461015357600080fd5b8063181f5a77146100c357806379ba50971461010b575b600080fd5b604080518082018252601a81527f5265636f76657279426561636f6e20312e302e302d616c706861000000000000602082015290516101029190611452565b60405180910390f35b6101136101ee565b005b600c54600e54604080516000815264010000000090930463ffffffff166020840152820152606001610102565b61011361015036600461146c565b50565b6008546001600160a01b03165b6040516001600160a01b039091168152602001610102565b600454610160906001600160a01b031681565b600e546010546040805160008152602081019390935263ffffffff90911690820152606001610102565b6101136101c3366004611512565b6102a8565b6101136101d63660046115f7565b61077c565b6101136101e93660046116fa565b610ebb565b6009546001600160a01b0316331461024d5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600880546001600160a01b0319808216339081179093556009805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b604080516101008082018352600b5460ff808216845291810464ffffffffff166020808501919091526601000000000000820463ffffffff908116858701526a01000000000000000000008304811660608601526e01000000000000000000000000000083048116608086015272010000000000000000000000000000000000008304811660a086015276010000000000000000000000000000000000000000000083041660c08501527a01000000000000000000000000000000000000000000000000000090910462ffffff1660e084015233600090815260118252939093205491928b013591166103dd5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401610244565b600e548a351461042f5760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401610244565b61043d898989898989610ecc565b815161044a906001611743565b60ff16861461049b5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401610244565b8584146104ea5760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401610244565b600089896040516104fc929190611768565b604051908190038120610513918d90602001611778565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b898110156106b957600060018589846020811061056257610562611717565b61056f91901a601b611743565b8e8e8681811061058157610581611717565b905060200201358d8d8781811061059a5761059a611717565b90506020020135604051600081526020016040526040516105d7949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156105f9573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526012602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506106925760405162461bcd60e51b815260206004820152600f60248201527f7369676e6174757265206572726f7200000000000000000000000000000000006044820152606401610244565b826020015160080260ff166001901b840193505080806106b190611794565b915050610543565b5081827e01010101010101010101010101010101010101010101010101010101010101161461072a5760405162461bcd60e51b815260206004820152601060248201527f6475706c6963617465207369676e6572000000000000000000000000000000006044820152606401610244565b5050604080516020601f8c018190048102820181019092528a815261077092508c35918491908d908d9081908401838280828437600092019190915250610f6992505050565b50505050505050505050565b6107846110af565b601f8911156107d55760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79206f7261636c6573000000000000000000000000000000006044820152606401610244565b8887146108245760405162461bcd60e51b815260206004820152601660248201527f6f7261636c65206c656e677468206d69736d61746368000000000000000000006044820152606401610244565b886108308760036117ad565b60ff16106108805760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610244565b61088c8660ff1661110b565b6040805160e060208c02808301820190935260c082018c815260009383928f918f918291908601908490808284376000920191909152505050908252506040805160208c810282810182019093528c82529283019290918d918d91829185019084908082843760009201919091525050509082525060ff891660208083019190915260408051601f8a01839004830281018301825289815292019190899089908190840183828082843760009201919091525050509082525067ffffffffffffffff861660208083019190915260408051601f870183900483028101830182528681529201919086908690819084018382808284376000920182905250939094525050600b805465ffffffffff001916905560135492935090505b81811015610a52576000601382815481106109c4576109c4611717565b6000918252602082200154601480546001600160a01b03909216935090849081106109f1576109f1611717565b60009182526020808320909101546001600160a01b039485168352601282526040808420805461ffff1916905594168252601190529190912080546dffffffffffffffffffffffffffff191690555080610a4a81611794565b9150506109a7565b50610a5f60136000611292565b610a6b60146000611292565b60005b825151811015610cf3576012600084600001518381518110610a9257610a92611717565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615610b065760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610244565b604080518082019091526001815260ff821660208201528351805160129160009185908110610b3757610b37611717565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051601192919084908110610ba257610ba2611717565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615610c165760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610244565b60405180606001604052806001151581526020018260ff16815260200160006bffffffffffffffffffffffff168152506011600085602001518481518110610c6057610c60611717565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516bffffffffffffffffffffffff1662010000026dffffffffffffffffffffffff00001960ff959095166101000261ff00199315159390931661ffff1990941693909317919091179290921617905580610ceb81611794565b915050610a6e565b5081518051610d0a916013916020909101906112b0565b506020808301518051610d219260149201906112b0565b506040820151600b805460ff191660ff909216919091179055600c805467ffffffff0000000019811664010000000063ffffffff43811682029283179094558204831692600092610d799290821691161760016117d6565b905080600c60006101000a81548163ffffffff021916908363ffffffff1602179055506000610dcd46308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a0015161115b565b905080600e600001819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05838284886000015189602001518a604001518b606001518c608001518d60a00151604051610e3099989796959493929190611842565b60405180910390a1600b546601000000000000900463ffffffff1660005b865151811015610ea85781601582601f8110610e6c57610e6c611717565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff1602179055508080610ea090611794565b915050610e4e565b5050505050505050505050505050505050565b610ec36110af565b610150816111e8565b6000610ed98260206118d8565b610ee48560206118d8565b610ef0886101446118f7565b610efa91906118f7565b610f0491906118f7565b610f0f9060006118f7565b9050368114610f605760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401610244565b50505050505050565b600254600003610f7857505050565b600081806020019051810190610f8e919061190f565b9050806040015115610fd45780516020808301516001600160a01b039081166000908152600a909252604090912080546001600160a01b03191691909216179055611039565b80516001600160a01b031660009081526003602052604090206001810154600160a01b900460ff161561103757600181018054600160a01b900460ff1690601461101d83611994565b91906101000a81548160ff021916908360ff160217905550505b505b61104560026000611315565b8051602080830151604080513381526001600160a01b039485169381019390935292168183015260608101869052905164ffffffffff8516917ffb8260fafd97df85b3771f6532a09a8af1daeac85d6938ee302237174c91a9c9919081900360800190a250505050565b6008546001600160a01b031633146111095760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610244565b565b806000106101505760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610244565b6000808a8a8a8a8a8a8a8a8a60405160200161117f999897969594939291906119b1565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036112405760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610244565b600980546001600160a01b0319166001600160a01b03838116918217909255600854604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b50805460008255906000526020600020908101906101509190611336565b828054828255906000526020600020908101928215611305579160200282015b8281111561130557825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906112d0565b50611311929150611336565b5090565b5080546000825560080290600052602060002090810190610150919061134b565b5b808211156113115760008155600101611337565b8082111561131157805460ff19168155600061136a60018301826113cb565b6113786002830160006113cb565b6113866003830160006113cb565b6113946004830160006113cb565b6113a26005830160006113cb565b506006810180546001600160a01b0319908116909155600782018054909116905560080161134b565b5080546113d790611a39565b6000825580601f106113e7575050565b601f0160209004906000526020600020908101906101509190611336565b6000815180845260005b8181101561142b5760208185018101518683018201520161140f565b8181111561143d576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006114656020830184611405565b9392505050565b60006060828403121561147e57600080fd5b50919050565b60008083601f84011261149657600080fd5b50813567ffffffffffffffff8111156114ae57600080fd5b6020830191508360208285010111156114c657600080fd5b9250929050565b60008083601f8401126114df57600080fd5b50813567ffffffffffffffff8111156114f757600080fd5b6020830191508360208260051b85010111156114c657600080fd5b60008060008060008060008060e0898b03121561152e57600080fd5b606089018a81111561153f57600080fd5b8998503567ffffffffffffffff8082111561155957600080fd5b6115658c838d01611484565b909950975060808b013591508082111561157e57600080fd5b61158a8c838d016114cd565b909750955060a08b01359150808211156115a357600080fd5b506115b08b828c016114cd565b999c989b50969995989497949560c00135949350505050565b803560ff811681146115da57600080fd5b919050565b803567ffffffffffffffff811681146115da57600080fd5b60008060008060008060008060008060c08b8d03121561161657600080fd5b8a3567ffffffffffffffff8082111561162e57600080fd5b61163a8e838f016114cd565b909c509a5060208d013591508082111561165357600080fd5b61165f8e838f016114cd565b909a50985088915061167360408e016115c9565b975060608d013591508082111561168957600080fd5b6116958e838f01611484565b90975095508591506116a960808e016115df565b945060a08d01359150808211156116bf57600080fd5b506116cc8d828e01611484565b915080935050809150509295989b9194979a5092959850565b6001600160a01b038116811461015057600080fd5b60006020828403121561170c57600080fd5b8135611465816116e5565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff038211156117605761176061172d565b019392505050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b6000600182016117a6576117a661172d565b5060010190565b600060ff821660ff84168160ff04811182151516156117ce576117ce61172d565b029392505050565b600063ffffffff8083168185168083038211156117f5576117f561172d565b01949350505050565b600081518084526020808501945080840160005b838110156118375781516001600160a01b031687529582019590820190600101611812565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526118728184018a6117fe565b9050828103608084015261188681896117fe565b905060ff871660a084015282810360c08401526118a38187611405565b905067ffffffffffffffff851660e08401528281036101008401526118c88185611405565b9c9b505050505050505050505050565b60008160001904831182151516156118f2576118f261172d565b500290565b6000821982111561190a5761190a61172d565b500190565b60006060828403121561192157600080fd5b6040516060810181811067ffffffffffffffff8211171561195257634e487b7160e01b600052604160045260246000fd5b6040528251611960816116e5565b81526020830151611970816116e5565b60208201526040830151801515811461198857600080fd5b60408201529392505050565b600060ff8216806119a7576119a761172d565b6000190192915050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526119eb8285018b6117fe565b915083820360808501526119ff828a6117fe565b915060ff881660a085015283820360c0850152611a1c8288611405565b90861660e085015283810361010085015290506118c88185611405565b600181811c90821680611a4d57607f821691505b60208210810361147e57634e487b7160e01b600052602260045260246000fdfea164736f6c634300080f000a",
}

var RecoveryBeaconOCRABI = RecoveryBeaconOCRMetaData.ABI

var RecoveryBeaconOCRBin = RecoveryBeaconOCRMetaData.Bin

func DeployRecoveryBeaconOCR(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RecoveryBeaconOCR, error) {
	parsed, err := RecoveryBeaconOCRMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecoveryBeaconOCRBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecoveryBeaconOCR{RecoveryBeaconOCRCaller: RecoveryBeaconOCRCaller{contract: contract}, RecoveryBeaconOCRTransactor: RecoveryBeaconOCRTransactor{contract: contract}, RecoveryBeaconOCRFilterer: RecoveryBeaconOCRFilterer{contract: contract}}, nil
}

type RecoveryBeaconOCR struct {
	RecoveryBeaconOCRCaller
	RecoveryBeaconOCRTransactor
	RecoveryBeaconOCRFilterer
}

type RecoveryBeaconOCRCaller struct {
	contract *bind.BoundContract
}

type RecoveryBeaconOCRTransactor struct {
	contract *bind.BoundContract
}

type RecoveryBeaconOCRFilterer struct {
	contract *bind.BoundContract
}

type RecoveryBeaconOCRSession struct {
	Contract     *RecoveryBeaconOCR
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconOCRCallerSession struct {
	Contract *RecoveryBeaconOCRCaller
	CallOpts bind.CallOpts
}

type RecoveryBeaconOCRTransactorSession struct {
	Contract     *RecoveryBeaconOCRTransactor
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconOCRRaw struct {
	Contract *RecoveryBeaconOCR
}

type RecoveryBeaconOCRCallerRaw struct {
	Contract *RecoveryBeaconOCRCaller
}

type RecoveryBeaconOCRTransactorRaw struct {
	Contract *RecoveryBeaconOCRTransactor
}

func NewRecoveryBeaconOCR(address common.Address, backend bind.ContractBackend) (*RecoveryBeaconOCR, error) {
	contract, err := bindRecoveryBeaconOCR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCR{RecoveryBeaconOCRCaller: RecoveryBeaconOCRCaller{contract: contract}, RecoveryBeaconOCRTransactor: RecoveryBeaconOCRTransactor{contract: contract}, RecoveryBeaconOCRFilterer: RecoveryBeaconOCRFilterer{contract: contract}}, nil
}

func NewRecoveryBeaconOCRCaller(address common.Address, caller bind.ContractCaller) (*RecoveryBeaconOCRCaller, error) {
	contract, err := bindRecoveryBeaconOCR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCRCaller{contract: contract}, nil
}

func NewRecoveryBeaconOCRTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoveryBeaconOCRTransactor, error) {
	contract, err := bindRecoveryBeaconOCR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCRTransactor{contract: contract}, nil
}

func NewRecoveryBeaconOCRFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoveryBeaconOCRFilterer, error) {
	contract, err := bindRecoveryBeaconOCR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCRFilterer{contract: contract}, nil
}

func bindRecoveryBeaconOCR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoveryBeaconOCRABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconOCR.Contract.RecoveryBeaconOCRCaller.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.RecoveryBeaconOCRTransactor.contract.Transfer(opts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.RecoveryBeaconOCRTransactor.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconOCR.Contract.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.contract.Transfer(opts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _RecoveryBeaconOCR.contract.Call(opts, &out, "latestConfigDetails")

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

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _RecoveryBeaconOCR.Contract.LatestConfigDetails(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _RecoveryBeaconOCR.Contract.LatestConfigDetails(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _RecoveryBeaconOCR.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

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

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _RecoveryBeaconOCR.Contract.LatestConfigDigestAndEpoch(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _RecoveryBeaconOCR.Contract.LatestConfigDigestAndEpoch(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeaconOCR.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) Owner() (common.Address, error) {
	return _RecoveryBeaconOCR.Contract.Owner(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCallerSession) Owner() (common.Address, error) {
	return _RecoveryBeaconOCR.Contract.Owner(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCaller) SMostRecentAccountToRecover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeaconOCR.contract.Call(opts, &out, "s_mostRecentAccountToRecover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeaconOCR.Contract.SMostRecentAccountToRecover(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCallerSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeaconOCR.Contract.SMostRecentAccountToRecover(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RecoveryBeaconOCR.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) TypeAndVersion() (string, error) {
	return _RecoveryBeaconOCR.Contract.TypeAndVersion(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRCallerSession) TypeAndVersion() (string, error) {
	return _RecoveryBeaconOCR.Contract.TypeAndVersion(&_RecoveryBeaconOCR.CallOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.contract.Transact(opts, "acceptOwnership")
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) AcceptOwnership() (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.AcceptOwnership(&_RecoveryBeaconOCR.TransactOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.AcceptOwnership(&_RecoveryBeaconOCR.TransactOpts)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactor) ExposeType(opts *bind.TransactOpts, arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.contract.Transact(opts, "exposeType", arg0)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) ExposeType(arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.ExposeType(&_RecoveryBeaconOCR.TransactOpts, arg0)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactorSession) ExposeType(arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.ExposeType(&_RecoveryBeaconOCR.TransactOpts, arg0)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.SetConfig(&_RecoveryBeaconOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.SetConfig(&_RecoveryBeaconOCR.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.contract.Transact(opts, "transferOwnership", to)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.TransferOwnership(&_RecoveryBeaconOCR.TransactOpts, to)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.TransferOwnership(&_RecoveryBeaconOCR.TransactOpts, to)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.Transmit(&_RecoveryBeaconOCR.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _RecoveryBeaconOCR.Contract.Transmit(&_RecoveryBeaconOCR.TransactOpts, reportContext, report, rs, ss, rawVs)
}

type RecoveryBeaconOCRConfigSetIterator struct {
	Event *RecoveryBeaconOCRConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconOCRConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconOCRConfigSet)
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
		it.Event = new(RecoveryBeaconOCRConfigSet)
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

func (it *RecoveryBeaconOCRConfigSetIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconOCRConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconOCRConfigSet struct {
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

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) FilterConfigSet(opts *bind.FilterOpts) (*RecoveryBeaconOCRConfigSetIterator, error) {

	logs, sub, err := _RecoveryBeaconOCR.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCRConfigSetIterator{contract: _RecoveryBeaconOCR.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconOCRConfigSet) (event.Subscription, error) {

	logs, sub, err := _RecoveryBeaconOCR.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconOCRConfigSet)
				if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) ParseConfigSet(log types.Log) (*RecoveryBeaconOCRConfigSet, error) {
	event := new(RecoveryBeaconOCRConfigSet)
	if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconOCRNewTransmissionIterator struct {
	Event *RecoveryBeaconOCRNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconOCRNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconOCRNewTransmission)
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
		it.Event = new(RecoveryBeaconOCRNewTransmission)
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

func (it *RecoveryBeaconOCRNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconOCRNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconOCRNewTransmission struct {
	EpochAndRound    *big.Int
	Transmitter      common.Address
	AccountToRecover common.Address
	Recoverer        common.Address
	ConfigDigest     [32]byte
	Raw              types.Log
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) FilterNewTransmission(opts *bind.FilterOpts, epochAndRound []*big.Int) (*RecoveryBeaconOCRNewTransmissionIterator, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _RecoveryBeaconOCR.contract.FilterLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCRNewTransmissionIterator{contract: _RecoveryBeaconOCR.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconOCRNewTransmission, epochAndRound []*big.Int) (event.Subscription, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _RecoveryBeaconOCR.contract.WatchLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconOCRNewTransmission)
				if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) ParseNewTransmission(log types.Log) (*RecoveryBeaconOCRNewTransmission, error) {
	event := new(RecoveryBeaconOCRNewTransmission)
	if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconOCROwnershipTransferRequestedIterator struct {
	Event *RecoveryBeaconOCROwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconOCROwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconOCROwnershipTransferRequested)
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
		it.Event = new(RecoveryBeaconOCROwnershipTransferRequested)
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

func (it *RecoveryBeaconOCROwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconOCROwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconOCROwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RecoveryBeaconOCROwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconOCR.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCROwnershipTransferRequestedIterator{contract: _RecoveryBeaconOCR.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconOCROwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconOCR.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconOCROwnershipTransferRequested)
				if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) ParseOwnershipTransferRequested(log types.Log) (*RecoveryBeaconOCROwnershipTransferRequested, error) {
	event := new(RecoveryBeaconOCROwnershipTransferRequested)
	if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconOCROwnershipTransferredIterator struct {
	Event *RecoveryBeaconOCROwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconOCROwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconOCROwnershipTransferred)
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
		it.Event = new(RecoveryBeaconOCROwnershipTransferred)
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

func (it *RecoveryBeaconOCROwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconOCROwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconOCROwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RecoveryBeaconOCROwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconOCR.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconOCROwnershipTransferredIterator{contract: _RecoveryBeaconOCR.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconOCROwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconOCR.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconOCROwnershipTransferred)
				if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RecoveryBeaconOCR *RecoveryBeaconOCRFilterer) ParseOwnershipTransferred(log types.Log) (*RecoveryBeaconOCROwnershipTransferred, error) {
	event := new(RecoveryBeaconOCROwnershipTransferred)
	if err := _RecoveryBeaconOCR.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var RecoveryBeaconReportMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recoverer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structRecoveryBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_mostRecentAccountToRecover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600880546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610149565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600980546001600160a01b0319166001600160a01b03838116918217909255600854604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b610316806101586000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80638da5cb5b116100505780638da5cb5b14610087578063acf2a156146100b0578063f2fde38b146100c357600080fd5b806379ba50971461006c578063878a1d8814610076575b600080fd5b6100746100d6565b005b6100746100843660046102c1565b50565b6008546001600160a01b03165b6040516001600160a01b03909116815260200160405180910390f35b600454610094906001600160a01b031681565b6100746100d13660046102d9565b61019d565b6009546001600160a01b031633146101355760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b6008805473ffffffffffffffffffffffffffffffffffffffff19808216339081179093556009805490911690556040516001600160a01b03909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6101a56101ae565b6100848161020a565b6008546001600160a01b031633146102085760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161012c565b565b336001600160a01b038216036102625760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161012c565b6009805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03838116918217909255600854604051919216907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127890600090a350565b6000606082840312156102d357600080fd5b50919050565b6000602082840312156102eb57600080fd5b81356001600160a01b038116811461030257600080fd5b939250505056fea164736f6c634300080f000a",
}

var RecoveryBeaconReportABI = RecoveryBeaconReportMetaData.ABI

var RecoveryBeaconReportBin = RecoveryBeaconReportMetaData.Bin

func DeployRecoveryBeaconReport(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RecoveryBeaconReport, error) {
	parsed, err := RecoveryBeaconReportMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecoveryBeaconReportBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecoveryBeaconReport{RecoveryBeaconReportCaller: RecoveryBeaconReportCaller{contract: contract}, RecoveryBeaconReportTransactor: RecoveryBeaconReportTransactor{contract: contract}, RecoveryBeaconReportFilterer: RecoveryBeaconReportFilterer{contract: contract}}, nil
}

type RecoveryBeaconReport struct {
	RecoveryBeaconReportCaller
	RecoveryBeaconReportTransactor
	RecoveryBeaconReportFilterer
}

type RecoveryBeaconReportCaller struct {
	contract *bind.BoundContract
}

type RecoveryBeaconReportTransactor struct {
	contract *bind.BoundContract
}

type RecoveryBeaconReportFilterer struct {
	contract *bind.BoundContract
}

type RecoveryBeaconReportSession struct {
	Contract     *RecoveryBeaconReport
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconReportCallerSession struct {
	Contract *RecoveryBeaconReportCaller
	CallOpts bind.CallOpts
}

type RecoveryBeaconReportTransactorSession struct {
	Contract     *RecoveryBeaconReportTransactor
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconReportRaw struct {
	Contract *RecoveryBeaconReport
}

type RecoveryBeaconReportCallerRaw struct {
	Contract *RecoveryBeaconReportCaller
}

type RecoveryBeaconReportTransactorRaw struct {
	Contract *RecoveryBeaconReportTransactor
}

func NewRecoveryBeaconReport(address common.Address, backend bind.ContractBackend) (*RecoveryBeaconReport, error) {
	contract, err := bindRecoveryBeaconReport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconReport{RecoveryBeaconReportCaller: RecoveryBeaconReportCaller{contract: contract}, RecoveryBeaconReportTransactor: RecoveryBeaconReportTransactor{contract: contract}, RecoveryBeaconReportFilterer: RecoveryBeaconReportFilterer{contract: contract}}, nil
}

func NewRecoveryBeaconReportCaller(address common.Address, caller bind.ContractCaller) (*RecoveryBeaconReportCaller, error) {
	contract, err := bindRecoveryBeaconReport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconReportCaller{contract: contract}, nil
}

func NewRecoveryBeaconReportTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoveryBeaconReportTransactor, error) {
	contract, err := bindRecoveryBeaconReport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconReportTransactor{contract: contract}, nil
}

func NewRecoveryBeaconReportFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoveryBeaconReportFilterer, error) {
	contract, err := bindRecoveryBeaconReport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconReportFilterer{contract: contract}, nil
}

func bindRecoveryBeaconReport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoveryBeaconReportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_RecoveryBeaconReport *RecoveryBeaconReportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconReport.Contract.RecoveryBeaconReportCaller.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.RecoveryBeaconReportTransactor.contract.Transfer(opts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.RecoveryBeaconReportTransactor.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconReport.Contract.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.contract.Transfer(opts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeaconReport.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeaconReport *RecoveryBeaconReportSession) Owner() (common.Address, error) {
	return _RecoveryBeaconReport.Contract.Owner(&_RecoveryBeaconReport.CallOpts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportCallerSession) Owner() (common.Address, error) {
	return _RecoveryBeaconReport.Contract.Owner(&_RecoveryBeaconReport.CallOpts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportCaller) SMostRecentAccountToRecover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeaconReport.contract.Call(opts, &out, "s_mostRecentAccountToRecover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeaconReport *RecoveryBeaconReportSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeaconReport.Contract.SMostRecentAccountToRecover(&_RecoveryBeaconReport.CallOpts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportCallerSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeaconReport.Contract.SMostRecentAccountToRecover(&_RecoveryBeaconReport.CallOpts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconReport.contract.Transact(opts, "acceptOwnership")
}

func (_RecoveryBeaconReport *RecoveryBeaconReportSession) AcceptOwnership() (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.AcceptOwnership(&_RecoveryBeaconReport.TransactOpts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.AcceptOwnership(&_RecoveryBeaconReport.TransactOpts)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactor) ExposeType(opts *bind.TransactOpts, arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeaconReport.contract.Transact(opts, "exposeType", arg0)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportSession) ExposeType(arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.ExposeType(&_RecoveryBeaconReport.TransactOpts, arg0)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactorSession) ExposeType(arg0 RecoveryBeaconReportReport) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.ExposeType(&_RecoveryBeaconReport.TransactOpts, arg0)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RecoveryBeaconReport.contract.Transact(opts, "transferOwnership", to)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.TransferOwnership(&_RecoveryBeaconReport.TransactOpts, to)
}

func (_RecoveryBeaconReport *RecoveryBeaconReportTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RecoveryBeaconReport.Contract.TransferOwnership(&_RecoveryBeaconReport.TransactOpts, to)
}

type RecoveryBeaconReportNewTransmissionIterator struct {
	Event *RecoveryBeaconReportNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconReportNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconReportNewTransmission)
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
		it.Event = new(RecoveryBeaconReportNewTransmission)
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

func (it *RecoveryBeaconReportNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconReportNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconReportNewTransmission struct {
	EpochAndRound    *big.Int
	Transmitter      common.Address
	AccountToRecover common.Address
	Recoverer        common.Address
	ConfigDigest     [32]byte
	Raw              types.Log
}

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) FilterNewTransmission(opts *bind.FilterOpts, epochAndRound []*big.Int) (*RecoveryBeaconReportNewTransmissionIterator, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _RecoveryBeaconReport.contract.FilterLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconReportNewTransmissionIterator{contract: _RecoveryBeaconReport.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconReportNewTransmission, epochAndRound []*big.Int) (event.Subscription, error) {

	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _RecoveryBeaconReport.contract.WatchLogs(opts, "NewTransmission", epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconReportNewTransmission)
				if err := _RecoveryBeaconReport.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) ParseNewTransmission(log types.Log) (*RecoveryBeaconReportNewTransmission, error) {
	event := new(RecoveryBeaconReportNewTransmission)
	if err := _RecoveryBeaconReport.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconReportOwnershipTransferRequestedIterator struct {
	Event *RecoveryBeaconReportOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconReportOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconReportOwnershipTransferRequested)
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
		it.Event = new(RecoveryBeaconReportOwnershipTransferRequested)
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

func (it *RecoveryBeaconReportOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconReportOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconReportOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RecoveryBeaconReportOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconReport.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconReportOwnershipTransferRequestedIterator{contract: _RecoveryBeaconReport.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconReportOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconReport.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconReportOwnershipTransferRequested)
				if err := _RecoveryBeaconReport.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) ParseOwnershipTransferRequested(log types.Log) (*RecoveryBeaconReportOwnershipTransferRequested, error) {
	event := new(RecoveryBeaconReportOwnershipTransferRequested)
	if err := _RecoveryBeaconReport.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RecoveryBeaconReportOwnershipTransferredIterator struct {
	Event *RecoveryBeaconReportOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RecoveryBeaconReportOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryBeaconReportOwnershipTransferred)
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
		it.Event = new(RecoveryBeaconReportOwnershipTransferred)
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

func (it *RecoveryBeaconReportOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RecoveryBeaconReportOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RecoveryBeaconReportOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RecoveryBeaconReportOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconReport.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconReportOwnershipTransferredIterator{contract: _RecoveryBeaconReport.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RecoveryBeaconReportOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RecoveryBeaconReport.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RecoveryBeaconReportOwnershipTransferred)
				if err := _RecoveryBeaconReport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RecoveryBeaconReport *RecoveryBeaconReportFilterer) ParseOwnershipTransferred(log types.Log) (*RecoveryBeaconReportOwnershipTransferred, error) {
	event := new(RecoveryBeaconReportOwnershipTransferred)
	if err := _RecoveryBeaconReport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var RecoveryBeaconTypesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"s_mostRecentAccountToRecover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5060828061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063acf2a15614602d575b600080fd5b600454604c9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f3fea164736f6c634300080f000a",
}

var RecoveryBeaconTypesABI = RecoveryBeaconTypesMetaData.ABI

var RecoveryBeaconTypesBin = RecoveryBeaconTypesMetaData.Bin

func DeployRecoveryBeaconTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RecoveryBeaconTypes, error) {
	parsed, err := RecoveryBeaconTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecoveryBeaconTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecoveryBeaconTypes{RecoveryBeaconTypesCaller: RecoveryBeaconTypesCaller{contract: contract}, RecoveryBeaconTypesTransactor: RecoveryBeaconTypesTransactor{contract: contract}, RecoveryBeaconTypesFilterer: RecoveryBeaconTypesFilterer{contract: contract}}, nil
}

type RecoveryBeaconTypes struct {
	RecoveryBeaconTypesCaller
	RecoveryBeaconTypesTransactor
	RecoveryBeaconTypesFilterer
}

type RecoveryBeaconTypesCaller struct {
	contract *bind.BoundContract
}

type RecoveryBeaconTypesTransactor struct {
	contract *bind.BoundContract
}

type RecoveryBeaconTypesFilterer struct {
	contract *bind.BoundContract
}

type RecoveryBeaconTypesSession struct {
	Contract     *RecoveryBeaconTypes
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconTypesCallerSession struct {
	Contract *RecoveryBeaconTypesCaller
	CallOpts bind.CallOpts
}

type RecoveryBeaconTypesTransactorSession struct {
	Contract     *RecoveryBeaconTypesTransactor
	TransactOpts bind.TransactOpts
}

type RecoveryBeaconTypesRaw struct {
	Contract *RecoveryBeaconTypes
}

type RecoveryBeaconTypesCallerRaw struct {
	Contract *RecoveryBeaconTypesCaller
}

type RecoveryBeaconTypesTransactorRaw struct {
	Contract *RecoveryBeaconTypesTransactor
}

func NewRecoveryBeaconTypes(address common.Address, backend bind.ContractBackend) (*RecoveryBeaconTypes, error) {
	contract, err := bindRecoveryBeaconTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconTypes{RecoveryBeaconTypesCaller: RecoveryBeaconTypesCaller{contract: contract}, RecoveryBeaconTypesTransactor: RecoveryBeaconTypesTransactor{contract: contract}, RecoveryBeaconTypesFilterer: RecoveryBeaconTypesFilterer{contract: contract}}, nil
}

func NewRecoveryBeaconTypesCaller(address common.Address, caller bind.ContractCaller) (*RecoveryBeaconTypesCaller, error) {
	contract, err := bindRecoveryBeaconTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconTypesCaller{contract: contract}, nil
}

func NewRecoveryBeaconTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoveryBeaconTypesTransactor, error) {
	contract, err := bindRecoveryBeaconTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconTypesTransactor{contract: contract}, nil
}

func NewRecoveryBeaconTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoveryBeaconTypesFilterer, error) {
	contract, err := bindRecoveryBeaconTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoveryBeaconTypesFilterer{contract: contract}, nil
}

func bindRecoveryBeaconTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoveryBeaconTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconTypes.Contract.RecoveryBeaconTypesCaller.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconTypes.Contract.RecoveryBeaconTypesTransactor.contract.Transfer(opts)
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconTypes.Contract.RecoveryBeaconTypesTransactor.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryBeaconTypes.Contract.contract.Call(opts, result, method, params...)
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryBeaconTypes.Contract.contract.Transfer(opts)
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryBeaconTypes.Contract.contract.Transact(opts, method, params...)
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesCaller) SMostRecentAccountToRecover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecoveryBeaconTypes.contract.Call(opts, &out, "s_mostRecentAccountToRecover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeaconTypes.Contract.SMostRecentAccountToRecover(&_RecoveryBeaconTypes.CallOpts)
}

func (_RecoveryBeaconTypes *RecoveryBeaconTypesCallerSession) SMostRecentAccountToRecover() (common.Address, error) {
	return _RecoveryBeaconTypes.Contract.SMostRecentAccountToRecover(&_RecoveryBeaconTypes.CallOpts)
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

package mockv3aggregator

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

var AggregatorInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

var AggregatorInterfaceABI = AggregatorInterfaceMetaData.ABI

type AggregatorInterface struct {
	AggregatorInterfaceCaller
	AggregatorInterfaceTransactor
	AggregatorInterfaceFilterer
}

type AggregatorInterfaceCaller struct {
	contract *bind.BoundContract
}

type AggregatorInterfaceTransactor struct {
	contract *bind.BoundContract
}

type AggregatorInterfaceFilterer struct {
	contract *bind.BoundContract
}

type AggregatorInterfaceSession struct {
	Contract     *AggregatorInterface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AggregatorInterfaceCallerSession struct {
	Contract *AggregatorInterfaceCaller
	CallOpts bind.CallOpts
}

type AggregatorInterfaceTransactorSession struct {
	Contract     *AggregatorInterfaceTransactor
	TransactOpts bind.TransactOpts
}

type AggregatorInterfaceRaw struct {
	Contract *AggregatorInterface
}

type AggregatorInterfaceCallerRaw struct {
	Contract *AggregatorInterfaceCaller
}

type AggregatorInterfaceTransactorRaw struct {
	Contract *AggregatorInterfaceTransactor
}

func NewAggregatorInterface(address common.Address, backend bind.ContractBackend) (*AggregatorInterface, error) {
	contract, err := bindAggregatorInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorInterface{AggregatorInterfaceCaller: AggregatorInterfaceCaller{contract: contract}, AggregatorInterfaceTransactor: AggregatorInterfaceTransactor{contract: contract}, AggregatorInterfaceFilterer: AggregatorInterfaceFilterer{contract: contract}}, nil
}

func NewAggregatorInterfaceCaller(address common.Address, caller bind.ContractCaller) (*AggregatorInterfaceCaller, error) {
	contract, err := bindAggregatorInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorInterfaceCaller{contract: contract}, nil
}

func NewAggregatorInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorInterfaceTransactor, error) {
	contract, err := bindAggregatorInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorInterfaceTransactor{contract: contract}, nil
}

func NewAggregatorInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorInterfaceFilterer, error) {
	contract, err := bindAggregatorInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorInterfaceFilterer{contract: contract}, nil
}

func bindAggregatorInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AggregatorInterface *AggregatorInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorInterface.Contract.AggregatorInterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_AggregatorInterface *AggregatorInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorInterface.Contract.AggregatorInterfaceTransactor.contract.Transfer(opts)
}

func (_AggregatorInterface *AggregatorInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorInterface.Contract.AggregatorInterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_AggregatorInterface *AggregatorInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorInterface.Contract.contract.Call(opts, result, method, params...)
}

func (_AggregatorInterface *AggregatorInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorInterface.Contract.contract.Transfer(opts)
}

func (_AggregatorInterface *AggregatorInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorInterface.Contract.contract.Transact(opts, method, params...)
}

func (_AggregatorInterface *AggregatorInterfaceCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorInterface.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorInterface *AggregatorInterfaceSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _AggregatorInterface.Contract.GetAnswer(&_AggregatorInterface.CallOpts, roundId)
}

func (_AggregatorInterface *AggregatorInterfaceCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _AggregatorInterface.Contract.GetAnswer(&_AggregatorInterface.CallOpts, roundId)
}

func (_AggregatorInterface *AggregatorInterfaceCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorInterface.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorInterface *AggregatorInterfaceSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _AggregatorInterface.Contract.GetTimestamp(&_AggregatorInterface.CallOpts, roundId)
}

func (_AggregatorInterface *AggregatorInterfaceCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _AggregatorInterface.Contract.GetTimestamp(&_AggregatorInterface.CallOpts, roundId)
}

func (_AggregatorInterface *AggregatorInterfaceCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorInterface.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorInterface *AggregatorInterfaceSession) LatestAnswer() (*big.Int, error) {
	return _AggregatorInterface.Contract.LatestAnswer(&_AggregatorInterface.CallOpts)
}

func (_AggregatorInterface *AggregatorInterfaceCallerSession) LatestAnswer() (*big.Int, error) {
	return _AggregatorInterface.Contract.LatestAnswer(&_AggregatorInterface.CallOpts)
}

func (_AggregatorInterface *AggregatorInterfaceCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorInterface.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorInterface *AggregatorInterfaceSession) LatestRound() (*big.Int, error) {
	return _AggregatorInterface.Contract.LatestRound(&_AggregatorInterface.CallOpts)
}

func (_AggregatorInterface *AggregatorInterfaceCallerSession) LatestRound() (*big.Int, error) {
	return _AggregatorInterface.Contract.LatestRound(&_AggregatorInterface.CallOpts)
}

func (_AggregatorInterface *AggregatorInterfaceCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorInterface.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorInterface *AggregatorInterfaceSession) LatestTimestamp() (*big.Int, error) {
	return _AggregatorInterface.Contract.LatestTimestamp(&_AggregatorInterface.CallOpts)
}

func (_AggregatorInterface *AggregatorInterfaceCallerSession) LatestTimestamp() (*big.Int, error) {
	return _AggregatorInterface.Contract.LatestTimestamp(&_AggregatorInterface.CallOpts)
}

type AggregatorInterfaceAnswerUpdatedIterator struct {
	Event *AggregatorInterfaceAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AggregatorInterfaceAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorInterfaceAnswerUpdated)
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
		it.Event = new(AggregatorInterfaceAnswerUpdated)
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

func (it *AggregatorInterfaceAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *AggregatorInterfaceAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AggregatorInterfaceAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_AggregatorInterface *AggregatorInterfaceFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*AggregatorInterfaceAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AggregatorInterface.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorInterfaceAnswerUpdatedIterator{contract: _AggregatorInterface.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_AggregatorInterface *AggregatorInterfaceFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorInterfaceAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AggregatorInterface.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AggregatorInterfaceAnswerUpdated)
				if err := _AggregatorInterface.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_AggregatorInterface *AggregatorInterfaceFilterer) ParseAnswerUpdated(log types.Log) (*AggregatorInterfaceAnswerUpdated, error) {
	event := new(AggregatorInterfaceAnswerUpdated)
	if err := _AggregatorInterface.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AggregatorInterfaceNewRoundIterator struct {
	Event *AggregatorInterfaceNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AggregatorInterfaceNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorInterfaceNewRound)
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
		it.Event = new(AggregatorInterfaceNewRound)
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

func (it *AggregatorInterfaceNewRoundIterator) Error() error {
	return it.fail
}

func (it *AggregatorInterfaceNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AggregatorInterfaceNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_AggregatorInterface *AggregatorInterfaceFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AggregatorInterfaceNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AggregatorInterface.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorInterfaceNewRoundIterator{contract: _AggregatorInterface.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_AggregatorInterface *AggregatorInterfaceFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *AggregatorInterfaceNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AggregatorInterface.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AggregatorInterfaceNewRound)
				if err := _AggregatorInterface.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_AggregatorInterface *AggregatorInterfaceFilterer) ParseNewRound(log types.Log) (*AggregatorInterfaceNewRound, error) {
	event := new(AggregatorInterfaceNewRound)
	if err := _AggregatorInterface.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var AggregatorV2V3InterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

var AggregatorV2V3InterfaceABI = AggregatorV2V3InterfaceMetaData.ABI

type AggregatorV2V3Interface struct {
	AggregatorV2V3InterfaceCaller
	AggregatorV2V3InterfaceTransactor
	AggregatorV2V3InterfaceFilterer
}

type AggregatorV2V3InterfaceCaller struct {
	contract *bind.BoundContract
}

type AggregatorV2V3InterfaceTransactor struct {
	contract *bind.BoundContract
}

type AggregatorV2V3InterfaceFilterer struct {
	contract *bind.BoundContract
}

type AggregatorV2V3InterfaceSession struct {
	Contract     *AggregatorV2V3Interface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AggregatorV2V3InterfaceCallerSession struct {
	Contract *AggregatorV2V3InterfaceCaller
	CallOpts bind.CallOpts
}

type AggregatorV2V3InterfaceTransactorSession struct {
	Contract     *AggregatorV2V3InterfaceTransactor
	TransactOpts bind.TransactOpts
}

type AggregatorV2V3InterfaceRaw struct {
	Contract *AggregatorV2V3Interface
}

type AggregatorV2V3InterfaceCallerRaw struct {
	Contract *AggregatorV2V3InterfaceCaller
}

type AggregatorV2V3InterfaceTransactorRaw struct {
	Contract *AggregatorV2V3InterfaceTransactor
}

func NewAggregatorV2V3Interface(address common.Address, backend bind.ContractBackend) (*AggregatorV2V3Interface, error) {
	contract, err := bindAggregatorV2V3Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorV2V3Interface{AggregatorV2V3InterfaceCaller: AggregatorV2V3InterfaceCaller{contract: contract}, AggregatorV2V3InterfaceTransactor: AggregatorV2V3InterfaceTransactor{contract: contract}, AggregatorV2V3InterfaceFilterer: AggregatorV2V3InterfaceFilterer{contract: contract}}, nil
}

func NewAggregatorV2V3InterfaceCaller(address common.Address, caller bind.ContractCaller) (*AggregatorV2V3InterfaceCaller, error) {
	contract, err := bindAggregatorV2V3Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV2V3InterfaceCaller{contract: contract}, nil
}

func NewAggregatorV2V3InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorV2V3InterfaceTransactor, error) {
	contract, err := bindAggregatorV2V3Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV2V3InterfaceTransactor{contract: contract}, nil
}

func NewAggregatorV2V3InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorV2V3InterfaceFilterer, error) {
	contract, err := bindAggregatorV2V3Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorV2V3InterfaceFilterer{contract: contract}, nil
}

func bindAggregatorV2V3Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorV2V3InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV2V3Interface.Contract.AggregatorV2V3InterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV2V3Interface.Contract.AggregatorV2V3InterfaceTransactor.contract.Transfer(opts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV2V3Interface.Contract.AggregatorV2V3InterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV2V3Interface.Contract.contract.Call(opts, result, method, params...)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV2V3Interface.Contract.contract.Transfer(opts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV2V3Interface.Contract.contract.Transact(opts, method, params...)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) Decimals() (uint8, error) {
	return _AggregatorV2V3Interface.Contract.Decimals(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) Decimals() (uint8, error) {
	return _AggregatorV2V3Interface.Contract.Decimals(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) Description() (string, error) {
	return _AggregatorV2V3Interface.Contract.Description(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) Description() (string, error) {
	return _AggregatorV2V3Interface.Contract.Description(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.GetAnswer(&_AggregatorV2V3Interface.CallOpts, roundId)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.GetAnswer(&_AggregatorV2V3Interface.CallOpts, roundId)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV2V3Interface.Contract.GetRoundData(&_AggregatorV2V3Interface.CallOpts, _roundId)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV2V3Interface.Contract.GetRoundData(&_AggregatorV2V3Interface.CallOpts, _roundId)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.GetTimestamp(&_AggregatorV2V3Interface.CallOpts, roundId)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.GetTimestamp(&_AggregatorV2V3Interface.CallOpts, roundId)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) LatestAnswer() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.LatestAnswer(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) LatestAnswer() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.LatestAnswer(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) LatestRound() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.LatestRound(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) LatestRound() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.LatestRound(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV2V3Interface.Contract.LatestRoundData(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV2V3Interface.Contract.LatestRoundData(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) LatestTimestamp() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.LatestTimestamp(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) LatestTimestamp() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.LatestTimestamp(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV2V3Interface.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceSession) Version() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.Version(&_AggregatorV2V3Interface.CallOpts)
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceCallerSession) Version() (*big.Int, error) {
	return _AggregatorV2V3Interface.Contract.Version(&_AggregatorV2V3Interface.CallOpts)
}

type AggregatorV2V3InterfaceAnswerUpdatedIterator struct {
	Event *AggregatorV2V3InterfaceAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AggregatorV2V3InterfaceAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorV2V3InterfaceAnswerUpdated)
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
		it.Event = new(AggregatorV2V3InterfaceAnswerUpdated)
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

func (it *AggregatorV2V3InterfaceAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *AggregatorV2V3InterfaceAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AggregatorV2V3InterfaceAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*AggregatorV2V3InterfaceAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AggregatorV2V3Interface.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorV2V3InterfaceAnswerUpdatedIterator{contract: _AggregatorV2V3Interface.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorV2V3InterfaceAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _AggregatorV2V3Interface.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AggregatorV2V3InterfaceAnswerUpdated)
				if err := _AggregatorV2V3Interface.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceFilterer) ParseAnswerUpdated(log types.Log) (*AggregatorV2V3InterfaceAnswerUpdated, error) {
	event := new(AggregatorV2V3InterfaceAnswerUpdated)
	if err := _AggregatorV2V3Interface.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AggregatorV2V3InterfaceNewRoundIterator struct {
	Event *AggregatorV2V3InterfaceNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AggregatorV2V3InterfaceNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorV2V3InterfaceNewRound)
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
		it.Event = new(AggregatorV2V3InterfaceNewRound)
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

func (it *AggregatorV2V3InterfaceNewRoundIterator) Error() error {
	return it.fail
}

func (it *AggregatorV2V3InterfaceNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AggregatorV2V3InterfaceNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AggregatorV2V3InterfaceNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AggregatorV2V3Interface.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorV2V3InterfaceNewRoundIterator{contract: _AggregatorV2V3Interface.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *AggregatorV2V3InterfaceNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _AggregatorV2V3Interface.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AggregatorV2V3InterfaceNewRound)
				if err := _AggregatorV2V3Interface.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_AggregatorV2V3Interface *AggregatorV2V3InterfaceFilterer) ParseNewRound(log types.Log) (*AggregatorV2V3InterfaceNewRound, error) {
	event := new(AggregatorV2V3InterfaceNewRound)
	if err := _AggregatorV2V3Interface.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var AggregatorV3InterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

var AggregatorV3InterfaceABI = AggregatorV3InterfaceMetaData.ABI

type AggregatorV3Interface struct {
	AggregatorV3InterfaceCaller
	AggregatorV3InterfaceTransactor
	AggregatorV3InterfaceFilterer
}

type AggregatorV3InterfaceCaller struct {
	contract *bind.BoundContract
}

type AggregatorV3InterfaceTransactor struct {
	contract *bind.BoundContract
}

type AggregatorV3InterfaceFilterer struct {
	contract *bind.BoundContract
}

type AggregatorV3InterfaceSession struct {
	Contract     *AggregatorV3Interface
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AggregatorV3InterfaceCallerSession struct {
	Contract *AggregatorV3InterfaceCaller
	CallOpts bind.CallOpts
}

type AggregatorV3InterfaceTransactorSession struct {
	Contract     *AggregatorV3InterfaceTransactor
	TransactOpts bind.TransactOpts
}

type AggregatorV3InterfaceRaw struct {
	Contract *AggregatorV3Interface
}

type AggregatorV3InterfaceCallerRaw struct {
	Contract *AggregatorV3InterfaceCaller
}

type AggregatorV3InterfaceTransactorRaw struct {
	Contract *AggregatorV3InterfaceTransactor
}

func NewAggregatorV3Interface(address common.Address, backend bind.ContractBackend) (*AggregatorV3Interface, error) {
	contract, err := bindAggregatorV3Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3Interface{AggregatorV3InterfaceCaller: AggregatorV3InterfaceCaller{contract: contract}, AggregatorV3InterfaceTransactor: AggregatorV3InterfaceTransactor{contract: contract}, AggregatorV3InterfaceFilterer: AggregatorV3InterfaceFilterer{contract: contract}}, nil
}

func NewAggregatorV3InterfaceCaller(address common.Address, caller bind.ContractCaller) (*AggregatorV3InterfaceCaller, error) {
	contract, err := bindAggregatorV3Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3InterfaceCaller{contract: contract}, nil
}

func NewAggregatorV3InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorV3InterfaceTransactor, error) {
	contract, err := bindAggregatorV3Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3InterfaceTransactor{contract: contract}, nil
}

func NewAggregatorV3InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorV3InterfaceFilterer, error) {
	contract, err := bindAggregatorV3Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3InterfaceFilterer{contract: contract}, nil
}

func bindAggregatorV3Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorV3InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_AggregatorV3Interface *AggregatorV3InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV3Interface.Contract.AggregatorV3InterfaceCaller.contract.Call(opts, result, method, params...)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.AggregatorV3InterfaceTransactor.contract.Transfer(opts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.AggregatorV3InterfaceTransactor.contract.Transact(opts, method, params...)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV3Interface.Contract.contract.Call(opts, result, method, params...)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.contract.Transfer(opts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV3Interface.Contract.contract.Transact(opts, method, params...)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_AggregatorV3Interface *AggregatorV3InterfaceSession) Decimals() (uint8, error) {
	return _AggregatorV3Interface.Contract.Decimals(&_AggregatorV3Interface.CallOpts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) Decimals() (uint8, error) {
	return _AggregatorV3Interface.Contract.Decimals(&_AggregatorV3Interface.CallOpts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_AggregatorV3Interface *AggregatorV3InterfaceSession) Description() (string, error) {
	return _AggregatorV3Interface.Contract.Description(&_AggregatorV3Interface.CallOpts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) Description() (string, error) {
	return _AggregatorV3Interface.Contract.Description(&_AggregatorV3Interface.CallOpts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AggregatorV3Interface *AggregatorV3InterfaceSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.GetRoundData(&_AggregatorV3Interface.CallOpts, _roundId)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.GetRoundData(&_AggregatorV3Interface.CallOpts, _roundId)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_AggregatorV3Interface *AggregatorV3InterfaceSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.LatestRoundData(&_AggregatorV3Interface.CallOpts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3Interface.Contract.LatestRoundData(&_AggregatorV3Interface.CallOpts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AggregatorV3Interface.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AggregatorV3Interface *AggregatorV3InterfaceSession) Version() (*big.Int, error) {
	return _AggregatorV3Interface.Contract.Version(&_AggregatorV3Interface.CallOpts)
}

func (_AggregatorV3Interface *AggregatorV3InterfaceCallerSession) Version() (*big.Int, error) {
	return _AggregatorV3Interface.Contract.Version(&_AggregatorV3Interface.CallOpts)
}

var MockV3AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"_initialAnswer\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_answer\",\"type\":\"int256\"}],\"name\":\"updateAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"_answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startedAt\",\"type\":\"uint256\"}],\"name\":\"updateRoundData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161055838038061055883398101604081905261002f916100a4565b6000805460ff191660ff84161790556100478161004e565b50506100ff565b60018190554260025560038054906000610067836100d8565b9091555050600380546000908152600460209081526040808320949094558254825260058152838220429081905592548252600690529190912055565b600080604083850312156100b757600080fd5b825160ff811681146100c857600080fd5b6020939093015192949293505050565b6000600182016100f857634e487b7160e01b600052601160045260246000fd5b5060010190565b61044a8061010e6000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80638205bf6a11610081578063b5ab58dc1161005b578063b5ab58dc1461025b578063b633620c1461027b578063feaf968c1461029b57600080fd5b80638205bf6a146101c15780639a6fc8f5146101ca578063a87a20ce1461024857600080fd5b806354fd4d50116100b257806354fd4d5014610171578063668a0f02146101795780637284e4161461018257600080fd5b8063313ce567146100d95780634aa2011f146100fd57806350d25bcd1461015a575b600080fd5b6000546100e69060ff1681565b60405160ff90911681526020015b60405180910390f35b61015861010b36600461033b565b69ffffffffffffffffffff90931660038181556001849055600283905560009182526004602090815260408084209590955581548352600581528483209390935554815260069091522055565b005b61016360015481565b6040519081526020016100f4565b610163600081565b61016360035481565b604080518082018252601f81527f76302e382f74657374732f4d6f636b563341676772656761746f722e736f6c00602082015290516100f49190610374565b61016360025481565b6102116101d83660046103c2565b69ffffffffffffffffffff8116600090815260046020908152604080832054600683528184205460059093529220549293919290918490565b6040805169ffffffffffffffffffff968716815260208101959095528401929092526060830152909116608082015260a0016100f4565b6101586102563660046103e4565b6102c6565b6101636102693660046103e4565b60046020526000908152604090205481565b6101636102893660046103e4565b60056020526000908152604090205481565b6003546000818152600460209081526040808320546006835281842054600590935292205483610211565b600181905542600255600380549060006102df836103fd565b9091555050600380546000908152600460209081526040808320949094558254825260058152838220429081905592548252600690529190912055565b803569ffffffffffffffffffff8116811461033657600080fd5b919050565b6000806000806080858703121561035157600080fd5b61035a8561031c565b966020860135965060408601359560600135945092505050565b600060208083528351808285015260005b818110156103a157858101830151858201604001528201610385565b506000604082860101526040601f19601f8301168501019250505092915050565b6000602082840312156103d457600080fd5b6103dd8261031c565b9392505050565b6000602082840312156103f657600080fd5b5035919050565b600060018201610436577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c6343000813000a",
}

var MockV3AggregatorABI = MockV3AggregatorMetaData.ABI

var MockV3AggregatorBin = MockV3AggregatorMetaData.Bin

func DeployMockV3Aggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _decimals uint8, _initialAnswer *big.Int) (common.Address, *types.Transaction, *MockV3Aggregator, error) {
	parsed, err := MockV3AggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockV3AggregatorBin), backend, _decimals, _initialAnswer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockV3Aggregator{MockV3AggregatorCaller: MockV3AggregatorCaller{contract: contract}, MockV3AggregatorTransactor: MockV3AggregatorTransactor{contract: contract}, MockV3AggregatorFilterer: MockV3AggregatorFilterer{contract: contract}}, nil
}

type MockV3Aggregator struct {
	MockV3AggregatorCaller
	MockV3AggregatorTransactor
	MockV3AggregatorFilterer
}

type MockV3AggregatorCaller struct {
	contract *bind.BoundContract
}

type MockV3AggregatorTransactor struct {
	contract *bind.BoundContract
}

type MockV3AggregatorFilterer struct {
	contract *bind.BoundContract
}

type MockV3AggregatorSession struct {
	Contract     *MockV3Aggregator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockV3AggregatorCallerSession struct {
	Contract *MockV3AggregatorCaller
	CallOpts bind.CallOpts
}

type MockV3AggregatorTransactorSession struct {
	Contract     *MockV3AggregatorTransactor
	TransactOpts bind.TransactOpts
}

type MockV3AggregatorRaw struct {
	Contract *MockV3Aggregator
}

type MockV3AggregatorCallerRaw struct {
	Contract *MockV3AggregatorCaller
}

type MockV3AggregatorTransactorRaw struct {
	Contract *MockV3AggregatorTransactor
}

func NewMockV3Aggregator(address common.Address, backend bind.ContractBackend) (*MockV3Aggregator, error) {
	contract, err := bindMockV3Aggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockV3Aggregator{MockV3AggregatorCaller: MockV3AggregatorCaller{contract: contract}, MockV3AggregatorTransactor: MockV3AggregatorTransactor{contract: contract}, MockV3AggregatorFilterer: MockV3AggregatorFilterer{contract: contract}}, nil
}

func NewMockV3AggregatorCaller(address common.Address, caller bind.ContractCaller) (*MockV3AggregatorCaller, error) {
	contract, err := bindMockV3Aggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorCaller{contract: contract}, nil
}

func NewMockV3AggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*MockV3AggregatorTransactor, error) {
	contract, err := bindMockV3Aggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorTransactor{contract: contract}, nil
}

func NewMockV3AggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*MockV3AggregatorFilterer, error) {
	contract, err := bindMockV3Aggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorFilterer{contract: contract}, nil
}

func bindMockV3Aggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockV3AggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_MockV3Aggregator *MockV3AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockV3Aggregator.Contract.MockV3AggregatorCaller.contract.Call(opts, result, method, params...)
}

func (_MockV3Aggregator *MockV3AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.MockV3AggregatorTransactor.contract.Transfer(opts)
}

func (_MockV3Aggregator *MockV3AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.MockV3AggregatorTransactor.contract.Transact(opts, method, params...)
}

func (_MockV3Aggregator *MockV3AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockV3Aggregator.Contract.contract.Call(opts, result, method, params...)
}

func (_MockV3Aggregator *MockV3AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.contract.Transfer(opts)
}

func (_MockV3Aggregator *MockV3AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.contract.Transact(opts, method, params...)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) Decimals() (uint8, error) {
	return _MockV3Aggregator.Contract.Decimals(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) Decimals() (uint8, error) {
	return _MockV3Aggregator.Contract.Decimals(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) Description() (string, error) {
	return _MockV3Aggregator.Contract.Description(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) Description() (string, error) {
	return _MockV3Aggregator.Contract.Description(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) GetAnswer(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "getAnswer", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) GetAnswer(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetAnswer(&_MockV3Aggregator.CallOpts, arg0)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) GetAnswer(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetAnswer(&_MockV3Aggregator.CallOpts, arg0)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockV3Aggregator.Contract.GetRoundData(&_MockV3Aggregator.CallOpts, _roundId)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockV3Aggregator.Contract.GetRoundData(&_MockV3Aggregator.CallOpts, _roundId)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) GetTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "getTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) GetTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetTimestamp(&_MockV3Aggregator.CallOpts, arg0)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) GetTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _MockV3Aggregator.Contract.GetTimestamp(&_MockV3Aggregator.CallOpts, arg0)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestAnswer(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestAnswer(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) LatestRound() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestRound(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestRound(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockV3Aggregator.Contract.LatestRoundData(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _MockV3Aggregator.Contract.LatestRoundData(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestTimestamp(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _MockV3Aggregator.Contract.LatestTimestamp(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockV3Aggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockV3Aggregator *MockV3AggregatorSession) Version() (*big.Int, error) {
	return _MockV3Aggregator.Contract.Version(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorCallerSession) Version() (*big.Int, error) {
	return _MockV3Aggregator.Contract.Version(&_MockV3Aggregator.CallOpts)
}

func (_MockV3Aggregator *MockV3AggregatorTransactor) UpdateAnswer(opts *bind.TransactOpts, _answer *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.contract.Transact(opts, "updateAnswer", _answer)
}

func (_MockV3Aggregator *MockV3AggregatorSession) UpdateAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.UpdateAnswer(&_MockV3Aggregator.TransactOpts, _answer)
}

func (_MockV3Aggregator *MockV3AggregatorTransactorSession) UpdateAnswer(_answer *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.UpdateAnswer(&_MockV3Aggregator.TransactOpts, _answer)
}

func (_MockV3Aggregator *MockV3AggregatorTransactor) UpdateRoundData(opts *bind.TransactOpts, _roundId *big.Int, _answer *big.Int, _timestamp *big.Int, _startedAt *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.contract.Transact(opts, "updateRoundData", _roundId, _answer, _timestamp, _startedAt)
}

func (_MockV3Aggregator *MockV3AggregatorSession) UpdateRoundData(_roundId *big.Int, _answer *big.Int, _timestamp *big.Int, _startedAt *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.UpdateRoundData(&_MockV3Aggregator.TransactOpts, _roundId, _answer, _timestamp, _startedAt)
}

func (_MockV3Aggregator *MockV3AggregatorTransactorSession) UpdateRoundData(_roundId *big.Int, _answer *big.Int, _timestamp *big.Int, _startedAt *big.Int) (*types.Transaction, error) {
	return _MockV3Aggregator.Contract.UpdateRoundData(&_MockV3Aggregator.TransactOpts, _roundId, _answer, _timestamp, _startedAt)
}

type MockV3AggregatorAnswerUpdatedIterator struct {
	Event *MockV3AggregatorAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockV3AggregatorAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockV3AggregatorAnswerUpdated)
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
		it.Event = new(MockV3AggregatorAnswerUpdated)
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

func (it *MockV3AggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockV3AggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockV3AggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_MockV3Aggregator *MockV3AggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*MockV3AggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _MockV3Aggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorAnswerUpdatedIterator{contract: _MockV3Aggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_MockV3Aggregator *MockV3AggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *MockV3AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _MockV3Aggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockV3AggregatorAnswerUpdated)
				if err := _MockV3Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_MockV3Aggregator *MockV3AggregatorFilterer) ParseAnswerUpdated(log types.Log) (*MockV3AggregatorAnswerUpdated, error) {
	event := new(MockV3AggregatorAnswerUpdated)
	if err := _MockV3Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockV3AggregatorNewRoundIterator struct {
	Event *MockV3AggregatorNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockV3AggregatorNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockV3AggregatorNewRound)
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
		it.Event = new(MockV3AggregatorNewRound)
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

func (it *MockV3AggregatorNewRoundIterator) Error() error {
	return it.fail
}

func (it *MockV3AggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockV3AggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_MockV3Aggregator *MockV3AggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*MockV3AggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _MockV3Aggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &MockV3AggregatorNewRoundIterator{contract: _MockV3Aggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_MockV3Aggregator *MockV3AggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *MockV3AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _MockV3Aggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockV3AggregatorNewRound)
				if err := _MockV3Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_MockV3Aggregator *MockV3AggregatorFilterer) ParseNewRound(log types.Log) (*MockV3AggregatorNewRound, error) {
	event := new(MockV3AggregatorNewRound)
	if err := _MockV3Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

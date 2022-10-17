package testbeaconvrfconsumersmoke

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

var IVRFCoordinatorConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IVRFCoordinatorConsumer.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerSession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFCoordinatorConsumer.Contract.NUMCONFDELAYS(&_IVRFCoordinatorConsumer.CallOpts)
}

func (_IVRFCoordinatorConsumer *IVRFCoordinatorConsumerCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFCoordinatorConsumer.Contract.NUMCONFDELAYS(&_IVRFCoordinatorConsumer.CallOpts)
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

var IVRFCoordinatorExternalAPIMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IVRFCoordinatorExternalAPIABI = IVRFCoordinatorExternalAPIMetaData.ABI

type IVRFCoordinatorExternalAPI struct {
	IVRFCoordinatorExternalAPICaller
	IVRFCoordinatorExternalAPITransactor
	IVRFCoordinatorExternalAPIFilterer
}

type IVRFCoordinatorExternalAPICaller struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorExternalAPITransactor struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorExternalAPIFilterer struct {
	contract *bind.BoundContract
}

type IVRFCoordinatorExternalAPISession struct {
	Contract     *IVRFCoordinatorExternalAPI
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type IVRFCoordinatorExternalAPICallerSession struct {
	Contract *IVRFCoordinatorExternalAPICaller
	CallOpts bind.CallOpts
}

type IVRFCoordinatorExternalAPITransactorSession struct {
	Contract     *IVRFCoordinatorExternalAPITransactor
	TransactOpts bind.TransactOpts
}

type IVRFCoordinatorExternalAPIRaw struct {
	Contract *IVRFCoordinatorExternalAPI
}

type IVRFCoordinatorExternalAPICallerRaw struct {
	Contract *IVRFCoordinatorExternalAPICaller
}

type IVRFCoordinatorExternalAPITransactorRaw struct {
	Contract *IVRFCoordinatorExternalAPITransactor
}

func NewIVRFCoordinatorExternalAPI(address common.Address, backend bind.ContractBackend) (*IVRFCoordinatorExternalAPI, error) {
	contract, err := bindIVRFCoordinatorExternalAPI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorExternalAPI{IVRFCoordinatorExternalAPICaller: IVRFCoordinatorExternalAPICaller{contract: contract}, IVRFCoordinatorExternalAPITransactor: IVRFCoordinatorExternalAPITransactor{contract: contract}, IVRFCoordinatorExternalAPIFilterer: IVRFCoordinatorExternalAPIFilterer{contract: contract}}, nil
}

func NewIVRFCoordinatorExternalAPICaller(address common.Address, caller bind.ContractCaller) (*IVRFCoordinatorExternalAPICaller, error) {
	contract, err := bindIVRFCoordinatorExternalAPI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorExternalAPICaller{contract: contract}, nil
}

func NewIVRFCoordinatorExternalAPITransactor(address common.Address, transactor bind.ContractTransactor) (*IVRFCoordinatorExternalAPITransactor, error) {
	contract, err := bindIVRFCoordinatorExternalAPI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorExternalAPITransactor{contract: contract}, nil
}

func NewIVRFCoordinatorExternalAPIFilterer(address common.Address, filterer bind.ContractFilterer) (*IVRFCoordinatorExternalAPIFilterer, error) {
	contract, err := bindIVRFCoordinatorExternalAPI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorExternalAPIFilterer{contract: contract}, nil
}

func bindIVRFCoordinatorExternalAPI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVRFCoordinatorExternalAPIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFCoordinatorExternalAPI.Contract.IVRFCoordinatorExternalAPICaller.contract.Call(opts, result, method, params...)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.IVRFCoordinatorExternalAPITransactor.contract.Transfer(opts)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.IVRFCoordinatorExternalAPITransactor.contract.Transact(opts, method, params...)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVRFCoordinatorExternalAPI.Contract.contract.Call(opts, result, method, params...)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.contract.Transfer(opts)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.contract.Transact(opts, method, params...)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPICaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IVRFCoordinatorExternalAPI.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPISession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFCoordinatorExternalAPI.Contract.NUMCONFDELAYS(&_IVRFCoordinatorExternalAPI.CallOpts)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPICallerSession) NUMCONFDELAYS() (uint8, error) {
	return _IVRFCoordinatorExternalAPI.Contract.NUMCONFDELAYS(&_IVRFCoordinatorExternalAPI.CallOpts)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactor) RedeemRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.contract.Transact(opts, "redeemRandomness", requestID)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPISession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.RedeemRandomness(&_IVRFCoordinatorExternalAPI.TransactOpts, requestID)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactorSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.RedeemRandomness(&_IVRFCoordinatorExternalAPI.TransactOpts, requestID)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactor) RequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.contract.Transact(opts, "requestRandomness", numWords, subID, confirmationDelayArg)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPISession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.RequestRandomness(&_IVRFCoordinatorExternalAPI.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactorSession) RequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.RequestRandomness(&_IVRFCoordinatorExternalAPI.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactor) RequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.contract.Transact(opts, "requestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPISession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.RequestRandomnessFulfillment(&_IVRFCoordinatorExternalAPI.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_IVRFCoordinatorExternalAPI *IVRFCoordinatorExternalAPITransactorSession) RequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _IVRFCoordinatorExternalAPI.Contract.RequestRandomnessFulfillment(&_IVRFCoordinatorExternalAPI.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

var TestBeaconVRFConsumerSmokeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_coordinator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"name\":\"makeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"randomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestID\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516107b63803806107b683398101604081905261002f91610052565b6001600160a01b03166080526000805465ffffffffffff19166001179055610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b60805161070b6100ab6000396000818161016b0152818161020a01526102e5015261070b6000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c80635a47dd71116100505780635a47dd71146100c157806384d1233b146100d45780638f779201146100e757600080fd5b80631b1d5b27146100775780631f6dd8bb1461009d5780632f7527cc146100a7575b600080fd5b61008a610085366004610431565b610110565b6040519081526020015b60405180910390f35b6100a5610131565b005b6100af600881565b60405160ff9091168152602001610094565b6100a56100cf366004610552565b610208565b6100a56100e2366004610624565b6102a2565b6000546100f99065ffffffffffff1681565b60405165ffffffffffff9091168152602001610094565b6001818154811061012057600080fd5b600091825260209091200154905081565b6000546040517f74d8461100000000000000000000000000000000000000000000000000000000815265ffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906374d84611906024016000604051808303816000875af11580156101c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101f19190810190610650565b8051610205916001916020909101906103d1565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146102925760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c7920636f6f7264696e61746f722063616e2066756c66696c6c0000000060448201526064015b60405180910390fd5b61029d838383610389565b505050565b6040517fdc92accf000000000000000000000000000000000000000000000000000000008152600160048201526000602482015262ffffff8216604482015281907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063dc92accf906064016020604051808303816000875af1158015610343573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036791906106e1565b6000805465ffffffffffff191665ffffffffffff929092169190911790555050565b60405162461bcd60e51b815260206004820152600d60248201527f756e696d706c656d656e746564000000000000000000000000000000000000006044820152606401610289565b82805482825590600052602060002090810192821561040c579160200282015b8281111561040c5782518255916020019190600101906103f1565b5061041892915061041c565b5090565b5b80821115610418576000815560010161041d565b60006020828403121561044357600080fd5b5035919050565b65ffffffffffff8116811461020557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156104b6576104b661045e565b604052919050565b600067ffffffffffffffff8211156104d8576104d861045e565b5060051b60200190565b600082601f8301126104f357600080fd5b813567ffffffffffffffff81111561050d5761050d61045e565b610520601f8201601f191660200161048d565b81815284602083860101111561053557600080fd5b816020850160208301376000918101602001919091529392505050565b60008060006060848603121561056757600080fd5b83356105728161044a565b925060208481013567ffffffffffffffff8082111561059057600080fd5b818701915087601f8301126105a457600080fd5b81356105b76105b2826104be565b61048d565b81815260059190911b8301840190848101908a8311156105d657600080fd5b938501935b828510156105f4578435825293850193908501906105db565b96505050604087013592508083111561060c57600080fd5b505061061a868287016104e2565b9150509250925092565b60006020828403121561063657600080fd5b813562ffffff8116811461064957600080fd5b9392505050565b6000602080838503121561066357600080fd5b825167ffffffffffffffff81111561067a57600080fd5b8301601f8101851361068b57600080fd5b80516106996105b2826104be565b81815260059190911b820183019083810190878311156106b857600080fd5b928401925b828410156106d6578351825292840192908401906106bd565b979650505050505050565b6000602082840312156106f357600080fd5b81516106498161044a56fea164736f6c634300080f000a",
}

var TestBeaconVRFConsumerSmokeABI = TestBeaconVRFConsumerSmokeMetaData.ABI

var TestBeaconVRFConsumerSmokeBin = TestBeaconVRFConsumerSmokeMetaData.Bin

func DeployTestBeaconVRFConsumerSmoke(auth *bind.TransactOpts, backend bind.ContractBackend, _coordinator common.Address) (common.Address, *types.Transaction, *TestBeaconVRFConsumerSmoke, error) {
	parsed, err := TestBeaconVRFConsumerSmokeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBeaconVRFConsumerSmokeBin), backend, _coordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestBeaconVRFConsumerSmoke{TestBeaconVRFConsumerSmokeCaller: TestBeaconVRFConsumerSmokeCaller{contract: contract}, TestBeaconVRFConsumerSmokeTransactor: TestBeaconVRFConsumerSmokeTransactor{contract: contract}, TestBeaconVRFConsumerSmokeFilterer: TestBeaconVRFConsumerSmokeFilterer{contract: contract}}, nil
}

type TestBeaconVRFConsumerSmoke struct {
	TestBeaconVRFConsumerSmokeCaller
	TestBeaconVRFConsumerSmokeTransactor
	TestBeaconVRFConsumerSmokeFilterer
}

type TestBeaconVRFConsumerSmokeCaller struct {
	contract *bind.BoundContract
}

type TestBeaconVRFConsumerSmokeTransactor struct {
	contract *bind.BoundContract
}

type TestBeaconVRFConsumerSmokeFilterer struct {
	contract *bind.BoundContract
}

type TestBeaconVRFConsumerSmokeSession struct {
	Contract     *TestBeaconVRFConsumerSmoke
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TestBeaconVRFConsumerSmokeCallerSession struct {
	Contract *TestBeaconVRFConsumerSmokeCaller
	CallOpts bind.CallOpts
}

type TestBeaconVRFConsumerSmokeTransactorSession struct {
	Contract     *TestBeaconVRFConsumerSmokeTransactor
	TransactOpts bind.TransactOpts
}

type TestBeaconVRFConsumerSmokeRaw struct {
	Contract *TestBeaconVRFConsumerSmoke
}

type TestBeaconVRFConsumerSmokeCallerRaw struct {
	Contract *TestBeaconVRFConsumerSmokeCaller
}

type TestBeaconVRFConsumerSmokeTransactorRaw struct {
	Contract *TestBeaconVRFConsumerSmokeTransactor
}

func NewTestBeaconVRFConsumerSmoke(address common.Address, backend bind.ContractBackend) (*TestBeaconVRFConsumerSmoke, error) {
	contract, err := bindTestBeaconVRFConsumerSmoke(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestBeaconVRFConsumerSmoke{TestBeaconVRFConsumerSmokeCaller: TestBeaconVRFConsumerSmokeCaller{contract: contract}, TestBeaconVRFConsumerSmokeTransactor: TestBeaconVRFConsumerSmokeTransactor{contract: contract}, TestBeaconVRFConsumerSmokeFilterer: TestBeaconVRFConsumerSmokeFilterer{contract: contract}}, nil
}

func NewTestBeaconVRFConsumerSmokeCaller(address common.Address, caller bind.ContractCaller) (*TestBeaconVRFConsumerSmokeCaller, error) {
	contract, err := bindTestBeaconVRFConsumerSmoke(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestBeaconVRFConsumerSmokeCaller{contract: contract}, nil
}

func NewTestBeaconVRFConsumerSmokeTransactor(address common.Address, transactor bind.ContractTransactor) (*TestBeaconVRFConsumerSmokeTransactor, error) {
	contract, err := bindTestBeaconVRFConsumerSmoke(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestBeaconVRFConsumerSmokeTransactor{contract: contract}, nil
}

func NewTestBeaconVRFConsumerSmokeFilterer(address common.Address, filterer bind.ContractFilterer) (*TestBeaconVRFConsumerSmokeFilterer, error) {
	contract, err := bindTestBeaconVRFConsumerSmoke(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestBeaconVRFConsumerSmokeFilterer{contract: contract}, nil
}

func bindTestBeaconVRFConsumerSmoke(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestBeaconVRFConsumerSmokeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestBeaconVRFConsumerSmoke.Contract.TestBeaconVRFConsumerSmokeCaller.contract.Call(opts, result, method, params...)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.TestBeaconVRFConsumerSmokeTransactor.contract.Transfer(opts)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.TestBeaconVRFConsumerSmokeTransactor.contract.Transact(opts, method, params...)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestBeaconVRFConsumerSmoke.Contract.contract.Call(opts, result, method, params...)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.contract.Transfer(opts)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.contract.Transact(opts, method, params...)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestBeaconVRFConsumerSmoke.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeSession) NUMCONFDELAYS() (uint8, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.NUMCONFDELAYS(&_TestBeaconVRFConsumerSmoke.CallOpts)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.NUMCONFDELAYS(&_TestBeaconVRFConsumerSmoke.CallOpts)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeCaller) Randomness(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestBeaconVRFConsumerSmoke.contract.Call(opts, &out, "randomness", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeSession) Randomness(arg0 *big.Int) (*big.Int, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.Randomness(&_TestBeaconVRFConsumerSmoke.CallOpts, arg0)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeCallerSession) Randomness(arg0 *big.Int) (*big.Int, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.Randomness(&_TestBeaconVRFConsumerSmoke.CallOpts, arg0)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeCaller) RequestID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestBeaconVRFConsumerSmoke.contract.Call(opts, &out, "requestID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeSession) RequestID() (*big.Int, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.RequestID(&_TestBeaconVRFConsumerSmoke.CallOpts)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeCallerSession) RequestID() (*big.Int, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.RequestID(&_TestBeaconVRFConsumerSmoke.CallOpts)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactor) MakeRequest(opts *bind.TransactOpts, confirmationDelay *big.Int) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.contract.Transact(opts, "makeRequest", confirmationDelay)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeSession) MakeRequest(confirmationDelay *big.Int) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.MakeRequest(&_TestBeaconVRFConsumerSmoke.TransactOpts, confirmationDelay)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactorSession) MakeRequest(confirmationDelay *big.Int) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.MakeRequest(&_TestBeaconVRFConsumerSmoke.TransactOpts, confirmationDelay)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords, arguments)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.RawFulfillRandomWords(&_TestBeaconVRFConsumerSmoke.TransactOpts, requestID, randomWords, arguments)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.RawFulfillRandomWords(&_TestBeaconVRFConsumerSmoke.TransactOpts, requestID, randomWords, arguments)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactor) RedeemRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.contract.Transact(opts, "redeemRequest")
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeSession) RedeemRequest() (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.RedeemRequest(&_TestBeaconVRFConsumerSmoke.TransactOpts)
}

func (_TestBeaconVRFConsumerSmoke *TestBeaconVRFConsumerSmokeTransactorSession) RedeemRequest() (*types.Transaction, error) {
	return _TestBeaconVRFConsumerSmoke.Contract.RedeemRequest(&_TestBeaconVRFConsumerSmoke.TransactOpts)
}

var VRFBeaconTypesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

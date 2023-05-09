package vrfbeacon

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

type VRFBeaconReportReport struct {
	Outputs            []VRFBeaconTypesVRFOutput
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	RecentBlockHeight  uint64
	RecentBlockHash    [32]byte
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
}

var AccessControllerInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

var AccessControllerInterfaceABI = AccessControllerInterfaceMetaData.ABI

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

var IVRFCoordinatorProducerAPIMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"paymentsInJuels\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubscriptionLinkBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconTypes.VRFOutput[]\",\"name\":\"vrfOutputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"processVRFOutputs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"confDelays\",\"type\":\"uint24[8]\"}],\"name\":\"setConfirmationDelays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"gasPrice\",\"type\":\"uint64\"}],\"name\":\"setReasonableGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"juelsAmount\",\"type\":\"uint256\"}],\"name\":\"transferLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactor) ProcessVRFOutputs(opts *bind.TransactOpts, vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64, blockHash [32]byte) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.contract.Transact(opts, "processVRFOutputs", vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight, blockHash)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) ProcessVRFOutputs(vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64, blockHash [32]byte) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.ProcessVRFOutputs(&_IVRFCoordinatorProducerAPI.TransactOpts, vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight, blockHash)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorSession) ProcessVRFOutputs(vrfOutputs []VRFBeaconTypesVRFOutput, juelsPerFeeCoin *big.Int, reasonableGasPrice uint64, blockHeight uint64, blockHash [32]byte) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.ProcessVRFOutputs(&_IVRFCoordinatorProducerAPI.TransactOpts, vrfOutputs, juelsPerFeeCoin, reasonableGasPrice, blockHeight, blockHash)
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactor) SetReasonableGasPrice(opts *bind.TransactOpts, gasPrice uint64) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.contract.Transact(opts, "setReasonableGasPrice", gasPrice)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPISession) SetReasonableGasPrice(gasPrice uint64) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.SetReasonableGasPrice(&_IVRFCoordinatorProducerAPI.TransactOpts, gasPrice)
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPITransactorSession) SetReasonableGasPrice(gasPrice uint64) (*types.Transaction, error) {
	return _IVRFCoordinatorProducerAPI.Contract.SetReasonableGasPrice(&_IVRFCoordinatorProducerAPI.TransactOpts, gasPrice)
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

type IVRFCoordinatorProducerAPIConfigSetIterator struct {
	Event *IVRFCoordinatorProducerAPIConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorProducerAPIConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorProducerAPIConfigSet)
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
		it.Event = new(IVRFCoordinatorProducerAPIConfigSet)
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

func (it *IVRFCoordinatorProducerAPIConfigSetIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorProducerAPIConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorProducerAPIConfigSet struct {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterConfigSet(opts *bind.FilterOpts) (*IVRFCoordinatorProducerAPIConfigSetIterator, error) {

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIConfigSetIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIConfigSet) (event.Subscription, error) {

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorProducerAPIConfigSet)
				if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) ParseConfigSet(log types.Log) (*IVRFCoordinatorProducerAPIConfigSet, error) {
	event := new(IVRFCoordinatorProducerAPIConfigSet)
	if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type IVRFCoordinatorProducerAPINewTransmissionIterator struct {
	Event *IVRFCoordinatorProducerAPINewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *IVRFCoordinatorProducerAPINewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVRFCoordinatorProducerAPINewTransmission)
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
		it.Event = new(IVRFCoordinatorProducerAPINewTransmission)
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

func (it *IVRFCoordinatorProducerAPINewTransmissionIterator) Error() error {
	return it.fail
}

func (it *IVRFCoordinatorProducerAPINewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type IVRFCoordinatorProducerAPINewTransmission struct {
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*IVRFCoordinatorProducerAPINewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPINewTransmissionIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPINewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(IVRFCoordinatorProducerAPINewTransmission)
				if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) ParseNewTransmission(log types.Log) (*IVRFCoordinatorProducerAPINewTransmission, error) {
	event := new(IVRFCoordinatorProducerAPINewTransmission)
	if err := _IVRFCoordinatorProducerAPI.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Raw                    types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIRandomnessFulfillmentRequestedIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
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
	Raw                    types.Log
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*IVRFCoordinatorProducerAPIRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &IVRFCoordinatorProducerAPIRandomnessRequestedIterator{contract: _IVRFCoordinatorProducerAPI.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_IVRFCoordinatorProducerAPI *IVRFCoordinatorProducerAPIFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *IVRFCoordinatorProducerAPIRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IVRFCoordinatorProducerAPI.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
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

var VRFBeaconMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"coordinator\",\"type\":\"address\"},{\"internalType\":\"contractDKG\",\"name\":\"keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"CalldataLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptPayeeship\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfFaultyOracles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTransmitters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numPayees\",\"type\":\"uint256\"}],\"name\":\"IncorrectNumberOfPayees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expectedNumSignatures\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ssLength\",\"type\":\"uint256\"}],\"name\":\"IncorrectNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeftGasExceedsInitialGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeOwnerOrBillingAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numFaultyOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"}],\"name\":\"NumberOfFaultyOraclesTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OnchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"OnlyActiveSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"OnlyActiveTransmitters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCurrentPayee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"existingPayee\",\"type\":\"address\"}],\"name\":\"PayeeAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedSignerAddress\",\"type\":\"address\"}],\"name\":\"RepeatedSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedTransmitterAddress\",\"type\":\"address\"}],\"name\":\"RepeatedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"}],\"name\":\"SeenReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numTransmitters\",\"type\":\"uint256\"}],\"name\":\"SignersTransmittersMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maxOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"providedOracles\",\"type\":\"uint256\"}],\"name\":\"TooManyOracles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ocrVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconTypes.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_coordinator\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyProvider\",\"outputs\":[{\"internalType\":\"contractDKG\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162004b5838038062004b588339810160408190526200003491620001c7565b8181858581813380600081620000915760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c457620000c48162000103565b5050506001600160a01b03918216608052811660a052601380546001600160a01b03191695909116949094179093555060c05250620002219350505050565b336001600160a01b038216036200015d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000088565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b0381168114620001c457600080fd5b50565b60008060008060808587031215620001de57600080fd5b8451620001eb81620001ae565b6020860151909450620001fe81620001ae565b60408601519093506200021181620001ae565b6060959095015193969295505050565b60805160a05160c05161487b620002dd600039600061045201526000818161035c01528181610fdc01528181611081015281816111510152818161121f015281816112bd01528181611ef501528181611f9a01528181612135015281816124fd015281816125e801528181612a2b01528181612ad001528181612fd9015261348b015260008181610305015281816110af0152818161124c01528181611fc8015281816121b401528181612afe0152612ee9015261487b6000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c8063b121e14711610104578063d09dc339116100a2578063e53bbc9a11610071578063e53bbc9a146104c0578063eb5dcd6c146104d3578063f2fde38b146104e6578063fbffd2c1146104f957600080fd5b8063d09dc33914610474578063d57fc45a1461047c578063e3d0e71214610485578063e4902f821461049857600080fd5b8063bf2732c7116100de578063bf2732c714610416578063c107532914610429578063c4c92b371461043c578063cc31f7dd1461044d57600080fd5b8063b121e147146103df578063b1dc65a4146103f2578063b8be03cd1461040557600080fd5b80637d253aff116101715780638ac28d5a1161014b5780638ac28d5a1461037e5780638da5cb5b146103915780639c849b30146103a2578063afcb95d7146103b557600080fd5b80637d253aff1461030057806381ff7048146103275780638a1b17721461035757600080fd5b80632f7527cc116101ad5780632f7527cc146102a957806355e48749146102c35780635f27026f146102cd57806379ba5097146102f857600080fd5b80630eafb25b146101d4578063181f5a77146101fa5780632993726814610239575b600080fd5b6101e76101e2366004613638565b61050c565b6040519081526020015b60405180910390f35b604080518082018252600f81527f565246426561636f6e20312e302e300000000000000000000000000000000000602082015290516101f191906136a5565b6002546003546040805165010000000000840467ffffffffffffffff90811682526d0100000000000000000000000000850481166020830152600160a81b909404841691810191909152918116606083015268010000000000000000900462ffffff16608082015260a0016101f1565b6102b1600881565b60405160ff90911681526020016101f1565b6102cb610605565b005b6013546102e0906001600160a01b031681565b6040516001600160a01b0390911681526020016101f1565b6102cb61064f565b6102e07f000000000000000000000000000000000000000000000000000000000000000081565b6004546005546040805163ffffffff808516825264010000000090940490931660208401528201526060016101f1565b6102e07f000000000000000000000000000000000000000000000000000000000000000081565b6102cb61038c366004613638565b610700565b6000546001600160a01b03166102e0565b6102cb6103b0366004613704565b610743565b6005546006546040805160008152602081019390935263ffffffff909116908201526060016101f1565b6102cb6103ed366004613638565b610915565b6102cb6104003660046137b2565b6109db565b6102cb610413366004613869565b50565b6102cb610424366004613a58565b610eb2565b6102cb610437366004613b25565b610f20565b6012546001600160a01b03166102e0565b6101e77f000000000000000000000000000000000000000000000000000000000000000081565b6101e7611208565b6101e760145481565b6102cb610493366004613b88565b611368565b6104ab6104a6366004613638565b611ad1565b60405163ffffffff90911681526020016101f1565b6102cb6104ce366004613c87565b611b85565b6102cb6104e1366004613cf8565b611d71565b6102cb6104f4366004613638565b611e65565b6102cb610507366004613638565b611e76565b6001600160a01b03811660009081526008602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b0316918101919091529061056e5750600092915050565b6002546020820151600091600160a81b900467ffffffffffffffff1690600c9060ff16601f81106105a1576105a1613d31565b6008810491909101546002546105d29260071660040261010090810a90920463ffffffff9081169290910416613d5d565b63ffffffff166105e29190613d81565b905081604001516001600160601b0316816105fd9190613d98565b949350505050565b6013546001600160a01b03163381146106475760405163292f4fb560e01b81523360048201526001600160a01b03821660248201526044015b60405180910390fd5b506000601455565b6001546001600160a01b031633146106a95760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161063e565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b0381811660009081526010602052604090205416331461073a57604051633738e30960e21b815260040160405180910390fd5b61041381611e87565b61074b612233565b82811461078e576040517f36d20459000000000000000000000000000000000000000000000000000000008152600481018490526024810182905260440161063e565b60005b8381101561090e5760008585838181106107ad576107ad613d31565b90506020020160208101906107c29190613638565b905060008484848181106107d8576107d8613d31565b90506020020160208101906107ed9190613638565b6001600160a01b0380841660009081526010602052604090205491925016801580158161082c5750826001600160a01b0316826001600160a01b031614155b15610876576040517febdf17560000000000000000000000000000000000000000000000000000000081526001600160a01b0380861660048301528316602482015260440161063e565b6001600160a01b03848116600090815260106020526040902080546001600160a01b031916858316908117909155908316146108f757826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b50505050808061090690613dab565b915050610791565b5050505050565b6001600160a01b03818116600090815260116020526040902054163314610968576040517f9d12ec4f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381811660008181526010602090815260408083208054336001600160a01b031980831682179093556011909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a6040805160e08101825260025460ff8082168352610100820463ffffffff1660208085019190915265010000000000830467ffffffffffffffff908116858701526d0100000000000000000000000000840481166060860152600160a81b9093048316608085015260035492831660a08501526801000000000000000090920462ffffff1660c08401523360009081526008835293909320549394509092908c01359116610aba576040517fb1c1f68e00000000000000000000000000000000000000000000000000000000815233600482015260240161063e565b6005548b3514610b04576005546040517f93df584c00000000000000000000000000000000000000000000000000000000815260048101919091528b35602482015260440161063e565b610b128a8a8a8a8a8a61228f565b8151610b1f906001613dc4565b60ff1687141580610b305750868514155b15610b88578151610b42906001613dc4565b6040517ffc33647500000000000000000000000000000000000000000000000000000000815260ff9091166004820152602481018890526044810186905260640161063e565b60008a8a604051610b9a929190613ddd565b604051908190038120610bb1918e90602001613ded565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a811015610d4c5760006001858a8460208110610c0057610c00613d31565b610c0d91901a601b613dc4565b8f8f86818110610c1f57610c1f613d31565b905060200201358e8e87818110610c3857610c38613d31565b9050602002013560405160008152602001604052604051610c75949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610c97573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526009602090815290849020838501909452925460ff8082161515808552610100909204169383019390935290955092509050610d25576040517f20fb74ee0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161063e565b826020015160080260ff166001901b84019350508080610d4490613dab565b915050610be1565b5081827e010101010101010101010101010101010101010101010101010101010101011614610da7576040517fc103be2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505060008a8a604051610dbc929190613ddd565b604080519182900390912060008181526007602052919091205490915060ff1615610e16576040517f0b8d39d50000000000000000000000000000000000000000000000000000000081526004810182905260240161063e565b6000908152600760205260408120805460ff1916600117905580610e78848e836020020135858f8f8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061231f92505050565b6006805463ffffffff191663ffffffff600888901c161790559092509050610ea384838388336126f1565b50505050505050505050505050565b6013546001600160a01b0316338114610eef5760405163292f4fb560e01b81523360048201526001600160a01b038216602482015260440161063e565b8151604051610f019190602001613e01565b60408051601f1981840301815291905280516020909101206014555050565b6000546001600160a01b03163314801590610fae5750601254604051630d629b5f60e31b81526001600160a01b0390911690636b14daf890610f6b9033906000903690600401613e46565b602060405180830381865afa158015610f88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fac9190613e69565b155b15610fcc57604051631809d98560e31b815260040160405180910390fd5b6000610fd6612809565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611038573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105c9190613e8b565b9050600061106a8284613d98565b6040516370a0823160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156110f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111a9190613e8b565b9050818110156111475760405163cf47918160e01b8152600481018290526024810183905260440161063e565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663f99b1d688761118a6111848686613ea4565b896129c7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b1580156111e857600080fd5b505af11580156111fc573d6000803e3d6000fd5b50505050505050505050565b6040516370a0823160e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116600483015260009182917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611293573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b79190613e8b565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611319573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061133d9190613e8b565b90506000611349612809565b9050816113568285613eb7565b6113609190613eb7565b935050505090565b888787601f8311156113b0576040517f809fc428000000000000000000000000000000000000000000000000000000008152601f60048201526024810184905260440161063e565b8183146113f3576040517f988a0804000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260440161063e565b6113fe816003613ed7565b60ff168311611445576040517ffda9db7800000000000000000000000000000000000000000000000000000000815260ff821660048201526024810184905260440161063e565b6114518160ff166129e1565b611459612233565b60006040518060c001604052808f8f80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505081526020018d8d8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050509082525060ff8c1660208083019190915260408051601f8d0183900483028101830182528c8152920191908c908c908190840183828082843760009201919091525050509082525067ffffffffffffffff891660208083019190915260408051601f8a0183900483028101830182528981529201919089908990819084018382808284376000920191909152505050915250905061157b612a1b565b600a5460005b8181101561162c576000600a828154811061159e5761159e613d31565b6000918252602082200154600b80546001600160a01b03909216935090849081106115cb576115cb613d31565b60009182526020808320909101546001600160a01b039485168352600982526040808420805461ffff1916905594168252600890529190912080546dffffffffffffffffffffffffffff19169055508061162481613dab565b915050611581565b50611639600a60006134f8565b611645600b60006134f8565b60005b8251518110156118f157600960008460000151838151811061166c5761166c613d31565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156116f65782518051829081106116ac576116ac613d31565b60200260200101516040517f7451f83e00000000000000000000000000000000000000000000000000000000815260040161063e91906001600160a01b0391909116815260200190565b604080518082019091526001815260ff82166020820152835180516009916000918590811061172757611727613d31565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff9095169490940293909317909255840151805160089291908490811061179257611792613d31565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff161561181e57826020015181815181106117d4576117d4613d31565b60200260200101516040517fe8d2989900000000000000000000000000000000000000000000000000000000815260040161063e91906001600160a01b0391909116815260200190565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b0316815250600860008560200151848151811061186357611863613d31565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516001600160601b031662010000026dffffffffffffffffffffffff00001960ff959095166101000261ff00199315159390931661ffff19909416939093179190911792909216179055806118e981613dab565b915050611648565b508151805161190891600a91602090910190613516565b50602080830151805161191f92600b920190613516565b5060408201516002805460ff191660ff909216919091179055600454640100000000900463ffffffff16611951613050565b6004805463ffffffff9283166401000000000267ffffffff0000000019821681179092556000926119889281169116176001613ef3565b905080600460006101000a81548163ffffffff021916908363ffffffff16021790555060006119dc46308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a001516130da565b9050806005819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e058360055484886000015189602001518a604001518b606001518c608001518d60a00151604051611a3e99989796959493929190613f54565b60405180910390a1600254610100900463ffffffff1660005b865151811015611ab15781600c82601f8110611a7557611a75613d31565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff1602179055508080611aa990613dab565b915050611a57565b50611abc8e8e613167565b50505050505050505050505050505050505050565b6001600160a01b03811660009081526008602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b03169181019190915290611b335750600092915050565b600c816020015160ff16601f8110611b4d57611b4d613d31565b600881049190910154600254611b7e9260071660040261010090810a90920463ffffffff9081169290910416613d5d565b9392505050565b6000546001600160a01b03163314801590611c135750601254604051630d629b5f60e31b81526001600160a01b0390911690636b14daf890611bd09033906000903690600401613e46565b602060405180830381865afa158015611bed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c119190613e69565b155b15611c3157604051631809d98560e31b815260040160405180910390fd5b611c39612a1b565b6002805467ffffffffffffffff858116600160a81b81027fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff8984166d010000000000000000000000000081027fffffffffffffffffffffff0000000000000000ffffffffffffffffffffffffff8d8716650100000000008102919091167fffffffffffffffffffffff00000000000000000000000000000000ffffffffff909816979097171791909116919091179094556003805462ffffff87166801000000000000000081026affffffffffffffffffffff199092169489169485179190911790915560408051948552602085019590955293830152606082015260808101919091527f49275ddcdfc9c0519b3d094308c8bf675f06070a754ce90c152163cb6e66e8a09060a00160405180910390a15050505050565b6001600160a01b03828116600090815260106020526040902054163314611dab57604051633738e30960e21b815260040160405180910390fd5b6001600160a01b0381163303611ded576040517fb387a23800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03808316600090815260116020526040902080548383166001600160a01b031982168117909255909116908114611e60576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a45b505050565b611e6d612233565b61041381613175565b611e7e612233565b6104138161321e565b6001600160a01b0381166000908152600860209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b031692810192909252611ee4575050565b6000611eef8361050c565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611f51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f759190613e8b565b90506000611f838284613d98565b6040516370a0823160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561200f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120339190613e8b565b9050818110156120605760405163cf47918160e01b8152600481018290526024810183905260440161063e565b831561222b576001600160a01b038681166000908152601060209081526040909120546002549188015192169161010090910463ffffffff1690600c9060ff16601f81106120b0576120b0613d31565b6008808204909201805463ffffffff9485166004600790941684026101000a90810295021916939093179092556001600160a01b03808a16600090815260209290925260409182902080546dffffffffffffffffffffffff00001916905590517ff99b1d680000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000009091169163f99b1d68916121809185918a91016001600160a01b03929092168252602082015260400190565b600060405180830381600087803b15801561219a57600080fd5b505af11580156121ae573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316886001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8860405161222191815260200190565b60405180910390a4505b505050505050565b6000546001600160a01b0316331461228d5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161063e565b565b600061229c826020613d81565b6122a7856020613d81565b6122b388610144613d98565b6122bd9190613d98565b6122c79190613d98565b6122d2906000613d98565b9050368114612316576040517ff7b94f0a0000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161063e565b50505050505050565b60008060008380602001905181019061233891906141db565b60208801805191925061234a826143c1565b63ffffffff1663ffffffff168152505086600260008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548163ffffffff021916908363ffffffff16021790555060408201518160000160056101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160086101000a81548162ffffff021916908362ffffff16021790555090505060006124988260600151613294565b9050808260800151146124fb57608082015160608301516040517faed0afe500000000000000000000000000000000000000000000000000000000815260048101929092526024820183905267ffffffffffffffff16604482015260640161063e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316625bd524836000015184602001518560400151866060015187608001516040518663ffffffff1660e01b81526004016125629594939291906144d5565b6000604051808303816000875af1158015612581573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526125a991908101906145d8565b5060408281015190517f05f4acc600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906305f4acc690602401600060405180830381600087803b15801561263457600080fd5b505af1158015612648573d6000803e3d6000fd5b505050508564ffffffffff16886020015163ffffffff167f27bf3f1077f091da6885751ba10f5775d06657fd59e47a6ab1f7635e5a115afe33856020015186604001518c6040516126d094939291906001600160a01b039490941684526001600160c01b0392909216602084015267ffffffffffffffff166040830152606082015260800190565b60405180910390a38160200151826040015193509350505094509492505050565b600061271d3a67ffffffffffffffff86161561270d5785612713565b87606001515b8860400151613370565b90506010360260005a905060006127468663ffffffff1685858c60c0015162ffffff16866133c1565b90506000670de0b6b3a76400006001600160c01b038a1683026001600160a01b03881660009081526008602052604090205460a08d01519290910492506201000090046001600160601b039081169167ffffffffffffffff16828401019081168211156127b9575050505050505061090e565b6001600160a01b038816600090815260086020526040902080546001600160601b0390921662010000026dffffffffffffffffffffffff0000199092169190911790555050505050505050505050565b600080600b80548060200260200160405190810160405280929190818152602001828054801561286257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612844575b50508351600254604080516103e0810191829052969750919561010090910463ffffffff169450600093509150600c90601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116128995790505050505050905060005b8381101561292c578181601f81106128f9576128f9613d31565b60200201516129089084613d5d565b6129189063ffffffff1687613d98565b95508061292481613dab565b9150506128df565b5060025461294b90600160a81b900467ffffffffffffffff1686613d81565b945060005b838110156129bf576008600086838151811061296e5761296e613d31565b6020908102919091018101516001600160a01b03168252810191909152604001600020546129ab906201000090046001600160601b031687613d98565b9550806129b781613dab565b915050612950565b505050505090565b6000818310156129d85750816129db565b50805b92915050565b80600003610413576040517fe77dba5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000612a25612809565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aab9190613e8b565b90506000612ab98284613d98565b6040516370a0823160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015612b45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b699190613e8b565b905081811015612b965760405163cf47918160e01b8152600481018290526024810183905260440161063e565b600254604080516103e081019182905261010090920463ffffffff1691600091600c90601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612bc0579050505050505090506000600b805480602002602001604051908101604052809291908181526020018280548015612c5b57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612c3d575b5050505050905060008151905060008167ffffffffffffffff811115612c8357612c836138a4565b604051908082528060200260200182016040528015612cac578160200160208202803683370190505b50905060008267ffffffffffffffff811115612cca57612cca6138a4565b604051908082528060200260200182016040528015612cf3578160200160208202803683370190505b5090506000805b84811015612f8457600060086000888481518110612d1a57612d1a613d31565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b03169050600060086000898581518110612d7c57612d7c613d31565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008883601f8110612ddf57612ddf613d31565b6020020151600254908b0363ffffffff169150600160a81b900467ffffffffffffffff16810282018015612f79576000601060008b8781518110612e2557612e25613d31565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b0316905080888781518110612e7657612e76613d31565b60200260200101906001600160a01b031690816001600160a01b03168152505081878781518110612ea957612ea9613d31565b6020026020010181815250508b8b86601f8110612ec857612ec8613d31565b602002019063ffffffff16908163ffffffff168152505085806001019650507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b03168b8781518110612f2d57612f2d613d31565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c85604051612f6f91815260200190565b60405180910390a4505b505050600101612cfa565b5081518114612f94578082528083525b612fa1600c87601f61357b565b50815115613043576040517f73433a2f0000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906373433a2f9061301090869086906004016146ca565b600060405180830381600087803b15801561302a57600080fd5b505af115801561303e573d6000803e3d6000fd5b505050505b5050505050505050505050565b60004661a4b1811480613065575062066eed81145b156130d35760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130cd9190613e8b565b91505090565b4391505090565b6000808a8a8a8a8a8a8a8a8a6040516020016130fe99989796959493929190614721565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6131718282613409565b5050565b336001600160a01b038216036131cd5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161063e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b03908116908216811461317157601280546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b60004661a4b18114806132a9575062066eed81145b15613360576101008367ffffffffffffffff166132c4613050565b6132ce9190613ea4565b11156132dd5750600092915050565b6040517f2b407a8200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152606490632b407a8290602401602060405180830381865afa15801561333c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b7e9190613e8b565b505067ffffffffffffffff164090565b60008367ffffffffffffffff84168110156133a4576002858567ffffffffffffffff1603816133a1576133a16146b4565b04015b6133b8818467ffffffffffffffff166129c7565b95945050505050565b6000818610156133fd576040517f3fef97df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50909303019091010290565b61010081811461344b578282826040517f418a179b00000000000000000000000000000000000000000000000000000000815260040161063e939291906147a9565b6000613459838501856147cd565b90506040517f8eef585f0000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eef585f906134c0908490600401614837565b600060405180830381600087803b1580156134da57600080fd5b505af11580156134ee573d6000803e3d6000fd5b5050505050505050565b5080546000825590600052602060002090810190610413919061360e565b82805482825590600052602060002090810192821561356b579160200282015b8281111561356b57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613536565b5061357792915061360e565b5090565b60048301918390821561356b5791602002820160005b838211156135d557835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302613591565b80156136055782816101000a81549063ffffffff02191690556004016020816003010492830192600103026135d5565b50506135779291505b5b80821115613577576000815560010161360f565b6001600160a01b038116811461041357600080fd5b60006020828403121561364a57600080fd5b8135611b7e81613623565b60005b83811015613670578181015183820152602001613658565b50506000910152565b60008151808452613691816020860160208601613655565b601f01601f19169290920160200192915050565b602081526000611b7e6020830184613679565b60008083601f8401126136ca57600080fd5b50813567ffffffffffffffff8111156136e257600080fd5b6020830191508360208260051b85010111156136fd57600080fd5b9250929050565b6000806000806040858703121561371a57600080fd5b843567ffffffffffffffff8082111561373257600080fd5b61373e888389016136b8565b9096509450602087013591508082111561375757600080fd5b50613764878288016136b8565b95989497509550505050565b60008083601f84011261378257600080fd5b50813567ffffffffffffffff81111561379a57600080fd5b6020830191508360208285010111156136fd57600080fd5b60008060008060008060008060e0898b0312156137ce57600080fd5b606089018a8111156137df57600080fd5b8998503567ffffffffffffffff808211156137f957600080fd5b6138058c838d01613770565b909950975060808b013591508082111561381e57600080fd5b61382a8c838d016136b8565b909750955060a08b013591508082111561384357600080fd5b506138508b828c016136b8565b999c989b50969995989497949560c00135949350505050565b60006020828403121561387b57600080fd5b813567ffffffffffffffff81111561389257600080fd5b820160a08185031215611b7e57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156138dd576138dd6138a4565b60405290565b604051610100810167ffffffffffffffff811182821017156138dd576138dd6138a4565b60405160a0810167ffffffffffffffff811182821017156138dd576138dd6138a4565b6040516080810167ffffffffffffffff811182821017156138dd576138dd6138a4565b6040516020810167ffffffffffffffff811182821017156138dd576138dd6138a4565b604051601f8201601f1916810167ffffffffffffffff81118282101715613999576139996138a4565b604052919050565b600067ffffffffffffffff8211156139bb576139bb6138a4565b50601f01601f191660200190565b600067ffffffffffffffff8211156139e3576139e36138a4565b5060051b60200190565b600082601f8301126139fe57600080fd5b81356020613a13613a0e836139c9565b613970565b82815260059290921b84018101918181019086841115613a3257600080fd5b8286015b84811015613a4d5780358352918301918301613a36565b509695505050505050565b60006020808385031215613a6b57600080fd5b823567ffffffffffffffff80821115613a8357600080fd5b9084019060408287031215613a9757600080fd5b613a9f6138ba565b823582811115613aae57600080fd5b8301601f81018813613abf57600080fd5b8035613acd613a0e826139a1565b8181528987838501011115613ae157600080fd5b818784018883013760008783830101528084525050508383013582811115613b0857600080fd5b613b14888286016139ed565b948201949094529695505050505050565b60008060408385031215613b3857600080fd5b8235613b4381613623565b946020939093013593505050565b803560ff81168114613b6257600080fd5b919050565b67ffffffffffffffff8116811461041357600080fd5b8035613b6281613b67565b60008060008060008060008060008060c08b8d031215613ba757600080fd5b8a3567ffffffffffffffff80821115613bbf57600080fd5b613bcb8e838f016136b8565b909c509a5060208d0135915080821115613be457600080fd5b613bf08e838f016136b8565b909a509850889150613c0460408e01613b51565b975060608d0135915080821115613c1a57600080fd5b613c268e838f01613770565b9097509550859150613c3a60808e01613b7d565b945060a08d0135915080821115613c5057600080fd5b50613c5d8d828e01613770565b915080935050809150509295989b9194979a5092959850565b62ffffff8116811461041357600080fd5b600080600080600060a08688031215613c9f57600080fd5b8535613caa81613b67565b94506020860135613cba81613b67565b93506040860135613cca81613b67565b92506060860135613cda81613b67565b91506080860135613cea81613c76565b809150509295509295909350565b60008060408385031215613d0b57600080fd5b8235613d1681613623565b91506020830135613d2681613623565b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b63ffffffff828116828216039080821115613d7a57613d7a613d47565b5092915050565b80820281158282048414176129db576129db613d47565b808201808211156129db576129db613d47565b600060018201613dbd57613dbd613d47565b5060010190565b60ff81811683821601908111156129db576129db613d47565b8183823760009101908152919050565b828152606082602083013760800192915050565b60008251613e13818460208701613655565b9190910192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03841681526040602082015260006133b8604083018486613e1d565b600060208284031215613e7b57600080fd5b81518015158114611b7e57600080fd5b600060208284031215613e9d57600080fd5b5051919050565b818103818111156129db576129db613d47565b8181036000831280158383131683831282161715613d7a57613d7a613d47565b60ff8181168382160290811690818114613d7a57613d7a613d47565b63ffffffff818116838216019080821115613d7a57613d7a613d47565b600081518084526020808501945080840160005b83811015613f495781516001600160a01b031687529582019590820190600101613f24565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152613f848184018a613f10565b90508281036080840152613f988189613f10565b905060ff871660a084015282810360c0840152613fb58187613679565b905067ffffffffffffffff851660e0840152828103610100840152613fda8185613679565b9c9b505050505050505050505050565b8051613b6281613b67565b805161ffff81168114613b6257600080fd5b8051613b6281613623565b600082601f83011261402357600080fd5b8151614031613a0e826139a1565b81815284602083860101111561404657600080fd5b6105fd826020830160208701613655565b80516001600160601b0381168114613b6257600080fd5b600082601f83011261407f57600080fd5b8151602061408f613a0e836139c9565b82815260059290921b840181019181810190868411156140ae57600080fd5b8286015b84811015613a4d57805167ffffffffffffffff808211156140d257600080fd5b90880190601f196040838c03820112156140eb57600080fd5b6140f36138ba565b878401518381111561410457600080fd5b8401610100818e038401121561411957600080fd5b6141216138e3565b925088810151835261413560408201613ff5565b8984015261414560608201614007565b604084015260808101518481111561415c57600080fd5b61416a8e8b83850101614012565b60608501525061417c60a08201614057565b608084015260c081015160a084015260e081015160c084015261010081015160e0840152508181526141b060408501614057565b8189015286525050509183019183016140b2565b80516001600160c01b0381168114613b6257600080fd5b6000602082840312156141ed57600080fd5b815167ffffffffffffffff8082111561420557600080fd5b9083019060a0828603121561421957600080fd5b614221613907565b82518281111561423057600080fd5b8301601f8101871361424157600080fd5b805161424f613a0e826139c9565b8082825260208201915060208360051b85010192508983111561427157600080fd5b602084015b838110156143715780518781111561428d57600080fd5b850160a0818d03601f190112156142a357600080fd5b6142ab61392a565b60208201516142b981613b67565b815260408201516142c981613c76565b60208201526040828e03605f190112156142e257600080fd5b6142ea61394d565b8d607f8401126142f957600080fd5b6143016138ba565b808f60a08601111561431257600080fd5b606085015b60a08601811015614332578051835260209283019201614317565b50825250604082015260a08201518981111561434d57600080fd5b61435c8e60208386010161406e565b60608301525084525060209283019201614276565b50845250614384915050602084016141c4565b602082015261439560408401613fea565b60408201526143a660608401613fea565b60608201526080830151608082015280935050505092915050565b600063ffffffff8083168181036143da576143da613d47565b6001019392505050565b600081518084526020808501808196508360051b8101915082860160005b858110156144c857828403895281516040815181875280518288015287810151606061ffff8216818a0152838301519350608091506001600160a01b038416828a01528083015193505061010060a081818b01526144646101408b0186613679565b9284015192945060c06144818b8201856001600160601b03169052565b9084015160e08b81019190915290840151918a01919091529091015161012088015250908601516001600160601b0316948601949094529784019790840190600101614402565b5091979650505050505050565b600060a080830181845280895180835260c08601915060c08160051b87010192506020808c016000805b848110156145865789870360bf190186528251805167ffffffffffffffff1688528481015162ffffff1685890152604080820151519084908a015b600282101561455957825181529187019160019190910190870161453a565b5050506060015160808801899052614573888a01826143e4565b97505094830194918301916001016144ff565b5050508395506145a08188018c6001600160c01b03169052565b50505050506145bb604083018667ffffffffffffffff169052565b67ffffffffffffffff939093166060820152608001529392505050565b600060208083850312156145eb57600080fd5b825167ffffffffffffffff81111561460257600080fd5b8301601f8101851361461357600080fd5b8051614621613a0e826139c9565b81815260079190911b8201830190838101908783111561464057600080fd5b928401925b828410156146a9576080848903121561465e5760008081fd5b61466661392a565b845161467181613b67565b81528486015161468081613c76565b818701526040858101519082015260608086015190820152825260809093019290840190614645565b979650505050505050565b634e487b7160e01b600052601260045260246000fd5b6040815260006146dd6040830185613f10565b82810360208481019190915284518083528582019282019060005b81811015614714578451835293830193918301916001016146f8565b5090979650505050505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b16604085015281606085015261475b8285018b613f10565b9150838203608085015261476f828a613f10565b915060ff881660a085015283820360c085015261478c8288613679565b90861660e08501528381036101008501529050613fda8185613679565b6040815260006147bd604083018587613e1d565b9050826020830152949350505050565b60006101008083850312156147e157600080fd5b83601f8401126147f057600080fd5b6147f86138e3565b90830190808583111561480a57600080fd5b845b8381101561482d57803561481f81613c76565b83526020928301920161480c565b5095945050505050565b6101008101818360005b600881101561486557815162ffffff16835260209283019290910190600101614841565b5050509291505056fea164736f6c6343000813000a",
}

var VRFBeaconABI = VRFBeaconMetaData.ABI

var VRFBeaconBin = VRFBeaconMetaData.Bin

func DeployVRFBeacon(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, coordinator common.Address, keyProvider common.Address, keyID [32]byte) (common.Address, *types.Transaction, *VRFBeacon, error) {
	parsed, err := VRFBeaconMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconBin), backend, link, coordinator, keyProvider, keyID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFBeacon{VRFBeaconCaller: VRFBeaconCaller{contract: contract}, VRFBeaconTransactor: VRFBeaconTransactor{contract: contract}, VRFBeaconFilterer: VRFBeaconFilterer{contract: contract}}, nil
}

type VRFBeacon struct {
	VRFBeaconCaller
	VRFBeaconTransactor
	VRFBeaconFilterer
}

type VRFBeaconCaller struct {
	contract *bind.BoundContract
}

type VRFBeaconTransactor struct {
	contract *bind.BoundContract
}

type VRFBeaconFilterer struct {
	contract *bind.BoundContract
}

type VRFBeaconSession struct {
	Contract     *VRFBeacon
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFBeaconCallerSession struct {
	Contract *VRFBeaconCaller
	CallOpts bind.CallOpts
}

type VRFBeaconTransactorSession struct {
	Contract     *VRFBeaconTransactor
	TransactOpts bind.TransactOpts
}

type VRFBeaconRaw struct {
	Contract *VRFBeacon
}

type VRFBeaconCallerRaw struct {
	Contract *VRFBeaconCaller
}

type VRFBeaconTransactorRaw struct {
	Contract *VRFBeaconTransactor
}

func NewVRFBeacon(address common.Address, backend bind.ContractBackend) (*VRFBeacon, error) {
	contract, err := bindVRFBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFBeacon{VRFBeaconCaller: VRFBeaconCaller{contract: contract}, VRFBeaconTransactor: VRFBeaconTransactor{contract: contract}, VRFBeaconFilterer: VRFBeaconFilterer{contract: contract}}, nil
}

func NewVRFBeaconCaller(address common.Address, caller bind.ContractCaller) (*VRFBeaconCaller, error) {
	contract, err := bindVRFBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCaller{contract: contract}, nil
}

func NewVRFBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFBeaconTransactor, error) {
	contract, err := bindVRFBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconTransactor{contract: contract}, nil
}

func NewVRFBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFBeaconFilterer, error) {
	contract, err := bindVRFBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconFilterer{contract: contract}, nil
}

func bindVRFBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFBeaconABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_VRFBeacon *VRFBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeacon.Contract.VRFBeaconCaller.contract.Call(opts, result, method, params...)
}

func (_VRFBeacon *VRFBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.Contract.VRFBeaconTransactor.contract.Transfer(opts)
}

func (_VRFBeacon *VRFBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeacon.Contract.VRFBeaconTransactor.contract.Transact(opts, method, params...)
}

func (_VRFBeacon *VRFBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFBeacon.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFBeacon *VRFBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.Contract.contract.Transfer(opts)
}

func (_VRFBeacon *VRFBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFBeacon.Contract.contract.Transact(opts, method, params...)
}

func (_VRFBeacon *VRFBeaconCaller) NUMCONFDELAYS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "NUM_CONF_DELAYS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeacon.Contract.NUMCONFDELAYS(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) NUMCONFDELAYS() (uint8, error) {
	return _VRFBeacon.Contract.NUMCONFDELAYS(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) GetBilling(opts *bind.CallOpts) (struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
}, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "getBilling")

	outstruct := new(struct {
		MaximumGasPrice     uint64
		ReasonableGasPrice  uint64
		ObservationPayment  uint64
		TransmissionPayment uint64
		AccountingGas       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPrice = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReasonableGasPrice = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ObservationPayment = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.TransmissionPayment = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFBeacon *VRFBeaconSession) GetBilling() (struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
}, error) {
	return _VRFBeacon.Contract.GetBilling(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) GetBilling() (struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
}, error) {
	return _VRFBeacon.Contract.GetBilling(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) GetBillingAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "getBillingAccessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeacon.Contract.GetBillingAccessController(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) GetBillingAccessController() (common.Address, error) {
	return _VRFBeacon.Contract.GetBillingAccessController(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) ICoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "i_coordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) ICoordinator() (common.Address, error) {
	return _VRFBeacon.Contract.ICoordinator(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) ICoordinator() (common.Address, error) {
	return _VRFBeacon.Contract.ICoordinator(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) ILink() (common.Address, error) {
	return _VRFBeacon.Contract.ILink(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) ILink() (common.Address, error) {
	return _VRFBeacon.Contract.ILink(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "latestConfigDetails")

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

func (_VRFBeacon *VRFBeaconSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRFBeacon.Contract.LatestConfigDetails(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _VRFBeacon.Contract.LatestConfigDetails(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

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

func (_VRFBeacon *VRFBeaconSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRFBeacon.Contract.LatestConfigDigestAndEpoch(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _VRFBeacon.Contract.LatestConfigDigestAndEpoch(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeacon.Contract.LinkAvailableForPayment(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _VRFBeacon.Contract.LinkAvailableForPayment(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) OracleObservationCount(opts *bind.CallOpts, transmitterAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "oracleObservationCount", transmitterAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _VRFBeacon.Contract.OracleObservationCount(&_VRFBeacon.CallOpts, transmitterAddress)
}

func (_VRFBeacon *VRFBeaconCallerSession) OracleObservationCount(transmitterAddress common.Address) (uint32, error) {
	return _VRFBeacon.Contract.OracleObservationCount(&_VRFBeacon.CallOpts, transmitterAddress)
}

func (_VRFBeacon *VRFBeaconCaller) OwedPayment(opts *bind.CallOpts, transmitterAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "owedPayment", transmitterAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeacon.Contract.OwedPayment(&_VRFBeacon.CallOpts, transmitterAddress)
}

func (_VRFBeacon *VRFBeaconCallerSession) OwedPayment(transmitterAddress common.Address) (*big.Int, error) {
	return _VRFBeacon.Contract.OwedPayment(&_VRFBeacon.CallOpts, transmitterAddress)
}

func (_VRFBeacon *VRFBeaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) Owner() (common.Address, error) {
	return _VRFBeacon.Contract.Owner(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) Owner() (common.Address, error) {
	return _VRFBeacon.Contract.Owner(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) SKeyID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "s_keyID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) SKeyID() ([32]byte, error) {
	return _VRFBeacon.Contract.SKeyID(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) SKeyID() ([32]byte, error) {
	return _VRFBeacon.Contract.SKeyID(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) SKeyProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "s_keyProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) SKeyProvider() (common.Address, error) {
	return _VRFBeacon.Contract.SKeyProvider(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) SKeyProvider() (common.Address, error) {
	return _VRFBeacon.Contract.SKeyProvider(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) SProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "s_provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeacon.Contract.SProvingKeyHash(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) SProvingKeyHash() ([32]byte, error) {
	return _VRFBeacon.Contract.SProvingKeyHash(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFBeacon.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFBeacon *VRFBeaconSession) TypeAndVersion() (string, error) {
	return _VRFBeacon.Contract.TypeAndVersion(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconCallerSession) TypeAndVersion() (string, error) {
	return _VRFBeacon.Contract.TypeAndVersion(&_VRFBeacon.CallOpts)
}

func (_VRFBeacon *VRFBeaconTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeacon *VRFBeaconSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptOwnership(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptOwnership(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactor) AcceptPayeeship(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "acceptPayeeship", transmitter)
}

func (_VRFBeacon *VRFBeaconSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptPayeeship(&_VRFBeacon.TransactOpts, transmitter)
}

func (_VRFBeacon *VRFBeaconTransactorSession) AcceptPayeeship(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.AcceptPayeeship(&_VRFBeacon.TransactOpts, transmitter)
}

func (_VRFBeacon *VRFBeaconTransactor) ExposeType(opts *bind.TransactOpts, arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "exposeType", arg0)
}

func (_VRFBeacon *VRFBeaconSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeacon.Contract.ExposeType(&_VRFBeacon.TransactOpts, arg0)
}

func (_VRFBeacon *VRFBeaconTransactorSession) ExposeType(arg0 VRFBeaconReportReport) (*types.Transaction, error) {
	return _VRFBeacon.Contract.ExposeType(&_VRFBeacon.TransactOpts, arg0)
}

func (_VRFBeacon *VRFBeaconTransactor) KeyGenerated(opts *bind.TransactOpts, kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "keyGenerated", kd)
}

func (_VRFBeacon *VRFBeaconSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeacon.Contract.KeyGenerated(&_VRFBeacon.TransactOpts, kd)
}

func (_VRFBeacon *VRFBeaconTransactorSession) KeyGenerated(kd KeyDataStructKeyData) (*types.Transaction, error) {
	return _VRFBeacon.Contract.KeyGenerated(&_VRFBeacon.TransactOpts, kd)
}

func (_VRFBeacon *VRFBeaconTransactor) NewKeyRequested(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "newKeyRequested")
}

func (_VRFBeacon *VRFBeaconSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeacon.Contract.NewKeyRequested(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactorSession) NewKeyRequested() (*types.Transaction, error) {
	return _VRFBeacon.Contract.NewKeyRequested(&_VRFBeacon.TransactOpts)
}

func (_VRFBeacon *VRFBeaconTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setBilling", maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeacon *VRFBeaconSession) SetBilling(maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBilling(&_VRFBeacon.TransactOpts, maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetBilling(maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBilling(&_VRFBeacon.TransactOpts, maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeacon *VRFBeaconTransactor) SetBillingAccessController(opts *bind.TransactOpts, _billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setBillingAccessController", _billingAccessController)
}

func (_VRFBeacon *VRFBeaconSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBillingAccessController(&_VRFBeacon.TransactOpts, _billingAccessController)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetBillingAccessController(_billingAccessController common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetBillingAccessController(&_VRFBeacon.TransactOpts, _billingAccessController)
}

func (_VRFBeacon *VRFBeaconTransactor) SetConfig(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setConfig", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeacon *VRFBeaconSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetConfig(&_VRFBeacon.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetConfig(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetConfig(&_VRFBeacon.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_VRFBeacon *VRFBeaconTransactor) SetPayees(opts *bind.TransactOpts, transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "setPayees", transmitters, payees)
}

func (_VRFBeacon *VRFBeaconSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetPayees(&_VRFBeacon.TransactOpts, transmitters, payees)
}

func (_VRFBeacon *VRFBeaconTransactorSession) SetPayees(transmitters []common.Address, payees []common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.SetPayees(&_VRFBeacon.TransactOpts, transmitters, payees)
}

func (_VRFBeacon *VRFBeaconTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeacon *VRFBeaconSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferOwnership(&_VRFBeacon.TransactOpts, to)
}

func (_VRFBeacon *VRFBeaconTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferOwnership(&_VRFBeacon.TransactOpts, to)
}

func (_VRFBeacon *VRFBeaconTransactor) TransferPayeeship(opts *bind.TransactOpts, transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "transferPayeeship", transmitter, proposed)
}

func (_VRFBeacon *VRFBeaconSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferPayeeship(&_VRFBeacon.TransactOpts, transmitter, proposed)
}

func (_VRFBeacon *VRFBeaconTransactorSession) TransferPayeeship(transmitter common.Address, proposed common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.TransferPayeeship(&_VRFBeacon.TransactOpts, transmitter, proposed)
}

func (_VRFBeacon *VRFBeaconTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_VRFBeacon *VRFBeaconSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.Transmit(&_VRFBeacon.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeacon *VRFBeaconTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _VRFBeacon.Contract.Transmit(&_VRFBeacon.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_VRFBeacon *VRFBeaconTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "withdrawFunds", recipient, amount)
}

func (_VRFBeacon *VRFBeaconSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawFunds(&_VRFBeacon.TransactOpts, recipient, amount)
}

func (_VRFBeacon *VRFBeaconTransactorSession) WithdrawFunds(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawFunds(&_VRFBeacon.TransactOpts, recipient, amount)
}

func (_VRFBeacon *VRFBeaconTransactor) WithdrawPayment(opts *bind.TransactOpts, transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.contract.Transact(opts, "withdrawPayment", transmitter)
}

func (_VRFBeacon *VRFBeaconSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawPayment(&_VRFBeacon.TransactOpts, transmitter)
}

func (_VRFBeacon *VRFBeaconTransactorSession) WithdrawPayment(transmitter common.Address) (*types.Transaction, error) {
	return _VRFBeacon.Contract.WithdrawPayment(&_VRFBeacon.TransactOpts, transmitter)
}

type VRFBeaconBillingAccessControllerSetIterator struct {
	Event *VRFBeaconBillingAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingAccessControllerSet)
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
		it.Event = new(VRFBeaconBillingAccessControllerSet)
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

func (it *VRFBeaconBillingAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingAccessControllerSet struct {
	Old     common.Address
	Current common.Address
	Raw     types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterBillingAccessControllerSet(opts *bind.FilterOpts) (*VRFBeaconBillingAccessControllerSetIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingAccessControllerSetIterator{contract: _VRFBeacon.contract, event: "BillingAccessControllerSet", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchBillingAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "BillingAccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingAccessControllerSet)
				if err := _VRFBeacon.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseBillingAccessControllerSet(log types.Log) (*VRFBeaconBillingAccessControllerSet, error) {
	event := new(VRFBeaconBillingAccessControllerSet)
	if err := _VRFBeacon.contract.UnpackLog(event, "BillingAccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconBillingSetIterator struct {
	Event *VRFBeaconBillingSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconBillingSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconBillingSet)
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
		it.Event = new(VRFBeaconBillingSet)
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

func (it *VRFBeaconBillingSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconBillingSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconBillingSet struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
	Raw                 types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterBillingSet(opts *bind.FilterOpts) (*VRFBeaconBillingSetIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconBillingSetIterator{contract: _VRFBeacon.contract, event: "BillingSet", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchBillingSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconBillingSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "BillingSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconBillingSet)
				if err := _VRFBeacon.contract.UnpackLog(event, "BillingSet", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseBillingSet(log types.Log) (*VRFBeaconBillingSet, error) {
	event := new(VRFBeaconBillingSet)
	if err := _VRFBeacon.contract.UnpackLog(event, "BillingSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconConfigSetIterator struct {
	Event *VRFBeaconConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconConfigSet)
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
		it.Event = new(VRFBeaconConfigSet)
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

func (it *VRFBeaconConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconConfigSet struct {
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

func (_VRFBeacon *VRFBeaconFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFBeaconConfigSetIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconConfigSetIterator{contract: _VRFBeacon.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconConfigSet)
				if err := _VRFBeacon.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseConfigSet(log types.Log) (*VRFBeaconConfigSet, error) {
	event := new(VRFBeaconConfigSet)
	if err := _VRFBeacon.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconNewTransmissionIterator struct {
	Event *VRFBeaconNewTransmission

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconNewTransmissionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconNewTransmission)
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
		it.Event = new(VRFBeaconNewTransmission)
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

func (it *VRFBeaconNewTransmissionIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconNewTransmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconNewTransmission struct {
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*VRFBeaconNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconNewTransmissionIterator{contract: _VRFBeacon.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconNewTransmission)
				if err := _VRFBeacon.contract.UnpackLog(event, "NewTransmission", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseNewTransmission(log types.Log) (*VRFBeaconNewTransmission, error) {
	event := new(VRFBeaconNewTransmission)
	if err := _VRFBeacon.contract.UnpackLog(event, "NewTransmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOraclePaidIterator struct {
	Event *VRFBeaconOraclePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOraclePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOraclePaid)
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
		it.Event = new(VRFBeaconOraclePaid)
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

func (it *VRFBeaconOraclePaidIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOraclePaid struct {
	Transmitter common.Address
	Payee       common.Address
	Amount      *big.Int
	LinkToken   common.Address
	Raw         types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOraclePaid(opts *bind.FilterOpts, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (*VRFBeaconOraclePaidIterator, error) {

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

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOraclePaidIterator{contract: _VRFBeacon.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *VRFBeaconOraclePaid, transmitter []common.Address, payee []common.Address, linkToken []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OraclePaid", transmitterRule, payeeRule, linkTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOraclePaid)
				if err := _VRFBeacon.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOraclePaid(log types.Log) (*VRFBeaconOraclePaid, error) {
	event := new(VRFBeaconOraclePaid)
	if err := _VRFBeacon.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOutputsServedIterator struct {
	Event *VRFBeaconOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOutputsServed)
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
		it.Event = new(VRFBeaconOutputsServed)
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

func (it *VRFBeaconOutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOutputsServed struct {
	RecentBlockHeight  uint64
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFBeaconOutputsServedIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOutputsServedIterator{contract: _VRFBeacon.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFBeaconOutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOutputsServed)
				if err := _VRFBeacon.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOutputsServed(log types.Log) (*VRFBeaconOutputsServed, error) {
	event := new(VRFBeaconOutputsServed)
	if err := _VRFBeacon.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOwnershipTransferRequestedIterator struct {
	Event *VRFBeaconOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOwnershipTransferRequested)
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
		it.Event = new(VRFBeaconOwnershipTransferRequested)
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

func (it *VRFBeaconOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOwnershipTransferRequestedIterator{contract: _VRFBeacon.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOwnershipTransferRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconOwnershipTransferRequested, error) {
	event := new(VRFBeaconOwnershipTransferRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconOwnershipTransferredIterator struct {
	Event *VRFBeaconOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOwnershipTransferred)
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
		it.Event = new(VRFBeaconOwnershipTransferred)
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

func (it *VRFBeaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOwnershipTransferredIterator{contract: _VRFBeacon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOwnershipTransferred)
				if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseOwnershipTransferred(log types.Log) (*VRFBeaconOwnershipTransferred, error) {
	event := new(VRFBeaconOwnershipTransferred)
	if err := _VRFBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconPayeeshipTransferRequestedIterator struct {
	Event *VRFBeaconPayeeshipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconPayeeshipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconPayeeshipTransferRequested)
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
		it.Event = new(VRFBeaconPayeeshipTransferRequested)
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

func (it *VRFBeaconPayeeshipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconPayeeshipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconPayeeshipTransferRequested struct {
	Transmitter common.Address
	Current     common.Address
	Proposed    common.Address
	Raw         types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterPayeeshipTransferRequested(opts *bind.FilterOpts, transmitter []common.Address, current []common.Address, proposed []common.Address) (*VRFBeaconPayeeshipTransferRequestedIterator, error) {

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

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconPayeeshipTransferRequestedIterator{contract: _VRFBeacon.contract, event: "PayeeshipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchPayeeshipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconPayeeshipTransferRequested, transmitter []common.Address, current []common.Address, proposed []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "PayeeshipTransferRequested", transmitterRule, currentRule, proposedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconPayeeshipTransferRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParsePayeeshipTransferRequested(log types.Log) (*VRFBeaconPayeeshipTransferRequested, error) {
	event := new(VRFBeaconPayeeshipTransferRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconPayeeshipTransferredIterator struct {
	Event *VRFBeaconPayeeshipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconPayeeshipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconPayeeshipTransferred)
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
		it.Event = new(VRFBeaconPayeeshipTransferred)
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

func (it *VRFBeaconPayeeshipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconPayeeshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconPayeeshipTransferred struct {
	Transmitter common.Address
	Previous    common.Address
	Current     common.Address
	Raw         types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterPayeeshipTransferred(opts *bind.FilterOpts, transmitter []common.Address, previous []common.Address, current []common.Address) (*VRFBeaconPayeeshipTransferredIterator, error) {

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

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconPayeeshipTransferredIterator{contract: _VRFBeacon.contract, event: "PayeeshipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchPayeeshipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconPayeeshipTransferred, transmitter []common.Address, previous []common.Address, current []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "PayeeshipTransferred", transmitterRule, previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconPayeeshipTransferred)
				if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParsePayeeshipTransferred(log types.Log) (*VRFBeaconPayeeshipTransferred, error) {
	event := new(VRFBeaconPayeeshipTransferred)
	if err := _VRFBeacon.contract.UnpackLog(event, "PayeeshipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomWordsFulfilledIterator struct {
	Event *VRFBeaconRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomWordsFulfilled)
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
		it.Event = new(VRFBeaconRandomWordsFulfilled)
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

func (it *VRFBeaconRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomWordsFulfilled struct {
	RequestIDs            []*big.Int
	SuccessfulFulfillment []byte
	TruncatedErrorData    [][]byte
	Raw                   types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts) (*VRFBeaconRandomWordsFulfilledIterator, error) {

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomWordsFulfilledIterator{contract: _VRFBeacon.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomWordsFulfilled) (event.Subscription, error) {

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomWordsFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomWordsFulfilled)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFBeaconRandomWordsFulfilled, error) {
	event := new(VRFBeaconRandomWordsFulfilled)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomnessFulfillmentRequestedIterator struct {
	Event *VRFBeaconRandomnessFulfillmentRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomnessFulfillmentRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomnessFulfillmentRequested)
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
		it.Event = new(VRFBeaconRandomnessFulfillmentRequested)
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

func (it *VRFBeaconRandomnessFulfillmentRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomnessFulfillmentRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomnessFulfillmentRequested struct {
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
	Raw                    types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomnessFulfillmentRequestedIterator{contract: _VRFBeacon.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomnessFulfillmentRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomnessFulfillmentRequested(log types.Log) (*VRFBeaconRandomnessFulfillmentRequested, error) {
	event := new(VRFBeaconRandomnessFulfillmentRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessFulfillmentRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomnessRedeemedIterator struct {
	Event *VRFBeaconRandomnessRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomnessRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomnessRedeemed)
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
		it.Event = new(VRFBeaconRandomnessRedeemed)
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

func (it *VRFBeaconRandomnessRedeemedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomnessRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomnessRedeemed struct {
	RequestID *big.Int
	Requester common.Address
	SubID     *big.Int
	Raw       types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconRandomnessRedeemedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomnessRedeemedIterator{contract: _VRFBeacon.contract, event: "RandomnessRedeemed", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomnessRedeemed)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomnessRedeemed(log types.Log) (*VRFBeaconRandomnessRedeemed, error) {
	event := new(VRFBeaconRandomnessRedeemed)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconRandomnessRequestedIterator struct {
	Event *VRFBeaconRandomnessRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconRandomnessRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconRandomnessRequested)
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
		it.Event = new(VRFBeaconRandomnessRequested)
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

func (it *VRFBeaconRandomnessRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconRandomnessRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconRandomnessRequested struct {
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	CostJuels              *big.Int
	Raw                    types.Log
}

func (_VRFBeacon *VRFBeaconFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconRandomnessRequestedIterator{contract: _VRFBeacon.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeacon *VRFBeaconFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeacon.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconRandomnessRequested)
				if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
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

func (_VRFBeacon *VRFBeaconFilterer) ParseRandomnessRequested(log types.Log) (*VRFBeaconRandomnessRequested, error) {
	event := new(VRFBeaconRandomnessRequested)
	if err := _VRFBeacon.contract.UnpackLog(event, "RandomnessRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var VRFBeaconDKGClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDKG\",\"name\":\"_keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyProvider\",\"outputs\":[{\"internalType\":\"contractDKG\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161048338038061048383398101604081905261002f91610058565b600080546001600160a01b0319166001600160a01b039390931692909217909155608052610092565b6000806040838503121561006b57600080fd5b82516001600160a01b038116811461008257600080fd5b6020939093015192949293505050565b6080516103d76100ac600039600060be01526103d76000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063bf2732c711610050578063bf2732c7146100a6578063cc31f7dd146100b9578063d57fc45a146100ee57600080fd5b806355e487491461006c5780635f27026f14610076575b600080fd5b6100746100f7565b005b600054610089906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100746100b43660046102b8565b610141565b6100e07f000000000000000000000000000000000000000000000000000000000000000081565b60405190815260200161009d565b6100e060015481565b6000546001600160a01b03163381146101395760405163292f4fb560e01b81523360048201526001600160a01b03821660248201526044015b60405180910390fd5b506000600155565b6000546001600160a01b031633811461017e5760405163292f4fb560e01b81523360048201526001600160a01b0382166024820152604401610130565b8151604051610190919060200161039b565b60408051601f1981840301815291905280516020909101206001555050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610201576102016101af565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610230576102306101af565b604052919050565b600082601f83011261024957600080fd5b8135602067ffffffffffffffff821115610265576102656101af565b8160051b610274828201610207565b928352848101820192828101908785111561028e57600080fd5b83870192505b848310156102ad57823582529183019190830190610294565b979650505050505050565b600060208083850312156102cb57600080fd5b823567ffffffffffffffff808211156102e357600080fd5b90840190604082870312156102f757600080fd5b6102ff6101de565b82358281111561030e57600080fd5b8301601f8101881361031f57600080fd5b803583811115610331576103316101af565b610343601f8201601f19168701610207565b818152898783850101111561035757600080fd5b81878401888301376000878383010152808452505050838301358281111561037e57600080fd5b61038a88828601610238565b948201949094529695505050505050565b6000825160005b818110156103bc57602081860181015185830152016103a2565b50600092019182525091905056fea164736f6c6343000813000a",
}

var VRFBeaconDKGClientABI = VRFBeaconDKGClientMetaData.ABI

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

func (_VRFBeaconDKGClient *VRFBeaconDKGClientCaller) SKeyProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconDKGClient.contract.Call(opts, &out, "s_keyProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientSession) SKeyProvider() (common.Address, error) {
	return _VRFBeaconDKGClient.Contract.SKeyProvider(&_VRFBeaconDKGClient.CallOpts)
}

func (_VRFBeaconDKGClient *VRFBeaconDKGClientCallerSession) SKeyProvider() (common.Address, error) {
	return _VRFBeaconDKGClient.Contract.SKeyProvider(&_VRFBeaconDKGClient.CallOpts)
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

var VRFBeaconOCRMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"coordinator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"}],\"name\":\"CalldataLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAcceptPayeeship\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectNumberOfFaultyOracles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTransmitters\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numPayees\",\"type\":\"uint256\"}],\"name\":\"IncorrectNumberOfPayees\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expectedNumSignatures\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ssLength\",\"type\":\"uint256\"}],\"name\":\"IncorrectNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeftGasExceedsInitialGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeOwnerOrBillingAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numFaultyOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"}],\"name\":\"NumberOfFaultyOraclesTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OnchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"OnlyActiveSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"OnlyActiveTransmitters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCurrentPayee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"existingPayee\",\"type\":\"address\"}],\"name\":\"PayeeAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedSignerAddress\",\"type\":\"address\"}],\"name\":\"RepeatedSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"repeatedTransmitterAddress\",\"type\":\"address\"}],\"name\":\"RepeatedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"}],\"name\":\"SeenReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numSigners\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numTransmitters\",\"type\":\"uint256\"}],\"name\":\"SignersTransmittersMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maxOracles\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"providedOracles\",\"type\":\"uint256\"}],\"name\":\"TooManyOracles\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ocrVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconTypes.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_coordinator\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"maximumGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"observationPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"transmissionPayment\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200481c3803806200481c8339810160408190526200003491620001a2565b818133806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620000de565b5050506001600160a01b039182166080521660a05250620001e19050565b336001600160a01b03821603620001385760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b03811681146200019f57600080fd5b50565b60008060408385031215620001b657600080fd5b8251620001c38162000189565b6020840151909250620001d68162000189565b809150509250929050565b60805160a05161458962000293600039600081816102c001528181610e4a01528181610eef01528181610fbf0152818161108d0152818161112b01528181611d6301528181611e0801528181611fa30152818161236b01528181612456015281816128990152818161293e01528181612e4701526132f901526000818161025101528181610f1d015281816110ba01528181611e36015281816120220152818161296c0152612d5701526145896000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c8063b121e147116100e3578063e3d0e7121161008c578063eb5dcd6c11610066578063eb5dcd6c146103f4578063f2fde38b14610407578063fbffd2c11461041a57600080fd5b8063e3d0e712146103a6578063e4902f82146103b9578063e53bbc9a146103e157600080fd5b8063c1075329116100bd578063c10753291461037a578063c4c92b371461038d578063d09dc3391461039e57600080fd5b8063b121e14714610343578063b1dc65a414610356578063b8be03cd1461036957600080fd5b806381ff7048116101455780638da5cb5b1161011f5780638da5cb5b146102f55780639c849b3014610306578063afcb95d71461031957600080fd5b806381ff70481461028b5780638a1b1772146102bb5780638ac28d5a146102e257600080fd5b80632f7527cc116101765780632f7527cc1461022857806379ba5097146102425780637d253aff1461024c57600080fd5b80630eafb25b1461019257806329937268146101b8575b600080fd5b6101a56101a03660046134a6565b61042d565b6040519081526020015b60405180910390f35b6002546003546040805165010000000000840467ffffffffffffffff90811682526d0100000000000000000000000000850481166020830152600160a81b909404841691810191909152918116606083015268010000000000000000900462ffffff16608082015260a0016101af565b610230600881565b60405160ff90911681526020016101af565b61024a610526565b005b6102737f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101af565b6004546005546040805163ffffffff808516825264010000000090940490931660208401528201526060016101af565b6102737f000000000000000000000000000000000000000000000000000000000000000081565b61024a6102f03660046134a6565b6105dc565b6000546001600160a01b0316610273565b61024a61031436600461350f565b61061f565b6005546006546040805160008152602081019390935263ffffffff909116908201526060016101af565b61024a6103513660046134a6565b6107f1565b61024a6103643660046135bd565b6108b7565b61024a610377366004613674565b50565b61024a6103883660046136af565b610d8e565b6012546001600160a01b0316610273565b6101a5611076565b61024a6103b4366004613712565b6111d6565b6103cc6103c73660046134a6565b61193f565b60405163ffffffff90911681526020016101af565b61024a6103ef366004613811565b6119f3565b61024a610402366004613882565b611bdf565b61024a6104153660046134a6565b611cd3565b61024a6104283660046134a6565b611ce4565b6001600160a01b03811660009081526008602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b0316918101919091529061048f5750600092915050565b6002546020820151600091600160a81b900467ffffffffffffffff1690600c9060ff16601f81106104c2576104c26138bb565b6008810491909101546002546104f39260071660040261010090810a90920463ffffffff90811692909104166138e7565b63ffffffff16610503919061390b565b905081604001516001600160601b03168161051e9190613922565b949350505050565b6001546001600160a01b031633146105855760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b0381811660009081526010602052604090205416331461061657604051633738e30960e21b815260040160405180910390fd5b61037781611cf5565b6106276120a1565b82811461066a576040517f36d20459000000000000000000000000000000000000000000000000000000008152600481018490526024810182905260440161057c565b60005b838110156107ea576000858583818110610689576106896138bb565b905060200201602081019061069e91906134a6565b905060008484848181106106b4576106b46138bb565b90506020020160208101906106c991906134a6565b6001600160a01b038084166000908152601060205260409020549192501680158015816107085750826001600160a01b0316826001600160a01b031614155b15610752576040517febdf17560000000000000000000000000000000000000000000000000000000081526001600160a01b0380861660048301528316602482015260440161057c565b6001600160a01b03848116600090815260106020526040902080546001600160a01b031916858316908117909155908316146107d357826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b5050505080806107e290613935565b91505061066d565b5050505050565b6001600160a01b03818116600090815260116020526040902054163314610844576040517f9d12ec4f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381811660008181526010602090815260408083208054336001600160a01b031980831682179093556011909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a6040805160e08101825260025460ff8082168352610100820463ffffffff1660208085019190915265010000000000830467ffffffffffffffff908116858701526d0100000000000000000000000000840481166060860152600160a81b9093048316608085015260035492831660a08501526801000000000000000090920462ffffff1660c08401523360009081526008835293909320549394509092908c01359116610996576040517fb1c1f68e00000000000000000000000000000000000000000000000000000000815233600482015260240161057c565b6005548b35146109e0576005546040517f93df584c00000000000000000000000000000000000000000000000000000000815260048101919091528b35602482015260440161057c565b6109ee8a8a8a8a8a8a6120fd565b81516109fb90600161394e565b60ff1687141580610a0c5750868514155b15610a64578151610a1e90600161394e565b6040517ffc33647500000000000000000000000000000000000000000000000000000000815260ff9091166004820152602481018890526044810186905260640161057c565b60008a8a604051610a76929190613967565b604051908190038120610a8d918e90602001613977565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a811015610c285760006001858a8460208110610adc57610adc6138bb565b610ae991901a601b61394e565b8f8f86818110610afb57610afb6138bb565b905060200201358e8e87818110610b1457610b146138bb565b9050602002013560405160008152602001604052604051610b51949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015610b73573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526009602090815290849020838501909452925460ff8082161515808552610100909204169383019390935290955092509050610c01576040517f20fb74ee0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161057c565b826020015160080260ff166001901b84019350508080610c2090613935565b915050610abd565b5081827e010101010101010101010101010101010101010101010101010101010101011614610c83576040517fc103be2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505060008a8a604051610c98929190613967565b604080519182900390912060008181526007602052919091205490915060ff1615610cf2576040517f0b8d39d50000000000000000000000000000000000000000000000000000000081526004810182905260240161057c565b6000908152600760205260408120805460ff1916600117905580610d54848e836020020135858f8f8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061218d92505050565b6006805463ffffffff191663ffffffff600888901c161790559092509050610d7f848383883361255f565b50505050505050505050505050565b6000546001600160a01b03163314801590610e1c5750601254604051630d629b5f60e31b81526001600160a01b0390911690636b14daf890610dd990339060009036906004016139b4565b602060405180830381865afa158015610df6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1a91906139d7565b155b15610e3a57604051631809d98560e31b815260040160405180910390fd5b6000610e44612677565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ea6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eca91906139f9565b90506000610ed88284613922565b6040516370a0823160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610f64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8891906139f9565b905081811015610fb55760405163cf47918160e01b8152600481018290526024810183905260440161057c565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663f99b1d6887610ff8610ff28686613a12565b89612835565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b0390921660048301526024820152604401600060405180830381600087803b15801561105657600080fd5b505af115801561106a573d6000803e3d6000fd5b50505050505050505050565b6040516370a0823160e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116600483015260009182917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611101573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112591906139f9565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611187573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ab91906139f9565b905060006111b7612677565b9050816111c48285613a25565b6111ce9190613a25565b935050505090565b888787601f83111561121e576040517f809fc428000000000000000000000000000000000000000000000000000000008152601f60048201526024810184905260440161057c565b818314611261576040517f988a0804000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260440161057c565b61126c816003613a45565b60ff1683116112b3576040517ffda9db7800000000000000000000000000000000000000000000000000000000815260ff821660048201526024810184905260440161057c565b6112bf8160ff1661284f565b6112c76120a1565b60006040518060c001604052808f8f80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505081526020018d8d8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050509082525060ff8c1660208083019190915260408051601f8d0183900483028101830182528c8152920191908c908c908190840183828082843760009201919091525050509082525067ffffffffffffffff891660208083019190915260408051601f8a018390048302810183018252898152920191908990899081908401838280828437600092019190915250505091525090506113e9612889565b600a5460005b8181101561149a576000600a828154811061140c5761140c6138bb565b6000918252602082200154600b80546001600160a01b0390921693509084908110611439576114396138bb565b60009182526020808320909101546001600160a01b039485168352600982526040808420805461ffff1916905594168252600890529190912080546dffffffffffffffffffffffffffff19169055508061149281613935565b9150506113ef565b506114a7600a6000613366565b6114b3600b6000613366565b60005b82515181101561175f5760096000846000015183815181106114da576114da6138bb565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff161561156457825180518290811061151a5761151a6138bb565b60200260200101516040517f7451f83e00000000000000000000000000000000000000000000000000000000815260040161057c91906001600160a01b0391909116815260200190565b604080518082019091526001815260ff821660208201528351805160099160009185908110611595576115956138bb565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051600892919084908110611600576116006138bb565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff161561168c5782602001518181518110611642576116426138bb565b60200260200101516040517fe8d2989900000000000000000000000000000000000000000000000000000000815260040161057c91906001600160a01b0391909116815260200190565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b031681525060086000856020015184815181106116d1576116d16138bb565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516001600160601b031662010000026dffffffffffffffffffffffff00001960ff959095166101000261ff00199315159390931661ffff199094169390931791909117929092161790558061175781613935565b9150506114b6565b508151805161177691600a91602090910190613384565b50602080830151805161178d92600b920190613384565b5060408201516002805460ff191660ff909216919091179055600454640100000000900463ffffffff166117bf612ebe565b6004805463ffffffff9283166401000000000267ffffffff0000000019821681179092556000926117f69281169116176001613a61565b905080600460006101000a81548163ffffffff021916908363ffffffff160217905550600061184a46308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a00151612f48565b9050806005819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e058360055484886000015189602001518a604001518b606001518c608001518d60a001516040516118ac99989796959493929190613b12565b60405180910390a1600254610100900463ffffffff1660005b86515181101561191f5781600c82601f81106118e3576118e36138bb565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550808061191790613935565b9150506118c5565b5061192a8e8e612fd5565b50505050505050505050505050505050505050565b6001600160a01b03811660009081526008602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906119a15750600092915050565b600c816020015160ff16601f81106119bb576119bb6138bb565b6008810491909101546002546119ec9260071660040261010090810a90920463ffffffff90811692909104166138e7565b9392505050565b6000546001600160a01b03163314801590611a815750601254604051630d629b5f60e31b81526001600160a01b0390911690636b14daf890611a3e90339060009036906004016139b4565b602060405180830381865afa158015611a5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7f91906139d7565b155b15611a9f57604051631809d98560e31b815260040160405180910390fd5b611aa7612889565b6002805467ffffffffffffffff858116600160a81b81027fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff8984166d010000000000000000000000000081027fffffffffffffffffffffff0000000000000000ffffffffffffffffffffffffff8d8716650100000000008102919091167fffffffffffffffffffffff00000000000000000000000000000000ffffffffff909816979097171791909116919091179094556003805462ffffff87166801000000000000000081026affffffffffffffffffffff199092169489169485179190911790915560408051948552602085019590955293830152606082015260808101919091527f49275ddcdfc9c0519b3d094308c8bf675f06070a754ce90c152163cb6e66e8a09060a00160405180910390a15050505050565b6001600160a01b03828116600090815260106020526040902054163314611c1957604051633738e30960e21b815260040160405180910390fd5b6001600160a01b0381163303611c5b576040517fb387a23800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03808316600090815260116020526040902080548383166001600160a01b031982168117909255909116908114611cce576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a45b505050565b611cdb6120a1565b61037781612fe3565b611cec6120a1565b6103778161308c565b6001600160a01b0381166000908152600860209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b031692810192909252611d52575050565b6000611d5d8361042d565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611dbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de391906139f9565b90506000611df18284613922565b6040516370a0823160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611e7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea191906139f9565b905081811015611ece5760405163cf47918160e01b8152600481018290526024810183905260440161057c565b8315612099576001600160a01b038681166000908152601060209081526040909120546002549188015192169161010090910463ffffffff1690600c9060ff16601f8110611f1e57611f1e6138bb565b6008808204909201805463ffffffff9485166004600790941684026101000a90810295021916939093179092556001600160a01b03808a16600090815260209290925260409182902080546dffffffffffffffffffffffff00001916905590517ff99b1d680000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000009091169163f99b1d6891611fee9185918a91016001600160a01b03929092168252602082015260400190565b600060405180830381600087803b15801561200857600080fd5b505af115801561201c573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316886001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8860405161208f91815260200190565b60405180910390a4505b505050505050565b6000546001600160a01b031633146120fb5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161057c565b565b600061210a82602061390b565b61211585602061390b565b61212188610144613922565b61212b9190613922565b6121359190613922565b612140906000613922565b9050368114612184576040517ff7b94f0a0000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161057c565b50505050505050565b6000806000838060200190518101906121a69190613ee9565b6020880180519192506121b8826140cf565b63ffffffff1663ffffffff168152505086600260008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548163ffffffff021916908363ffffffff16021790555060408201518160000160056101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600001600d6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160000160156101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160086101000a81548162ffffff021916908362ffffff16021790555090505060006123068260600151613102565b90508082608001511461236957608082015160608301516040517faed0afe500000000000000000000000000000000000000000000000000000000815260048101929092526024820183905267ffffffffffffffff16604482015260640161057c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316625bd524836000015184602001518560400151866060015187608001516040518663ffffffff1660e01b81526004016123d09594939291906141e3565b6000604051808303816000875af11580156123ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261241791908101906142e6565b5060408281015190517f05f4acc600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906305f4acc690602401600060405180830381600087803b1580156124a257600080fd5b505af11580156124b6573d6000803e3d6000fd5b505050508564ffffffffff16886020015163ffffffff167f27bf3f1077f091da6885751ba10f5775d06657fd59e47a6ab1f7635e5a115afe33856020015186604001518c60405161253e94939291906001600160a01b039490941684526001600160c01b0392909216602084015267ffffffffffffffff166040830152606082015260800190565b60405180910390a38160200151826040015193509350505094509492505050565b600061258b3a67ffffffffffffffff86161561257b5785612581565b87606001515b88604001516131de565b90506010360260005a905060006125b48663ffffffff1685858c60c0015162ffffff168661322f565b90506000670de0b6b3a76400006001600160c01b038a1683026001600160a01b03881660009081526008602052604090205460a08d01519290910492506201000090046001600160601b039081169167ffffffffffffffff168284010190811682111561262757505050505050506107ea565b6001600160a01b038816600090815260086020526040902080546001600160601b0390921662010000026dffffffffffffffffffffffff0000199092169190911790555050505050505050505050565b600080600b8054806020026020016040519081016040528092919081815260200182805480156126d057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116126b2575b50508351600254604080516103e0810191829052969750919561010090910463ffffffff169450600093509150600c90601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116127075790505050505050905060005b8381101561279a578181601f8110612767576127676138bb565b602002015161277690846138e7565b6127869063ffffffff1687613922565b95508061279281613935565b91505061274d565b506002546127b990600160a81b900467ffffffffffffffff168661390b565b945060005b8381101561282d57600860008683815181106127dc576127dc6138bb565b6020908102919091018101516001600160a01b0316825281019190915260400160002054612819906201000090046001600160601b031687613922565b95508061282581613935565b9150506127be565b505050505090565b600081831015612846575081612849565b50805b92915050565b80600003610377576040517fe77dba5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000612893612677565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663597d2f3c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061291991906139f9565b905060006129278284613922565b6040516370a0823160e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156129b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129d791906139f9565b905081811015612a045760405163cf47918160e01b8152600481018290526024810183905260440161057c565b600254604080516103e081019182905261010090920463ffffffff1691600091600c90601f908285855b82829054906101000a900463ffffffff1663ffffffff1681526020019060040190602082600301049283019260010382029150808411612a2e579050505050505090506000600b805480602002602001604051908101604052809291908181526020018280548015612ac957602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612aab575b5050505050905060008151905060008167ffffffffffffffff811115612af157612af1613ba8565b604051908082528060200260200182016040528015612b1a578160200160208202803683370190505b50905060008267ffffffffffffffff811115612b3857612b38613ba8565b604051908082528060200260200182016040528015612b61578160200160208202803683370190505b5090506000805b84811015612df257600060086000888481518110612b8857612b886138bb565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b03169050600060086000898581518110612bea57612bea6138bb565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008883601f8110612c4d57612c4d6138bb565b6020020151600254908b0363ffffffff169150600160a81b900467ffffffffffffffff16810282018015612de7576000601060008b8781518110612c9357612c936138bb565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b0316905080888781518110612ce457612ce46138bb565b60200260200101906001600160a01b031690816001600160a01b03168152505081878781518110612d1757612d176138bb565b6020026020010181815250508b8b86601f8110612d3657612d366138bb565b602002019063ffffffff16908163ffffffff168152505085806001019650507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b03168b8781518110612d9b57612d9b6138bb565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c85604051612ddd91815260200190565b60405180910390a4505b505050600101612b68565b5081518114612e02578082528083525b612e0f600c87601f6133e9565b50815115612eb1576040517f73433a2f0000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906373433a2f90612e7e90869086906004016143d8565b600060405180830381600087803b158015612e9857600080fd5b505af1158015612eac573d6000803e3d6000fd5b505050505b5050505050505050505050565b60004661a4b1811480612ed3575062066eed81145b15612f415760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f3b91906139f9565b91505090565b4391505090565b6000808a8a8a8a8a8a8a8a8a604051602001612f6c9998979695949392919061442f565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b612fdf8282613277565b5050565b336001600160a01b0382160361303b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161057c565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6012546001600160a01b039081169082168114612fdf57601280546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b60004661a4b1811480613117575062066eed81145b156131ce576101008367ffffffffffffffff16613132612ebe565b61313c9190613a12565b111561314b5750600092915050565b6040517f2b407a8200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152606490632b407a8290602401602060405180830381865afa1580156131aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119ec91906139f9565b505067ffffffffffffffff164090565b60008367ffffffffffffffff8416811015613212576002858567ffffffffffffffff16038161320f5761320f6143c2565b04015b613226818467ffffffffffffffff16612835565b95945050505050565b60008186101561326b576040517f3fef97df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50909303019091010290565b6101008181146132b9578282826040517f418a179b00000000000000000000000000000000000000000000000000000000815260040161057c939291906144b7565b60006132c7838501856144db565b90506040517f8eef585f0000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638eef585f9061332e908490600401614545565b600060405180830381600087803b15801561334857600080fd5b505af115801561335c573d6000803e3d6000fd5b5050505050505050565b5080546000825590600052602060002090810190610377919061347c565b8280548282559060005260206000209081019282156133d9579160200282015b828111156133d957825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906133a4565b506133e592915061347c565b5090565b6004830191839082156133d95791602002820160005b8382111561344357835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026133ff565b80156134735782816101000a81549063ffffffff0219169055600401602081600301049283019260010302613443565b50506133e59291505b5b808211156133e5576000815560010161347d565b6001600160a01b038116811461037757600080fd5b6000602082840312156134b857600080fd5b81356119ec81613491565b60008083601f8401126134d557600080fd5b50813567ffffffffffffffff8111156134ed57600080fd5b6020830191508360208260051b850101111561350857600080fd5b9250929050565b6000806000806040858703121561352557600080fd5b843567ffffffffffffffff8082111561353d57600080fd5b613549888389016134c3565b9096509450602087013591508082111561356257600080fd5b5061356f878288016134c3565b95989497509550505050565b60008083601f84011261358d57600080fd5b50813567ffffffffffffffff8111156135a557600080fd5b60208301915083602082850101111561350857600080fd5b60008060008060008060008060e0898b0312156135d957600080fd5b606089018a8111156135ea57600080fd5b8998503567ffffffffffffffff8082111561360457600080fd5b6136108c838d0161357b565b909950975060808b013591508082111561362957600080fd5b6136358c838d016134c3565b909750955060a08b013591508082111561364e57600080fd5b5061365b8b828c016134c3565b999c989b50969995989497949560c00135949350505050565b60006020828403121561368657600080fd5b813567ffffffffffffffff81111561369d57600080fd5b820160a081850312156119ec57600080fd5b600080604083850312156136c257600080fd5b82356136cd81613491565b946020939093013593505050565b803560ff811681146136ec57600080fd5b919050565b67ffffffffffffffff8116811461037757600080fd5b80356136ec816136f1565b60008060008060008060008060008060c08b8d03121561373157600080fd5b8a3567ffffffffffffffff8082111561374957600080fd5b6137558e838f016134c3565b909c509a5060208d013591508082111561376e57600080fd5b61377a8e838f016134c3565b909a50985088915061378e60408e016136db565b975060608d01359150808211156137a457600080fd5b6137b08e838f0161357b565b90975095508591506137c460808e01613707565b945060a08d01359150808211156137da57600080fd5b506137e78d828e0161357b565b915080935050809150509295989b9194979a5092959850565b62ffffff8116811461037757600080fd5b600080600080600060a0868803121561382957600080fd5b8535613834816136f1565b94506020860135613844816136f1565b93506040860135613854816136f1565b92506060860135613864816136f1565b9150608086013561387481613800565b809150509295509295909350565b6000806040838503121561389557600080fd5b82356138a081613491565b915060208301356138b081613491565b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b63ffffffff828116828216039080821115613904576139046138d1565b5092915050565b8082028115828204841417612849576128496138d1565b80820180821115612849576128496138d1565b600060018201613947576139476138d1565b5060010190565b60ff8181168382160190811115612849576128496138d1565b8183823760009101908152919050565b828152606082602083013760800192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b038416815260406020820152600061322660408301848661398b565b6000602082840312156139e957600080fd5b815180151581146119ec57600080fd5b600060208284031215613a0b57600080fd5b5051919050565b81810381811115612849576128496138d1565b8181036000831280158383131683831282161715613904576139046138d1565b60ff8181168382160290811690818114613904576139046138d1565b63ffffffff818116838216019080821115613904576139046138d1565b600081518084526020808501945080840160005b83811015613ab75781516001600160a01b031687529582019590820190600101613a92565b509495945050505050565b60005b83811015613add578181015183820152602001613ac5565b50506000910152565b60008151808452613afe816020860160208601613ac2565b601f01601f19169290920160200192915050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152613b428184018a613a7e565b90508281036080840152613b568189613a7e565b905060ff871660a084015282810360c0840152613b738187613ae6565b905067ffffffffffffffff851660e0840152828103610100840152613b988185613ae6565b9c9b505050505050505050505050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613be157613be1613ba8565b60405290565b604051610100810167ffffffffffffffff81118282101715613be157613be1613ba8565b60405160a0810167ffffffffffffffff81118282101715613be157613be1613ba8565b6040516080810167ffffffffffffffff81118282101715613be157613be1613ba8565b6040516020810167ffffffffffffffff81118282101715613be157613be1613ba8565b604051601f8201601f1916810167ffffffffffffffff81118282101715613c9d57613c9d613ba8565b604052919050565b600067ffffffffffffffff821115613cbf57613cbf613ba8565b5060051b60200190565b80516136ec816136f1565b805161ffff811681146136ec57600080fd5b80516136ec81613491565b600082601f830112613d0257600080fd5b815167ffffffffffffffff811115613d1c57613d1c613ba8565b613d2f601f8201601f1916602001613c74565b818152846020838601011115613d4457600080fd5b61051e826020830160208701613ac2565b80516001600160601b03811681146136ec57600080fd5b600082601f830112613d7d57600080fd5b81516020613d92613d8d83613ca5565b613c74565b82815260059290921b84018101918181019086841115613db157600080fd5b8286015b84811015613ec757805167ffffffffffffffff80821115613dd557600080fd5b90880190601f196040838c0382011215613dee57600080fd5b613df6613bbe565b8784015183811115613e0757600080fd5b8401610100818e0384011215613e1c57600080fd5b613e24613be7565b9250888101518352613e3860408201613cd4565b89840152613e4860608201613ce6565b6040840152608081015184811115613e5f57600080fd5b613e6d8e8b83850101613cf1565b606085015250613e7f60a08201613d55565b608084015260c081015160a084015260e081015160c084015261010081015160e084015250818152613eb360408501613d55565b818901528652505050918301918301613db5565b509695505050505050565b80516001600160c01b03811681146136ec57600080fd5b600060208284031215613efb57600080fd5b815167ffffffffffffffff80821115613f1357600080fd5b9083019060a08286031215613f2757600080fd5b613f2f613c0b565b825182811115613f3e57600080fd5b8301601f81018713613f4f57600080fd5b8051613f5d613d8d82613ca5565b8082825260208201915060208360051b850101925089831115613f7f57600080fd5b602084015b8381101561407f57805187811115613f9b57600080fd5b850160a0818d03601f19011215613fb157600080fd5b613fb9613c2e565b6020820151613fc7816136f1565b81526040820151613fd781613800565b60208201526040828e03605f19011215613ff057600080fd5b613ff8613c51565b8d607f84011261400757600080fd5b61400f613bbe565b808f60a08601111561402057600080fd5b606085015b60a08601811015614040578051835260209283019201614025565b50825250604082015260a08201518981111561405b57600080fd5b61406a8e602083860101613d6c565b60608301525084525060209283019201613f84565b5084525061409291505060208401613ed2565b60208201526140a360408401613cc9565b60408201526140b460608401613cc9565b60608201526080830151608082015280935050505092915050565b600063ffffffff8083168181036140e8576140e86138d1565b6001019392505050565b600081518084526020808501808196508360051b8101915082860160005b858110156141d657828403895281516040815181875280518288015287810151606061ffff8216818a0152838301519350608091506001600160a01b038416828a01528083015193505061010060a081818b01526141726101408b0186613ae6565b9284015192945060c061418f8b8201856001600160601b03169052565b9084015160e08b81019190915290840151918a01919091529091015161012088015250908601516001600160601b0316948601949094529784019790840190600101614110565b5091979650505050505050565b600060a080830181845280895180835260c08601915060c08160051b87010192506020808c016000805b848110156142945789870360bf190186528251805167ffffffffffffffff1688528481015162ffffff1685890152604080820151519084908a015b6002821015614267578251815291870191600191909101908701614248565b5050506060015160808801899052614281888a01826140f2565b975050948301949183019160010161420d565b5050508395506142ae8188018c6001600160c01b03169052565b50505050506142c9604083018667ffffffffffffffff169052565b67ffffffffffffffff939093166060820152608001529392505050565b600060208083850312156142f957600080fd5b825167ffffffffffffffff81111561431057600080fd5b8301601f8101851361432157600080fd5b805161432f613d8d82613ca5565b81815260079190911b8201830190838101908783111561434e57600080fd5b928401925b828410156143b7576080848903121561436c5760008081fd5b614374613c2e565b845161437f816136f1565b81528486015161438e81613800565b818701526040858101519082015260608086015190820152825260809093019290840190614353565b979650505050505050565b634e487b7160e01b600052601260045260246000fd5b6040815260006143eb6040830185613a7e565b82810360208481019190915284518083528582019282019060005b8181101561442257845183529383019391830191600101614406565b5090979650505050505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526144698285018b613a7e565b9150838203608085015261447d828a613a7e565b915060ff881660a085015283820360c085015261449a8288613ae6565b90861660e08501528381036101008501529050613b988185613ae6565b6040815260006144cb60408301858761398b565b9050826020830152949350505050565b60006101008083850312156144ef57600080fd5b83601f8401126144fe57600080fd5b614506613be7565b90830190808583111561451857600080fd5b845b8381101561453b57803561452d81613800565b83526020928301920161451a565b5095945050505050565b6101008101818360005b600881101561457357815162ffffff1683526020928301929091019060010161454f565b5050509291505056fea164736f6c6343000813000a",
}

var VRFBeaconOCRABI = VRFBeaconOCRMetaData.ABI

var VRFBeaconOCRBin = VRFBeaconOCRMetaData.Bin

func DeployVRFBeaconOCR(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, coordinator common.Address) (common.Address, *types.Transaction, *VRFBeaconOCR, error) {
	parsed, err := VRFBeaconOCRMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconOCRBin), backend, link, coordinator)
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
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
}, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "getBilling")

	outstruct := new(struct {
		MaximumGasPrice     uint64
		ReasonableGasPrice  uint64
		ObservationPayment  uint64
		TransmissionPayment uint64
		AccountingGas       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaximumGasPrice = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReasonableGasPrice = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ObservationPayment = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.TransmissionPayment = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.AccountingGas = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) GetBilling() (struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
}, error) {
	return _VRFBeaconOCR.Contract.GetBilling(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) GetBilling() (struct {
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
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

func (_VRFBeaconOCR *VRFBeaconOCRCaller) ICoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "i_coordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) ICoordinator() (common.Address, error) {
	return _VRFBeaconOCR.Contract.ICoordinator(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) ICoordinator() (common.Address, error) {
	return _VRFBeaconOCR.Contract.ICoordinator(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) ILink() (common.Address, error) {
	return _VRFBeaconOCR.Contract.ILink(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) ILink() (common.Address, error) {
	return _VRFBeaconOCR.Contract.ILink(&_VRFBeaconOCR.CallOpts)
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

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) SetBilling(opts *bind.TransactOpts, maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "setBilling", maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) SetBilling(maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetBilling(&_VRFBeaconOCR.TransactOpts, maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) SetBilling(maximumGasPrice uint64, reasonableGasPrice uint64, observationPayment uint64, transmissionPayment uint64, accountingGas *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.SetBilling(&_VRFBeaconOCR.TransactOpts, maximumGasPrice, reasonableGasPrice, observationPayment, transmissionPayment, accountingGas)
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
	MaximumGasPrice     uint64
	ReasonableGasPrice  uint64
	ObservationPayment  uint64
	TransmissionPayment uint64
	AccountingGas       *big.Int
	Raw                 types.Log
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
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*VRFBeaconOCRNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRNewTransmissionIterator{contract: _VRFBeaconOCR.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
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

type VRFBeaconOCROutputsServedIterator struct {
	Event *VRFBeaconOCROutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCROutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCROutputsServed)
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
		it.Event = new(VRFBeaconOCROutputsServed)
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

func (it *VRFBeaconOCROutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCROutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCROutputsServed struct {
	RecentBlockHeight  uint64
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFBeaconOCROutputsServedIterator, error) {

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCROutputsServedIterator{contract: _VRFBeaconOCR.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCROutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCROutputsServed)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseOutputsServed(log types.Log) (*VRFBeaconOCROutputsServed, error) {
	event := new(VRFBeaconOCROutputsServed)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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
	Raw                    types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconOCRRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconOCR.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
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

type VRFBeaconOCRRandomnessRedeemedIterator struct {
	Event *VRFBeaconOCRRandomnessRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconOCRRandomnessRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconOCRRandomnessRedeemed)
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
		it.Event = new(VRFBeaconOCRRandomnessRedeemed)
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

func (it *VRFBeaconOCRRandomnessRedeemedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconOCRRandomnessRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconOCRRandomnessRedeemed struct {
	RequestID *big.Int
	Requester common.Address
	SubID     *big.Int
	Raw       types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconOCRRandomnessRedeemedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRRandomnessRedeemedIterator{contract: _VRFBeaconOCR.contract, event: "RandomnessRedeemed", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconOCRRandomnessRedeemed)
				if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) ParseRandomnessRedeemed(log types.Log) (*VRFBeaconOCRRandomnessRedeemed, error) {
	event := new(VRFBeaconOCRRandomnessRedeemed)
	if err := _VRFBeaconOCR.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	CostJuels              *big.Int
	Raw                    types.Log
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconOCRRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconOCRRandomnessRequestedIterator{contract: _VRFBeaconOCR.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconOCR *VRFBeaconOCRFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconOCRRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconOCR.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"coordinator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconTypes.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_coordinator\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorProducerAPI\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_link\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161059138038061059183398101604081905261002f91610190565b33806000816100855760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b5576100b5816100cf565b5050506001600160a01b039182166080521660a0526101ca565b336001600160a01b038216036101275760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007c565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b038116811461018d57600080fd5b50565b600080604083850312156101a357600080fd5b82516101ae81610178565b60208401519092506101bf81610178565b809150509250929050565b60805160a0516103a46101ed600039600060ef0152600060b001526103a46000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638a1b17721161005b5780638a1b1772146100ea5780638da5cb5b14610111578063b8be03cd14610122578063f2fde38b1461013357600080fd5b80632f7527cc1461008257806379ba5097146100a15780637d253aff146100ab575b600080fd5b61008a600881565b60405160ff90911681526020015b60405180910390f35b6100a9610146565b005b6100d27f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610098565b6100d27f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b03166100d2565b6100a961013036600461032c565b50565b6100a961014136600461036e565b610209565b6001546001600160a01b031633146101a55760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61021161021a565b61013081610276565b6000546001600160a01b031633146102745760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161019c565b565b336001600160a01b038216036102ce5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161019c565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561033e57600080fd5b813567ffffffffffffffff81111561035557600080fd5b820160a0818503121561036757600080fd5b9392505050565b60006020828403121561038057600080fd5b81356001600160a01b038116811461036757600080fdfea164736f6c6343000813000a",
}

var VRFBeaconReportABI = VRFBeaconReportMetaData.ABI

var VRFBeaconReportBin = VRFBeaconReportMetaData.Bin

func DeployVRFBeaconReport(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address, coordinator common.Address) (common.Address, *types.Transaction, *VRFBeaconReport, error) {
	parsed, err := VRFBeaconReportMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconReportBin), backend, link, coordinator)
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

func (_VRFBeaconReport *VRFBeaconReportCaller) ICoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "i_coordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) ICoordinator() (common.Address, error) {
	return _VRFBeaconReport.Contract.ICoordinator(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) ICoordinator() (common.Address, error) {
	return _VRFBeaconReport.Contract.ICoordinator(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) ILink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "i_link")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) ILink() (common.Address, error) {
	return _VRFBeaconReport.Contract.ILink(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) ILink() (common.Address, error) {
	return _VRFBeaconReport.Contract.ILink(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) Owner() (common.Address, error) {
	return _VRFBeaconReport.Contract.Owner(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) Owner() (common.Address, error) {
	return _VRFBeaconReport.Contract.Owner(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeaconReport *VRFBeaconReportSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.AcceptOwnership(&_VRFBeaconReport.TransactOpts)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.AcceptOwnership(&_VRFBeaconReport.TransactOpts)
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

func (_VRFBeaconReport *VRFBeaconReportTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeaconReport *VRFBeaconReportSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.TransferOwnership(&_VRFBeaconReport.TransactOpts, to)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.TransferOwnership(&_VRFBeaconReport.TransactOpts, to)
}

type VRFBeaconReportConfigSetIterator struct {
	Event *VRFBeaconReportConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportConfigSet)
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
		it.Event = new(VRFBeaconReportConfigSet)
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

func (it *VRFBeaconReportConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportConfigSet struct {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFBeaconReportConfigSetIterator, error) {

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportConfigSetIterator{contract: _VRFBeaconReport.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportConfigSet)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseConfigSet(log types.Log) (*VRFBeaconReportConfigSet, error) {
	event := new(VRFBeaconReportConfigSet)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	AggregatorRoundId  uint32
	EpochAndRound      *big.Int
	Transmitter        common.Address
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	ConfigDigest       [32]byte
	Raw                types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*VRFBeaconReportNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportNewTransmissionIterator{contract: _VRFBeaconReport.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
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

type VRFBeaconReportOutputsServedIterator struct {
	Event *VRFBeaconReportOutputsServed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportOutputsServedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportOutputsServed)
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
		it.Event = new(VRFBeaconReportOutputsServed)
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

func (it *VRFBeaconReportOutputsServedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportOutputsServedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportOutputsServed struct {
	RecentBlockHeight  uint64
	JuelsPerFeeCoin    *big.Int
	ReasonableGasPrice uint64
	OutputsServed      []VRFBeaconTypesOutputServed
	Raw                types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterOutputsServed(opts *bind.FilterOpts) (*VRFBeaconReportOutputsServedIterator, error) {

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportOutputsServedIterator{contract: _VRFBeaconReport.contract, event: "OutputsServed", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchOutputsServed(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportOutputsServed) (event.Subscription, error) {

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "OutputsServed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportOutputsServed)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "OutputsServed", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseOutputsServed(log types.Log) (*VRFBeaconReportOutputsServed, error) {
	event := new(VRFBeaconReportOutputsServed)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "OutputsServed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconReportOwnershipTransferRequestedIterator struct {
	Event *VRFBeaconReportOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportOwnershipTransferRequested)
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
		it.Event = new(VRFBeaconReportOwnershipTransferRequested)
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

func (it *VRFBeaconReportOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconReportOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportOwnershipTransferRequestedIterator{contract: _VRFBeaconReport.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportOwnershipTransferRequested)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconReportOwnershipTransferRequested, error) {
	event := new(VRFBeaconReportOwnershipTransferRequested)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconReportOwnershipTransferredIterator struct {
	Event *VRFBeaconReportOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportOwnershipTransferred)
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
		it.Event = new(VRFBeaconReportOwnershipTransferred)
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

func (it *VRFBeaconReportOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconReportOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportOwnershipTransferredIterator{contract: _VRFBeaconReport.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportOwnershipTransferred)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseOwnershipTransferred(log types.Log) (*VRFBeaconReportOwnershipTransferred, error) {
	event := new(VRFBeaconReportOwnershipTransferred)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
	Raw                    types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterRandomnessFulfillmentRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconReportRandomnessFulfillmentRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportRandomnessFulfillmentRequestedIterator{contract: _VRFBeaconReport.contract, event: "RandomnessFulfillmentRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchRandomnessFulfillmentRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportRandomnessFulfillmentRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "RandomnessFulfillmentRequested", requestIDRule, requesterRule)
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

type VRFBeaconReportRandomnessRedeemedIterator struct {
	Event *VRFBeaconReportRandomnessRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconReportRandomnessRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconReportRandomnessRedeemed)
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
		it.Event = new(VRFBeaconReportRandomnessRedeemed)
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

func (it *VRFBeaconReportRandomnessRedeemedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconReportRandomnessRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconReportRandomnessRedeemed struct {
	RequestID *big.Int
	Requester common.Address
	SubID     *big.Int
	Raw       types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterRandomnessRedeemed(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconReportRandomnessRedeemedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportRandomnessRedeemedIterator{contract: _VRFBeaconReport.contract, event: "RandomnessRedeemed", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchRandomnessRedeemed(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportRandomnessRedeemed, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "RandomnessRedeemed", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconReportRandomnessRedeemed)
				if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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

func (_VRFBeaconReport *VRFBeaconReportFilterer) ParseRandomnessRedeemed(log types.Log) (*VRFBeaconReportRandomnessRedeemed, error) {
	event := new(VRFBeaconReportRandomnessRedeemed)
	if err := _VRFBeaconReport.contract.UnpackLog(event, "RandomnessRedeemed", log); err != nil {
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
	RequestID              *big.Int
	Requester              common.Address
	NextBeaconOutputHeight uint64
	ConfDelay              *big.Int
	SubID                  *big.Int
	NumWords               uint16
	CostJuels              *big.Int
	Raw                    types.Log
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) FilterRandomnessRequested(opts *bind.FilterOpts, requestID []*big.Int, requester []common.Address) (*VRFBeaconReportRandomnessRequestedIterator, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.FilterLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconReportRandomnessRequestedIterator{contract: _VRFBeaconReport.contract, event: "RandomnessRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconReport *VRFBeaconReportFilterer) WatchRandomnessRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconReportRandomnessRequested, requestID []*big.Int, requester []common.Address) (event.Subscription, error) {

	var requestIDRule []interface{}
	for _, requestIDItem := range requestID {
		requestIDRule = append(requestIDRule, requestIDItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _VRFBeaconReport.contract.WatchLogs(opts, "RandomnessRequested", requestIDRule, requesterRule)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reasonableGasPrice\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"proofG1X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proofG1Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"OutputsServed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAllowance\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weiPerUnitLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"}],\"name\":\"RandomnessRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"costJuels\",\"type\":\"uint256\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	CostJuels              *big.Int
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

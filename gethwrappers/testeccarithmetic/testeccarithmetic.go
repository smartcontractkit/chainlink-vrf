package testeccarithmetic

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

var TestECCArithmeticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"p1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"p2\",\"type\":\"tuple\"}],\"name\":\"testAddG1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"sum\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"g1Base\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"signature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[4]\",\"name\":\"p\",\"type\":\"uint256[4]\"}],\"internalType\":\"structECCArithmetic.G2Point\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"name\":\"testDiscreteLogsMatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"testNegation\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testPairingSmokeTest\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testPairingSmokeTest2\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610fcc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80633ce82ca6116100505780633ce82ca61461009e578063f1374692146100a6578063fafa4f3b146100b957600080fd5b80631bace06d1461006c57806330a267b414610094575b600080fd5b61007f61007a366004610e13565b6100d9565b60405190151581526020015b60405180910390f35b61009c6100ee565b005b61009c610217565b61009c6100b4366004610eb9565b610454565b6100cc6100c7366004610ed5565b61050b565b60405161008b9190610f0a565b60006100e6848484610526565b949350505050565b6100f6610ca2565b8051600190528051600260209091015261010e610cba565b80517f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2905280517f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60209091015280517f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60409091015280517f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6060909101526101b782610454565b6101c2828383610526565b6102135760405162461bcd60e51b815260206004820152601160248201527f736d6f6b652074657374206661696c656400000000000000000000000000000060448201526064015b60405180910390fd5b5050565b61021f610ca2565b80516001905280516002602090910152610237610ca2565b80516001905261026860027f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610f69565b81516020015260408051600280825260608201909252600091816020015b61028e610ca2565b81526020019060019003908161028657905050905081816000815181106102b7576102b7610f3d565b602002602001018190525082816001815181106102d6576102d6610f3d565b60200260200101819052506102e9610cba565b80517f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2905280517f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed60209182015281517f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b60409182015282517f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6060918201528151600280825291810190925260009282015b6103a4610cba565b81526020019060019003908161039c57905050905081816000815181106103cd576103cd610f3d565b602002602001018190525081816001815181106103ec576103ec610f3d565b60200260200101819052506104018382610721565b61044d5760405162461bcd60e51b815260206004820152601360248201527f736d6f6b6520746573742032206661696c656400000000000000000000000000604482015260640161020a565b5050505050565b61045c610ca2565b815151815152815160200151610492907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610f69565b81516020015260006104a483836109d2565b8051519091501580156104ba5750805160200151155b6105065760405162461bcd60e51b815260206004820152600d60248201527f70202b20282d702920213d203000000000000000000000000000000000000000604482015260640161020a565b505050565b610513610ca2565b61051d8383610a96565b90505b92915050565b60408051600280825260608201909252600091829190816020015b610549610ca2565b8152602001906001900390816105415750506040805160028082526060820190925291925060009190602082015b61057f610cba565b8152602001906001900390816105775750506040805160608101825288515160208083019182528a510151939450909283928301906105de907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610f69565b815250815250826000815181106105f7576105f7610f3d565b6020026020010181905250848260018151811061061657610616610f3d565b6020026020010181905250838160008151811061063557610635610f3d565b6020026020010181905250604051806020016040528060405180608001604052807f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c281526020017f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed81526020017f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b81526020017f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa8152508152508160018151811061070257610702610f3d565b60200260200101819052506107178282610721565b9695505050505050565b6000815183511461073157600080fd5b82516000610740826006610f7c565b905060008167ffffffffffffffff81111561075d5761075d610d27565b604051908082528060200260200182016040528015610786578160200160208202803683370190505b50905060005b8381101561099b578681815181106107a6576107a6610f3d565b60209081029190910101515151826107bf836006610f7c565b6107ca906000610f93565b815181106107da576107da610f3d565b6020026020010181815250508681815181106107f8576107f8610f3d565b60209081029190910181015151015182610813836006610f7c565b61081e906001610f93565b8151811061082e5761082e610f3d565b60200260200101818152505085818151811061084c5761084c610f3d565b6020908102919091010151515182610865836006610f7c565b610870906002610f93565b8151811061088057610880610f3d565b60200260200101818152505085818151811061089e5761089e610f3d565b602090810291909101810151510151826108b9836006610f7c565b6108c4906003610f93565b815181106108d4576108d4610f3d565b6020026020010181815250508581815181106108f2576108f2610f3d565b602090810291909101015151604001518261090e836006610f7c565b610919906004610f93565b8151811061092957610929610f3d565b60200260200101818152505085818151811061094757610947610f3d565b6020908102919091010151516060015182610963836006610f7c565b61096e906005610f93565b8151811061097e5761097e610f3d565b60209081029190910101528061099381610fa6565b91505061078c565b506109a4610ccd565b6000602082602086026020860160086201b968fa9050806109c457600080fd5b505115159695505050505050565b6109da610ca2565b6109e383610b0e565b6109ec82610b0e565b6109f4610ceb565b83515181528351602090810151828201528351516040830152835101516060820152610a1e610d09565b600060408260808560066096fa905080600003610a7d5760405162461bcd60e51b815260206004820152601160248201527f61646467312063616c6c206661696c6564000000000000000000000000000000604482015260640161020a565b5080518351526020908101518351909101525092915050565b610a9e610ca2565b6000610aaa84846109d2565b80515190915015801590610ac2575080516020015115155b61051d5760405162461bcd60e51b815260206004820152601b60248201527f6164646731206661696c65643a207a65726f206f7264696e6174650000000000604482015260640161020a565b8051517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4711610b7f5760405162461bcd60e51b815260206004820152600c60248201527f78206e6f7420696e20465f500000000000000000000000000000000000000000604482015260640161020a565b8051602001517f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4711610bf35760405162461bcd60e51b815260206004820152600c60248201527f79206e6f7420696e20465f500000000000000000000000000000000000000000604482015260640161020a565b8051516000907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790600390829081818009090882516020015190915081907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47908009146102135760405162461bcd60e51b815260206004820152601260248201527f706f696e74206e6f74206f6e2063757276650000000000000000000000000000604482015260640161020a565b6040518060200160405280610cb5610d09565b905290565b6040518060200160405280610cb5610ceb565b60405180602001604052806001906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b6040516020810167ffffffffffffffff81118282101715610d6057610d60610d27565b60405290565b6040516080810167ffffffffffffffff81118282101715610d6057610d60610d27565b600060408284031215610d9b57600080fd5b610da3610d3d565b905082601f830112610db457600080fd5b6040516040810181811067ffffffffffffffff82111715610dd757610dd7610d27565b8060405250806040840185811115610dee57600080fd5b845b81811015610e08578035835260209283019201610df0565b505050815292915050565b600080600083850361010080821215610e2b57600080fd5b610e358787610d89565b9450610e448760408801610d89565b93506080607f1983011215610e5857600080fd5b610e60610d3d565b915086609f870112610e7157600080fd5b610e79610d66565b908601908088831115610e8b57600080fd5b608088015b83811015610ea8578035835260209283019201610e90565b508352509396929550935090915050565b600060408284031215610ecb57600080fd5b61051d8383610d89565b60008060808385031215610ee857600080fd5b610ef28484610d89565b9150610f018460408501610d89565b90509250929050565b815160408201908260005b6002811015610f34578251825260209283019290910190600101610f15565b50505092915050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561052057610520610f53565b808202811582820484141761052057610520610f53565b8082018082111561052057610520610f53565b600060018201610fb857610fb8610f53565b506001019056fea164736f6c6343000813000a",
}

var TestECCArithmeticABI = TestECCArithmeticMetaData.ABI

var TestECCArithmeticBin = TestECCArithmeticMetaData.Bin

func DeployTestECCArithmetic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestECCArithmetic, error) {
	parsed, err := TestECCArithmeticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestECCArithmeticBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestECCArithmetic{TestECCArithmeticCaller: TestECCArithmeticCaller{contract: contract}, TestECCArithmeticTransactor: TestECCArithmeticTransactor{contract: contract}, TestECCArithmeticFilterer: TestECCArithmeticFilterer{contract: contract}}, nil
}

type TestECCArithmetic struct {
	TestECCArithmeticCaller
	TestECCArithmeticTransactor
	TestECCArithmeticFilterer
}

type TestECCArithmeticCaller struct {
	contract *bind.BoundContract
}

type TestECCArithmeticTransactor struct {
	contract *bind.BoundContract
}

type TestECCArithmeticFilterer struct {
	contract *bind.BoundContract
}

type TestECCArithmeticSession struct {
	Contract     *TestECCArithmetic
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TestECCArithmeticCallerSession struct {
	Contract *TestECCArithmeticCaller
	CallOpts bind.CallOpts
}

type TestECCArithmeticTransactorSession struct {
	Contract     *TestECCArithmeticTransactor
	TransactOpts bind.TransactOpts
}

type TestECCArithmeticRaw struct {
	Contract *TestECCArithmetic
}

type TestECCArithmeticCallerRaw struct {
	Contract *TestECCArithmeticCaller
}

type TestECCArithmeticTransactorRaw struct {
	Contract *TestECCArithmeticTransactor
}

func NewTestECCArithmetic(address common.Address, backend bind.ContractBackend) (*TestECCArithmetic, error) {
	contract, err := bindTestECCArithmetic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestECCArithmetic{TestECCArithmeticCaller: TestECCArithmeticCaller{contract: contract}, TestECCArithmeticTransactor: TestECCArithmeticTransactor{contract: contract}, TestECCArithmeticFilterer: TestECCArithmeticFilterer{contract: contract}}, nil
}

func NewTestECCArithmeticCaller(address common.Address, caller bind.ContractCaller) (*TestECCArithmeticCaller, error) {
	contract, err := bindTestECCArithmetic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestECCArithmeticCaller{contract: contract}, nil
}

func NewTestECCArithmeticTransactor(address common.Address, transactor bind.ContractTransactor) (*TestECCArithmeticTransactor, error) {
	contract, err := bindTestECCArithmetic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestECCArithmeticTransactor{contract: contract}, nil
}

func NewTestECCArithmeticFilterer(address common.Address, filterer bind.ContractFilterer) (*TestECCArithmeticFilterer, error) {
	contract, err := bindTestECCArithmetic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestECCArithmeticFilterer{contract: contract}, nil
}

func bindTestECCArithmetic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestECCArithmeticABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_TestECCArithmetic *TestECCArithmeticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestECCArithmetic.Contract.TestECCArithmeticCaller.contract.Call(opts, result, method, params...)
}

func (_TestECCArithmetic *TestECCArithmeticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestECCArithmetic.Contract.TestECCArithmeticTransactor.contract.Transfer(opts)
}

func (_TestECCArithmetic *TestECCArithmeticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestECCArithmetic.Contract.TestECCArithmeticTransactor.contract.Transact(opts, method, params...)
}

func (_TestECCArithmetic *TestECCArithmeticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestECCArithmetic.Contract.contract.Call(opts, result, method, params...)
}

func (_TestECCArithmetic *TestECCArithmeticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestECCArithmetic.Contract.contract.Transfer(opts)
}

func (_TestECCArithmetic *TestECCArithmeticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestECCArithmetic.Contract.contract.Transact(opts, method, params...)
}

func (_TestECCArithmetic *TestECCArithmeticCaller) TestAddG1(opts *bind.CallOpts, p1 ECCArithmeticG1Point, p2 ECCArithmeticG1Point) (ECCArithmeticG1Point, error) {
	var out []interface{}
	err := _TestECCArithmetic.contract.Call(opts, &out, "testAddG1", p1, p2)

	if err != nil {
		return *new(ECCArithmeticG1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(ECCArithmeticG1Point)).(*ECCArithmeticG1Point)

	return out0, err

}

func (_TestECCArithmetic *TestECCArithmeticSession) TestAddG1(p1 ECCArithmeticG1Point, p2 ECCArithmeticG1Point) (ECCArithmeticG1Point, error) {
	return _TestECCArithmetic.Contract.TestAddG1(&_TestECCArithmetic.CallOpts, p1, p2)
}

func (_TestECCArithmetic *TestECCArithmeticCallerSession) TestAddG1(p1 ECCArithmeticG1Point, p2 ECCArithmeticG1Point) (ECCArithmeticG1Point, error) {
	return _TestECCArithmetic.Contract.TestAddG1(&_TestECCArithmetic.CallOpts, p1, p2)
}

func (_TestECCArithmetic *TestECCArithmeticCaller) TestDiscreteLogsMatch(opts *bind.CallOpts, g1Base ECCArithmeticG1Point, signature ECCArithmeticG1Point, pubkey ECCArithmeticG2Point) (bool, error) {
	var out []interface{}
	err := _TestECCArithmetic.contract.Call(opts, &out, "testDiscreteLogsMatch", g1Base, signature, pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TestECCArithmetic *TestECCArithmeticSession) TestDiscreteLogsMatch(g1Base ECCArithmeticG1Point, signature ECCArithmeticG1Point, pubkey ECCArithmeticG2Point) (bool, error) {
	return _TestECCArithmetic.Contract.TestDiscreteLogsMatch(&_TestECCArithmetic.CallOpts, g1Base, signature, pubkey)
}

func (_TestECCArithmetic *TestECCArithmeticCallerSession) TestDiscreteLogsMatch(g1Base ECCArithmeticG1Point, signature ECCArithmeticG1Point, pubkey ECCArithmeticG2Point) (bool, error) {
	return _TestECCArithmetic.Contract.TestDiscreteLogsMatch(&_TestECCArithmetic.CallOpts, g1Base, signature, pubkey)
}

func (_TestECCArithmetic *TestECCArithmeticCaller) TestNegation(opts *bind.CallOpts, p ECCArithmeticG1Point) error {
	var out []interface{}
	err := _TestECCArithmetic.contract.Call(opts, &out, "testNegation", p)

	if err != nil {
		return err
	}

	return err

}

func (_TestECCArithmetic *TestECCArithmeticSession) TestNegation(p ECCArithmeticG1Point) error {
	return _TestECCArithmetic.Contract.TestNegation(&_TestECCArithmetic.CallOpts, p)
}

func (_TestECCArithmetic *TestECCArithmeticCallerSession) TestNegation(p ECCArithmeticG1Point) error {
	return _TestECCArithmetic.Contract.TestNegation(&_TestECCArithmetic.CallOpts, p)
}

func (_TestECCArithmetic *TestECCArithmeticCaller) TestPairingSmokeTest(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestECCArithmetic.contract.Call(opts, &out, "testPairingSmokeTest")

	if err != nil {
		return err
	}

	return err

}

func (_TestECCArithmetic *TestECCArithmeticSession) TestPairingSmokeTest() error {
	return _TestECCArithmetic.Contract.TestPairingSmokeTest(&_TestECCArithmetic.CallOpts)
}

func (_TestECCArithmetic *TestECCArithmeticCallerSession) TestPairingSmokeTest() error {
	return _TestECCArithmetic.Contract.TestPairingSmokeTest(&_TestECCArithmetic.CallOpts)
}

func (_TestECCArithmetic *TestECCArithmeticCaller) TestPairingSmokeTest2(opts *bind.CallOpts) error {
	var out []interface{}
	err := _TestECCArithmetic.contract.Call(opts, &out, "testPairingSmokeTest2")

	if err != nil {
		return err
	}

	return err

}

func (_TestECCArithmetic *TestECCArithmeticSession) TestPairingSmokeTest2() error {
	return _TestECCArithmetic.Contract.TestPairingSmokeTest2(&_TestECCArithmetic.CallOpts)
}

func (_TestECCArithmetic *TestECCArithmeticCallerSession) TestPairingSmokeTest2() error {
	return _TestECCArithmetic.Contract.TestPairingSmokeTest2(&_TestECCArithmetic.CallOpts)
}

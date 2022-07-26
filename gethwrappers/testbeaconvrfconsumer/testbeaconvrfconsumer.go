package testbeaconvrfconsumer

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

var BeaconVRFConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"shouldFail\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"computeAndStoreExpectedOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fail\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_ExpectedRandomnessByBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"s_ExpectedSeeds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_ReceivedRandomnessByRequestID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_gasAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"name\":\"s_myBeaconRequests\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.SlotNumber\",\"name\":\"slotNumber\",\"type\":\"uint32\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_randomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"s_requestsIDs\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_subId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"}],\"name\":\"setExpectedSeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"shouldFail\",\"type\":\"bool\"}],\"name\":\"setFail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"testRedeemRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"testRequestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"testRequestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161147f38038061147f83398101604081905261002f9161006c565b6001600160a01b03929092166080819052600680546001600160a01b03191690911790556009805460ff1916911515919091179055600a556100c0565b60008060006060848603121561008157600080fd5b83516001600160a01b038116811461009857600080fd5b602085015190935080151581146100ae57600080fd5b80925050604084015190509250925092565b6080516113a46100db60003960006106e101526113a46000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c8063706da1ca116100b2578063cd0593df11610081578063f08c5daa11610066578063f08c5daa14610354578063f339c7941461035d578063f6eaffc81461037057600080fd5b8063cd0593df14610338578063e4d47bfa1461034157600080fd5b8063706da1ca146102a25780639c9cd015146102e75780639d769402146102fa578063a9cc47181461031b57600080fd5b8063563db24c116100ee578063563db24c146102105780635a47dd71146102315780635f15cccc14610244578063678d38f71461027757600080fd5b806319a5fa22146101205780633d8b70aa146101be57806340e836ab146101d357806345907626146101e6575b600080fd5b61017c61012e366004610ce4565b60026020526000908152604090205463ffffffff811690640100000000810462ffffff1690670100000000000000810461ffff1690690100000000000000000090046001600160a01b031684565b6040805163ffffffff909516855262ffffff909316602085015261ffff909116918301919091526001600160a01b031660608201526080015b60405180910390f35b6101d16101cc366004610ce4565b610383565b005b6101d16101e1366004610dcb565b610449565b6101f96101f4366004610ef0565b6104a4565b60405165ffffffffffff90911681526020016101b5565b61022361021e366004610f7b565b6106ae565b6040519081526020016101b5565b6101d161023f366004610fcb565b6106df565b6101f961025236600461109d565b600160209081526000928352604080842090915290825290205465ffffffffffff1681565b61022361028536600461109d565b600560209081526000928352604080842090915290825290205481565b6007546102ce9074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016101b5565b6101f96102f53660046110c9565b610767565b6101d161030836600461110c565b6009805460ff1916911515919091179055565b6009546103289060ff1681565b60405190151581526020016101b5565b610223600a5481565b6101d161034f36600461112e565b610981565b61022360085481565b61022361036b366004610f7b565b610baf565b61022361037e366004611158565b610bcb565b6006546040517f74d8461100000000000000000000000000000000000000000000000000000000815265ffffffffffff831660048201526000916001600160a01b0316906374d84611906024016000604051808303816000875af11580156103ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104179190810190611171565b65ffffffffffff831660009081526003602090815260409091208251929350610444929091840190610c6d565b505050565b60008160405160200161045c9190611202565b60408051601f19818403018152918152815160209283012067ffffffffffffffff90961660009081526005835281812062ffffff909616815294909152909220929092555050565b600080600a54436104b5919061124b565b9050600081600a54436104c89190611275565b6104d2919061128d565b6006546040517ff645dcb10000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063f645dcb190610529908c908c908c908c908c906004016112a4565b6020604051808303816000875af1158015610548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c919061132f565b600083815260016020908152604080832062ffffff8c1684529091528120805465ffffffffffff191665ffffffffffff8416179055600a54919250906105b2908461134c565b6040805160808101825263ffffffff928316815262ffffff9a8b16602080830191825261ffff9d8e16838501908152306060850190815265ffffffffffff8916600090815260029093529490912092518354925191519451951666ffffffffffffff199092169190911764010000000091909c16029a909a177fffffff00000000000000000000000000000000000000000000ffffffffffffff1667010000000000000091909b16027fffffff0000000000000000000000000000000000000000ffffffffffffffffff169990991769010000000000000000006001600160a01b03909a16999099029890981790965550939695505050505050565b600460205281600052604060002081815481106106ca57600080fd5b90600052602060002001600091509150505481565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316331461075c5760405162461bcd60e51b815260206004820152601c60248201527f6f6e6c7920636f6f7264696e61746f722063616e2066756c66696c6c0000000060448201526064015b60405180910390fd5b610444838383610bec565b600080600a5443610778919061124b565b9050600081600a544361078b9190611275565b610795919061128d565b6006546040517fdc92accf00000000000000000000000000000000000000000000000000000000815261ffff8916600482015267ffffffffffffffff8816602482015262ffffff871660448201529192506000916001600160a01b039091169063dc92accf906064016020604051808303816000875af115801561081d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610841919061132f565b600083815260016020908152604080832062ffffff8a1684529091528120805465ffffffffffff191665ffffffffffff8416179055600a5491925090610887908461134c565b6040805160808101825263ffffffff928316815262ffffff988916602080830191825261ffff9c8d16838501908152306060850190815265ffffffffffff8916600090815260029093529490912092518354925191519451951666ffffffffffffff199092169190911764010000000091909a1602989098177fffffff00000000000000000000000000000000000000000000ffffffffffffff1667010000000000000091909a16027fffffff0000000000000000000000000000000000000000ffffffffffffffffff169890981769010000000000000000006001600160a01b0390991698909802979097179094555091949350505050565b67ffffffffffffffff8216600081815260016020908152604080832062ffffff8681168086529184528285205465ffffffffffff1680865260028552838620845160808082018752915463ffffffff808216835264010000000082048616838a0190815261ffff67010000000000000084048116858b019081526001600160a01b036901000000000000000000909504851660608088019182529e8e5260058d528b8e209a8e52998c528a8d20548b519c8d0189905286519094169a8c019a909a5290519096169a89019a909a52955190931690860152915190921660a084015260c08301859052939092909160e0016040516020818303038152906040528051906020012090506000836040015161ffff1667ffffffffffffffff811115610aac57610aac610d38565b604051908082528060200260200182016040528015610ad5578160200160208202803683370190505b50905060005b846040015161ffff168161ffff161015610b7c578281604051602001610b3092919091825260f01b7fffff00000000000000000000000000000000000000000000000000000000000016602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff1681518110610b5f57610b5f611360565b602090810291909101015280610b7481611376565b915050610adb565b5065ffffffffffff851660009081526004602090815260409091208251610ba592840190610c6d565b5050505050505050565b600360205281600052604060002081815481106106ca57600080fd5b60008181548110610bdb57600080fd5b600091825260209091200154905081565b60095460ff1615610c3f5760405162461bcd60e51b815260206004820152601d60248201527f206661696c656420696e2066756c66696c6c52616e646f6d576f7264730000006044820152606401610753565b65ffffffffffff831660009081526003602090815260409091208351610c6792850190610c6d565b50505050565b828054828255906000526020600020908101928215610ca8579160200282015b82811115610ca8578251825591602001919060010190610c8d565b50610cb4929150610cb8565b5090565b5b80821115610cb45760008155600101610cb9565b65ffffffffffff81168114610ce157600080fd5b50565b600060208284031215610cf657600080fd5b8135610d0181610ccd565b9392505050565b803567ffffffffffffffff81168114610d2057600080fd5b919050565b803562ffffff81168114610d2057600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516020810167ffffffffffffffff81118282101715610d7157610d71610d38565b60405290565b6040805190810167ffffffffffffffff81118282101715610d7157610d71610d38565b604051601f8201601f1916810167ffffffffffffffff81118282101715610dc357610dc3610d38565b604052919050565b60008060008385036080811215610de157600080fd5b610dea85610d08565b93506020610df9818701610d25565b93506040603f1983011215610e0d57600080fd5b610e15610d4e565b915086605f870112610e2657600080fd5b610e2e610d77565b806080880189811115610e4057600080fd5b604089015b81811015610e5c5780358452928401928401610e45565b50508352509396929550935090915050565b803561ffff81168114610d2057600080fd5b600082601f830112610e9157600080fd5b813567ffffffffffffffff811115610eab57610eab610d38565b610ebe601f8201601f1916602001610d9a565b818152846020838601011115610ed357600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a08688031215610f0857600080fd5b610f1186610d08565b9450610f1f60208701610e6e565b9350610f2d60408701610d25565b9250606086013563ffffffff81168114610f4657600080fd5b9150608086013567ffffffffffffffff811115610f6257600080fd5b610f6e88828901610e80565b9150509295509295909350565b60008060408385031215610f8e57600080fd5b8235610f9981610ccd565b946020939093013593505050565b600067ffffffffffffffff821115610fc157610fc1610d38565b5060051b60200190565b600080600060608486031215610fe057600080fd5b8335610feb81610ccd565b925060208481013567ffffffffffffffff8082111561100957600080fd5b818701915087601f83011261101d57600080fd5b813561103061102b82610fa7565b610d9a565b81815260059190911b8301840190848101908a83111561104f57600080fd5b938501935b8285101561106d57843582529385019390850190611054565b96505050604087013592508083111561108557600080fd5b505061109386828701610e80565b9150509250925092565b600080604083850312156110b057600080fd5b823591506110c060208401610d25565b90509250929050565b6000806000606084860312156110de57600080fd5b6110e784610e6e565b92506110f560208501610d08565b915061110360408501610d25565b90509250925092565b60006020828403121561111e57600080fd5b81358015158114610d0157600080fd5b6000806040838503121561114157600080fd5b61114a83610d08565b91506110c060208401610d25565b60006020828403121561116a57600080fd5b5035919050565b6000602080838503121561118457600080fd5b825167ffffffffffffffff81111561119b57600080fd5b8301601f810185136111ac57600080fd5b80516111ba61102b82610fa7565b81815260059190911b820183019083810190878311156111d957600080fd5b928401925b828410156111f7578351825292840192908401906111de565b979650505050505050565b815160408201908260005b600281101561122c57825182526020928301929091019060010161120d565b50505092915050565b634e487b7160e01b600052601260045260246000fd5b60008261125a5761125a611235565b500690565b634e487b7160e01b600052601160045260246000fd5b600082198211156112885761128861125f565b500190565b60008282101561129f5761129f61125f565b500390565b67ffffffffffffffff861681526000602061ffff87168184015262ffffff8616604084015263ffffffff8516606084015260a0608084015283518060a085015260005b818110156113035785810183015185820160c0015282016112e7565b8181111561131557600060c083870101525b50601f01601f19169290920160c001979650505050505050565b60006020828403121561134157600080fd5b8151610d0181610ccd565b60008261135b5761135b611235565b500490565b634e487b7160e01b600052603260045260246000fd5b600061ffff80831681810361138d5761138d61125f565b600101939250505056fea164736f6c634300080f000a",
}

var BeaconVRFConsumerABI = BeaconVRFConsumerMetaData.ABI

var BeaconVRFConsumerBin = BeaconVRFConsumerMetaData.Bin

func DeployBeaconVRFConsumer(auth *bind.TransactOpts, backend bind.ContractBackend, coordinator common.Address, shouldFail bool, beaconPeriodBlocks *big.Int) (common.Address, *types.Transaction, *BeaconVRFConsumer, error) {
	parsed, err := BeaconVRFConsumerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeaconVRFConsumerBin), backend, coordinator, shouldFail, beaconPeriodBlocks)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeaconVRFConsumer{BeaconVRFConsumerCaller: BeaconVRFConsumerCaller{contract: contract}, BeaconVRFConsumerTransactor: BeaconVRFConsumerTransactor{contract: contract}, BeaconVRFConsumerFilterer: BeaconVRFConsumerFilterer{contract: contract}}, nil
}

type BeaconVRFConsumer struct {
	BeaconVRFConsumerCaller
	BeaconVRFConsumerTransactor
	BeaconVRFConsumerFilterer
}

type BeaconVRFConsumerCaller struct {
	contract *bind.BoundContract
}

type BeaconVRFConsumerTransactor struct {
	contract *bind.BoundContract
}

type BeaconVRFConsumerFilterer struct {
	contract *bind.BoundContract
}

type BeaconVRFConsumerSession struct {
	Contract     *BeaconVRFConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BeaconVRFConsumerCallerSession struct {
	Contract *BeaconVRFConsumerCaller
	CallOpts bind.CallOpts
}

type BeaconVRFConsumerTransactorSession struct {
	Contract     *BeaconVRFConsumerTransactor
	TransactOpts bind.TransactOpts
}

type BeaconVRFConsumerRaw struct {
	Contract *BeaconVRFConsumer
}

type BeaconVRFConsumerCallerRaw struct {
	Contract *BeaconVRFConsumerCaller
}

type BeaconVRFConsumerTransactorRaw struct {
	Contract *BeaconVRFConsumerTransactor
}

func NewBeaconVRFConsumer(address common.Address, backend bind.ContractBackend) (*BeaconVRFConsumer, error) {
	contract, err := bindBeaconVRFConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumer{BeaconVRFConsumerCaller: BeaconVRFConsumerCaller{contract: contract}, BeaconVRFConsumerTransactor: BeaconVRFConsumerTransactor{contract: contract}, BeaconVRFConsumerFilterer: BeaconVRFConsumerFilterer{contract: contract}}, nil
}

func NewBeaconVRFConsumerCaller(address common.Address, caller bind.ContractCaller) (*BeaconVRFConsumerCaller, error) {
	contract, err := bindBeaconVRFConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerCaller{contract: contract}, nil
}

func NewBeaconVRFConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconVRFConsumerTransactor, error) {
	contract, err := bindBeaconVRFConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerTransactor{contract: contract}, nil
}

func NewBeaconVRFConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconVRFConsumerFilterer, error) {
	contract, err := bindBeaconVRFConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconVRFConsumerFilterer{contract: contract}, nil
}

func bindBeaconVRFConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BeaconVRFConsumerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_BeaconVRFConsumer *BeaconVRFConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconVRFConsumer.Contract.BeaconVRFConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.BeaconVRFConsumerTransactor.contract.Transfer(opts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.BeaconVRFConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconVRFConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.contract.Transfer(opts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) Fail(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "fail")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) Fail() (bool, error) {
	return _BeaconVRFConsumer.Contract.Fail(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) Fail() (bool, error) {
	return _BeaconVRFConsumer.Contract.Fail(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) IBeaconPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "i_beaconPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.IBeaconPeriodBlocks(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) IBeaconPeriodBlocks() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.IBeaconPeriodBlocks(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SExpectedRandomnessByBlock(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_ExpectedRandomnessByBlock", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SExpectedRandomnessByBlock(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SExpectedRandomnessByBlock(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SExpectedRandomnessByBlock(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SExpectedRandomnessByBlock(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SExpectedSeeds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_ExpectedSeeds", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SExpectedSeeds(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _BeaconVRFConsumer.Contract.SExpectedSeeds(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SExpectedSeeds(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _BeaconVRFConsumer.Contract.SExpectedSeeds(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SReceivedRandomnessByRequestID(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_ReceivedRandomnessByRequestID", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SReceivedRandomnessByRequestID(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SReceivedRandomnessByRequestID(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SReceivedRandomnessByRequestID(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SReceivedRandomnessByRequestID(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SGasAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_gasAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SGasAvailable() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SGasAvailable(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SGasAvailable() (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SGasAvailable(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SMyBeaconRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_myBeaconRequests", arg0)

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

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SMyBeaconRequests(arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	return _BeaconVRFConsumer.Contract.SMyBeaconRequests(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SMyBeaconRequests(arg0 *big.Int) (struct {
	SlotNumber        uint32
	ConfirmationDelay *big.Int
	NumWords          uint16
	Requester         common.Address
}, error) {
	return _BeaconVRFConsumer.Contract.SMyBeaconRequests(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_randomWords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRandomWords(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRandomWords(&_BeaconVRFConsumer.CallOpts, arg0)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SRequestsIDs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_requestsIDs", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SRequestsIDs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRequestsIDs(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SRequestsIDs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _BeaconVRFConsumer.Contract.SRequestsIDs(&_BeaconVRFConsumer.CallOpts, arg0, arg1)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCaller) SSubId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BeaconVRFConsumer.contract.Call(opts, &out, "s_subId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SSubId() (uint64, error) {
	return _BeaconVRFConsumer.Contract.SSubId(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerCallerSession) SSubId() (uint64, error) {
	return _BeaconVRFConsumer.Contract.SSubId(&_BeaconVRFConsumer.CallOpts)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) ComputeAndStoreExpectedOutput(opts *bind.TransactOpts, height uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "computeAndStoreExpectedOutput", height, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) ComputeAndStoreExpectedOutput(height uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.ComputeAndStoreExpectedOutput(&_BeaconVRFConsumer.TransactOpts, height, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) ComputeAndStoreExpectedOutput(height uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.ComputeAndStoreExpectedOutput(&_BeaconVRFConsumer.TransactOpts, height, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "rawFulfillRandomWords", requestID, randomWords, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.RawFulfillRandomWords(&_BeaconVRFConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) RawFulfillRandomWords(requestID *big.Int, randomWords []*big.Int, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.RawFulfillRandomWords(&_BeaconVRFConsumer.TransactOpts, requestID, randomWords, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) SetExpectedSeed(opts *bind.TransactOpts, height uint64, confirmationDelayArg *big.Int, vrfOutput ECCArithmeticG1Point) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "setExpectedSeed", height, confirmationDelayArg, vrfOutput)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SetExpectedSeed(height uint64, confirmationDelayArg *big.Int, vrfOutput ECCArithmeticG1Point) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.SetExpectedSeed(&_BeaconVRFConsumer.TransactOpts, height, confirmationDelayArg, vrfOutput)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) SetExpectedSeed(height uint64, confirmationDelayArg *big.Int, vrfOutput ECCArithmeticG1Point) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.SetExpectedSeed(&_BeaconVRFConsumer.TransactOpts, height, confirmationDelayArg, vrfOutput)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) SetFail(opts *bind.TransactOpts, shouldFail bool) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "setFail", shouldFail)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) SetFail(shouldFail bool) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.SetFail(&_BeaconVRFConsumer.TransactOpts, shouldFail)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) SetFail(shouldFail bool) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.SetFail(&_BeaconVRFConsumer.TransactOpts, shouldFail)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) TestRedeemRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "testRedeemRandomness", requestID)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) TestRedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRedeemRandomness(&_BeaconVRFConsumer.TransactOpts, requestID)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) TestRedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRedeemRandomness(&_BeaconVRFConsumer.TransactOpts, requestID)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) TestRequestRandomness(opts *bind.TransactOpts, numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "testRequestRandomness", numWords, subID, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) TestRequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomness(&_BeaconVRFConsumer.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) TestRequestRandomness(numWords uint16, subID uint64, confirmationDelayArg *big.Int) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomness(&_BeaconVRFConsumer.TransactOpts, numWords, subID, confirmationDelayArg)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactor) TestRequestRandomnessFulfillment(opts *bind.TransactOpts, subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.contract.Transact(opts, "testRequestRandomnessFulfillment", subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerSession) TestRequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomnessFulfillment(&_BeaconVRFConsumer.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
}

func (_BeaconVRFConsumer *BeaconVRFConsumerTransactorSession) TestRequestRandomnessFulfillment(subID uint64, numWords uint16, confirmationDelayArg *big.Int, callbackGasLimit uint32, arguments []byte) (*types.Transaction, error) {
	return _BeaconVRFConsumer.Contract.TestRequestRandomnessFulfillment(&_BeaconVRFConsumer.TransactOpts, subID, numWords, confirmationDelayArg, callbackGasLimit, arguments)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractDKGClient\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"DKGClientError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"indexed\":false,\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"KeyGenerated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"addClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_configDigest\",\"type\":\"bytes32\"}],\"name\":\"getKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"},{\"internalType\":\"contractDKGClient\",\"name\":\"clientAddress\",\"type\":\"address\"}],\"name\":\"removeClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_uint8\",\"type\":\"uint8\"}],\"name\":\"toASCII\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b503380600081620000695760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156200009c576200009c81620000a5565b50505062000150565b336001600160a01b03821603620000ff5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000060565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61289380620001606000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063b1dc65a411610066578063b1dc65a414610223578063c3105a6b14610236578063e3d0e71214610256578063f2fde38b1461026957600080fd5b80638da5cb5b146101cb5780639201de55146101e6578063afcb95d7146101f957600080fd5b80635429a79e116100c85780635429a79e1461016e57806379ba5097146101835780637bf1ffc51461018b57806381ff70481461019e57600080fd5b80630bc643e8146100ef578063181f5a771461011957806339614e4f1461015b575b600080fd5b6101026100fd366004611def565b61027c565b60405160ff90911681526020015b60405180910390f35b60408051808201909152600981527f444b4720302e302e31000000000000000000000000000000000000000000000060208201525b6040516101109190611e66565b61014e610169366004611f3e565b6102ab565b61018161017c366004611f90565b61042d565b005b61018161066e565b610181610199366004611f90565b610724565b600754600554604080516000815264010000000090930463ffffffff166020840152820152606001610110565b6000546040516001600160a01b039091168152602001610110565b61014e6101f4366004611fc0565b61076b565b6005546004546040805160008152602081019390935263ffffffff90911690820152606001610110565b610181610231366004612025565b6107f7565b61024961024436600461210a565b610943565b604051610110919061212c565b610181610264366004612234565b610a6b565b610181610277366004612301565b6111e7565b6000600a8260ff16101561029b57610295826030612334565b92915050565b610295826057612334565b919050565b6060600080835160026102be9190612359565b67ffffffffffffffff8111156102d6576102d6611e79565b6040519080825280601f01601f191660200182016040528015610300576020820181803683370190505b509050600091505b80518260ff16101561042657600084610322600285612378565b60ff1681518110610335576103356123a8565b60209101015160f81c600f1690506000600486610353600287612378565b60ff1681518110610366576103666123a8565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c905061039d8161027c565b60f81b838560ff16815181106103b5576103b56123a8565b60200101906001600160f81b031916908160001a9053506103d7846001612334565b93506103e28261027c565b60f81b838560ff16815181106103fa576103fa6123a8565b60200101906001600160f81b031916908160001a9053505050818061041e906123be565b925050610308565b9392505050565b6104356111fb565b60008281526002602090815260408083208054825181850281018501909352808352919290919083018282801561049557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610477575b505050505090506000815167ffffffffffffffff8111156104b8576104b8611e79565b6040519080825280602002602001820160405280156104e1578160200160208202803683370190505b5090506000805b835181101561058457846001600160a01b031684828151811061050d5761050d6123a8565b60200260200101516001600160a01b03161461056457848361052f84846123dd565b8151811061053f5761053f6123a8565b60200260200101906001600160a01b031690816001600160a01b031681525050610572565b8161056e816123f4565b9250505b8061057c816123f4565b9150506104e8565b50600081845161059491906123dd565b67ffffffffffffffff8111156105ac576105ac611e79565b6040519080825280602002602001820160405280156105d5578160200160208202803683370190505b50905060005b8285516105e891906123dd565b81101561064557838181518110610601576106016123a8565b602002602001015182828151811061061b5761061b6123a8565b6001600160a01b03909216602092830291909101909101528061063d816123f4565b9150506105db565b506000868152600260209081526040909120825161066592840190611d0a565b50505050505050565b6001546001600160a01b031633146106cd5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61072c6111fb565b600091825260026020908152604083208054600181018255908452922090910180546001600160a01b0319166001600160a01b03909216919091179055565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156107ed578381602081106107ae576107ae6123a8565b1a60f81b8282815181106107c4576107c46123a8565b60200101906001600160f81b031916908160001a905350806107e5816123f4565b915050610793565b50610426816102ab565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916108479184918491908e908e908190840183828082843760009201919091525061125792505050565b6040805183815263ffffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260055480825260065460ff808216602085015261010090910416928201929092529083146109035760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016106c4565b6109118b8b8b8b8b8b6114b1565b6109228c8c8c8c8c8c8c8c89611545565b50505063ffffffff81106109385761093861240d565b505050505050505050565b6040805180820190915260608082526020820152600083815260036020908152604080832085845290915290819020815180830190925280548290829061098990612423565b80601f01602080910402602001604051908101604052809291908181526020018280546109b590612423565b8015610a025780601f106109d757610100808354040283529160200191610a02565b820191906000526020600020905b8154815290600101906020018083116109e557829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610a5a57602002820191906000526020600020905b815481526020019060010190808311610a46575b505050505081525050905092915050565b855185518560ff16601f831115610ac45760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016106c4565b60008111610b145760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016106c4565b818314610b885760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016106c4565b610b93816003612359565b8311610be15760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016106c4565b610be96111fb565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60095415610d3957600954600090610c41906001906123dd565b9050600060098281548110610c5857610c586123a8565b6000918252602082200154600a80546001600160a01b0390921693509084908110610c8557610c856123a8565b60009182526020808320909101546001600160a01b03858116845260089092526040808420805461ffff1990811690915592909116808452922080549091169055600980549192509080610cdb57610cdb61245d565b600082815260209020810160001990810180546001600160a01b0319169055019055600a805480610d0e57610d0e61245d565b600082815260209020810160001990810180546001600160a01b031916905501905550610c27915050565b60005b8151518110156110785760006008600084600001518481518110610d6257610d626123a8565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610d9f57610d9f612473565b14610dec5760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016106c4565b6040805180820190915260ff82168152600160208201528251805160089160009185908110610e1d57610e1d6123a8565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610e7657610e76612473565b021790555060009150610e869050565b6008600084602001518481518110610ea057610ea06123a8565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115610edd57610edd612473565b14610f2a5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016106c4565b6040805180820190915260ff821681526020810160028152506008600084602001518481518110610f5d57610f5d6123a8565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610fb657610fb6612473565b021790555050825180516009925083908110610fd457610fd46123a8565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b03909316929092179091558201518051600a91908390811061102b5761102b6123a8565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b0390921691909117905580611070816123f4565b915050610d3c565b5060408101516006805460ff191660ff9092169190911790556007805467ffffffff0000000019811664010000000063ffffffff4381168202928317855590830481169360019390926000926110d5928692908216911617612489565b92506101000a81548163ffffffff021916908363ffffffff16021790555060006111364630600760009054906101000a900463ffffffff1663ffffffff1686600001518760200151886040015189606001518a608001518b60a001516119d4565b6005819055835180516006805460ff9092166101000261ff00199092169190911790556007546020860151604080880151606089015160808a015160a08b015193519798507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05976111be978b978b9763ffffffff9091169691959094909390929091906124f5565b60405180910390a16111d98360400151846060015183611a61565b505050505050505050505050565b6111ef6111fb565b6111f881611c61565b50565b6000546001600160a01b031633146112555760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106c4565b565b600060608083806020019051810190611270919061258b565b60408051808201825283815260208082018490526000868152600282528381208054855181850281018501909652808652979a50959850939650909492939192908301828280156112ea57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112cc575b5050505050905060005b81518110156114075781818151811061130f5761130f6123a8565b60200260200101516001600160a01b031663bf2732c7846040518263ffffffff1660e01b8152600401611342919061212c565b600060405180830381600087803b15801561135c57600080fd5b505af192505050801561136d575060015b6113f5573d80801561139b576040519150601f19603f3d011682016040523d82523d6000602084013e6113a0565b606091505b507f116391732f5df106193bda7cedf1728f3b07b62f6cdcdd611c9eeec44efcae548383815181106113d4576113d46123a8565b6020026020010151826040516113eb929190612689565b60405180910390a1505b806113ff816123f4565b9150506112f4565b5060008581526003602090815260408083208b845290915290208251839190819061143290826126fa565b50602082810151805161144b9260018501920190611d6f565b5090505084887fc8db841f5b2231ccf7190311f440aa197b161e369f3b40b023508160cc55565684604051611480919061212c565b60405180910390a350506004805460089690961c63ffffffff1663ffffffff19909616959095179094555050505050565b60006114be826020612359565b6114c9856020612359565b6114d5886101446127ba565b6114df91906127ba565b6114e991906127ba565b6114f49060006127ba565b90503681146106655760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016106c4565b600060028260200151836040015161155d9190612334565b6115679190612378565b611572906001612334565b60408051600180825281830190925260ff929092169250600091906020820181803683370190505090508160f81b816000815181106115b3576115b36123a8565b60200101906001600160f81b031916908160001a9053508682146115d6826102ab565b906115f45760405162461bcd60e51b81526004016106c49190611e66565b508685146116445760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016106c4565b3360009081526008602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561168757611687612473565b600281111561169857611698612473565b90525090506002816020015160028111156116b5576116b5612473565b1480156116ef5750600a816000015160ff16815481106116d7576116d76123a8565b6000918252602090912001546001600160a01b031633145b61173b5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016106c4565b505050600088886040516117509291906127d2565b604051908190038120611767918c906020016127e2565b604051602081830303815290604052805190602001209050611787611daa565b604080518082019091526000808252602082015260005b888110156119c55760006001858884602081106117bd576117bd6123a8565b6117ca91901a601b612334565b8d8d868181106117dc576117dc6123a8565b905060200201358c8c878181106117f5576117f56123a8565b9050602002013560405160008152602001604052604051611832949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611854573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526008602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156118a9576118a9612473565b60028111156118ba576118ba612473565b90525092506001836020015160028111156118d7576118d7612473565b146119245760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016106c4565b8251849060ff16601f811061193b5761193b6123a8565b60200201511561198d5760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016106c4565b600184846000015160ff16601f81106119a8576119a86123a8565b9115156020909202015250806119bd816123f4565b91505061179e565b50505050505050505050505050565b6000808a8a8a8a8a8a8a8a8a6040516020016119f8999897969594939291906127fe565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6000808351602014611ab55760405162461bcd60e51b815260206004820152601e60248201527f77726f6e67206c656e67746820666f72206f6e636861696e436f6e666967000060448201526064016106c4565b60208401519150808203611b0b5760405162461bcd60e51b815260206004820152601460248201527f6661696c656420746f20636f7079206b6579494400000000000000000000000060448201526064016106c4565b60408051808201909152606080825260208201526000838152600360209081526040808320878452909152902081518291908190611b4990826126fa565b506020828101518051611b629260018501920190611d6f565b505050600083815260026020908152604080832080548251818502810185019093528083529192909190830182828015611bc557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611ba7575b5050505050905060005b8151811015611c5757818181518110611bea57611bea6123a8565b60200260200101516001600160a01b03166355e487496040518163ffffffff1660e01b8152600401600060405180830381600087803b158015611c2c57600080fd5b505af1158015611c40573d6000803e3d6000fd5b505050508080611c4f906123f4565b915050611bcf565b5050505050505050565b336001600160a01b03821603611cb95760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106c4565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215611d5f579160200282015b82811115611d5f57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611d2a565b50611d6b929150611dc9565b5090565b828054828255906000526020600020908101928215611d5f579160200282015b82811115611d5f578251825591602001919060010190611d8f565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115611d6b5760008155600101611dca565b803560ff811681146102a657600080fd5b600060208284031215611e0157600080fd5b61042682611dde565b60005b83811015611e25578181015183820152602001611e0d565b83811115611e34576000848401525b50505050565b60008151808452611e52816020860160208601611e0a565b601f01601f19169290920160200192915050565b6020815260006104266020830184611e3a565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611eb857611eb8611e79565b604052919050565b600067ffffffffffffffff821115611eda57611eda611e79565b50601f01601f191660200190565b600082601f830112611ef957600080fd5b8135611f0c611f0782611ec0565b611e8f565b818152846020838601011115611f2157600080fd5b816020850160208301376000918101602001919091529392505050565b600060208284031215611f5057600080fd5b813567ffffffffffffffff811115611f6757600080fd5b611f7384828501611ee8565b949350505050565b6001600160a01b03811681146111f857600080fd5b60008060408385031215611fa357600080fd5b823591506020830135611fb581611f7b565b809150509250929050565b600060208284031215611fd257600080fd5b5035919050565b60008083601f840112611feb57600080fd5b50813567ffffffffffffffff81111561200357600080fd5b6020830191508360208260051b850101111561201e57600080fd5b9250929050565b60008060008060008060008060e0898b03121561204157600080fd5b606089018a81111561205257600080fd5b8998503567ffffffffffffffff8082111561206c57600080fd5b818b0191508b601f83011261208057600080fd5b81358181111561208f57600080fd5b8c60208285010111156120a157600080fd5b6020830199508098505060808b01359150808211156120bf57600080fd5b6120cb8c838d01611fd9565b909750955060a08b01359150808211156120e457600080fd5b506120f18b828c01611fd9565b999c989b50969995989497949560c00135949350505050565b6000806040838503121561211d57600080fd5b50508035926020909101359150565b6000602080835283516040828501526121486060850182611e3a565b85830151858203601f19016040870152805180835290840192506000918401905b808310156121895783518252928401926001929092019190840190612169565b509695505050505050565b600067ffffffffffffffff8211156121ae576121ae611e79565b5060051b60200190565b600082601f8301126121c957600080fd5b813560206121d9611f0783612194565b82815260059290921b840181019181810190868411156121f857600080fd5b8286015b8481101561218957803561220f81611f7b565b83529183019183016121fc565b803567ffffffffffffffff811681146102a657600080fd5b60008060008060008060c0878903121561224d57600080fd5b863567ffffffffffffffff8082111561226557600080fd5b6122718a838b016121b8565b9750602089013591508082111561228757600080fd5b6122938a838b016121b8565b96506122a160408a01611dde565b955060608901359150808211156122b757600080fd5b6122c38a838b01611ee8565b94506122d160808a0161221c565b935060a08901359150808211156122e757600080fd5b506122f489828a01611ee8565b9150509295509295509295565b60006020828403121561231357600080fd5b813561042681611f7b565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff038211156123515761235161231e565b019392505050565b60008160001904831182151516156123735761237361231e565b500290565b600060ff83168061239957634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036123d4576123d461231e565b60010192915050565b6000828210156123ef576123ef61231e565b500390565b6000600182016124065761240661231e565b5060010190565b634e487b7160e01b600052600160045260246000fd5b600181811c9082168061243757607f821691505b60208210810361245757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b600063ffffffff8083168185168083038211156124a8576124a861231e565b01949350505050565b600081518084526020808501945080840160005b838110156124ea5781516001600160a01b0316875295820195908201906001016124c5565b509495945050505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526125258184018a6124b1565b9050828103608084015261253981896124b1565b905060ff871660a084015282810360c08401526125568187611e3a565b905067ffffffffffffffff851660e084015282810361010084015261257b8185611e3a565b9c9b505050505050505050505050565b6000806000606084860312156125a057600080fd5b8351925060208085015167ffffffffffffffff808211156125c057600080fd5b818701915087601f8301126125d457600080fd5b81516125e2611f0782611ec0565b81815289858386010111156125f657600080fd5b61260582868301878701611e0a565b60408901519096509250508082111561261d57600080fd5b508501601f8101871361262f57600080fd5b805161263d611f0782612194565b81815260059190911b8201830190838101908983111561265c57600080fd5b928401925b8284101561267a57835182529284019290840190612661565b80955050505050509250925092565b6001600160a01b0383168152604060208201526000611f736040830184611e3a565b601f8211156126f557600081815260208120601f850160051c810160208610156126d25750805b601f850160051c820191505b818110156126f1578281556001016126de565b5050505b505050565b815167ffffffffffffffff81111561271457612714611e79565b612728816127228454612423565b846126ab565b602080601f83116001811461275d57600084156127455750858301515b600019600386901b1c1916600185901b1785556126f1565b600085815260208120601f198616915b8281101561278c5788860151825594840194600190910190840161276d565b50858210156127aa5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600082198211156127cd576127cd61231e565b500190565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526128388285018b6124b1565b9150838203608085015261284c828a6124b1565b915060ff881660a085015283820360c08501526128698288611e3a565b90861660e0850152838103610100850152905061257b8185611e3a56fea164736f6c634300080f000a",
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
	Bin: "0x608060405234801561001057600080fd5b50610518806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630bc643e81461004657806339614e4f146100705780639201de5514610090575b600080fd5b6100596100543660046102db565b6100a3565b60405160ff90911681526020015b60405180910390f35b61008361007e366004610314565b6100cd565b60405161006791906103c5565b61008361009e36600461041a565b61024f565b6000600a8260ff1610156100c2576100bc826030610449565b92915050565b6100bc826057610449565b6060600080835160026100e0919061046e565b67ffffffffffffffff8111156100f8576100f86102fe565b6040519080825280601f01601f191660200182016040528015610122576020820181803683370190505b509050600091505b80518260ff1610156102485760008461014460028561048d565b60ff1681518110610157576101576104bd565b60209101015160f81c600f169050600060048661017560028761048d565b60ff1681518110610188576101886104bd565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016901c60f81c90506101bf816100a3565b60f81b838560ff16815181106101d7576101d76104bd565b60200101906001600160f81b031916908160001a9053506101f9846001610449565b9350610204826100a3565b60f81b838560ff168151811061021c5761021c6104bd565b60200101906001600160f81b031916908160001a90535050508180610240906104d3565b92505061012a565b9392505050565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b60208110156102d157838160208110610292576102926104bd565b1a60f81b8282815181106102a8576102a86104bd565b60200101906001600160f81b031916908160001a905350806102c9816104f2565b915050610277565b50610248816100cd565b6000602082840312156102ed57600080fd5b813560ff8116811461024857600080fd5b634e487b7160e01b600052604160045260246000fd5b60006020828403121561032657600080fd5b813567ffffffffffffffff8082111561033e57600080fd5b818401915084601f83011261035257600080fd5b813581811115610364576103646102fe565b604051601f8201601f19908116603f0116810190838211818310171561038c5761038c6102fe565b816040528281528760208487010111156103a557600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156103f2578581018301518582016040015282016103d6565b81811115610404576000604083870101525b50601f01601f1916929092016040019392505050565b60006020828403121561042c57600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff84168060ff0382111561046657610466610433565b019392505050565b600081600019048311821515161561048857610488610433565b500290565b600060ff8316806104ae57634e487b7160e01b600052601260045260246000fd5b8060ff84160491505092915050565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036104e9576104e9610433565b60010192915050565b60006001820161050457610504610433565b506001019056fea164736f6c634300080f000a",
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

var IVRFBeaconConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var IVRFBeaconConsumerABI = IVRFBeaconConsumerMetaData.ABI

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

var VRFBeaconBillingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"forgetConsumerSubscriptionID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161061438038061061483398101604081905261002f91610172565b33806000816100855760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03848116919091179091558116156100b5576100b5816100c9565b5050506001600160a01b03166080526101a2565b336001600160a01b038216036101215760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007c565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561018457600080fd5b81516001600160a01b038116811461019b57600080fd5b9392505050565b6080516104586101bc6000396000607101526104586000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80638da5cb5b116100505780638da5cb5b146100b95780639e3616f4146100ca578063f2fde38b146100dd57600080fd5b80631b6b6d231461006c57806379ba5097146100af575b600080fd5b6100937f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b6100b76100f0565b005b6000546001600160a01b0316610093565b6100b76100d8366004610369565b6101b3565b6100b76100eb3660046103de565b610243565b6001546001600160a01b0316331461014f5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6101bb610257565b60005b8181101561023e576000600560008585858181106101de576101de61040e565b90506020020160208101906101f391906103de565b6001600160a01b031681526020810191909152604001600020805467ffffffffffffffff191667ffffffffffffffff929092169190911790558061023681610424565b9150506101be565b505050565b61024b610257565b610254816102b3565b50565b6000546001600160a01b031633146102b15760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610146565b565b336001600160a01b0382160361030b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610146565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000806020838503121561037c57600080fd5b823567ffffffffffffffff8082111561039457600080fd5b818501915085601f8301126103a857600080fd5b8135818111156103b757600080fd5b8660208260051b85010111156103cc57600080fd5b60209290920196919550909350505050565b6000602082840312156103f057600080fd5b81356001600160a01b038116811461040757600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161044457634e487b7160e01b600052601160045260246000fd5b506001019056fea164736f6c634300080f000a",
}

var VRFBeaconBillingABI = VRFBeaconBillingMetaData.ABI

var VRFBeaconBillingBin = VRFBeaconBillingMetaData.Bin

func DeployVRFBeaconBilling(auth *bind.TransactOpts, backend bind.ContractBackend, link common.Address) (common.Address, *types.Transaction, *VRFBeaconBilling, error) {
	parsed, err := VRFBeaconBillingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconBillingBin), backend, link)
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

func (_VRFBeaconBilling *VRFBeaconBillingCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconBilling.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconBilling *VRFBeaconBillingSession) LINK() (common.Address, error) {
	return _VRFBeaconBilling.Contract.LINK(&_VRFBeaconBilling.CallOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingCallerSession) LINK() (common.Address, error) {
	return _VRFBeaconBilling.Contract.LINK(&_VRFBeaconBilling.CallOpts)
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

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AcceptOwnership(&_VRFBeaconBilling.TransactOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.AcceptOwnership(&_VRFBeaconBilling.TransactOpts)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactor) ForgetConsumerSubscriptionID(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.contract.Transact(opts, "forgetConsumerSubscriptionID", consumers)
}

func (_VRFBeaconBilling *VRFBeaconBillingSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconBilling.TransactOpts, consumers)
}

func (_VRFBeaconBilling *VRFBeaconBillingTransactorSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconBilling.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconBilling.TransactOpts, consumers)
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

var VRFBeaconCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"contractDKG\",\"name\":\"keyProvider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"keyID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keyProvider\",\"type\":\"address\"}],\"name\":\"KeyInfoMustComeFromProvider\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OffchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"occVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconReport.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconReport.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"forgetConsumerSubscriptionID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfirmationDelays\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"\",\"type\":\"uint24[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"internalType\":\"structKeyDataStruct.KeyData\",\"name\":\"kd\",\"type\":\"tuple\"}],\"name\":\"keyGenerated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxErrorMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newKeyRequested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005c7538038062005c7583398101604081905262000034916200022f565b8181848681818181803380600081620000945760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c757620000c7816200016b565b5050506001600160a01b03166080526000829003620000f957604051632abc297960e01b815260040160405180910390fd5b60a082905260006200010c83436200027d565b905060008160a051620001209190620002b6565b90506200012e8143620002d0565b60c0525050601d80546001600160a01b0319166001600160a01b039990991698909817909755505050601e9290925550620002eb95505050505050565b336001600160a01b03821603620001c55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008b565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6001600160a01b03811681146200022c57600080fd5b50565b600080600080608085870312156200024657600080fd5b8451620002538162000216565b6020860151604087015191955093506200026d8162000216565b6060959095015193969295505050565b6000826200029b57634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b600082821015620002cb57620002cb620002a0565b500390565b60008219821115620002e657620002e6620002a0565b500190565b60805160a05160c05161590d62000368600039600061058a01526000818161056301528181610b1c01528181613611015281816136400152818161367801526140140152600081816102e601528181611715015281816117dc015281816119050152818161252d01528181612b7e0152612cc5015261590d6000f3fe608060405234801561001057600080fd5b50600436106102775760003560e01c8063b1dc65a411610160578063cf7e754a116100d8578063e4902f821161008c578063f2fde38b11610071578063f2fde38b14610635578063f645dcb114610648578063fbffd2c11461065b57600080fd5b8063e4902f82146105fa578063eb5dcd6c1461062257600080fd5b8063d57fc45a116100bd578063d57fc45a146105b4578063dc92accf146105bd578063e3d0e712146105e757600080fd5b8063cf7e754a14610585578063d09dc339146105ac57600080fd5b8063c278e5b71161012f578063c63c4e9b11610114578063c63c4e9b1461053a578063cc31f7dd14610555578063cd0593df1461055e57600080fd5b8063c278e5b714610518578063c4c92b371461052957600080fd5b8063b1dc65a4146104d6578063bbcdd0d8146104e9578063bf2732c7146104f2578063c10753291461050557600080fd5b80637a464944116101f35780638da5cb5b116101c25780639e3616f4116101a75780639e3616f414610486578063afcb95d714610499578063b121e147146104c357600080fd5b80638da5cb5b146104625780639c849b301461047357600080fd5b80637a4649441461040557806381ff70481461040d57806385c64e111461043a5780638ac28d5a1461044f57600080fd5b80632f7527cc1161024a578063643dc1051161022f578063643dc105146103ca57806374d84611146103dd57806379ba5097146103fd57600080fd5b80632f7527cc146103a657806355e48749146103c057600080fd5b80630eafb25b1461027c578063181f5a77146102a25780631b6b6d23146102e15780632993726814610320575b600080fd5b61028f61028a3660046144ef565b61066e565b6040519081526020015b60405180910390f35b604080518082018252601581527f565246426561636f6e20312e302e302d616c7068610000000000000000000000602082015290516102999190614564565b6103087f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610299565b61036a600c546a0100000000000000000000810463ffffffff90811692600160701b8304821692600160901b8104831692600160b01b82041691600160d01b90910462ffffff1690565b6040805163ffffffff9687168152948616602086015292851692840192909252909216606082015262ffffff909116608082015260a001610299565b6103ae600881565b60405160ff9091168152602001610299565b6103c8610776565b005b6103c86103d83660046145a1565b6107c0565b6103f06103eb36600461461e565b610a1e565b6040516102999190614676565b6103c8610c22565b61028f608081565b600d54600f54604080516000815264010000000090930463ffffffff166020840152820152606001610299565b610442610cd3565b60405161029991906146b1565b6103c861045d3660046144ef565b610d38565b6000546001600160a01b0316610308565b6103c861048136600461470c565b610daa565b6103c8610494366004614778565b610f88565b600f546011546040805160008152602081019390935263ffffffff90911690820152606001610299565b6103c86104d13660046144ef565b611018565b6103c86104e43660046147fc565b6110f4565b61028f6103e881565b6103c8610500366004614a2e565b6115a4565b6103c8610513366004614b17565b611612565b6103c8610526366004614b43565b50565b601c546001600160a01b0316610308565b610542600381565b60405161ffff9091168152602001610299565b61028f601e5481565b61028f7f000000000000000000000000000000000000000000000000000000000000000081565b61028f7f000000000000000000000000000000000000000000000000000000000000000081565b61028f6118e3565b61028f601f5481565b6105d06105cb366004614baf565b61198f565b60405165ffffffffffff9091168152602001610299565b6103c86105f5366004614c0b565b611ac8565b61060d6106083660046144ef565b61220e565b60405163ffffffff9091168152602001610299565b6103c8610630366004614cf9565b6122c0565b6103c86106433660046144ef565b6123f8565b6105d0610656366004614d32565b612409565b6103c86106693660046144ef565b61250a565b6001600160a01b03811660009081526012602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906106d05750600092915050565b600c546020820151600091600160901b900463ffffffff169060169060ff16601f81106106ff576106ff614db8565b600881049190910154600c54610735926007166004026101000a90910463ffffffff908116916601000000000000900416614de4565b63ffffffff166107459190614e09565b61075390633b9aca00614e09565b905081604001516001600160601b03168161076e9190614e28565b949350505050565b601d546001600160a01b03163381146107b85760405163292f4fb560e01b81523360048201526001600160a01b03821660248201526044015b60405180910390fd5b506000601f55565b601c546001600160a01b03166107de6000546001600160a01b031690565b6001600160a01b0316336001600160a01b0316148061086a5750604051630d629b5f60e31b81526001600160a01b03821690636b14daf8906108299033906000903690600401614e69565b602060405180830381865afa158015610846573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086a9190614e8c565b6108b65760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c60448201526064016107af565b6108be61251b565b600c80547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff166a010000000000000000000063ffffffff8981169182027fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1692909217600160701b898416908102919091177fffffffffffff0000000000000000ffffffffffffffffffffffffffffffffffff16600160901b8985169081027fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff1691909117600160b01b948916948502177fffffff000000ffffffffffffffffffffffffffffffffffffffffffffffffffff16600160d01b62ffffff89169081029190911790955560408051938452602084019290925290820152606081019190915260808101919091527f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f9060a00160405180910390a1505050505050565b65ffffffffffff81166000818152600a602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff670100000000000000820416938201939093526001600160a01b03690100000000000000000084048116606083810191825298909752949093527fffffff000000000000000000000000000000000000000000000000000000000090911690559151163314610b115760608101516040517f8e30e8230000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201523360248201526044016107af565b8051600090610b47907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16614e09565b90506000826020015162ffffff1643610b609190614eae565b9050808210610ba4576040517f15ad27c3000000000000000000000000000000000000000000000000000000008152600481018390524360248201526044016107af565b67ffffffffffffffff821115610be9576040517f058ddf02000000000000000000000000000000000000000000000000000000008152600481018390526024016107af565b60008281526007602090815260408083208287015162ffffff168452909152902054610c199086908590856128cb565b95945050505050565b6001546001600160a01b03163314610c7c5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016107af565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610cdb614305565b6040805161010081019182905290600990600890826000855b82829054906101000a900462ffffff1662ffffff1681526020019060030190602082600201049283019260010382029150808411610cf45790505050505050905090565b6001600160a01b038181166000908152601a6020526040902054163314610da15760405162461bcd60e51b815260206004820152601760248201527f4f6e6c792070617965652063616e20776974686472617700000000000000000060448201526064016107af565b61052681612ad2565b610db2612d1c565b828114610e015760405162461bcd60e51b815260206004820181905260248201527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a6560448201526064016107af565b60005b83811015610f81576000858583818110610e2057610e20614db8565b9050602002016020810190610e3591906144ef565b90506000848484818110610e4b57610e4b614db8565b9050602002016020810190610e6091906144ef565b6001600160a01b038084166000908152601a60205260409020549192501680158080610e9d5750826001600160a01b0316826001600160a01b0316145b610ee95760405162461bcd60e51b815260206004820152601160248201527f706179656520616c72656164792073657400000000000000000000000000000060448201526064016107af565b6001600160a01b038481166000908152601a6020526040902080546001600160a01b03191685831690811790915590831614610f6a57826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b505050508080610f7990614ec5565b915050610e04565b5050505050565b610f90612d1c565b60005b8181101561101357600060056000858585818110610fb357610fb3614db8565b9050602002016020810190610fc891906144ef565b6001600160a01b031681526020810191909152604001600020805467ffffffffffffffff191667ffffffffffffffff929092169190911790558061100b81614ec5565b915050610f93565b505050565b6001600160a01b038181166000908152601b60205260409020541633146110815760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792070726f706f736564207061796565732063616e206163636570740060448201526064016107af565b6001600160a01b038181166000818152601a602090815260408083208054336001600160a01b03198083168217909355601b909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a604080516101008082018352600c5460ff808216845291810464ffffffffff166020808501919091526601000000000000820463ffffffff908116858701526a0100000000000000000000830481166060860152600160701b830481166080860152600160901b8304811660a0860152600160b01b83041660c0850152600160d01b90910462ffffff1660e08401523360009081526012825293909320549394509092918c013591166111ec5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016107af565b600f548b351461123e5760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016107af565b61124c8a8a8a8a8a8a612d78565b8151611259906001614ede565b60ff1687146112aa5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016107af565b8685146112f95760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016107af565b60008a8a60405161130b929190614f03565b604051908190038120611322918e90602001614f13565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a8110156114c85760006001858a846020811061137157611371614db8565b61137e91901a601b614ede565b8f8f8681811061139057611390614db8565b905060200201358e8e878181106113a9576113a9614db8565b90506020020135604051600081526020016040526040516113e6949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611408573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526013602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506114a15760405162461bcd60e51b815260206004820152600f60248201527f7369676e6174757265206572726f72000000000000000000000000000000000060448201526064016107af565b826020015160080260ff166001901b840193505080806114c090614ec5565b915050611352565b5081827e0101010101010101010101010101010101010101010101010101010101010116146115395760405162461bcd60e51b815260206004820152601060248201527f6475706c6963617465207369676e65720000000000000000000000000000000060448201526064016107af565b50600091506115889050838d836020020135848e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612e1592505050565b905061159683828633613266565b505050505050505050505050565b601d546001600160a01b03163381146115e15760405163292f4fb560e01b81523360048201526001600160a01b03821660248201526044016107af565b81516040516115f39190602001614f2f565b60408051601f198184030181529190528051602090910120601f555050565b6000546001600160a01b031633148061169c5750601c54604051630d629b5f60e31b81526001600160a01b0390911690636b14daf89061165b9033906000903690600401614e69565b602060405180830381865afa158015611678573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061169c9190614e8c565b6116e85760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c60448201526064016107af565b60006116f261338b565b6040516370a0823160e01b81523060048201529091506000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561175c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117809190614f4b565b9050818110156117d25760405162461bcd60e51b815260206004820152601460248201527f696e73756666696369656e742062616c616e636500000000000000000000000060448201526064016107af565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb8561181561180f8686614eae565b87613558565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015611878573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061189c9190614e8c565b6118dd5760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b60448201526064016107af565b50505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561194c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119709190614f4b565b9050600061197c61338b565b90506119888183614f64565b9250505090565b60008060008061199f8786613572565b92509250925065ffffffffffff83166000908152600a602090815260409182902084518154928601518487015160608801516001600160a01b03166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff90921667010000000000000002919091167fffffff00000000000000000000000000000000000000000000ffffffffffffff62ffffff9093166401000000000266ffffffffffffff1990961663ffffffff909416939093179490941716179190911790555167ffffffffffffffff8216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc590611ab390889062ffffff91909116815260200190565b60405180910390a250909150505b9392505050565b611ad0612d1c565b601f891115611b215760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79206f7261636c65730000000000000000000000000000000060448201526064016107af565b888714611b705760405162461bcd60e51b815260206004820152601660248201527f6f7261636c65206c656e677468206d69736d617463680000000000000000000060448201526064016107af565b88611b7c876003614fd8565b60ff1610611bcc5760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016107af565b611bd88660ff16613891565b6040805160e060208c02808301820190935260c082018c815260009383928f918f918291908601908490808284376000920191909152505050908252506040805160208c810282810182019093528c82529283019290918d918d91829185019084908082843760009201919091525050509082525060ff891660208083019190915260408051601f8a01839004830281018301825289815292019190899089908190840183828082843760009201919091525050509082525067ffffffffffffffff861660208083019190915260408051601f870183900483028101830182528681529201919086908690819084018382808284376000920191909152505050915250600c805465ffffffffff00191690559050611cf461251b565b60145460005b81811015611da557600060148281548110611d1757611d17614db8565b6000918252602082200154601580546001600160a01b0390921693509084908110611d4457611d44614db8565b60009182526020808320909101546001600160a01b039485168352601382526040808420805461ffff1916905594168252601290529190912080546dffffffffffffffffffffffffffff191690555080611d9d81614ec5565b915050611cfa565b50611db260146000614324565b611dbe60156000614324565b60005b82515181101561203c576013600084600001518381518110611de557611de5614db8565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611e595760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016107af565b604080518082019091526001815260ff821660208201528351805160139160009185908110611e8a57611e8a614db8565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051601292919084908110611ef557611ef5614db8565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611f695760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016107af565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b03168152506012600085602001518481518110611fae57611fae614db8565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516001600160601b031662010000026dffffffffffffffffffffffff00001960ff959095166101000261ff00199315159390931661ffff199094169390931791909117929092161790558061203481614ec5565b915050611dc1565b508151805161205391601491602090910190614342565b50602080830151805161206a926015920190614342565b506040820151600c805460ff191660ff909216919091179055600d805467ffffffff0000000019811664010000000063ffffffff438116820292831790945582048316926000926120c2929082169116176001615001565b905080600d60006101000a81548163ffffffff021916908363ffffffff160217905550600061211646308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a001516138e1565b905080600f600001819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05838284886000015189602001518a604001518b606001518c608001518d60a0015160405161217999989796959493929190615062565b60405180910390a1600c546601000000000000900463ffffffff1660005b8651518110156121f15781601682601f81106121b5576121b5614db8565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff16021790555080806121e990614ec5565b915050612197565b506121fc8b8b61396e565b50505050505050505050505050505050565b6001600160a01b03811660009081526012602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906122705750600092915050565b6016816020015160ff16601f811061228a5761228a614db8565b600881049190910154600c54611ac1926007166004026101000a90910463ffffffff908116916601000000000000900416614de4565b6001600160a01b038281166000908152601a60205260409020541633146123295760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792063757272656e742070617965652063616e2075706461746500000060448201526064016107af565b6001600160a01b03811633036123815760405162461bcd60e51b815260206004820152601760248201527f63616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107af565b6001600160a01b038083166000908152601b6020526040902080548383166001600160a01b031982168117909255909116908114611013576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b612400612d1c565b6105268161397c565b60008060006124188787613572565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a67ffffffffffffffff1681526020018763ffffffff166001600160601b0316815250905081878a8360405160200161249394939291906150f8565b60408051601f19818403018152828252805160209182012065ffffffffffff871660009081526006909252919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906124f59084908a908d9086906150f8565b60405180910390a15090979650505050505050565b612512612d1c565b61052681613a25565b600c54604080516103e08101918290527f0000000000000000000000000000000000000000000000000000000000000000926601000000000000900463ffffffff169160009190601690601f908285855b82829054906101000a900463ffffffff1663ffffffff168152602001906004019060208260030104928301926001038202915080841161256c579050505050505090506000601580548060200260200160405190810160405280929190818152602001828054801561260757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116125e9575b5050505050905060005b81518110156128bd5760006012600084848151811061263257612632614db8565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b0316905060006012600085858151811061269457612694614db8565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008483601f81106126f7576126f7614db8565b6020020151600c5490870363ffffffff9081169250600160901b909104168102633b9aca0002820180156128b2576000601a600087878151811061273d5761273d614db8565b6020908102919091018101516001600160a01b03908116835290820192909252604090810160002054905163a9059cbb60e01b815290821660048201819052602482018590529250908a169063a9059cbb906044016020604051808303816000875af11580156127b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127d59190614e8c565b6128165760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b60448201526064016107af565b878786601f811061282957612829614db8565b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b031687878151811061286657612866614db8565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c856040516128a891815260200190565b60405180910390a4505b505050600101612611565b50610f81601683601f6143a7565b60608261291e576040517fc7d41b1b00000000000000000000000000000000000000000000000000000000815265ffffffffffff8616600482015267ffffffffffffffff831660248201526044016107af565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff1611156129d4576040808601519051634a90778560e01b815261ffff90911660048201526103e860248201526044016107af565b6000856040015161ffff1667ffffffffffffffff8111156129f7576129f76148b3565b604051908082528060200260200182016040528015612a20578160200160208202803683370190505b50905060005b866040015161ffff168161ffff161015612ac7578281604051602001612a7b92919091825260f01b7fffff00000000000000000000000000000000000000000000000000000000000016602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff1681518110612aaa57612aaa614db8565b602090810291909101015280612abf8161519c565b915050612a26565b509695505050505050565b6001600160a01b0381166000908152601260209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b031692810192909252612b2f575050565b6000612b3a8361066e565b90508015611013576001600160a01b038381166000908152601a60205260409081902054905163a9059cbb60e01b81529082166004820181905260248201849052917f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af1158015612bc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612beb9190614e8c565b612c2c5760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b60448201526064016107af565b600c60000160069054906101000a900463ffffffff166016846020015160ff16601f8110612c5c57612c5c614db8565b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b0384811660008181526012602090815260409182902080546dffffffffffffffffffffffff00001916905590518581527f0000000000000000000000000000000000000000000000000000000000000000841693851692917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c910160405180910390a450505050565b6000546001600160a01b03163314612d765760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107af565b565b6000612d85826020614e09565b612d90856020614e09565b612d9c88610144614e28565b612da69190614e28565b612db09190614e28565b612dbb906000614e28565b9050368114612e0c5760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016107af565b50505050505050565b60008082806020019051810190612e2c91906153a0565b64ffffffffff85166020880152604087018051919250612e4b82615575565b63ffffffff1663ffffffff168152505085600c60008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff1602179055509050506000816040015167ffffffffffffffff164090508082606001511461300857606082015160408084015190517faed0afe500000000000000000000000000000000000000000000000000000000815260048101929092526024820183905267ffffffffffffffff1660448201526064016107af565b60008083600001515167ffffffffffffffff811115613029576130296148b3565b60405190808252806020026020018201604052801561306e57816020015b60408051808201909152600080825260208201528152602001906001900390816130475790505b50905060005b84515181101561313f5760008560000151828151811061309657613096614db8565b602002602001015190506130b38187604001518860200151613a9b565b604081015151511515806130cf57506040810151516020015115155b1561312c576040518060400160405280826000015167ffffffffffffffff168152602001826020015162ffffff1681525083838151811061311257613112614db8565b602002602001018190525083806131289061519c565b9450505b508061313781614ec5565b915050613074565b5060008261ffff1667ffffffffffffffff81111561315f5761315f6148b3565b6040519080825280602002602001820160405280156131a457816020015b604080518082019091526000808252602082015281526020019060019003908161317d5790505b50905060005b8361ffff16811015613200578281815181106131c8576131c8614db8565b60200260200101518282815181106131e2576131e2614db8565b602002602001018190525080806131f890614ec5565b9150506131aa565b508764ffffffffff168a6040015163ffffffff167fe0c90b8e55243fcba0f8b68b201983b97f7a3d5aebd6dfa1a4082a07925cc7443388602001518d8660405161324d949392919061558e565b60405180910390a3505050506020015195945050505050565b600061328d633b9aca003a04866080015163ffffffff16876060015163ffffffff16613ead565b90506010360260005a905060006132b68663ffffffff1685858b60e0015162ffffff1686613eca565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff891683026001600160a01b03881660009081526012602052604090205460c08c01519290910492506201000090046001600160601b039081169163ffffffff16633b9aca00028284010190811682111561333c57505050505050506118dd565b6001600160a01b038816600090815260126020526040902080546001600160601b0390921662010000026dffffffffffffffffffffffff00001990921691909117905550505050505050505050565b60008060158054806020026020016040519081016040528092919081815260200182805480156133e457602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116133c6575b50508351600c54604080516103e08101918290529697509195660100000000000090910463ffffffff169450600093509150601690601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116134205790505050505050905060005b838110156134b3578181601f811061348057613480614db8565b602002015161348f9084614de4565b61349f9063ffffffff1687614e28565b9550806134ab81614ec5565b915050613466565b50600c546134d290600160901b900463ffffffff16633b9aca00614e09565b6134dc9086614e09565b945060005b8381101561355057601260008683815181106134ff576134ff614db8565b6020908102919091018101516001600160a01b031682528101919091526040016000205461353c906201000090046001600160601b031687614e28565b95508061354881614ec5565b9150506134e1565b505050505090565b60008183101561356957508161356c565b50805b92915050565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff1611156135cc57604051634a90778560e01b815261ffff861660048201526103e860248201526044016107af565b8461ffff1660000361360a576040517f08fad2a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006136367f00000000000000000000000000000000000000000000000000000000000000004361563d565b90506000816136657f000000000000000000000000000000000000000000000000000000000000000043614e28565b61366f9190614eae565b9050600061369d7f000000000000000000000000000000000000000000000000000000000000000083615651565b905063ffffffff81106136dc576040517f7b2a523000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526008805465ffffffffffff168252825161010081019384905284936000939291602084019160099084908288855b82829054906101000a900462ffffff1662ffffff168152602001906003019060208260020104928301926001038202915080841161371357905050505091909252505081519192505065ffffffffffff8082161061379d576040517f2b4655b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6137a8816001615665565b6008805465ffffffffffff191665ffffffffffff9290921691909117905560005b600881101561380f578a62ffffff16836020015182600881106137ee576137ee614db8565b602002015162ffffff161461380f578061380781614ec5565b9150506137c9565b600881106138505760208301516040517fc4f769b00000000000000000000000000000000000000000000000000000000081526107af918d91600401615686565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b806000106105265760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016107af565b6000808a8a8a8a8a8a8a8a8a604051602001613905999897969594939291906156a0565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6139788282613f2e565b5050565b336001600160a01b038216036139d45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107af565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b601c546001600160a01b03908116908216811461397857601c80546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b825167ffffffffffffffff80841691161115613afa5782516040517f012d824d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808516600483015290911660248201526044016107af565b60408301515151600090158015613b18575060408401515160200151155b15613b515750825167ffffffffffffffff1660009081526007602090815260408083208287015162ffffff168452909152902054613bac565b8360400151604051602001613b669190615728565b60408051601f198184030181529181528151602092830120865167ffffffffffffffff166000908152600784528281208885015162ffffff168252909352912081905590505b60608401515160008167ffffffffffffffff811115613bcd57613bcd6148b3565b604051908082528060200260200182016040528015613bf6578160200160208202803683370190505b50905060008267ffffffffffffffff811115613c1457613c146148b3565b6040519080825280601f01601f191660200182016040528015613c3e576020820181803683370190505b50905060008367ffffffffffffffff811115613c5c57613c5c6148b3565b604051908082528060200260200182016040528015613c8f57816020015b6060815260200190600190039081613c7a5790505b5090506000805b85811015613daa5760008a606001518281518110613cb657613cb6614db8565b60209081029190910101519050600080613cda8d600001518e602001518c8661400a565b915091508115613d195780868661ffff1681518110613cfb57613cfb614db8565b60200260200101819052508480613d119061519c565b955050613d60565b600160f81b878581518110613d3057613d30614db8565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505b8251518851899086908110613d7757613d77614db8565b602002602001019065ffffffffffff16908165ffffffffffff168152505050505080613da281614ec5565b915050613c96565b5060608901515115613ea25760008161ffff1667ffffffffffffffff811115613dd557613dd56148b3565b604051908082528060200260200182016040528015613e0857816020015b6060815260200190600190039081613df35790505b50905060005b8261ffff16811015613e6457838181518110613e2c57613e2c614db8565b6020026020010151828281518110613e4657613e46614db8565b60200260200101819052508080613e5c90614ec5565b915050613e0e565b507f47ddf7bb0cbd94c1b43c5097f1352a80db0ceb3696f029d32b24f32cd631d2b7858583604051613e989392919061575b565b60405180910390a1505b505050505050505050565b60008383811015613ec057600285850304015b610c198184613558565b600081861015613f1c5760405162461bcd60e51b815260206004820181905260248201527f6c6566744761732063616e6e6f742065786365656420696e697469616c47617360448201526064016107af565b50633b9aca0094039190910101020290565b610100818114613f70578282826040517fb93aa5de0000000000000000000000000000000000000000000000000000000081526004016107af93929190615801565b613f78614305565b8181604051602001613f8a91906146b1565b6040516020818303038152906040525114613fa757613fa7615825565b6040805180820190915260085465ffffffffffff16815260208101613fce8587018761583b565b905280516008805465ffffffffffff191665ffffffffffff9092169190911781556020820151614001906009908361443e565b506118dd915050565b60006060816140437f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff8916615651565b845160808101516040519293509091600091614067918b918b9186906020016150f8565b60408051601f198184030181529181528151602092830120845165ffffffffffff166000908152600690935291205490915081146140e35760016040518060400160405280601081526020017f756e6b6e6f776e2063616c6c6261636b00000000000000000000000000000000815250945094505050506142b0565b6040805160808101825263ffffffff8516815262ffffff8a1660208083019190915284015161ffff1681830152908301516001600160a01b03166060820152825160009061413390838b8e6128cb565b60608084015186519187015160405193945090926000927f5a47dd71000000000000000000000000000000000000000000000000000000009261417b928791906024016158c3565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252600b805466ff000000000000191666010000000000001790558b5160a0015191880151909250600091614219916001600160601b0390911690846142b9565b600b805466ff000000000000191690559050801561426b575050935165ffffffffffff1660009081526006602090815260408083208390558051918201905281815290975095506142b0945050505050565b60016040518060400160405280601081526020017f657865637574696f6e206661696c6564000000000000000000000000000000008152509950995050505050505050505b94509492505050565b60005a6113888110156142cb57600080fd5b6113888103905084604082048203116142e357600080fd5b50823b6142ef57600080fd5b60008083516020850160008789f1949350505050565b6040518061010001604052806008906020820280368337509192915050565b508054600082559060005260206000209081019061052691906144c5565b828054828255906000526020600020908101928215614397579160200282015b8281111561439757825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614362565b506143a39291506144c5565b5090565b6004830191839082156143975791602002820160005b8382111561440157835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026143bd565b80156144315782816101000a81549063ffffffff0219169055600401602081600301049283019260010302614401565b50506143a39291506144c5565b6001830191839082156143975791602002820160005b8382111561449657835183826101000a81548162ffffff021916908362ffffff1602179055509260200192600301602081600201049283019260010302614454565b80156144315782816101000a81549062ffffff0219169055600301602081600201049283019260010302614496565b5b808211156143a357600081556001016144c6565b6001600160a01b038116811461052657600080fd5b60006020828403121561450157600080fd5b8135611ac1816144da565b60005b8381101561452757818101518382015260200161450f565b838111156118dd5750506000910152565b6000815180845261455081602086016020860161450c565b601f01601f19169290920160200192915050565b602081526000611ac16020830184614538565b803563ffffffff8116811461458b57600080fd5b919050565b62ffffff8116811461052657600080fd5b600080600080600060a086880312156145b957600080fd5b6145c286614577565b94506145d060208701614577565b93506145de60408701614577565b92506145ec60608701614577565b915060808601356145fc81614590565b809150509295509295909350565b65ffffffffffff8116811461052657600080fd5b60006020828403121561463057600080fd5b8135611ac18161460a565b600081518084526020808501945080840160005b8381101561466b5781518752958201959082019060010161464f565b509495945050505050565b602081526000611ac1602083018461463b565b8060005b60088110156118dd57815162ffffff1684526020938401939091019060010161468d565b610100810161356c8284614689565b60008083601f8401126146d257600080fd5b50813567ffffffffffffffff8111156146ea57600080fd5b6020830191508360208260051b850101111561470557600080fd5b9250929050565b6000806000806040858703121561472257600080fd5b843567ffffffffffffffff8082111561473a57600080fd5b614746888389016146c0565b9096509450602087013591508082111561475f57600080fd5b5061476c878288016146c0565b95989497509550505050565b6000806020838503121561478b57600080fd5b823567ffffffffffffffff8111156147a257600080fd5b6147ae858286016146c0565b90969095509350505050565b60008083601f8401126147cc57600080fd5b50813567ffffffffffffffff8111156147e457600080fd5b60208301915083602082850101111561470557600080fd5b60008060008060008060008060e0898b03121561481857600080fd5b606089018a81111561482957600080fd5b8998503567ffffffffffffffff8082111561484357600080fd5b61484f8c838d016147ba565b909950975060808b013591508082111561486857600080fd5b6148748c838d016146c0565b909750955060a08b013591508082111561488d57600080fd5b5061489a8b828c016146c0565b999c989b50969995989497949560c00135949350505050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156148ec576148ec6148b3565b60405290565b60405160c0810167ffffffffffffffff811182821017156148ec576148ec6148b3565b6040516080810167ffffffffffffffff811182821017156148ec576148ec6148b3565b6040516020810167ffffffffffffffff811182821017156148ec576148ec6148b3565b604051601f8201601f1916810167ffffffffffffffff81118282101715614984576149846148b3565b604052919050565b600067ffffffffffffffff8211156149a6576149a66148b3565b50601f01601f191660200190565b600082601f8301126149c557600080fd5b81356149d86149d38261498c565b61495b565b8181528460208386010111156149ed57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115614a2457614a246148b3565b5060051b60200190565b60006020808385031215614a4157600080fd5b823567ffffffffffffffff80821115614a5957600080fd5b9084019060408287031215614a6d57600080fd5b614a756148c9565b823582811115614a8457600080fd5b614a90888286016149b4565b8252508383013582811115614aa457600080fd5b80840193505086601f840112614ab957600080fd5b82359150614ac96149d383614a0a565b82815260059290921b83018401918481019088841115614ae857600080fd5b938501935b83851015614b0657843582529385019390850190614aed565b948201949094529695505050505050565b60008060408385031215614b2a57600080fd5b8235614b35816144da565b946020939093013593505050565b600060208284031215614b5557600080fd5b813567ffffffffffffffff811115614b6c57600080fd5b820160808185031215611ac157600080fd5b61ffff8116811461052657600080fd5b67ffffffffffffffff8116811461052657600080fd5b803561458b81614b8e565b600080600060608486031215614bc457600080fd5b8335614bcf81614b7e565b92506020840135614bdf81614b8e565b91506040840135614bef81614590565b809150509250925092565b803560ff8116811461458b57600080fd5b60008060008060008060008060008060c08b8d031215614c2a57600080fd5b8a3567ffffffffffffffff80821115614c4257600080fd5b614c4e8e838f016146c0565b909c509a5060208d0135915080821115614c6757600080fd5b614c738e838f016146c0565b909a509850889150614c8760408e01614bfa565b975060608d0135915080821115614c9d57600080fd5b614ca98e838f016147ba565b9097509550859150614cbd60808e01614ba4565b945060a08d0135915080821115614cd357600080fd5b50614ce08d828e016147ba565b915080935050809150509295989b9194979a5092959850565b60008060408385031215614d0c57600080fd5b8235614d17816144da565b91506020830135614d27816144da565b809150509250929050565b600080600080600060a08688031215614d4a57600080fd5b8535614d5581614b8e565b94506020860135614d6581614b7e565b93506040860135614d7581614590565b9250614d8360608701614577565b9150608086013567ffffffffffffffff811115614d9f57600080fd5b614dab888289016149b4565b9150509295509295909350565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600063ffffffff83811690831681811015614e0157614e01614dce565b039392505050565b6000816000190483118215151615614e2357614e23614dce565b500290565b60008219821115614e3b57614e3b614dce565b500190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0384168152604060208201526000610c19604083018486614e40565b600060208284031215614e9e57600080fd5b81518015158114611ac157600080fd5b600082821015614ec057614ec0614dce565b500390565b600060018201614ed757614ed7614dce565b5060010190565b600060ff821660ff84168060ff03821115614efb57614efb614dce565b019392505050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60008251614f4181846020870161450c565b9190910192915050565b600060208284031215614f5d57600080fd5b5051919050565b6000808312837f800000000000000000000000000000000000000000000000000000000000000001831281151615614f9e57614f9e614dce565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018313811615614fd257614fd2614dce565b50500390565b600060ff821660ff84168160ff0481118215151615614ff957614ff9614dce565b029392505050565b600063ffffffff80831681851680830382111561502057615020614dce565b01949350505050565b600081518084526020808501945080840160005b8381101561466b5781516001600160a01b03168752958201959082019060010161503d565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526150928184018a615029565b905082810360808401526150a68189615029565b905060ff871660a084015282810360c08401526150c38187614538565b905067ffffffffffffffff851660e08401528281036101008401526150e88185614538565b9c9b505050505050505050505050565b600067ffffffffffffffff808716835262ffffff8616602084015280851660408401526080606084015265ffffffffffff845116608084015261ffff60208501511660a08401526001600160a01b0360408501511660c0840152606084015160c060e085015261516c610140850182614538565b60808601519092166101008501525060a0909301516001600160601b031661012090920191909152509392505050565b600061ffff8083168181036151b3576151b3614dce565b6001019392505050565b805161458b81614b8e565b600082601f8301126151d957600080fd5b81516151e76149d38261498c565b8181528460208386010111156151fc57600080fd5b61076e82602083016020870161450c565b80516001600160601b038116811461458b57600080fd5b600082601f83011261523557600080fd5b815160206152456149d383614a0a565b82815260059290921b8401810191818101908684111561526457600080fd5b8286015b84811015612ac757805167ffffffffffffffff8082111561528857600080fd5b90880190601f196040838c03820112156152a157600080fd5b6152a96148c9565b87840151838111156152ba57600080fd5b840160c0818e03840112156152ce57600080fd5b6152d66148f2565b9250888101516152e58161460a565b835260408101516152f581614b7e565b838a01526060810151615307816144da565b604084015260808101518481111561531e57600080fd5b61532c8e8b838501016151c8565b60608501525061533e60a082016151bd565b608084015261534f60c0820161520d565b60a0840152508181526153646040850161520d565b818901528652505050918301918301615268565b805177ffffffffffffffffffffffffffffffffffffffffffffffff8116811461458b57600080fd5b6000602082840312156153b257600080fd5b815167ffffffffffffffff808211156153ca57600080fd5b90830190608082860312156153de57600080fd5b6153e6614915565b8251828111156153f557600080fd5b8301601f8101871361540657600080fd5b80516154146149d382614a0a565b8082825260208201915060208360051b85010192508983111561543657600080fd5b602084015b838110156155365780518781111561545257600080fd5b850160a0818d03601f1901121561546857600080fd5b615470614915565b602082015161547e81614b8e565b8152604082015161548e81614590565b60208201526040828e03605f190112156154a757600080fd5b6154af614938565b8d607f8401126154be57600080fd5b6154c66148c9565b808f60a0860111156154d757600080fd5b606085015b60a086018110156154f75780518352602092830192016154dc565b50825250604082015260a08201518981111561551257600080fd5b6155218e602083860101615224565b6060830152508452506020928301920161543b565b5084525061554991505060208401615378565b602082015261555a604084016151bd565b60408201526060830151606082015280935050505092915050565b600063ffffffff8083168181036151b3576151b3614dce565b6000608082016001600160a01b0387168352602077ffffffffffffffffffffffffffffffffffffffffffffffff871681850152604086818601526080606086015282865180855260a087019150838801945060005b81811015615617578551805167ffffffffffffffff16845285015162ffffff168584015294840194918301916001016155e3565b50909a9950505050505050505050565b634e487b7160e01b600052601260045260246000fd5b60008261564c5761564c615627565b500690565b60008261566057615660615627565b500490565b600065ffffffffffff80831681851680830382111561502057615020614dce565b62ffffff831681526101208101611ac16020830184614689565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526156da8285018b615029565b915083820360808501526156ee828a615029565b915060ff881660a085015283820360c085015261570b8288614538565b90861660e085015283810361010085015290506150e88185614538565b815160408201908260005b6002811015615752578251825260209283019290910190600101615733565b50505092915050565b606080825284519082018190526000906020906080840190828801845b8281101561579c57815165ffffffffffff1684529284019290840190600101615778565b505050838103828501526157b08187614538565b905083810360408501528085518083528383019150838160051b84010184880160005b8381101561561757601f198684030185526157ef838351614538565b948701949250908601906001016157d3565b604081526000615815604083018587614e40565b9050826020830152949350505050565b634e487b7160e01b600052600160045260246000fd5b600061010080838503121561584f57600080fd5b83601f84011261585e57600080fd5b60405181810181811067ffffffffffffffff82111715615880576158806148b3565b60405290830190808583111561589557600080fd5b845b838110156158b85780356158aa81614590565b825260209182019101615897565b509095945050505050565b65ffffffffffff841681526060602082015260006158e4606083018561463b565b82810360408401526158f68185614538565b969550505050505056fea164736f6c634300080f000a",
}

var VRFBeaconCoordinatorABI = VRFBeaconCoordinatorMetaData.ABI

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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) LINK() (common.Address, error) {
	return _VRFBeaconCoordinator.Contract.LINK(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) LINK() (common.Address, error) {
	return _VRFBeaconCoordinator.Contract.LINK(&_VRFBeaconCoordinator.CallOpts)
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCaller) GetConfirmationDelays(opts *bind.CallOpts) ([8]*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconCoordinator.contract.Call(opts, &out, "getConfirmationDelays")

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) GetConfirmationDelays() ([8]*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.GetConfirmationDelays(&_VRFBeaconCoordinator.CallOpts)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorCallerSession) GetConfirmationDelays() ([8]*big.Int, error) {
	return _VRFBeaconCoordinator.Contract.GetConfirmationDelays(&_VRFBeaconCoordinator.CallOpts)
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) ForgetConsumerSubscriptionID(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "forgetConsumerSubscriptionID", consumers)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconCoordinator.TransactOpts, consumers)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconCoordinator.TransactOpts, consumers)
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

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactor) RedeemRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.contract.Transact(opts, "redeemRandomness", requestID)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.RedeemRandomness(&_VRFBeaconCoordinator.TransactOpts, requestID)
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorTransactorSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconCoordinator.Contract.RedeemRandomness(&_VRFBeaconCoordinator.TransactOpts, requestID)
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
	EpochAndRound     *big.Int
	Transmitter       common.Address
	JuelsPerFeeCoin   *big.Int
	ConfigDigest      [32]byte
	OutputsServed     []VRFBeaconReportOutputServed
	Raw               types.Log
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) FilterNewTransmission(opts *bind.FilterOpts, aggregatorRoundId []uint32, epochAndRound []*big.Int) (*VRFBeaconCoordinatorNewTransmissionIterator, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.FilterLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconCoordinatorNewTransmissionIterator{contract: _VRFBeaconCoordinator.contract, event: "NewTransmission", logs: logs, sub: sub}, nil
}

func (_VRFBeaconCoordinator *VRFBeaconCoordinatorFilterer) WatchNewTransmission(opts *bind.WatchOpts, sink chan<- *VRFBeaconCoordinatorNewTransmission, aggregatorRoundId []uint32, epochAndRound []*big.Int) (event.Subscription, error) {

	var aggregatorRoundIdRule []interface{}
	for _, aggregatorRoundIdItem := range aggregatorRoundId {
		aggregatorRoundIdRule = append(aggregatorRoundIdRule, aggregatorRoundIdItem)
	}
	var epochAndRoundRule []interface{}
	for _, epochAndRoundItem := range epochAndRound {
		epochAndRoundRule = append(epochAndRoundRule, epochAndRoundItem)
	}

	logs, sub, err := _VRFBeaconCoordinator.contract.WatchLogs(opts, "NewTransmission", aggregatorRoundIdRule, epochAndRoundRule)
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
	Bin: "0x608060405234801561001057600080fd5b5060405161045338038061045383398101604081905261002f91610058565b600080546001600160a01b0319166001600160a01b039390931692909217909155600155610092565b6000806040838503121561006b57600080fd5b82516001600160a01b038116811461008257600080fd5b6020939093015192949293505050565b6103b2806100a16000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806355e4874914610051578063bf2732c71461005b578063cc31f7dd1461006e578063d57fc45a14610089575b600080fd5b610059610092565b005b610059610069366004610287565b6100f6565b61007760015481565b60405190815260200160405180910390f35b61007760025481565b60005473ffffffffffffffffffffffffffffffffffffffff163381146100ee5760405163292f4fb560e01b815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044015b60405180910390fd5b506000600255565b60005473ffffffffffffffffffffffffffffffffffffffff1633811461014d5760405163292f4fb560e01b815233600482015273ffffffffffffffffffffffffffffffffffffffff821660248201526044016100e5565b815160405161015f919060200161036a565b60408051601f1981840301815291905280516020909101206002555050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156101d0576101d061017e565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156101ff576101ff61017e565b604052919050565b600082601f83011261021857600080fd5b8135602067ffffffffffffffff8211156102345761023461017e565b8160051b6102438282016101d6565b928352848101820192828101908785111561025d57600080fd5b83870192505b8483101561027c57823582529183019190830190610263565b979650505050505050565b6000602080838503121561029a57600080fd5b823567ffffffffffffffff808211156102b257600080fd5b90840190604082870312156102c657600080fd5b6102ce6101ad565b8235828111156102dd57600080fd5b8301601f810188136102ee57600080fd5b8035838111156103005761030061017e565b610312601f8201601f191687016101d6565b818152898783850101111561032657600080fd5b81878401888301376000878383010152808452505050838301358281111561034d57600080fd5b61035988828601610207565b948201949094529695505050505050565b6000825160005b8181101561038b5760208186018101518583015201610371565b8181111561039a576000828501525b50919091019291505056fea164736f6c634300080f000a",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"forgetConsumerSubscriptionID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200161a3803806200161a8339810160408190526200003491620001df565b8033806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000134565b5050506001600160a01b03166080526000829003620000f157604051632abc297960e01b815260040160405180910390fd5b60a082905260006200010483436200021e565b905060008160a05162000118919062000257565b905062000126814362000271565b60c052506200028c92505050565b336001600160a01b038216036200018e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008060408385031215620001f357600080fd5b825160208401519092506001600160a01b03811681146200021357600080fd5b809150509250929050565b6000826200023c57634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b6000828210156200026c576200026c62000241565b500390565b6000821982111562000287576200028762000241565b500190565b60805160a05160c051611343620002d760003960006101ee0152600081816101c70152818161036301528181610b0a01528181610b390152610b710152600060e901526113436000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c8063bbcdd0d81161008c578063cf7e754a11610066578063cf7e754a146101e9578063dc92accf14610210578063f2fde38b1461023a578063f645dcb11461024d57600080fd5b8063bbcdd0d814610190578063c63c4e9b146101a7578063cd0593df146101c257600080fd5b806379ba5097116100bd57806379ba5097146101625780638da5cb5b1461016c5780639e3616f41461017d57600080fd5b80631b6b6d23146100e45780632f7527cc1461012857806374d8461114610142575b600080fd5b61010b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b610130600881565b60405160ff909116815260200161011f565b610155610150366004610e40565b610260565b60405161011f9190610e6f565b61016a610469565b005b6000546001600160a01b031661010b565b61016a61018b366004610eb3565b610527565b6101996103e881565b60405190815260200161011f565b6101af600381565b60405161ffff909116815260200161011f565b6101997f000000000000000000000000000000000000000000000000000000000000000081565b6101997f000000000000000000000000000000000000000000000000000000000000000081565b61022361021e366004610f6a565b6105b7565b60405165ffffffffffff909116815260200161011f565b61016a610248366004610fad565b6106ee565b61022361025b366004610fec565b610702565b65ffffffffffff81166000818152600a602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff670100000000000000820416938201939093526001600160a01b03690100000000000000000084048116606083810191825298909752949093527fffffff0000000000000000000000000000000000000000000000000000000000909116905591511633146103585760608101516040517f8e30e8230000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201523360248201526044015b60405180910390fd5b805160009061038e907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16611101565b90506000826020015162ffffff16436103a79190611120565b90508082106103eb576040517f15ad27c30000000000000000000000000000000000000000000000000000000081526004810183905243602482015260440161034f565b67ffffffffffffffff821115610430576040517f058ddf020000000000000000000000000000000000000000000000000000000081526004810183905260240161034f565b60008281526007602090815260408083208287015162ffffff168452909152902054610460908690859085610808565b95945050505050565b6001546001600160a01b031633146104c35760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161034f565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61052f610a0f565b60005b818110156105b25760006005600085858581811061055257610552611137565b90506020020160208101906105679190610fad565b6001600160a01b031681526020810191909152604001600020805467ffffffffffffffff191667ffffffffffffffff92909216919091179055806105aa8161114d565b915050610532565b505050565b6000806000806105c78786610a6b565b92509250925065ffffffffffff83166000908152600a602090815260409182902084518154928601518487015160608801516001600160a01b03166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff90921667010000000000000002919091167fffffff00000000000000000000000000000000000000000000ffffffffffffff62ffffff9093166401000000000266ffffffffffffff1990961663ffffffff909416939093179490941716179190911790555167ffffffffffffffff8216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc5906106db90889062ffffff91909116815260200190565b60405180910390a2509095945050505050565b6106f6610a0f565b6106ff81610d8a565b50565b60008060006107118787610a6b565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a67ffffffffffffffff1681526020018763ffffffff166bffffffffffffffffffffffff16815250905081878a836040516020016107919493929190611166565b60408051601f19818403018152828252805160209182012065ffffffffffff871660009081526006909252919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906107f39084908a908d908690611166565b60405180910390a15090979650505050505050565b60608261085b576040517fc7d41b1b00000000000000000000000000000000000000000000000000000000815265ffffffffffff8616600482015267ffffffffffffffff8316602482015260440161034f565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff161115610911576040808601519051634a90778560e01b815261ffff90911660048201526103e8602482015260440161034f565b6000856040015161ffff1667ffffffffffffffff81111561093457610934610fd6565b60405190808252806020026020018201604052801561095d578160200160208202803683370190505b50905060005b866040015161ffff168161ffff161015610a045782816040516020016109b892919091825260f01b7fffff00000000000000000000000000000000000000000000000000000000000016602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff16815181106109e7576109e7611137565b6020908102919091010152806109fc81611252565b915050610963565b509695505050505050565b6000546001600160a01b03163314610a695760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161034f565b565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff161115610ac557604051634a90778560e01b815261ffff861660048201526103e8602482015260440161034f565b8461ffff16600003610b03576040517f08fad2a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610b2f7f000000000000000000000000000000000000000000000000000000000000000043611289565b9050600081610b5e7f00000000000000000000000000000000000000000000000000000000000000004361129d565b610b689190611120565b90506000610b967f0000000000000000000000000000000000000000000000000000000000000000836112b5565b905063ffffffff8110610bd5576040517f7b2a523000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526008805465ffffffffffff168252825161010081019384905284936000939291602084019160099084908288855b82829054906101000a900462ffffff1662ffffff1681526020019060030190602082600201049283019260010382029150808411610c0c57905050505091909252505081519192505065ffffffffffff80821610610c96576040517f2b4655b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ca18160016112c9565b6008805465ffffffffffff191665ffffffffffff9290921691909117905560005b6008811015610d08578a62ffffff1683602001518260088110610ce757610ce7611137565b602002015162ffffff1614610d085780610d008161114d565b915050610cc2565b60088110610d495760208301516040517fc4f769b000000000000000000000000000000000000000000000000000000000815261034f918d916004016112f3565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b336001600160a01b03821603610de25760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161034f565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215610e5257600080fd5b813565ffffffffffff81168114610e6857600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610ea757835183529284019291840191600101610e8b565b50909695505050505050565b60008060208385031215610ec657600080fd5b823567ffffffffffffffff80821115610ede57600080fd5b818501915085601f830112610ef257600080fd5b813581811115610f0157600080fd5b8660208260051b8501011115610f1657600080fd5b60209290920196919550909350505050565b803561ffff81168114610f3a57600080fd5b919050565b803567ffffffffffffffff81168114610f3a57600080fd5b803562ffffff81168114610f3a57600080fd5b600080600060608486031215610f7f57600080fd5b610f8884610f28565b9250610f9660208501610f3f565b9150610fa460408501610f57565b90509250925092565b600060208284031215610fbf57600080fd5b81356001600160a01b0381168114610e6857600080fd5b634e487b7160e01b600052604160045260246000fd5b600080600080600060a0868803121561100457600080fd5b61100d86610f3f565b945061101b60208701610f28565b935061102960408701610f57565b9250606086013563ffffffff8116811461104257600080fd5b9150608086013567ffffffffffffffff8082111561105f57600080fd5b818801915088601f83011261107357600080fd5b81358181111561108557611085610fd6565b604051601f8201601f19908116603f011681019083821181831017156110ad576110ad610fd6565b816040528281528b60208487010111156110c657600080fd5b8260208601602083013760006020848301015280955050505050509295509295909350565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561111b5761111b6110eb565b500290565b600082821015611132576111326110eb565b500390565b634e487b7160e01b600052603260045260246000fd5b60006001820161115f5761115f6110eb565b5060010190565b600067ffffffffffffffff8087168352602062ffffff87168185015281861660408501526080606085015265ffffffffffff855116608085015261ffff818601511660a08501526001600160a01b0360408601511660c08501526060850151915060c060e085015281518061014086015260005b818110156111f757838101830151868201610160015282016111da565b8181111561120a57600061016083880101525b50608086015167ffffffffffffffff1661010086015260a0909501516bffffffffffffffffffffffff16610120850152505050601f909101601f191601610160019392505050565b600061ffff808316818103611269576112696110eb565b6001019392505050565b634e487b7160e01b600052601260045260246000fd5b60008261129857611298611273565b500690565b600082198211156112b0576112b06110eb565b500190565b6000826112c4576112c4611273565b500490565b600065ffffffffffff8083168185168083038211156112ea576112ea6110eb565b01949350505050565b62ffffff838116825261012082019060208084018560005b600881101561132a57815185168352918301919083019060010161130b565b5050505050939250505056fea164736f6c634300080f000a",
}

var VRFBeaconExternalAPIABI = VRFBeaconExternalAPIMetaData.ABI

var VRFBeaconExternalAPIBin = VRFBeaconExternalAPIMetaData.Bin

func DeployVRFBeaconExternalAPI(auth *bind.TransactOpts, backend bind.ContractBackend, beaconPeriodBlocksArg *big.Int, linkToken common.Address) (common.Address, *types.Transaction, *VRFBeaconExternalAPI, error) {
	parsed, err := VRFBeaconExternalAPIMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconExternalAPIBin), backend, beaconPeriodBlocksArg, linkToken)
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

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconExternalAPI.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) LINK() (common.Address, error) {
	return _VRFBeaconExternalAPI.Contract.LINK(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerSession) LINK() (common.Address, error) {
	return _VRFBeaconExternalAPI.Contract.LINK(&_VRFBeaconExternalAPI.CallOpts)
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

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconExternalAPI.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) Owner() (common.Address, error) {
	return _VRFBeaconExternalAPI.Contract.Owner(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPICallerSession) Owner() (common.Address, error) {
	return _VRFBeaconExternalAPI.Contract.Owner(&_VRFBeaconExternalAPI.CallOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.contract.Transact(opts, "acceptOwnership")
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.AcceptOwnership(&_VRFBeaconExternalAPI.TransactOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.AcceptOwnership(&_VRFBeaconExternalAPI.TransactOpts)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactor) ForgetConsumerSubscriptionID(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.contract.Transact(opts, "forgetConsumerSubscriptionID", consumers)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconExternalAPI.TransactOpts, consumers)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconExternalAPI.TransactOpts, consumers)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactor) RedeemRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.contract.Transact(opts, "redeemRandomness", requestID)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.RedeemRandomness(&_VRFBeaconExternalAPI.TransactOpts, requestID)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.RedeemRandomness(&_VRFBeaconExternalAPI.TransactOpts, requestID)
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

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPISession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.TransferOwnership(&_VRFBeaconExternalAPI.TransactOpts, to)
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPITransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconExternalAPI.Contract.TransferOwnership(&_VRFBeaconExternalAPI.TransactOpts, to)
}

type VRFBeaconExternalAPIOwnershipTransferRequestedIterator struct {
	Event *VRFBeaconExternalAPIOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconExternalAPIOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconExternalAPIOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconExternalAPIOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFBeaconExternalAPIOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconExternalAPIOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconExternalAPIOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconExternalAPIOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconExternalAPI.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPIOwnershipTransferRequestedIterator{contract: _VRFBeaconExternalAPI.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFBeaconExternalAPIOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconExternalAPI.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconExternalAPIOwnershipTransferRequested)
				if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFBeaconExternalAPIOwnershipTransferRequested, error) {
	event := new(VRFBeaconExternalAPIOwnershipTransferRequested)
	if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFBeaconExternalAPIOwnershipTransferredIterator struct {
	Event *VRFBeaconExternalAPIOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFBeaconExternalAPIOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFBeaconExternalAPIOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(VRFBeaconExternalAPIOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *VRFBeaconExternalAPIOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFBeaconExternalAPIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFBeaconExternalAPIOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFBeaconExternalAPIOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconExternalAPI.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFBeaconExternalAPIOwnershipTransferredIterator{contract: _VRFBeaconExternalAPI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFBeaconExternalAPIOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFBeaconExternalAPI.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFBeaconExternalAPIOwnershipTransferred)
				if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_VRFBeaconExternalAPI *VRFBeaconExternalAPIFilterer) ParseOwnershipTransferred(log types.Log) (*VRFBeaconExternalAPIOwnershipTransferred, error) {
	event := new(VRFBeaconExternalAPIOwnershipTransferred)
	if err := _VRFBeaconExternalAPI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"OffchainConfigHasWrongLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"occVersion\",\"type\":\"uint64\"}],\"name\":\"UnknownConfigVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"old\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAccessControllerInterface\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"BillingAccessControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"BillingSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconReport.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractLinkTokenInterface\",\"name\":\"linkToken\",\"type\":\"address\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"PayeeshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"acceptPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconReport.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"forgetConsumerSubscriptionID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBilling\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBillingAccessController\",\"outputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfirmationDelays\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"\",\"type\":\"uint24[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"availableBalance\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxErrorMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"oracleObservationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitterAddress\",\"type\":\"address\"}],\"name\":\"owedPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maximumGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"reasonableGasPriceGwei\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"observationPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionPaymentGjuels\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"accountingGas\",\"type\":\"uint24\"}],\"name\":\"setBilling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAccessControllerInterface\",\"name\":\"_billingAccessController\",\"type\":\"address\"}],\"name\":\"setBillingAccessController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"}],\"name\":\"setPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposed\",\"type\":\"address\"}],\"name\":\"transferPayeeship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620059ba380380620059ba8339810160408190526200003491620001e7565b81818181803380600081620000905760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c357620000c3816200013c565b5050506001600160a01b03166080526000829003620000f557604051632abc297960e01b815260040160405180910390fd5b60a0829052600062000108834362000226565b905060008160a0516200011c91906200025f565b90506200012a814362000279565b60c05250620002949650505050505050565b336001600160a01b03821603620001965760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000087565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008060408385031215620001fb57600080fd5b825160208401519092506001600160a01b03811681146200021b57600080fd5b809150509250929050565b6000826200024457634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60008282101562000274576200027462000249565b500390565b600082198211156200028f576200028f62000249565b500190565b60805160a05160c0516156a96200031160003960006104fa0152600081816104d301528181610a3e015281816134c5015281816134f40152818161352c0152613ec801526000818161027a015281816115c901528181611690015281816117b9015281816123e101528181612a320152612b7901526156a96000f3fe608060405234801561001057600080fd5b506004361061020b5760003560e01c8063b121e1471161012a578063cf7e754a116100bd578063e4902f821161008c578063f2fde38b11610071578063f2fde38b1461059c578063f645dcb1146105af578063fbffd2c1146105c257600080fd5b8063e4902f8214610561578063eb5dcd6c1461058957600080fd5b8063cf7e754a146104f5578063d09dc3391461051c578063dc92accf14610524578063e3d0e7121461054e57600080fd5b8063c278e5b7116100f9578063c278e5b714610491578063c4c92b37146104a2578063c63c4e9b146104b3578063cd0593df146104ce57600080fd5b8063b121e1471461044f578063b1dc65a414610462578063bbcdd0d814610475578063c10753291461047e57600080fd5b80637a464944116101a25780638da5cb5b116101715780638da5cb5b146103ee5780639c849b30146103ff5780639e3616f414610412578063afcb95d71461042557600080fd5b80637a4649441461039157806381ff70481461039957806385c64e11146103c65780638ac28d5a146103db57600080fd5b80632f7527cc116101de5780632f7527cc1461033a578063643dc1051461035457806374d846111461036957806379ba50971461038957600080fd5b80630eafb25b14610210578063181f5a77146102365780631b6b6d231461027557806329937268146102b4575b600080fd5b61022361021e3660046143a3565b6105d5565b6040519081526020015b60405180910390f35b604080518082018252601581527f565246426561636f6e20312e302e302d616c70686100000000000000000000006020820152905161022d9190614418565b61029c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161022d565b6102fe600c546a0100000000000000000000810463ffffffff90811692600160701b8304821692600160901b8104831692600160b01b82041691600160d01b90910462ffffff1690565b6040805163ffffffff9687168152948616602086015292851692840192909252909216606082015262ffffff909116608082015260a00161022d565b610342600881565b60405160ff909116815260200161022d565b610367610362366004614455565b6106dd565b005b61037c6103773660046144d2565b610940565b60405161022d919061452a565b610367610b44565b610223608081565b600d54600f54604080516000815264010000000090930463ffffffff16602084015282015260600161022d565b6103ce610bf5565b60405161022d9190614565565b6103676103e93660046143a3565b610c5a565b6000546001600160a01b031661029c565b61036761040d3660046145c0565b610ccc565b61036761042036600461462c565b610eaa565b600f546011546040805160008152602081019390935263ffffffff9091169082015260600161022d565b61036761045d3660046143a3565b610f3a565b6103676104703660046146b0565b611016565b6102236103e881565b61036761048c366004614767565b6114c6565b61036761049f366004614793565b50565b601c546001600160a01b031661029c565b6104bb600381565b60405161ffff909116815260200161022d565b6102237f000000000000000000000000000000000000000000000000000000000000000081565b6102237f000000000000000000000000000000000000000000000000000000000000000081565b610223611797565b6105376105323660046147ff565b611843565b60405165ffffffffffff909116815260200161022d565b61036761055c36600461485b565b61197c565b61057461056f3660046143a3565b6120c2565b60405163ffffffff909116815260200161022d565b610367610597366004614949565b612174565b6103676105aa3660046143a3565b6122ac565b6105376105bd366004614a83565b6122bd565b6103676105d03660046143a3565b6123be565b6001600160a01b03811660009081526012602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906106375750600092915050565b600c546020820151600091600160901b900463ffffffff169060169060ff16601f811061066657610666614b4c565b600881049190910154600c5461069c926007166004026101000a90910463ffffffff908116916601000000000000900416614b78565b63ffffffff166106ac9190614b9d565b6106ba90633b9aca00614b9d565b905081604001516001600160601b0316816106d59190614bbc565b949350505050565b601c546001600160a01b03166106fb6000546001600160a01b031690565b6001600160a01b0316336001600160a01b031614806107875750604051630d629b5f60e31b81526001600160a01b03821690636b14daf8906107469033906000903690600401614bfd565b602060405180830381865afa158015610763573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107879190614c20565b6107d85760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c60448201526064015b60405180910390fd5b6107e06123cf565b600c80547fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff166a010000000000000000000063ffffffff8981169182027fffffffffffffffffffffffffffff00000000ffffffffffffffffffffffffffff1692909217600160701b898416908102919091177fffffffffffff0000000000000000ffffffffffffffffffffffffffffffffffff16600160901b8985169081027fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff1691909117600160b01b948916948502177fffffff000000ffffffffffffffffffffffffffffffffffffffffffffffffffff16600160d01b62ffffff89169081029190911790955560408051938452602084019290925290820152606081019190915260808101919091527f0bf184bf1bba9699114bdceddaf338a1b364252c5e497cc01918dde92031713f9060a00160405180910390a1505050505050565b65ffffffffffff81166000818152600a602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff670100000000000000820416938201939093526001600160a01b03690100000000000000000084048116606083810191825298909752949093527fffffff000000000000000000000000000000000000000000000000000000000090911690559151163314610a335760608101516040517f8e30e8230000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201523360248201526044016107cf565b8051600090610a69907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16614b9d565b90506000826020015162ffffff1643610a829190614c42565b9050808210610ac6576040517f15ad27c3000000000000000000000000000000000000000000000000000000008152600481018390524360248201526044016107cf565b67ffffffffffffffff821115610b0b576040517f058ddf02000000000000000000000000000000000000000000000000000000008152600481018390526024016107cf565b60008281526007602090815260408083208287015162ffffff168452909152902054610b3b90869085908561277f565b95945050505050565b6001546001600160a01b03163314610b9e5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016107cf565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610bfd6141b9565b6040805161010081019182905290600990600890826000855b82829054906101000a900462ffffff1662ffffff1681526020019060030190602082600201049283019260010382029150808411610c165790505050505050905090565b6001600160a01b038181166000908152601a6020526040902054163314610cc35760405162461bcd60e51b815260206004820152601760248201527f4f6e6c792070617965652063616e20776974686472617700000000000000000060448201526064016107cf565b61049f81612986565b610cd4612bd0565b828114610d235760405162461bcd60e51b815260206004820181905260248201527f7472616e736d6974746572732e73697a6520213d207061796565732e73697a6560448201526064016107cf565b60005b83811015610ea3576000858583818110610d4257610d42614b4c565b9050602002016020810190610d5791906143a3565b90506000848484818110610d6d57610d6d614b4c565b9050602002016020810190610d8291906143a3565b6001600160a01b038084166000908152601a60205260409020549192501680158080610dbf5750826001600160a01b0316826001600160a01b0316145b610e0b5760405162461bcd60e51b815260206004820152601160248201527f706179656520616c72656164792073657400000000000000000000000000000060448201526064016107cf565b6001600160a01b038481166000908152601a6020526040902080546001600160a01b03191685831690811790915590831614610e8c57826001600160a01b0316826001600160a01b0316856001600160a01b03167f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b360405160405180910390a45b505050508080610e9b90614c59565b915050610d26565b5050505050565b610eb2612bd0565b60005b81811015610f3557600060056000858585818110610ed557610ed5614b4c565b9050602002016020810190610eea91906143a3565b6001600160a01b031681526020810191909152604001600020805467ffffffffffffffff191667ffffffffffffffff9290921691909117905580610f2d81614c59565b915050610eb5565b505050565b6001600160a01b038181166000908152601b6020526040902054163314610fa35760405162461bcd60e51b815260206004820152601f60248201527f6f6e6c792070726f706f736564207061796565732063616e206163636570740060448201526064016107cf565b6001600160a01b038181166000818152601a602090815260408083208054336001600160a01b03198083168217909355601b909452828520805490921690915590519416939092849290917f78af32efdcad432315431e9b03d27e6cd98fb79c405fdc5af7c1714d9c0f75b39190a45050565b60005a604080516101008082018352600c5460ff808216845291810464ffffffffff166020808501919091526601000000000000820463ffffffff908116858701526a0100000000000000000000830481166060860152600160701b830481166080860152600160901b8304811660a0860152600160b01b83041660c0850152600160d01b90910462ffffff1660e08401523360009081526012825293909320549394509092918c0135911661110e5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016107cf565b600f548b35146111605760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016107cf565b61116e8a8a8a8a8a8a612c2c565b815161117b906001614c72565b60ff1687146111cc5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016107cf565b86851461121b5760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016107cf565b60008a8a60405161122d929190614c97565b604051908190038120611244918e90602001614ca7565b60408051601f19818403018152828252805160209182012083830190925260008084529083018190529092509060005b8a8110156113ea5760006001858a846020811061129357611293614b4c565b6112a091901a601b614c72565b8f8f868181106112b2576112b2614b4c565b905060200201358e8e878181106112cb576112cb614b4c565b9050602002013560405160008152602001604052604051611308949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561132a573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526013602090815290849020838501909452925460ff80821615158085526101009092041693830193909352909550925090506113c35760405162461bcd60e51b815260206004820152600f60248201527f7369676e6174757265206572726f72000000000000000000000000000000000060448201526064016107cf565b826020015160080260ff166001901b840193505080806113e290614c59565b915050611274565b5081827e01010101010101010101010101010101010101010101010101010101010101161461145b5760405162461bcd60e51b815260206004820152601060248201527f6475706c6963617465207369676e65720000000000000000000000000000000060448201526064016107cf565b50600091506114aa9050838d836020020135848e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612cc992505050565b90506114b88382863361311a565b505050505050505050505050565b6000546001600160a01b03163314806115505750601c54604051630d629b5f60e31b81526001600160a01b0390911690636b14daf89061150f9033906000903690600401614bfd565b602060405180830381865afa15801561152c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115509190614c20565b61159c5760405162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722662696c6c696e6741646d696e2063616e2063616c6c60448201526064016107cf565b60006115a661323f565b6040516370a0823160e01b81523060048201529091506000906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611610573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116349190614cc3565b9050818110156116865760405162461bcd60e51b815260206004820152601460248201527f696e73756666696369656e742062616c616e636500000000000000000000000060448201526064016107cf565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb856116c96116c38686614c42565b8761340c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561172c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117509190614c20565b6117915760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b60448201526064016107cf565b50505050565b6040516370a0823160e01b815230600482015260009081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611800573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118249190614cc3565b9050600061183061323f565b905061183c8183614cdc565b9250505090565b6000806000806118538786613426565b92509250925065ffffffffffff83166000908152600a602090815260409182902084518154928601518487015160608801516001600160a01b03166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff90921667010000000000000002919091167fffffff00000000000000000000000000000000000000000000ffffffffffffff62ffffff9093166401000000000266ffffffffffffff1990961663ffffffff909416939093179490941716179190911790555167ffffffffffffffff8216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc59061196790889062ffffff91909116815260200190565b60405180910390a250909150505b9392505050565b611984612bd0565b601f8911156119d55760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79206f7261636c65730000000000000000000000000000000060448201526064016107cf565b888714611a245760405162461bcd60e51b815260206004820152601660248201527f6f7261636c65206c656e677468206d69736d617463680000000000000000000060448201526064016107cf565b88611a30876003614d50565b60ff1610611a805760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016107cf565b611a8c8660ff16613745565b6040805160e060208c02808301820190935260c082018c815260009383928f918f918291908601908490808284376000920191909152505050908252506040805160208c810282810182019093528c82529283019290918d918d91829185019084908082843760009201919091525050509082525060ff891660208083019190915260408051601f8a01839004830281018301825289815292019190899089908190840183828082843760009201919091525050509082525067ffffffffffffffff861660208083019190915260408051601f870183900483028101830182528681529201919086908690819084018382808284376000920191909152505050915250600c805465ffffffffff00191690559050611ba86123cf565b60145460005b81811015611c5957600060148281548110611bcb57611bcb614b4c565b6000918252602082200154601580546001600160a01b0390921693509084908110611bf857611bf8614b4c565b60009182526020808320909101546001600160a01b039485168352601382526040808420805461ffff1916905594168252601290529190912080546dffffffffffffffffffffffffffff191690555080611c5181614c59565b915050611bae565b50611c66601460006141d8565b611c72601560006141d8565b60005b825151811015611ef0576013600084600001518381518110611c9957611c99614b4c565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611d0d5760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016107cf565b604080518082019091526001815260ff821660208201528351805160139160009185908110611d3e57611d3e614b4c565b6020908102919091018101516001600160a01b03168252818101929092526040016000908120835181549484015161ffff1990951690151561ff0019161761010060ff90951694909402939093179092558401518051601292919084908110611da957611da9614b4c565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff1615611e1d5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016107cf565b60405180606001604052806001151581526020018260ff16815260200160006001600160601b03168152506012600085602001518481518110611e6257611e62614b4c565b6020908102919091018101516001600160a01b03168252818101929092526040908101600020835181549385015194909201516001600160601b031662010000026dffffffffffffffffffffffff00001960ff959095166101000261ff00199315159390931661ffff1990941693909317919091179290921617905580611ee881614c59565b915050611c75565b5081518051611f07916014916020909101906141f6565b506020808301518051611f1e9260159201906141f6565b506040820151600c805460ff191660ff909216919091179055600d805467ffffffff0000000019811664010000000063ffffffff43811682029283179094558204831692600092611f76929082169116176001614d79565b905080600d60006101000a81548163ffffffff021916908363ffffffff1602179055506000611fca46308463ffffffff16886000015189602001518a604001518b606001518c608001518d60a00151613795565b905080600f600001819055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05838284886000015189602001518a604001518b606001518c608001518d60a0015160405161202d99989796959493929190614dda565b60405180910390a1600c546601000000000000900463ffffffff1660005b8651518110156120a55781601682601f811061206957612069614b4c565b600891828204019190066004026101000a81548163ffffffff021916908363ffffffff160217905550808061209d90614c59565b91505061204b565b506120b08b8b613822565b50505050505050505050505050505050565b6001600160a01b03811660009081526012602090815260408083208151606081018352905460ff80821615158084526101008304909116948301949094526201000090046001600160601b031691810191909152906121245750600092915050565b6016816020015160ff16601f811061213e5761213e614b4c565b600881049190910154600c54611975926007166004026101000a90910463ffffffff908116916601000000000000900416614b78565b6001600160a01b038281166000908152601a60205260409020541633146121dd5760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c792063757272656e742070617965652063616e2075706461746500000060448201526064016107cf565b6001600160a01b03811633036122355760405162461bcd60e51b815260206004820152601760248201527f63616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107cf565b6001600160a01b038083166000908152601b6020526040902080548383166001600160a01b031982168117909255909116908114610f35576040516001600160a01b038084169133918616907f84f7c7c80bb8ed2279b4aab5f61cd05e6374073d38f46d7f32de8c30e9e3836790600090a4505050565b6122b4612bd0565b61049f81613830565b60008060006122cc8787613426565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a67ffffffffffffffff1681526020018763ffffffff166001600160601b0316815250905081878a836040516020016123479493929190614e70565b60408051601f19818403018152828252805160209182012065ffffffffffff871660009081526006909252919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c906123a99084908a908d908690614e70565b60405180910390a15090979650505050505050565b6123c6612bd0565b61049f816138d9565b600c54604080516103e08101918290527f0000000000000000000000000000000000000000000000000000000000000000926601000000000000900463ffffffff169160009190601690601f908285855b82829054906101000a900463ffffffff1663ffffffff168152602001906004019060208260030104928301926001038202915080841161242057905050505050509050600060158054806020026020016040519081016040528092919081815260200182805480156124bb57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161249d575b5050505050905060005b8151811015612771576000601260008484815181106124e6576124e6614b4c565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160029054906101000a90046001600160601b03166001600160601b0316905060006012600085858151811061254857612548614b4c565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160026101000a8154816001600160601b0302191690836001600160601b0316021790555060008483601f81106125ab576125ab614b4c565b6020020151600c5490870363ffffffff9081169250600160901b909104168102633b9aca000282018015612766576000601a60008787815181106125f1576125f1614b4c565b6020908102919091018101516001600160a01b03908116835290820192909252604090810160002054905163a9059cbb60e01b815290821660048201819052602482018590529250908a169063a9059cbb906044016020604051808303816000875af1158015612665573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126899190614c20565b6126ca5760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b60448201526064016107cf565b878786601f81106126dd576126dd614b4c565b602002019063ffffffff16908163ffffffff1681525050886001600160a01b0316816001600160a01b031687878151811061271a5761271a614b4c565b60200260200101516001600160a01b03167fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c8560405161275c91815260200190565b60405180910390a4505b5050506001016124c5565b50610ea3601683601f61425b565b6060826127d2576040517fc7d41b1b00000000000000000000000000000000000000000000000000000000815265ffffffffffff8616600482015267ffffffffffffffff831660248201526044016107cf565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff161115612888576040808601519051634a90778560e01b815261ffff90911660048201526103e860248201526044016107cf565b6000856040015161ffff1667ffffffffffffffff8111156128ab576128ab614982565b6040519080825280602002602001820160405280156128d4578160200160208202803683370190505b50905060005b866040015161ffff168161ffff16101561297b57828160405160200161292f92919091825260f01b7fffff00000000000000000000000000000000000000000000000000000000000016602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff168151811061295e5761295e614b4c565b60209081029190910101528061297381614f14565b9150506128da565b509695505050505050565b6001600160a01b0381166000908152601260209081526040918290208251606081018452905460ff80821615158084526101008304909116938301939093526201000090046001600160601b0316928101929092526129e3575050565b60006129ee836105d5565b90508015610f35576001600160a01b038381166000908152601a60205260409081902054905163a9059cbb60e01b81529082166004820181905260248201849052917f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af1158015612a7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a9f9190614c20565b612ae05760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b60448201526064016107cf565b600c60000160069054906101000a900463ffffffff166016846020015160ff16601f8110612b1057612b10614b4c565b6008810491909101805460079092166004026101000a63ffffffff8181021990931693909216919091029190911790556001600160a01b0384811660008181526012602090815260409182902080546dffffffffffffffffffffffff00001916905590518581527f0000000000000000000000000000000000000000000000000000000000000000841693851692917fd0b1dac935d85bd54cf0a33b0d41d39f8cf53a968465fc7ea2377526b8ac712c910160405180910390a450505050565b6000546001600160a01b03163314612c2a5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107cf565b565b6000612c39826020614b9d565b612c44856020614b9d565b612c5088610144614bbc565b612c5a9190614bbc565b612c649190614bbc565b612c6f906000614bbc565b9050368114612cc05760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d61746368000000000000000060448201526064016107cf565b50505050505050565b60008082806020019051810190612ce0919061513c565b64ffffffffff85166020880152604087018051919250612cff82615311565b63ffffffff1663ffffffff168152505085600c60008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160000160066101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001600a6101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600e6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160000160126101000a81548163ffffffff021916908363ffffffff16021790555060c08201518160000160166101000a81548163ffffffff021916908363ffffffff16021790555060e082015181600001601a6101000a81548162ffffff021916908362ffffff1602179055509050506000816040015167ffffffffffffffff1640905080826060015114612ebc57606082015160408084015190517faed0afe500000000000000000000000000000000000000000000000000000000815260048101929092526024820183905267ffffffffffffffff1660448201526064016107cf565b60008083600001515167ffffffffffffffff811115612edd57612edd614982565b604051908082528060200260200182016040528015612f2257816020015b6040805180820190915260008082526020820152815260200190600190039081612efb5790505b50905060005b845151811015612ff357600085600001518281518110612f4a57612f4a614b4c565b60200260200101519050612f67818760400151886020015161394f565b60408101515151151580612f8357506040810151516020015115155b15612fe0576040518060400160405280826000015167ffffffffffffffff168152602001826020015162ffffff16815250838381518110612fc657612fc6614b4c565b60200260200101819052508380612fdc90614f14565b9450505b5080612feb81614c59565b915050612f28565b5060008261ffff1667ffffffffffffffff81111561301357613013614982565b60405190808252806020026020018201604052801561305857816020015b60408051808201909152600080825260208201528152602001906001900390816130315790505b50905060005b8361ffff168110156130b45782818151811061307c5761307c614b4c565b602002602001015182828151811061309657613096614b4c565b602002602001018190525080806130ac90614c59565b91505061305e565b508764ffffffffff168a6040015163ffffffff167fe0c90b8e55243fcba0f8b68b201983b97f7a3d5aebd6dfa1a4082a07925cc7443388602001518d86604051613101949392919061532a565b60405180910390a3505050506020015195945050505050565b6000613141633b9aca003a04866080015163ffffffff16876060015163ffffffff16613d61565b90506010360260005a9050600061316a8663ffffffff1685858b60e0015162ffffff1686613d7e565b90506000670de0b6b3a764000077ffffffffffffffffffffffffffffffffffffffffffffffff891683026001600160a01b03881660009081526012602052604090205460c08c01519290910492506201000090046001600160601b039081169163ffffffff16633b9aca0002828401019081168211156131f05750505050505050611791565b6001600160a01b038816600090815260126020526040902080546001600160601b0390921662010000026dffffffffffffffffffffffff00001990921691909117905550505050505050505050565b600080601580548060200260200160405190810160405280929190818152602001828054801561329857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161327a575b50508351600c54604080516103e08101918290529697509195660100000000000090910463ffffffff169450600093509150601690601f908285855b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116132d45790505050505050905060005b83811015613367578181601f811061333457613334614b4c565b60200201516133439084614b78565b6133539063ffffffff1687614bbc565b95508061335f81614c59565b91505061331a565b50600c5461338690600160901b900463ffffffff16633b9aca00614b9d565b6133909086614b9d565b945060005b8381101561340457601260008683815181106133b3576133b3614b4c565b6020908102919091018101516001600160a01b03168252810191909152604001600020546133f0906201000090046001600160601b031687614bbc565b9550806133fc81614c59565b915050613395565b505050505090565b60008183101561341d575081613420565b50805b92915050565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff16111561348057604051634a90778560e01b815261ffff861660048201526103e860248201526044016107cf565b8461ffff166000036134be576040517f08fad2a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006134ea7f0000000000000000000000000000000000000000000000000000000000000000436153d9565b90506000816135197f000000000000000000000000000000000000000000000000000000000000000043614bbc565b6135239190614c42565b905060006135517f0000000000000000000000000000000000000000000000000000000000000000836153ed565b905063ffffffff8110613590576040517f7b2a523000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526008805465ffffffffffff168252825161010081019384905284936000939291602084019160099084908288855b82829054906101000a900462ffffff1662ffffff16815260200190600301906020826002010492830192600103820291508084116135c757905050505091909252505081519192505065ffffffffffff80821610613651576040517f2b4655b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61365c816001615401565b6008805465ffffffffffff191665ffffffffffff9290921691909117905560005b60088110156136c3578a62ffffff16836020015182600881106136a2576136a2614b4c565b602002015162ffffff16146136c357806136bb81614c59565b91505061367d565b600881106137045760208301516040517fc4f769b00000000000000000000000000000000000000000000000000000000081526107cf918d91600401615422565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b8060001061049f5760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016107cf565b6000808a8a8a8a8a8a8a8a8a6040516020016137b99998979695949392919061543c565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b61382c8282613de2565b5050565b336001600160a01b038216036138885760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107cf565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b601c546001600160a01b03908116908216811461382c57601c80546001600160a01b0319166001600160a01b0384811691821790925560408051928416835260208301919091527f793cb73064f3c8cde7e187ae515511e6e56d1ee89bf08b82fa60fb70f8d48912910160405180910390a15050565b825167ffffffffffffffff808416911611156139ae5782516040517f012d824d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808516600483015290911660248201526044016107cf565b604083015151516000901580156139cc575060408401515160200151155b15613a055750825167ffffffffffffffff1660009081526007602090815260408083208287015162ffffff168452909152902054613a60565b8360400151604051602001613a1a91906154c4565b60408051601f198184030181529181528151602092830120865167ffffffffffffffff166000908152600784528281208885015162ffffff168252909352912081905590505b60608401515160008167ffffffffffffffff811115613a8157613a81614982565b604051908082528060200260200182016040528015613aaa578160200160208202803683370190505b50905060008267ffffffffffffffff811115613ac857613ac8614982565b6040519080825280601f01601f191660200182016040528015613af2576020820181803683370190505b50905060008367ffffffffffffffff811115613b1057613b10614982565b604051908082528060200260200182016040528015613b4357816020015b6060815260200190600190039081613b2e5790505b5090506000805b85811015613c5e5760008a606001518281518110613b6a57613b6a614b4c565b60209081029190910101519050600080613b8e8d600001518e602001518c86613ebe565b915091508115613bcd5780868661ffff1681518110613baf57613baf614b4c565b60200260200101819052508480613bc590614f14565b955050613c14565b600160f81b878581518110613be457613be4614b4c565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505b8251518851899086908110613c2b57613c2b614b4c565b602002602001019065ffffffffffff16908165ffffffffffff168152505050505080613c5681614c59565b915050613b4a565b5060608901515115613d565760008161ffff1667ffffffffffffffff811115613c8957613c89614982565b604051908082528060200260200182016040528015613cbc57816020015b6060815260200190600190039081613ca75790505b50905060005b8261ffff16811015613d1857838181518110613ce057613ce0614b4c565b6020026020010151828281518110613cfa57613cfa614b4c565b60200260200101819052508080613d1090614c59565b915050613cc2565b507f47ddf7bb0cbd94c1b43c5097f1352a80db0ceb3696f029d32b24f32cd631d2b7858583604051613d4c939291906154f7565b60405180910390a1505b505050505050505050565b60008383811015613d7457600285850304015b610b3b818461340c565b600081861015613dd05760405162461bcd60e51b815260206004820181905260248201527f6c6566744761732063616e6e6f742065786365656420696e697469616c47617360448201526064016107cf565b50633b9aca0094039190910101020290565b610100818114613e24578282826040517fb93aa5de0000000000000000000000000000000000000000000000000000000081526004016107cf9392919061559d565b613e2c6141b9565b8181604051602001613e3e9190614565565b6040516020818303038152906040525114613e5b57613e5b6155c1565b6040805180820190915260085465ffffffffffff16815260208101613e82858701876155d7565b905280516008805465ffffffffffff191665ffffffffffff9092169190911781556020820151613eb590600990836142f2565b50611791915050565b6000606081613ef77f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff89166153ed565b845160808101516040519293509091600091613f1b918b918b918690602001614e70565b60408051601f198184030181529181528151602092830120845165ffffffffffff16600090815260069093529120549091508114613f975760016040518060400160405280601081526020017f756e6b6e6f776e2063616c6c6261636b0000000000000000000000000000000081525094509450505050614164565b6040805160808101825263ffffffff8516815262ffffff8a1660208083019190915284015161ffff1681830152908301516001600160a01b031660608201528251600090613fe790838b8e61277f565b60608084015186519187015160405193945090926000927f5a47dd71000000000000000000000000000000000000000000000000000000009261402f9287919060240161565f565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252600b805466ff000000000000191666010000000000001790558b5160a00151918801519092506000916140cd916001600160601b03909116908461416d565b600b805466ff000000000000191690559050801561411f575050935165ffffffffffff166000908152600660209081526040808320839055805191820190528181529097509550614164945050505050565b60016040518060400160405280601081526020017f657865637574696f6e206661696c6564000000000000000000000000000000008152509950995050505050505050505b94509492505050565b60005a61138881101561417f57600080fd5b61138881039050846040820482031161419757600080fd5b50823b6141a357600080fd5b60008083516020850160008789f1949350505050565b6040518061010001604052806008906020820280368337509192915050565b508054600082559060005260206000209081019061049f9190614379565b82805482825590600052602060002090810192821561424b579160200282015b8281111561424b57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614216565b50614257929150614379565b5090565b60048301918390821561424b5791602002820160005b838211156142b557835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302614271565b80156142e55782816101000a81549063ffffffff02191690556004016020816003010492830192600103026142b5565b5050614257929150614379565b60018301918390821561424b5791602002820160005b8382111561434a57835183826101000a81548162ffffff021916908362ffffff1602179055509260200192600301602081600201049283019260010302614308565b80156142e55782816101000a81549062ffffff021916905560030160208160020104928301926001030261434a565b5b80821115614257576000815560010161437a565b6001600160a01b038116811461049f57600080fd5b6000602082840312156143b557600080fd5b81356119758161438e565b60005b838110156143db5781810151838201526020016143c3565b838111156117915750506000910152565b600081518084526144048160208601602086016143c0565b601f01601f19169290920160200192915050565b60208152600061197560208301846143ec565b803563ffffffff8116811461443f57600080fd5b919050565b62ffffff8116811461049f57600080fd5b600080600080600060a0868803121561446d57600080fd5b6144768661442b565b94506144846020870161442b565b93506144926040870161442b565b92506144a06060870161442b565b915060808601356144b081614444565b809150509295509295909350565b65ffffffffffff8116811461049f57600080fd5b6000602082840312156144e457600080fd5b8135611975816144be565b600081518084526020808501945080840160005b8381101561451f57815187529582019590820190600101614503565b509495945050505050565b60208152600061197560208301846144ef565b8060005b600881101561179157815162ffffff16845260209384019390910190600101614541565b6101008101613420828461453d565b60008083601f84011261458657600080fd5b50813567ffffffffffffffff81111561459e57600080fd5b6020830191508360208260051b85010111156145b957600080fd5b9250929050565b600080600080604085870312156145d657600080fd5b843567ffffffffffffffff808211156145ee57600080fd5b6145fa88838901614574565b9096509450602087013591508082111561461357600080fd5b5061462087828801614574565b95989497509550505050565b6000806020838503121561463f57600080fd5b823567ffffffffffffffff81111561465657600080fd5b61466285828601614574565b90969095509350505050565b60008083601f84011261468057600080fd5b50813567ffffffffffffffff81111561469857600080fd5b6020830191508360208285010111156145b957600080fd5b60008060008060008060008060e0898b0312156146cc57600080fd5b606089018a8111156146dd57600080fd5b8998503567ffffffffffffffff808211156146f757600080fd5b6147038c838d0161466e565b909950975060808b013591508082111561471c57600080fd5b6147288c838d01614574565b909750955060a08b013591508082111561474157600080fd5b5061474e8b828c01614574565b999c989b50969995989497949560c00135949350505050565b6000806040838503121561477a57600080fd5b82356147858161438e565b946020939093013593505050565b6000602082840312156147a557600080fd5b813567ffffffffffffffff8111156147bc57600080fd5b82016080818503121561197557600080fd5b61ffff8116811461049f57600080fd5b67ffffffffffffffff8116811461049f57600080fd5b803561443f816147de565b60008060006060848603121561481457600080fd5b833561481f816147ce565b9250602084013561482f816147de565b9150604084013561483f81614444565b809150509250925092565b803560ff8116811461443f57600080fd5b60008060008060008060008060008060c08b8d03121561487a57600080fd5b8a3567ffffffffffffffff8082111561489257600080fd5b61489e8e838f01614574565b909c509a5060208d01359150808211156148b757600080fd5b6148c38e838f01614574565b909a5098508891506148d760408e0161484a565b975060608d01359150808211156148ed57600080fd5b6148f98e838f0161466e565b909750955085915061490d60808e016147f4565b945060a08d013591508082111561492357600080fd5b506149308d828e0161466e565b915080935050809150509295989b9194979a5092959850565b6000806040838503121561495c57600080fd5b82356149678161438e565b915060208301356149778161438e565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156149bb576149bb614982565b60405290565b60405160c0810167ffffffffffffffff811182821017156149bb576149bb614982565b6040516080810167ffffffffffffffff811182821017156149bb576149bb614982565b6040516020810167ffffffffffffffff811182821017156149bb576149bb614982565b604051601f8201601f1916810167ffffffffffffffff81118282101715614a5357614a53614982565b604052919050565b600067ffffffffffffffff821115614a7557614a75614982565b50601f01601f191660200190565b600080600080600060a08688031215614a9b57600080fd5b8535614aa6816147de565b94506020860135614ab6816147ce565b93506040860135614ac681614444565b9250614ad46060870161442b565b9150608086013567ffffffffffffffff811115614af057600080fd5b8601601f81018813614b0157600080fd5b8035614b14614b0f82614a5b565b614a2a565b818152896020838501011115614b2957600080fd5b816020840160208301376000602083830101528093505050509295509295909350565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600063ffffffff83811690831681811015614b9557614b95614b62565b039392505050565b6000816000190483118215151615614bb757614bb7614b62565b500290565b60008219821115614bcf57614bcf614b62565b500190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0384168152604060208201526000610b3b604083018486614bd4565b600060208284031215614c3257600080fd5b8151801515811461197557600080fd5b600082821015614c5457614c54614b62565b500390565b600060018201614c6b57614c6b614b62565b5060010190565b600060ff821660ff84168060ff03821115614c8f57614c8f614b62565b019392505050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b600060208284031215614cd557600080fd5b5051919050565b6000808312837f800000000000000000000000000000000000000000000000000000000000000001831281151615614d1657614d16614b62565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018313811615614d4a57614d4a614b62565b50500390565b600060ff821660ff84168160ff0481118215151615614d7157614d71614b62565b029392505050565b600063ffffffff808316818516808303821115614d9857614d98614b62565b01949350505050565b600081518084526020808501945080840160005b8381101561451f5781516001600160a01b031687529582019590820190600101614db5565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614e0a8184018a614da1565b90508281036080840152614e1e8189614da1565b905060ff871660a084015282810360c0840152614e3b81876143ec565b905067ffffffffffffffff851660e0840152828103610100840152614e6081856143ec565b9c9b505050505050505050505050565b600067ffffffffffffffff808716835262ffffff8616602084015280851660408401526080606084015265ffffffffffff845116608084015261ffff60208501511660a08401526001600160a01b0360408501511660c0840152606084015160c060e0850152614ee46101408501826143ec565b60808601519092166101008501525060a0909301516001600160601b031661012090920191909152509392505050565b600061ffff808316818103614f2b57614f2b614b62565b6001019392505050565b600067ffffffffffffffff821115614f4f57614f4f614982565b5060051b60200190565b805161443f816147de565b600082601f830112614f7557600080fd5b8151614f83614b0f82614a5b565b818152846020838601011115614f9857600080fd5b6106d58260208301602087016143c0565b80516001600160601b038116811461443f57600080fd5b600082601f830112614fd157600080fd5b81516020614fe1614b0f83614f35565b82815260059290921b8401810191818101908684111561500057600080fd5b8286015b8481101561297b57805167ffffffffffffffff8082111561502457600080fd5b90880190601f196040838c038201121561503d57600080fd5b615045614998565b878401518381111561505657600080fd5b840160c0818e038401121561506a57600080fd5b6150726149c1565b925088810151615081816144be565b83526040810151615091816147ce565b838a015260608101516150a38161438e565b60408401526080810151848111156150ba57600080fd5b6150c88e8b83850101614f64565b6060850152506150da60a08201614f59565b60808401526150eb60c08201614fa9565b60a08401525081815261510060408501614fa9565b818901528652505050918301918301615004565b805177ffffffffffffffffffffffffffffffffffffffffffffffff8116811461443f57600080fd5b60006020828403121561514e57600080fd5b815167ffffffffffffffff8082111561516657600080fd5b908301906080828603121561517a57600080fd5b6151826149e4565b82518281111561519157600080fd5b8301601f810187136151a257600080fd5b80516151b0614b0f82614f35565b8082825260208201915060208360051b8501019250898311156151d257600080fd5b602084015b838110156152d2578051878111156151ee57600080fd5b850160a0818d03601f1901121561520457600080fd5b61520c6149e4565b602082015161521a816147de565b8152604082015161522a81614444565b60208201526040828e03605f1901121561524357600080fd5b61524b614a07565b8d607f84011261525a57600080fd5b615262614998565b808f60a08601111561527357600080fd5b606085015b60a08601811015615293578051835260209283019201615278565b50825250604082015260a0820151898111156152ae57600080fd5b6152bd8e602083860101614fc0565b606083015250845250602092830192016151d7565b508452506152e591505060208401615114565b60208201526152f660408401614f59565b60408201526060830151606082015280935050505092915050565b600063ffffffff808316818103614f2b57614f2b614b62565b6000608082016001600160a01b0387168352602077ffffffffffffffffffffffffffffffffffffffffffffffff871681850152604086818601526080606086015282865180855260a087019150838801945060005b818110156153b3578551805167ffffffffffffffff16845285015162ffffff1685840152948401949183019160010161537f565b50909a9950505050505050505050565b634e487b7160e01b600052601260045260246000fd5b6000826153e8576153e86153c3565b500690565b6000826153fc576153fc6153c3565b500490565b600065ffffffffffff808316818516808303821115614d9857614d98614b62565b62ffffff831681526101208101611975602083018461453d565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526154768285018b614da1565b9150838203608085015261548a828a614da1565b915060ff881660a085015283820360c08501526154a782886143ec565b90861660e08501528381036101008501529050614e6081856143ec565b815160408201908260005b60028110156154ee5782518252602092830192909101906001016154cf565b50505092915050565b606080825284519082018190526000906020906080840190828801845b8281101561553857815165ffffffffffff1684529284019290840190600101615514565b5050508381038285015261554c81876143ec565b905083810360408501528085518083528383019150838160051b84010184880160005b838110156153b357601f1986840301855261558b8383516143ec565b9487019492509086019060010161556f565b6040815260006155b1604083018587614bd4565b9050826020830152949350505050565b634e487b7160e01b600052600160045260246000fd5b60006101008083850312156155eb57600080fd5b83601f8401126155fa57600080fd5b60405181810181811067ffffffffffffffff8211171561561c5761561c614982565b60405290830190808583111561563157600080fd5b845b8381101561565457803561564681614444565b825260209182019101615633565b509095945050505050565b65ffffffffffff8416815260606020820152600061568060608301856144ef565b828103604084015261569281856143ec565b969550505050505056fea164736f6c634300080f000a",
}

var VRFBeaconOCRABI = VRFBeaconOCRMetaData.ABI

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

func (_VRFBeaconOCR *VRFBeaconOCRCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) LINK() (common.Address, error) {
	return _VRFBeaconOCR.Contract.LINK(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) LINK() (common.Address, error) {
	return _VRFBeaconOCR.Contract.LINK(&_VRFBeaconOCR.CallOpts)
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

func (_VRFBeaconOCR *VRFBeaconOCRCaller) GetConfirmationDelays(opts *bind.CallOpts) ([8]*big.Int, error) {
	var out []interface{}
	err := _VRFBeaconOCR.contract.Call(opts, &out, "getConfirmationDelays")

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

func (_VRFBeaconOCR *VRFBeaconOCRSession) GetConfirmationDelays() ([8]*big.Int, error) {
	return _VRFBeaconOCR.Contract.GetConfirmationDelays(&_VRFBeaconOCR.CallOpts)
}

func (_VRFBeaconOCR *VRFBeaconOCRCallerSession) GetConfirmationDelays() ([8]*big.Int, error) {
	return _VRFBeaconOCR.Contract.GetConfirmationDelays(&_VRFBeaconOCR.CallOpts)
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

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) ForgetConsumerSubscriptionID(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "forgetConsumerSubscriptionID", consumers)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconOCR.TransactOpts, consumers)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconOCR.TransactOpts, consumers)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactor) RedeemRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.contract.Transact(opts, "redeemRandomness", requestID)
}

func (_VRFBeaconOCR *VRFBeaconOCRSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.RedeemRandomness(&_VRFBeaconOCR.TransactOpts, requestID)
}

func (_VRFBeaconOCR *VRFBeaconOCRTransactorSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconOCR.Contract.RedeemRandomness(&_VRFBeaconOCR.TransactOpts, requestID)
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
	EpochAndRound     *big.Int
	Transmitter       common.Address
	JuelsPerFeeCoin   *big.Int
	ConfigDigest      [32]byte
	OutputsServed     []VRFBeaconReportOutputServed
	Raw               types.Log
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beaconPeriodBlocksArg\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BeaconPeriodMustBePositive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earliestAllowed\",\"type\":\"uint256\"}],\"name\":\"BlockTooRecent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"firstDelay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minDelay\",\"type\":\"uint16\"}],\"name\":\"ConfirmationDelayBlocksTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confirmationDelays\",\"type\":\"uint16[10]\"},{\"internalType\":\"uint8\",\"name\":\"violatingIndex\",\"type\":\"uint8\"}],\"name\":\"ConfirmationDelaysNotIncreasing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"separatorHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"providedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"onchainHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"HistoryDomainSeparatorWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWordsRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16[10]\",\"name\":\"confDelays\",\"type\":\"uint16[10]\"}],\"name\":\"NonZeroDelayAfterZeroDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"requestHeight\",\"type\":\"uint256\"}],\"name\":\"RandomnessNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ResponseMustBeRetrievedByRequester\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequestsReplaceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySlotsReplaceContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"TooManyWords\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"name\":\"UniverseHasEndedBangBangBang\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"givenDelay\",\"type\":\"uint24\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay[8]\",\"name\":\"knownDelays\",\"type\":\"uint24[8]\"}],\"name\":\"UnknownConfirmationDelay\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"aggregatorRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconReport.OutputServed[]\",\"name\":\"outputsServed\",\"type\":\"tuple[]\"}],\"name\":\"NewTransmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.RequestID[]\",\"name\":\"requestIDs\",\"type\":\"uint48[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"successfulFulfillment\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"truncatedErrorData\",\"type\":\"bytes[]\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"}],\"name\":\"RandomnessFulfillmentRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nextBeaconOutputHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confDelay\",\"type\":\"uint24\"}],\"name\":\"RandomnessRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NUM_CONF_DELAYS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelay\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"p\",\"type\":\"uint256[2]\"}],\"internalType\":\"structECCArithmetic.G1Point\",\"name\":\"vrfOutput\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"gasAllowance\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.Callback\",\"name\":\"callback\",\"type\":\"tuple\"},{\"internalType\":\"uint96\",\"name\":\"price\",\"type\":\"uint96\"}],\"internalType\":\"structVRFBeaconTypes.CostedCallback[]\",\"name\":\"callbacks\",\"type\":\"tuple[]\"}],\"internalType\":\"structVRFBeaconReport.VRFOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint192\",\"name\":\"juelsPerFeeCoin\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"recentBlockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"recentBlockHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVRFBeaconReport.Report\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"exposeType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"name\":\"forgetConsumerSubscriptionID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_StartSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"i_beaconPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxErrorMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDelay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"requestID\",\"type\":\"uint48\"}],\"name\":\"redeemRandomness\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"randomness\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"}],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subID\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numWords\",\"type\":\"uint16\"},{\"internalType\":\"VRFBeaconTypes.ConfirmationDelay\",\"name\":\"confirmationDelayArg\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"arguments\",\"type\":\"bytes\"}],\"name\":\"requestRandomnessFulfillment\",\"outputs\":[{\"internalType\":\"VRFBeaconTypes.RequestID\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001685380380620016858339810160408190526200003491620001e3565b81818033806000816200008e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c157620000c18162000138565b5050506001600160a01b03166080526000829003620000f357604051632abc297960e01b815260040160405180910390fd5b60a0829052600062000106834362000222565b905060008160a0516200011a91906200025b565b905062000128814362000275565b60c0525062000290945050505050565b336001600160a01b03821603620001925760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000085565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008060408385031215620001f757600080fd5b825160208401519092506001600160a01b03811681146200021757600080fd5b809150509250929050565b6000826200024057634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b60008282101562000270576200027062000245565b500390565b600082198211156200028b576200028b62000245565b500190565b60805160a05160c0516113aa620002db600039600061021d0152600081816101f60152818161039201528181610b3601528181610b650152610b9d0152600060ff01526113aa6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063bbcdd0d811610097578063cf7e754a11610066578063cf7e754a14610218578063dc92accf1461023f578063f2fde38b14610269578063f645dcb11461027c57600080fd5b8063bbcdd0d8146101bc578063c278e5b7146101c5578063c63c4e9b146101d6578063cd0593df146101f157600080fd5b806379ba5097116100d357806379ba5097146101785780637a464944146101825780638da5cb5b146101985780639e3616f4146101a957600080fd5b80631b6b6d23146100fa5780632f7527cc1461013e57806374d8461114610158575b600080fd5b6101217f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b610146600881565b60405160ff9091168152602001610135565b61016b610166366004610e6c565b61028f565b6040516101359190610e9b565b610180610498565b005b61018a608081565b604051908152602001610135565b6000546001600160a01b0316610121565b6101806101b7366004610edf565b610556565b61018a6103e881565b6101806101d3366004610f54565b50565b6101de600381565b60405161ffff9091168152602001610135565b61018a7f000000000000000000000000000000000000000000000000000000000000000081565b61018a7f000000000000000000000000000000000000000000000000000000000000000081565b61025261024d366004610fd1565b6105e6565b60405165ffffffffffff9091168152602001610135565b610180610277366004611014565b61071d565b61025261028a366004611053565b61072e565b65ffffffffffff81166000818152600a602081815260408084208151608081018352815463ffffffff8116825262ffffff6401000000008204168286015261ffff670100000000000000820416938201939093526001600160a01b03690100000000000000000084048116606083810191825298909752949093527fffffff0000000000000000000000000000000000000000000000000000000000909116905591511633146103875760608101516040517f8e30e8230000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201523360248201526044015b60405180910390fd5b80516000906103bd907f00000000000000000000000000000000000000000000000000000000000000009063ffffffff16611168565b90506000826020015162ffffff16436103d69190611187565b905080821061041a576040517f15ad27c30000000000000000000000000000000000000000000000000000000081526004810183905243602482015260440161037e565b67ffffffffffffffff82111561045f576040517f058ddf020000000000000000000000000000000000000000000000000000000081526004810183905260240161037e565b60008281526007602090815260408083208287015162ffffff16845290915290205461048f908690859085610834565b95945050505050565b6001546001600160a01b031633146104f25760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161037e565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61055e610a3b565b60005b818110156105e1576000600560008585858181106105815761058161119e565b90506020020160208101906105969190611014565b6001600160a01b031681526020810191909152604001600020805467ffffffffffffffff191667ffffffffffffffff92909216919091179055806105d9816111b4565b915050610561565b505050565b6000806000806105f68786610a97565b92509250925065ffffffffffff83166000908152600a602090815260409182902084518154928601518487015160608801516001600160a01b03166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff61ffff90921667010000000000000002919091167fffffff00000000000000000000000000000000000000000000ffffffffffffff62ffffff9093166401000000000266ffffffffffffff1990961663ffffffff909416939093179490941716179190911790555167ffffffffffffffff8216907fc334d6f57be304c8192da2e39220c48e35f7e9afa16c541e68a6a859eff4dbc59061070a90889062ffffff91909116815260200190565b60405180910390a2509095945050505050565b610725610a3b565b6101d381610db6565b600080600061073d8787610a97565b925050915060006040518060c001604052808465ffffffffffff1681526020018961ffff168152602001336001600160a01b031681526020018681526020018a67ffffffffffffffff1681526020018763ffffffff166bffffffffffffffffffffffff16815250905081878a836040516020016107bd94939291906111cd565b60408051601f19818403018152828252805160209182012065ffffffffffff871660009081526006909252919020557fa62e84e206cb87e2f6896795353c5358ff3d415d0bccc24e45c5fad83e17d03c9061081f9084908a908d9086906111cd565b60405180910390a15090979650505050505050565b606082610887576040517fc7d41b1b00000000000000000000000000000000000000000000000000000000815265ffffffffffff8616600482015267ffffffffffffffff8316602482015260440161037e565b6040805165ffffffffffff8716602080830191909152865163ffffffff168284015286015162ffffff166060808301919091529186015161ffff166080820152908501516001600160a01b031660a082015260c0810184905260009060e0016040516020818303038152906040528051906020012090506103e8856040015161ffff16111561093d576040808601519051634a90778560e01b815261ffff90911660048201526103e8602482015260440161037e565b6000856040015161ffff1667ffffffffffffffff8111156109605761096061103d565b604051908082528060200260200182016040528015610989578160200160208202803683370190505b50905060005b866040015161ffff168161ffff161015610a305782816040516020016109e492919091825260f01b7fffff00000000000000000000000000000000000000000000000000000000000016602082015260220190565b6040516020818303038152906040528051906020012060001c828261ffff1681518110610a1357610a1361119e565b602090810291909101015280610a28816112b9565b91505061098f565b509695505050505050565b6000546001600160a01b03163314610a955760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161037e565b565b604080516080810182526000808252602082018190529181018290526060810182905260006103e88561ffff161115610af157604051634a90778560e01b815261ffff861660048201526103e8602482015260440161037e565b8461ffff16600003610b2f576040517f08fad2a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610b5b7f0000000000000000000000000000000000000000000000000000000000000000436112f0565b9050600081610b8a7f000000000000000000000000000000000000000000000000000000000000000043611304565b610b949190611187565b90506000610bc27f00000000000000000000000000000000000000000000000000000000000000008361131c565b905063ffffffff8110610c01576040517f7b2a523000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526008805465ffffffffffff168252825161010081019384905284936000939291602084019160099084908288855b82829054906101000a900462ffffff1662ffffff1681526020019060030190602082600201049283019260010382029150808411610c3857905050505091909252505081519192505065ffffffffffff80821610610cc2576040517f2b4655b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ccd816001611330565b6008805465ffffffffffff191665ffffffffffff9290921691909117905560005b6008811015610d34578a62ffffff1683602001518260088110610d1357610d1361119e565b602002015162ffffff1614610d345780610d2c816111b4565b915050610cee565b60088110610d755760208301516040517fc4f769b000000000000000000000000000000000000000000000000000000000815261037e918d9160040161135a565b506040805160808101825263ffffffff909416845262ffffff8b16602085015261ffff8c169084015233606084015297509095509193505050509250925092565b336001600160a01b03821603610e0e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161037e565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215610e7e57600080fd5b813565ffffffffffff81168114610e9457600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610ed357835183529284019291840191600101610eb7565b50909695505050505050565b60008060208385031215610ef257600080fd5b823567ffffffffffffffff80821115610f0a57600080fd5b818501915085601f830112610f1e57600080fd5b813581811115610f2d57600080fd5b8660208260051b8501011115610f4257600080fd5b60209290920196919550909350505050565b600060208284031215610f6657600080fd5b813567ffffffffffffffff811115610f7d57600080fd5b820160808185031215610e9457600080fd5b803561ffff81168114610fa157600080fd5b919050565b803567ffffffffffffffff81168114610fa157600080fd5b803562ffffff81168114610fa157600080fd5b600080600060608486031215610fe657600080fd5b610fef84610f8f565b9250610ffd60208501610fa6565b915061100b60408501610fbe565b90509250925092565b60006020828403121561102657600080fd5b81356001600160a01b0381168114610e9457600080fd5b634e487b7160e01b600052604160045260246000fd5b600080600080600060a0868803121561106b57600080fd5b61107486610fa6565b945061108260208701610f8f565b935061109060408701610fbe565b9250606086013563ffffffff811681146110a957600080fd5b9150608086013567ffffffffffffffff808211156110c657600080fd5b818801915088601f8301126110da57600080fd5b8135818111156110ec576110ec61103d565b604051601f8201601f19908116603f011681019083821181831017156111145761111461103d565b816040528281528b602084870101111561112d57600080fd5b8260208601602083013760006020848301015280955050505050509295509295909350565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561118257611182611152565b500290565b60008282101561119957611199611152565b500390565b634e487b7160e01b600052603260045260246000fd5b6000600182016111c6576111c6611152565b5060010190565b600067ffffffffffffffff8087168352602062ffffff87168185015281861660408501526080606085015265ffffffffffff855116608085015261ffff818601511660a08501526001600160a01b0360408601511660c08501526060850151915060c060e085015281518061014086015260005b8181101561125e5783810183015186820161016001528201611241565b8181111561127157600061016083880101525b50608086015167ffffffffffffffff1661010086015260a0909501516bffffffffffffffffffffffff16610120850152505050601f909101601f191601610160019392505050565b600061ffff8083168181036112d0576112d0611152565b6001019392505050565b634e487b7160e01b600052601260045260246000fd5b6000826112ff576112ff6112da565b500690565b6000821982111561131757611317611152565b500190565b60008261132b5761132b6112da565b500490565b600065ffffffffffff80831681851680830382111561135157611351611152565b01949350505050565b62ffffff838116825261012082019060208084018560005b6008811015611391578151851683529183019190830190600101611372565b5050505050939250505056fea164736f6c634300080f000a",
}

var VRFBeaconReportABI = VRFBeaconReportMetaData.ABI

var VRFBeaconReportBin = VRFBeaconReportMetaData.Bin

func DeployVRFBeaconReport(auth *bind.TransactOpts, backend bind.ContractBackend, beaconPeriodBlocksArg *big.Int, link common.Address) (common.Address, *types.Transaction, *VRFBeaconReport, error) {
	parsed, err := VRFBeaconReportMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFBeaconReportBin), backend, beaconPeriodBlocksArg, link)
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

func (_VRFBeaconReport *VRFBeaconReportCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFBeaconReport.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFBeaconReport *VRFBeaconReportSession) LINK() (common.Address, error) {
	return _VRFBeaconReport.Contract.LINK(&_VRFBeaconReport.CallOpts)
}

func (_VRFBeaconReport *VRFBeaconReportCallerSession) LINK() (common.Address, error) {
	return _VRFBeaconReport.Contract.LINK(&_VRFBeaconReport.CallOpts)
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

func (_VRFBeaconReport *VRFBeaconReportTransactor) ForgetConsumerSubscriptionID(opts *bind.TransactOpts, consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "forgetConsumerSubscriptionID", consumers)
}

func (_VRFBeaconReport *VRFBeaconReportSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconReport.TransactOpts, consumers)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) ForgetConsumerSubscriptionID(consumers []common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.ForgetConsumerSubscriptionID(&_VRFBeaconReport.TransactOpts, consumers)
}

func (_VRFBeaconReport *VRFBeaconReportTransactor) RedeemRandomness(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "redeemRandomness", requestID)
}

func (_VRFBeaconReport *VRFBeaconReportSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.RedeemRandomness(&_VRFBeaconReport.TransactOpts, requestID)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) RedeemRandomness(requestID *big.Int) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.RedeemRandomness(&_VRFBeaconReport.TransactOpts, requestID)
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

func (_VRFBeaconReport *VRFBeaconReportTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFBeaconReport *VRFBeaconReportSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.TransferOwnership(&_VRFBeaconReport.TransactOpts, to)
}

func (_VRFBeaconReport *VRFBeaconReportTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFBeaconReport.Contract.TransferOwnership(&_VRFBeaconReport.TransactOpts, to)
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
	EpochAndRound     *big.Int
	Transmitter       common.Address
	JuelsPerFeeCoin   *big.Int
	ConfigDigest      [32]byte
	OutputsServed     []VRFBeaconReportOutputServed
	Raw               types.Log
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
	Bin: "0x6080604052348015600f57600080fd5b50601680601d6000396000f3fe6080604052600080fdfea164736f6c634300080f000a",
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

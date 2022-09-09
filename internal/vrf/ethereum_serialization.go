package vrf

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"github.com/pkg/errors"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/mod"

	"github.com/smartcontractkit/ocr2vrf/altbn_128"
	"github.com/smartcontractkit/ocr2vrf/gethwrappers/vrfbeaconcoordinator"

	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

type EthereumReportSerializer struct {
	G kyber.Group
}

func (e *EthereumReportSerializer) DeserializeReport(bytes []byte) (vrf_types.AbstractReport, error) {
	arguments := vrfABI().Methods["exposeType"].Inputs
	report := make(map[string]interface{}, 0)
	err := arguments.UnpackIntoMap(report, bytes)
	if err != nil {
		return vrf_types.AbstractReport{}, err
	}

	rep := report[""].(struct {
		Outputs []struct {
			BlockHeight       uint64   "json:\"blockHeight\""
			ConfirmationDelay *big.Int "json:\"confirmationDelay\""
			VrfOutput         struct {
				P [2]*big.Int "json:\"p\""
			} "json:\"vrfOutput\""
			Callbacks []struct {
				Callback struct {
					RequestID    *big.Int       "json:\"requestID\""
					NumWords     uint16         "json:\"numWords\""
					Requester    common.Address "json:\"requester\""
					Arguments    []uint8        "json:\"arguments\""
					SubID        uint64         "json:\"subID\""
					GasAllowance *big.Int       "json:\"gasAllowance\""
				} "json:\"callback\""
				Price *big.Int "json:\"price\""
			} "json:\"callbacks\""
		} "json:\"outputs\""
		JuelsPerFeeCoin   *big.Int  "json:\"juelsPerFeeCoin\""
		RecentBlockHeight uint64    "json:\"recentBlockHeight\""
		RecentBlockHash   [32]uint8 "json:\"recentBlockHash\""
	})
	outputs := make([]vrfbeaconcoordinator.VRFBeaconReportVRFOutput, 0)
	for _, out := range rep.Outputs {
		callbacks := make([]vrfbeaconcoordinator.VRFBeaconTypesCostedCallback, 0)
		for _, c := range out.Callbacks {
			callback := vrfbeaconcoordinator.VRFBeaconTypesCallback{
				c.Callback.RequestID,
				c.Callback.NumWords,
				c.Callback.Requester,
				c.Callback.Arguments,
				c.Callback.SubID,
				c.Callback.GasAllowance,
			}
			costedCallback := vrfbeaconcoordinator.VRFBeaconTypesCostedCallback{
				callback,
				c.Price,
			}
			callbacks = append(callbacks, costedCallback)
		}
		output := vrfbeaconcoordinator.VRFBeaconReportVRFOutput{
			out.BlockHeight,
			out.ConfirmationDelay,
			vrfbeaconcoordinator.ECCArithmeticG1Point{
				[2]*big.Int{
					out.VrfOutput.P[0],
					out.VrfOutput.P[1]},
			},
			callbacks,
		}
		outputs = append(outputs, output)
	}
	beaconReport := vrfbeaconcoordinator.VRFBeaconReportReport{
		outputs,
		rep.JuelsPerFeeCoin,
		rep.RecentBlockHeight,
		rep.RecentBlockHash,
	}
	abstractReport, err := e.ConvertToAbstractReport(beaconReport)
	if err != nil {
		return vrf_types.AbstractReport{}, err
	}
	return abstractReport, err
}

func (e *EthereumReportSerializer) MaxReportLength() uint {

	panic("implement me")
}

func (e *EthereumReportSerializer) ReportLength(abstractReport vrf_types.AbstractReport) uint {

	panic("implement me")
}

var _ vrf_types.ReportSerializer = &EthereumReportSerializer{}

func (e *EthereumReportSerializer) SerializeReport(
	report vrf_types.AbstractReport,
) ([]byte, error) {
	beaconReport, err := e.ConvertToBeaconReport(report)
	if err != nil {
		return nil, err
	}
	arguments := vrfABI().Methods["exposeType"].Inputs
	serialziedReport, err := arguments.Pack(beaconReport)
	if err != nil {
		return nil, err
	}
	return serialziedReport, err
}

func (e *EthereumReportSerializer) ConvertToBeaconReport(
	report vrf_types.AbstractReport,
) (vrfbeaconcoordinator.VRFBeaconReportReport, error) {
	vrfOutputs := make(
		[]vrfbeaconcoordinator.VRFBeaconReportVRFOutput, 0, len(report.Outputs),
	)
	emptyReport := vrfbeaconcoordinator.VRFBeaconReportReport{}
	for _, output := range report.Outputs {
		p := e.G.Point()
		if err := p.UnmarshalBinary(output.VRFProof[:]); err != nil {
			return emptyReport, errors.Wrap(err, "while unmarshalling vrf proof")
		}
		x, y := big.NewInt(0), big.NewInt(0)
		if !p.Equal(e.G.Point().Null()) {
			x, y = affineCoordinates(p)
		}
		callbacks := make(
			[]vrfbeaconcoordinator.VRFBeaconTypesCostedCallback, 0,
			len(output.Callbacks),
		)
		for _, callback := range output.Callbacks {
			costedCallback := vrfbeaconcoordinator.VRFBeaconTypesCallback{
				big.NewInt(0).SetUint64(callback.RequestID),
				callback.NumWords,
				callback.Requester,
				callback.Arguments,
				callback.SubscriptionID,
				callback.GasAllowance,
			}
			beaconCostedCallback := vrfbeaconcoordinator.VRFBeaconTypesCostedCallback{
				costedCallback, callback.Price,
			}
			callbacks = append(callbacks, beaconCostedCallback)
		}
		vrfOutput := vrfbeaconcoordinator.VRFBeaconReportVRFOutput{
			output.BlockHeight,
			big.NewInt(int64(output.ConfirmationDelay)),
			vrfbeaconcoordinator.ECCArithmeticG1Point{[2]*big.Int{x, y}},
			callbacks,
		}
		vrfOutputs = append(vrfOutputs, vrfOutput)
	}
	onchainReport := vrfbeaconcoordinator.VRFBeaconReportReport{
		vrfOutputs,
		report.JulesPerFeeCoin,
		report.RecentBlockHeight,
		report.RecentBlockHash,
	}
	return onchainReport, nil
}

func (e *EthereumReportSerializer) ConvertToAbstractReport(
	report vrfbeaconcoordinator.VRFBeaconReportReport,
) (vrf_types.AbstractReport, error) {
	abstracOutputs := make([]vrf_types.AbstractVRFOutput, 0)
	for _, out := range report.Outputs {
		x := mod.NewInt(out.VrfOutput.P[0], bn256.P)
		y := mod.NewInt(out.VrfOutput.P[1], bn256.P)
		p := altbn_128.CoordinatesToG1(x, y)
		temp, err := p.MarshalBinary()
		if err != nil {
			return vrf_types.AbstractReport{}, errors.Wrap(err, "while unmarshalling vrf proof")
		}
		abstractCallbacks := make([]vrf_types.AbstractCostedCallbackRequest, 0)
		for _, c := range out.Callbacks {
			abstractCallback := vrf_types.AbstractCostedCallbackRequest{
				out.BlockHeight,
				uint32(out.ConfirmationDelay.Uint64()),
				c.Callback.SubID,
				c.Price,
				c.Callback.RequestID.Uint64(),
				c.Callback.NumWords,
				c.Callback.Requester,
				c.Callback.Arguments,
				c.Callback.GasAllowance,
				0,
				common.Hash{},
			}
			abstractCallbacks = append(abstractCallbacks, abstractCallback)
		}
		var vrfProof [32]byte
		copy(vrfProof[:], temp[:])
		abstractOutput := vrf_types.AbstractVRFOutput{
			out.BlockHeight,
			uint32(out.ConfirmationDelay.Uint64()),
			vrfProof,
			abstractCallbacks,
		}
		abstracOutputs = append(abstracOutputs, abstractOutput)
	}
	abstractReport := vrf_types.AbstractReport{
		abstracOutputs,
		report.JuelsPerFeeCoin,
		report.RecentBlockHeight,
		report.RecentBlockHash,
	}
	return abstractReport, nil
}

func vrfABI() *abi.ABI {
	rv, err := abi.JSON(
		strings.NewReader(vrfbeaconcoordinator.VRFBeaconReportMetaData.ABI),
	)
	if err != nil {
		panic(err)
	}
	return &rv
}

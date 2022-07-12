package vrf

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"go.dedis.ch/kyber/v3"

	"github.com/smartcontractkit/ocr2vrf/gethwrappers/vrfbeaconcoordinator"

	vrf_types "github.com/smartcontractkit/ocr2vrf/types"
)

type EthereumReportSerializer struct {
	G kyber.Group
}

func (e *EthereumReportSerializer) DeserializeReport(bytes []byte) (vrf_types.AbstractReport, error) {

	panic("implement me")
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
		err := p.UnmarshalBinary(output.VRFProof[:])
		if err != nil {
			return emptyReport, errors.Wrap(err, "while unmarshalling vrf proof")
		}
		x, y := big.NewInt(0), big.NewInt(0)
		if !p.Equal(e.G.Point().Null()) {
			x, y, err = affineCoordinates(p)
			if err != nil {
				return emptyReport, errors.Wrap(
					err, "while computing the affine coordinates",
				)
			}
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

func vrfABI() *abi.ABI {
	rv, err := abi.JSON(strings.NewReader(vrfbeaconcoordinator.VRFBeaconReportMetaData.ABI))
	if err != nil {
		panic(err)
	}
	return &rv
}

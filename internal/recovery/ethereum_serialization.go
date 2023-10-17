package recovery

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/ocr2vrf/gethwrappers/recoverybeacon"
	recovery_types "github.com/smartcontractkit/ocr2vrf/types"

	"go.dedis.ch/kyber/v3"
)

type EthereumReportSerializer struct {
	G kyber.Group
}

func (e *EthereumReportSerializer) DeserializeReport(
	rawReport []byte,
) (recovery_types.AbstractReport, error) {
	arguments := recoveryABI().Methods["exposeType"].Inputs
	parsedReport, err := arguments.Unpack(rawReport)
	if err != nil {
		return recovery_types.AbstractReport{}, err
	}
	type reportType = recovery_types.AbstractReport
	report := reportType{}
	abi.ConvertType(parsedReport[0], &report)
	return report, err
}

var _ recovery_types.ReportSerializer = &EthereumReportSerializer{}

func (e *EthereumReportSerializer) SerializeReport(
	report recovery_types.AbstractReport,
) ([]byte, error) {
	arguments := recoveryABI().Methods["exposeType"].Inputs
	serializedReport, err := arguments.Pack(report)
	if err != nil {
		return nil, err
	}
	return serializedReport, err
}

func recoveryABI() *abi.ABI {
	rv, err := abi.JSON(
		strings.NewReader(recoverybeacon.RecoveryBeaconReportMetaData.ABI),
	)
	if err != nil {
		panic(err)
	}
	return &rv
}

func (e *EthereumReportSerializer) MaxReportLength() uint {

	panic("implement me")
}

func (e *EthereumReportSerializer) ReportLength(abstractReport recovery_types.AbstractReport) uint {

	panic("implement me")
}

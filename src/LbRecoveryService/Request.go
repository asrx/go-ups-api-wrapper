package LbRecoveryService

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/LbRecoveryService/ComplexType"
	"github.com/asrx/gowsdl/soap"
)

type PortType interface {

	// Error can be either of the following types:
	//
	//   - LabelRecoveryError

	ProcessLabelRecovery(request *LabelRecoveryRequest) (*LabelRecoveryResponse, *ErrorDetail)
}

type portType struct {
	client *soap.Client
}

var _ProductionUrl = "https://onlinetools.ups.com/webservices/LBRecovery"
var _TestUrl = "https://wwwcie.ups.com/webservices/LBRecovery"

func NewPortType(security *UPSSecurity, testEnv bool) PortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}
	client.AddHeader(security)

	return &portType{
		client: client,
	}
}

func (service *portType) ProcessLabelRecovery(request *LabelRecoveryRequest) (*LabelRecoveryResponse, *ErrorDetail) {
	response := new(LabelRecoveryResponse)

	errResponse := new(ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/ShipBinding/v1.1", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}

	return response, nil
}

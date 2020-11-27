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

func NewPortType(client *soap.Client) PortType {
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

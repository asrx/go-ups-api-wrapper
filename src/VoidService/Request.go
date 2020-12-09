package VoidService

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/VoidService/ComplexType"
	"github.com/asrx/gowsdl/soap"
)

type VoidPortType interface {

	// Error can be either of the following types:
	//
	//   - VoidError

	ProcessVoid(request *VoidShipmentRequest) (*VoidShipmentResponse, *ErrorDetail)
}

type voidPortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://onlinetools.ups.com/webservices/Void"
var _TestUrl = "https://wwwcie.ups.com/webservices/Void"

func NewVoidPortType(security *UPSSecurity, testEnv bool) VoidPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_ProductionUrl)
	}else {
		client = soap.NewClient(_TestUrl)
	}
	client.AddHeader(security)

	return &voidPortType{
		client: client,
	}
}

func (service *voidPortType) ProcessVoid(request *VoidShipmentRequest) (*VoidShipmentResponse, *ErrorDetail) {
	response := new(VoidShipmentResponse)

	errResponse := new(ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/VoidBinding/v1.1", request, response, errResponse)

	if err != nil {
		return nil, errResponse
	}

	return response, nil
}

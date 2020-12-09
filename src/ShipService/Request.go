package ShipService

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/ShipService/ComplexType"
	"github.com/asrx/gowsdl/soap"
)

type ShipPortType interface {

	// Error can be either of the following types:
	//
	//   - ShipmentError

	ProcessShipment(request *ShipmentRequest) (*ShipmentResponse, *ErrorDetail)

	// Error can be either of the following types:
	//
	//   - ShipConfirmError

	ProcessShipConfirm(request *ShipConfirmRequest) (*ShipConfirmResponse, *ErrorDetail)

	// Error can be either of the following types:
	//
	//   - ShipAcceptError

	ProcessShipAccept(request *ShipAcceptRequest) (*ShipAcceptResponse, *ErrorDetail)
}

type shipPortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://onlinetools.ups.com/webservices/Ship"
var _TestUrl = "https://wwwcie.ups.com/webservices/Ship"

func NewShipPortType(security *UPSSecurity, testEnv bool) ShipPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_ProductionUrl)
	}else {
		client = soap.NewClient(_TestUrl)
	}
	client.AddHeader(security)

	return &shipPortType{
		client: client,
	}
}

func (service *shipPortType) ProcessShipment(request *ShipmentRequest) (*ShipmentResponse, *ErrorDetail) {
	response := new(ShipmentResponse)

	errResponse := new(ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/ShipBinding/v1.0", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}

	return response, nil
}

func (service *shipPortType) ProcessShipConfirm(request *ShipConfirmRequest) (*ShipConfirmResponse, *ErrorDetail) {
	response := new(ShipConfirmResponse)
	errResponse := new(ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/ShipBinding/v1.0", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}

	return response, nil
}

func (service *shipPortType) ProcessShipAccept(request *ShipAcceptRequest) (*ShipAcceptResponse, *ErrorDetail) {
	response := new(ShipAcceptResponse)
	errResponse := new(ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/ShipBinding/v1.0", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}

	return response, nil
}

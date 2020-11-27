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

func NewVoidPortType(client *soap.Client) VoidPortType {
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

package RateService

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/RateService/ComplexType"
	"github.com/asrx/gowsdl/soap"
)

type RatePortType interface {
	// Error can be either of the following types:
	//
	//   - RateError

	ProcessRate(request *RateRequest) (*RateResponse, *ErrorDetail)
}

type ratePortType struct {
	client *soap.Client
}

func NewRatePortType(client *soap.Client) RatePortType {
	return &ratePortType{
		client: client,
	}
}

func (service *ratePortType) ProcessRate(request *RateRequest) (*RateResponse, *ErrorDetail) {
	response := new(RateResponse)

	errResponse := new(ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/RateBinding/v1.1", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}

	return response, nil
}
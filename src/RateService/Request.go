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

var _ProductionUrl = "https://onlinetools.ups.com/webservices/Rate"
var _TestUrl = "https://wwwcie.ups.com/webservices/Rate"

func NewRatePortType(security *UPSSecurity, testEnv bool) RatePortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}
	client.AddHeader(security)

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
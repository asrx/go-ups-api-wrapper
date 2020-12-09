package RateService

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/RateService/ComplexType"
	"github.com/asrx/gowsdl/soap"
)

type FreightRatePortType interface {

	// Error can be either of the following types:
	//
	//   - RateError

	ProcessFreightRate(request *FreightRateRequest) (*FreightRateResponse, *ComplexType.ErrorDetail)
}

type freightRatePortType struct {
	client *soap.Client
}

var _FreightProductionUrl = "https://onlinetools.ups.com/webservices/FreightRate"
var _FreightTestUrl = "https://wwwcie.ups.com/webservices/FreightRate"

func NewFreightRatePortType(security *ComplexType.UPSSecurity, testEnv bool) FreightRatePortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_FreightTestUrl)
	}else {
		client = soap.NewClient(_FreightProductionUrl)
	}
	client.AddHeader(security)

	return &freightRatePortType{
		client: client,
	}
}

func (service *freightRatePortType) ProcessFreightRate(request *FreightRateRequest) (*FreightRateResponse, *ComplexType.ErrorDetail) {
	response := new(FreightRateResponse)

	errResponse := new(ComplexType.ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/FreightRateBinding/v1.0", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}
	return response, nil
}

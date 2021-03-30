package AddressValidationService



import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/AddressValidationService/ComplexType"
	"github.com/asrx/gowsdl/soap"
)

var _ProductionUrl = "https://onlinetools.ups.com/webservices/XAV"
var _TestUrl = "https://wwwcie.ups.com/webservices/XAV"

type XAVPortType interface {

	// Error can be either of the following types:
	//
	//   - XAVError

	ProcessXAV(request *XAVRequest) (*XAVResponse, error)
}

type xAVPortType struct {
	client *soap.Client
}

func NewXAVPortType(security *UPSSecurity, testEnv bool) *xAVPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else{
		client = soap.NewClient(_ProductionUrl)
	}
	client.AddHeader(security)

	return &xAVPortType{
		client: client,
	}
}

func (service *xAVPortType) ProcessXAV(request *XAVRequest) (*XAVResponse, *ErrorDetail) {
	response := new(XAVResponse)

	errResponse := new(ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/XAVBinding/v1.0", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}

	return response, nil
}
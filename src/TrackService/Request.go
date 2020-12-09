package TrackService

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/TrackService/ComplexType"
	"github.com/asrx/gowsdl/soap"
)
type TrackPortType interface {

	// Error can be either of the following types:
	//
	//   - TrackError

	ProcessTrack(request *TrackRequest) (*TrackResponse, *ComplexType.ErrorDetail)
}

type trackPortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://onlinetools.ups.com/webservices/Track"
var _TestUrl = "https://wwwcie.ups.com/webservices/Track"

func NewTrackPortType(security *ComplexType.UPSSecurity, testEnv bool) TrackPortType {

	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}
	client.AddHeader(security)

	return &trackPortType{
		client: client,
	}
}

func (service *trackPortType) ProcessTrack(request *TrackRequest) (*TrackResponse, *ComplexType.ErrorDetail) {
	response := new(TrackResponse)
	errResponse := new(ComplexType.ErrorDetail)

	err := service.client.Call("http://onlinetools.ups.com/webservices/TrackBinding/v2.0", request, response, errResponse)
	if err != nil {
		return nil, errResponse
	}

	return response, nil
}

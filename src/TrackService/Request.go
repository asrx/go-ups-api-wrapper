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

func NewTrackPortType(client *soap.Client) TrackPortType {
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

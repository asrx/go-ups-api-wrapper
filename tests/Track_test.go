package tests

import (
	"github.com/asrx/go-ups-api-wrapper/src/TrackService"
	. "github.com/asrx/go-ups-api-wrapper/src/TrackService/ComplexType"
	"encoding/json"
	"fmt"
	"github.com/asrx/gowsdl/soap"
	"testing"
)

func Test_Track(t *testing.T) {

	var addr = "Track"

	var trackNumber = "1Z12345E0305271640"
	fmt.Println(trackNumber)

	request := &TrackRequest{
		InquiryNumber:                  trackNumber,
	}

	client := soap.NewClient(GetRequestUrl(addr))
	client.AddHeader(GetSoapHeaderSecurity())
	c := TrackService.NewTrackPortType(client)
	resp, fault := c.ProcessTrack(request)

	if fault != nil {

		fmt.Println("==================================================")
		for _, exp := range fault.Errors.ErrorDetail{
			fmt.Printf("%s: %s\n", exp.PrimaryErrorCode.Code, exp.PrimaryErrorCode.Description)
		}
		fmt.Println("**************************************************")
		return
	}

	for _,p := range resp.Shipment.Package {
		fmt.Printf("%s:\n",p.TrackingNumber)
		for _, p2 :=range p.Activity{
			fmt.Printf("%s[%s] \t - \t %s %s\n",p2.ActivityLocation.Description, p2.Status.Description, p2.Date, p2.Time)
		}
		fmt.Println("===========================")
	}

	byte, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n",resp)
	}else{
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}
}

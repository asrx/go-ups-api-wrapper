package tests

import (
	"encoding/json"
	"fmt"
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-ups-api-wrapper/src/TrackService"
	. "github.com/asrx/go-ups-api-wrapper/src/TrackService/ComplexType"
	"testing"
)

func Test_Track(t *testing.T) {
	var trackNumber = "1Z7A03740307372858"

	//TrackRequest{
	//	Request:                        nil,
	//	InquiryNumber:                  "",
	//	IncludeMailInnovationIndicator: "",
	//	TrackingOption:                 "",
	//	CandidateBookmark:              "",
	//	ShipperAccountInfo:             nil,
	//	PreauthorizedReturnIndicator:   "",
	//	Locale:                         "",
	//}
	request := &TrackRequest{
		Request:                        &ComplexType.RequestType{
			RequestOption:        []string{"activity"},
		},
		InquiryNumber:                  trackNumber,
	}

	c := TrackService.NewTrackPortType(GetSoapHeaderSecurity(), true)
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
			switch p2.Status.Code {
			case "DP", "AR", "FS":
				fmt.Printf("%s %s | %s : %s %s\n", p2.Date, p2.Time, p2.Status.Description, p2.ActivityLocation.Address.City, p2.ActivityLocation.Address.StateProvinceCode)
			default:
				fmt.Printf("%s %s | %s\n", p2.Date, p2.Time, p2.Status.Description)
			}
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

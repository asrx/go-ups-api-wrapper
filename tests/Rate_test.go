package tests

import (
	ComplexType2 "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-ups-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-ups-api-wrapper/src/RateService"
	"github.com/asrx/go-ups-api-wrapper/src/RateService/ComplexType"
	"encoding/json"
	"fmt"
	"github.com/asrx/gowsdl/soap"
	"testing"
)

func Test_Rate(t *testing.T) {

	var addr = "Rate"
	request := &ComplexType.RateRequest{
		Request: 				&ComplexType2.RequestType{
			RequestOption:        []string{"Rate"},
		},
		Shipment:               &ComplexType.ShipmentType{
			Shipper:                            Shipper,
			ShipTo:                             Recipient,
			ShipFrom:							ShippFrom,
			//FRSPaymentInformation:              &ComplexType.FRSPaymentInfoType{
			//	Type:          &ComplexType2.CodeDescriptionType{
			//		Code:        SimpleType.CODE_PREPAID,
			//	},
			//	AccountNumber: _ShipperNumber,
			//},
			Service:                            &ComplexType2.CodeDescriptionType{
				Code:        SimpleType.S_GROUND,
			},
			Package:                            GetRatePackages(),
			ShipmentRatingOptions: &ComplexType2.ShipmentRatingOptionsType{
				NegotiatedRatesIndicator:   "true",
				//FRSShipmentIndicator:       "true",
			},
			RatingMethodRequestedIndicator:     "true",
			TaxInformationIndicator:            "true",
		},
	}

	client := soap.NewClient(GetRequestUrl(addr))
	client.AddHeader(GetSoapHeaderSecurity())
	c := RateService.NewRatePortType(client)
	resp, fault := c.ProcessRate(request)

	if fault != nil {

		fmt.Println("==================================================")
		for _, exp := range fault.Errors.ErrorDetail{
			fmt.Printf("%s: %s\n", exp.PrimaryErrorCode.Code, exp.PrimaryErrorCode.Description)
		}
		fmt.Println("**************************************************")
		return
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

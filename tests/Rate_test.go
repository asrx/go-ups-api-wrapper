package tests

import (
	"encoding/json"
	"fmt"
	ComplexType2 "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-ups-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-ups-api-wrapper/src/RateService"
	"github.com/asrx/go-ups-api-wrapper/src/RateService/ComplexType"
	"testing"
)

func Test_RateGround(t *testing.T) {
	request := &ComplexType.RateRequest{
		Request: 				&ComplexType2.RequestType{
			RequestOption:        []string{"Rate"},
		},
		Shipment:               &ComplexType.ShipmentType{
			Shipper:                            Shipper,
			ShipTo:                             Recipient,
			ShipFrom:							ShippFrom,
			FRSPaymentInformation:              &ComplexType2.FRSPaymentInfoType{
				Type:          &ComplexType2.CodeDescriptionType{
					Code:        SimpleType.CODE_PREPAID,
				},
				AccountNumber: _ShipperNumber,
			},
			Service:                            &ComplexType2.CodeDescriptionType{
				Code:        SimpleType.S_GROUND,
			},
			Package:                            GetRatePackages(),
			ShipmentRatingOptions: &ComplexType2.ShipmentRatingOptionsType{
				NegotiatedRatesIndicator:   "true",
			//	FRSShipmentIndicator:       "true",
			},

			RatingMethodRequestedIndicator:     "true",
			TaxInformationIndicator:            "true",
		},
	}

	c := RateService.NewRatePortType(GetSoapHeaderSecurity(), false)
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
		if resp.RatedShipment[0].NegotiatedRateCharges != nil {
			fmt.Printf("协议价：%v\n",resp.RatedShipment[0].NegotiatedRateCharges.TotalCharge.MonetaryValue)
		}
		fmt.Println(resp.RatedShipment[0].TotalCharges.MonetaryValue)
	}
}

func Test_RateGFP(t *testing.T) {
	request := &ComplexType.RateRequest{
		Request: 				&ComplexType2.RequestType{
			RequestOption:        []string{"Rate"},
		},
		Shipment:               &ComplexType.ShipmentType{
			Shipper:                            Shipper,
			ShipTo:                             Recipient,
			ShipFrom:							ShippFrom,
			// must
			FRSPaymentInformation:              &ComplexType2.FRSPaymentInfoType{
				Type:          &ComplexType2.CodeDescriptionType{
					Code:        SimpleType.CODE_PREPAID,
				},
				AccountNumber: _ShipperNumber,
			},
			Service:                            &ComplexType2.CodeDescriptionType{
				Code:        SimpleType.S_GROUND,
			},
			Package:                            GetRatePackagesGfp(),
			ShipmentRatingOptions: &ComplexType2.ShipmentRatingOptionsType{
				NegotiatedRatesIndicator:   "true",
				// must
				FRSShipmentIndicator:       "true",
			},

			//RatingMethodRequestedIndicator:     "true",
			//TaxInformationIndicator:            "true",
		},
	}
	security := GetSoapHeaderSecurity()
	c := RateService.NewRatePortType(security, false)
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
		if resp.RatedShipment[0].NegotiatedRateCharges != nil {
			fmt.Printf("协议价：%v\n",resp.RatedShipment[0].NegotiatedRateCharges.TotalCharge.MonetaryValue)
		}

		fmt.Printf("TotalCharges: %v\n", resp.RatedShipment[0].TotalCharges.MonetaryValue)
	}
}

// Not through
func Test_RateHundredweight(t *testing.T) {

	request := &ComplexType.RateRequest{
		Request: 				&ComplexType2.RequestType{
			RequestOption:        []string{"Rate"},
		},
		Shipment:               &ComplexType.ShipmentType{
			Shipper:                            Shipper,
			ShipTo:                             Recipient,
			ShipFrom:							ShippFrom,
			FRSPaymentInformation:              &ComplexType2.FRSPaymentInfoType{
				Type:          &ComplexType2.CodeDescriptionType{
					Code:        SimpleType.CODE_PREPAID,
				},
				AccountNumber: _ShipperNumber,
			},
			Service:                            &ComplexType2.CodeDescriptionType{
				Code:        SimpleType.S_GROUND,
			},
			Package:                            GetRatePackagesHundredweight(),
			ShipmentRatingOptions: &ComplexType2.ShipmentRatingOptionsType{
				NegotiatedRatesIndicator:   "1",
				//	FRSShipmentIndicator:       "true",
			},

			RatingMethodRequestedIndicator:     "true",
			TaxInformationIndicator:            "true",
		},
	}

	c := RateService.NewRatePortType(GetSoapHeaderSecurity(), false)
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
		fmt.Println(resp.RatedShipment[0].NegotiatedRateCharges.TotalCharge.MonetaryValue)
		fmt.Println(resp.RatedShipment[0].TotalCharges.MonetaryValue)
	}
}
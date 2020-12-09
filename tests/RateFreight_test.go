package tests

import (
	"encoding/json"
	"fmt"
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-ups-api-wrapper/src/RateService"
	. "github.com/asrx/go-ups-api-wrapper/src/RateService/ComplexType"
	. "github.com/asrx/go-ups-api-wrapper/src/RateService/SimpleType"
	"testing"
)

func Test_RateFreight(t *testing.T) {
	request := &FreightRateRequest{
		Request:                  &RequestType{
			RequestOption:        []string{"Rate Freight"},
		},
		ShipFrom:                 ShippFrom,
		ShipTo:                   Recipient,
		PaymentInformation:       &PaymentInformationType{
			Payer:                 Shipper,
			ShipmentBillingOption: &CodeDescriptionType{
				Code:        ShipmentBillingOption_Prepaid,
			},
		},
		Service:                  &CodeDescriptionType{
			Code:        FreightServiceCode_FreightLTL,
		},
		//HandlingUnitOne:          &HandlingUnitType{
		//	Quantity: "20",
		//	Type:     &CodeDescriptionType{
		//		Code:        "PLT",
		//		Description: "PALLET",// 托盘
		//	},
		//},
		Commodity:                []*FreightCommodityType{
			&FreightCommodityType{
				CommodityID:             "",
				Description:			 "NO Description",
				Weight:                  &WeightType{
					Value:             "200",
					UnitOfMeasurement: &UnitOfMeasurementType{
						Code:        UOM_LBS,
					},
				},
				Dimensions:              &DimensionsType{
					UnitOfMeasurement: &CodeDescriptionType{
						Code:        UOM_IN,
					},
					Length:            "4",
					Width:             "4",
					Height:            "4",
				},
				NumberOfPieces:          "45",
				PackagingType:           &CodeDescriptionType{
					Code:        PROD_BAG,
				},
				DangerousGoodsIndicator: "",
				CommodityValue:          &CommodityValueType{
					CurrencyCode:  "USD",
					MonetaryValue: "5670",
				},
				FreightClass:            CLASS_50,
				NMFCCommodityCode:       "",
				//NMFCCommodity:           nil,
			},
		},
		//ShipmentServiceOptions:   nil,
		//PickupRequest:            nil,
		AlternateRateOptions:     &CodeDescriptionType{
			Code:        AlternateRateOptions_ALL,
		},
		//GFPOptions:               &GFPOptionsType{
		//	GPFAccesorialRateIndicator: "",
		//	OnCallInformation:          &OnCallInformationType{
		//		OnCallPickupIndicator:UOM_LBS,
		//	},
		//},
		//AccountType:              nil,
		//ShipmentTotalWeight:      nil,
		//HandlingUnitWeight:       nil,
		//AdjustedWeightIndicator:  "",
		//TimeInTransitIndicator:   "",
		//HandlingUnits:            nil,
		//AdjustedHeightIndicator:  "",
		//DensityEligibleIndicator: "",
		//QuoteNumberIndicator:     "",
	}

	//request := &ComplexType.FreightRateRequest{
	//	Request: 				&ComplexType2.RequestType{
	//		RequestOption:        []string{"Rate"},
	//	},
	//	Shipment:               &ComplexType.ShipmentType{
	//		Shipper:                            Shipper,
	//		ShipTo:                             Recipient,
	//		ShipFrom:							ShippFrom,
	//		// must
	//		FRSPaymentInformation:              &ComplexType2.FRSPaymentInfoType{
	//			Type:          &ComplexType2.CodeDescriptionType{
	//				Code:        SimpleType.CODE_PREPAID,
	//			},
	//			AccountNumber: _ShipperNumber,
	//		},
	//		Service:                            &ComplexType2.CodeDescriptionType{
	//			Code:        SimpleType.S_GROUND,
	//		},
	//		Package:                            GetRatePackagesGfp(),
	//		ShipmentRatingOptions: &ComplexType2.ShipmentRatingOptionsType{
	//			NegotiatedRatesIndicator:   "true",
	//			// must
	//			FRSShipmentIndicator:       "true",
	//		},
	//
	//		//RatingMethodRequestedIndicator:     "true",
	//		//TaxInformationIndicator:            "true",
	//	},
	//}

	security := GetSoapHeaderSecurity()
	c := RateService.NewFreightRatePortType(security, true)
	resp, fault := c.ProcessFreightRate(request)

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

		//fmt.Printf("TotalCharges: %v\n", resp.)
	}
}

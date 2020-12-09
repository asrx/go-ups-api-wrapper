package tests

import (
	"encoding/json"
	"fmt"
	ComplexType2 "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-ups-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-ups-api-wrapper/src/ShipService"
	. "github.com/asrx/go-ups-api-wrapper/src/ShipService/ComplexType"
	"testing"
)

func Test_Ship(t *testing.T) {

	request := &ShipmentRequest{
		Request:              &ComplexType2.RequestType{
			RequestOption:        []string{SimpleType.REQ_NONVALIDATE},
		},
		Shipment:             &ShipmentType{
			Shipper:                            Shipper,
			ShipTo:                             Recipient,
			//AlternateDeliveryAddress:           nil,
			ShipFrom:                           ShippFrom,
			PaymentInformation:                 &PaymentInfoType{
				ShipmentCharge:        &ShipmentChargeType{
					Type:                     "01",
					BillShipper:              &BillShipperType{
						AccountNumber:          _ShipperNumber,
					},
				},
			},

			//FRSPaymentInformation:              &ComplexType.FRSPaymentInfoType{
			//	Type:          &ComplexType2.CodeDescriptionType{
			//		Code:        SimpleType.CODE_PREPAID,
			//	},
			//	AccountNumber: _ShipperNumber,
			//},
			//FreightShipmentInformation:         nil,
			//GoodsNotInFreeCirculationIndicator: "",
			ShipmentRatingOptions:              &ComplexType2.ShipmentRatingOptionsType{
				NegotiatedRatesIndicator:   "true",
				//FRSShipmentIndicator:       "",
			},
			RatingMethodRequestedIndicator:     "true",
			TaxInformationIndicator:            "true",
			Service:                            &ComplexType2.CodeDescriptionType{
				Code:        SimpleType.S_GROUND,
			},
			//Package: 							GetRatePackages(),
			Package: 							GetShipPackages(),


			//MovementReferenceNumber:            "",
			//ReferenceNumber:                    &ComplexType2.ReferenceNumberType{
			//	Code:  SimpleType.CODE_PRODUCTION_CODE,
			//	//Code:  SimpleType.CODE_INVOICE_NUMBER,
			//	Value: "XXXTestTTTT",
			//},
			//InvoiceLineTotal:                   nil,
			//NumOfPiecesInShipment:              "",
			//USPSEndorsement:                    "",
			//MILabelCN22Indicator:               "",
			//SubClassification:                  "",
			//CostCenter:                         "",
			//CostCenterBarcodeIndicator:         "",
			//PackageID:                          "",
			//PackageIDBarcodeIndicator:          "",
			//IrregularIndicator:                 "",
			//MIDualReturnShipmentKey:            "",
			//MIDualReturnShipmentIndicator:      "",

			//PromotionalDiscountInformation:     nil,
			//DGSignatoryInfo:                    nil,
			//MasterCartonID:                     "",
			//MasterCartonIndicator:              "",
			//ShipmentServiceOptions:     nil,
		},
		LabelSpecification:   &ComplexType2.LabelSpecificationType{
			LabelImageFormat: &ComplexType2.LabelImageFormatType{
				Code:        SimpleType.IMG_FORMAT_CODE_GIF,
			},
			//LabelStockSize:   &LabelStockSizeType{
			//	Height: "",
			//	Width:  "",
			//},
		},
	}

	c := ShipService.NewShipPortType(GetSoapHeaderSecurity(), true)
	resp, fault := c.ProcessShipment(request)

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

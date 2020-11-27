package tests

import (
	ComplexType2 "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-ups-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-ups-api-wrapper/src/LbRecoveryService"
	"github.com/asrx/go-ups-api-wrapper/src/LbRecoveryService/ComplexType"
	"encoding/json"
	"fmt"
	"github.com/asrx/gowsdl/soap"
	"testing"
)

func Test_Recovery(t *testing.T) {

	addr := "LBRecovery"

	request := &ComplexType.LabelRecoveryRequest{
		LabelSpecification:   &ComplexType2.LabelSpecificationType{
			LabelImageFormat: &ComplexType2.LabelImageFormatType{
				Code:        SimpleType.IMG_FORMAT_CODE_GIF,
			},
		},
		TrackingNumber:                "1Z557FW50337911210",
	}

	client := soap.NewClient(GetRequestUrl(addr))
	client.AddHeader(GetSoapHeaderSecurity())
	c := LbRecoveryService.NewPortType(client)
	resp, fault := c.ProcessLabelRecovery(request)

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

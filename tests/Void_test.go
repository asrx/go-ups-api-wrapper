package tests

import (
	ComplexType2 "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-ups-api-wrapper/src/VoidService"
	"github.com/asrx/go-ups-api-wrapper/src/VoidService/ComplexType"
	"encoding/json"
	"fmt"
	"github.com/asrx/gowsdl/soap"
	"testing"
)

func Test_Cancel(t *testing.T) {

	var addr = "Void"

	request := &ComplexType.VoidShipmentRequest{
		Request:      &ComplexType2.RequestType{
			RequestOption:        []string{"1"},
		},
		VoidShipment: &ComplexType.VoidShipment{
			ShipmentIdentificationNumber: "1Z557FW50331231279",
			//TrackingNumber:               []string{"1Z557FW50393224947","1Z557FW50331443602"},
		},
	}

	client := soap.NewClient(GetRequestUrl(addr))
	client.AddHeader(GetSoapHeaderSecurity())
	c := VoidService.NewVoidPortType(client)
	resp, fault := c.ProcessVoid(request)

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
		fmt.Println("Result:", resp.Response.ResponseStatus.Description)
	}
}

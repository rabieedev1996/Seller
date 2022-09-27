package Service

import (
	"Seller/Seller.Domain/ServiceModel"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FarazSMSService struct{}

func (service FarazSMSService) SendSms(destphone string, verificationCode string) {

}
func (service FarazSMSService) SendVerificationCode(destphone string, verificationCode string) {
	url := "http://rest.ippanel.com/v1/messages/patterns/send"
	requestBody := ServiceModel.FarazSmsVerificationSMSRequest{
		Originator:  "3000505",
		PatternCode: "t7vcvo91jx",
		Recipient:   destphone,
		Value: ServiceModel.FarazSmsVerificationSMSValues{
			VerificationCode: verificationCode,
		},
	}

	jsonBytes, _ := json.Marshal(requestBody)
	fmt.Println("request Body:", string(jsonBytes))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "AccessKey hOam4QxZrawuabSWeIqY3z5H8atz7hOTy4sUJ_CEoZg=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

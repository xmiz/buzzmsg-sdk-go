package auth

import (
	"github.com/xmiz/buzzmsg-sdk-go/sdk/requests"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/utils"
	"net/url"
	"strings"
)

var hookGetNonce = func(fn func() string) string {
	return fn()
}

var hookGetDate = func(fn func() string) string {
	return fn()
}

func signHttpRequest(request requests.Request, signer Signer, regionId string) (err error) {
	err = completeHttpSignParams(request, signer, regionId)
	if err != nil {
		return
	}
	// remove while retry
	if _, containsSign := request.GetQueryParams()["Signature"]; containsSign {
		delete(request.GetQueryParams(), "Signature")
	}
	stringToSign := buildHttpStringToSign(request)
	request.SetStringToSign(stringToSign)
	signature := signer.Sign(stringToSign, "&")
	request.GetQueryParams()["Signature"] = signature
	return
}

func completeHttpSignParams(request requests.Request, signer Signer, regionId string) (err error) {
	queryParams := request.GetQueryParams()
	queryParams["Version"] = request.GetVersion()
	queryParams["Action"] = request.GetActionName()
	queryParams["Timestamp"] = hookGetDate(utils.GetTimeInFormatISO8601)
	queryParams["SignatureMethod"] = signer.GetName()
	queryParams["SignatureType"] = signer.GetType()
	queryParams["SignatureVersion"] = signer.GetVersion()
	queryParams["SignatureNonce"] = hookGetNonce(utils.GetUUID)
	queryParams["AccessKeyId"], err = signer.GetAccessKeyId()

	if err != nil {
		return
	}

	if _, contains := queryParams["RegionId"]; !contains {
		queryParams["RegionId"] = regionId
	}
	if extraParam := signer.GetExtraParam(); extraParam != nil {
		for key, value := range extraParam {
			queryParams[key] = value
		}
	}

	request.GetHeaders()["Content-Type"] = requests.Form
	formString := utils.GetUrlFormedMap(request.GetFormParams())
	request.SetContent([]byte(formString))

	return
}

func buildHttpStringToSign(request requests.Request) (stringToSign string) {
	signParams := make(map[string]string)
	for key, value := range request.GetQueryParams() {
		signParams[key] = value
	}
	for key, value := range request.GetFormParams() {
		signParams[key] = value
	}
	stringToSign = utils.GetUrlFormedMap(signParams)
	stringToSign = strings.Replace(stringToSign, "+", "%20", -1)
	stringToSign = strings.Replace(stringToSign, "*", "%2A", -1)
	stringToSign = strings.Replace(stringToSign, "%7E", "~", -1)
	stringToSign = url.QueryEscape(stringToSign)
	stringToSign = request.GetMethod() + "&%2F&" + stringToSign
	return
}
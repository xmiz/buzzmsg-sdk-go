package requests

type HttpRequest struct {
	*baseRequest
}

func (request *HttpRequest) init() {
	request.baseRequest = defaultBaseRequest()
	request.Method = POST
}

func (request *HttpRequest) InitWithApiInfo(product, version, action, serviceCode, endpointType string) {
	request.init()
	request.product = product
	request.version = version
	request.actionName = action
	request.locationServiceCode = serviceCode
	request.locationEndpointType = endpointType
	request.Headers["x-http-version"] = version
	request.Headers["x-http-action"] = action
}

func (request *baseRequest) SetStringToSign(stringToSign string) {
	request.stringToSign = stringToSign
}

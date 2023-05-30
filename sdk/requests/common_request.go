package requests
type CommonRequest struct {
	*baseRequest

	Version      string
	ApiName      string
	Product      string
	ServiceCode  string
	EndpointType string

	PathPattern string
	PathParams  map[string]string

	Ontology Request
}

func NewCommonRequest() (request *CommonRequest) {
	request = &CommonRequest{
		baseRequest: defaultBaseRequest(),
	}
	request.Headers["x-sdk-invoke-type"] = "common"
	request.PathParams = make(map[string]string)
	return
}

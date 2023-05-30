package requests

import (
	"encoding/json"
	"fmt"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	HTTP  = "HTTP"
	HTTPS = "HTTPS"

	DefaultHttpPort = "5000"

	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"

	Json = "application/json"
	Xml  = "application/xml"
	Raw  = "application/octet-stream"
	Form = "application/x-www-form-urlencoded"

	Header = "Header"
	Query  = "Query"
	Body   = "Body"
	Path   = "Path"

	HeaderSeparator = "\n"
	//ImsdkServerDebugHost = "http://10.26.0.10:7500/"
	ImsdkServerDebugHost   = "http://172.17.0.1:7500/"
	ImsdkServerReleaseHost = "http://sdkserver-api.tmmtmm.internal:5500/"
	//ImsdkServerDebugHost = "http://localhost:7500/"
)

// interface
type Request interface {
	GetScheme() string
	GetMethod() string
	GetDomain() string
	GetPort() string
	GetRegionId() string
	GetHeaders() map[string]string
	GetQueryParams() map[string]string
	GetFormParams() map[string]string
	GetContent() []byte
	GetBodyReader() io.Reader
	GetStyle() string
	GetProduct() string
	GetVersion() string
	SetVersion(version string)
	GetActionName() string
	GetAcceptFormat() string
	GetLocationServiceCode() string
	GetLocationEndpointType() string
	GetReadTimeout() time.Duration
	GetConnectTimeout() time.Duration
	SetReadTimeout(readTimeout time.Duration)
	SetConnectTimeout(connectTimeout time.Duration)
	SetHTTPSInsecure(isInsecure bool)
	GetHTTPSInsecure() *bool

	GetUserAgent() map[string]string

	SetStringToSign(stringToSign string)
	GetStringToSign() string

	SetDomain(domain string)
	SetContent(content []byte)
	SetScheme(scheme string)
	BuildUrl() string
	BuildQueries() string

	addHeaderParam(key, value string)
	addQueryParam(key, value string)
	addFormParam(key, value string)
	addPathParam(key, value string)
}

type baseRequest struct {
	Scheme               string
	Host                 string
	Method               string
	Domain               string
	Port                 string
	ReadTimeout          time.Duration
	ConnectTimeout       time.Duration
	userAgent            map[string]string
	version              string
	product              string
	actionName           string
	AcceptFormat         string
	QueryParams          map[string]string
	Headers              map[string]string
	FormParams           map[string]string
	Content              []byte
	locationServiceCode  string
	locationEndpointType string
	stringToSign         string
}

func New() *baseRequest {
	return &baseRequest{}
}

func defaultBaseRequest() (request *baseRequest) {
	request = &baseRequest{
		Scheme:       "",
		AcceptFormat: "JSON",
		Method:       POST,
		QueryParams:  make(map[string]string),
		Headers: map[string]string{
			"x-sdk-client":      "golang/1.0.0",
			"x-sdk-invoke-type": "normal",
			"Accept-Encoding":   "identity",
		},
		FormParams: make(map[string]string),
	}
	return
}

//func GetImsdkServerHost(ctx context.Context, model imsdkV2.ModelType) string {
//	//imsdkServerHost, _ := config.GetIMSDKServer()
//	//return imsdkServerHost
//	res := ImsdkServerDebugHost
//	//if model == imsdkV2.ModelRelease {
//	//	res = ImsdkServerReleaseHost
//	//}
//	return res
//}

func (request *baseRequest) Post() ([]byte, error) {
	params := request.GetContent()
	//byteData, _ := json.Marshal(params)
	//fmt.Println(" strings.NewReader(string(byteData)):", strings.NewReader(string(byteData)))
	req, err := http.NewRequest(request.Method, request.Host, strings.NewReader(string(params)))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", Json)
	req.Header.Set("req-id", utils.GenerateNonce())
	if request.Headers != nil {
		for k, v := range request.Headers {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req) //
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (request *baseRequest) Get() ([]byte, error) {
	params := request.GetQueryParams()
	byteData, _ := json.Marshal(params)
	fmt.Println(" strings.NewReader(string(byteData)):", strings.NewReader(string(byteData)))
	req, err := http.NewRequest(request.Method, request.Host, strings.NewReader(string(byteData)))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", Json)
	if request.Headers != nil {
		for k, v := range request.Headers {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req) //
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (request *baseRequest) SetContentType(contentType string) *baseRequest {
	if contentType == "" {
		contentType = Json
	}
	request.addHeaderParam("Content-Type", contentType)
	return request
}

func (request *baseRequest) GetContentType() (contentType string, contains bool) {
	contentType, contains = request.Headers["Content-Type"]
	return
}

func (request *baseRequest) SetHost(host string) *baseRequest {
	request.Host = host
	return request
}

func (request *baseRequest) GetHost() string {
	return request.Host
}

func (request *baseRequest) SetMethod(method string) *baseRequest {
	request.Method = method
	return request
}

func (request *baseRequest) GetMethod() string {
	return request.Method
}

func (request *baseRequest) SetContent(content []byte) *baseRequest {
	request.Content = content
	return request
}

func (request *baseRequest) GetContent() []byte {
	return request.Content
}

func (request *baseRequest) SetVersion(version string) *baseRequest {
	request.version = version
	return request
}

func (request *baseRequest) GetVersion() string {
	return request.version
}

func (request *baseRequest) GetActionName() string {
	return request.actionName
}

func (request *baseRequest) GetQueryParams() map[string]string {
	return request.QueryParams
}

func (request *baseRequest) GetFormParams() map[string]string {
	return request.FormParams
}

func (request *baseRequest) GetUserAgent() map[string]string {
	return request.userAgent
}

func (request *baseRequest) GetHeaders() map[string]string {
	return request.Headers
}

func (request *baseRequest) AppendUserAgent(key, value string) {
	newkey := true
	if request.userAgent == nil {
		request.userAgent = make(map[string]string)
	}
	if strings.ToLower(key) != "core" && strings.ToLower(key) != "go" {
		for tag, _ := range request.userAgent {
			if tag == key {
				request.userAgent[tag] = value
				newkey = false
			}
		}
		if newkey {
			request.userAgent[key] = value
		}
	}
}

func (request *baseRequest) addHeaderParam(key, value string) {
	request.Headers[key] = value
}

func (request *baseRequest) addQueryParam(key, value string) {
	request.QueryParams[key] = value
}

func (request *baseRequest) addFormParam(key, value string) {
	request.FormParams[key] = value
}

func (request *baseRequest) GetDomain() string {
	return request.Domain
}

func (request *baseRequest) SetDomain(host string) *baseRequest {
	request.Domain = host
	return request
}

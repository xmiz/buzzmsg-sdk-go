/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sdk

import (
	"fmt"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/auth"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/auth/credentials"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/requests"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/responses"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/utils"
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var debug utils.Debug

func init() {
	debug = utils.Init("sdk")
}

// Version this value will be replaced while build: -ldflags="-X sdk.version=x.x.x"
var Version = "0.0.1"
var defaultConnectTimeout = 5 * time.Second
var defaultReadTimeout = 10 * time.Second

var DefaultUserAgent = fmt.Sprintf("Buzzmsg (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)

var hookDo = func(fn func(req *http.Request) (*http.Response, error)) func(req *http.Request) (*http.Response, error) {
	return fn
}

// Client the type Client
type Client struct {
	SourceIp        string
	SecureTransport string
	isInsecure      bool
	regionId        string
	config          *Config
	httpProxy       string
	httpsProxy      string
	noProxy         string
	//logger          *Logger
	userAgent      map[string]string
	signer         auth.Signer
	httpClient     *http.Client
	asyncTaskQueue chan func()
	readTimeout    time.Duration
	connectTimeout time.Duration
	EndpointMap    map[string]string
	EndpointType   string
	Network        string
	Domain         string
	isOpenAsync    bool
}

func (client *Client) Init() (err error) {
	panic("not support yet")
}

func (client *Client) InitWithOptions(regionId string, config *Config, credential auth.Credential) (err error) {
	if regionId != "" {
		match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", regionId)
		if !match {
			return fmt.Errorf("regionId contains invalid characters")
		}
	}

	client.regionId = regionId
	client.config = config
	client.httpClient = &http.Client{}

	if config.Transport != nil {
		client.httpClient.Transport = config.Transport
	} else if config.HttpTransport != nil {
		client.httpClient.Transport = config.HttpTransport
	}

	if config.Timeout > 0 {
		client.httpClient.Timeout = config.Timeout
	}

	client.signer, err = auth.NewSignerWithCredential(credential)

	return
}

func (client *Client) InitWithAccessKey(regionId, accessKeyId, accessKeySecret string) (err error) {
	config := client.InitClientConfig()
	credential := &credentials.AccessKeyCredential{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	return client.InitWithOptions(regionId, config, credential)
}

func (client *Client) InitWithStsToken(regionId, accessKeyId, accessKeySecret, securityToken string) (err error) {
	config := client.InitClientConfig()
	credential := &credentials.StsTokenCredential{
		AccessKeyId:       accessKeyId,
		AccessKeySecret:   accessKeySecret,
		AccessKeyStsToken: securityToken,
	}
	return client.InitWithOptions(regionId, config, credential)
}

func (client *Client) InitClientConfig() (config *Config) {
	if client.config != nil {
		return client.config
	} else {
		return NewConfig()
	}
}

func (client *Client) DoAction(request requests.Request, response responses.Response) (err error) {
	if (client.SecureTransport == "false" || client.SecureTransport == "true") && client.SourceIp != "" {
		t := reflect.TypeOf(request).Elem()
		v := reflect.ValueOf(request).Elem()
		for i := 0; i < t.NumField(); i++ {
			value := v.FieldByName(t.Field(i).Name)
			if t.Field(i).Name == "requests.RoaRequest" || t.Field(i).Name == "RoaRequest" {
				request.GetHeaders()["x-acs-proxy-source-ip"] = client.SourceIp
				request.GetHeaders()["x-acs-proxy-secure-transport"] = client.SecureTransport
				//return client.DoActionWithSigner(request, response, nil)
			} else if t.Field(i).Name == "PathPattern" && !value.IsZero() {
				request.GetHeaders()["x-acs-proxy-source-ip"] = client.SourceIp
				request.GetHeaders()["x-acs-proxy-secure-transport"] = client.SecureTransport
				//return client.DoActionWithSigner(request, response, nil)
			} else if i == t.NumField()-1 {
				request.GetQueryParams()["SourceIp"] = client.SourceIp
				request.GetQueryParams()["SecureTransport"] = client.SecureTransport
				//return client.DoActionWithSigner(request, response, nil)
			}
		}
	}
	//return client.DoActionWithSigner(request, response, nil)
	return nil
}

func (client *Client) ProcessCommonRequestWithSigner(request *requests.CommonRequest, signerInterface interface{}) (response *responses.CommonResponse, err error) {
	if signer, isSigner := signerInterface.(auth.Signer); isSigner {
		//request.TransToAcsRequest()
		response = responses.NewCommonResponse()
		//err = client.DoActionWithSigner(request, response, signer)
		fmt.Println("signer:", signer)
		return
	}
	panic("should not be here")
}

func (client *Client) DoActionWithSigner(request requests.CommonRequest, response responses.Response, signer auth.Signer) (err error) {
	if client.Network != "" {
		match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", client.Network)
		if !match {
			return fmt.Errorf("netWork contains invalid characters")
		}
	}
	//fieldMap := make(map[string]string)
	////initLogMsg(fieldMap)
	////defer func() {
	////	client.printLog(fieldMap, err)
	////}()
	//httpRequest, err := client.buildRequestWithSigner(request, signer)
	//if err != nil {
	//	return
	//}
	//
	//client.setTimeout(request)
	//proxy, err := client.getHttpProxy(httpRequest.URL.Scheme)
	//if err != nil {
	//	return err
	//}
	//
	//noProxy := client.getNoProxy(httpRequest.URL.Scheme)
	//
	//var flag bool
	//for _, value := range noProxy {
	//	if strings.HasPrefix(value, "*") {
	//		value = fmt.Sprintf(".%s", value)
	//	}
	//	noProxyReg, err := regexp.Compile(value)
	//	if err != nil {
	//		return err
	//	}
	//	if noProxyReg.MatchString(httpRequest.Host) {
	//		flag = true
	//		break
	//	}
	//}
	//
	//// Set whether to ignore certificate validation.
	//// Default InsecureSkipVerify is false.
	//if trans, ok := client.httpClient.Transport.(*http.Transport); ok && trans != nil {
	//	if trans.TLSClientConfig != nil {
	//		trans.TLSClientConfig.InsecureSkipVerify = client.getHTTPSInsecure(request)
	//	} else {
	//		trans.TLSClientConfig = &tls.Config{
	//			InsecureSkipVerify: client.getHTTPSInsecure(request),
	//		}
	//	}
	//	if proxy != nil && !flag {
	//		trans.Proxy = http.ProxyURL(proxy)
	//	}
	//	client.httpClient.Transport = trans
	//}
	//
	//var httpResponse *http.Response
	//for retryTimes := 0; retryTimes <= client.config.MaxRetryTime; retryTimes++ {
	//	if retryTimes > 0 {
	//		//client.printLog(fieldMap, err)
	//		//initLogMsg(fieldMap)
	//	}
	//	putMsgToMap(fieldMap, httpRequest)
	//	debug("> %s %s %s", httpRequest.Method, httpRequest.URL.RequestURI(), httpRequest.Proto)
	//	debug("> Host: %s", httpRequest.Host)
	//	for key, value := range httpRequest.Header {
	//		debug("> %s: %v", key, strings.Join(value, ""))
	//	}
	//	debug(">")
	//	debug(" Retry Times: %d.", retryTimes)
	//
	//	startTime := time.Now()
	//	fieldMap["{start_time}"] = startTime.Format("2006-01-02 15:04:05")
	//	httpResponse, err = hookDo(client.httpClient.Do)(httpRequest)
	//	fieldMap["{cost}"] = time.Since(startTime).String()
	//	if err == nil {
	//		fieldMap["{code}"] = strconv.Itoa(httpResponse.StatusCode)
	//		fieldMap["{res_headers}"] = TransToString(httpResponse.Header)
	//		debug("< %s %s", httpResponse.Proto, httpResponse.Status)
	//		for key, value := range httpResponse.Header {
	//			debug("< %s: %v", key, strings.Join(value, ""))
	//		}
	//	}
	//	debug("<")
	//	// receive error
	//	if err != nil {
	//		debug(" Error: %s.", err.Error())
	//		if !client.config.AutoRetry {
	//			return
	//		} else if retryTimes >= client.config.MaxRetryTime {
	//			// timeout but reached the max retry times, return
	//			times := strconv.Itoa(retryTimes + 1)
	//			timeoutErrorMsg := fmt.Sprintf(errors.TimeoutErrorMessage, times, times)
	//			if strings.Contains(err.Error(), "Client.Timeout") {
	//				timeoutErrorMsg += " Read timeout. Please set a valid ReadTimeout."
	//			} else {
	//				timeoutErrorMsg += " Connect timeout. Please set a valid ConnectTimeout."
	//			}
	//			err = errors.NewClientError(errors.TimeoutErrorCode, timeoutErrorMsg, err)
	//			return
	//		}
	//	}
	//	if isCertificateError(err) {
	//		return
	//	}
	//
	//	//  if status code >= 500 or timeout, will trigger retry
	//	if client.config.AutoRetry && (err != nil || isServerError(httpResponse)) {
	//		client.setTimeout(request)
	//		// rewrite signatureNonce and signature
	//		httpRequest, err = client.buildRequestWithSigner(request, signer)
	//		// buildHttpRequest(request, finalSigner, regionId)
	//		if err != nil {
	//			return
	//		}
	//		continue
	//	}
	//	break
	//}
	//
	//err = responses.Unmarshal(response, httpResponse, request.GetAcceptFormat())
	//fieldMap["{res_body}"] = response.GetHttpContentString()
	//debug("%s", response.GetHttpContentString())
	//// wrap server errors
	//if serverErr, ok := err.(*errors.ServerError); ok {
	//	var wrapInfo = map[string]string{}
	//	serverErr.RespHeaders = response.GetHttpHeaders()
	//	wrapInfo["StringToSign"] = request.GetStringToSign()
	//	err = errors.WrapServerError(serverErr, wrapInfo)
	//}
	return
}


func isCertificateError(err error) bool {
	if err != nil && strings.Contains(err.Error(), " certificate signed by unknown authority") {
		return true
	}
	return false
}

func buildHttpRequest(request requests.Request, singer auth.Signer, regionId string) (httpRequest *http.Request, err error) {
	err = auth.Sign(request, singer, regionId)
	if err != nil {
		return
	}
	requestMethod := request.GetMethod()
	requestUrl := request.BuildUrl()
	body := request.GetBodyReader()
	httpRequest, err = http.NewRequest(requestMethod, requestUrl, body)
	if err != nil {
		return
	}
	for key, value := range request.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}
	// host is a special case
	if host, containsHost := request.GetHeaders()["Host"]; containsHost {
		httpRequest.Host = host
	}
	return
}

func isServerError(httpResponse *http.Response) bool {
	return httpResponse.StatusCode >= http.StatusInternalServerError
}


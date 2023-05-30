package auth

import (
	"fmt"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/auth/credentials"
	signers "github.com/xmiz/buzzmsg-sdk-go/sdk/auth/singers"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/errors"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/requests"
	"reflect"
)

type Signer interface {
	GetName() string
	GetType() string
	GetVersion() string
	GetAccessKeyId() (string, error)
	GetExtraParam() map[string]string
	Sign(stringToSign, secretSuffix string) string
}

func NewSignerWithCredential(credential Credential) (signer Signer, err error) {
	switch instance := credential.(type) {
	case *credentials.AccessKeyCredential:
		{
			signer = signers.NewAccessKeySigner(instance)
		}
	default:
		message := fmt.Sprintf(errors.UnsupportedCredentialErrorMessage, reflect.TypeOf(credential))
		err = errors.NewClientError(errors.UnsupportedCredentialErrorCode, message, nil)
	}
	return
}

func Sign(request requests.Request, signer Signer, regionId string) (err error) {
	switch request.GetStyle() {
	case requests.HTTP:
		{
			err = signHttpRequest(request, signer, regionId)
		}
	default:
		message := fmt.Sprintf(errors.UnknownRequestTypeErrorMessage, reflect.TypeOf(request))
		err = errors.NewClientError(errors.UnknownRequestTypeErrorCode, message, nil)
	}

	return
}

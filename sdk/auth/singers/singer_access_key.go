package signers

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/auth/credentials"
)

type AccessKeySigner struct {
	credential *credentials.AccessKeyCredential
}

func (signer *AccessKeySigner) GetExtraParam() map[string]string {
	return nil
}

func NewAccessKeySigner(credential *credentials.AccessKeyCredential) *AccessKeySigner {
	return &AccessKeySigner{
		credential: credential,
	}
}

func (*AccessKeySigner) GetName() string {
	return "HMAC-SHA1"
}

func (*AccessKeySigner) GetType() string {
	return ""
}

func (*AccessKeySigner) GetVersion() string {
	return "1.0"
}

func (signer *AccessKeySigner) GetAccessKeyId() (accessKeyId string, err error) {
	return signer.credential.AccessKeyId, nil
}

func (signer *AccessKeySigner) Sign(stringToSign, secretSuffix string) string {
	secret := signer.credential.AccessKeySecret
	sign, _ := SignByPrivateKeyStr([]byte(stringToSign), secret)
	return base58.Encode(sign)
}

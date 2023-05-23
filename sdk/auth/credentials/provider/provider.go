package provider

import (
	"github.com/xmiz/buzzmsg-sdk-go/sdk/auth"
)

//Environmental virables that may be used by the provider
const (
	ENVAccessKeyID     = "ABUZZMSG_ACCESS_KEY_ID"
	ENVAccessKeySecret = "ABUZZMSG_ACCESS_KEY_SECRET"
	ENVCredentialFile  = "ABUZZMSG_CREDENTIALS_FILE"
	ENVEcsMetadata     = "ABUZZMSG_ECS_METADATA"
)

// When you want to customize the provider, you only need to implement the method of the interface.
type Provider interface {
	Resolve() (auth.Credential, error)
}

package chat

import (
	"github.com/xmiz/buzzmsg-sdk-go/sdk"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/auth"

	//"github.com/xmiz/buzzmsg-sdk-go/sdk/auth"
)

// Client is the sdk client struct
type Client struct {
	sdk.Client
}

// NewClientWithOptions creates a sdk client with regionId/sdkConfig/credential
// this is the common api to create a sdk client
func NewClientWithOptions(regionId string, config *sdk.Config, credential auth.Credential) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithOptions(regionId, config, credential)
	return
}

// NewClientWithAccessKey is a shortcut to create sdk client with accesskey
func NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(regionId, accessKeyId, accessKeySecret)
	return
}

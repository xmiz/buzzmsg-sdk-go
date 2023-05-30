package main

import (
	"fmt"
	"github.com/xmiz/buzzmsg-sdk-go/sdk"
	"github.com/xmiz/buzzmsg-sdk-go/sdk/auth/credentials"
	"github.com/xmiz/buzzmsg-sdk-go/services/chat"
)

func main() {
		config := sdk.NewConfig().WithDebug(true)
		accessKeyId := "TMMTMM"
		accessKeySecret := "51coYQESAvfJaRdXvbei3GkfZL2D9YXvQwsE3suycVZj"
		Credential := credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
		client, _ := chat.NewClientWithOptions("ap-southeast-1", config, Credential)
		res, err := client.Chat(chat.CreateChatRequest())
		fmt.Println(res, err)
}

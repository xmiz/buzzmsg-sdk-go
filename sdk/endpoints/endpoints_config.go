package endpoints

import (
	"encoding/json"
	"fmt"
	"sync"
)

var initOnce sync.Once
var data interface{}

const endpointsJson = `{
	"products":[
			{
				"code":"chat",
				"location_service_code":"chat",
				"endpoints":[
					{
						"region":"ap-southeast-1",
						"endpoint":"172.17.0.1:7500",
					}
				],
				"global_endpoint":"172.17.0.1:7500"
			},
			{
				"code":"message",
				"location_service_code":"message",
				"endpoints":[
					{
						"region":"ap-southeast-1",
						"endpoint":"172.17.0.1:7500",
					}
				],
				"global_endpoint":"172.17.0.1:7500"
			},
			{
				"code":"command",
				"location_service_code":"command",
				"endpoints":[
					{
						"region":"ap-southeast-1",
						"endpoint":"172.17.0.1:7500",
					}
				],
				"global_endpoint":"172.17.0.1:7500"
			},
	]
}`

func getEndpointConfigData() interface{} {
	initOnce.Do(func() {
		err := json.Unmarshal([]byte(endpointsJson), &data)
		if err != nil {
			panic(fmt.Sprintf("init endpoint config data failed. %s", err))
		}
	})
	return data
}

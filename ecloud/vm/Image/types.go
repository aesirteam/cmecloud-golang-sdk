package Image

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type ImageInterface interface {
}

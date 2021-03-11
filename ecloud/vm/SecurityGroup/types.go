package SecurityGroup

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func NewForConfig(conf *global.Config) (*APIv2, error) {
	return &APIv2{client: global.New(conf)}, nil
}

type SecurityGroupInterface interface {
}

package SecurityGroup

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type SecurityGroupInterface interface {
	//创建安全组
	CreateSecurityGroup(name, desc string) (SecurityGroupResult, error)

	//删除安全组
	DeleteSecurityGroup(id string) error

	//查询安全组列表
	GetSecurityGroupList(name string, page, size int) ([]SecurityGroupResult, error)

	//安全组
	ModifySecurityGroup(id, name, desc string) (SecurityGroupResult, error)

	//
	AddSecurityGroupRules()

	//
	DeleteSecurityGroupRules()

	//
	GetSecurityGroupRules()
}

type SecurityGroupResult struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"description"`
	CreatedTime  string `json:"createdTime"`
	ModifiedTime string `json:"modifiedTime,omitempty"`
}

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
	GetSecurityGroupList(name string, ruleShow bool, page, size int) ([]SecurityGroupResult, error)

	//修改安全组名称及备注
	ModifySecurityGroup(id, name, desc string) (SecurityGroupResult, error)

	//创建安全组规则
	AddSecurityGroupRules(sg *global.SecurityGroupRuleSpec) (SecurityGroupRuleResult, error)

	//查询安全组规则
	GetSecurityGroupRules(securityGroupId, ruleId string) ([]SecurityGroupRuleResult, error)

	//删除安全组规则
	DeleteSecurityGroupRules(ruleId string) error
}

type SecurityGroupResult struct {
	Id                 string                    `json:"id"`
	Name               string                    `json:"name"`
	Desc               string                    `json:"description,omitempty"`
	CreatedTime        string                    `json:"createdTime"`
	ModifiedTime       string                    `json:"modifiedTime,omitempty"`
	SecurityGroupRules []SecurityGroupRuleResult `json:"securityGroupRules,omitempty"`
}

type SecurityGroupRuleResult struct {
	Id              string `json:"id"`
	SecurityGroupId string `json:"securityGroupId"`
	EtherType       string `json:"ethertype"`
	Protocol        string `json:"protocol,omitempty"`
	PortRangeMin    int    `json:"portRangeMin,omitempty"`
	PortRangeMax    int    `json:"portRangeMax,omitempty"`
	Direction       string `json:"direction"`
	RemoteGroupId   string `json:"remoteGroupId,omitempty"`
	RemoteIpPrefix  string `json:"remoteIpPrefix,omitempty"`
	CreatedTime     string `json:"createdTime"`
	ModifiedTime    string `json:"modifiedTime,omitempty"`
}

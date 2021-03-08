package VirtualPrivateCloud

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
)

type APIv2 struct {
	client *core.EcloudClient
}

func NewForConfig(conf *core.Config) (*APIv2, error) {
	return &APIv2{client: core.New(conf)}, nil
}

type Interface interface {
	/*
		查看VPC列表
		  natGatewayBind	VPC是否绑定了NAT网关
		  visible		是否可见
		  scale			VPC规格
		  region		可用区
		  tagIds		标签ID列表
	*/
	GetVPCList(natGatewayBind, visible bool, scale VpcScale, region string, tagIds []string, page, size int) (VpcResultArray, error)

	//根据vpcId查看VPC详情，包含路由详情
	GetVPCInfo(vpcId string) (VpcResult, error)

	//根据vpcId查看VPC详情，包含路由详情
	GetVPCInfoByName(name string) (VpcResult, error)
}

type VpcResult struct {
	Id           string `json: "id"`
	Name         string `json:"name"`
	Desc         string `json:"description,omitempty"`
	CreateTime   string `json:"createdTime"`
	RouteId      string `json:"routerId"`
	Deleted      bool   `json:"deleted"`
	Scale        string `json:"scale"`
	EcStatus     string `json:"ecStatus"`
	AdminStateUp bool   `json:"adminStateUp"`
	NetworkId    string `json:"firstNetworkId"`
	UserId       string `json:"userId"`
	UserName     string `json:"userName,omitempty"`
	Region       string `json:"region"`
}
type VpcResultArray []VpcResult

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
	//创建VPC
	CreatVpc(spec VpcSpec) (string, error)

	//退订VPC
	DeleteVpc(vpcId string)

	/*
		查看VPC列表
			natGatewayBind	VPC是否绑定了NAT网关
			visible		是否可见
			scale			VPC规格
			region		可用区
			tagIds		标签ID列表
	*/
	GetVpcList(natGatewayBind, visible bool, scale VpcScale, region string, tagIds []string, page, size int) (VpcResultArray, error)

	//根据vpcId查看VPC详情，包含路由详情
	GetVpcInfo(vpcId string) (VpcResult, error)

	//根据name查看VPC详情，包含路由详情
	GetVpcInfoByName(name string) (VpcResult, error)

	//根据routerId查看VPC详情
	GetVpcInfoByRouterId(id string)

	//更新VPC名称和描述信息
	ModifyVpcInfo(vpcId, name, desc string)

	//根据routerId查询VPC下防火墙
	GetVpcFirewall(routerId string)

	//根据routerId查询VPC下的网络
	GetVpcNetwork(routerId string)

	//根据routerId查询VPC下IPSecVPN
	GetVpcVPN(routerId string)

	//根据routerId查询VPC下虚拟网卡
	GetVpcNIC(routerId string)

	//创建子网
	CreateSubnet(spec VpcSpec) (string, error)

	//删除子网
	DeleteSubnet(networkId string) error

	//修改子网名称
	ModifySubnet(networkId, name string)

	//根据subnetId查询子网详情
	GetSubnetInfo(subnetId string)

	//创建虚拟网卡
	CreateNic()

	//删除虚拟网卡
	DeleteNic(portId string)

	//查询虚拟网卡详情
	GetNicDetail(portId string)

	//修改虚拟网卡名称
	ModifyNicName(id, name string)

	//修改虚拟网卡安全组
	ModifyNicSecurityGroup(id, name string, securityGroupIds []string)
}

/*
虚拟私有云
  Name		【必填】vpc名称,路由器名称
  NetworkName	【必填】网络名称
  RouterExternal	【必填】路由器是否开启网关
  Region		【必填】可用区
  Specs		VPC规格
*/
type VpcSpec struct {
	Name           string
	NetworkName    string
	Cidr           string
	CidrV6         string
	RouterExternal bool
	Region         string
	Specs          VpcScale
	Subnets        []SubnetSpec

	RouterId string
}

/*
子网信息
  Cidr		【必填】子网cid
  IpVersion	【必填】enum(4,6)
  SubnetName	【必填】子网名称
*/
type SubnetSpec struct {
	Cidr       string
	IpVersion  int
	SubnetName string
}

type VpcResult struct {
	Id           string `json:"id"`
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

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
	CreateVpc(spec *VpcSpec) (string, error)

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
	GetVpcList(natGatewayBind, visible bool, scale VpcScale, region string, tagIds []string, page, size int) ([]VpcResult, error)

	//根据vpcId查看VPC详情，包含路由详情
	GetVpcInfo(vpcId string) (VpcResult, error)

	//根据name查看VPC详情，包含路由详情
	GetVpcInfoByName(name string) (VpcResult, error)

	//根据routerId查看VPC详情
	GetVpcInfoByRouterId(routerId string)

	//更新VPC名称和描述信息
	ModifyVpcInfo(vpcId, name, desc string)

	//根据routerId查询VPC下防火墙
	GetVpcFirewall(routerId string)

	//根据routerId查询VPC下的网络
	GetVpcNetwork(routerId string) ([]VpcNetResult, error)

	//根据routerId查询VPC下IPSecVPN
	GetVpcVPN(routerId string)

	//根据routerId查询VPC下虚拟网卡
	GetVpcNIC(routerId string)

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

	//查询路由所关联的子网列表
	GetRouterNetList(routerId string) ([]RouterNetResult, error)

	//查看路由器详情
	GetRouterInfo(routerId string) (RouterResult, error)

	//创建子网
	CreateSubnet(spec *SubnetSpec) (string, error)

	//删除子网
	DeleteSubnet(networkId string) error

	//修改子网名称
	ModifySubnet(networkId, name string) error

	//根据subnetId查询子网详情
	GetSubnetInfo(networkId string) ([]SubnetResult, error)
}

/*
虚拟私有云
  Cidr	【必填】cidr
  Name		【必填】vpc名称,路由器名称
  NetworkName	【必填】网络名称
  RouterExternal	【必填】路由器是否开启网关
  Region		【必填】可用区
  Specs		VPC规格
*/
type VpcSpec struct {
	Cidr        string
	CidrV6      string
	Name        string
	NetworkName string
	Region      string
	Specs       VpcScale
}

/*
子网信息
  Cidr		【必填】子网cid
  IpVersion	【必填】enum(4,6)
  SubnetName	【必填】子网名称
*/
type SubnetSpec struct {
	RouterId    string
	NetworkName string
	Region      string
	Subnets     []struct {
		Cidr       string
		IpVersion  int
		SubnetName string
	}
}

type VpcResult struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"description,omitempty"`
	CreateTime   string `json:"createdTime"`
	RouterId     string `json:"routerId"`
	Deleted      bool   `json:"deleted"`
	Scale        string `json:"scale"`
	EcStatus     string `json:"ecStatus"`
	AdminStateUp bool   `json:"adminStateUp"`
	NetworkId    string `json:"firstNetworkId"`
	UserId       string `json:"userId"`
	UserName     string `json:"userName,omitempty"`
	Region       string `json:"region"`
}

type VpcNetResult struct {
	Id              string         `json:"id"`
	Name            string         `json:"name"`
	Shared          bool           `json:"shared"`
	Enabled         bool           `json:"enabled"`
	EcStatus        string         `json:"ecStatus"`
	CreateTime      string         `json:"createdTime"`
	Deleted         bool           `json:"deleted"`
	OrderSource     string         `json:"orderSource"`
	Subnets         []SubnetResult `json:"subnets"`
	NetworkTypeEnum string         `json:"networkTypeEnum"`
	Region          string         `json:"region"`
}

type RouterResult struct {
	Id                  string              `json:"id"`
	Name                string              `json:"name"`
	CreateTime          string              `json:"createdTime"`
	Deleted             bool                `json:"deleted"`
	EcStatus            string              `json:"ecStatus"`
	AdminStateUp        bool                `json:"adminStateUp"`
	ExternalGatewayInfo interface{}         `json:"externalGatewayInfo,omitempty"`
	ExternalNet         string              `json:"externalNet,omitempty"`
	SubnetIds           []string            `json:"subnetIds"`
	Routes              []map[string]string `json:"routes,omitempty"`
	Region              string              `json:"region"`
}

type RouterNetResult struct {
	RouterId   string `json:"routerId"`
	SubnetId   string `json:"subnetId"`
	SubnetName string `json:"subnetName"`
	Cidr       string `json:"cidr"`
	GatewayIp  string `json:"gatewayIP"`
	PortId     string `json:"portId"`
}

type SubnetResult struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	NetworkId  string   `json:"networkId"`
	Cidr       string   `json:"cidr"`
	IpVersion  int      `json:"ipVersion"`
	GatewayIP  string   `json:"gatewayIp"`
	CreateTime string   `json:"createdTime"`
	Deleted    bool     `json:"deleted"`
	EnableDHCP bool     `json:"enableDHCP,omitempty"`
	DnsNames   []string `json:"dnsNames,omitempty"`
	Pools      []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"pools,omitempty"`
	NetworkTypeEnum string `json:"networkTypeEnum,omitempty"`
	Region          string `json:"region"`
}

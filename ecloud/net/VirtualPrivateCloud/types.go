package VirtualPrivateCloud

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Server"
)

type APIv2 struct {
	client *global.EcloudClient
}

func NewForConfig(conf *global.Config) (*APIv2, error) {
	return &APIv2{client: global.New(conf)}, nil
}

type VPCInterface interface {
	//创建VPC
	CreateVpc(spec *VpcSpec) (string, error)

	//退订VPC
	DeleteVpc(vpcId string)

	/*
		查看VPC列表
			natGatewayBind	VPC是否绑定了NAT网关
			scale			VPC规格
			region		可用区
			tagIds		标签ID列表
	*/
	GetVpcList(natGatewayBind bool, scale VpcScale, region string, tagIds []string, page, size int) ([]VpcResult, error)

	//根据vpcId查看VPC详情，包含路由详情
	GetVpcInfo(vpcId string) (VpcResult, error)

	//根据name查看VPC详情，包含路由详情
	GetVpcInfoByName(name string) (VpcResult, error)

	//根据routerId查看VPC详情
	GetVpcInfoByRouterId(routerId string) (VpcResult, error)

	//更新VPC名称和描述信息
	ModifyVpcInfo(vpcId, name, desc string) error

	//根据routerId查询VPC下防火墙
	GetVpcFirewall(routerId string) (string, error)

	//根据routerId查询VPC下的网络
	GetVpcNetwork(routerId string) ([]VpcNetResult, error)

	//根据routerId查询VPC下IPSecVPN
	GetVpcVPN(routerId string) ([]VpnResult, error)

	//根据routerId查询VPC下虚拟网卡
	GetVpcNic(routerId string) ([]NicResult, error)

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

	//查询路由器所关联的子网列表
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

	//根据routerId绑定云防火墙
	//FirewallBindRouter(firewallId, routerId string) (result string, err error)

	//查看IPSecVPN列表
	GetIpsecVpnList(vpcName, routerId string, scale VpcScale, region string, page, size int) ([]VpnResult, error)

	//根据vpnId查看IPSecVPN服务的详情
	GetIpsecVpnInfo(vpnId string) (result VpnResult, err error)
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

type NicResult struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Status         int    `json:"status"`
	AdminStateUp   int    `json:"adminStateUp"`
	MacAddress     string `json:"macAddress"`
	NetworkId      string `json:"networkId"`
	ResourceId     string `json:"resourceId,omitempty"`
	Proposer       string `json:"proposer,omitempty"`
	CustomerId     string `json:"customerId"`
	PoolId         string `json:"poolId,omitempty"`
	CreatedBy      string `json:"createdBy"`
	CreateTime     string `json:"createdTime"`
	IsBasic        int    `json:"isBasic"`
	Source         int    `json:"source"`
	HostName       string `json:"hostName,omitempty"`
	SecurityGroups []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"securityGroups,omitempty"`
	FixedIps      Server.ServerFixedIpDetailArray `json:"fixedIps,omitempty"`
	IpId          string                          `json:"ipId,omitempty"`
	PublicIp      string                          `json:"publicIp,omitempty"`
	BandWidthsize string                          `json:"bandWidthsize,omitempty"`
}

type VpnResult struct {
	Id                 string         `json:"id"`
	Name               string         `json:"name"`
	Desc               string         `json:"description,omitempty"`
	RouterId           string         `json:"routerId"`
	EcStatus           string         `json:"ecStatus"`
	AdminStateUp       bool           `json:"adminStateUp"`
	FloatingIpId       string         `json:"floatingipId,omitempty"`
	Region             string         `json:"region"`
	MaxSiteConnections string         `json:"maxSiteConnections,omitempty"`
	LocalGateway       string         `json:"localGateway"`
	BandwidthSize      int            `json:"bandwidthSize"`
	BandwidthType      string         `json:"bandwidthType"`
	CreateTime         string         `json:"createdTime"`
	Deleted            bool           `json:"deleted"`
	Visible            bool           `json:"visible"`
	SubnetIds          []string       `json:"subnetIds"`
	SubnetResps        []SubnetResult `json:"subnetResps,omitempty"`
	SiteConnIds        []string       `json:"siteConnIds"`
}

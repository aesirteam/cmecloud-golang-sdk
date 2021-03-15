package VirtualPrivateCloud

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/IPSecVpn"
)

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type VPCInterface interface {
	//创建VPC
	CreateVpc(spec *global.VpcSpec) (string, error)

	//退订VPC
	DeleteVpc(vpcId string) error

	/*
		查看VPC列表
			natGatewayBind	VPC是否绑定了NAT网关
			scale			VPC规格
			region		可用区
			tagIds		标签ID列表
	*/
	GetVpcList(natGatewayBind bool, scale global.VpcScale, region string, tagIds []string, page, size int) ([]VpcResult, error)

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
	GetVpcVPN(routerId string) ([]IPSecVpn.VpnResult, error)

	//根据routerId查询VPC下虚拟网卡
	GetVpcNic(routerId string) ([]NicResult, error)

	//创建虚拟网卡
	CreateNic(spec *global.NicSpec) (string, error)

	//删除虚拟网卡
	DeleteNic(portId string) error

	//查询虚拟网卡详情
	GetNicDetail(portId string) (NicResult, error)

	//修改虚拟网卡名称
	ModifyNicName(portId, portName string) error

	//修改虚拟网卡安全组
	ModifyNicSecurityGroup(portId string, securityGroupIds []string) error

	//查询路由器所关联的子网列表
	GetRouterNetList(routerId string) ([]RouterNetResult, error)

	//查看路由器详情
	GetRouterInfo(routerId string) (RouterResult, error)

	//创建子网
	CreateSubnet(spec *global.SubnetSpec) (string, error)

	//删除子网
	DeleteSubnet(networkId string) error

	//修改子网名称
	ModifySubnet(networkId, networkName string) error

	//根据subnetId查询子网详情
	GetSubnetInfo(networkId string) ([]SubnetResult, error)
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
	Id              string                `json:"id"`
	Name            string                `json:"name"`
	Status          int                   `json:"status,omitempty"`
	AdminStateUp    interface{}           `json:"adminStateUp"`
	MacAddress      string                `json:"macAddress"`
	NetworkId       string                `json:"networkId,omitempty"`
	ResourceId      string                `json:"resourceId,omitempty"`
	ResourceName    string                `json:"resourceName,omitempty"`
	FipBind         bool                  `json:"fipBind,omitempty"`
	OperationStatus string                `json:"operationStatus,omitempty"`
	Proposer        string                `json:"proposer,omitempty"`
	CustomerId      string                `json:"customerId,omitempty"`
	PoolId          string                `json:"poolId,omitempty"`
	CreatedBy       string                `json:"createdBy,omitempty"`
	CreateTime      string                `json:"createdTime,omitempty"`
	IsBasic         interface{}           `json:"isBasic" newtag:"basic"`
	Source          interface{}           `json:"source"`
	AddressCheck    bool                  `json:"addressCheck,omitempty"`
	Region          string                `json:"region,omitempty"`
	HostName        string                `json:"hostName,omitempty"`
	SecurityGroups  []interface{}         `json:"securityGroups,omitempty" newtag:"sgIds"`
	FixedIps        []ServerFixedIpDetail `json:"fixedIps,omitempty" newtag:"fixedIpResps"`
	IpId            string                `json:"ipId,omitempty"`
	PublicIp        string                `json:"publicIp,omitempty"`
	BandWidthsize   string                `json:"bandWidthsize,omitempty"`
}

/*
绑定信息
  IpAddress		内网ip
  IpVersion		v4/v6
  PublicIP		公网ip
  BandwidthSize		公网ip带宽（Mb）
  BandwidthType		带宽类型
  Ipv6BandwidthSize	Ipv6带宽值（Mb）
  SubnetId		子网id
  SubnetName		子网名称
*/
type ServerFixedIpDetail struct {
	IpAddress         string `json:"ipAddress"`
	IpVersion         int    `json:"ipVersion"`
	PublicIP          string `json:"publicIP,omitempty"`
	BandwidthSize     int    `json:"bandwidthSize,omitempty"`
	BandwidthType     string `json:"bandwidthType,omitempty"`
	Ipv6BandwidthSize int    `json:"ipv6BandwidthSize,omitempty"`
	SubnetId          string `json:"subnetId"`
	SubnetName        string `json:"subnetName,omitempty"`
	SubnetCidr        string `json:"subnetCidr,omitempty"`
	PortId            string `json:"portId,omitempty"`
}

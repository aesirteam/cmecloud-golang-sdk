package IPSecVpn

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type VPNInterface interface {
	// 创建IPSecVPN
	CreateIpsecVpn(spec *global.IPSecVpnSpec) (string, error)

	// 查看IPSecVPN列表
	GetIpsecVpnList(vpcName, routerId string, scale global.VpcScale, region string, page, size int) ([]VpnResult, error)

	// 根据vpnId查看IPSecVPN服务的详情
	GetIpsecVpnInfo(vpnId string) (result VpnResult, err error)

	// 查看VPN配额情况
	GetIpsecVpnQuota() (VpnQuotaResult, error)

	// 删除IPSecVPN
	DeleteIpsecVpn(vpnId string) error

	// 创建VPN站点连接
	CreateIpsecVpnConnection(cs *global.IPSecVpnSiteConnectionSpec) error

	// 查看VPN站点连接列表
	GetIpsecVpnConnectionList(queryWord, name, networkId string, serviceIdsInRange []string, page, size int) ([]VpnConnectionResult, error)

	// 查看VPN站点连接详情
	GetIpsecVpnConnectionInfo(siteConnectionId string) (VpnConnectionResult, error)

	// 查看ikePolicy详情
	GetIkePolicyInfo(siteConnectionId, ikePolicyId string) (VpnConnectionPolicyResult, error)

	// 查看ipsecPolicy详情
	GetIpsecPolicyInfo(siteConnectionId, ipsecPolicyId string) (VpnConnectionPolicyResult, error)

	// 修改VPN站点连接信息
	ModifyIpsecVpnConnection(siteConnectionId string, cs *global.IPSecVpnSiteConnectionSpec) error

	// 删除VPN站点连接
	DeleteIpsecVpnConnection(siteConnectionId string) error
}

type VpnResult struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Desc               string   `json:"description,omitempty"`
	RouterId           string   `json:"routerId"`
	EcStatus           string   `json:"ecStatus"`
	AdminStateUp       bool     `json:"adminStateUp"`
	FloatingIpId       string   `json:"floatingipId,omitempty"`
	Region             string   `json:"region"`
	MaxSiteConnections string   `json:"maxSiteConnections,omitempty"`
	LocalGateway       string   `json:"localGateway"`
	BandwidthSize      int      `json:"bandwidthSize"`
	BandwidthType      string   `json:"bandwidthType"`
	CreateTime         string   `json:"createdTime"`
	Deleted            bool     `json:"deleted"`
	Visible            bool     `json:"visible"`
	SubnetIds          []string `json:"subnetIds"`
	//SubnetResps        []VirtualPrivateCloud.SubnetResult `json:"subnetResps,omitempty"`
	SubnetResps interface{} `json:"subnetResps,omitempty"`
	SiteConnIds []string    `json:"siteConnIds"`
}

type VpnQuotaResult struct {
	UsedQuota  int `json:"usedQuota"`
	TotalQuota int `json:"totalQuota"`
}

type VpnConnectionResult struct {
	Id                   string      `json:"id"`
	Name                 string      `json:"name"`
	Desc                 string      `json:"description,omitempty"`
	EcStatus             string      `json:"ecStatus"`
	AdminStateUp         bool        `json:"adminStateUp"`
	BandwithSize         int         `json:"bandwithSize,omitempty"`
	VpnServiceId         string      `json:"vpnserviceId"`
	IkePolicyId          string      `json:"ikepolicyId"`
	IpsecPolicyId        string      `json:"ipsecpolicyId"`
	PeerAddress          string      `json:"peerAddress"`
	PeerId               string      `json:"peerId,omitempty"`
	LocalId              string      `json:"localId,omitempty"`
	PeerCidrs            string      `json:"peerCidrs,omitempty"`
	AuthMode             string      `json:"authMode"`
	RouteMode            string      `json:"routeMode"`
	Mtu                  int         `json:"mtu"`
	Initiator            string      `json:"initiator"`
	Psk                  string      `json:"psk"`
	Dpd                  interface{} `json:"dpd,omitempty"`
	CreateTime           string      `json:"createdTime"`
	Deleted              bool        `json:"deleted"`
	Region               string      `json:"region"`
	LocalEndpointGroupId string      `json:"localEndpointGroupId"`
	PeerEndpointGroupId  string      `json:"peerEndpointGroupId"`
	LocalSubnets         []string    `json:"localSubnets"`
	PeerSubnets          []string    `json:"peerSubnets"`
}

type VpnConnectionPolicyResult struct {
	Id                    string      `json:"id"`
	ServiceName           string      `json:"name"`
	Desc                  string      `json:"description,omitempty"`
	Lifetime              interface{} `json:"lifetime"`
	Pfs                   string      `json:"pfs"`
	TenantId              string      `json:"tenant_id"`
	AuthAlgorithm         string      `json:"auth_algorithm"`
	EncryptionAlgorithm   string      `json:"encryption_algorithm"`
	Phase1NegotiationMode string      `json:"phase1_negotiation_mode,omitempty"`
	Version               string      `json:"ike_version,omitempty"`

	TransformProtocol string `json:"transform_protocol,omitempty"`
	EncapsulationMode string `json:"encapsulation_mode,omitempty"`
}

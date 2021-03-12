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
	//查看IPSecVPN列表
	GetIpsecVpnList(vpcName, routerId string, scale global.VpcScale, region string, page, size int) ([]VpnResult, error)

	//根据vpnId查看IPSecVPN服务的详情
	GetIpsecVpnInfo(vpnId string) (result VpnResult, err error)
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

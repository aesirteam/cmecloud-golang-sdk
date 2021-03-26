package FloatingIP

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type FIPInterface interface {
	// 订购弹性公网IP和带宽
	CreateFloatingIp(fs *global.FloatingIpSpec) (FloatingIpOrderResult, error)

	/*
		查看IP带宽的列表
		  queryWord	查询关键字
		  routerId	共享带宽所在VPC的routerId
		  natGatewayId	Nat网关Id
		  bound	绑定或未绑定
		  frozen	冻结或未冻结
		  openIpv6Converter	是否开通IPv6转换
		  availableForSbw	公网IP是否可加入到共享带宽
		  tagIds	标签Id列表
	*/
	GetFloatingIpList(queryWord, routerId, natGatewayId string, bound, frozen, openIpv6Converter, availableForSbw bool, tagIds []string, page, size int) ([]FloatingIpResult, error)

	//  查看公网IP详情
	GetFloatingIpDetail(ipId string) (FloatingIpResult, error)

	// 查看IP带宽配额 (k版)
	GetFloatingIpQuotaWithK() (FloatingIpQuotaResult, error)

	// 查看IP带宽配额 (p版)
	GetFloatingIpQuotaWithP(quotaName global.FloatingIpQuotaName, resourceId string) (FloatingIpQuotaResult, error)

	/*
		弹性公网IP绑定资源
		  bindType	浮动ip绑定类型
		  ipId	浮动ip Id
		  resourceId	资源id
		  portId	虚拟网卡id (仅绑定类型为vm时需设置)
	*/
	AttachFloatingIp(bindType global.FloatingIpBindType, ipId, resourceId, portId string) error

	//  弹性公网IP解绑资源
	DetachFloatingIp(ipId string) error
}

type FloatingIpOrderResult struct {
	OrderId       string `json:"orderId"`
	OrderExtResps []struct {
		ProductType       string `json:"productType"`
		OrderExtId        string `json:"orderExtId"`
		OrderExtStatus    int    `json:"orderExtStatus"`
		RelProductType    string `json:"relProductType"`
		RelOrderExtId     string `json:"relOrderExtId"`
		RelOrderExtStatus int    `json:"relOrderExtStatus"`
	} `json:"orderExtResps"`
}

type FloatingIpResult struct {
	Id                         string `json:"id"`
	Ipv4                       string `json:"name"`
	Desc                       string `json:"description,omitempty"`
	Ipv6                       string `json:"ipv6,omitempty"`
	IpType                     string `json:"ipType"`
	ChargeModeEnum             string `json:"chargeModeEnum"`
	IcpStatus                  string `json:"icpStatus"`
	Status                     string `json:"status"`
	BandwidthSize              int    `json:"bandwidthSize"`
	BandwidthType              string `json:"bandwidthType"`
	BandwidthId                string `json:"bandwidthId"`
	BandwidthMbSize            int    `json:"bandwidthMbSize"`
	Visible                    bool   `json:"visible"`
	DummyFip                   string `json:"dummyFip,omitempty"`
	Region                     string `json:"region"`
	OrderVersion               string `json:"orderVersion,omitempty"`
	OrderType                  string `json:"orderType,omitempty"`
	BindType                   string `json:"bindType,omitempty"`
	ResourceId                 string `json:"resourceId,omitempty"`
	ResourceName               string `json:"resourceName,omitempty"`
	PortNetworkId              string `json:"portNetworkId,omitempty"`
	NicName                    string `json:"nicName,omitempty"`
	Frozen                     bool   `json:"frozen"`
	Bound                      bool   `json:"bound"`
	Occupy                     bool   `json:"occupy"`
	RouterId                   string `json:"routerId,omitempty"`
	UserId                     string `json:"userId"`
	UserName                   string `json:"userName"`
	CreatedTime                string `json:"createdTime"`
	ChargingMode               string `json:"chargingMode,omitempty"`
	FloatingIpBindResourceResp *struct {
		ResourceName string `json:"resourceName,omitempty"`
		NetworkId    string `json:"networkId,omitempty"`
		RouterId     string `json:"routerId,omitempty"`
		SubnetId     string `json:"subnetId,omitempty"`
		PortId       string `json:"portId,omitempty"`
	} `json:"floatingIpBindResourceResp,omitempty"`
	NatGatewayDnatRuleResps interface{} `json:"natGatewayDnatRuleResps,omitempty"`
}

type FloatingIpQuotaResult struct {
	// k-Version
	TotalIPNum          int `json:"totalIPNum,omitempty"`
	UsedIPNum           int `json:"usedIPNum,omitempty"`
	TotalBandWidthNum   int `json:"totalBandWidthNum,omitempty"`
	UsedBandWidthNum    int `json:"usedBandWidthNum,omitempty"`
	TotalBandWidthValue int `json:"totalBandWidthValue,omitempty"`
	UsedBandWidthValue  int `json:"usedBandWidthValue,omitempty"`

	// p-Version
	Quantity int `json:"quantity,omitempty"`
	Remain   int `json:"remain,omitempty"`
}

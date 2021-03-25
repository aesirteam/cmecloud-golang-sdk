package ELB

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type ELBInterface interface {
	// 创建弹性负载均衡
	CreateELB(es *global.ELBSpec) (ELBOrderResult, error)

	// 查看弹性负载均衡列表
	GetELBList(queryWord, routerId, networkId string, fipBind, visible bool, page, size int) ([]ELBResult, error)

	// 查看弹性负载均衡配额情况
	GetELBQuota() (ELBQuotaResult, error)

	//  删除弹性负载均衡
	DeleteELB(loadBalanceId string) error

	// 创建弹性负载均衡监听器
	CreateELBListener(el *global.ELBListenerSpec) (string, error)

	// 查看弹性负载均衡监听器列表
	GetELBListenerList(loadbalanceId string) ([]ELBListenerResult, error)

	// 查看弹性负载均衡监听器详情
	GetELBListenerInfo(listenerId string) (ELBListenerResult, error)

	// 修改弹性负载均衡监听器名称
	ModifyELBListenerName(el *global.ELBListenerSpec) error

	// 删除弹性负载均衡监听器
	DeleteELBListener(listenerId string) error

	// 向弹性负载均衡添加业务主机
	AddELBMember(ms *global.ELBMemberSpec) (string, error)

	// 查看弹性负载均衡中的业务主机列表
	GetELBMemberList(poolId string) ([]ELBMemberResult, error)

	// 查看弹性负载均衡中的业务主机详情
	GetELBMemberInfo(poolId, memberId string) (ELBMemberResult, error)

	// 从弹性负载均衡中删除一台业务主机
	DeleteELBMember(poolId, memberId string) error
}

type ELBOrderResult struct {
	OrderId       string `json:"orderId"`
	OrderExtResps []struct {
		ProductType    string `json:"productType"`
		OrderExtId     string `json:"orderExtId"`
		OrderExtStatus int    `json:"orderExtStatus"`
	} `json:"orderExtResps"`
}

type ELBResult struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	SubnetId     string `json:"subnetId"`
	PrivateIp    string `json:"privateIp"`
	VipPortId    string `json:"vipPortId"`
	Desc         string `json:"description,omitempty"`
	EcStatus     string `json:"ecStatus"`
	CreateTime   string `json:"createdTime"`
	Proposer     string `json:"proposer"`
	AdminStateUp bool   `json:"adminStateUp"`
	Deleted      bool   `json:"deleted"`
	Visible      bool   `json:"visible"`
	Region       string `json:"region"`
	Provider     string `jsson:"provider"`
	Flavor       int    `json:"flavor"`
	UserName     string `json:"userName"`
	FloatingIpId string `json:"ipId,omitempty"`
	PublicIp     string `json:"publicIp"`
	RouterId     string `json:"routerId"`
	IpVersion    int    `json:"ipVersion"`
	SubnetName   string `json:"subnetName"`
	VpcName      string `json:"vpcName"`
}

type ELBQuotaResult struct {
	UsedQuota  int `json:"usedQuota"`
	TotalQuota int `json:"totalQuota"`
}

type ELBListenerResult struct {
	Id                     string `json:"id"`
	Name                   string `json:"name"`
	Desc                   string `json:"description,omitempty"`
	CreateTime             string `json:"createdTime"`
	ModifiedTime           string `json:"modifiedTime"`
	Deleted                bool   `json:"deleted"`
	OpStatus               string `json:"opStatus"`
	LoadbalanceId          string `json:"loadbalanceId"`
	PoolId                 string `json:"poolId"`
	PoolName               string `json:"poolName"`
	Protocol               string `json:"protocol"`
	ProtocolPort           int    `json:"protocolPort"`
	ConnectionLimit        int    `json:"connectionLimit"`
	Algorithm              string `json:"lbAlgorithm"`
	SessionPersistence     string `json:"sessionPersistence,omitempty"`
	HealthId               string `json:"healthId"`
	HealthType             string `json:"healthType,omitempty"`
	HealthDelay            int    `json:"healthDelay"`
	HealthMaxRetries       int    `json:"healthMaxRetries"`
	HealthTimeout          int    `json:"healthTimeout"`
	HealthHttpMethod       string `json:"healthHttpMethod,omitempty"`
	HealthUrlPath          string `json:"healthUrlPath,omitempty"`
	HealthExpectedCode     string `json:"healthExpectedCode,omitempty"`
	CookieName             string `json:"cookieName,omitempty"`
	DefaultTlsContainerId  string `json:"defaultTlsContainerId,omitempty"`
	MutualAuthenticationUp bool   `json:"mutualAuthenticationUp"`
	CaContainerId          string `json:"caContainerId,omitempty"`
	ControlGroupId         string `json:"controlGroupId,omitempty"`
	GroupType              string `json:"groupType,omitempty"`
	GroupEnabled           string `json:"groupEnabled,omitempty"`
	GroupName              string `json:"groupName,omitempty"`
	LoadBalanceFlavor      int    `json:"loadBalanceFlavor"`
	RedirectToListenerId   string `json:"redirectToListenerId,omitempty"`
	RedirectToListenerName string `json:"redirectToListenerName,omitempty"`
}

type ELBMemberResult struct {
	Id           string `json:"id"`
	Desc         string `json:"description,omitempty"`
	CreateTime   string `json:"createdTime"`
	Status       string `json:"status,omitempty"`
	Ip           string `json:"ip"`
	Port         int    `json:"port"`
	HealthStatus string `json:"healthStatus"`
	PoolId       string `json:"poolId"`
	SubnetId     string `json:"subnetId"`
	VmHostId     string `json:"vmHostId"`
	VmName       string `json:"vmName,omitempty"`
	Weight       int    `json:"weight"`
	Type         int    `json:"type"`
	Proposer     string `json:"proposer"`
	IsDeleted    int    `json:"isDelete"`
	Region       string `json:"region,omitempty"`
}

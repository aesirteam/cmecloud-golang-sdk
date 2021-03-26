package global

/*
虚拟私有云(VPC)
  Cidr	【必填】cidr
  Name		【必填】vpc名称,路由器名称
  NetworkName	【必填】网络名称
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

// VPC子网
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

//VPC 虚拟网卡
type NicSpec struct {
	Name           string
	NetworkId      string
	SecurityGroups []string
	Subnets        []struct {
		IpAddress string
		SubnetId  string
	}
}

/*
IPSecVpn服务
  ChargePeriod	计费周期
  Duration	购买时长
  Quantity	购买数量
  Name	VPN服务名称
  RouterId 关联VPC的routerId
  FloatingIpId 关联弹性公网IP的IpId
*/
type IPSecVpnSpec struct {
	ChargePeriod BillingType
	Duration     int
	Quantity     int
	Name         string
	RouterId     string
	FloatingIpId string
}

/*
IPSecVpn连接
    VpnServiceId	ipsec站点连接所属的ipsecVPN服务id
  	Name	ipsec站点连接的名称
	LocalSubnetId	本端子网id
	PeerAddress	对端网关公网地址 or FQDN
	PeerSubnets	对端子网cidr列表
	Psk	共享秘钥字符串
	IkePolicy:
	    AuthAlgorithm	认证算法
		EncryptionAlgorithm	加密算法
		Pfs	完美向前加密
		Version	IKE策略版本
		Lifetime	生命期
		Phase1NegotiationMode	协商模式
	IpsecPolicy:
	    AuthAlgorithm	认证算法
		EncryptionAlgorithm	加密算法
		EncapsulationMode	传输封装模式
		Pfs	完美向前加密
		Lifetime	生命期
*/
type IPSecVpnSiteConnectionSpec struct {
	VpnServiceId   string
	Name           string
	LocalSubnetIds []string
	PeerAddress    string
	PeerSubnets    []string
	Psk            string
	IkePolicy      struct {
		AuthAlgorithm         VpnAuthAlgorithm
		EncryptionAlgorithm   VpnEncryptionAlgorithm
		Pfs                   VpnPfs
		Version               VpnIkeVersion
		Lifetime              int
		Phase1NegotiationMode VpnIkePhase1NegotiationMode
	}
	IpsecPolicy struct {
		AuthAlgorithm       VpnAuthAlgorithm
		EncryptionAlgorithm VpnEncryptionAlgorithm
		EncapsulationMode   VpnEncapsulationMode
		Pfs                 VpnPfs
		Lifetime            int
	}
}

/*
弹性公网IP
  ChargeMode	带宽计费方式
  ChargePeriod 计费周期
  Duration	购买时长
  Quantity	购买数量
  BandwidthSize	购买带宽值大小, 单位:MB (1M~10240M)
  FloatingIpAddress	指定公网IP创建的IP地址(只能指定当前用户历史订购且已退订的公网IP)
*/
type FloatingIpSpec struct {
	ChargeMode        FloatingIpChargeMode
	ChargePeriod      BillingType
	Duration          int
	Quantity          int
	BandwidthSize     int
	FloatingIpAddress string
}

/*
弹性负载均衡器
  ChargePeriod	计费方式
  LoadBalanceName	负载均衡器名称
  SubnetId	关联VPC子网的id
  FloatingIpId	关联弹性公网IP的IpId
*/
type ELBSpec struct {
	ChargePeriod    BillingType
	LoadBalanceName string
	SubnetId        string
	FloatingIpId    string
}

/*
弹性负载均衡器 监听器
  LoadbalanceId	负载均衡器id
  Name	监听器名称
  Protocol 监听器协议类型
  ProtocolPort	监听器的服务端口
  Algorithm	负载均衡算法
  SessionPersistence	会话保持类型
  CookieName	cookie名称
  HealthType	健康检查方式
  HealthDelay	健康检查间隔
  HealthMaxRetries   最大尝试次数
  HealthTimeout	超时时间
  HealthHttpMethod	用于监控的http方法
  HealthUrlPath	用于监控的url
  HealthExpectedCode 期望http状态代码
  MutualAuthenticationUp	开启双向认证
*/
type ELBListenerSpec struct {
	LoadbalanceId          string
	Name                   string
	Protocol               LoadbalanceProtocol
	ProtocolPort           int
	Algorithm              LoadbalanceAlgorithm
	SessionPersistence     LoadbalanceSessionPersistence
	CookieName             string
	HealthType             LoadbalanceHealthCheck
	HealthDelay            int
	HealthMaxRetries       int
	HealthTimeout          int
	HealthHttpMethod       string
	HealthUrlPath          string
	HealthExpectedCode     string
	MutualAuthenticationUp bool
}

/*
弹性负载均衡器 业务主机
  PoolId	监听器poolid
  Ip	服务器ip
  PortId	服务器端口
  Weight	权重，1-100
  Type	类型 (1:云主机 2:云数据库)
  VmHostId	虚拟机id
  Desc	描述
*/
type ELBMemberSpec struct {
	PoolId   string
	Ip       string
	Port     int
	Weight   int
	Type     int
	VmHostId string
	Desc     string
}

/*
云主机
  Name		云主机名称 (名称为5-22位数字、字母、“-”组合，且“-”不可在名称的开头或结尾)
  Cpu		主机cpu
  Ram		主机ram
  Disk		20/40（创建linux操作系统云主机仅可填写20，windows操作系统仅可填写40）
  BootVolume	云主机系统盘(VolumeType: 系统盘类型，Size: 云硬盘大小(GB)，linux主机系统盘最小为20，windows主机系统盘最小为40，最大均为500)
  Password	密文（采用RSA对明文加密，password和keypairName能且仅能填写一个）
  ImageName	镜像名称
  KeypairName	SSH密钥名称
  Networks	云主机网络 (NetworkId:云主机网络id, PortId: 网卡id)
  SecurityGroupIds 安全组id
  UserData	云主机自定义数据
  VmType		主机规格类型
  Region		可用区
  BillingType	计费类型
  Duration		订购时长（月），取值范围[0,60]。按小时付费与按月后付费填写0，按月预付费按需填写
  Quantity	订购数量，取值范围[1,10]
  DataVolumes	云主机数据盘(ResourceType: 数据盘类型，Size: 数据盘大小(G)，取值范围[10,32768]，IsShare: 是否为共享数据盘，默认为非共享数据盘)
  OsType		操作系统类型
*/
type ServerSpec struct {
	Name       string
	Cpu        int
	Ram        int
	Disk       int
	BootVolume struct {
		VolumeType BootVolumeType
		Size       int
	}
	Password    string
	ImageName   string
	KeypairName string
	Networks    *struct {
		NetworkId string
		PortId    string
	}
	SecurityGroupIds []string
	UserData         string
	VmType
	Region string
	BillingType
	Duration    int
	Quantity    int
	DataVolumes []struct {
		ResourceType DataVolumeType
		Size         int
		IsShare      bool
	}
	OsType
}

/*
安全组规则
  SecurityGroupId	安全组id
  Protocol	访问协议
  MinPortRange	端口范围(最小值)
  MaxPortRange	端口范围(最大值)
  Direction 方向
  EtherType	IP类型
  RemoteType	授权类型
  RemoteSecurityGroupId	授权对象,当RemoteType为security_group时，此值为所选安全组id
  RemoteIpPrefix	授权对象,当RemoteType为cidr时，此值为cidr
*/
type SecurityGroupRuleSpec struct {
	SecurityGroupId       string
	Protocol              SecurityGroupProtocol
	MinPortRange          int
	MaxPortRange          int
	Direction             SecurityGroupDirection
	EtherType             SecurityGroupEtherType
	RemoteType            SecurityGroupRemoteType
	RemoteSecurityGroupId string
	RemoteIpPrefix        string
}

/*
自定义镜像导入
  	ImageUrl	制作好的镜像url
	ImageName	镜像名称
	OsType	操作系统类型
	MinDisk	磁盘容量 (Linux系统最小为20，Windows系统最小为40)
	ImageFormat	镜像格式
	OsVersion	操作系统版本
	Desc	描述
*/
type UserImageImportSpec struct {
	ImageUrl  string
	ImageName string
	OsType
	MinDisk int64
	ImageFormat
	OsVersion string
	Desc      string
}

/*
云硬盘
  CinderType	硬盘cinder类型
  Name	硬盘名称
  Share	是否设置为共享盘
  Size	硬盘大小 (最小值10，最大值视硬盘类型不同)
  Region	可用区
  PeriodType	订购周期类型
  PeriodNum	订购周期数
  BackupId	备份id (基于备份id创建云硬盘)
  BackupPolicyId	硬盘订购成功后需要直接绑定的备份策略id
  ServerId 	硬盘订购成功后需要直接挂载的主机id
*/
type CloudBlockStorageSpec struct {
	CinderType string
	Name       string
	Share      bool
	Size       int
	Region     string
	PeriodType BillingType
	PeriodNum  int
	// 选填
	BackupId       string
	BackupPolicyId string
	ServerId       string
}

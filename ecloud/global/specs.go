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

/*
子网
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

/*
虚拟网卡
*/
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
云主机
  Name		云主机名称，名称为5-22位数字、字母、“-”组合，且“-”不可在名称的开头或结尾
  Cpu		主机cpu
  Ram		主机ram
  Disk		20/40（创建linux操作系统云主机仅可填写20，windows操作系统仅可填写40）
  BootVolume	云主机系统盘(VolumeType: 系统盘类型，Size: 云硬盘大小(GB)，linux主机系统盘最小为20，windows主机系统盘最小为40，最大均为500)
  Password	密文（采用RSA对明文加密，password和keypairName能且仅能填写一个）
  ImageType	镜像类型
  KeypairName	SSH密钥名称
  Networks	云主机网络 (NetworkId:云主机网络id, PortId: 网卡id)
  SecurityGroupIds 安全组id
  UserData	云主机自定义数据
  VmType		主机规格类型
  Region		可用区
  BillingType	计费类型
  Dration		订购时长（月），取值范围[0,60]。按小时付费与按月后付费填写0，按月预付费按需填写
  Quantity	订购数量，取值范围[1,10]
  DataVolumes	云主机数据盘(ResourceType: 数据盘类型，Size: 数据盘大小(G)，取值范围[10,32768]，IsShare: 是否为共享数据盘，默认为非共享数据盘)
  OsType		操作系统类型

  //查询请求参数(非必填)：
  Name		云主机名称result
  Ram		主机ram
  Disk		系统盘大小
  EcStatus	云主机底层状态
  OpStatus	云主机op侧状态
  PublicIp	公网ip地址
  PrivateIp	内网ip地址
  SecurityGroupId	SecurityGroupIds[0]
  NetworkId	Networks[0].NetworkId
  KeypairName	SSH密钥名称
  Region		可用区
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
	Password string
	ImageType
	KeypairName string
	Networks    []struct {
		NetworkId string
		PortId    string
	}
	SecurityGroupIds []string
	UserData         string
	VmType
	Region string
	BillingType
	Dration     int
	Quantity    int
	DataVolumes []struct {
		ResourceType DataVolumeType
		Size         int
		IsShare      bool
	}
	OsType

	ServerId  string
	EcStatus  string
	OpStatus  string
	PublicIp  string
	PrivateIp string
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

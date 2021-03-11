package global

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

/*
云主机
  Name		云主机名称，名称为5-22位数字、字母、“-”组合，且“-”不可在名称的开头或结尾
  Cpu		主机cpu
  Ram		主机ram
  Disk		20/40（创建linux操作系统云主机仅可填写20，windows操作系统仅可填写40；创建前需确认该资源池内是否支持该种虚机）
  BootVolume	云主机系统盘
  Password	密文（采用RSA对明文加密，公钥内容详见请求样例最后，用户可用该公钥自行加密明文得到密文。明文8-16位字符，同时包括数字、大小写字母和特殊字符，其中特殊字符最多不能超过3个，且需要在“~@#$%*_-+=:,.?[]{}”范围内）（password和keypairName能且仅能填写一个）
  ImageType	镜像类型
  KeypairName	SSH密钥名称
  Networks	云主机网络
  SecurityGroupIds 安全组id
  UserData	云主机自定义数据
  VmType		主机规格类型
  Region		可用区
  BillingType	计费类型
  Dration		订购时长（月），取值范围[0,60]。按小时付费与按月后付费填写0，按月预付费按需填写
  Quantity	订购数量，取值范围[1,10]
  DataVolumes	云主机数据盘集合
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
	Name string
	Cpu  int
	Ram  int
	Disk int
	BootVolumeSpec
	Password string
	ImageType
	KeypairName string
	NetworkSpec
	SecurityGroupIds []string
	UserData         string
	VmType
	Region string
	BillingType
	Dration     int
	Quantity    int
	DataVolumes []DataVolumeSpec
	OsType

	ServerId  string
	EcStatus  string
	OpStatus  string
	PublicIp  string
	PrivateIp string
}

/*
云主机系统盘
  VolumeType	系统盘类型
  Size		云硬盘大小(GB)，linux主机系统盘最小为20，windows主机系统盘最小为40，最大均为500
*/
type BootVolumeSpec struct {
	VolumeType BootVolumeType
	Size       int
}

/*
云主机数据盘
  ResourceType	数据盘类型
  Size		数据盘大小(G)，取值范围[10,32768]
  IsShare		是否为共享数据盘，默认为非共享数据盘
*/
type DataVolumeSpec struct {
	ResourceType DataVolumeType
	Size         int
	IsShare      bool
}

/*
云主机网络
	NetworkId	云主机网络id
	PortId		网卡id
*/
type NetworkSpec struct {
	NetworkId string
	PortId    string
}

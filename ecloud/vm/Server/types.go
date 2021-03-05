package Server

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
	//查询可用区
	GetRegionList() (result RegionResultArray, err error)

	//创建云主机
	CreatServer(spec *ServerSpec) (result ServerOrderResult, err error)

	//查询云主机列表
	GetServerList(req *ServerSpec, page, size int) (string, error)

	//查看云主机详情
	GetServerInfo(serverId string, detail bool)

	//获取VNC登录云主机地址
	GetServerVNCAddress(serverId string)

	//修改云主机名称信息
	UpdateServerName(id, name string)

	//重置云主机的密码
	UpdateServerPassword(id, password string)

	//重启云主机
	RebootServer(serverId string)

	//启动云主机
	StartServer(serverId string)

	//停止云主机
	StopServer(serverId string)

	//查看可重装云主机镜像列表 (imageType: 1:用户创建镜像,2:系统镜像)
	GetRebuildImageList(serverId string, imageType int)

	//重装指定云主机上的操作系统 (非必填: adminPass, userData)
	RebuildServer(serverId, imageId string, adminPass, userData string)

	//为指定云主机挂载云硬盘
	AttachStorageDisk(serverId, volumeId string)

	//为指定云主机卸载云硬盘
	DetachStorageDisk(serverId, volumeId string)

	//查看可绑定虚拟网卡列表 (resourceType：云主机0，物理机1)
	GetUnbindNicList(serverId string, resourceType, page, size int)

	//为云主机绑定虚拟网卡
	AttachNic(nicId, serverId string)

	//为云主机解绑虚拟网卡
	DetachNic(nicId, serverId string)

	//查询云主机可用规格
	GetProductFlavorList(spec *ServerSpec, page, size int) (result ProductResultArray, err error)
}

type RegionResult struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"region"`
	Note      string `json:"name,omitempty"`
	Component string `json:"component,omitempty"`
	PoolId    string `json:"poolId,omitempty"`
	Deleted   bool   `json:"deleted,omitempty"`
	Visible   bool   `json:"visible,omitempty"`
}
type RegionResultArray []RegionResult

type ProductResult struct {
	ProductId  string `json:"productId"`
	FlavorId   string `json:"flavorId"`
	OsType     string `json:"osType"`
	ServerType string `json:"serverType"`
	VmType     string `json:"vmType"`
	Generation string `json:"generation,omitempty"`
	Cpu        int    `json:"cpu"`
	Ram        int    `json:"ram"`
	Disk       int    `json:"disk"`
	Gpu        string `json:"gpu,omitempty"`
	BootVolume bool   `json:"bootVolume"`
	Zone       string `json:"zone"`
	Deleted    bool   `json:"deleted"`
}
type ProductResultArray []ProductResult

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
  DataVolumeArray	云主机数据盘
  OsType		操作系统类型

  //查询请求参数(非必填)：
  Name		云主机名称
  ServerId	云主机id
  Cpu		主机cpu
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
	BootVolume
	Password string
	ImageType
	KeypairName string
	Networks
	SecurityGroupIds []string
	UserData         string
	VmType
	Region string
	BillingType
	Dration  int
	Quantity int
	DataVolumeArray
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
type BootVolume struct {
	VolumeType BootVolumeType
	Size       int
}

/*
云主机数据盘
  ResourceType	数据盘类型
  Size		数据盘大小(G)，取值范围[10,32768]
  IsShare		是否为共享数据盘，默认为非共享数据盘
*/
type DataVolume struct {
	ResourceType DataVolumeType
	Size         int
	IsShare      bool
}

type DataVolumeArray []DataVolume

/*
云主机网络
	NetworkId	云主机网络id
	PortId		网卡id
*/
type Networks struct {
	NetworkId string
	PortId    string
}

type ServerOrderResult struct {
	ProcedureCode string   `json:"procedureCode"`
	ReturnUrl     string   `json:"returnUrl,omitempty"`
	OrderId       string   `json:"orderId"`
	OrderExts     []string `json:"orderExts"`
}

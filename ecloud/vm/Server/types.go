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
	//查询云主机可用规格
	GetProductFlavorList(spec ServerSpec, page, size int) (ProductResultArray, error)

	//创建云主机
	CreatServer(spec ServerSpec) (ServerOrderResult, error)

	//查询云主机列表
	GetServerList(req *ServerSpec, page, size int) (ServerResultArray, error)

	//查看云主机详情
	GetServerInfo(serverId string, detail bool) (ServerResult, error)

	//获取VNC登录云主机地址
	GetServerVNCAddress(serverId string) (string, error)

	//修改云主机名称信息
	UpdateServerName(serverId, name string) (ServerResult, error)

	//重置云主机的密码
	UpdateServerPassword(serverId, password string) (ServerResult, error)

	//重启云主机
	RebootServer(serverId string) (ServerResult, error)

	//启动云主机
	StartServer(serverId string) (ServerResult, error)

	//停止云主机
	StopServer(serverId string) (ServerResult, error)

	//查看可重装云主机镜像列表 (imageType: 1:用户创建镜像,2:系统镜像)
	GetRebuildImageList(serverId string, imageType int) (ServerRebuildImageArray, error)

	//重装指定云主机上的操作系统 (非必填: adminPass, userData)
	RebuildServer(serverId, imageId string, adminPass, userData string) error

	//查看可绑定虚拟网卡列表 (resourceType：云主机0，物理机1)
	GetUnbindNicList(serverId string, resourceType, page, size int)

	//为云主机绑定虚拟网卡
	AttachNic(nicId, serverId string)

	//为云主机解绑虚拟网卡
	DetachNic(nicId, serverId string)
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
  DataVolumeArray	云主机数据盘
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

type ServerOrderResult struct {
	ProcedureCode string   `json:"procedureCode"`
	ReturnUrl     string   `json:"returnUrl,omitempty"`
	OrderId       string   `json:"orderId"`
	OrderExts     []string `json:"orderExts"`
}

/*
ServerResult
  ServerId	云主机id
  Name		云主机名称
  Cpu		云主机cpu数量
  Ram		云主机内存大小
  Disk		云主机硬盘大小
  HwHostId	所属物理机id
  HwHost		所属物理机编号
  KeyName		密钥对名称
  ImageRef	镜像id
  ImageName	镜像名称
  ImageOsType	镜像操作系统类型
  FlavorRef	规格id
  SystemDiskId	系统盘id
  Task		底层任务描述
  Deleted		是否已删除
  AvailabilityZone 所属区域
  ServerType	主机类型
  ServerVmType	主机规格类型
  ProductType	云主机产品类型
  OpStatus	云主机op侧状态
  EcStatus	云主机底层状态
  BootVolumeType	云主机系统盘类型
  UserName	订购用户名称
  Visible		是否可见
  Region		可用区
*/
type ServerResult struct {
	ServerId                string `json:"id" newtag:"id"`
	Name                    string `json:"name"`
	Cpu                     int    `json:"vcpu" newtag:"cpu"`
	Ram                     int    `json:"vmemory" newtag:"memory"`
	Disk                    int    `json:"vdisk" newtag:"disk"`
	HwHostId                string `json:"hwHostId,omitempty"`
	HwHost                  string `json:"hwHost,omitempty"`
	KeyName                 string `json:"keyName,omitempty" newtag:"keypairName"`
	Status                  int    `json:"status,omitempty"`
	ImageRef                string `json:"imageRef" newtag:"imageId"`
	ImageName               string `json:"imageName" newtag:"imageName"`
	ImageOsType             string `json:"imageOsType,omitempty"`
	FlavorRef               string `json:"flavorRef" newtag:"flavorId"`
	FlavorName              string `json:"flavorName,omitempty"`
	SystemDiskId            string `json:"systemDiskId" newtag:"systemDisk"`
	Task                    string `json:"task,omitempty"`
	CreatedTime             string `json:"createdTime"`
	ModifiedTime            string `json:"modifiedTime"`
	Deleted                 int    `json:"deleted,omitempty"`
	AvailabilityZone        string `json:"availabilityZone"`
	ServerType              string `json:"serverType,omitempty"`
	ServerVmType            string `json:"serverVmType,omitempty"`
	ProductType             string `json:"productType,omitempty"`
	OpStatus                string `json:"opStatus,omitempty"`
	EcStatus                string `json:"ecStatus,omitempty"`
	BootVolumeType          string `json:"bootVolumeType,omitempty"`
	AdminPaused             bool   `json:"adminPaused,omitempty"`
	ServerPortDetailArray   `json:"portDetail,omitempty" newtag:"ports"`
	ServerVolumeDetailArray `json:"volumeDetail,omitempty" newtag:"volumes"`
	UserName                string `json:"userName,omitempty"`
	Visible                 bool   `json:"visible,omitempty"`
	Region                  string `json:"region"`
	PoolId                  string `json:"poolId,omitempty"`
	OperationFlag           int    `json:"operationFlag,omitempty"`
	DeCloudServerName       string `json:"deCloudServerName,omitempty"`
}
type ServerResultArray []ServerResult

/*
云主机绑定的网络信息
  PrivateIp		内网ip
  FipAddress 		公网ip地址
  FipV6Address 		公网ipv6地址
  FipBandwidthSize	带宽大小
  PortId			网卡id
  PortName		网卡名称
*/
type ServerPortDetail struct {
	IpId                     string `json:"ipId,omitempty"`
	PrivateIp                string `json:"privateIp"`
	FipAddress               string `json:"fipAddress,omitempty" newtag:"publicIp"`
	FipV6Address             string `json:"fipV6Address,omitempty"`
	FipBandwidthSize         int    `json:"fipBandwidthSize,omitempty"`
	PortId                   string `json:"portId" newtag:"id"`
	PortName                 string `json:"portName,omitempty" newtag:"name"`
	ServerFixedIpDetailArray `json:"fixedIpDetailResps,omitempty"`
}
type ServerPortDetailArray []ServerPortDetail

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
}
type ServerFixedIpDetailArray []ServerFixedIpDetail

/*
ServerVolumeDetail
  Id	数据盘id
  Name	数据盘名称
  Size	数据盘大小(G)
  Type	数据盘类型
  Status	使用状态
*/
type ServerVolumeDetail struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Type   string `json:"type"`
	Status string `json:"status"`
}
type ServerVolumeDetailArray []ServerVolumeDetail

/*
ServerRebuildImage
  Id		镜像id
  ServerId	虚机id
  Name		镜像名称
  Type		镜像来源类型
  Url		镜像路径
  SourceImagePoolId 原镜像所属资源池ID
  SourceImageId	原镜像ID
  Status		镜像状态
  Size		镜像大小，单位M
  CreatedTime	创建时间
  ModifiedTime	更新时间
  Note		备注
  OsType		操作系统类型：linux,windows
  MinDisk		flavor需要的最小大小，单位G
*/
type ServerRebuildImage struct {
	Id                string `json:"id"`
	ServerId          string `json:"serverId,omitempty"`
	Name              string `json:"name"`
	Type              int    `json:"type"`
	Url               string `json:"url,omitempty"`
	SourceImagePoolId string `json:"sourceImagePoolId,omitempty"`
	SourceImageId     string `json:"sourceImageId,omitempty"`
	Status            string `json:"status"`
	Size              int    `json:"size"`
	CreatedTime       string `json:"createdTime,omitempty"`
	ModifiedTime      string `json:"modifiedTime,omitempty"`
	Note              string `json:"note,omitempty"`
	OsType            string `json:"osType"`
	MinDisk           int    `json:"minDisk"`
}
type ServerRebuildImageArray []ServerRebuildImage

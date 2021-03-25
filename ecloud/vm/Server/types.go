package Server

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type ServerInterface interface {
	//查询可用区
	GetRegionList() ([]RegionResult, error)

	//查询云主机可用规格
	GetProductFlavorList(cpu, ram int, osType global.OsType, vmType global.VmType, page, size int) ([]ProductResult, error)

	//创建云主机
	CreatServer(spec *global.ServerSpec) (ServerOrderResult, error)

	//查询云主机列表
	GetServerList(name, networkId, region string, page, size int) ([]ServerResult, error)

	//查看云主机详情
	GetServerInfo(serverId string, detail bool) (*ServerResult, error)

	//获取VNC登录云主机地址
	GetServerVNCAddress(serverId string) (string, error)

	//修改云主机名称信息
	ModifyServerName(serverId, name string) (ServerResult, error)

	//重置云主机的密码
	ModifyServerPassword(serverId, password string) (ServerResult, error)

	//重启云主机
	RebootServer(serverId string) (ServerResult, error)

	//启动云主机
	StartServer(serverId string) (ServerResult, error)

	//停止云主机
	StopServer(serverId string) (ServerResult, error)

	//查看可重装云主机镜像列表 (imageType: 1:用户创建镜像,2:系统镜像)
	GetRebuildImageList(serverId string, imageType int) ([]ServerRebuildImageResult, error)

	//重装指定云主机上的操作系统 (非必填: adminPass, userData)
	RebuildServer(serverId, imageId string, adminPass, userData string) error

	//查看可绑定虚拟网卡列表 (resourceType：云主机0，物理机1)
	GetUnbindNicList(serverId string, resourceType, page, size int) ([]VirtualPrivateCloud.NicResult, error)

	//为云主机绑定虚拟网卡
	AttachNic(serverId, portId string) (VirtualPrivateCloud.NicResult, error)

	//为云主机解绑虚拟网卡
	DetachNic(serverId, portId string) (VirtualPrivateCloud.NicResult, error)
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
	SubnetCidr        string `json:"subnetCidr,omitempty"`
	PortId            string `json:"portId,omitempty"`
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
ServerRebuildImageResult
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
type ServerRebuildImageResult struct {
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

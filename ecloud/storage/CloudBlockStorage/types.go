package CloudBlockStorage

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type CBSInterface interface {
	//云硬盘订购
	CreateVolume()

	//云硬盘退订
	DeleteVolume()

	//查询云硬盘列表 (带硬盘的基本信息)
	GetVolumeList()

	//查询云硬盘列表（带挂载主机信息）
	GetVolumeListWithServer()

	//查询云硬盘详情
	GetVolumeInfo()

	//查询云硬盘的全量类型
	GetVolumeTypeList()

	//云硬盘挂载
	MountVolume()

	//云硬盘卸载
	UnmountVolume()

	//查询云主机可挂载云硬盘列表
	GetUnmountList()
}

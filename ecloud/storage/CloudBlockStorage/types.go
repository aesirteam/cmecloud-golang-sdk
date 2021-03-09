package CloudBlockStorage

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

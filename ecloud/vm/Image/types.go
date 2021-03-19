package Image

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

type APIv1 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv1 {
	return &APIv1{client}
}

type ImageInterface interface {
	// 云主机制作自定义镜像
	CreateUserImage(serverId, name string, note string) (string, error)

	// 查询用户创建的所有自定义镜像
	GetUserImageList(imageId, serverId, name string, imageOsTypes, tagIds []string, page, size int) ([]UserImageResult, error)

	// 查询用户创建的自定义镜像详情信息
	GetUserImageInfo(imageId string) (*UserImageResult, error)

	// 更新镜像属性，包括镜像名称和镜像描述
	UpdateUserImageInfo(imageId, name string, note string) (UserImageResult, error)

	// 删除镜像,正在被主机使用中的镜像不允许删除
	DeleteUserImage(imageId string) error

	// 导入自定义镜像信息
	ImportUserImage(cs *global.UserImageImportSpec) (string, error)
}

type UserImageResult struct {
	Id              string `json:"imageId" newtag:"id"`
	ServerId        string `json:"serverId"`
	ImageAlias      string `json:"imageAlias,omitempty"`
	Name            string `json:"name"`
	Url             string `json:"url,omitempty"`
	SourceImageId   string `json:"sourceImageId,omitempty"`
	Status          string `json:"status"`
	Size            int    `json:"size"`
	IsPublic        int    `json:"isPublic,omitempty"`
	CreatedTime     string `json:"createdTime"`
	ModifiedTime    string `json:"modifiedTime,omitempty"`
	Note            string `json:"note,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	OsType          string `json:"osType,omitempty"`
	MinDisk         int    `json:"minDisk"`
	ImageType       string `json:"imageType,omitempty"`
	PublicImageType string `json:"publicImageType,omitempty"`
	BackupType      string `json:"backupType,omitempty"`
	BackupWay       string `json:"backupWay,omitempty"`
	SnapshotId      string `json:"snapshotId,omitempty"`
	Region          string `json:"region,omitempty"`
	IsAscendOrder   bool   `json:"isAscendOrder,omitempty"`
	OsName          string `json:"osName,omitempty"`
}

package ServerBackup

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type ServerBackupInterface interface {
	// 开通云主机备份服务
	OrderServerBackup() (string, error)

	// 创建云主机备份
	CreateServerBackup(serverId, name string) (ServerBackupResult, error)

	// 查看云主机备份列表
	GetServerBackupList(serverId, name string, page, size int) ([]ServerBackupResult, error)

	// 使用云主机备份恢复云主机
	RestoreServerFromBackup(backupId, serverId string) error

	// 删除云主机备份
	DeleteServerBackup(backupId string) error
}

type ServerBackupResult struct {
	Id           string `json:"id"`
	ServerId     string `json:"serverId"`
	Name         string `json:"name"`
	BackupSize   int    `json:"backupSize"`
	Status       int    `json:"status"`
	BackupStatus int    `json:"backupStatus"`
	SystemDisk   string `json:"systemDisk"`
	CreatedTime  string `json:"createdTime"`
	ModifiedTime string `json:"modifiedTime,omitempty"`
}

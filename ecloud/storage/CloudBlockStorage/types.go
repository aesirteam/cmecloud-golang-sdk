package CloudBlockStorage

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

type APIv2 struct {
	client *global.EcloudClient
}

func New(client *global.EcloudClient) *APIv2 {
	return &APIv2{client}
}

type CBSInterface interface {
	// 云硬盘订购
	CreateVolume(spec *global.CloudBlockStorageSpec) (VolumeOrderResult, error)

	// 查询云硬盘列表 (带硬盘的基本信息或主机信息)
	GetVolumeList(queryWord, volumeId string, showServer bool, tagIds []string, page, size int) ([]VolumeResult, error)

	// 查询云硬盘详情
	GetVolumeInfo(volumeId string) (VolumeResult, error)

	// 云硬盘退订
	DeleteVolume(volumeId string) error

	// 查询云硬盘的全量类型
	GetVolumeTypeList(opType, region string) ([]VolumeTypeResult, error)
}

type VolumeTypeResult struct {
	CinderType        string   `json:"cinderType"`
	BackupType        string   `json:"backupType"`
	Priority          int      `json:"priority"`
	AttachServerTypes []string `json:"attachServerTypes"`
	OpType            string   `json:"opType"`
	CustomBack        bool     `json:"customBack"`
	Iscsi             bool     `json:"iscsi"`
	Region            string   `json:"region"`
}

type VolumeOrderResult struct {
	ResourceId   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	OrderId      string `json:"orderId"`
	OrderExtId   string `json:"orderExtId"`
	ReturnUrl    string `json:"returnUrl,omitempty"`
}

type VolumeResult struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Desc              string   `json:"description,omitempty"`
	Status            string   `json:"status"`
	SourceVolumeId    string   `json:"sourceVolumeId,omitempty"`
	AvailabilityZone  string   `json:"availabilityZone,omitempty"`
	BackupId          string   `json:"backupId,omitempty"`
	Size              int      `json:"size"`
	VolumeType        string   `json:"volumeType"`
	Metadata          string   `json:"metadata,omitempty"`
	CreatedTime       string   `json:"createdTime"`
	IsDelete          bool     `json:"isDelete"`
	OperationFlag     string   `json:"operationFlag"`
	UserId            string   `json:"userId,omitempty"`
	Type              string   `json:"type"`
	IsShare           bool     `json:"isShare"`
	ServerIds         []string `json:"serverIds,omitempty"`
	AttachServerTypes []string `json:"attachServerTypes,omitempty"`
	AttachServers     []struct {
		ServerId    string `json:"serverId"`
		ServerName  string `json:"serverName"`
		IsAuthority bool   `json:"isAuthority"`
		ServerType  string `json:"serverType"`
		ProductType string `json:"productType"`
	} `json:"attachSevers,omitempty"`
	Iscsi       bool   `json:"iscsi,omitempty"`
	ProductType string `json:"productType"`
	Region      string `json:"region"`
	UserName    string `json:"userName,omitempty"`
	Recycle     bool   `json:"recycle"`
	RecycleTime string `json:"recycleTime,omitempty"`
}

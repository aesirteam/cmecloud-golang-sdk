package KeyPair

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
	//创建密钥
	CreateKeypair(name, region string) (result *KeypairResult, err error)

	//查询SSH密钥对列表
	GetKeypairList(name, region string, page, size int) (result KeypairResultArray, err error)

	//删除SSH密钥对
	DeleteKeypair(keyId string) error
}

type KeypairResult struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Type         string `json:"ssh,omitempty"`
	FingerPrint  string `json:"fingerPrint,omitempty"`
	Keyaddress   string `json:"keyaddress,omitempty"`
	Status       string `json:"status,omitempty"`
	Material     string `json:"material,omitempty"`
	Proposer     string `json:"proposer,omitempty"`
	CustomerId   string `json:"customerId,omitempty"`
	PoolId       string `json:"poolId,omitempty"`
	CreatedBy    string `json:"createdBy,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty"`
	ModifiedBy   string `json:"modifiedBy,omitempty"`
	ModifiedTime string `json:"modifiedTime,omitempty"`
	IsDelete     int    `json:"isDelete,omitempty"`
	Region       string `json:"region,omitempty"`
	privateKey   string `json:"privateKey,omitempty"`
}
type KeypairResultArray []KeypairResult

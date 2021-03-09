package core

type APIv2 struct {
	client *EcloudClient
}

func NewForConfig(conf *Config) (*APIv2, error) {
	return &APIv2{client: New(conf)}, nil
}

type Interface interface {
	//查询可用区
	GetRegionList() (RegionResultArray, error)
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

package global

type APIv2 struct {
	client *EcloudClient
}

func New(client *EcloudClient) *APIv2 {
	return &APIv2{client}
}

type CoreInterface interface {
	//查询可用区
	GetRegionList() ([]RegionResult, error)
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

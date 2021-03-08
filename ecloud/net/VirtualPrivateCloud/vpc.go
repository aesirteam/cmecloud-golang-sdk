package VirtualPrivateCloud

import (
	"errors"
	"fmt"

	json "github.com/json-iterator/go"
)

func (a *APIv2) CreatVPC() {

}

func (a *APIv2) GetVPCList(natGatewayBind, visible bool, scale VpcScale, region string, tagIds []string, page, size int) (result VpcResultArray, err error) {
	params := map[string]interface{}{
		"scale": scale.String(),
	}

	if natGatewayBind {
		params["natGatewayBind"] = true
	}

	if visible {
		params["visible"] = true
	}

	if region != "" {
		params["region"] = region
	}

	if tagIds != nil && len(tagIds) > 0 {
		params["tagIds"] = tagIds
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/vpc", nil, params, nil)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	obj := json.Get([]byte(resp.Body), "content")
	if obj.LastError() != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, obj.LastError())
		return
	}

	if _err := json.UnmarshalFromString(obj.ToString(), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) GetVPCInfo(vpcId string) (result VpcResult, err error) {
	if vpcId == "" {
		err = errors.New("No vpcId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/vpc/"+vpcId, nil, nil, nil)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := json.UnmarshalFromString(resp.Body, &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) GetVPCInfoByName(name string) (result VpcResult, err error) {
	arr, err := a.GetVPCList(false, false, VPC_SCALE_HIGH, "", nil, 0, 0)
	if err != nil {
		return
	}

	for _, r := range arr {
		if r.Name == name {
			return a.GetVPCInfo(r.Id)
		}
	}

	err = errors.New("No match vpc info")
	return
}

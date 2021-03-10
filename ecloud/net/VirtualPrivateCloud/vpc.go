package VirtualPrivateCloud

import (
	"errors"

	json "github.com/json-iterator/go"
)

func (a *APIv2) CreateVpc(vs *VpcSpec) (vpcId string, err error) {
	if vs.Cidr == "" {
		err = errors.New("No cidr is available")
		return
	}

	if vs.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if vs.NetworkName == "" {
		err = errors.New("No networkName is available")
		return
	}

	if vs.Region == "" {
		err = errors.New("No region is available")
		return
	}

	body := map[string]interface{}{
		"cidr":            vs.Cidr,
		"cidrV6":          vs.CidrV6,
		"name":            vs.Name,
		"networkName":     vs.NetworkName,
		"region":          vs.Region,
		"networkTypeEnum": "VM",
		"specs":           vs.Specs.String(),
	}

	//fmt.Println(core.Dump(body))

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/vpc", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	vpcId = resp.Body

	return
}

func (a *APIv2) DeleteVpc(vpcId string) {

}

func (a *APIv2) GetVpcList(natGatewayBind, visible bool, scale VpcScale, region string, tagIds []string, page, size int) (result []VpcResult, err error) {
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
		err = resp.Error(err)
		return
	}

	obj := json.Get([]byte(resp.Body), "content")
	if obj.LastError() != nil {
		err = resp.Error(err)
		return
	}

	if _err := json.UnmarshalFromString(obj.ToString(), &result); _err != nil {
		err = resp.Error(_err)
		return
	}

	return
}

func (a *APIv2) GetVpcInfo(vpcId string) (result VpcResult, err error) {
	if vpcId == "" {
		err = errors.New("No vpcId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/vpc/"+vpcId, nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if _err := json.UnmarshalFromString(resp.Body, &result); _err != nil {
		err = resp.Error(_err)
		return
	}

	return
}

func (a *APIv2) GetVpcInfoByName(name string) (result VpcResult, err error) {
	arr, err := a.GetVpcList(false, false, VPC_SCALE_HIGH, "", nil, 0, 0)
	if err != nil {
		return
	}

	for _, vpc := range arr {
		if vpc.Name == name {
			return a.GetVpcInfo(vpc.Id)
		}
	}

	err = errors.New("No match vpc info")
	return
}

func (a *APIv2) GetVpcInfoByRouterId(routerId string) {

}

func (a *APIv2) ModifyVpcInfo(vpcId, name, desc string) {

}

func (a *APIv2) GetVpcFirewall(routerId string) {

}

func (a *APIv2) GetVpcNetwork(routerId string) (result []VpcNetResult, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	params := map[string]interface{}{
		"routerId": routerId,
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/network/NetworkResps", nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	obj := json.Get([]byte(resp.Body), "content")
	if obj.LastError() != nil {
		err = resp.Error(obj.LastError())
		return
	}

	if _err := json.UnmarshalFromString(obj.ToString(), &result); _err != nil {
		err = resp.Error(_err)
		return
	}

	return
}

func (a *APIv2) GetVpcVPN(routerId string) {

}

func (a *APIv2) GetVpcNIC(routerId string) {

}

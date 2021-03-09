package VirtualPrivateCloud

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	json "github.com/json-iterator/go"
)

func (a *APIv2) CreatVpc(vs VpcSpec) (result string, err error) {
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
		"name":            vs.Name,
		"networkName":     vs.NetworkName,
		"region":          vs.Region,
		"networkTypeEnum": "VM",
		"routerExternal":  vs.RouterExternal,
		"specs":           vs.Specs.String(),
	}

	var ns []map[string]string

	for _, v := range vs.Subnets {
		if v.Cidr != "" && v.SubnetName != "" {
			ns = append(ns, map[string]string{
				"cidr": v.Cidr,
				"ipVersion": func() string {
					if v.IpVersion == 0 {
						return "4"
					}
					return strconv.Itoa(v.IpVersion)
				}(),
				"subnetName": v.SubnetName,
			})
		}
	}

	if len(ns) == 0 {
		err = errors.New("No subnets is available")
		return
	}

	body["subnets"] = ns

	fmt.Println(core.Dump(body))

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/vpc", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv2) DeleteVpc(vpcId string) {

}

func (a *APIv2) GetVpcList(natGatewayBind, visible bool, scale VpcScale, region string, tagIds []string, page, size int) (result VpcResultArray, err error) {
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

func (a *APIv2) GetVpcInfo(vpcId string) (result VpcResult, err error) {
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

func (a *APIv2) GetVpcInfoByName(name string) (result VpcResult, err error) {
	arr, err := a.GetVpcList(false, false, VPC_SCALE_HIGH, "", nil, 0, 0)
	if err != nil {
		return
	}

	for _, r := range arr {
		if r.Name == name {
			return a.GetVpcInfo(r.Id)
		}
	}

	err = errors.New("No match vpc info")
	return
}

func (a *APIv2) GetVpcInfoByRouterId(id string) {

}

func (a *APIv2) ModifyVpcInfo(vpcId, name, desc string) {

}

func (a *APIv2) GetVpcFirewall(routerId string) {

}

func (a *APIv2) GetVpcNetwork(routerId string) {

}

func (a *APIv2) GetVpcVPN(routerId string) {

}

func (a *APIv2) GetVpcNIC(routerId string) {

}

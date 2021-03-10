package VirtualPrivateCloud

import (
	"errors"
	"strconv"

	json "github.com/json-iterator/go"
)

func (a *APIv2) CreateSubnet(ss *SubnetSpec) (networkId string, err error) {
	if ss.RouterId == "" {
		err = errors.New("No routerId is available")
		return
	}

	if ss.NetworkName == "" {
		err = errors.New("No networkName is available")
		return
	}

	if ss.Region == "" {
		err = errors.New("No region is available")
		return
	}

	body := map[string]interface{}{
		"routerId":              ss.RouterId,
		"networkName":           ss.NetworkName,
		"availabilityZoneHints": ss.Region,
		"region":                ss.Region,
		"networkTypeEnum":       "VM",
	}

	var ns []map[string]string

	for _, v := range ss.Subnets {
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

	//fmt.Println(core.Dump(body))

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/network", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	networkId = resp.Body

	return
}

func (a *APIv2) DeleteSubnet(networkId string) error {
	if networkId == "" {
		return errors.New("No networkId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/netcenter/network/"+networkId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) ModifySubnet(networkId, name string) error {
	if networkId == "" {
		return errors.New("No networkId is available")
	}

	if name == "" {
		return errors.New("No name is available")
	}

	body := map[string]interface{}{
		"networkId":   networkId,
		"networkName": name,
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/netcenter/network/networkName", nil, nil, body)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) GetSubnetInfo(networkId string) (result []SubnetResult, err error) {
	if networkId == "" {
		err = errors.New("No networkId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/subnet/networkId/"+networkId+"/SubnetResps", nil, nil, nil)
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
		err = resp.Error(err)
		return
	}

	return
}

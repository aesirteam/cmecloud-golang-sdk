package VirtualPrivateCloud

import (
	"errors"
	"strconv"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateSubnet(ss *global.SubnetSpec) (networkId string, err error) {
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

	subnets := make([]map[string]string, 0)

	for _, v := range ss.Subnets {
		if v.Cidr != "" && v.SubnetName != "" {
			subnets = append(subnets, map[string]string{
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

	if len(subnets) == 0 {
		err = errors.New("No subnets is available")
		return
	}

	body["subnets"] = subnets

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

func (a *APIv2) ModifySubnetName(networkId, networkName string) error {
	if networkId == "" {
		return errors.New("No networkId is available")
	}

	if networkName == "" {
		return errors.New("No networkName is available")
	}

	body := map[string]interface{}{
		"networkId":   networkId,
		"networkName": networkName,
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/netcenter/network/networkName", nil, nil, body)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) GetSubnetList(networkId string) (result []SubnetResult, err error) {
	if networkId == "" {
		err = errors.New("No networkId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/subnet/networkId/"+networkId+"/SubnetResps", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetSubnetInfo(subnetId string) (result SubnetResult, err error) {
	if subnetId == "" {
		err = errors.New("No subnetId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/subnet/subnetId/"+subnetId+"/SubnetDetailResp", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

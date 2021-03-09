package VirtualPrivateCloud

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
)

func (a *APIv2) CreateSubnet(vs VpcSpec) (result string, err error) {
	if vs.RouterId == "" {
		err = errors.New("No routerId is available")
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
		"routerId":              vs.RouterId,
		"networkName":           vs.NetworkName,
		"availabilityZoneHints": vs.Region,
		"region":                vs.Region,
		"networkTypeEnum":       "VM",
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

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/network", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv2) DeleteSubnet(networkId string) error {
	return nil
}

func (a *APIv2) ModifySubnet(networkId, name string) {

}

func (a *APIv2) GetSubnetInfo(subnetId string) {

}

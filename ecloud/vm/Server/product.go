package Server

import "github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"

func (a *APIv2) GetProductFlavorList(cpu, ram int, osType global.OsType, vmType global.VmType, page, size int) (result []ProductResult, err error) {
	params := map[string]interface{}{
		"serverType": "VM",
		"category":   "NORMAL",
	}

	if cpu > 0 {
		params["cpu"] = cpu
	}

	if ram > 0 {
		params["ram"] = ram
	}

	switch osType {
	case global.OS_TYPE_WINDOWS:
		params["disk"] = 40
	default:
		params["disk"] = 20
	}

	params["osType"] = osType.String()
	params["vmType"] = vmType.String()

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/server/product/flavor", nil, params, nil)
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

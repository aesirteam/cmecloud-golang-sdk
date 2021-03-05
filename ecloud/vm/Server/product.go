package Server

import (
	"fmt"
	json "github.com/json-iterator/go"
)

func (a *APIv2) GetProductFlavorList(ss *ServerSpec, page, size int) (result ProductResultArray, err error) {
	params := map[string]interface{}{
		"category": "NORMAL",
	}

	if ss.Cpu > 0 {
		params["cpu"] = ss.Cpu
	}

	if ss.Ram > 0 {
		params["ram"] = ss.Ram
	}

	switch ss.OsType {
	case OS_TYPE_WINDOWS:
		params["osType"] = OS_TYPE_WINDOWS.String()
		params["disk"] = 40
	default:
		params["osType"] = OS_TYPE_LINUX.String()
		params["disk"] = 20
	}

	if ss.VmType.String() != "" {
		params["vmType"] = ss.VmType.String()
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/server/product/flavor", nil, params, nil)
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

func (a *ProductResultArray) Dump() string {
	if bytes, err := json.MarshalIndent(&a, "", "  "); err == nil {
		return fmt.Sprintf("\n%s", bytes)
	}
	return ""
}

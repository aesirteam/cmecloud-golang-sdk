package Server

import (
	"fmt"

	json "github.com/json-iterator/go"
)

func (a *APIv2) GetRegionList() (result RegionResultArray, err error) {
	params := map[string]interface{}{
		"component": "NOVA",
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/region", nil, params, nil)
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

package core

import (
	json "github.com/json-iterator/go"
)

func (a *APIv2) GetRegionList() (result []RegionResult, err error) {
	params := map[string]interface{}{
		"component": "NOVA",
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/region", nil, params, nil)
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

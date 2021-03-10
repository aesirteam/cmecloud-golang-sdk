package VirtualPrivateCloud

import (
	"errors"

	json "github.com/json-iterator/go"
)

func (a *APIv2) GetRouterNetList(routerId string) (result []RouterNetResult, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/router/"+routerId+"/interface", nil, nil, nil)
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

func (a *APIv2) GetRouterInfo(routerId string) (result RouterResult, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/router/"+routerId+"/RouterDetailResp", nil, nil, nil)
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

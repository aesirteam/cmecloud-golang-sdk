package VirtualPrivateCloud

import (
	"errors"
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

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
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

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

//Deprecated: OpenRouterGateway
func (a *APIv2) OpenRouterGateway(routerId string) (result string, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	resp, err := a.client.NewRequest("PUT", "/api/router/"+routerId+"/gateway", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}

//Deprecated: CloseRouterGateway
func (a *APIv2) CloseRouterGateway(routerId string) error {
	if routerId == "" {
		return errors.New("No routerId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/router/"+routerId+"/gateway", nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

//Deprecated: RouterAttachSubnet
func (a *APIv2) RouterAttachSubnet() {

}

//Deprecated: RouterDetachSubnet
func (a *APIv2) RouterDetachSubnet() {

}

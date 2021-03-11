package global

func (a *APIv2) GetRegionList() (result []RegionResult, err error) {
	params := map[string]interface{}{
		"component": "NOVA",
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/region", nil, params, nil)
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

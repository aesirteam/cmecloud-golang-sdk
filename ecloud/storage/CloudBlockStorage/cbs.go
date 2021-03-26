package CloudBlockStorage

import (
	"errors"
	"strings"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateVolume(ss *global.CloudBlockStorageSpec) (result VolumeOrderResult, err error) {
	if ss.CinderType == "" {
		err = errors.New("No cinderType is available")
		return
	}

	if ss.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if ss.Size < 10 {
		err = errors.New("No size is available")
		return
	}

	if ss.Region == "" {
		err = errors.New("No region is available")
		return
	}

	body := map[string]interface{}{
		"cinderType":  ss.CinderType,
		"name":        ss.Name,
		"needConfirm": true,
		"productType": "NORMAL",
		"share":       ss.Share,
		"size":        ss.Size,
		"region":      ss.Region,
		"periodType":  strings.ToLower(ss.PeriodType.String()),
	}

	switch ss.PeriodType {
	case global.BILLING_TYPE_YEAR:
		if ss.PeriodNum == 0 {
			body["periodNum"] = 12
		} else if ss.PeriodNum > 0 && ss.PeriodNum <= 5*12 {
			body["periodNum"] = ss.PeriodNum
		}
	case global.BILLING_TYPE_MONTH:
		if ss.PeriodNum == 0 {
			body["periodNum"] = 1
		} else if ss.PeriodNum > 0 && ss.PeriodNum <= 12 {
			body["periodNum"] = ss.PeriodNum
		}
	}

	if ss.BackupId != "" {
		body["backupId"] = ss.BackupId
	}

	if ss.BackupPolicyId != "" {
		body["backupPolicyId"] = ss.BackupPolicyId
	}

	if ss.ServerId != "" {
		body["serverId"] = ss.ServerId
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/volume/volume/order/volume", nil, nil, body)
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

func (a *APIv2) DeleteVolume(volumeId string) (err error) {
	if volumeId == "" {
		err = errors.New("No volumeId is available")
		return
	}

	body := map[string]interface{}{
		"resourceId":   volumeId,
		"resourceType": "VOLUME",
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/volume/common/resource/preDelete", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetVolumeList(queryWord, volumeId string, showServer bool, tagIds []string, page, size int) (result []VolumeResult, err error) {
	params := map[string]interface{}{}

	if queryWord != "" {
		params["likeSearch"] = queryWord
	}

	if volumeId != "" {
		params["volumeId"] = volumeId
	}

	if tagIds != nil && len(tagIds) > 0 {
		params["tagIds"] = strings.Join(tagIds, ",")
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	var prefixPath = "/api/v2/volume/volume"

	if showServer {
		prefixPath = "/api/v2/volume/volume/volume/list/with/server"
	}

	resp, err := a.client.NewRequest("GET", prefixPath, nil, params, nil)
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

func (a *APIv2) GetVolumeInfo(volumeId string) (result VolumeResult, err error) {
	if volumeId == "" {
		err = errors.New("No volumeId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/volume/volume/volumeDetail/"+volumeId, nil, nil, nil)
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

func (a *APIv2) GetVolumeTypeList(opType, region string) (result []VolumeTypeResult, err error) {
	resp, err := a.client.NewRequest("GET", "/api/v2/volume/types", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	if opType != "" && region != "" {
		var obj []VolumeTypeResult
		for _, v := range result {
			if v.OpType == opType && v.Region == region {
				obj = append(obj, v)
				break
			}
		}

		result = obj
	}

	return
}

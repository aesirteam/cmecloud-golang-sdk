package Image

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv1) CreateUserImage(serverId, name string, note string) (result string, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	if name == "" {
		err = errors.New("No name is available")
		return
	}

	body := map[string]interface{}{
		"serverId": serverId,
		"name":     name,
	}

	if note != "" {
		body["note"] = note
	}

	resp, err := a.client.NewRequest("POST", "/api/image", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	var obj map[string]interface{}

	if err = resp.UnmarshalFromContent(&obj, ""); err != nil {
		err = resp.Error(err)
		return
	}

	result = obj["id"].(string)

	return
}

func (a *APIv1) GetUserImageList(imageId, serverId, name string, imageOsTypes, tagIds []string, page, size int) (result []UserImageResult, err error) {
	params := map[string]interface{}{}

	if imageId != "" {
		params["imageId"] = imageId
	}

	if serverId != "" {
		params["serverId"] = serverId
	}

	if name != "" {
		params["name"] = name
	}

	if imageOsTypes != nil && len(imageOsTypes) > 0 {
		params["imageOsTypes"] = strings.Join(imageOsTypes, ",")
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

	resp, err := a.client.NewRequest("GET", "/api/image", nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv1) GetUserImageInfo(imageId string) (result *UserImageResult, err error) {
	if imageId == "" {
		err = errors.New("No imageId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/image/"+imageId, nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv1) UpdateUserImageInfo(imageId, name string, note string) (result UserImageResult, err error) {
	if imageId == "" {
		err = errors.New("No imageId is available")
		return
	}

	if name == "" && note == "" {
		err = errors.New("No name and note is available")
		return
	}

	body := map[string]interface{}{
		"id": imageId,
	}

	if name != "" {
		body["name"] = name
	}

	if note != "" {
		body["note"] = note
	}

	resp, err := a.client.NewRequest("PUT", "/api/image", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv1) DeleteUserImage(imageId string) (err error) {
	if imageId == "" {
		err = errors.New("No imageId is available")
		return
	}

	resp, err := a.client.NewRequest("DELETE", "/api/image/"+imageId, nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv1) ImportUserImage(cs *global.UserImageImportSpec) (result string, err error) {
	if cs.ImageUrl == "" {
		err = errors.New("No ImageUrl is available")
		return
	}

	if cs.ImageName == "" {
		err = errors.New("No ImageName is available")
		return
	}

	body := map[string]interface{}{
		"imageURL":    cs.ImageUrl,
		"imageName":   cs.ImageName,
		"osversion":   cs.OsVersion,
		"imageFormat": cs.ImageFormat.String(),
		"mindisk":     cs.MinDisk,
		"description": cs.Desc,
	}

	switch cs.OsType {
	case global.OS_TYPE_LINUX:
		body["osType"] = "Linux"
		if cs.MinDisk < 20 {
			err = errors.New("Mindisk greater than 20 with os linux")
			return
		}
	case global.OS_TYPE_WINDOWS:
		body["osType"] = "Windows"
		if cs.MinDisk < 40 {
			err = errors.New("Mindisk greater than 40 with os windows")
			return
		}
	}

	if cs.OsVersion == "" {
		body["osversion"] = "other"
	}

	if cs.Desc == "" {
		body["description"] = fmt.Sprintf("%s_%s_%s.%s", cs.ImageName, cs.OsType.String(), cs.OsVersion, cs.ImageFormat.String())
	}

	resp, err := a.client.NewRequest("POST", "/api/image/inputImage", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	var obj map[string]string

	if err = resp.UnmarshalFromContent(&obj, ""); err != nil {
		err = resp.Error(err)
		return
	}

	result = obj["imageId"]

	return
}

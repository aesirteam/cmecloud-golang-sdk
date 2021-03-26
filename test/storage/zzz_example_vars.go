package storage

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

var cli = ecloud.NewForConfigDie(&global.Config{
	ApiGwHost: "api-guiyang-1.cmecloud.cn",
	//ApiGwPort:     8443,
	ApiGwProtocol: "https",
	// AccessKey:     "YOUR ACCESS KEY",
	// SecretKey:     "YOUR SECRET KEY",
})

var vm, storage = cli.VM(), cli.Storage()

var (
	cbsVolumeName = "BS99999"
	cbsVolumeId   string
)

var region = func() string {
	if result, err := vm.GetRegionList("CINDER"); err == nil && len(result) > 0 {
		return result[0].Name
	}
	return ""
}()

var getCinderType = func() string {
	if result, err := storage.GetVolumeTypeList(global.DATA_VOLUME_CAPEBS.String(), region); err == nil && len(result) > 0 {
		return result[0].CinderType
	}

	return ""
}

var getVolumeId = func() string {
	if cbsVolumeId == "" {
		if result, err := storage.GetVolumeList(cbsVolumeName, "", false, nil, 0, 0); err == nil && len(result) > 0 {
			cbsVolumeId = result[0].Id
		}
	}

	return cbsVolumeId
}

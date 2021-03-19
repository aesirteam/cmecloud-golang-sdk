package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestUserImage(t *testing.T) {
	vm := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).VM()

	var (
		serverName        = "api198994873"
		imageName         = "imageApi198994873"
		serverId, imageId string
	)

	getServerId := func() string {
		if serverId == "" {
			if result, err := vm.GetServerList(&global.ServerSpec{Name: serverName}, 0, 0); err == nil && len(result) > 0 {
				serverId = result[0].ServerId
			}
		}
		return serverId
	}

	getImageId := func() string {
		if imageId == "" {
			if result, err := vm.GetUserImageList("", "", imageName, nil, nil, 0, 0); err == nil && len(result) > 0 {
				imageId = result[0].Id
			}
		}

		return imageId
	}

	t.Run("CreateUserImage", func(t *testing.T) {
		serverId = getServerId()

		var err error
		imageId, err = vm.CreateUserImage(serverId, imageName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("imageId:%s\n", imageId)
	})

	t.Run("GetUserImageList", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.GetUserImageList("", "", "", nil, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetUserImageInfo", func(t *testing.T) {
		imageId = getImageId()

		result, err := vm.GetUserImageInfo(imageId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("UpdateUserImageInfo", func(t *testing.T) {
		imageId = getImageId()

		result, err := vm.UpdateUserImageInfo(imageId, imageName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteUserImage", func(t *testing.T) {
		imageId = getImageId()

		err := vm.DeleteUserImage(imageId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ImportUserImage", func(t *testing.T) {
		spec := global.UserImageImportSpec{
			ImageUrl:    "http://eos-guiyang-1.cmecloud.cn/custom-images/blj.qcow2",
			ImageName:   "blj_qcow2",
			OsType:      global.OS_TYPE_LINUX,
			MinDisk:     20,
			ImageFormat: global.IMAGE_FORMAT_QCOW2,
		}

		imageId, err := vm.ImportUserImage(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("imageId:%s\n", imageId)
	})
}

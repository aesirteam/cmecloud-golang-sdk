package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestUserImage(t *testing.T) {

	var (
		imageName = "image" + serverName
	)

	t.Run("CreateUserImage", func(t *testing.T) {
		serverId = getServerId()

		var err error
		userImageId, err = vm.CreateUserImage(serverId, imageName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("imageId:%s\n", userImageId)
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
		userImageId = getUserImageId(imageName)

		result, err := vm.GetUserImageInfo(userImageId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("UpdateUserImageInfo", func(t *testing.T) {
		userImageId = getUserImageId(imageName)

		result, err := vm.UpdateUserImageInfo(userImageId, imageName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteUserImage", func(t *testing.T) {
		userImageId = getUserImageId(imageName)

		err := vm.DeleteUserImage(userImageId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ImportUserImage", func(t *testing.T) {
		spec := global.UserImageImportSpec{
			ImageUrl:    "http://eos-guiyang-1.cmecloud.cn/custom-images/Ubuntu-1604.qcow2",
			ImageName:   "Ubuntu_1604",
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

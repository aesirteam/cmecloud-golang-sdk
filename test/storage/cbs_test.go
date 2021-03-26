package storage

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestCBS(t *testing.T) {
	t.Run("GetVolumeTypeList", func(t *testing.T) {
		result, err := storage.GetVolumeTypeList("", "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("CreateVolume", func(t *testing.T) {
		cinderType := getCinderType()

		spec := global.CloudBlockStorageSpec{
			CinderType: cinderType,
			Name:       cbsVolumeName,
			Size:       10,
			PeriodType: global.BILLING_TYPE_HOUR,
			Region:     region,
		}

		result, err := storage.CreateVolume(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVolumeList", func(t *testing.T) {
		result, err := storage.GetVolumeList("", "", true, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVolumeInfo", func(t *testing.T) {
		cbsVolumeId = getVolumeId()

		result, err := storage.GetVolumeInfo(cbsVolumeId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteVolume", func(t *testing.T) {
		cbsVolumeId = getVolumeId()

		err := storage.DeleteVolume(cbsVolumeId)
		if err != nil {
			t.Fatal(err)
		}
	})
}

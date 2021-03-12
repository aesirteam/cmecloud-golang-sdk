package vm

import (
	"strconv"
	"testing"
	"time"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestKeypair(t *testing.T) {
	vm := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).VM()

	var (
		name  = "kp" + strconv.Itoa(time.Now().Nanosecond())
		keyId string
	)

	t.Run("CreateKeypair", func(t *testing.T) {
		result, err := vm.CreateKeypair(name, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})

	t.Run("GetKeypairList", func(t *testing.T) {
		result, err := vm.GetKeypairList(name, "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		keyId = result[0].Id
	})

	t.Run("DeleteKeypair", func(t *testing.T) {
		err := vm.DeleteKeypair(keyId)
		if err != nil {
			t.Fatal(err)
		}
	})
}

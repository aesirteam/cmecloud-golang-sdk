package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestSecurityGroup(t *testing.T) {
	vm := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).VM()

	var (
		sgName = "SecurityGroup_test"
		sgId   string
	)

	getSecurityGroupId := func() string {
		result, err := vm.GetSecurityGroupList(sgName, 0, 0)
		if err != nil {
			t.Fatal(err)
		}
		return result[0].Id
	}

	t.Run("CreateSecurityGroup", func(t *testing.T) {
		result, err := vm.CreateSecurityGroup(sgName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		sgId = result.Id
	})

	t.Run("GetSecurityGroupList", func(t *testing.T) {
		result, err := vm.GetSecurityGroupList(sgName, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifySecurityGroup", func(t *testing.T) {
		if sgId == "" {
			sgId = getSecurityGroupId()
		}

		result, err := vm.ModifySecurityGroup(sgId, sgName+"_1", "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

	})

	t.Run("DeleteSecurityGroup", func(t *testing.T) {
		if sgId == "" {
			sgId = getSecurityGroupId()
		}

		err := vm.DeleteSecurityGroup(sgId)
		if err != nil {
			t.Fatal(err)
		}
	})

}

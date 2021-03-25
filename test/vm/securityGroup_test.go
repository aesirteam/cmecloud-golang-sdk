package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestSecurityGroup(t *testing.T) {

	t.Run("CreateSecurityGroup", func(t *testing.T) {
		result, err := vm.CreateSecurityGroup(securityGroupName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		securityGroupId = result.Id
	})

	t.Run("GetSecurityGroupList", func(t *testing.T) {
		result, err := vm.GetSecurityGroupList("", true, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifySecurityGroup", func(t *testing.T) {
		securityGroupId = getSecurityGroupId()

		newName := securityGroupName + "_1"
		result, err := vm.ModifySecurityGroup(securityGroupId, newName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		securityGroupName = newName
	})

	t.Run("AddSecurityGroupRules", func(t *testing.T) {
		securityGroupId = getSecurityGroupId()

		spec := global.SecurityGroupRuleSpec{
			SecurityGroupId: securityGroupId,
			Protocol:        global.SECURITYGROUP_PROTOCOL_TCP,
			//Direction:       global.SECURITYGROUP_DIRECTION_EGRESS,
			MinPortRange: securityGroupPortRange[0],
			MaxPortRange: securityGroupPortRange[1],
			EtherType:    global.SECURITYGROUP_ETHERTYPE_IPV4,
			//RemoteType: global.SECURITYGROUP_REMOTETYPE_SECURITYGROUP,
			//RemoteIpPrefix: "10.20.10.0/0",
			//RemoteSecurityGroupId: securityGroupId,
		}

		result, err := vm.AddSecurityGroupRules(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		securityGroupRuleId = result.Id
	})

	t.Run("GetSecurityGroupRules", func(t *testing.T) {
		securityGroupId = getSecurityGroupId()

		result, err := vm.GetSecurityGroupRules(securityGroupId, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteSecurityGroupRules", func(t *testing.T) {
		securityGroupRuleId = getSecurityGroupRuleId()

		err := vm.DeleteSecurityGroupRules(securityGroupRuleId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteSecurityGroup", func(t *testing.T) {
		securityGroupId = getSecurityGroupId()

		err := vm.DeleteSecurityGroup(securityGroupId)
		if err != nil {
			t.Fatal(err)
		}
	})

}

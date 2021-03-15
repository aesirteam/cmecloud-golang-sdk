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
		sgName                     = "SecurityGroup_test"
		minPortRange, maxPortRange = 22, 22
		sgId                       string
		ruleId                     string
	)

	getSecurityGroupId := func() string {
		result, err := vm.GetSecurityGroupList(sgName, false, 0, 0)
		if err != nil {
			t.Fatal(err)
		}
		return result[0].Id
	}

	getSecurityGroupRuleId := func(sgId string) string {
		result, err := vm.GetSecurityGroupRules(sgId, "")
		if err != nil {
			t.Fatal(err)
		}

		for _, rule := range result {
			if rule.PortRangeMin == minPortRange && rule.PortRangeMax == maxPortRange {
				return rule.Id
			}
		}

		return ""
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
		result, err := vm.GetSecurityGroupList(sgName, true, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifySecurityGroup", func(t *testing.T) {
		if sgId == "" {
			sgId = getSecurityGroupId()
		}

		newName := sgName + "_1"
		result, err := vm.ModifySecurityGroup(sgId, newName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		sgName = newName
	})

	t.Run("AddSecurityGroupRules", func(t *testing.T) {
		if sgId == "" {
			sgId = getSecurityGroupId()
		}

		spec := global.SecurityGroupRuleSpec{
			SecurityGroupId: sgId,
			Protocol:        global.SECURITYGROUP_PROTOCOL_TCP,
			//Direction:       global.SECURITYGROUP_DIRECTION_EGRESS,
			MinPortRange: minPortRange,
			MaxPortRange: maxPortRange,
			EtherType:    global.SECURITYGROUP_ETHERTYPE_IPV4,
			//RemoteType: global.SECURITYGROUP_REMOTETYPE_SECURITYGROUP,
			//RemoteIpPrefix: "10.20.10.0/0",
			//RemoteSecurityGroupId: sgId,
		}

		result, err := vm.AddSecurityGroupRules(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		ruleId = result.Id
	})

	t.Run("GetSecurityGroupRules", func(t *testing.T) {
		if sgId == "" {
			sgId = getSecurityGroupId()
		}

		result, err := vm.GetSecurityGroupRules(sgId, ruleId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteSecurityGroupRules", func(t *testing.T) {
		if sgId == "" {
			sgId = getSecurityGroupId()
		}

		if ruleId == "" {
			ruleId = getSecurityGroupRuleId(sgId)
		}

		err := vm.DeleteSecurityGroupRules(ruleId)
		if err != nil {
			t.Fatal(err)
		}
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

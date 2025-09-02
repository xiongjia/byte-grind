package util

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/adhoc"
)

func AnsibleExec() error {
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Args: `msg="
		{{ arg1 }}
		{{ arg2 }}
		{{ arg3 }}
		"`,
		Connection: "local",
		ExtraVars: map[string]interface{}{
			"arg1": map[string]interface{}{"subargument": "subargument_value"},
			"arg2": "arg2_value",
			"arg3": "arg3_value",
		},
		Inventory:  "192.168.71.180,",
		ModuleName: "debug",
	}

	return adhoc.NewAnsibleAdhocExecute("all").
		WithAdhocOptions(ansibleAdhocOptions).
		Execute(context.TODO())
}

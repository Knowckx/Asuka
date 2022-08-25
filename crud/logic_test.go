package crud

import (
	"testing"

	infa "github.com/Knowckx/ainfa"
)

func Test_DiffClientSecrets(t *testing.T) {
	old := []*AzureClient{
		&AzureClient{SubsName: "a10"},
		&AzureClient{SubsName: "a11"},
		&AzureClient{SubsName: "a12"},
	}
	new := []*AzureClient{
		&AzureClient{SubsName: "a11"},
		&AzureClient{SubsName: "a12"},
		&AzureClient{SubsName: "a13"},
	}
	unfound := DiffClientSecrets(old, new)
	infa.PrintJson(unfound)
}

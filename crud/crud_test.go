package crud

import (
	"testing"
)

/*
	把常用的CRUD都写成例子
	CRUD = create, read, update and delete.
*/

// insert
func Test_ClientSecretInsertTo(t *testing.T) {
	// InitPGForTest()
	ins := []*AzureClient{}
	data := NewAzureClientTese()
	ins = append(ins, data)
	Inserts(ins)
}

func NewAzureClientTese() *AzureClient {
	out := &AzureClient{}
	out.ClientID = "c79c6b22-7efb-4e09-a012-32f372613ef9"
	out.ClientSecret = "test"
	out.SubscriptionID = "9d864643-1d6b-4ce5-93f8-60eb1e6a2172"
	out.TenantID = "69b863e3-480a-4ee9-8bd0-20a8adb6909b"
	out.SubscriptionName = "sap-cic-polar-dev1"
	return out
}

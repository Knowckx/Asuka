package k8sclient

import (
	"fmt"
	"testing"

	infa "github.com/Knowckx/ainfa"
)

func Test_GetEAEvent(t *testing.T) {
	kubeconfig := GetTestKubeconfig()
	coreClient, err := CreateCoreClientFromKubeconfig(kubeconfig)
	infa.PanicIfErr(err)
	ea := &EphemeralAccess{}
	ea.Name = "cheng-postgres-292"
	ea.Namespace = "dx-test3"
	event, err := GetEALatestEvent(coreClient, ea)

	infa.PanicIfErr(err)
	if event == nil {
		fmt.Println("result event nil")
	} else {
		fmt.Println("result event ", event.Message, event.Count)
	}
}

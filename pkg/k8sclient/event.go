package k8sclient

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	infa "github.com/Knowckx/ainfa"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

/* 获取对象的event */

func GetTestKubeconfig() string {
	kubeconfigPath := `/Users/i544887/dev/repos/kubeconfig/uat.yml`
	res, err := infa.ReadFile(kubeconfigPath)
	infa.PanicIfErr(err)
	return res
}

func CreateCoreClientFromKubeconfig(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*clientcmdapi.Config, error) {
		return clientcmd.Load([]byte(kubeconfig))
	})
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

// GetEALatestEvent from kind=EphemeralAccess
func GetEALatestEvent(coreClient *kubernetes.Clientset, ea *EphemeralAccess) (*corev1.Event, error) {
	option := metav1.ListOptions{}
	option.FieldSelector = fmt.Sprintf("involvedObject.name=%s,involvedObject.kind=EphemeralAccess", ea.Name)
	eventList, err := coreClient.CoreV1().Events(ea.Namespace).List(context.TODO(), option)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var result corev1.Event
	if len(eventList.Items) > 0 {
		result = eventList.Items[len(eventList.Items)-1]
		return &result, nil
	}
	return nil, nil
}

type EphemeralAccess struct {
	Name      string
	Namespace string
}

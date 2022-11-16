package k8sclient

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	infa "github.com/Knowckx/ainfa"
	"github.com/Knowckx/ainfa/path"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.tools.sap/aeolia/in-fa/parallel"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

//-- 一堆kubeconfig的验证 用上了我们包的waitpool

// 常用：获取测试环境某个路径下的kubeconfig
func GetTestKubeConfig() string {
	projName := "dev-datastore-operator"
	pa := path.GetProjRootPath(projName)
	kubePath := pa + ".vscode/cic-dev1.yml"
	content, err := infa.ReadFile(kubePath)
	infa.PanicIfErr(err)
	return content
}

func Test_KubeconfigChecker(t *testing.T) {
	ch := NewKubeconfigChecker("", 100)
	ch.InPath = "/Users/i544887/dev/repos/infra/bom-ctl/kubeconfigcheck/draft/allconfigs"
	ch.Start()
}

type KubeconfigChecker struct {
	InPath     string
	wg         *parallel.WaitPool
	ErrorCount int
	Total      int
}

func NewKubeconfigChecker(inpath string, total int) *KubeconfigChecker {
	out := &KubeconfigChecker{}
	out.InPath = inpath
	out.wg = parallel.NewWaitPool(total) // 1 声明同时跑的最大的协程数量
	return out
}

func (t *KubeconfigChecker) Start() error {
	filepath.Walk(t.InPath, t.fnwalk)
	t.wg.Wait()
	log.Info().Msgf("Error Kubeconfig Count %d,  All Kubeconfig Count %d", t.ErrorCount, t.Total)
	return nil
}

type InputArgs struct {
	*parallel.ProcessPool
	Args string // 带入的参数
}

// 对每个文件的操作
func (t *KubeconfigChecker) fnwalk(path string, info os.FileInfo, err error) error {
	if !strings.Contains(info.Name(), "kubeconfig") {
		return nil
	}
	t.wg.Add(1)
	t.Total++
	go t.DoCheck(path)
	return nil
}

func (t *KubeconfigChecker) DoCheck(inPath string) {
	defer t.wg.Done()
	splitPoint := "allconfigs"
	tempPath := strings.Split(inPath, splitPoint)
	splitPoint2 := "kubeconfig"
	mmpInfo := strings.Split(tempPath[1], splitPoint2)
	mmpAndCluster := mmpInfo[0]
	// log.Info().Msgf("mmpInfo %s", mmpAndCluster)
	cli, err := ReadKubeconfig(inPath)
	if err != nil {
		log.Err(err).Msg("")
	}
	if !CheckClientIsValid(cli) {
		log.Info().Msgf("%s - kubeconfig is invalid!", mmpAndCluster)
		t.ErrorCount++
	}
}

func Test_ReadKubeconfig(t *testing.T) {
	path := "/Users/i544887/dev/repos/infra/bom-ctl/kubeconfigcheck/draft/allconfigs/eureka/itcm-dev/kubeconfig"
	cli, err := ReadKubeconfig(path)
	assert.Nil(t, err)
	bRes := CheckClientIsValid(cli)
	assert.True(t, bRes)
}

func ReadKubeconfig(inPath string) (*kubernetes.Clientset, error) {
	fi, err := infa.ReadFile(inPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cli, err := CreateCoreClientFromKubeconfig(fi)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cli, nil
}

func CheckClientIsValid(cli *kubernetes.Clientset) bool {
	res, err := cli.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err == nil && len(res.Items) > 0 {
		return true
	}
	if err == nil && len(res.Items) == 0 {
		log.Error().Msg("CheckClientIsValid got no err but len namespace is 0")
		return true
	}
	if err != nil {
		// log.Err(err).Msg("CheckClientIsValid got error.")
	}
	return false
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

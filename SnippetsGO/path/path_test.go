package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// 获得上一层的路径
func Test_GetUpPath(t *testing.T) {
	in := `a/b/c`
	in1 := filepath.Dir(in)
	in2 := filepath.Dir(in1)
	println(in, in1, in2) // output:a/b/c a/b a
}

// 获得执行文件的当前路径
func Test_CurPath(t *testing.T) {
	curPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println(curPath) // Asuka/SnippetsGO/path
}

// 获得项目根路径
func Test_CurProjPath(t *testing.T) {
	curPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	projName := `Asuka`
	paths := strings.Split(curPath, projName)
	execPath := fmt.Sprintf("%s%s/", paths[0], projName)
	println(execPath)
}

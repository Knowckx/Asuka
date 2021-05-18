package path

import (
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

// 路径的拼接
func Test_JoinPath(t *testing.T) {
	out := filepath.Join("a", "b")
	println(out) // a/b
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
	execPath := filepath.Join(paths[0], projName)
	println(execPath)
}

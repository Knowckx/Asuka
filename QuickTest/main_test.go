package sampletest

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/stretchr/testify/assert"
)

// 关于错误打印的总结
func Test_ForErrorPrint(t *testing.T) {
	err := GenErr02()
	fmt.Printf("%+v\n", err) // 测试里就这样打印。格式已是最佳。

	// zerolog 和 errors 包有点麻烦
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack // zerolog需要手动配置
	log.Info().Stack().Err(err).Send()                   // 正文中这么打印
}

func GenErr02() error {
	return errors.Errorf("OriError: %s", "err02") // 这样会带追踪
}

// 关于测试断言的总结
func Test_ForErrorAssert(t *testing.T) {
	err := GenErr02()
	// assert.NotNil(t, err)
	assert.EqualError(t, err, "OriError")

}

func GenErr01() error {
	err := GenErr02()
	return errors.WithStack(err)
}
func Test_MoreInputs(t *testing.T) {
	MoreInputs(1, 2, 3, 4)
}

func MoreInputs(a1 int, args ...int) {
	fmt.Println(a1)
	fmt.Println(args)
}

// err := errors.Errorf("whoops: %s", "foo")

func Test_PrintlnBool(t *testing.T) {
	var v1 bool
	fmt.Println(v1)
}

func Test_StartTest(t *testing.T) {
	StartTest()
}

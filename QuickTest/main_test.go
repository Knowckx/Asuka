package main

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func Test_Errors(t *testing.T) {
	err := GetErrorLayer2()
	fmt.Printf("%+v\n", err) // 必需是%+v 才能打印出stack信息
}

func GetErrorLayer2() error {
	err := GetBaseError()
	return errors.WithStack(err)
}

func GetErrorLayer1() error {
	err := GetBaseError()
	return errors.WithStack(err)
}

func GetBaseError() error {
	return fmt.Errorf("base error")
}

func Zlog() {
	// zerolog 和 errors 包有点麻烦
	// zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack // zerolog需要手动配置
	// log.Info().Stack().Err(err).Send()                   // 正文中这么打印
}

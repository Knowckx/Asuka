package sampletest

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func Test_ForOne(t *testing.T) {
	err := GenErr02()
	fmt.Printf("%+v\n", err)
}

func GenErr01() error {
	err := GenErr02()
	return errors.Wrap(err, "")
}

func GenErr02() error {
	// return fmt.Errorf("GenErr02")
	return errors.Errorf("whoops: %s", "foo")

}

// err := errors.Errorf("whoops: %s", "foo")

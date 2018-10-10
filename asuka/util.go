package asuka

import (
	"fmt"
)

// Printf auto \n
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}

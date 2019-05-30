package asuka

import (
	"fmt"
	"time"
)

var t1 time.Time

func Now() {
	t1 = time.Now()
}

func SinceSep(funcName string) time.Duration {
	elapsed := time.Since(t1)
	fmt.Printf("Elapsed Time:%s func name:%s", elapsed, funcName)
	t1 = time.Now()
	return elapsed
}

func Since() time.Duration {
	elapsed := time.Since(t1)
	fmt.Printf("Elapsed Time:%s", elapsed)
	t1 = time.Now()
	return elapsed
}

package asuka

import (
	"time"
)

var t1 time.Time

func Now() {
	t1 = time.Now()
}

func Since(funcName string) time.Duration {
	elapsed := time.Since(t1)
	Printf("Elapsed Time:%s func name:%s", elapsed, funcName)
	t1 = time.Now()
	return elapsed
}

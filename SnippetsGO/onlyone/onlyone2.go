package onlyone

import "fmt"

// mod 不可重入锁 | 使用channel作为并发控制入口，每次只允许一个协程进入临界区 | 重入报错
type OnlyOneCh chan bool

func GetOnlyOneCh() OnlyOneCh {
	one := make(OnlyOneCh, 1) // max limit = 1
	one <- true
	return one
}

var SyncError = fmt.Errorf("We are working on the task! please try again later")

func (in OnlyOneCh) TryLock() error {
	va := <-in
	in <- false
	if va == false {
		return SyncError
	}
	return nil
}

func (in OnlyOneCh) UnLock() {
	<-in
	in <- true
}

// mod 不可重入锁 End -----------------

package gosrc

type g struct {
	goid   int64  // 协程ID
	m      *m     // 目前的m
	status uint32 // 状态
	sched  gobuf  // 上下文信息
	stack  struct {
		lo uintptr // 该协程的栈低位
		hi uintptr // 该协程的栈高位
	}
	startfunc uintptr // 执行的函数地址
}

type m struct {
	procid uint64   // 底层的线程ID
	curg   *g       // 目标执行的G
	p      puintptr // 目标对应的P
}

type p struct {
	id     int32    // 执行单元ID
	status uint32   // 目前的状态
	m      muintptr // 目前绑定的M  idle状态就是空

	// 就绪状态下的协程队列. Accessed without lock.
	runq     [256]guintptr
	runqhead uint32
	runqtail uint32
}

// 调度器
type schedt struct {
	lock mutex // 操作全局队列的锁

	runq     gQueue // 全局可以运行的队列
	runqsize int32

	midle muintptr // 空闲的M
	pidle puintptr // 空闲的P
	gFree struct{} // 全局空闲的g
}

type guintptr struct{}
type puintptr struct{}
type gobuf struct{}
type muintptr struct{}
type mutex struct{}
type gQueue struct{}

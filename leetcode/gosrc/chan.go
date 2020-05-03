package gosrc

import "unsafe"

type hchan struct {
	qcount   uint           // len()
	dataqsiz uint           // cap()
	buf      unsafe.Pointer // 缓冲区的底层数组
	// elemtype *_type         // 表明channel的实际类型
	closed uint32 // 是否关闭状态close()

	elemsize uint16
	sendx    uint // 序号 send index
	recvx    uint // 序号 receive index
	// recvq    waitq // 读channel的协程队列
	// sendq    waitq // 写channel的协程队列

	lock mutex // 锁
}

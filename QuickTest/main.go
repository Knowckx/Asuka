package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/micro/go-micro/util/log"
)

func main() {
	StartTest()
}

// 模拟外部程序调用接口
func StartTest() {
	for {
		go Consumer()
		time.Sleep(1 * time.Second)
	}
}

func Consumer() {
	ok := DoMgr.AddOne()
	if !ok {
		log.Info("No Token.returned.")
		return
	}
	defer DoMgr.DoneOne()
	log.Info("Get Token.go.")
	time.Sleep(3 * time.Second)
	fmt.Println("Consumer End")
}

var DoMgr = NewMqCmr()

type MqCmr struct {
	data *sync.Map
}

func NewMqCmr() *MqCmr {
	out := &MqCmr{}
	out.data = &sync.Map{}
	return out
}

func (m *MqCmr) GetOrNewMgr(funcName string) *CountInfo {
	res, ok := m.data.Load(funcName)
	if ok {
		f, ok := res.(*CountInfo)
		if ok {
			return f
		} else {
			fmt.Errorf("type assertions failed")
		}
	}
	co := NewCountInfo(1)
	m.data.LoadOrStore(funcName, co)
	fmt.Printf("create limit for %s. Cap %d\n", funcName, 1)
	return co
}

// 准备干活
func (m *MqCmr) AddOne() bool {
	keyStr := printCallerName()
	res := m.GetOrNewMgr(keyStr)
	if res.IsFull() {
		return false
	}
	return res.AddLen()
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

// 干完了
func (m *MqCmr) DoneOne() {
	keyStr := printCallerName()
	res := m.GetOrNewMgr(keyStr)
	res.DoneOne()
}

type CountInfo struct {
	len    uint32
	cap    uint32
	isFull uint32 // 1 = full
	m      sync.Mutex
}

func NewCountInfo(capVal int) *CountInfo {
	out := &CountInfo{}
	out.cap = uint32(capVal)
	return out
}

func (c *CountInfo) IsFull() bool {
	return atomic.LoadUint32(&c.len) >= c.cap
}

func (c *CountInfo) AddLen() bool {
	c.m.Lock()
	defer c.m.Unlock()
	if c.IsFull() {
		return false
	}
	c.len++
	fmt.Println("use 1")
	return true
}

func (c *CountInfo) DoneOne() {
	c.m.Lock()
	defer c.m.Unlock()
	if c.len == 0 {
		return
	}
	c.len--
	fmt.Println("pop 1")
	return
}

var SyncError = fmt.Errorf("We are working on the task! please try again later")

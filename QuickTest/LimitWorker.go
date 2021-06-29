package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/micro/go-micro/util/log"
)

// 模拟外部程序调用接口
func StartTest() {
	for {
		go Consumer()
		time.Sleep(1 * time.Second)
	}
}

func Consumer() {
	ok := GlbMgr.AddOne()
	if !ok {
		log.Info("No Token.returned.")
		return
	}
	defer GlbMgr.Done()
	log.Info("Get Token.go.")
	time.Sleep(3 * time.Second)
	fmt.Println("Consumer End")
}

var GlbMgr = NewMqCmr()

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
	keyStr := getCallerName()
	res := m.GetOrNewMgr(keyStr)
	if res.IsFull() {
		return false
	}
	return res.Add()
}

func getCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

// 干完了
func (m *MqCmr) Done() {
	keyStr := getCallerName()
	res := m.GetOrNewMgr(keyStr)
	res.Done()
}

type CountInfo struct {
	len uint32
	cap uint32
	m   sync.Mutex
}

func NewCountInfo(capVal int) *CountInfo {
	out := &CountInfo{}
	if capVal < 1 {
		capVal = 1
	}
	out.cap = uint32(capVal)
	return out
}

func (c *CountInfo) IsFull() bool {
	return atomic.LoadUint32(&c.len) >= c.cap
}

func (c *CountInfo) Add() bool {
	if c.IsFull() {
		return false
	}
	c.m.Lock()
	defer c.m.Unlock()
	if c.IsFull() {
		return false
	}
	c.len++
	fmt.Println("use 1")
	return true
}

func (c *CountInfo) Done() {
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

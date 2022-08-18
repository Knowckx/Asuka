package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/Knowckx/Asuka/codemode/util"
)

// 内部缓存系统 - 代替全局变量
type CacheValue struct {
	Data       interface{}
	UpdateTime time.Time
	ExpireTime time.Duration

	Mu util.OnlyOneMu // once 内部更新本地缓存
}

// 缓存过期时间
func (in *CacheValue) IsValid() bool {
	return in.UpdateTime.Add(in.ExpireTime).After(time.Now())
}

// 刷新数据
func (in *CacheValue) resetUserData(newData interface{}) {
	in.Data = newData
	in.UpdateTime = time.Now()
}

// ----------------- 全局map start -----------------
var allCacheData = NewLocalCache()

type LocalCache struct {
	sync.Map                              // 存放数据 | 格式为[string - CacheValue]这样的K-V对
	InitOneMap map[string]*util.OnlyOneMu // 每个key的初始化锁,防止重复初始化
	InitOnce   sync.RWMutex               // key初始化锁对应的Map需要防止并发读写
}

func NewLocalCache() *LocalCache {
	out := &LocalCache{}
	out.InitOneMap = map[string]*util.OnlyOneMu{}
	return out
}

// 获取key对应的初始化锁
func (in *LocalCache) GetKeyInitMutex(key string) *util.OnlyOneMu {
	in.InitOnce.Lock()
	defer in.InitOnce.Unlock()

	mu, ok := in.InitOneMap[key]
	if ok {
		return mu //这个key对应的初始化锁已经生成过了
	}
	newMu := &util.OnlyOneMu{}
	in.InitOneMap[key] = newMu
	return newMu
}

/* ----------------- Error List Start ----------------- */
var (
	ErrorKeyEmpty    = fmt.Errorf("key is empty")
	ErrorDataOutTime = fmt.Errorf("Data in cache is out-of-time")
	ErrorLocked      = util.SyncError
)

/* ----------------- Error List End   ----------------- */

// ----------------- 全局map end-----------------

// 获取该key的唯一更新锁
func GetKeyMutex(keyName string) (*util.OnlyOneMu, error) {
	va, err := getLocalCacheValue(keyName)
	if err == nil { // ok
		return &va.Mu, nil
	}
	//need init this key
	if err == ErrorKeyEmpty {
		return InitKey(keyName)
	}
	return nil, err
}

// 新增key
func InitKey(keyName string) (*util.OnlyOneMu, error) {
	initMu := allCacheData.GetKeyInitMutex(keyName)

	if !initMu.TryLock() {
		return nil, util.SyncError
	}
	defer initMu.Unlock()

	va := &CacheValue{}
	va.ExpireTime = getDefaultExpireTime()
	va.Mu = util.OnlyOneMu{}
	allCacheData.Store(keyName, va)
	return &va.Mu, nil
}

// 更新数据
func RefreshKeyData(keyName string, newData interface{}) error {
	va, err := getLocalCacheValue(keyName)
	if err != nil {
		return err
	}
	// ok
	va.resetUserData(newData)
	return nil
}

// 取数据 UserData
func GetLocalCache(keyName string) (interface{}, error) {
	va, err := getLocalCacheValue(keyName)
	if err != nil {
		return nil, err
	}
	out := va.Data
	if !va.IsValid() { //已过期
		return nil, ErrorDataOutTime
	}
	return out, nil
}

// 取数据 CacheValue
func getLocalCacheValue(keyName string) (*CacheValue, error) {
	cacheVa, ok := allCacheData.Load(keyName)
	if !ok {
		return nil, ErrorKeyEmpty
	}
	va, ok := cacheVa.(*CacheValue)
	if !ok {
		return nil, fmt.Errorf("type assertions failed")
	}
	return va, nil
}

func getDefaultExpireTime() time.Duration {
	defaultMin := "5m"
	dur, err := time.ParseDuration(defaultMin)
	if err != nil {
		fmt.Printf("localcache.getDefaultExpireTime %s\n", err)
	}
	return dur
}

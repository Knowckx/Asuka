package SnippetsGO

import (
	"sync"

	"github.com/rs/zerolog/log"
)

/* 工作中经常有这样的需求，需要一个全局变量 或者 单例，用这个变量来记一些状态数据 -- 会有个问题，你的服务就不能多实例部署了*/

var (
	fInstance     *NotifyFilter // 用一个map来记录这个元素是不是已经出现过了
	fInstanceOnce sync.Once
)

// 入口
func NotifyFilterEntrence(accessId string) error {
	// each ID only use once
	if FilterAccessId(accessId) {
		log.Info().Msgf("filter ID. skip notify.", "id", accessId)
		return nil
	}
	return nil
}

type NotifyFilter struct {
	data map[string]string
}

// true: Repeated
func (t *NotifyFilter) IsRepeated(in string) bool {
	_, ok := t.data[in]
	if ok {
		return true
	}
	if len(t.data) > 100000 { // reset
		log.Info().Msgf("NotifyFilter sive is too big. reset.")
		t.data = make(map[string]string, 100000)
	}
	t.data[in] = ""
	log.Info().Msgf("NotifyFilter add new unit: %s. Now Cap: %d", in, len(t.data))
	return false
}

func NewNotifyFilter() *NotifyFilter {
	out := NotifyFilter{}
	out.data = make(map[string]string, 100000)
	return &out
}

func ResetfInstance() {
	fInstance = NewNotifyFilter()
}

func FilterAccessId(accessId string) bool {
	if fInstance == nil {
		fInstanceOnce.Do(ResetfInstance)
	}
	return fInstance.IsRepeated(accessId)
}

package main

import (
	"fmt"
	"time"
)

type NewObjMod struct {
	ID         int        `xorm:"ID"`
	UserID     int        `xorm:"UserID"`
	NickName   string     `xorm:"NickName"`
	Lang       string     `xorm:"Lang"`
	CreateTime *time.Time `xorm:"CreateTime"`
	MT4Account `xorm:"extends"`
}

func (in NewObjMod) GetMT4Account() *MT4Account {
	return &in.MT4Account
}

type NewObjMods []*NewObjMod

func (ins NewObjMods) GetMT4Accounts() MT4Accounts {
	outs := MT4Accounts{}
	for _, in := range ins {
		iIn := *in
		outs = append(outs, iIn.GetMT4Account())
	}
	return outs
}

//----------------- Mod MT4Data interface start -----------------
type IMT4AccData interface { //抽象，映射带mt4账号的数据
	GetMT4Account() *MT4Account
}

type MT4AccDatas []IMT4AccData //数据直接拼成这个

// 好处1,直接拿到数据里的mt4accs
func (ins MT4AccDatas) GetMT4Accounts() MT4Accounts {
	outs := MT4Accounts{}
	for _, in := range ins {
		iIn := in
		outs = append(outs, iIn.GetMT4Account())
	}
	return outs
}

// 好处2,组成map 方便索引
func (lis MT4AccDatas) ToMT4AccMap() MT4AccDataMap {
	rstMap := MT4AccDataMap{}
	for _, li := range lis {
		li := li
		mt4acc := li.GetMT4Account()
		rstMap[*mt4acc] = li
	}
	return rstMap
}

//抽象，映射mt4账户的map
type MT4AccDataMap map[MT4Account]IMT4AccData

//目标2:mt4账户的list
func (in MT4AccDataMap) GetKeys() MT4Accounts {
	keys := make([]*MT4Account, 0, len(in))
	for k := range in {
		key := k
		keys = append(keys, &key)
	}
	return keys
}

//目标3:原本data的list
func (in MT4AccDataMap) GetValues() MT4AccDatas {
	vs := make([]IMT4AccData, 0, len(in))
	for _, v := range in {
		v := v
		vs = append(vs, v)
	}
	return vs
}

//----------------- Mod MT4Data interface end -----------------

//----------------- Mod MT4Account Start -----------------
type MT4Account struct {
	BrokerID int32  `xorm:"'BrokerID'"`
	Account  string `xorm:"'Account'"`
}

type MT4Accounts []*MT4Account

func (in MT4Account) String() string {
	return fmt.Sprintf("MT4account:BrokerID %d Account %s", in.BrokerID, in.Account)
}

//----------------- Mod MT4Account End -----------------

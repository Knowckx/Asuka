package main

import (
	"fmt"
	"reflect"
	"time"
)

type ErrorLogOperator struct {
	ErrorLog *FollowOrderErrorLog `xorm:"extends"`
	Operator string               `xorm:"<-"`
}

type FollowOrderErrorLog struct {
	ID             int32     `xorm:"'ID' pk autoincr"`
	TraderUserID   int32     `xorm:"'TraderUserID'"`
	TraderAccount  string    `xorm:"'TraderAccount'"`
	TraderBrokerID int32     `xorm:"'TraderBrokerID'"`
	TraderOrderID  int32     `xorm:"'TraderOrderID'"`
	UserID         int32     `xorm:"'UserID'"`
	Account        string    `xorm:"'Account'"`
	BrokerID       int32     `xorm:"'BrokerID'"`
	OrderID        int32     `xorm:"'OrderID'"`
	ErrorCode      int32     `xorm:"'ErrorCode'"`
	ErrorMessage   string    `xorm:"'ErrorMessage'"`
	Status         int32     `xorm:"'Status'"`
	CreateTime     time.Time `xorm:"'CreateTime' timestamp created"`
}

func (f *FollowOrderErrorLog) ToStrSlice() []string {
	ss := []string{}
	v := reflect.ValueOf(f).Elem()
	for i := 0; i < v.NumField(); i++ {
		str := fmt.Sprintf("%v", v.Field(i).Interface())
		ss = append(ss, str)

	}
	return ss
}

func NewData() []*ErrorLogOperator {
	outs := []*ErrorLogOperator{}
	for i := 0; i < 999; i++ {
		out := &ErrorLogOperator{}
		out.Operator = "AA"
		errlog := &FollowOrderErrorLog{}
		errlog.ID = int32(i)
		errlog.Status = -20
		errlog.ErrorMessage = "None"
		errlog.CreateTime = time.Now()

		out.ErrorLog = errlog
		outs = append(outs, out)
	}
	return outs
}

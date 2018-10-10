package fmdebug

import (
	"fmt"

	"gitlab.followme.com/CopyTradingGo/monitor-srv/src/handler"
	pbmo "gitlab.followme.com/CopyTradingGo/proto/src/monitor"
	pbpage "gitlab.followme.com/Followme/proto/page"

	"golang.org/x/net/context"
)

func args_AddOperationLog(h *handler.MonitorSrv) {
	req := &pbmo.OperateLog{
		ID:            100,
		Operator:      "",
		Operation:     "Fix Sam",
		OperateStatus: "Fail",
		OperateTime:   101010001010,
	}
	rsp := new(pbmo.Empty)
	err := h.AddOperationLog(context.Background(), req, rsp)
	fmt.Printf("err:\n%v\nrsp:\n%#v", err, rsp)
}

func args_PageOperationLog(h *handler.MonitorSrv) {
	req := &pbmo.PageOperateLogRequest{
		Operator:         "",
		Operation:        "",
		OperateTimeBegin: 1538038998,
		OperateTimeEnd:   1538117010,
		Page: &pbpage.PageOption{
			Page:     0,
			PageSize: 15,
		},
	}
	rsp := new(pbmo.PageOperateLogReply)
	err := h.PageOperationLog(context.Background(), req, rsp)
	fmt.Printf("err:\n%v\nrsp:\n%#v", err, rsp)
}

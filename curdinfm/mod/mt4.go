package mod

import "github.com/Knowckx/Asuka/curdinfm/pb"

// 依赖mod
type MT4Account struct {
	BrokerID int32
	Account  string
}

func (in *MT4Account) ToProto() *pb.MT4Account {
	out := &pb.MT4Account{}
	out.BrokerID = in.BrokerID
	out.Account = in.Account
	return out
}

func NewMT4Account(BrokerID int32, Account string) *MT4Account {
	out := &MT4Account{}
	out.BrokerID = BrokerID
	out.Account = Account
	return out
}

type MT4Accounts []*MT4Account

type MT4AccData map[MT4Account]*NewObjMod

//----------------- Mod PageOp Start -----------------
type PageOp struct {
	OrderBy   string
	IsDESC    bool //true:倒序排列
	PageIndex int
	PageSize  int
}

func NewPageOpPB(page *pb.PageOp) *PageOp {
	if page == nil {
		return nil
	}
	p := &PageOp{}
	p.OrderBy = page.OrderBy
	p.IsDESC = page.IsDESC
	p.PageIndex = int(page.PageIndex)
	p.PageSize = int(page.PageSize)
	return p
}

//----------------- Mod PageOp End -----------------

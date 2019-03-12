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

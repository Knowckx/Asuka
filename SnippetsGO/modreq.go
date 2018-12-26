package SnippetsGO

import "github.com/Knowckx/Asuka/SnippetsGO/pbmod"

// 本地的req
type TestReq struct {
	UserID   int32
	NickName string
	Page     *PageOp
}

func (req *TestReq) NewFromProto(in *pbmod.TestReq) {
	if in == nil {
		return
	}
	req.UserID = in.UserID
	req.NickName = in.NickName
}

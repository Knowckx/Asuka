package pbmod

type pbPage struct {
	idx int
}

type TestReq struct {
	UserID   int32
	NickName string
}

// 映射proto
type pbMod struct {
	UserID int
	Text   string
	Page   pbPage
}

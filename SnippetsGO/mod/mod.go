package mod

type MT4Account struct {
	BrokerID int32  `bson:"brokerid" json:"B"`
	Account  string `bson:"login" json:"A"`
}

type AccInfo struct {
	MT4Account
	UserID int
	Score  int //缺这个字段
}

type AccScore struct {
	MT4Account
	Score int //有这个字段
}

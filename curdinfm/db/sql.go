package db

import (
	"fmt"

	"github.com/Knowckx/Asuka/curdinfm/mod"
	"github.com/xormplus/xorm"
)

// 拼SQL  fmt 裸语句   XX and (a or b or c)
func GetS4FollowProfits(tx *xorm.Session, accs mod.MT4Accounts) (mod.MT4Accounts, error) {
	out := mod.MT4Accounts{}
	if len(accs) == 0 {
		return out, nil
	}

	sqlStr := `SELECT fo.TraderBrokerID as BrokerID,fo.TraderAccount as Account,sum(td.profit+td.swaps+td.commission) as fProfit
	FROM t_followorder fo INNER JOIN  t_trades td
	on td.BrokerID = fo.BrokerID and td.Account = fo.Account and td.TradeID = fo.TradeID `

	whereSub := fmt.Sprintf("WHERE td.OpenTime > '%s' and td.CloseTime< '%s'", "2019-01-01", "2019-03-01")

	for i, acc := range accs {
		if i == 0 {
			whereSub = fmt.Sprintf("%s and (fo.TraderBrokerID = %d and fo.TraderAccount = %s)", whereSub, acc.BrokerID, acc.Account)
			continue
		}
		whereSub = fmt.Sprintf("%s or (fo.TraderBrokerID = %d and fo.TraderAccount = %s)", whereSub, acc.BrokerID, acc.Account)
	}

	sqlStr = fmt.Sprintf("%s %s GROUP BY fo.TraderBrokerID,fo.TraderAccount", sqlStr, whereSub)
	// fmt.Println(sqlStr)
	err := tx.SQL(sqlStr).Find(&out)

	return out, err
}

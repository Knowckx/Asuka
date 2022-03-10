package crud

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	infa "github.tools.sap/aeolia/in-fa"
)

/*
	gorm 把常用的CRUD都写成例子
	CRUD = create, read, update and delete.
*/

// read - one line
func GetSecretsHash() (string, error) {
	sec := &AzureClientSecret{}
	res := defaultDB.Where("name = ?", "123").First(sec)
	if res.Error != nil {
		return "", errors.WithStack(res.Error)
	}
	return sec.ClientSecret, nil
}

// create
func Inserts(ins []*AzureClientSecret) {
	if !CanWriteDB {
		log.Info().Msg("CanWriteDB is false.Skip insert.")
		fmt.Println(ins)
		// infa.PrintYaml(ins)
		return
	}

	leng := len(ins)
	if leng == 0 {
		log.Info().Msg("Insert return.len is 0")
		return
	}
	err := inserts(&ins)
	if err != nil {
		log.Err(err).Msg("ClientSecretInsertTo failed")
		infa.PrintJson(ins)
		return
	}
	startID := ins[0].ID
	endID := ins[leng-1].ID
	log.Info().Msgf("Inserts success. Total:%d. ID range [%d-%d]", leng, startID, endID)
}

func inserts(data interface{}) error {
	res := defaultDB.Create(data)
	if res.Error != nil {
		return errors.WithStack(res.Error)
	}
	af := res.RowsAffected
	log.Info().Msgf("Insert affectd lines %d", af)
	return nil
}

// update
func UpdateAzureClientSecret(in *AzureClientSecret) error {
	res := defaultDB.Model(&AzureClientSecret{}).Where("subscription_name = ?", in.SubscriptionName).Update("ClientSecret", in.ClientSecret)
	if res.Error != nil {
		return errors.WithStack(res.Error)
	}
	af := res.RowsAffected
	log.Info().Msgf("Insert affectd lines %d", af)
	return nil
}

// update or insert
func ClientSecretCreateOrUpdate(ins []*AzureClientSecret) {
	if !CanWriteDB {
		log.Info().Msg("CanWriteDB is false.Skip insert.")
		// infa.PrintYaml(ins)
		return
	}
	insertsLis := []*AzureClientSecret{}
	for _, in := range ins { // Update
		upAff := defaultDB.Model(&AzureClientSecret{}).Where("subscription_name = ?", in.SubscriptionName).Updates(in).RowsAffected
		if upAff == 0 {
			insertsLis = append(insertsLis, in)

		}
	}
	Inserts(insertsLis) // Insert
}

// get 自定义select
type ResGroupNameCost struct {
	ResGroupName string
	Cost         float64
}

func (t *ResGroupNameCost) TableName() string {
	return "azure_cost_report"
}

func GetDailyResGroupCost(in *AzureClientSecret) []*ResGroupNameCost {
	dayData := []*ResGroupNameCost{}
	tx := defaultDB.Where("date_type = ?", in.ClientID)
	tx.Group("res_group_name").Select("res_group_name, sum(cost)::NUMERIC as cost") // PG money type shoube convert to NUMERIC
	tx.Find(&dayData)
	return dayData
}

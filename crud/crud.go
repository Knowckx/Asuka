package crud

import (
	"fmt"

	infa "github.com/Knowckx/ainfa"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

/*
	gorm 把常用的CRUD都写成例子
	CRUD = create, read, update and delete.
*/

// Get - one line
func GetSecretsHash() (string, error) {
	sec := &AzureClient{}
	res := GetDefault().Where("name = ?", "123").First(sec)
	if res.Error != nil {
		return "", errors.WithStack(res.Error)
	}
	return sec.ClientSecret, nil
}

// Get - All
func GetRegisterClusters() ([]*AzureClient, error) {
	datas := []*AzureClient{}
	res := GetDefault().Where("id < ?", 10).Find(&datas)
	if res.Error != nil {
		return nil, errors.WithStack(res.Error)
	}
	return datas, nil
}

// count(*)
func KubeconfigGetCount() int {
	var count int
	op := GetDefault().Raw("SELECT count(*) FROM bom_kubeconfig").Scan(&count)
	if op.Error != nil {
		log.Err(op.Error).Msg("")
	}
	return count
}

// get 自定义select
type ResGroupNameCost struct {
	ResGroupName string
	Cost         float64
}

func (t *ResGroupNameCost) TableName() string {
	return "azure_cost_report"
}

func GetDailyResGroupCost(in *AzureClient) []*ResGroupNameCost {
	dayData := []*ResGroupNameCost{}
	tx := GetDefault().Where("date_type = ?", in.ClientID)
	tx.Group("res_group_name").Select("res_group_name, sum(cost)::NUMERIC as cost") // PG money type shoube convert to NUMERIC
	tx.Find(&dayData)
	return dayData
}

// create
func Inserts(ins []*AzureClient) {
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
	res := GetDefault().Create(data)
	if res.Error != nil {
		return errors.WithStack(res.Error)
	}
	af := res.RowsAffected
	log.Info().Msgf("Insert affectd lines %d", af)
	return nil
}

// update
func UpdateAzureClient(in *AzureClient) error {
	res := GetDefault().Model(&AzureClient{}).Where("subscription_name = ?", in.SubsName).Update("ClientSecret", in.ClientSecret)
	if res.Error != nil {
		return errors.WithStack(res.Error)
	}
	af := res.RowsAffected
	log.Info().Msgf("Insert affectd lines %d", af)
	return nil
}

// update or insert
func ClientUpdateOrInsert(ins []*AzureClient) {
	if !CanWriteDB {
		log.Info().Msg("CanWriteDB is false.Skip insert.")
		// infa.PrintYaml(ins)
		return
	}
	if len(ins) == 0 {
		return
	}
	insertsLis := []*AzureClient{}
	for _, in := range ins { // Update
		op := GetDefault().Where("name = ?", in.SubsName).Updates(in)
		if op.Error != nil {
			log.Err(op.Error).Str("name", in.SubsName).Msg("do update error")
			continue
		}
		if op.RowsAffected == 0 {
			insertsLis = append(insertsLis, in)
		}
	}
	updatedCount := len(ins) - len(insertsLis)
	log.Info().Msgf("UpdateOrInsert. Total len %d. Updated %d. Need insert rows %d.", len(ins), updatedCount, len(insertsLis))
	if len(insertsLis) > 0 {
		inserts(insertsLis)
	}
}

// delete one
func KubeconfigDeleteOne(sID string) error {
	// Unscoped 表示硬删除
	res := GetDefault().Unscoped().Where("server_id = ?", sID).Delete(&AzureClient{})
	if res.Error != nil {
		return errors.WithStack(res.Error)
	}
	log.Debug().Msgf("delete server_id = %s, Affected Rows %d", sID, res.RowsAffected)
	return nil
}

// delete list
func KubeconfigDelete(ins []*AzureClient) {
	deleted := 0
	for _, in := range ins {
		err := KubeconfigDeleteOne(in.ClientID)
		if err != nil {
			log.Err(err).Msg("KubeconfigDelete Error")
		}
		deleted++
	}
	log.Info().Msgf("delete ClientSecret. want delete %d. have deleted %d", len(ins), deleted)
}

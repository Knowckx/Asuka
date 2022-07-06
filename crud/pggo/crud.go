package pggo

import (
	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
)

/* go-pg 相关的CURD */

// Query by id
func GetRecordByID(ID int64) *User {
	tx := GetPgDb()
	r := User{}
	//language=sql
	_, err := tx.QueryOne(&r, `select * from tableAA where id=?`, ID)
	if err == pg.ErrNoRows {
		log.Error().Msgf("Query id %d: not found in db", ID)
		return nil
	}
	if err != nil {
		log.Err(err).Msgf("Query id %d: get error", ID)
		return nil
	}
	return &r
}

type User struct {
}

func GetPgDb() *pg.DB {
	out := &pg.DB{}
	return out
}

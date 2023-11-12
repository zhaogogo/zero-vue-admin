package types

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func NewMysql() (*gorm.DB, error) {
	dsn := "root:zhaO..123@tcp(127.0.0.1:3306)/t1?collation=utf8mb4_general_ci&parseTime=false&loc=Asia%2FShanghai&timeout=2s&readTimeout=1s&writeTimeout=1s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestTtt(t *testing.T) {
	db, err := NewMysql()
	if err != nil {
		t.Error(err)
		return
	}
	db = db.Debug()
	slience := []SlienceName{}
	//sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	//	return tx.Model(Host{Id: 4, Host: "127.0.0.1"}).Association("SlienceNames").
	//})
	//fmt.Println(sql)
	err = db.Model(&Host{Id: 4, Host: "127.0.0.1"}).Session(&gorm.Session{FullSaveAssociations: true}).Association("SlienceNames").Unscoped().Replace(&[]SlienceName{
		SlienceName{
			Id:          26,
			HostID:      4,
			SlienceName: "up_prometheus",
			Default:     true,
			To:          2,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	bytes, err := json.MarshalIndent(slience, "", "\t")
	fmt.Println(string(bytes), err)
}

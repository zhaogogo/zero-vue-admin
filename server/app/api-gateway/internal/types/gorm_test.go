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
	alertSlience := []SlienceName{}
	err = db.Where(SlienceName{HostID: 1}).Preload("Matchers", "host_id = ?", 1).Find(&alertSlience).Error
	if err != nil {
		t.Error(err)
		return
	}

	bytes, err := json.MarshalIndent(alertSlience, "", "\t")
	if err != nil {
		t.Error(bytes)
		return
	}
	fmt.Print(string(bytes))
}

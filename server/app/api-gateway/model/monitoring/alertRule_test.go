package monitoring

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestTtt(t *testing.T) {
	db, err := NewMysql()
	if err != nil {
		t.Error(err)
		return
	}
	err = db.Debug().AutoMigrate(Host{}, HostTag{})
	if err != nil {
		t.Error(err)
	}
	var hosts []Host
	err = db.Find(&hosts).Order("id").Limit(10).Offset(0).Association("Tags").Error
	t.Log(err)
	bytes, err := json.MarshalIndent(hosts, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	//db.Find(&hosts).Order("id").Limit(10).Offset(0).Association("Tags").Error
}

func NewMysql() (*gorm.DB, error) {
	dsn := "root:zhaO..123@tcp(127.0.0.1:3306)/t1?collation=utf8mb4_general_ci&parseTime=false&loc=Asia%2FShanghai&timeout=2s&readTimeout=1s&writeTimeout=1s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

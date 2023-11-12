package monitoring

import (
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
	db = db.Debug()
	//db.AutoMigrate(&Host{}, &HostTag{})
	//err = db.Save(&Host{Host: "test", To: 1}).Error
	err = db.Delete(&Host{ID: 1}).Error

	t.Log(err)
	//db.Find(&hosts).Order("id").Limit(10).Offset(0).Association("Tags").Error
}

func NewMysql() (*gorm.DB, error) {
	dsn := "root:zhaO..123@tcp(127.0.0.1:3306)/test?collation=utf8mb4_general_ci&parseTime=false&loc=Asia%2FShanghai&timeout=2s&readTimeout=1s&writeTimeout=1s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return nil, err
	}
	return db, nil
}

package mm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := "root:zhaO..123@tcp(127.0.0.1)/t1?collation=utf8mb4_general_ci&parseTime=true&loc=Asia%2FShanghai&timeout=2s&readTimeout=1s&writeTimeout=1s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		t.Error(err)
		return
	}
	db = db.Debug()
	err = db.AutoMigrate(&AlertRules{}, &AlertRuleTag{}, &AlertRuleLabel{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}

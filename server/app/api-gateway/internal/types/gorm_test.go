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
	bodyString := "{\"id\":2,\"host\":\"192.168.14.102\",\"sliences\":[{\"id\":3,\"host_id\":2,\"slience_name\":\"a\",\"default\":true,\"to\":1,\"matchers\":[{\"id\":6,\"host_id\":2,\"slience_name_id\":3,\"name\":\"env\",\"value\":\"aaa\",\"is_regex\":false,\"is_equal\":true},{\"id\":10,\"host_id\":2,\"slience_name_id\":3,\"name\":\"instance\",\"value\":\"{{.Host}}\",\"is_regex\":false,\"is_equal\":true}]},{\"id\":4,\"host_id\":2,\"slience_name\":\"b\",\"default\":false,\"to\":1,\"matchers\":[{\"id\":8,\"host_id\":2,\"slience_name_id\":4,\"name\":\"env\",\"value\":\"bbb\",\"is_regex\":false,\"is_equal\":true},{\"id\":11,\"host_id\":2,\"slience_name_id\":4,\"name\":\"instance\",\"value\":\"{{.Host}}\",\"is_regex\":false,\"is_equal\":true}]},{\"host_id\":2,\"slience_name\":\"c\",\"default\":false,\"to\":1,\"matchers\":[{\"host_id\":2,\"name\":\"env\",\"value\":\"ccc\",\"is_regex\":false,\"is_equal\":true},{\"host_id\":2,\"name\":\"instance\",\"value\":\"{{.Host}}\",\"is_regex\":false,\"is_equal\":true}]}]}"
	body := SliencePutRequest{}
	err = json.Unmarshal([]byte(bodyString), &body)
	if err != nil {
		t.Error(err)
		return
	}
	bytes, err := json.MarshalIndent(body, "", "\t")
	if err != nil {
		t.Error(bytes)
		return
	}
	fmt.Print(string(bytes))
	err = db.Transaction(func(db *gorm.DB) error {
		err = db.Model(&Host{Id: uint64(body.ID), Host: body.Host}).Association("Sliences").Replace(body.Sliences)
		if err != nil {
			t.Error(err)
			return err
		}
		for _, slience := range body.Sliences {
			fmt.Println(slience)
			s := slience
			err := db.Model(&s).Association("Matchers").Replace(slience.Matchers)
			if err != nil {
				t.Error(err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

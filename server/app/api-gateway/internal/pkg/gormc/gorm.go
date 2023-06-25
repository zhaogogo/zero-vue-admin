package gormc

import (
	"database/sql"
	"fmt"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/config"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"time"
)

func MustNewGrom(conns config.Mysql) *gorm.DB {
	DB, err := NewGorm(conns)
	if err != nil {
		panic(fmt.Sprintf("gorm create connect error: %w", err))
	}
	return DB
}

func NewGorm(conns config.Mysql) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: MustSqlConn(conns.System.DataSource)}), &gorm.Config{
		Logger: newGormLogger(conns.LogPath),
	})
	if err != nil {
		return nil, err
	}
	db.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:           []gorm.Dialector{mysql.New(mysql.Config{Conn: MustSqlConn(conns.Monitoring.DataSource)})},
			Replicas:          nil,
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		}, "monitoring"),
	)

	return db, nil
}

func MustSqlConn(dsn string) *sql.DB {
	conn, err := newSqlConn(dsn)
	if err != nil {
		panic(fmt.Sprintf("创建链接失败, err:%w", err))
	}
	return conn
}

func newSqlConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(64)
	db.SetConnMaxLifetime(2 * time.Minute)
	return db, nil
}

func newGormLogger(path string) logger.Interface {
	return logger.New(log.New(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    1024,
		MaxAge:     2,
		MaxBackups: 7,
		Compress:   false,
	}, "", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second,
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      true,
		LogLevel:                  logger.Silent,
	})
}

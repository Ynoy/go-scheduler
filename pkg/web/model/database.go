package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go-scheduler/pkg/init/database"
	"go-scheduler/pkg/web/conf"
	"time"
	"xorm.io/core"
)

const DefaultTimeFormat = "1970-01-01 00:00:00"

type (
	Seeder interface {
		Seed() error
	}
	Model interface {
		Store() error
		Update() error
		ToString() (string, error)
	}
)

var Engine *xorm.Engine

func Connection() (*xorm.Engine, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/charset=%s",
		conf.Conf.Database.User,
		conf.Conf.Database.Pass,
		conf.Conf.Database.Host,
		conf.Conf.Database.Port,
		conf.Conf.Database.Char,
	)
	engine, err := xorm.NewEngine(string(database.MYSQL), dsn)
	if engine != nil {
		engine.SetMaxIdleConns(10)
		engine.SetMaxOpenConns(30)
		engine.SetLogLevel(core.LOG_INFO)
		engine.SetConnMaxLifetime(30 * time.Second)
	}
	// 协程保持心跳
	go keepAlive()
	return engine, err
}

// 实现心跳
func keepAlive() {
	t := time.Tick(60 * time.Second)
	for {
		<-t
		if err := Engine.Ping(); err != nil {
			// Todo
		}
	}
}

// 数据库迁移
func Migrate() error {
	table := []interface{}{
		&User{},
		&Node{},
		&Log{},
		&PasswordResets{},
		&Pipeline{},
		&PipelineRecord{},
		&PipelienTaskPivot{},
		&PipelineNodePivot{},
		&Task{},
		&TaskRecord{},
	}

	if err := Engine.DropTables(table...); err != nil {
		return err
	}
	if err := Engine.Charset("utf8mb4").Sync2(table...); err != nil {
		return err
	}
	return nil
}

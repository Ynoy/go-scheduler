package database

import (
	"database/sql"
	"fmt"
	"go-scheduler/pkg/web/conf"
	"log"
)

var (
	DB  *sql.DB
	err error
)

type Type string

const (
	MYSQL Type = "mysql"
)

func InitDatabase() {
	// 构造数据库连接地址
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/charset=%s",
		conf.Conf.Database.User,
		conf.Conf.Database.Pass,
		conf.Conf.Database.Host,
		conf.Conf.Database.Port,
		conf.Conf.Database.Char,
	)
	DB, err = sql.Open(string(MYSQL), dsn)
}

func CheckDatabase() bool {
	InitDatabase()
	defer func() {
		if err := DB.Close(); err != nil {

		}
	}()

	statement := fmt.Sprintf("SHOW DATABASES LIKE '%s'", conf.Conf.Database.Name)

	var (
		rows     *sql.Rows
		database string
	)

	rows, err = DB.Query(statement)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&database); err != nil {
			log.Println(err)
		}
		if database == conf.Conf.Database.Name {
			return true
		}
	}
	return false
}

func CreateDatabase() error {
	InitDatabase()
	statement := fmt.Sprintf("CREATE DATABASE IF NOT EXIST %s DEFAULT CHARACTER SET %s DEFAULT COLLATE %s",
		conf.Conf.Database.Name,
		conf.Conf.Database.Char,
		"utf8mb4_unicode-ci",
	)
	_, err := DB.Query(statement)
	return err
}

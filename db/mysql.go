package db

import (
	// "fmt"
)

type MysqlOptions struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Db     string   // Database name
	Tables []string // If empty of nil, then all tables are fetched
}


func Backup(opts MysqlOptions) {

}
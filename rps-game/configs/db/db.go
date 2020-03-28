package db

import "time"

// Configurations of MySql connection pool
const (
	MysqlMaxOpenConns = 120
	MysqlMaxIdleConns = 20
	MysqlConnMaxLifetime = 12 * time.Hour
)

package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ssidletsky/esportal-twitter/app/config"
)

// Conn is a database connection
var Conn *sql.DB

// Initialize initializes mysql connection
func Initialize(cnf config.MySQL) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.DBname,
	)
	var err error
	Conn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Conn.Ping(); err != nil {
		panic(err)
	}

	Conn.SetConnMaxLifetime(time.Minute * cnf.MaxLifetime)
	Conn.SetMaxOpenConns(cnf.MaxOpenConns)
	Conn.SetMaxIdleConns(cnf.MaxIdleConns)
}

// Shutdown gracefully closes the connection
func Shutdown() error {
	return Conn.Close()
}

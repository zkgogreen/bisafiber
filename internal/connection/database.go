package connection

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/zkgogreen/bisago/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase(conf config.Database) *sql.DB {
	// MySQL DSN format: username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	// Use loc parameter for timezone (MySQL DSN doesn't support time_zone parameter)
	encodedLoc := url.QueryEscape(conf.Tz)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", 
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Name, conf.Charset, conf.ParseTime, encodedLoc)
	
	db, err := sql.Open(conf.Driver, dsn)
	if err != nil {
		log.Fatal("Error opening database connection", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database", err.Error())
	}
	return db
}

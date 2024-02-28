package configuration

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDb() *sqlx.DB {
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	host := viper.GetString("db.host")
	dba := viper.GetString("db.db")
	connStr := fmt.Sprintf("postgres://%v:%v@%v/%v", user, password, host, dba)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	return db
}
func InitTimeZone() {
	thaiTime, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
		panic("system none time")
	}
	time.Local = thaiTime
}

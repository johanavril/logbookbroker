package database

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dataSource struct {
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Dbname   string `yaml:"dbname"`
}

var instance *sqlx.DB
var once sync.Once

func (ds dataSource) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s", ds.User, ds.Password, ds.Host, ds.Dbname)
}

func getDataSource() (*dataSource, error) {
	ds := dataSource{}
	dburl := os.Getenv("DATABASE_URL")
	dburl = dburl[11:]

	s := strings.Split(dburl, "@")
	credential := strings.Split(s[0], ":")
	dbcon := strings.Split(s[1], "/")
	socket := strings.Split(dbcon[0], ":")

	ds.User = credential[0]
	ds.Password = credential[1]
	ds.Host = socket[0]
	ds.Dbname = dbcon[1]

	return &ds, nil
}

func new() (*sqlx.DB, error) {
	ds, err := getDataSource()
	if err != nil {
		return nil, err
	}

	return sqlx.Connect("postgres", ds.String())
}

func Get() *sqlx.DB {
	once.Do(func() {
		var err error
		instance, err = new()
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}

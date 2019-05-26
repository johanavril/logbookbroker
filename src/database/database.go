package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
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
	env := os.Getenv("APP_ENV")
	yamlFile, err := ioutil.ReadFile("../config/database/" + env + ".yml")
	if err != nil {
		return nil, err
	}

	ds := dataSource{}

	if err := yaml.Unmarshal(yamlFile, &ds); err != nil {
		return nil, err
	}

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

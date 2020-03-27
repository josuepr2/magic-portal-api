package helpers

import (
	"fmt"
	"log"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var dbs = map[string] *gorm.DB {
	"defaultDB": nil,
}

const (
	DEFAULTDB_CON = "defaultDB"
)

func GetDB(dbType string) (*gorm.DB, error) {
	err := connect(dbType)
	if err != nil {
		return nil, err
	}

	dbs[dbType].DB().SetConnMaxLifetime(time.Minute * 5)
	dbs[dbType].DB().SetMaxIdleConns(50)
	dbs[dbType].DB().SetMaxOpenConns(300)
	return dbs[dbType], nil
}


func connect(dbType string) error {
	var err error

	if dbType == DEFAULTDB_CON {
		err = DefaultConnection()
	}
	return err
}

func DefaultConnection() error {

	if dbs[DEFAULTDB_CON] != nil {
		return nil
	}

	var err error
	var connString string

	log.Println("conecting with data base")

	connString = fmt.Sprintf(
		"host=%s dbname=%s port=%d user=%s ",
		"ec2-50-17-90-177.compute-1.amazonaws.com",
		"dbr96altfcla6d",
		5432,
		"rfmtoblmdnhweo")

	connString = connString + " password=2bd73aef1c6d5b0950594c94414519a13e587aae63b6d82f5d051f49b2e0740a"
	connString = connString + " sslmode=require"

	dbs[DEFAULTDB_CON], err = gorm.Open("postgres", connString)
	if err != nil {
		log.Println("Connection error in Extra: ", err)
	}

	return nil
}

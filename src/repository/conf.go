package repository

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Pg is a client of database
type Pg struct {
	*gorm.DB
}

// instance of connection to database
var instance *Pg

var doOnce sync.Once

//Connect return only one instance
func Connect() *Pg {
	doOnce.Do(func() {
		instance = NewDBClient("localhost", "5432", "postgres", "demo-go", "root")
	})
	return instance
}

// NewDBClient creates a Postgres client
func NewDBClient(host string, port string, user string, dbName string, password string) *Pg {

	db, _ := gorm.Open(
		"postgres",
		"host="+host+
			" port="+port+
			" user="+user+
			" dbname="+dbName+
			" password="+password)

	return &Pg{db}
}

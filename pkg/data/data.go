package data

import (
	"database/sql"
	"flag"
	"log"
	"os"

	// imported to act as a driver for the db.
	_ "github.com/mattn/go-sqlite3"
)

// assets dir is a package scoped variable. It is accessed before the env
// is established and therefore needs to be accessed through a lookup and
// used in the database initalization throughtout this package.
var assetsDir string

// Datastore is an interface with data access methods that will
// allow persistence.
type Datastore interface {
	CrossSellMachines(int) ([]Pod, error)
	CrossSellPods(int) ([]Pod, error)
	initDB() error
	Machines(int) ([]CoffeeMachine, error)
	Pods(int, int) ([]Pod, error)
	insertPod(string) error
	insertMachine(string) error
	parseProducts() error
}

// DB is type which will implement Datastore and will be the receiver
// for our data access functions.
type DB struct {
	*sql.DB
}

// NewDB returns an instance of the DB struct.
func NewDB() (*DB, error) {
	// gets assetsDir flag.
	assetsDir = flag.Lookup("assetsDir").Value.(flag.Getter).Get().(string)
	// Recreates data on build.
	if err := os.Remove(assetsDir + "/data.db"); err != nil {
		log.Printf("either the this is this first time creating the db file or the assetDir is not set properly")
	}

	db, err := sql.Open("sqlite3", assetsDir+"/data.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	newDB := &DB{db}
	if err = newDB.initDB(); err != nil {
		return nil, err
	}
	return newDB, nil
}

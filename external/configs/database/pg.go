package database

import (
	"database/sql"
	"log"
	"sync"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/env"
	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func DbConn() *sql.DB {
	once.Do(func() {
		var err error
		env := env.GetEnv()
		if db, err = sql.Open("postgres", env.DATABASE_URL); err != nil {
			log.Panic(err)
		}
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
	})
	return db
}

func DbClose() {
	db.Close()
}

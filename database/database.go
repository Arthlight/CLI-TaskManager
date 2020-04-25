package database

import (
	bolt "go.etcd.io/bbolt"
	"log"
)

func connect() *bolt.DB {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AddTask() {
	d := connect()
	defer d.Close()

}
package database

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
)

func connect() *bolt.DB {
	db, err := bolt.Open("TaskManager.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AddTask(s string) {
	fmt.Println("in addTask in database.go")
	d := connect()
	defer func() {
		err := d.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("successfully connected to database in addTask in database.go")
	err := handleTask(d, s)
	if err != nil {
		log.Fatal(err)
	}

}

func handleTask(d *bolt.DB, s string) error {
	fmt.Println("in handleTask in database.go")
	err := d.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		v := b.Get([]byte("Tasks"))
		if v == nil {
			err := updateTasks(d, s)
			if err != nil {
				return err
			}
		} else {
			err := updateTasks(d, s + " " + string(v))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}


func updateTasks(d *bolt.DB, s string) error {
	fmt.Println("in updateTask in database.go")
	err := d.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("Tasks"), []byte(s))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("returned in updateTasks in database.go")
	return nil
}
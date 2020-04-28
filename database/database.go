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

	err := handleTask(d, s)
	if err != nil {
		log.Fatal(err)
	}

}

func handleTask(d *bolt.DB, s string) error {
	fmt.Println("in handleTask in database.go")
	err := d.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyTasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		fmt.Println(s)
		err = b.Put([]byte(s), []byte(s))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(s string) {
	d := connect()
	err := d.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyTasks"))
		err := b.Delete([]byte(s))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ListTasks() {
	d := connect()
	err := d.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyTasks"))

		fmt.Println("These are your current tasks:")
		err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("%s \n", v)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
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
	err := d.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyTasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
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
	// b.ForEach panics if there are no keys in the bucket to iterate over,
	// hence the deferred recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Oh, Seems like you have no tasks to do! Nice :)")
		}
	}()

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
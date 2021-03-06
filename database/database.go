package db

import (
	"encoding/binary"
	bolt "go.etcd.io/bbolt"
	"time"
)

var taskBucket = []byte("Tasks")
var db *bolt.DB

type Task struct {
	Key int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := intToByteSlice(int(id64))
		return bucket.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			tasks = append(tasks, Task{
				Key: byteSliceToInteger(key),
				Value: string(value),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		return bucket.Delete(intToByteSlice(key))
	})
	if err != nil {
		return err
	}
	return nil

}

func intToByteSlice(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func byteSliceToInteger(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
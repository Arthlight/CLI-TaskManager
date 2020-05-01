package db

import "go.etcd.io/bbolt"

var taskBucket = []byte("Tasks")
var db *bbolt.DB

type Task struct {
	Key int
	value string
}

package main

import (
	db "CLI-TaskManager/database"
	"CLI-TaskManager/cmd"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"log"
	"os"
)


func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "TaskManager.db")
	must(db.Init(dbPath))
	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
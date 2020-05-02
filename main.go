package main

import (
	db "CLI-TaskManager/database"
	"fmt"
	"CLI-TaskManager/cmd"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"log"
	"os"
)

func main() {
	home, _ := homedir.Dir()
	fmt.Println(home)
	dbPath:= filepath.Join(home, "TaskManager.db")
	fmt.Println(dbPath)
	must(db.Init(dbPath))
	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
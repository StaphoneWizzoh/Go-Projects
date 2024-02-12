package main

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
)

type Product struct{
  Id int
  Name string
  Inventory int
  Price int
}

func main(){
  db, err := sql.Open("sqlite3", "./database.db")
  if err != nil{
    log.Fatal(err.Error())
  }
  
}
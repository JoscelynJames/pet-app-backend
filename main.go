package main

import (  
  "database/sql"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

const (  
  host     = "localhost"
  port     = 5432
  user     = "joscelynjames"
  dbname   = "wheres_my_pet_db"
)

func main() {  
	getSession()

	r:= mux.NewRouter()

	r.Handle("/", r)
	r.HandleFunc("/users", getUsers).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func getSession() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s sslmode=disable",
    host, port, user, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
}

type active_user struct {
	ID         uint 	`json:"id"`
	Email 		 string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Password 	 string `json:"password"`
	Admin			 bool   `json:"admin"`
}

func getUsers(http.ResponseWriter, *http.Request) {
	
}




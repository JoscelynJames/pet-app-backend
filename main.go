package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/postgres"
)

type Active_Users struct {
	Id 				int 	 `gorm:"AUTO_INCREMENT" form:"id" json:"id`
	FirstName string `gorm:"not null" form:"first_name" json:"first_name"`
	Lastname 	string `gorm:"not null" form:"last_name" json:"last_name"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/users", PostUser)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run(":8080")
}

func InitDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=joscelynjames dbname=wheres_my_pet_db sslmode=disable")
	db.LogMode(true)
	
    // Error
    if err != nil {
        panic(err)
    }
    // Creating the table
    if !db.HasTable(&Active_Users{}) {
        db.CreateTable(&Active_Users{})
        db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Active_Users{})
    }

    return db
}

func PostUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user Active_Users
	c.Bind(&user)


}

func GetUsers(c *gin.Context) {
	var users = []Active_Users{
		Active_Users{Id: 1, FirstName: "Jo", Lastname: "James"},
		Active_Users{Id: 2, FirstName: "Bradford", Lastname: "Lamson_Scribner"},
	}

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	// code goes here
}

func UpdateUser(c *gin.Context) {
    // code goes here
}

func DeleteUser(c *gin.Context) {
    // code goes here
}











// import (  
//   "database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/mux"

//   _ "github.com/lib/pq"
// )

// const (  
//   host     = "localhost"
//   port     = 5432
//   user     = "joscelynjames"
//   dbname   = "wheres_my_pet_db"
// )

// func main() { 
// 	// getSession() 
// 	router := mux.NewRouter()
// 	router.HandleFunc("/users", GetUsers).Methods("GET")
// 	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
// 	router.HandleFunc("/user/{id}", CreateUser).Methods("POST")
// 	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
// 	log.Fatal(http.ListenAndServe(":8000", router))

// }


// func getSession() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//     "dbname=%s sslmode=disable",
//     host, port, user, dbname)
//   db, err := sql.Open("postgres", psqlInfo)
//   if err != nil {
//     panic(err)
//   }
//   defer db.Close()

//   err = db.Ping()
//   if err != nil {
//     panic(err)
//   }

// 	fmt.Println("Successfully connected!")
// }


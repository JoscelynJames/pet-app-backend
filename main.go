package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "log"	
	"github.com/joscelynjames/bird-app-backend/user"
	"github.com/joscelynjames/bird-app-backend/user/login"
	"github.com/joscelynjames/bird-app-backend/user/signup"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/active_users", PostUser)
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
	// defer db.Close()

	db.AutoMigrate(&Active_Users{})
	// Creating the table
	// if !db.HasTable(&Active_Users{}) {
	// 	db.CreateTable(&Active_Users{})
	// 	db.Set("gorm:table_options", "ON CONFLICT").CreateTable(&Active_Users{})
	// }
	return db
}


func PostUser(c *gin.Context) {
    db := InitDb()

    var active_user Active_Users
		c.Bind(&active_user)

		db.Create(active_user)
		// Display error
		c.JSON(201, gin.H{"success": active_user})
		
    // if active_user.first_name != "" && active_user.last_name != "" {
    //     db.Create(&active_user)
    //     // Display error
    //     c.JSON(201, gin.H{"success": active_user})
    // } else {
    //     // Display error
    //     c.JSON(422, gin.H{"error": "Fields are empty"})
    // }
}


func GetUsers(c *gin.Context) {
	// code goes here
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

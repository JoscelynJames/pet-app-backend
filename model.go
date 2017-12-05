package user

type Active_Users struct {
	Id         uint 	`gorm:"primary_key"`
	Email 		 string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Password 	 string `json:"password"`
	Admin			 bool   `json:"admin"`
}
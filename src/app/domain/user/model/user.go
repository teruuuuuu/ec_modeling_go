package user_model

type User struct {
	UserId   uint `gorm:"primary_key"`
	Name     string
	Password string
}

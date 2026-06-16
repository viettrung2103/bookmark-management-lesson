package main

import (
	"github.com/google/uuid"
	"github.com/viettrung2103/bookmark-management-lesson/pkg/sqldb"
	"gorm.io/gorm"
)

type User struct {
	ID       string
	Username string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

func main() {

	dbClient, err := sqldb.NewClient("")
	if err != nil {
		panic(err)
	}

	dbClient.AutoMigrate(&User{})

	user := &User{
		//ID:       "1",
		Username: "test",
		Password: "test",
	}
	dbClient.Create(user)

}

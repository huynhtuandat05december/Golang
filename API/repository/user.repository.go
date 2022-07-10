package repository

import (
	"golang_api/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	IsDuplicateEmail(email string) *gorm.DB
	CreateUser(user models.User) models.User
	VerifyAccount(email string, password string) interface{}
	ProfileUser(userID uint64) models.User
	UpdateUser(user models.User) models.User
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

/**************AUTH**************/
func (db *userConnection) IsDuplicateEmail(email string) *gorm.DB {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	return res
}

func (db *userConnection) CreateUser(user models.User) models.User {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	user.Password = string(hashPassword)
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyAccount(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		return false
	}
	// errorCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if errorCompare != nil {
	// 	log.Println(errorCompare)
	// 	return false
	// }
	return user

}

/**************USER**************/
func (db *userConnection) ProfileUser(userID uint64) models.User {
	var user models.User
	db.connection.Where("id = ?", userID).Take(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	var newInfoUser models.User
	if user.Password != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			log.Println(err)
			panic("Failed to hash a password")
		}
		user.Password = string(hashPassword)
	}
	db.connection.Model(&newInfoUser).Where("id = ?", user.ID).Updates(user)
	return newInfoUser
}

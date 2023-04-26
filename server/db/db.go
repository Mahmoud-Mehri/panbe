package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost port=123 dbname=bcdb user=bcuser password=bcuserpass sslmode=disable"
var DB *gorm.DB;

func ConnectToDB() bool {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err == nil {
		DB = db
		return true
	}
	log.Print(err)
	return false
}

func CreateUser(user *User) (*gorm.DB, bool) {
	result := DB.Create(&user)
	success := result.RowsAffected > 0
	return result, success
}

func UpdateUser(user *User) (*gorm.DB, bool) {
	result := DB.Update("*", &user)
	success := result.RowsAffected > 0
	return result, success
}

func DeleteUser(userId uint) (*gorm.DB, bool) {
	
	result := DB.Delete(&User{}, userId)
	success := result.RowsAffected > 0
	return result, success
}

func GetUserById(userId uint, includeWallet bool) (*User, *gorm.DB, bool) {
	user := &User{};
	result := DB.Model(User{ID: userId}).Find(&user)
	success := (user != nil) && (user.ID > 0)
	if success && includeWallet {
		DB.Model(User{}).Association("wallets").Find(&user.wallets)
	}
	return user, result, success
}

func GetUserByUsername(username string, includeWallet bool) (*User, *gorm.DB, bool) {
	user := &User{};
	result := DB.Model(&User{Username: username}).Find(&user)
	success := (user != nil) && (user.ID > 0)
	return user, result, success
}
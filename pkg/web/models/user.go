package models

import (
	"errors"
	"fmt"

	"github.com/Devil666face/goaccountant/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	PasswordLen = 8
)

var (
	ErrEmptyUsername     = fiber.NewError(fiber.StatusBadRequest, "Username is required")
	ErrPasswordMissmatch = fiber.NewError(fiber.StatusBadRequest, "Passwords mismatch")
	ErrPasswordRequired  = fiber.NewError(fiber.StatusBadRequest, "Password is required")
	ErrPasswordShort     = fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("The minimum len of password is %d", PasswordLen))
	ErrPasswordEncrypt   = fiber.ErrInternalServerError
	ErrUserNotUniq       = fiber.NewError(fiber.StatusBadRequest, "User already create")
)

type User struct {
	gorm.Model
	Username        string `gorm:"unique;not null" form:"username"`
	Password        string `gorm:"not null" form:"password"`
	PasswordConfirm string `gorm:"-" form:"password_confirm"`
	Admin           bool   `gorm:"default:false" form:"admin"`
}

func (u *User) Create(db *gorm.DB) error {
	// If user with this username is found return err
	if u.IsFound(db) {
		return ErrUserNotUniq
	}
	// if err := u.GetUser(db, u.Username); !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return ErrUserNotUniq
	// }
	return db.Create(u).Error
	// if err := db.Create(u); err.Error != nil {
	// 	return err.Error
	// }
	// return nil
}

func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *User) IsFound(db *gorm.DB) bool {
	return !errors.Is(u.GetUserByUsername(db, u.Username), gorm.ErrRecordNotFound)
}

func (u *User) validateInput() bool {
	return utils.ValidateUserInputs(u.Username, u.Password, u.PasswordConfirm)
}

func (u *User) Validate() error {
	if !u.validateInput() {
		return fiber.ErrInternalServerError
	}
	if u.Username == "" {
		return ErrEmptyUsername
	}
	if u.Password != u.PasswordConfirm {
		return ErrPasswordMissmatch
	}
	if u.Password == "" || u.PasswordConfirm == "" {
		return ErrPasswordRequired
	}
	if len([]rune(u.Password)) < PasswordLen {
		return ErrPasswordShort
	}
	// Hash password and do u.Password = password
	if u.hashPassword() != nil {
		return ErrPasswordEncrypt
	}
	return nil
}

func (u *User) hashPassword() error {
	password, err := utils.GenHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	return nil
}

func (u *User) ComparePassword(password string) bool {
	if err := utils.ComparePassword(u.Password, password); err == nil {
		return true
	}
	return false
}

func GetAllUsers(db *gorm.DB) []User {
	users := []User{}
	db.Find(&users)
	return users
}

func (u *User) GetUser(db *gorm.DB) error {
	return db.First(u, u.ID).Error
}

func (u *User) GetUserByUsername(db *gorm.DB, username string) error {
	u.ID = 0
	return db.Where("username = ?", username).First(&u).Error
	// return db.Where("username = ?", username).Take(&u).Error
}

// func (user *User) Set(username string, password string, admin bool) {
// 	user.Username = username
// 	user.Password = password
// 	user.Admin = admin
// }

// func CreateUser(user *User) *gorm.DB {
// 	return database.DB.Create(user)
// }

// func UpdateUser(user *User) *gorm.DB {
// 	return database.DB.Save(user)
// }

// func DeleteUser(user *User) *gorm.DB {
// 	return database.DB.Unscoped().Delete(user)
// }

// func GetUserByUsername(dest *User, username string) *gorm.DB {
// 	return database.DB.Where("username= ?", username).Take(&dest)
// }

// func GetAllUsers() []User {
// 	var users []User
// 	database.DB.Find(&users)
// 	return users
// }

// type UserForm struct {
// 	Username        string `form:"username"`
// 	Password        string `form:"password"`
// 	PasswordConfirm string `form:"password_confirm"`
// 	Admin           string `form:"admin"`
// }

// func (form *UserForm) IsAdmin() bool {
// 	if form.Admin != "" {
// 		return true
// 	}
// 	return false
// }

// func (form *UserForm) IsEmptyUsername() (string, bool) {
// 	if form.Username == "" {
// 		return "Username is required.", true
// 	}
// 	return "", false
// }

// func (form *UserForm) IsPasswordsMatch() (string, bool) {
// 	if form.Password != form.PasswordConfirm {
// 		return "The passwords don't match.", true
// 	}
// 	return "", false
// }

// func (form *UserForm) IsPasswordsEmpty() (string, bool) {
// 	if form.Password == "" || form.PasswordConfirm == "" {
// 		return "Password is required.", true
// 	}
// 	return "", false
// }

// func (form *UserForm) IsPasswordsShort() (string, bool) {
// 	if len([]rune(form.Password)) < utils.PASSWORDLEN {
// 		return fmt.Sprintf("The minimum len of password is %d", utils.PASSWORDLEN), true
// 	}
// 	return "", false
// }

// func (form *UserForm) CheckPasswordForCreate() (string, bool) {
// 	if message, ok := form.IsPasswordsMatch(); ok {
// 		return message, ok
// 	}
// 	if message, ok := form.IsPasswordsEmpty(); ok {
// 		return message, ok
// 	}
// 	if message, ok := form.IsPasswordsShort(); ok {
// 		return message, ok
// 	}
// 	return "", false
// }

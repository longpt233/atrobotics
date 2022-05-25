package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUserByEmail(string) (model.User, error)
	GetUser(string) (model.User, error)
	AddUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	GetAllUser(string) ([]model.User, error)
	GetAllUserPaging(userId string,adminId string, limit int ,offset int ) (listUser []model.User, err error)
}

type userRepository struct {
	connection *gorm.DB
}

//NewUserRepository --> returns new user repository
func NewUserRepository() UserRepository {
	 
	myclient := &MySQLClient{}
	return &userRepository{
		connection:myclient.GetConn(),
	}
}

func (db *userRepository) GetUser(id string) (user model.User, err error) { // TODO tại sao k xóa dc cái user thường ở cái return này đi như hàm bên dưới ?
	return user, db.connection.First(&user, "user_id=?", id).Error
}

func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) GetUserByEmail(email string) (user model.User, err error) {
	return user, db.connection.First(&user, "user_email=?", email).Error
}
func (db *userRepository) UpdateUser(user model.User) (model.User, error){
	var checkUser model.User
	if err := db.connection.First(&checkUser, "user_id = ?",user.UserID).Error; err != nil {
		return checkUser, err
	}
	return user, db.connection.Model(&user).Where(model.User{UserID: user.UserID}).Updates(&user).Error
}

func (db *userRepository) GetAllUser(roleId string) (listUser []model.User, err error){
	return listUser, db.connection.Find(&listUser, "user_role_id=?",roleId).Error
}


func (db *userRepository) GetAllUserPaging(userId string,adminId string,limit int ,offset int) (listUser []model.User, err error){
	return listUser, db.connection.Limit(limit).Offset(offset).Find(&listUser, "user_role_id=? or user_role_id=? ",userId, adminId).Error
}
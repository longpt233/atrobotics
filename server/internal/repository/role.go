package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

type RoleRepository interface {
	GetRole(int) (model.Role, error)
	AddRole(role model.Role) (model.Role, error)
	GetRoleByName(string) (model.Role, error)
}

type roleRepository struct{
	connection *gorm.DB
}

func NewRoleRepository() RoleRepository{
	return &roleRepository{
		connection: DB(),
	}
}

func (db* roleRepository) GetRole(id int) (role model.Role, err error){
	return role, db.connection.First(&role, id).Error
}

func (db* roleRepository) AddRole(role model.Role) (model.Role, error){
	return role, db.connection.Create(&role).Error
}

func (db* roleRepository) GetRoleByName(roleName string) (role model.Role, err error){
	return role, db.connection.First(&role, "role_name=?", roleName).Error
}
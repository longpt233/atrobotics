package repository

import (
	"atro/internal/model"

	"github.com/jinzhu/gorm"
)

//OrderRepository --> Repository for Order Model
type OrderRepository interface {
	OrderProduct(model.Order) (model.Order, error)
	GetAllOrder() ([]model.Order, error)
	GetOrder(string) (model.Order, error)
	UpdateOrder(model.Order) (model.Order, error)
	GetAllOrderOptions(filter map[string]interface{}, limit int, offset int, query string) ([]model.Order, error)
}

type orderRepository struct {
	connection *gorm.DB
}

//NewOrderRepository --> returns new order repository
func NewOrderRepository() OrderRepository {
	
	myclient := &MySQLClient{}
	return &orderRepository{
		connection:myclient.GetConn(),
	}
}

func (db *orderRepository) OrderProduct(order model.Order) (model.Order, error) {
	return order, db.connection.Create(&order).Error
}

func (db *orderRepository) GetAllOrder() (orders []model.Order, err error) {
	return orders, db.connection.Find(&orders).Error
}

func (db *orderRepository) GetOrder(id string) (order model.Order, err error) {
	return order, db.connection.First(&order, "order_id=?", id).Error
}

func (db *orderRepository) UpdateOrder(model.Order) (model.Order, error) {
	return model.Order{}, nil
}

func (db *orderRepository) GetAllOrderOptions(filter map[string]interface{}, limit int, offset int, query string) (orders []model.Order, err error) {
	return orders, db.connection.Where(filter).Limit(limit).Offset(offset).Order(query).Find(&orders).Error
}

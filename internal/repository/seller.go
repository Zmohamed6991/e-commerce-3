package repository

import (
	"e-commerce/internal/models"
	"os/user"
)

func (p *Postgres) FindSellerByEmail(email string) (*models.Seller, error) {
	seller := &models.Seller{}

	if err := p.DB.Where("email = ?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

// Create a user in the database
func (p *Postgres) CreateSeller(seller *models.Seller) error {
	if err := p.DB.Create(seller).Error; err != nil {
		return err
	}
	return nil
}

// Update a user in the database
func (p *Postgres) UpdateSeller(seller *models.Seller) error {
	if err := p.DB.Save(seller).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) CreateProduct(product *models.Product) error {
	if err := p.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// getting the order via the ID - returning order from this function
func (p *Postgres) GetOrderByOrderID(orderID uint) (*models.Order, error) {

	// declare variable to return the order
	order := &models.Order{} // one object [] is a slice

	// filter the orders using order ID
	if err := p.DB.Where("order_id = ?", orderID).Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

// type order - object trying to save
func (p *Postgres) UpdateOrder(order *models.Order) error {
	if err := p.DB.Save(order).Error; err != nil {
		return err
	}

}

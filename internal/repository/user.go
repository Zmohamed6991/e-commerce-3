package repository

import (
	"e-commerce/internal/models"
	"errors"
	"log"
)

func (p *Postgres) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p *Postgres) FindAllUsers() ([]models.User, error) {
	var users []models.User

	if err := p.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (p *Postgres) GetUserByID(userID uint) (*models.User, error) {
	user := &models.User{}

	if err := p.DB.Where("ID = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create a user in the database
func (p *Postgres) CreateUser(user *models.User) error {
	if err := p.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// Update a user in the database
func (p *Postgres) UpdateUser(user *models.User) error {
	if err := p.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetProductByID(productID uint) (*models.Product, error) {
	product := &models.Product{}

	if err := p.DB.Where("ID = ?", productID).First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Postgres) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	if err := p.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Postgres) AddToCart(cart *models.IndividualItemInCart) error {
	if err := p.DB.Save(cart).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetCartItemByProductID(productID uint) (*models.IndividualItemInCart, error) {
	cart := &models.IndividualItemInCart{}

	if err := p.DB.Where("product_id = ?", productID).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// Delete a product from the cart
func (p *Postgres) DeleteProductFromCart(cart *models.IndividualItemInCart) error {
	if err := p.DB.Delete(cart).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetCartsByUserID(userID uint) ([]*models.IndividualItemInCart, error) {
	var cartItems []*models.IndividualItemInCart

	// Fetch cart items for the user with null OrderID
	if err := p.DB.Where("user_id = ? AND order_id IS NULL", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}

	// If no items found, return an error
	if len(cartItems) == 0 {
		return nil, errors.New("no items found in cart")
	}

	return cartItems, nil
}

func (p *Postgres) OrderHistorybyUserID(userID uint) ([]*models.Order, error) {
	var orders []*models.Order

	if err := p.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (p *Postgres) GetOrderItemsByOrderID(orderID uint) ([]*models.OrderItem, error) {
	var orderDetails []*models.OrderItem

	if err := p.DB.Preload("Product").Where("order_id = ?", orderID).Find(&orderDetails).Error; err != nil {
		return nil, err
	}
	return orderDetails, nil
}


// create an order
func (p *Postgres) CreateOrder(order *models.Order) error {
	tx := p.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	// Attempt to create the order
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		log.Printf("Error creating order: %v", err) // Add this log
		return err
	}

	// Clear the cart
	if err := tx.Where("user_id = ?", order.UserID).Delete(&models.IndividualItemInCart{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error clearing cart: %v", err) // Add this log
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err) // Add this log
		return err
	}

	return nil
}
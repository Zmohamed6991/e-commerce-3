package repository

import (
	"e-commerce/internal/models"
	"errors"
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

	if err := p.DB.Where("ID = ?", productID).First(&cart).Error; err != nil {
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





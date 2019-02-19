package main

import "database/sql"

// Repository - ...
type Repository interface {
	GetAll() ([]*User, error)
	Get(id string) (*User, error)
	Create(*User) error
	GetByEmail(email string) (*User, error)
}

// UserRepository
type UserRepository struct {
	Db *sql.DB
}

// GetAll - ...
func (repo *UserRepository) GetAll() ([]*User, error) {
	var users []*User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get - ...
func (repo *UserRepository) Get(id string) (*User, error) {
	var user *User
	if err := repo.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmail - ...
func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user *User
	if err := repo.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Create - ...
func (repo *UserRepository) Create(user *User) error {
	if err := repo.Db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

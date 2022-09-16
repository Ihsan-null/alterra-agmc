package database

import (
	"echo-database/config"
	"echo-database/models"
)

func GetUsers() ([]models.User, error){
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id *int) (models.User, error) {
	var users models.User
	if err := config.DB.Where("id", id).First(&users).Error; err != nil {
		return users, err
	}
	
	return users, nil
}

func CreateUser(user models.User) (models.User, error){
	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id *int, userUpdate *models.User) (*models.User, error) {
	var users models.User
	
	if err := config.DB.Model(&users).
	Where("id = ?", id).
	Updates(models.User{Name: userUpdate.Name, Email: userUpdate.Email, Password: userUpdate.Password}).Error; err != nil {
		return userUpdate, err
	}
	return userUpdate, nil	
}

func DeleteUser(id *int) (models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var users models.User
	if err := config.DB.Where("email", email).First(&users).Error; err != nil {
		return users, err
	}
	
	return users, nil
}
package service

import (
		"Week_12/internal/models"
		"golang.org/x/crypto/bcrypt"
		"errors")

func (s *UserService)Login(email string,password string)(models.User,error){
	if email == "" || password == ""{
		return models.User{},errors.New("Missing required fields")
	}
	user,err := s.repo.GetUserDetails(email)
	if err!=nil{
		return models.User{},errors.New("Invalid Email or Password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err!=nil{
		return models.User{},errors.New("Invalid Email or Password")
	}
	return user,nil
	
}
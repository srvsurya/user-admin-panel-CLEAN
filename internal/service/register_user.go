package service

import ("Week_12/internal/repository"
		"errors"
		"Week_12/internal/models")

type UserService struct{
	repo *repository.UserRepository
}
func NewUserService(repo *repository.UserRepository) *UserService{
	return &UserService{repo:repo}
}

func (s *UserService)RegisterUser(name string,email string,password string,confirmPassword string)error{
	var user models.User
	var err error
	var emailexists bool
	role:="app_user"
	if name == "" || email == "" || password == "" || confirmPassword == ""{
		return errors.New("Missing required fields")
	}
	if password != confirmPassword{
		return errors.New("Passwords does not match")
	}

	user = models.User{
		Name:name,
		Email:email,
		Password:password,
		Role:role,
	}
	if emailexists,err=s.repo.EmailExists(email);err!=nil{
		return errors.New(err.Error())
	}else if emailexists{
		return errors.New("Email already exists in database")
	}
	return s.repo.Create(&user)
}
func (s *UserService)AdminRegisterUser(name string,email string,role string,password string)error{
	var user models.User
	var err error
	if name == "" || email == "" || password == "" {
		return errors.New("Missing required fields")
	}
	if err!=nil{
		return errors.New(err.Error())
	}
	user = models.User{
		Name:name,
		Email:email,
		Password:password,
		Role:role,
	}

	return s.repo.Create(&user)
}
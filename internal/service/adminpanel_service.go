package service

import ("Week_12/internal/models"
		"errors")

func (s *UserService)AdminViewService(id int,offset int,limit int)([]models.User,error){
	var users []models.User
	var admin models.User
	admin,err:=s.repo.GetUserById(id)
	if err!=nil{
		return nil,err
	}
	if admin.Role != "admin"{
		return nil,errors.New("Unauthorized")
	}

	users,err=s.repo.PageView(offset,limit)
	if err!=nil{
		return nil,err
	}
	return users,nil

}
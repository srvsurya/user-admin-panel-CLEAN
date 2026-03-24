package service

import "errors"

func (s *UserService)AdminEdit(id int, adminId int, current_role string)(error){
	var role string
	admin,err:=s.repo.GetUserById(adminId)
	if err!=nil{
		return err
	}
	if admin.Role != "admin"{
		return errors.New("Unauthorized")
	}
	if current_role == "app_user"{
		role = "admin"
	}else if current_role == "admin"{
		role = "app_user"
	}
	err = s.repo.AdminEdit(id,role)
	if err!=nil{
		return err
	}
	return nil
}
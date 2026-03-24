package service

import "errors"

func (s *UserService)HomeService(id int)(string,error){
	if id<0{
		return "",errors.New("Invalid ID")
	}
	name,err:=s.repo.HomePageName(id)
	if err!=nil{
		return "",err
	}
	return name,nil
}
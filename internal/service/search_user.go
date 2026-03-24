package service
import ("Week_12/internal/models"
		"strings"
		"errors")

func (s *UserService)SearchUsers(name string)([]models.User, error){
	var users []models.User

	name = strings.TrimSpace(name)
	if name == ""{
		return nil,errors.New("Must Enter a Name")
	}
	users,err:= s.repo.SearchUsers(name)
	if err!=nil{
		return nil,err
	}
	return users,nil 
}


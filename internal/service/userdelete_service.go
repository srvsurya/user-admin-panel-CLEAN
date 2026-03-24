package service
import "errors"

func (s *UserService)DeleteUser(id int,adminId int)(error){
	if id<0{
		return errors.New("Not a valid ID")
	}
	admin,err:=s.repo.GetUserById(adminId)
	if err!=nil{
		return err
	}
	if admin.Role != "admin"{
		return errors.New("Unauthorized")
	}
	err=s.repo.Delete(id)
	if err!=nil{
		return err
	}
	return nil
}
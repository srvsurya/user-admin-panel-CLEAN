package repository

import ("gorm.io/gorm"
		"Week_12/internal/models"
		)
type UserRepository struct{
	db *gorm.DB
}
func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{db:db}

}
func (r *UserRepository) EmailExists(email string)(bool,error){
	var count int64
	err:=r.db.Model(&models.User{}).Where("email = ?",email).Count(&count).Error
	if err!=nil{
		return false,err
	}
	return count>0,nil
}
func (r *UserRepository) Create(user *models.User)error{
	return r.db.Create(user).Error
}

func (r *UserRepository) Delete(id int)error{
	return r.db.Delete(&models.User{},id).Error
}
func (r *UserRepository) SearchUsers(search string)([]models.User, error){
	var users []models.User
	err:= r.db.Select("user_id,name,email,role").Where("name ILIKE ?","%"+search+"%").Find(&users).Error
	return users,err
}
func (r *UserRepository)GetUserDetails(email string)(models.User,error){
	var user models.User
	err:=r.db.Select("user_id,email,password,role").Where("email = ?",email).First(&user).Error
	return user,err
}
func (r *UserRepository)HomePageName(id int)(string,error){
	var user models.User
	err:=r.db.Select("name").Where("user_id = ?",id).First(&user).Error
	return user.Name,err
}
func (r *UserRepository)PageView(offset int, limit int)([]models.User,error){
	var users []models.User
	err:=r.db.Select("user_id,name,email,role").Order("user_id").Limit(limit).Offset(offset).Find(&users).Error
	return users,err
}	
func (r *UserRepository)AdminEdit(id int, role string)error{
	err:=r.db.Model(&models.User{}).Where("user_id = ?",id).Update("role",role).Error
	return err
}
func (r *UserRepository)GetUserById(id int)(models.User,error){
	var user models.User
	err:=r.db.Where("user_id = ?",id).First(&user).Error
	return user,err
}
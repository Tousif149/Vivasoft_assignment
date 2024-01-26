package repositories
import(
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	
	"gorm.io/gorm"
)
//user repo
type UserRepo struct{
	db *gorm.DB
}
//user repo 
func UserDBInstance(db *gorm.DB) domain.IUserRepo{
	return &UserRepo{
		db: db,
	}
}
//create user
func (ur *UserRepo) CreateUser(user *models.UserDetail) error{
	result := ur.db.Create(user)
	if result.Error != nil{
		return result.Error
	}
	return nil
}
//get user
func (ur *UserRepo) GetUser(username *string) (*models.UserDetail, error){
	user := &models.UserDetail{}
	result := ur.db.Where("username = ?", username).First(user)
	if result.Error != nil{
		return nil, result.Error
	}
	return user, nil
}
//login user	
func (ur *UserRepo) LoginUser(user *models.UserDetail) error{
	result := ur.db.Where("username = ? AND password_hash = ?", user.Username, user.PasswordHash).First(user)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

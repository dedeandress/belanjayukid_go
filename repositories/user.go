package repositories

import (
	"go_sample_login_register/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *models.User) (insertedUser *models.User, err error)
	GetUserByID(userID string) (user *models.User, err error)
	GetUserByEmail(email string) (user *models.User, err error)
	Update(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *DataSource
}

var userRepo *userRepository

func GetUserRepository() UserRepository {
	if DBTrx != nil {
		userRepo = &userRepository{db: DBTrx}
	} else {
		userRepo = &userRepository{db: DB}
	}
	return userRepo
}

func (userRepo *userRepository) Insert(user *models.User) (insertedUser *models.User, err error) {
	insertedUser = &models.User{}
	if err := userRepo.db.Create(user).Scan(insertedUser).Error; err != nil {
		// fmt.Errorf("Error while inserting record in database (%s) ", err.Error())
		return nil, err
	}
	return insertedUser, err
}

func (userRepo *userRepository) GetUserByID(userID string) (user *models.User, err error) {
	user = &models.User{}
	res := userRepo.db.Scopes(filterUsersByID(userID)).First(&user)
	if err := res.Error; err != nil {
		// fmt.Errorf("Error querying for user for user with ID %s and error: %s", userID, err.Error())
		return nil, err
	}
	return user, err
}

func (userRepo *userRepository) GetUserByEmail(email string) (user *models.User, err error) {
	user = &models.User{}
	res := userRepo.db.Scopes(filterUsersByEmail(email)).First(&user)
	if err := res.Error; err != nil {
		// fmt.Errorf("Error querying for user for user with ID %s and error: %s", userID, err.Error())
		return nil, err
	}
	return user, err
}

func (userRepo *userRepository) Update(user *models.User) (*models.User, error) {
	if err := userRepo.db.Model(&user).Updates(user).First(&user).Error; err != nil {
		// fmt.Errorf("Error while updating record in database (%s) ", err.Error())
		return nil, err
	}
	return user, nil
}

func filterUsersByID(userID string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("users.id = ?", userID)
}

func filterUsersByEmail(email string) func(db *gorm.DB) *gorm.DB {
	return makeFilterFunc("users.email = ?", email)
}

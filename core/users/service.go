package users

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/davisbento/gorm-api/config/jwtManager"
	"github.com/davisbento/gorm-api/utils"
	"github.com/gofrs/uuid"
)

type UserCreated struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UseCase interface {
	GetAll() ([]*UserList, error)
	Get(Id int64) (*UserList, error)
	getByEmail(email string) (*User, error)
	countByEmail(email string) (int64, error)
	Store(u *UserCreateModel) (UserCreated, error)
	Login(u *UserLogin) (string, error)
}

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetAll() ([]*UserList, error) {
	var users []*UserList

	result := s.DB.Model(&User{}).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *Service) Get(id int64) (*UserList, error) {
	var user *UserList

	result := s.DB.Model(&User{}).First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (s *Service) Store(u *UserCreateModel) (UserCreated, error) {
	countUsersByEmail, err := s.countByEmail(u.Email)

	newUser := UserCreated{}

	if countUsersByEmail > 0 || err != nil {
		return newUser, errors.New("e-mail already exists")
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		panic(err)
	}

	uuid, _ := uuid.NewV4()

	user := User{Id: uuid, Name: u.Name, Email: u.Email, Password: hashedPassword}

	result := s.DB.Create(&user)

	if result.Error != nil {
		return newUser, result.Error
	}

	newUser.Name = u.Name
	newUser.Email = u.Email

	return newUser, nil
}

func (s *Service) Login(u *UserLogin) (string, error) {
	user, err := s.getByEmail(u.Email)
	if err != nil {
		return "", err
	}

	isValid := utils.ComparePasswords(user.Password, u.Password)

	if !isValid {
		return "", fmt.Errorf("password-invalid")
	}

	token, err := jwtManager.GenerateToken(user.Id)

	if err != nil {
		return "", nil
	}

	return token, nil
}

func (s *Service) getByEmail(email string) (*User, error) {
	var user *User

	result := s.DB.Where(&User{Email: email}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (s *Service) countByEmail(email string) (int64, error) {
	var count int64

	result := s.DB.Model(&User{}).Where("email = ?", email).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

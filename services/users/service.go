package users

import (
	"context"
	"github.com/mercadofarma/services/codes"
	"github.com/mercadofarma/services/errors"
	"github.com/mercadofarma/services/repos/models"
	"github.com/mercadofarma/services/repos/users"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, password string, firstName string, lastName string, role string, phoneNumber string) (*models.User, error)
	ValidateUserInputs(email string, password string, firstName string, role string) error
}

type ServiceImpl struct {
	userRepo users.UserRepo
}

func NewUserService(userRepo users.UserRepo) *ServiceImpl {
	return &ServiceImpl{
		userRepo: userRepo,
	}
}

func (svc *ServiceImpl) CreateUser(ctx context.Context, email string, password string, firstName string, lastName string, role, phoneNumber string) (*models.User, error) {
	if err := svc.ValidateUserInputs(email, password, firstName, role); err != nil {
		return nil, err
	}

	if svc.userRepo.CheckIfUserExist(ctx, email, models.Role(role)) {
		return nil, errors.ErrorWithCode(codes.InvalidInput, "email is already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Hash:        string(hash),
		Role:        models.Role(role),
		Active:      true,
	}

	err = svc.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (svc *ServiceImpl) ValidateUserInputs(email string, password string, firstName string, role string) error {
	if len(strings.Trim(email, " ")) == 0 {
		return errors.ErrorWithCode(codes.InvalidInput, "invalid email")
	}

	if len(strings.Trim(password, " ")) == 0 {
		return errors.ErrorWithCode(codes.InvalidInput, "invalid password")
	}

	if len(strings.Trim(firstName, " ")) == 0 {
		return errors.ErrorWithCode(codes.InvalidInput, "invalid first name")
	}

	if !models.IsValidRole[models.Role(role)] || string(models.ShopperRole) != role {
		return errors.ErrorWithCode(codes.InvalidInput, "invalid role")
	}

	return nil
}

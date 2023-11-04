package users

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mercadofarma/services/codes"
	"github.com/mercadofarma/services/errors"
	"github.com/mercadofarma/services/repos/models"
	"github.com/mercadofarma/services/repos/users"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, password string, firstName string, lastName string, role string, phoneNumber string) (*models.User, error)
	Login(ctx context.Context, email string, role string, password string) (*models.Authentication, error)
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

func (svc *ServiceImpl) Login(ctx context.Context, email string, role string, password string) (*models.Authentication, error) {
	if email == "" {
		return nil, errors.ErrorWithCode(codes.RequiredInput, "email is required")
	}

	if role == "" {
		return nil, errors.ErrorWithCode(codes.RequiredInput, "role is required")
	}

	if password == "" {
		return nil, errors.ErrorWithCode(codes.RequiredInput, "password is required")
	}

	user, err := svc.userRepo.GetUserByEmail(ctx, email, models.Role(role))
	if err != nil {
		return nil, err
	}

	if user.UserId == "" {
		return nil, errors.ErrorWithCode(codes.ResourceNotFound, "email address not found")
	}

	const days = 30
	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		return nil, errors.ErrorWithCode(codes.Unauthorized, "invalid login credentials, please try again")
	}

	claims := models.CustomClaims{
		UserId: user.UserId,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(days))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "mercadofarma",
			ID:        user.UserId,
			Audience:  []string{string(user.Role)},
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	apiKey, err := token.SignedString([]byte(user.SecretKey))
	if err != nil {
		return nil, err
	}

	return &models.Authentication{
		User:  *user,
		Token: apiKey,
	}, nil
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

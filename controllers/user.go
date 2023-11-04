package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mercadofarma/services/codes"
	swaggerModels "github.com/mercadofarma/services/models"
	"github.com/mercadofarma/services/repos/models"
	"github.com/mercadofarma/services/restapi/operations/shopper"
	"github.com/mercadofarma/services/services/users"
)

type UserController struct {
	BaseController
	userService users.UserService
}

func NewUserController(userService users.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) Login(params shopper.LoginParams) middleware.Responder {
	if params.LoginRequest == nil {
		return shopper.NewLoginBadRequest().WithPayload(&swaggerModels.Error{
			Code:    string(codes.InvalidInput),
			Message: "loginRequest can not be nil",
		})
	}

	ctx := params.HTTPRequest.Context()
	email := params.LoginRequest.Email
	password := params.LoginRequest.Password
	role := params.LoginRequest.Role

	response, err := controller.userService.Login(ctx, string(email), role, password)
	if err != nil {
		code, convertedError := controller.ConvertErrorToSwaggerError(err)
		switch code {
		case codes.InvalidInput, codes.RequiredInput:
			return shopper.NewLoginBadRequest().WithPayload(&convertedError)
		case codes.Unauthorized:
			return shopper.NewLoginUnauthorized().WithPayload(&convertedError)
		case codes.ResourceNotFound:
			return shopper.NewLoginNotFound().WithPayload(&convertedError)
		default:
			return shopper.NewLoginInternalServerError().WithPayload(&convertedError)
		}
	}

	return shopper.NewLoginOK().WithPayload(convertToAuthenticationResponse(response))
}

func (controller *UserController) SignUp(params shopper.SignUpShopperParams) middleware.Responder {
	if params.SignUpShopperRequest == nil {
		return shopper.NewSignUpShopperBadRequest().WithPayload(&swaggerModels.Error{
			Code:    string(codes.InvalidInput),
			Message: "signUpRequest can not be nil",
		})
	}

	email := params.SignUpShopperRequest.Email
	password := params.SignUpShopperRequest.Password
	firstName := params.SignUpShopperRequest.FirstName
	lastName := params.SignUpShopperRequest.LastName
	role := ""
	if params.SignUpShopperRequest.Role != nil {
		role = *params.SignUpShopperRequest.Role
	}

	ctx := params.HTTPRequest.Context()
	_, err := controller.userService.CreateUser(ctx, string(email), password, firstName, lastName, role, "")
	if err != nil {
		code, convertedError := controller.ConvertErrorToSwaggerError(err)
		switch code {
		case codes.InvalidInput:
			return shopper.NewSignUpShopperBadRequest().WithPayload(&convertedError)
		default:
			return shopper.NewSignUpShopperInternalServerError().WithPayload(&convertedError)
		}
	}

	return shopper.NewSignUpShopperNoContent()
}

func convertToAuthenticationResponse(data *models.Authentication) *swaggerModels.AuthenticationResponse {
	if data == nil {
		return &swaggerModels.AuthenticationResponse{}
	}

	return &swaggerModels.AuthenticationResponse{
		Token: data.Token,
		User: &swaggerModels.AuthenticationResponseUser{
			Active:      data.User.Active,
			Email:       data.User.Email,
			FirstName:   data.User.FirstName,
			LastName:    data.User.LastName,
			PhoneNumber: data.User.PhoneNumber,
			Role:        string(data.User.Role),
			UserID:      data.User.UserId,
		},
	}
}

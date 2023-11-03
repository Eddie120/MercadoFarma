package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mercadofarma/services/codes"
	swaggerModels "github.com/mercadofarma/services/models"
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

	return nil
}

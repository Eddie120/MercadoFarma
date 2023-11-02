package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mercadofarma/services/codes"
	"github.com/mercadofarma/services/restapi/operations/business"
	businessService "github.com/mercadofarma/services/services/business"
	"github.com/mercadofarma/services/services/users"
)

type BusinessController struct {
	BaseController
	userService     users.UserService
	businessService businessService.BusinessService
}

func NewBusinessController(userService users.UserService, businessService businessService.BusinessService) *BusinessController {
	return &BusinessController{
		userService:     userService,
		businessService: businessService,
	}
}

func (ctrl *BusinessController) SignUp(params business.SignUpAdminParams) middleware.Responder {
	if params.SignUpAdminRequest == nil {
		return business.NewSignUpAdminBadRequest()
	}

	ctx := params.HTTPRequest.Context()

	_, err := ctrl.businessService.CreateBusiness(ctx, *params.SignUpAdminRequest)
	if err != nil {
		code, convertedError := ctrl.ConvertErrorToSwaggerError(err)
		switch code {
		case codes.InvalidInput:
			return business.NewSignUpAdminBadRequest().WithPayload(&convertedError)
		default:
			return business.NewSignUpAdminInternalServerError().WithPayload(&convertedError)
		}
	}

	return business.NewSignUpAdminNoContent()
}

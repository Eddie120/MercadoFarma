package controllers

import (
	"github.com/mercadofarma/services/codes"
	"github.com/mercadofarma/services/errors"
	swaggerModels "github.com/mercadofarma/services/models"
)

type BaseController struct{}

func (baseController *BaseController) ConvertErrorToSwaggerError(err error) (codes.ErrCode, swaggerModels.Error) {
	if err == nil {
		return "", swaggerModels.Error{}
	}

	defaultCode := codes.InternalServerError
	customErr, ok := err.(*errors.Error)
	if !ok {
		return defaultCode, swaggerModels.Error{
			Code:    string(defaultCode),
			Message: err.Error(),
		}
	}

	return customErr.Code, swaggerModels.Error{
		Code:    string(customErr.Code),
		Message: customErr.Message,
	}
}

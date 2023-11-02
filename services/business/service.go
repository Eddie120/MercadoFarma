package business

import (
	"context"
	"github.com/mercadofarma/services/codes"
	"github.com/mercadofarma/services/errors"
	swaggerModels "github.com/mercadofarma/services/models"
	"github.com/mercadofarma/services/repos/business"
	"github.com/mercadofarma/services/repos/models"
	userService "github.com/mercadofarma/services/services/users"
	"github.com/nyaruka/phonenumbers"
	"strings"
	"time"
)

type BusinessService interface {
	CreateBusiness(ctx context.Context, record swaggerModels.SignUpAdminRequest) (*models.Business, error)
	ValidateBusinessRecord(record *swaggerModels.SignUpAdminRequest) error
}

type ServiceImpl struct {
	businessRepo business.BusinessRepo
	userService  userService.UserService
}

func NewBusinessService(businessRepo business.BusinessRepo, userService userService.UserService) *ServiceImpl {
	return &ServiceImpl{
		businessRepo: businessRepo,
		userService:  userService,
	}
}

func (svc *ServiceImpl) CreateBusiness(ctx context.Context, record swaggerModels.SignUpAdminRequest) (*models.Business, error) {
	if err := svc.ValidateBusinessRecord(&record); err != nil {
		return nil, err
	}

	tx, err := svc.businessRepo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	email := record.Email
	password := record.Password
	firstName := record.FirstName
	lastName := record.LastName
	phoneNumber := record.PhoneNumber

	role := ""
	if record.Role != nil {
		role = *record.Role
	}

	user, err := svc.userService.CreateUser(ctx, string(email), password, firstName, lastName, role, phoneNumber)
	if err != nil {
		return nil, err
	}

	newBusiness := &models.Business{
		TaxId:        record.TaxID,
		Name:         record.BusinessName,
		BusinessType: models.BusinessType(record.BusinessType),
		Address:      record.Address,
		PhoneNumber:  record.PhoneNumber,
		UserId:       user.UserId,
		SectorId:     record.SectorID,
		Active:       true,
	}

	err = svc.businessRepo.CreateBusiness(ctx, newBusiness)
	if err != nil {
		return nil, err
	}

	err = svc.businessRepo.CreateBusinessOpeningHours(ctx, newBusiness.BusinessId, convertToBusinessOpeningHours(record.BusinessOpeningHours))
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return newBusiness, nil
}

func (svc *ServiceImpl) ValidateBusinessRecord(record *swaggerModels.SignUpAdminRequest) error {
	if record == nil {
		return errors.ErrorWithCode(codes.InvalidInput, "business record can not be nil")
	}

	if len(strings.Trim(record.TaxID, " ")) == 0 {
		return errors.ErrorWithCode(codes.InvalidInput, "tax id is required")
	}

	if len(strings.Trim(record.BusinessName, " ")) == 0 {
		return errors.ErrorWithCode(codes.InvalidInput, "business name is required")
	}

	if !models.IsValidBusinessTypes[models.BusinessType(record.BusinessType)] {
		return errors.ErrorWithCode(codes.InvalidInput, "invalid business type")
	}

	if len(strings.Trim(record.Address, " ")) == 0 {
		return errors.ErrorWithCode(codes.InvalidInput, "address is required")
	}

	_, err := phonenumbers.Parse(record.PhoneNumber, "COL")
	if len(strings.Trim(record.PhoneNumber, " ")) == 0 || err != nil {
		return errors.ErrorWithCode(codes.InvalidInput, "invalid phone number")
	}

	if err = checkBusinessOpeningHours(record.BusinessOpeningHours); err != nil {
		return err
	}

	return nil
}

func checkBusinessOpeningHours(hours []*swaggerModels.BusinessOpeningHour) error {
	for _, record := range hours {
		if !models.IsValidDayOfWeek[models.Day(record.Day)] {
			return errors.ErrorWithCode(codes.InvalidInput, "invalid day of week")
		}

		startTime, err1 := time.Parse(time.TimeOnly, record.StartTime)
		if err1 != nil {
			return errors.ErrWithCode(codes.InvalidInput, err1)
		}

		endingTime, err2 := time.Parse(time.TimeOnly, record.EndingTime)
		if err2 != nil {
			return errors.ErrWithCode(codes.InvalidInput, err2)
		}

		if startTime.After(endingTime) {
			return errors.ErrorWithCode(codes.InvalidInput, "invalid opening hours")
		}
	}

	return nil
}

func convertToBusinessOpeningHours(hours []*swaggerModels.BusinessOpeningHour) []models.BusinessOpeningHour {
	data := make([]models.BusinessOpeningHour, 0)
	for _, record := range hours {
		startTime, _ := time.Parse(time.TimeOnly, record.StartTime)
		endingTime, _ := time.Parse(time.TimeOnly, record.EndingTime)

		data = append(data, models.BusinessOpeningHour{
			Day:        record.Day,
			StartTime:  &startTime,
			EndingTime: &endingTime,
		})
	}

	return data
}

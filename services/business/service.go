package business

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	swaggerModels "github.com/mercadofarma/services/models"
	"github.com/mercadofarma/services/repos/business"
	"github.com/mercadofarma/services/repos/models"
	"github.com/nyaruka/phonenumbers"
	"strings"
	"time"
)

type BusinessService interface {
	CreateBusiness(ctx context.Context, userId string, sectorId string, record swaggerModels.SignUpAdminRequest) (*dynamodb.PutItemOutput, error)
	ValidateBusinessRecord(record *swaggerModels.SignUpAdminRequest) error
}

type ServiceImpl struct {
	businessRepo business.BusinessRepo
}

var (
	ErrInvalidBusinessRecord = errors.New("invalid business record")
	ErrInvalidTaxId          = errors.New("invalid tax id")
	ErrInvalidBusinessName   = errors.New("invalid business name")
	ErrInvalidBusinessType   = errors.New("invalid business type")
	ErrInvalidAddress        = errors.New("invalid address")
	ErrInvalidPhoneNumber    = errors.New("invalid phone number")
)

func NewBusinessService(businessRepo business.BusinessRepo) *ServiceImpl {
	return &ServiceImpl{
		businessRepo: businessRepo,
	}
}

func (svc *ServiceImpl) CreateBusiness(ctx context.Context, userId string, sectorId string, record swaggerModels.SignUpAdminRequest) (*dynamodb.PutItemOutput, error) {
	if err := svc.ValidateBusinessRecord(&record); err != nil {
		return nil, err
	}

	newBusiness := &models.Business{
		TaxId:                record.TaxID,
		Name:                 record.BusinessName,
		BusinessType:         models.BusinessType(record.BusinessType),
		Address:              record.Address,
		PhoneNumber:          record.PhoneNumber,
		UserId:               userId,
		SectorId:             sectorId,
		Active:               true,
		BusinessOpeningHours: convertToBusinessOpeningHours(record.BusinessOpeningHours),
	}

	return svc.businessRepo.CreateBusiness(ctx, newBusiness)
}

func (svc *ServiceImpl) ValidateBusinessRecord(record *swaggerModels.SignUpAdminRequest) error {
	if record == nil {
		return ErrInvalidBusinessRecord
	}

	if len(strings.Trim(record.TaxID, " ")) == 0 {
		return ErrInvalidTaxId
	}

	if len(strings.Trim(record.BusinessName, " ")) == 0 {
		return ErrInvalidBusinessName
	}

	if !models.IsValidBusinessTypes[models.BusinessType(record.BusinessType)] {
		return ErrInvalidBusinessType
	}

	if len(strings.Trim(record.Address, " ")) == 0 {
		return ErrInvalidAddress
	}

	_, err := phonenumbers.Parse(record.PhoneNumber, "COL")
	if len(strings.Trim(record.PhoneNumber, " ")) == 0 || err != nil {
		return ErrInvalidPhoneNumber
	}

	// we need to check if the userid and sectorid exist

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

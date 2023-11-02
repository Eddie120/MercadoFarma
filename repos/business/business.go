package business

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/mercadofarma/services/db/mysql"
	"github.com/mercadofarma/services/repos/models"
	"time"
)

type BusinessRepo interface {
	CreateBusiness(ctx context.Context, businessRecord *models.Business) error
	CreateBusinessOpeningHours(ctx context.Context, businessId string, records []models.BusinessOpeningHour) error
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type BusinessRepoImpl struct {
	db        mysql.DataAccess
	tableName string
}

const (
	businessPrefix             = "BNS"
	insertStatement            = `INSERT INTO business (business_id,user_id,sector_id,tax_id,company_name,business_type,active,creation_date,update_date,address,phone_number) values (?,?,?,?,?,?,?,?,?,?,?);`
	insertBusinessOpeningHours = `INSERT INTO business_opening_hours  (business_id,day,start_time,ending_time) values (?,?,?,?);`
)

func NewBusinessRepo(db mysql.DataAccess) *BusinessRepoImpl {
	return &BusinessRepoImpl{
		db:        db,
		tableName: "business",
	}
}

func (svc *BusinessRepoImpl) BeginTx(ctx context.Context) (*sql.Tx, error) {
	err := svc.db.Ping()
	if err != nil {
		return nil, err
	}

	return svc.db.BeginTx(ctx)
}

func (svc *BusinessRepoImpl) CreateBusiness(ctx context.Context, record *models.Business) error {
	current := time.Now()
	record.BusinessId = fmt.Sprintf("%s-%s", businessPrefix, uuid.New().String())
	record.CreationDate = &current
	record.UpdateDate = &current

	params := []interface{}{
		record.BusinessId,
		record.UserId,
		record.SectorId,
		record.TaxId,
		record.Name,
		record.BusinessType,
		record.Active,
		record.CreationDate,
		record.UpdateDate,
		record.Address,
		record.PhoneNumber,
	}

	_, err := svc.db.ExecWithContext(ctx, insertStatement, params...)
	if err != nil {
		return err
	}

	return nil
}

func (svc *BusinessRepoImpl) CreateBusinessOpeningHours(ctx context.Context, businessId string, records []models.BusinessOpeningHour) error {
	for _, record := range records {
		params := []interface{}{
			businessId,
			record.Day,
			record.StartTime.Format(time.TimeOnly),
			record.EndingTime.Format(time.TimeOnly),
		}

		_, err := svc.db.ExecWithContext(ctx, insertBusinessOpeningHours, params...)
		if err != nil {
			return err
		}
	}

	return nil
}

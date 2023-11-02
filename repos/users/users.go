package users

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mercadofarma/services/db/mysql"
	"github.com/mercadofarma/services/repos/models"
	"time"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
}

const (
	userPrefix = "USR"
)

type UserRepoImpl struct {
	db        mysql.DataAccess
	tableName string
}

func NewUserRepo(db mysql.DataAccess) *UserRepoImpl {
	return &UserRepoImpl{
		db:        db,
		tableName: "users",
	}
}

func (svc *UserRepoImpl) CreateUser(ctx context.Context, user *models.User) error {
	current := time.Now()
	user.UserId = fmt.Sprintf("%s-%s", userPrefix, uuid.New().String())
	user.CreationDate = &current
	user.UpdateDate = &current

	const query = "INSERT INTO users (user_id,email,hash,role,active,creation_date,update_date) values (?,?,?,?,?,?,?);"
	params := []interface{}{
		user.UserId,
		user.Email,
		user.Hash,
		user.Role,
		user.Active,
		user.CreationDate,
		user.UpdateDate,
	}

	_, err := svc.db.ExecWithContext(ctx, query, params...)
	if err != nil {
		return err
	}

	return nil
}

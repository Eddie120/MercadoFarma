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
	CheckIfUserExist(ctx context.Context, email string, role models.Role) bool
	GetUserByEmail(ctx context.Context, email string, role models.Role) (*models.User, error)
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

func (svc *UserRepoImpl) CheckIfUserExist(ctx context.Context, email string, role models.Role) bool {
	user, err := svc.GetUserByEmail(ctx, email, role)
	if err != nil {
		return true
	}

	return user.UserId != ""
}

func (svc *UserRepoImpl) GetUserByEmail(ctx context.Context, email string, role models.Role) (*models.User, error) {
	user := models.User{}
	params := []interface{}{
		email,
		role,
	}

	const query = "SELECT user_id,email,first_name,last_name,hash,role,active,creation_date,update_date FROM users WHERE email = ? AND ROLE = ?;"

	rows, err := svc.db.QueryWithContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&user.UserId, &user.Email, &user.FirstName, &user.LastName, &user.Hash, &user.Role, &user.Active, &user.CreationDate, &user.UpdateDate); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (svc *UserRepoImpl) CreateUser(ctx context.Context, user *models.User) error {
	current := time.Now()
	user.UserId = fmt.Sprintf("%s-%s", userPrefix, uuid.New().String())
	user.CreationDate = &current
	user.UpdateDate = &current

	const query = "INSERT INTO users (user_id,email,hash,first_name,last_name,role,active,creation_date,update_date) values (?,?,?,?,?,?,?,?,?);"
	params := []interface{}{
		user.UserId,
		user.Email,
		user.Hash,
		user.FirstName,
		user.LastName,
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

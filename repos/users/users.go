package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/mercadofarma/services/repos/models"
	"time"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, userId string) (*models.User, error)
}

var ErrUserNotFound = errors.New("user not found")

const userPrefix = "USR"

type UserRepoImpl struct {
	client    *dynamodb.Client
	tableName string
}

func NewUserRepo(client *dynamodb.Client) *UserRepoImpl {
	return &UserRepoImpl{
		client:    client,
		tableName: "users",
	}
}

func (svc *UserRepoImpl) CreateUser(ctx context.Context, user *models.User) error {
	current := time.Now()
	user.UserId = fmt.Sprintf("%s%s", userPrefix, uuid.New().String())
	user.CreationDate = &current
	user.UpdateDate = &current

	item, _ := attributevalue.MarshalMap(user)
	_, err := svc.client.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(svc.tableName),
	})

	if err != nil {
		return err
	}

	return nil
}

func (svc *UserRepoImpl) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	result, err := svc.client.GetItem(ctx, &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{
				Value: userId,
			},
		},
		TableName: aws.String(svc.tableName),
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, ErrUserNotFound
	}

	user := models.User{}

	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

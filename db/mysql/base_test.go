package mysql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewDataAccess(t *testing.T) {
	c := assert.New(t)

	db, _, err := sqlmock.New()
	c.Nil(err)

	defer db.Close()

	dataAccess, err := NewDataAccess(db, "mysql", "user:password@/dbname")
	c.Nil(err)
	c.NotNil(dataAccess)
}

func TestDataStore_Ping(t *testing.T) {
	c := assert.New(t)

	db, _, err := sqlmock.New()
	c.Nil(err)

	defer db.Close()

	dataAccess, err := NewDataAccess(db, "mysql", "user:password@/dbname")
	err = dataAccess.Ping()
	c.Nil(err)
}

func TestDataStore_ExecWithContext(t *testing.T) {
	c := assert.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)

	defer db.Close()

	const expectedQuery = "INSERT INTO product_viewers"
	mock.ExpectExec(expectedQuery).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(1, 1))

	dataAccess, err := NewDataAccess(db, "mysql", "user:password@/dbname")
	c.Nil(err)

	ctx := context.Background()
	result, err := dataAccess.ExecWithContext(ctx, expectedQuery, 1, 2)
	c.Nil(err)

	lastId, _ := result.LastInsertId()
	affectedRows, _ := result.RowsAffected()

	c.Equal(int64(1), lastId)
	c.Equal(int64(1), affectedRows)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDataStore_QueryWithContext(t *testing.T) {
	c := assert.New(t)

	db, mock, err := sqlmock.New()
	c.Nil(err)

	defer db.Close()

	const categoryID = 10
	const expectedQuery = "SELECT id,product_name,description,category_id FROM products WHERE category_id=?"
	rows := sqlmock.NewRows([]string{"id", "product_name", "description", "category_id"}).
		AddRow(1, "product name 1", "description 1", 10).
		AddRow(2, "product name 2", "description 2", 10).
		AddRow(3, "product name 3", "description 3", 10).
		AddRow(4, "product name 4", "description 4", 10)
	mock.ExpectQuery(expectedQuery).WithArgs(categoryID).WillReturnRows(rows)

	dataAccess, err := NewDataAccess(db, "mysql", "user:password@/dbname")
	c.Nil(err)

	ctx := context.Background()
	rows_, err := dataAccess.QueryWithContext(ctx, expectedQuery, categoryID)
	c.Nil(err)

	type product struct {
		Id          string `sql:"id"`
		ProductName string `sql:"product_name"`
		Description string `sql:"description"`
		CategoryId  int    `sql:"category_id"`
	}

	products := make([]*product, 0)
	for rows_.Next() {
		var item = &product{}
		if err := rows_.Scan(&item.Id, &item.ProductName, &item.Description, &item.CategoryId); err != nil {
			log.Fatal(err)
		}

		products = append(products, item)
	}

	c.Len(products, 4)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

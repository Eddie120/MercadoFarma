package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"log"
	"net"
	"time"
)

const dbIsNil = "db is nil"

type DataAccess interface {
	Ping() error
	ExecWithContext(ctx context.Context, query string, arg ...interface{}) (sql.Result, error)
	QueryWithContext(ctx context.Context, query string, arg ...interface{}) (*sql.Rows, error)
	Close() error
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type DataStore struct {
	Db         *sql.DB
	DriverName string
	DataSource string
	Database   string
	Host       string
	Port       string
	log        *log.Logger
}

func CreateDBConnection(driverName, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func NewDataAccess(db *sql.DB, driverName, dataSource string) DataAccess {
	dsn, err := mysql.ParseDSN(dataSource)
	if err != nil {
		dsn = &mysql.Config{}
	}

	host, port, err := net.SplitHostPort(dsn.Addr)
	if err != nil {
		port = "3306"
	}

	return &DataStore{
		Db:         db,
		DriverName: driverName,
		DataSource: dataSource, // connection string
		Database:   dsn.DBName,
		Host:       host,
		Port:       port,
		log:        log.Default(),
	}
}

func (d *DataStore) Ping() error {
	if d.Db != nil {
		return d.Db.Ping()
	} else {
		return errors.New(dbIsNil)
	}
}

func (d *DataStore) ExecWithContext(ctx context.Context, query string, arg ...interface{}) (sql.Result, error) {
	if d.Db == nil {
		return nil, errors.New(dbIsNil)
	}

	return d.Db.ExecContext(ctx, query, arg...)
}

func (d *DataStore) QueryWithContext(ctx context.Context, query string, arg ...interface{}) (*sql.Rows, error) {
	if d.Db == nil {
		return nil, errors.New(dbIsNil)
	}

	return d.Db.QueryContext(ctx, query, arg...)
}

func (d *DataStore) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return d.Db.BeginTx(ctx, nil)
}

func (d *DataStore) Close() error {
	if d.Db != nil {
		return d.Db.Close()
	} else {
		return errors.New(dbIsNil)
	}
}

package config

import (
	"database/sql"
	"fmt"

	"github.com/Parking/errorApi/logger"
	_ "github.com/go-sql-driver/mysql"
)

const (
	connectionErrorMessage   = "error when trying to connect with mysql database"
	notAvailableErrorMessage = "the database is not available or accessible"
	driverTypeMysql          = "mysql"
)

// GetWriteConnection database connection client
func GetWriteConnection() (sqlDB *sql.DB, err error) {
	connection := getWriteDBConnectionString()
	return getConnection(connection)
}

// GetReadConnection database connection client
func GetReadConnection() (sqlDB *sql.DB, err error) {
	connection := getReadDBConnectionString()
	return getConnection(connection)
}

func getConnection(connection string) (sqlDB *sql.DB, err error) {
	sqlDB, err = sql.Open(driverTypeMysql, connection)
	if err != nil {
		logger.Error(connectionErrorMessage, err)
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		logger.Error(notAvailableErrorMessage, err)
		return
	}

	return
}

func CloseConnections(err error, tx *sql.Tx, stmt *sql.Stmt, rows *sql.Rows) {
	if tx != nil {
		switch err {
		case nil:
			_ = tx.Commit()
		default:
			_ = tx.Rollback()
		}
	}

	if stmt != nil {
		_ = stmt.Close()
	}

	if rows != nil {
		_ = rows.Close()
	}
}

// getReadDBConnectionString gets the Translations DB connection string depending the actual scope for read
func getReadDBConnectionString() string {
	var (
		user, password string
	)

	if IsProductiveScope() {
		user = readDBAdminUser
		password = dbAdminPwd
	} else {
		user = readDBTestUser
		password = readDBTestPwd
	}

	return getDBConnectionString(user, password)
}

func getWriteDBConnectionString() string {
	var (
		user, password string
	)

	if IsProductiveScope() {
		user = writeDBAdminUser
		password = writeDBAdminPwd
	} else {
		user = writeDBTestUser
		password = writeDBTestPwd
	}

	return getDBConnectionString(user, password)
}
func getDBConnectionString(user, password string) string {
	connection := "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true"

	var host, database string

	if IsProductiveScope() {
		host = dbHost
		database = dbName
	} else {
		host = dbTestHost
		database = dbTestName
	}

	return fmt.Sprintf(connection, user, password, host, database)
}

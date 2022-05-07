package AutoFSecBackend

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
)

type DbManager struct {
	m_mainDbConnection *pgx.Conn
}

func (dbConnectionManager *DbManager) ConnectToDb() bool {
	var connectionError error
	dbConnectionManager.m_mainDbConnection, connectionError = pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")

	return connectionError == nil
}

func (dbConnectionManager *DbManager) DisconnectFromDb() {
	if dbConnectionManager.m_mainDbConnection != nil {
		dbConnectionManager.m_mainDbConnection.Close(context.Background())
	}
}

func (dbConnectionManager DbManager) PerformRequest(query string) string {
	var returnFromDb string
	if dbConnectionManager.m_mainDbConnection != nil {
		dbConnectionManager.m_mainDbConnection.QueryRow(context.Background(), query).Scan(&returnFromDb)
	} else {
		fmt.Println("dbConnectionManager::PerformRequest: Db connection error")
	}
	return returnFromDb
}

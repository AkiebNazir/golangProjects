package handler_func

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectoToDB() (*gorm.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "njasm786"
		dbname   = "employees"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, port, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// This function will execute the query on a perticular table
//
// Parameters:
//
// table_name(string): name of the table on which you want to query
//
// query(string): SQL query in a string format
//
// query_type(string): query to be performed e,g [ SELECT, DELETE, INSERT, UPDATE, CREATE ]
//
// RETURN:
//
// []byte data:

type QueryType interface {
	queryTable() (interface{}, error)
}

type SelectQuery struct {
	app       *App
	query     string
	dataStore interface{}
}

func (s *SelectQuery) queryTable() (interface{}, error) {
	err := s.app.Db.Raw(s.query).Scan(&s.dataStore).Error
	if err != nil {
		return nil, err
	}
	return s.dataStore, nil
}

type OtherQueries struct {
	app   *App
	query string
}

func (s *OtherQueries) queryTable() (interface{}, error) {
	err := s.app.Db.Exec(s.query).Error
	if err != nil {
		return nil, err
	}
	msg := "Sql Transaction run successfully"
	s.app.Db.Commit()
	return msg, nil

}

func (app *App) CreateTable() (interface{}, error) {

	var queryType = OtherQueries{
		app: app,
		query: `CREATE TABLE IF NOT EXISTS employees_details(
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(200),
			contact VARCHAR(20) NOT NULL,
			expreience int
		)`,
	}
	data, err := queryType.queryTable()
	if err != nil {
		return nil, err
	}
	return data, nil

}

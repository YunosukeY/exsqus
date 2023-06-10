package util

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func GetConfig() (*mysql.Config, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	protocol := os.Getenv("MYSQL_PROTOCOL")

	if host == "" {
		return nil, fmt.Errorf("The environment variable MYSQL_HOST is required")
	}
	if dbname == "" {
		return nil, fmt.Errorf("The environment variable MYSQL_DATABASE is required")
	}
	if user == "" {
		return nil, fmt.Errorf("The environment variable MYSQL_USER is required")
	}
	if pass == "" {
		return nil, fmt.Errorf("The environment variable MYSQL_PASSWORD is required")
	}

	if port == "" {
		port = "3306"
	}
	if protocol == "" {
		protocol = "tcp"
	}

	return &mysql.Config{
		Addr:   fmt.Sprintf("%s:%s", host, port),
		Net:    protocol,
		DBName: dbname,
		User:   user,
		Passwd: pass,
	}, nil
}

func GetDB(c mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		db.Close()
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

type Row struct {
	Id                                                                                               int
	SelectType, Table, Partitions, AccessType, PossibleKeys, Key, KeyLen, Ref, Rows, Filtered, Extra sql.NullString
}

type Plan struct {
	Rows []Row
	Id   string
}

func GetPlan(db *sql.DB, query string) (*Plan, error) {
	rows, err := db.Query("EXPLAIN " + query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rs := []Row{}
	for rows.Next() {
		r := Row{}

		err := rows.Scan(&r.Id, &r.SelectType, &r.Table, &r.Partitions, &r.AccessType, &r.PossibleKeys, &r.Key, &r.KeyLen, &r.Ref, &r.Rows, &r.Filtered, &r.Extra)
		if err != nil {
			return nil, err
		}

		rs = append(rs, r)
	}
	return &Plan{rs, ""}, nil
}

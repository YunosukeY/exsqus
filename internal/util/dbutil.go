package util

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func getConfig() mysql.Config {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	protocol := os.Getenv("MYSQL_PROTOCOL")

	if port == "" {
		port = "3306"
	}
	if protocol == "" {
		protocol = "tcp"
	}

	return mysql.Config{
		Addr:   fmt.Sprintf("%s:%s", host, port),
		Net:    protocol,
		DBName: dbname,
		User:   user,
		Passwd: pass,
	}
}

func GetDB() (*sql.DB, error) {
	c := getConfig()
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
	id                                                                                               int
	selectType, table, partitions, accessType, possibleKeys, key, keyLen, ref, rows, filtered, extra sql.NullString
}

func GetPlan(db *sql.DB, query string) ([]Row, error) {
	rows, err := db.Query("EXPLAIN " + query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rs := []Row{}
	for rows.Next() {
		r := Row{}
		err := rows.Scan(&r.id, &r.selectType, &r.table, &r.partitions, &r.accessType, &r.possibleKeys, &r.key, &r.keyLen, &r.ref, &r.rows, &r.filtered, &r.extra)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

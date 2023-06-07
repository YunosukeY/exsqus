package util

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	t.Skip()

	var c mysql.Config

	c = GetConfig()
	assert.Equal(t, mysql.Config{Addr: ":3306", Net: "tcp"}, c)

	os.Setenv("MYSQL_HOST", "localhost")
	os.Setenv("MYSQL_PORT", "33060")
	os.Setenv("MYSQL_USER", "user")
	os.Setenv("MYSQL_PASSWORD", "pass")
	os.Setenv("MYSQL_DATABASE", "db")
	os.Setenv("MYSQL_PROTOCOL", "udp")
	c = GetConfig()
	assert.Equal(t, mysql.Config{Addr: "localhost:33060", Net: "udp", DBName: "db", User: "user", Passwd: "pass"}, c)
}

func TestGetPlan(t *testing.T) {
	err := loadEnv()
	if err != nil {
		t.Fatal(err)
	}

	c := GetConfig()
	db, err := GetDB(c)
	if err != nil {
		t.Fatal(err)
	}

	plan, err := GetPlan(db, "SELECT SLEEP(2);")
	assert.Nil(t, err)
	assert.Len(t, plan.Rows, 1)

	st := sql.NullString{}
	st.Scan("SIMPLE")
	e := sql.NullString{}
	e.Scan("No tables used")
	assert.Equal(t, Row{1, st, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, sql.NullString{}, e}, plan.Rows[0])
}

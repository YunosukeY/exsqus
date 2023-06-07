package util

import (
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
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
}

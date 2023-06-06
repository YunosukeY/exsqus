package main

import (
	"os"

	app "github.com/YunosukeY/explain-slow-query/internal"
)

func main() {
	os.Setenv("MYSQL_HOST", "localhost")
	os.Setenv("MYSQL_USER", "root")
	os.Setenv("MYSQL_PASSWORD", "root")
	os.Setenv("MYSQL_DATABASE", "test")
	os.Setenv("LOG_FILE_PATH", "/Users/kimitsu/work/slow-query-trigger-explain/logs/slow.log")

	app.Run()
}

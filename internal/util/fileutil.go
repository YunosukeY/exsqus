package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func GetLogFilePath() string {
	path := os.Getenv("LOG_FILE_PATH")
	if path == "" {
		path = "/tmp/slow.log"
	}

	return path
}

var timePattern = regexp.MustCompile(`# Time: (.*)`)
var queryTimePattern = regexp.MustCompile(`# Query_time: ([\d\.]*).*Lock_time: ([\d\.]*).*Rows_sent: (\d*).*Rows_examined: (\d+)`)

type Log struct {
	Time, QueryTime, LockTime, RowsSent, RowsExamined, Query string
}

func GetLastQueryLog(reader *bufio.Reader) (*Log, error) {
	var time, queryTime, lockTime, rowsSent, rowsExamined, query string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if timePattern.MatchString(line) {
			match := timePattern.FindStringSubmatch(line)
			time = match[1]
		} else if queryTimePattern.MatchString(line) {
			match := queryTimePattern.FindStringSubmatch(line)
			queryTime = match[1]
			lockTime = match[2]
			rowsSent = match[3]
			rowsExamined = match[4]
		} else {
			query = line
		}
	}

	if time == "" || queryTime == "" || lockTime == "" || rowsSent == "" || rowsExamined == "" || query == "" {
		return nil, fmt.Errorf("No query log found")
	}
	return &Log{time, queryTime, lockTime, rowsSent, rowsExamined, query}, nil
}

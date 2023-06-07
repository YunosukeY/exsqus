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

type Log struct {
	Id, Time, QueryTime, LockTime, RowsSent, RowsExamined, Query string
}

func GetLastQueryLog(reader *bufio.Reader) (*Log, error) {
	var line string
	var err error

	// A slow query log consists of five rows.

	// The first row is like "# Time: 2023-06-07T11:58:58.688716Z".
	line, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	time := getTime(line)
	if time == nil {
		return nil, fmt.Errorf("Time not found")
	}

	// The second row is like "# User@Host: root[root] @  [192.168.16.1]  Id:    10".
	// Skip it.
	if _, err = reader.ReadString('\n'); err != nil {
		return nil, err
	}

	// The third row is like "# Query_time: 2.001390  Lock_time: 0.000000 Rows_sent: 1  Rows_examined: 1".
	line, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	queryTime, lockTime, rowsSent, rowsExamined := getQueryTime(line)
	if queryTime == nil {
		return nil, fmt.Errorf("Query_time, Lock_time, Rows_sent, Rows_examined not found")
	}

	// The forth row is like "SET timestamp=1686139143;".
	// Skip it.
	if _, err = reader.ReadString('\n'); err != nil {
		return nil, err
	}

	// The last row is a query.
	line, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	query := getQuery(line)
	if query == nil {
		return nil, fmt.Errorf("Query not found")
	}

	// Skip remaining rows
	for {
		_, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return &Log{"", *time, *queryTime, *lockTime, *rowsSent, *rowsExamined, *query}, nil
}

var timePattern = regexp.MustCompile(`# Time: (.*)`)

func getTime(line string) *string {
	if timePattern.MatchString(line) {
		match := timePattern.FindStringSubmatch(line)
		return &match[1]
	}
	return nil
}

var queryTimePattern = regexp.MustCompile(`# Query_time: ([\d\.]*).*Lock_time: ([\d\.]*).*Rows_sent: (\d*).*Rows_examined: (\d+)`)

func getQueryTime(line string) (*string, *string, *string, *string) {
	if queryTimePattern.MatchString(line) {
		match := queryTimePattern.FindStringSubmatch(line)
		return &match[1], &match[2], &match[3], &match[4]
	}
	return nil, nil, nil, nil
}

var queryPattern = regexp.MustCompile(`(.*)\n`)

func getQuery(line string) *string {
	if queryPattern.MatchString(line) {
		match := queryPattern.FindStringSubmatch(line)
		return &match[1]
	}
	return nil
}

package util

import (
	"bufio"
	"fmt"
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
	// A slow query log consists of five rows.

	// The first row is like "# Time: 2023-06-07T11:58:58.688716Z".
	time, err := getTime(reader)
	if err != nil {
		return nil, err
	}

	// The second row is like "# User@Host: root[root] @  [192.168.16.1]  Id:    10".
	// Skip it.
	if _, err = reader.ReadString('\n'); err != nil {
		return nil, err
	}

	// The third row is like "# Query_time: 2.001390  Lock_time: 0.000000 Rows_sent: 1  Rows_examined: 1".
	queryTime, lockTime, rowsSent, rowsExamined, err := getQueryTime(reader)
	if err != nil {
		return nil, err
	}

	// The forth row is like "SET timestamp=1686139143;".
	// Skip it.
	if _, err = reader.ReadString('\n'); err != nil {
		return nil, err
	}

	// The last row is a query.
	query, err := getQuery(reader)
	if err != nil {
		return nil, err
	}

	// Skip remaining rows
	if err = SkipAll(reader); err != nil {
		return nil, err
	}

	return &Log{"", time, queryTime, lockTime, rowsSent, rowsExamined, query}, nil
}

var timePattern = regexp.MustCompile(`# Time: (.*)`)

func getTime(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if timePattern.MatchString(line) {
		match := timePattern.FindStringSubmatch(line)
		return match[1], nil
	}
	return "", fmt.Errorf("Time not found")
}

var queryTimePattern = regexp.MustCompile(`# Query_time: ([\d\.]*).*Lock_time: ([\d\.]*).*Rows_sent: (\d*).*Rows_examined: (\d+)`)

func getQueryTime(reader *bufio.Reader) (string, string, string, string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", "", "", "", err
	}

	if queryTimePattern.MatchString(line) {
		match := queryTimePattern.FindStringSubmatch(line)
		return match[1], match[2], match[3], match[4], nil
	}
	return "", "", "", "", fmt.Errorf("Query_time, Lock_time, Rows_sent, Rows_examined not found")
}

var queryPattern = regexp.MustCompile(`(.*)\n`)

func getQuery(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	if queryPattern.MatchString(line) {
		match := queryPattern.FindStringSubmatch(line)
		return match[1], nil
	}
	return "", fmt.Errorf("Query not found")
}

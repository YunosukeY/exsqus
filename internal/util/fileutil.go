package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func GetLogFilePaths() []string {
	paths := os.Getenv("LOG_FILE_PATH")
	if paths == "" {
		paths = "/tmp/slow.log"
	}

	return strings.Split(paths, ",")
}

type Log struct {
	Id, Time, QueryTime, LockTime, RowsSent, RowsExamined, Query string
}

func GetLastQueryLog(reader io.Reader) (*Log, error) {
	scanner := bufio.NewScanner(reader)

	// A slow query log consists of five rows.

	// The first row is like "# Time: 2023-06-07T11:58:58.688716Z".
	time, err := getTime(scanner)
	if err != nil {
		return nil, err
	}

	// The second row is like "# User@Host: root[root] @  [192.168.16.1]  Id:    10".
	// Skip it.
	if !scanner.Scan() {
		return nil, fmt.Errorf("Failed to get new row")
	}

	// The third row is like "# Query_time: 2.001390  Lock_time: 0.000000 Rows_sent: 1  Rows_examined: 1".
	queryTime, lockTime, rowsSent, rowsExamined, err := getQueryTime(scanner)
	if err != nil {
		return nil, err
	}

	// The forth row is like "SET timestamp=1686139143;".
	// Skip it.
	if !scanner.Scan() {
		return nil, fmt.Errorf("Failed to get new row")
	}

	// The last row is a query.
	query, err := getQuery(scanner)
	if err != nil {
		return nil, err
	}

	// Skip remaining rows
	SkipAll(reader)

	return &Log{"", time, queryTime, lockTime, rowsSent, rowsExamined, query}, nil
}

var timePattern = regexp.MustCompile(`# Time: (.*)`)

func getTime(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", fmt.Errorf("Failed to get new row")
	}
	line := scanner.Text()

	if timePattern.MatchString(line) {
		match := timePattern.FindStringSubmatch(line)
		return match[1], nil
	}
	return "", fmt.Errorf("Time not found")
}

var queryTimePattern = regexp.MustCompile(`# Query_time: ([\d\.]*).*Lock_time: ([\d\.]*).*Rows_sent: (\d*).*Rows_examined: (\d+)`)

func getQueryTime(scanner *bufio.Scanner) (string, string, string, string, error) {
	if !scanner.Scan() {
		return "", "", "", "", fmt.Errorf("Failed to get new row")
	}
	line := scanner.Text()

	if queryTimePattern.MatchString(line) {
		match := queryTimePattern.FindStringSubmatch(line)
		return match[1], match[2], match[3], match[4], nil
	}
	return "", "", "", "", fmt.Errorf("Query_time, Lock_time, Rows_sent, Rows_examined not found")
}

func getQuery(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", fmt.Errorf("Failed to get new row")
	}
	return scanner.Text(), nil
}

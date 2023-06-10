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
	Id, QueryTime, LockTime, RowsSent, RowsExamined, Query string
}

func GetLastQueryLog(reader io.Reader) (*Log, error) {
	scanner := bufio.NewScanner(reader)

	// Usually, a slow query log consists of five rows.
	// when log-short-format is enabled, it consists of three rows.

	scanner.Scan()
	infoRow := scanner.Text()
	if strings.HasPrefix(infoRow, "# Time:") {
		scanner.Scan()
		scanner.Scan()
		infoRow = scanner.Text()
	}

	// Expect  "# Query_time: 2.001390  Lock_time: 0.000000 Rows_sent: 1  Rows_examined: 1".
	queryTime, lockTime, rowsSent, rowsExamined, err := getExecInfo(infoRow)
	if err != nil {
		return nil, err
	}

	// Expect "SET timestamp=1686139143;".
	// Skip it.
	if !scanner.Scan() {
		return nil, fmt.Errorf("Failed to get new row")
	}

	// Expect a query.
	query, err := getQuery(scanner)
	if err != nil {
		return nil, err
	}

	return &Log{"", queryTime, lockTime, rowsSent, rowsExamined, query}, nil
}

var execInfoPattern = regexp.MustCompile(`# Query_time: ([\d\.]*).*Lock_time: ([\d\.]*).*Rows_sent: (\d*).*Rows_examined: (\d+)`)

func getExecInfo(line string) (string, string, string, string, error) {
	if execInfoPattern.MatchString(line) {
		match := execInfoPattern.FindStringSubmatch(line)
		return match[1], match[2], match[3], match[4], nil
	}
	return "", "", "", "", fmt.Errorf("Query_time, Lock_time, Rows_sent, Rows_examined not found")
}

func getQuery(scanner *bufio.Scanner) (string, error) {
	var query string
	for scanner.Scan() {
		query += scanner.Text()
	}

	if query == "" {
		return "", fmt.Errorf("Failed to get query")
	}
	return query, nil
}

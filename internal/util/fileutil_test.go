package util

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastQueryLog(t *testing.T) {
	test := `# Time: 2023-06-07T11:59:05.164053Z
# User@Host: root[root] @  [192.168.16.1]  Id:    12
# Query_time: 2.001390  Lock_time: 0.000000 Rows_sent: 1  Rows_examined: 1
SET timestamp=1686139143;
SELECT
SLEEP(2);
`
	reader := bufio.NewReader(strings.NewReader(test))

	log, err := GetLastQueryLog(reader)
	assert.Nil(t, err)
	assertEqual(t, &Log{QueryTime: "2.001390", LockTime: "0.000000", RowsSent: "1", RowsExamined: "1", Query: "SELECT SLEEP(2);"}, log)

	_, err = reader.ReadString('\n')
	assertEqual(t, io.EOF, err)
}

func TestGetQueryTime(t *testing.T) {
	test := "# Query_time: 2.001390  Lock_time: 0.000000 Rows_sent: 1  Rows_examined: 1"

	queryTime, lockTime, rowsSent, rowsExamined, err := getExecInfo(test)
	assert.Nil(t, err)
	assertEqual(t, "2.001390", queryTime)
	assertEqual(t, "0.000000", lockTime)
	assertEqual(t, "1", rowsSent)
	assertEqual(t, "1", rowsExamined)
}

func TestGetQuery(t *testing.T) {
	test := "SELECT SLEEP(2);\n"
	scanner := bufio.NewScanner(strings.NewReader(test))

	time, err := getQuery(scanner)
	assert.Nil(t, err)
	assertEqual(t, "SELECT SLEEP(2);", time)
}

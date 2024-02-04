package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'RemainderSorting' function below (and other code for sorting if needed).
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */

// A custom type that implements sort.Interface
type ByRemainderAndName []string

func (s ByRemainderAndName) Len() int {
	return len(s)
}
func (s ByRemainderAndName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByRemainderAndName) Less(i, j int) bool {
	// Compare the remainders of the lengths when divided by 3
	remainderI, remainderJ := len(s[i])%3, len(s[j])%3
	if remainderI == remainderJ {
		// If the remainders are equal, sort alphabetically
		return s[i] < s[j]
	}
	// Otherwise, sort by the remainder
	return remainderI < remainderJ
}

func RemainderSorting(strArr []string) []string {
	sort.Sort(ByRemainderAndName(strArr))
	return strArr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	strArrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var strArr []string

	for i := 0; i < int(strArrCount); i++ {
		strArrItem := readLine(reader)
		strArr = append(strArr, strArrItem)
	}

	result := RemainderSorting(strArr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

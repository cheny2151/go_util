package untils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(path string) ([][]string, bool) {
	inputStream, err := os.Open(path)
	if err != nil {
		return dealReadCsvError(err)
	}
	defer inputStream.Close()
	reader := csv.NewReader(bufio.NewReader(inputStream))
	reader.Comma = ';'
	reader.LazyQuotes = true
	all, err := reader.ReadAll()
	if err != nil {
		return dealReadCsvError(err)
	}
	return all, true
}

func dealReadCsvError(err error) ([][]string, bool) {
	fmt.Print("io异常:", err)
	return nil, false
}

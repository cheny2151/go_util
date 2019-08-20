package untils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type ExcelReader struct {
	filepath string
	file     *excelize.File
	e        error
}

func ReadExcel(filepath string) (ExcelReader, error) {
	file, e := excelize.OpenFile(filepath)
	return ExcelReader{filepath, file, e}, e
}

func (reader *ExcelReader) ReadCell(cellNumber []string, startRow int, endRow int) [][]interface{} {
	e := reader.e
	if e != nil {
		fmt.Println(e)
		return nil
	}

	file := reader.file

	values := make([][]interface{}, endRow-startRow+1)
	index := 0
	for i := startRow; i <= endRow; i++ {
		cellToFind := make([]string, len(cellNumber))
		for i2 := range cellNumber {
			cell := cellNumber[i2] + strconv.Itoa(i)
			cellToFind[i2] = cell
		}

		value := make([]interface{}, len(cellToFind))
		for j := 0; j < len(cellToFind); j++ {
			s, e := file.GetCellValue("sheet", cellToFind[j])
			if e != nil {
				fmt.Println(e)
				return nil
			}
			value[j] = s
		}
		values[index] = value
		index++
	}
	return values
}

package main

import (
	"com/cheney/untils"
	"fmt"
)

var cellArray = []string{"A", "B", "C", "D", "H", "I", "J", "K", "L", "M"}
var filepath = "C:\\Users\\Cheney\\Documents\\base_repository.xlsx"

func main() {
	reader, e := untils.ReadExcel(filepath)

	if e != nil {
		fmt.Print(e)
		return
	}

	sql := "INSERT INTO `base_repository` (`id`, `code`, `name`, `repository_type`, `province`, `city`, `district`, `detail_address`, `data_source`," +
		" `default_repo_location_code`, `ship_doc_type`, `status`, `del_flag`, `create_time`, `update_time`)" +
		" VALUES (?, ?, ?, ?, NULL, NULL, NULL, ?, ?, ?, ?, ?,?, sysdate(), sysdate());"

	db, _ := untils.OpenConnection()
	stmt, _ := db.Prepare(sql)

	cell := reader.ReadCell(cellArray, 30677, 30677)
	for i := range cell {
		args := cell[i]
		// 传入数组/切片转不定参数时，加上...
		result, e := stmt.Exec(args...)
		if e != nil {
			fmt.Println(e)
		} else {
			id, _ := result.LastInsertId()
			fmt.Println(id * 100 / 30676)
		}
	}
	e = db.Close()

}

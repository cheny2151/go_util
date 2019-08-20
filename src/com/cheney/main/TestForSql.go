package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*db, e := sql.Open("mysql", "root:cheney@mysql@tcp(localhost:3306)/lexue")
	if dealError(e) {
		return
	}
	fmt.Println("connect to mysql:success")
	rows,e :=db.Query("select id,grade,cost_id,remark from t_hl_grade_cost")
	if dealError(e) {
		return
	}
	for rows.Next()  {
		var gc entiry.GradeCost
		e = rows.Scan(&gc.Id,&gc.Grade,&gc.CostId,&gc.Remark)
		if dealError(e) {
			return
		}
		bytes, _ := json.Marshal(gc)
		fmt.Print(string(bytes))
		fmt.Println()
	}*/

	// insert
	db, e := sql.Open("mysql", "finenter:finenter@tcp(192.168.240.20:3306)/finenter")
	if dealError(e) {
		return
	}
	fmt.Println("connect to mysql:success")
	stmt, e := db.Prepare("INSERT INTO invoice_sales_invoice_line (`line_no`, `order_no`, `product_no`, `product_name`, `unit`, `gift_flag`, `amount`, `source_document`, `source_order_no`, `source_line_number`, `del_flag`) VALUES ('10', ?, '111', 'test', '个', '0', '100', 'sale_order', ? , '10', '0')")
	if dealError(e) {
		return
	}
	for i := 1; i < 1000; i++ {
		_, e := stmt.Exec(i, i)
		if e != nil {
			fmt.Println(e)
		}
	}
}

func dealError(err error) bool {
	if err != nil {
		fmt.Print(err)
		return true
	}
	return false
}

package untils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func OpenConnection() (*sql.DB, error) {
	db, e := sql.Open("mysql", "finenter:finenter@tcp(192.168.240.20:3306)/finenter")
	if dealError(e) {
		return nil, e
	}
	return db, nil
}

func Execute(sql string, args []string) (*sql.Rows, sql.Result, error) {

	db, e := OpenConnection()
	if e != nil {
		fmt.Println(e)
		return nil, nil, e
	}
	stmt, e := db.Prepare(sql)
	if e != nil {
		fmt.Println(e)
		return nil, nil, e
	}

	fmt.Println(args)

	if strings.HasPrefix(sql, "select") {
		rows, e := stmt.Query(args)
		e = db.Close()
		return rows, nil, e
	} else {
		result, e := stmt.Exec(args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7],args[8],args[9])
		e = db.Close()
		return nil, result, e
	}

}

func dealError(err error) bool {
	if err != nil {
		fmt.Print(err)
		return true
	}
	return false
}

package main

import (
	"com/cheney/untils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Product struct {
	Id          int64  `db:"id"`
	ProductName string `db:"product_name"`
	ProductEn   string `db:"product_en"`
	RelationNum int32  `db:"relation_num"`
}

type Category struct {
	Id           int64  `db:"id"`
	Depth        int64  `db:"depth"`
	ParentId     int64  `db:"parent_id"`
	CategoryName string `db:"category_name"`
	CategoryEn   string `db:"category_en"`
	FullClass    string `db:"full_class"`
	FullClassEn  string `db:"full_class_en"`
}

type AttrKey struct {
	Id            int64  `db:"id"`
	AttributeName string `db:"attribute_name"`
	AttributeEn   string `db:"attribute_en"`
}

type AttrVal struct {
	Id               int64  `db:"id"`
	AttributeId      int64  `db:"attribute_id"`
	AttributeValue   string `db:"attribute_value"`
	AttributeValueEn string `db:"attribute_value_en"`
}

var db *sqlx.DB

var proEnIndex = 0
var proIndex = 1

var cateEnIndex = 4
var cateIndex = 5

var attrKeyEnIndex = 0
var attrKeyIndex = 1
var attrValEnIndex = 2
var attrValIndex = 3

var productCache = make(map[string]Product)
var categoryCache = make(map[string]Category)
var attrKeyCache = make(map[string]AttrKey)
var attrValCache = make(map[string]AttrVal)

var proUpdater *sqlx.NamedStmt
var pcrInserter *sql.Stmt
var cavrInserter *sql.Stmt
var proInserter *sqlx.NamedStmt
var cateInserter *sqlx.NamedStmt
var attrKeyInserter *sqlx.NamedStmt
var attrValInserter *sqlx.NamedStmt

func init() {
	database, err := sqlx.Open("mysql", "sl_rec_data:xx@tcp(xx:xx)/sl_rec_data")
	if err != nil {
		fmt.Println("open mysql error:", err)
		return
	}
	db = database
	proUpdater, _ = db.PrepareNamed("update product1 set product_name = :product_name,product_en = :product_en where id = :id")
	pcrInserter, _ = db.Prepare("insert into product_category_relation1(product_id,category_id) values(?,?)")
	cavrInserter, _ = db.Prepare("insert into category_attribute_value_relation1(category_id,attribute_id,attribute_value_id) values(?,?,?)")
	proInserter, _ = db.PrepareNamed("insert into product1(product_name,product_en,relation_num) values(:product_name,:product_en,:relation_num)")
	cateInserter, _ = db.PrepareNamed("insert into category1(depth,parent_id,category_name,category_en,full_class,full_class_en) values(:depth,:parent_id,:category_name,:category_en,:full_class,:full_class_en)")
	attrKeyInserter, _ = db.PrepareNamed("insert into attribute1(attribute_name,attribute_en) values(:attribute_name,:attribute_en)")
	attrValInserter, _ = db.PrepareNamed("insert into attribute_value1(attribute_id,attribute_value,attribute_value_en) values(:attribute_id,:attribute_value,:attribute_value_en)")

	r, _ := db.Queryx("select id,full_class_en,full_class from category1")
	for r.Next() {
		category := Category{}
		r.StructScan(&category)
		categoryCache[category.FullClassEn] = category
	}
	fmt.Println("finish category", categoryCache)
	r, _ = db.Queryx("select id,attribute_en from attribute1")
	for r.Next() {
		attrKey := AttrKey{}
		r.StructScan(&attrKey)
		attrKeyCache[attrKey.AttributeEn] = attrKey
	}
	fmt.Println("finish attr attr key")
	r, _ = db.Queryx("select id,attribute_value_en from attribute_value1")
	for r.Next() {
		attrVal := AttrVal{}
		r.StructScan(&attrVal)
		attrValCache[attrVal.AttributeValueEn] = attrVal
	}
	fmt.Println("finish attr val")
}

func main() {
	import2()
}

func import4() {
	reader, err := untils.ReadExcel("D:\\work\\df_goods_name_new_white.xlsx")
	var cellArray = []string{"A", "B", "C", "D", "E", "F"}
	startRow := 2
	cell := reader.ReadCell(cellArray, startRow, 5045)
	if err != nil {
		goto _ERR
	}
	for _, row := range cell {
		fmt.Printf("执行第%v行\r", startRow)
		startRow++
		proEnStr := row[proEnIndex].(string)
		proStr := row[proIndex].(string)
		key := proEnStr
		if key == "NULL" || key == "" {
			key = proStr
		}
		product := Product{
			RelationNum: 30,
			Id:          (int64)(startRow - 2),
		}
		if !(proEnStr == "NULL" || proEnStr == "") {
			product.ProductEn = proEnStr
		}
		if !(proStr == "NULL" || proStr == "") {
			product.ProductName = proStr
		}
		proUpdater.Exec(product)
	}

	defer db.Close()
	return
_ERR:
	fmt.Println(err)
}

func import3() {
	reader, err := untils.ReadExcel("D:\\work\\df_goods_name_new_white.xlsx")
	var cellArray = []string{"A", "B", "C", "D", "E", "F"}
	startRow := 2
	cell := reader.ReadCell(cellArray, startRow, 5045)
	if err != nil {
		goto _ERR
	}
	for _, row := range cell {
		fmt.Printf("执行第%v行\r", startRow)
		startRow++
		product, isNew := insertProduct(row)
		category := insertCategory(row)
		if category == nil {
			continue
		}
		categoryId := category.Id
		if categoryId == 0 {
			fmt.Println(category)
		}
		if isNew {
			pcrInserter.Exec(product.Id, categoryId)
		}
	}

	defer db.Close()
	return
_ERR:
	fmt.Println(err)
}

func import2() {
	reader, err := untils.ReadExcel("/Users/chenyi/yy/product-knowledge/df_attr_new_to_chenyi_white_list.xlsx")
	var cellArray = []string{"A", "B", "C", "D", "E", "F"}
	startRow := 2231
	cell := reader.ReadCell(cellArray, startRow, 4519)
	if err != nil {
		goto _ERR
	}
	for _, row := range cell {
		fmt.Printf("执行第%v行：", startRow)
		startRow++
		categoryId := insertCategory(row).Id
		attr := insertAttr(row)
		if attr != nil {
			vals := insertAttrVal(attr.Id, row)
			if vals != nil && len(vals) > 0 {
				for _, val := range vals {
					cavrInserter.Exec(categoryId, attr.Id, val.Id)
				}
			}
		}
	}

	defer db.Close()
	return
_ERR:
	fmt.Println(err)
}

func import1() {
	reader, err := untils.ReadExcel("/Users/chenyi/yy/product-knowledge/good_name_attr_20201230_limit5000.xlsx")
	var cellArray = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	startRow := 4334
	cell := reader.ReadCell(cellArray, startRow, 19170)
	if err != nil {
		goto _ERR
	}
	for _, row := range cell {
		fmt.Printf("执行第%v行\n", startRow)
		startRow++
		product, isNew := insertProduct(row)
		categoryId := insertCategory(row).Id
		if isNew {
			pcrInserter.Exec(product.Id, categoryId)
		}
		attr := insertAttr(row)
		if attr != nil {
			vals := insertAttrVal(attr.Id, row)
			if vals != nil && len(vals) > 0 {
				for _, val := range vals {
					cavrInserter.Exec(categoryId, attr.Id, val.Id)
				}
			}
		}
	}

	defer db.Close()
	return
_ERR:
	fmt.Println(err)
}

func insertProduct(row []interface{}) (Product, bool) {
	proEnStr := row[proEnIndex].(string)
	proStr := row[proIndex].(string)
	key := proEnStr
	if key == "NULL" || key == "" {
		key = proStr
	}
	if product, ok := productCache[key]; ok {
		return product, false
	}
	product := Product{
		RelationNum: 30,
	}
	if proEnStr != "NULL" && proEnStr != "" {
		product.ProductEn = proEnStr
	} else {
		product.ProductEn = ""
	}
	if proStr != "NULL" && proStr != "" {
		product.ProductName = proStr
	} else {
		product.ProductName = ""
	}
	result, err := proInserter.Exec(product)
	if err != nil {
		log.Error("商品插入异常", err)
	}
	product.Id, _ = result.LastInsertId()
	productCache[product.ProductName] = product
	fmt.Println(product)
	return product, true
}

func insertCategory(row []interface{}) *Category {
	cateEn := row[cateEnIndex].(string)
	if cateEn == "NULL" || cateEn == "" {
		return nil
	}
	cateEnArray := strings.Split(cateEn, ">")
	cacheKey := "," + strings.Join(cateEnArray, ",") + ","
	if category, ok := categoryCache[cacheKey]; ok {
		return &category
	}
	var result Category
	cate := row[cateIndex].(string)
	cateArray := strings.Split(cate, ">")
	var parentId int64
	fullClass := ","
	fullClassEn := ","
	for i, v := range cateEnArray {
		fullClassEn += v + ","
		if category, ok := categoryCache[fullClassEn]; !ok {
			result = Category{
				Depth:       int64(i + 1),
				ParentId:    parentId,
				CategoryEn:  v,
				FullClassEn: fullClassEn,
			}
			if i > len(cateArray)-1 {
				if i != len(cateEnArray)-1 {
					panic("数据缺失")
				} else {
					result.CategoryName = cateArray[len(cateArray)-1]
				}
			} else {
				result.CategoryName = cateArray[i]
			}
			fullClass += result.CategoryName + ","
			result.FullClass = fullClass
			r, err := cateInserter.Exec(result)
			if err != nil {
				log.Error("分类插入异常", err)
			}
			fmt.Println("new category:", result)
			result.Id, _ = r.LastInsertId()
			parentId = result.Id
			categoryCache[fullClassEn] = result
		} else {
			parentId = category.Id
			fullClass = category.FullClass
		}
	}

	return &result
}

func insertAttr(row []interface{}) *AttrKey {
	keyStr := row[attrKeyEnIndex].(string)
	if keyStr == "NULL" || keyStr == "" {
		return nil
	}
	if attrKey, ok := attrKeyCache[keyStr]; ok {
		return &attrKey
	} else {
		attrKey = AttrKey{
			AttributeName: row[attrKeyIndex].(string),
			AttributeEn:   keyStr,
		}
		r, err := attrKeyInserter.Exec(attrKey)
		if err != nil {
			log.Error("分类插入异常", err)
		}
		attrKey.Id, _ = r.LastInsertId()
		attrKeyCache[keyStr] = attrKey
		return &attrKey
	}
}

func insertAttrVal(keyId int64, row []interface{}) []AttrVal {
	attrValStr := row[attrValIndex].(string)
	var attrValArray []string
	if attrValStr != "NULL" && attrValStr != "" {
		attrValArray = strings.Split(attrValStr, "，")
	}
	attrValEnStr := row[attrValEnIndex].(string)
	if attrValEnStr == "NULL" || attrValEnStr == "" {
		return nil
	}
	results := make([]AttrVal, 0)
	attrValEnArray := strings.Split(attrValEnStr, "，")
	fmt.Println(len(attrValEnArray))
	for i, v := range attrValEnArray {
		if _, ok := attrValCache[v]; !ok {
			attrVal := AttrVal{
				AttributeId:      keyId,
				AttributeValueEn: v,
			}
			if attrValArray != nil {
				attrVal.AttributeValue = attrValArray[i]
			}
			r, err := attrValInserter.Exec(attrVal)
			if err != nil {
				log.Error("插入属性值异常", err)
			}
			attrVal.Id, _ = r.LastInsertId()
			attrValCache[v] = attrVal
			results = append(results, attrVal)
		}
	}
	return results
}

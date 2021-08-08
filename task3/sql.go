package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Clothings struct {
	id    string
	size  string
	price int
	style string
}

type Store struct {
	id       string
	capacity int
}

type Suppliers struct {
	id   string
	name string
}

type Condition struct {
	id          string
	supplier_id int
	quality     string
}

var db *sql.DB

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Other wrong:", err)
		}
	}()
	dsn := "root:qazpl.123456@tcp(127.0.0.1:3306)/now"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open:", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping:", err.Error())
		return
	}
	fmt.Println("连接数据库成功!")
}

func QueryClothingsByPS(price int, size string) ([]Clothings, error) {
	str := "select * from `clothings` where price < ? and size = ?;"
	rows, err := db.Query(str, price, size)
	if err != nil {
		return nil, err
	}
	clothings := make([]Clothings, 0)
	for rows.Next() {
		var tmp Clothings
		_ = rows.Scan(&tmp.id, &tmp.size, &tmp.price, &tmp.style)
		clothings = append(clothings, tmp)
	}
	return clothings, nil
}

func QueryClothingsById(id string) ([]Clothings, error) {
	str := "select * from `clothings` where id like ?;"
	stmt, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id + "%")
	if err != nil {
		return nil, err
	}
	clothings := make([]Clothings, 0)
	for rows.Next() {
		var tmp Clothings
		_ = rows.Scan(&tmp.id, &tmp.size, &tmp.price, &tmp.style)
		clothings = append(clothings, tmp)
	}
	return clothings, nil
}

func QueryClothingsByQ(style string) (int, error) {
	str := "select count(id) from `clothings` where style = ?"
	stmt, err := db.Prepare(str)
	var count int
	if err != nil {
		return count, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(style).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func QueryMaxStore() (Store, error) {
	str := "select * from `store` order by capacity desc limit 1;"
	tmp := Store{}
	err := db.QueryRow(str).Scan(&tmp.id, &tmp.capacity)
	if err != nil {
		return tmp, err
	}
	return tmp, nil
}

func QuerySuppliersByQ(quality string) ([]Suppliers, error) {
	str := "select * from `suppliers` where id = (select supplier_id from `condition` where quality = ?);"
	stmt, err := db.Prepare(str)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(quality)
	if err != nil {
		return nil, err
	}
	res := make([]Suppliers, 0)
	for rows.Next() {
		var tmp Suppliers
		err = rows.Scan(&tmp.id, &tmp.name)
		res = append(res, tmp)
	}
	return res, nil
}

func DealResult(msg interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}

func UpdatePriceByS(size string) error {
	str := "Update `clothings` set price = price*1.1 where size = ?"
	stmt, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(size)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConditionByQ(quality string) error {
	str := "Delete from `condition` where quality = ?"
	stmt, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(quality)
	if err != nil {
		return err
	}
	return nil
}

func InsertStore(store *Store) error {
	str := "Insert into `store` value(?,?);"
	stmt, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(store.id, store.capacity)
	if err != nil {
		return err
	}
	return nil
}

func InsertClothing(clothing *Clothings) error {
	str := "Insert into `clothings` value(?,?,?,?);"
	stmt, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(clothing.id, clothing.size, clothing.price, clothing.style)
	if err != nil {
		return err
	}
	return nil
}

func InsertSupplier(supplier *Suppliers) error {
	str := "Insert into `suppliers` value(?,?);"
	stmt, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(supplier.id, supplier.name)
	if err != nil {
		return err
	}
	return nil
}

func InsertCondition(condition *Condition) error {
	str := "Insert into `condition` value(?,?,?);"
	stmt, err := db.Prepare(str)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(condition.id, condition.supplier_id, condition.quality)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	defer db.Close()
	var (
		res interface{}
		err error
	)
	clothing := &Clothings{id: "ACL001", size: "S", price: 50, style: "A"}
	store := &Store{id: "AST001", capacity: 30}
	supplier := &Suppliers{id: "1", name: "Black"}
	condition := &Condition{id: "ACO001", supplier_id: 1, quality: "不合格"}
	err = InsertStore(store)
	DealResult(nil, err)
	err = InsertClothing(clothing)
	DealResult(nil, err)
	err = InsertSupplier(supplier)
	DealResult(nil, err)
	err = InsertCondition(condition)
	DealResult(nil, err)
	res, err = QueryClothingsByPS(100, "S")
	DealResult(res, err)
	res, err = QueryMaxStore()
	DealResult(res, err)
	res, err = QueryClothingsByQ("A")
	DealResult(res, err)
	res, err = QueryClothingsById("A")
	DealResult(res, err)
	res, err = QuerySuppliersByQ("不合格")
	DealResult(res, err)
	err = UpdatePriceByS("S")
	DealResult(nil, err)
	err = DeleteConditionByQ("不合格")
	DealResult(nil, err)
}

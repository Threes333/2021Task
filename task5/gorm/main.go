package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Clothing struct {
	Id    string `gorm:"primary_key"`
	Size  string `gorm:"type:varchar(20)"`
	Price int    `gorm:"type:int"`
	Style string `gorm:"type:varchar(20)"`
}

type Store struct {
	Id       string `gorm:"primary_key"`
	Capacity int    `gorm:"type:int"`
}

type Supplier struct {
	Id   string `gorm:"primary_key"`
	Name string `gorm:"type:varchar(20)"`
}

type Condition struct {
	Id          string `gorm:"primary_key"`
	Supplier_id int    `gorm:"primary_key"`
	Quality     string `gorm:"type:varchar(20)"`
}

var db *gorm.DB

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Other wrong:", err)
		}
	}()
	dsn := "root:qazpl.123456@tcp(127.0.0.1:3306)/next?charset=utf8mb4&parseTime=true&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect SqlDB failed")
	}
	err = db.AutoMigrate(Clothing{}, Supplier{}, Store{}, Condition{})
	fmt.Println(err)
	fmt.Println("连接数据库成功!")
}

func main() {
	clothing := &Clothing{Id: "ACL001", Size: "S", Price: 50, Style: "A"}
	store := &Store{Id: "AST001", Capacity: 30}
	supplier := &Supplier{Id: "1", Name: "Black"}
	condition := &Condition{Id: "ACO001", Supplier_id: 1, Quality: "不合格"}
	db.Create(clothing)
	db.Create(supplier)
	db.Create(store)
	db.Create(condition)
	clothings := make([]Clothing, 10)
	db.Debug().Find(&clothings, "price < ? and size = ?", 100, "S")
	fmt.Println(clothings)
	db.Debug().Order("capacity DESC").First(store)
	fmt.Println(store)
	var count int64
	db.Debug().Model(&Clothing{}).Select("id").Where("style = ?", "A").Count(&count)
	fmt.Println(count)
	db.Debug().Where("id like ?", "A%").Find(&clothings)
	fmt.Println(clothings)
	suppliers := make([]Supplier, 10)
	db.Raw("select * from `suppliers` where id = (select supplier_id from `conditions` where quality = ?)", "不合格").Find(&suppliers)
	fmt.Println(suppliers)
	db.Debug().Table("clothings").Where("size = ?", "S").Update("price", gorm.Expr("price * 1.1"))
	db.Debug().Where("quality = ?", "不合格").Delete(Condition{})
}

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Adiciona uma categoria ao produto
type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

func main() {
	//Foi alterado o DSN para atualizar o produto criado
	dsn := "root:root@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// A autoMigrate cria a tabela de Product
	db.AutoMigrate(&Product{}, Category{}, SerialNumber{})

	//create category
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	//create product
	db.Create(&Product{
		Name:       "Celular",
		Price:      5799.00,
		CategoryID: 1,
	})

	//create serial number
	db.Create(&SerialNumber{
		Number:    "23456",
		ProductID: 1,
	})

	//trazer todos os produtos
	///o Preload()vai instanciar a struct de Category e SerialNumber para conseguir mostrar no CLI
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}

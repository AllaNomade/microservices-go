package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
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
	db.AutoMigrate(&Product{}, Category{})

	//create category
	category := Category{Name: "Eletrônicos"}
	db.Create(&category)

	//create product
	db.Create(&Product{
		Name:       "Celular",
		Price:      5799.00,
		CategoryID: category.ID,
	})

	//trazer todos os produtos
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}

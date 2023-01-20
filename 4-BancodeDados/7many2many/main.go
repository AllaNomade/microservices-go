package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Adiciona uma categoria ao produto
type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

// // type SerialNumber struct {
// // 	ID        int `gorm:"primaryKey"`
// // 	Number    string
// // 	ProductID int
// // }

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
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
	// // category := Category{Name: "Cozinha"}
	// // db.Create(&category)

	// // category2 := Category{Name: "Eletronicos"}
	// // db.Create(&category2)

	//create product
	// // db.Create(&Product{
	// // 	Name:       "Panela",
	// // 	Price:      88.00,
	// // 	Categories: []Category{category, category2},
	// // })

	//trazer todos os produtos
	///o Preload()vai instanciar a struct de Category para conseguir mostrar no CLI
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}
}

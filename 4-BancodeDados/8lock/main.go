package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	//.Begin() inicia uma transaction
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletrônicos"
	tx.Debug().Save(&c)
	tx.Commit()
}

//select * from product where id = 1 FOR UPDATE

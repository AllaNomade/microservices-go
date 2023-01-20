package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
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
	db.AutoMigrate(&Product{})

	//criação de produto
	// // db.Create(&Product{
	// // 	Name:  "Carro",
	// // 	Price: 34.666,
	// // })

	//criação vários produtos
	// // products := []Product{
	// // 	{Name: "Moto", Price: 12.999},
	// // 	{Name: "Avião", Price: 99.999},
	// // 	{Name: "Bike", Price: 1.999},
	// // }
	// // db.Create(&products)

	//seleciona um produto
	// // var product Product
	// // db.First(&product, 3)
	// // defer fmt.Println(product)
	// // db.First(&product, "name = ?", "Bike")
	// // fmt.Println("algo aqui", product)

	//seleciona todos os produtos
	// // var products []Product
	// // db.Find(&products)
	// // for _, product := range products {
	// // 	fmt.Println(product)
	// // }

	//seleciona e instaura um limite total de produtos requisitados na função
	// // var products []Product
	// // db.Limit(2).Offset(2).Find(&products)
	// // for _, product := range products {
	// // 	fmt.Println(product)
	// // }

	//WHERE
	// // var products []Product
	// // db.Where("price >= ?", 99.999).Find(&products)
	// // for _, product := range products {
	// // 	fmt.Println(product)
	// // }

	//WHERE LIKE traz um filtro pela letra do nome
	// // var products []Product
	// // db.Where("name LIKE ?", "%e%").Find(&products)
	// // for _, product := range products {
	// // 	fmt.Println(product)
	// // }

	//Atualiza e deleta
	// // var p Product
	// // db.First(&p, 1)
	// // p.Name = "Notebook Acer I9"
	// // db.Save(&p)

	//Utilizando gorm.Model na struct, ao deletar nesta variável, há o Soft Delete,
	//Que deleta o produto, mas não exclui do banco de dados, mantendo o registro
	// da remoção.
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)
}

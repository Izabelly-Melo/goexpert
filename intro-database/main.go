package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Products struct {
	ID    string
	Name  string
	Price float64
}

func NewProducts(name string, price float64) *Products {
	return &Products{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProducts("Notebook", 3000)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 2500
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %s, Price: %.2f\n", p.Name, p.Price)

	fmt.Println("-----------------------")
	products, err := selectProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Product: %s, Price: %.2f\n", p.Name, p.Price)
	}
}

func insertProduct(db *sql.DB, p *Products) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)") // Prepara a query para inserção de um produto seguindo o padrão de segurança para evitar SQL Injection
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, p *Products) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Products, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	p := &Products{} // pontando o endereço de memória da struct para receber os dados retornados da query

	// queryrowcontext é utilizado para passar o contexto da aplicação, como tempo de execução, cancelamento, etc.
	// Ele é recomendado para evitar que a aplicação fique bloqueada em uma query que demora muito para ser executada.
	//err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)

	// queryrow é utilizado para retornar apenas um registro, e o scan é utilizado para mapear os campos retornados para a struct
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func selectProducts(db *sql.DB) ([]*Products, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []*Products{}

	for rows.Next() {
		p := &Products{}
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

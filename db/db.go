package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {

	conexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"0.0.0.0", 5051, "postgres", "qwerty", "postgres")
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Postgres com docker

// -- Postgres DB
// docker run -it --name postgres-db -e POSTGRES_PASSWORD=qwerty -d -p 5051:5432 postgres

// -- pgAdmin
// docker run -it --name pgadmin-web -e PGADMIN_DEFAULT_EMAIL=admin@admin.com -e PGADMIN_DEFAULT_PASSWORD=qwerty -d -p 5050:80  dpage/pgadmin4

// No resultado busque por networks e verifique o IP vinculado ao container criado.
// docker inspect postgres-db

// "IPAddress": "172.17.0.2",

// URL Browser:
// 0.0.0:5050/browser/

// Login / Senha psAdmin
// admin@admin.com
// qwerty

// Conexão com o Banco, dentro do psAdmin - Navegador
// Hostname: 172.17.0.2
// Port: 5432
// Database: postgres
// Username: postgres
// Password: qwerty

// Banco de dados - Golang - main.go
// nc -vz 0.0.0.0 5051
// Connection to 0.0.0.0 port 5051 [tcp/ita-agent] succeeded!

// conexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//         "0.0.0.0", 5051, "postgres", "qwerty", "postgres")

// create table produtos (
// 	id serial primary key,
// 	nome varchar,
// 	descricao varchar,
// 	preco decimal,
// 	quantidade integer
// )

// insert into produtos (nome, descricao, preco, quantidade) values
// ('Camiseta', 'Azul', 19, 10),
// ('Fone', 'Muito Bom', 99, 5);

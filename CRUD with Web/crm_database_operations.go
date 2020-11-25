package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "anhyeuhanoi"
	dbname   = "sanpham"
)

type SanPham struct{
	TenSanPham string
	Gia int
}

func getAll() []SanPham {
	var db *sql.DB
	db = connectGate()
	defer db.Close()
	rows, err := db.Query("select * from docs")

	if err != nil {
		log.Fatal("@getAll: ", err)
	}

	listSP := []SanPham{}
	for rows.Next(){
		sp := SanPham{}
		rows.Scan(&sp.TenSanPham, &sp.Gia)
		listSP = append(listSP, sp)

		if error := rows.Err(); error != nil{ 
			panic(error)
		}
	}
	fmt.Println(listSP)
	return listSP
}

func getByID (id string) SanPham {
	var db *sql.DB
	db = connectGate()
	defer db.Close()
	rows, err := db.Query("select * from docs where tensanpham=$1", id)

	if err != nil {
		log.Fatal("@getByID: ", err)
	}

	sp := SanPham{}
	for rows.Next(){
		var TenSanPham string
		var Gia int
		rows.Scan(&TenSanPham, &Gia)

		if error := rows.Err(); error != nil{ 
			panic(error)
		}
		sp.TenSanPham = TenSanPham
		sp.Gia = Gia
	}
	return sp
}

func create(sp SanPham) {
	fmt.Println(sp)
	var db *sql.DB
	db = connectGate()
	defer db.Close()
	query, err := db.Prepare("insert into docs values($1, $2)")
	if err != nil{
		log.Fatal("@create: ", err)
	}	
	query.Exec(sp.TenSanPham, sp.Gia)

	fmt.Println("@create success")
}

func delete(TenSanPham string){
	var db *sql.DB
	db = connectGate()
	defer db.Close()
	query, err := db.Prepare("delete from docs where TenSanPham = $1")
	if err != nil{
		log.Fatal("@delete: ", err)
	}
	query.Exec(TenSanPham)

	fmt.Println("@delete success")
}

func edit(sp SanPham){
	var db *sql.DB
	db = connectGate()
	defer db.Close()
	query, err := db.Prepare("update docs set gia = $1 where tensanpham = $2")
	if err != nil{
		log.Fatal("@edit: ", err)
	}
	query.Exec(sp.Gia, sp.TenSanPham)

	fmt.Println("@edit success")
}

func connectGate() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  return
}


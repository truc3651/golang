package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "anhyeuhanoi"
  dbname   = "sanpham"
)

type sanpham struct{
	TenSanPham string
	Gia int
}

func main() {
	for{
		// show 
		fmt.Println("*****CRUD*****")
		fmt.Println("1. select all")
		fmt.Println("2. create record")
		fmt.Println("3. edit record")
		fmt.Println("4. delete record")
		fmt.Println("**************")
		fmt.Print("choose: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := scanner.Text()

		// 
		switch option{
			case "1":
				getAll()
			case "2":
				create()
			case "3":
				edit()
			default:
				delete()
		}

		// wanna continue
		fmt.Println()
		fmt.Print("Continue? y/n: ")
		scanner.Scan()
		if strings.Compare(strings.ToLower(scanner.Text()), "n") == 0 {
			break
		}
		time.Sleep(time.Second)
	}
}

func getAll()  {
	db := connectGate()
	defer db.Close()
	rs, err := db.Query("select * from docs")

	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	listSP := []sanpham{}
	for rs.Next(){
		sp := sanpham{}
		rs.Scan(&sp.TenSanPham, &sp.Gia)
		listSP = append(listSP, sp)

		if error := rs.Err(); error != nil{ 
			panic(error)
		}
	}
	fmt.Println(listSP)
}

func create(){
	fmt.Println("@create: Type name and price")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TenSanPham := scanner.Text()
	scanner.Scan()
	Gia, _ := strconv.Atoi(scanner.Text())

	db := connectGate()
	defer db.Close()
	query, err := db.Prepare("insert into docs values($1, $2)")
	if err != nil{
		log.Fatal("@create: ", err)
	}	
	query.Exec(TenSanPham, Gia)

	fmt.Println("@create success")
}

func delete(){
	fmt.Println("@delete: Type name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TenSanPham := scanner.Text()

	db := connectGate()
	defer db.Close()
	query, err := db.Prepare("delete from docs where tensanpham = $1")
	if err != nil{
		log.Fatal("@delete: ", err)
	}
	query.Exec(TenSanPham)

	fmt.Println("@delete success")
}

func edit(){
	fmt.Println("@edit: Type name and price")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TenSanPham := scanner.Text()
	scanner.Scan()
	Gia, _ := strconv.Atoi(scanner.Text())

	db := connectGate()
	defer db.Close()
	query, err := db.Prepare("update docs set gia = $1 where tensanpham = $2")
	if err != nil{
		log.Fatal("@edit: ", err)
	}
	query.Exec(Gia, TenSanPham)

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

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var template_html = template.Must(template.ParseGlob("templates/*"))

func Home(writer http.ResponseWriter, request *http.Request) {
	listSP := getAll()
	template_html.ExecuteTemplate(writer, "Home", listSP)
}

// navigate to the create page
func Create(writer http.ResponseWriter, request *http.Request) {
	template_html.ExecuteTemplate(writer, "Create", nil)
}

// now you enter the create button
func Insert(writer http.ResponseWriter, request *http.Request) {
	var sp SanPham
	sp.TenSanPham = request.FormValue("TenSanPham")
	var Gia int
	fmt.Sscanf(request.FormValue("Gia"), "%d", &Gia)
	sp.Gia = Gia

	create(sp)
	var listSP []SanPham
	listSP = getAll()
	template_html.ExecuteTemplate(writer, "Home", listSP)
}

// show all the current activated record data to the edit page
func Update(writer http.ResponseWriter, request *http.Request) {
	sp := getByID(request.FormValue("id"))
	fmt.Println(sp)
	template_html.ExecuteTemplate(writer, "Update", sp)
}

// enter edit record button
func Alter(writer http.ResponseWriter, request *http.Request) {
	var sp SanPham
	sp.TenSanPham = request.FormValue("TenSanPham")
	var Gia int
	fmt.Sscanf(request.FormValue("Gia"), "%d", &Gia)
	sp.Gia = Gia
	edit(sp)

	var listSP []SanPham
	listSP = getAll()
	template_html.ExecuteTemplate(writer, "Home", listSP)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	TenSanPham := request.FormValue("id")
	delete(TenSanPham)
	var listSP []SanPham
	listSP = getAll()
	template_html.ExecuteTemplate(writer, "Home", listSP)

}

func View(writer http.ResponseWriter, request *http.Request) {
	TenSanPham := request.FormValue("id")
	sp := getByID(TenSanPham)
	listSP := []SanPham{sp}
	
	template_html.ExecuteTemplate(writer, "View", listSP)
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}

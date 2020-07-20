package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
	router.HandleFunc("/",IndexPage)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", router)
}
func IndexPage(w http.ResponseWriter, r *http.Request)  {
	tmpl,_:=template.ParseFiles("wwwroot/index.html")
	tmpl.Execute(w,"")
}
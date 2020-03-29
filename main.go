package main

import (
	"net/http"

	"github.com/chaitanyamaili/go-webservice/controllers"
)

func main() {
	// u := models.User{
	// 	ID:        1,
	// 	FirstName: "Chaitanya",
	// 	LastName:  "Maili",
	// }
	// fmt.Println(u)
	controllers.RegisterControllers()
	http.ListenAndServe(`:3000`, nil)
}

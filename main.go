package main

import (
	"fmt"

	"github.com/chaitanyamaili/go-webservice/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Chaitanya",
		LastName:  "Maili",
	}
	fmt.Println(u)
}

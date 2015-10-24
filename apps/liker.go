package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("starting service")

	http.ListenAndServe(":8080", nil)

}

package handler

import "fmt"

type Handler struct {
	// add some fields
}

func (*Handler) initialize() {
	fmt.Println("Initialization")
	var a, b int
	fmt.Scan(&a, &b)
	c := a + b
	fmt.Print(c)
}

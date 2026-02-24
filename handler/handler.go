package handler

import "fmt"

type Handler struct {
	// add some fields
}

func (*Handler) initialize() {
	fmt.Println("Fixing handler")
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	d := a + b + c
	fmt.Print(d)
}

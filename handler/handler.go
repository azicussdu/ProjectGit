package handler

import "fmt"

type Handler struct {
	// add some fields
}

func (*Handler) initialize() {
	fmt.Println("Hey hey in main")
	var a, b int
	fmt.Scan(&a, &b)
	c := a + b
	fmt.Print(c)
}

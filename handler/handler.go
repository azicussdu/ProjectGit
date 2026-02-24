package handler

import "fmt"

type Handler struct {
	// add some fields
}

func (*Handler) initialize() {
	fmt.Println("Hey hey in main")
	fmt.Println("I fixed it in main")
}

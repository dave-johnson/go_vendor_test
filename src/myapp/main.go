package main

import (
	"fmt"

	ms "./modules/mystuff"
	gm "github.com/deprepo"
)

func main() {
	fmt.Println("start main")
	gm.Stuff("main")
	ms.Stuff()
	fmt.Println("end main")
}

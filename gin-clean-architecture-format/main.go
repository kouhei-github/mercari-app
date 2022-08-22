package main

import (
	"fmt"
	"kouhei-github/sample-gin/route"
)

func main() {
	router := route.GetRouter()
	fmt.Println(router.Run(":8080"))
}

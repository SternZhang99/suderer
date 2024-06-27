package main

import (
	"fmt"
	"sunderer/routers"
)

func main() {
	r := routers.Router()
	err := r.Run(":8081")
	if err != nil {
		fmt.Println("running error:", err)
		return
	}
}

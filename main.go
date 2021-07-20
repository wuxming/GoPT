package main

import (
	"gin1/router"
)

func main() {
	r := router.SetupRouter()
	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}

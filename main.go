package main

import (
	"gin1/initialization"
	"gin1/router"
)

func init() {
	initialization.MongoInit()
}
func main() {
	r := router.SetupRouter()
	err := r.Run(":8000")
	if err != nil {
		panic(err)
	}
}

package main

import "gin1/routers"

func main() {
	r := routers.SetupRouter()
	err:=r.Run(":8000")
	if err!=nil {
		panic(err)
	}
}

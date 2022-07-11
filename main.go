package main

import "gin-JWT/route"

func main() {
	r := route.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		return 
	}
}

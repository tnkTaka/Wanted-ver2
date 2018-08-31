package main

import "github.com/Wanted-release/router"

func main() {
	r := router.GetRouter()
	r.Run(":5000")
}

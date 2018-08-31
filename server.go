package main

import "github.com/Wanted-ver2/router"

func main() {
	r := router.GetRouter()
	r.Run(":80")
}

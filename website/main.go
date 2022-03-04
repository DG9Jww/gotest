package main

import (
	"website/route"
)

func main() {
	r := route.Setup()
	r.Run(":50001")
}

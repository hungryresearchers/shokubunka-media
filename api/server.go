package main

import "api/infrastructure"

func main() {
	infrastructure.Router.Run(":3000")
}

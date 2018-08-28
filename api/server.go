package main

import "shokubunka-media/api/src/infrastructure"

func main() {
	infrastructure.Router.Run(":3000")
}
